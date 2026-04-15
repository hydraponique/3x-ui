package sub

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
	"time"
)

// RoscomVPN routing sources served from hydraponique/roscomvpn-routing (HAPP profile).
// Each .DEEPLINK file contains a single-line happ://routing/onadd/<base64> value that
// is placed into the Happ `Routing` response header.
const (
	RoscomVPNSourceDefault   = "default"
	RoscomVPNSourceJsonSub   = "jsonsub"
	RoscomVPNSourceWhitelist = "whitelist"
	RoscomVPNSourceCustom    = "custom"

	roscomvpnCacheTTL       = 10 * time.Minute
	roscomvpnNegativeCache  = 30 * time.Second // back-off after a fetch failure to avoid blocking sub responses
	roscomvpnHTTPTimeout    = 4 * time.Second
	roscomvpnMaxBodyBytes   = 1 << 20 // 1 MiB cap for .DEEPLINK content
)

var roscomvpnSourceURLs = map[string]string{
	RoscomVPNSourceDefault:   "https://raw.githubusercontent.com/hydraponique/roscomvpn-routing/main/HAPP/DEFAULT.DEEPLINK",
	RoscomVPNSourceJsonSub:   "https://raw.githubusercontent.com/hydraponique/roscomvpn-routing/main/HAPP/JSONSUB.DEEPLINK",
	RoscomVPNSourceWhitelist: "https://raw.githubusercontent.com/hydraponique/roscomvpn-routing/main/HAPP/WHITELIST.DEEPLINK",
}

type roscomvpnCacheEntry struct {
	value     string
	fetchedAt time.Time
	lastFail  time.Time // zero if the last attempt succeeded
}

var (
	roscomvpnMu     sync.RWMutex
	roscomvpnCache  = map[string]roscomvpnCacheEntry{}
	roscomvpnClient = &http.Client{Timeout: roscomvpnHTTPTimeout}

	// Per-source fetch lock — serializes refreshes so simultaneous subscription
	// requests coalesce into a single outbound HTTP call.
	roscomvpnFetchLocks sync.Map // map[string]*sync.Mutex
)

func roscomvpnLockFor(src string) *sync.Mutex {
	if m, ok := roscomvpnFetchLocks.Load(src); ok {
		return m.(*sync.Mutex)
	}
	m, _ := roscomvpnFetchLocks.LoadOrStore(src, &sync.Mutex{})
	return m.(*sync.Mutex)
}

// ResolveRoutingRules returns the value to put in the "Routing" response header
// based on admin-selected source. For "custom" (or unknown value) it returns the
// supplied custom string as-is. For known RoscomVPN sources it fetches the
// matching .DEEPLINK content from GitHub with a TTL cache; on fetch failure it
// falls back to the last known good value, or to the custom string if the cache
// is cold. A short negative-cache window prevents blocking every subscription
// response during a routing.help outage.
func ResolveRoutingRules(source, custom string) string {
	src := strings.ToLower(strings.TrimSpace(source))
	if src == "" {
		src = RoscomVPNSourceDefault
	}
	if src == RoscomVPNSourceCustom {
		return custom
	}
	url, ok := roscomvpnSourceURLs[src]
	if !ok {
		return custom
	}

	// Fast path: warm cache entry still within TTL.
	roscomvpnMu.RLock()
	entry, hit := roscomvpnCache[src]
	roscomvpnMu.RUnlock()
	if hit && time.Since(entry.fetchedAt) < roscomvpnCacheTTL {
		return entry.value
	}

	// Negative cache: a recent failure — don't try again yet, serve stale.
	if hit && !entry.lastFail.IsZero() && time.Since(entry.lastFail) < roscomvpnNegativeCache {
		if entry.value != "" {
			return entry.value
		}
		return custom
	}

	// Singleflight: one concurrent fetch per source.
	mu := roscomvpnLockFor(src)
	mu.Lock()
	defer mu.Unlock()

	// Re-check under the lock — another goroutine may have refreshed the entry.
	roscomvpnMu.RLock()
	entry, hit = roscomvpnCache[src]
	roscomvpnMu.RUnlock()
	if hit && time.Since(entry.fetchedAt) < roscomvpnCacheTTL {
		return entry.value
	}

	if v, err := fetchRoscomVPNDeepLink(url); err == nil {
		roscomvpnMu.Lock()
		roscomvpnCache[src] = roscomvpnCacheEntry{value: v, fetchedAt: time.Now()}
		roscomvpnMu.Unlock()
		return v
	}

	// Record the failure so subsequent requests take the negative-cache fast path.
	roscomvpnMu.Lock()
	prev := roscomvpnCache[src]
	prev.lastFail = time.Now()
	roscomvpnCache[src] = prev
	roscomvpnMu.Unlock()

	if hit && entry.value != "" {
		return entry.value
	}
	return custom
}

func fetchRoscomVPNDeepLink(url string) (string, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("Accept", "text/plain")

	resp, err := roscomvpnClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return "", fmt.Errorf("roscomvpn deeplink fetch failed: HTTP %d", resp.StatusCode)
	}

	body, err := io.ReadAll(io.LimitReader(resp.Body, roscomvpnMaxBodyBytes))
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(body)), nil
}
