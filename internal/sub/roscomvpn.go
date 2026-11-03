package sub

import "strings"

// RoscomVPN routing sources. Every named source maps to a permanent HTTPS URL
// of a Happ .DEEPLINK routing profile published in the roscomvpn-routing
// repository. The resolved URL is handed to the regular remote-routing
// resolver (see remote_routing.go), which fetches, caches and persists it —
// nothing here performs network I/O.
const (
	RoscomVPNSourceDefault   = "default"
	RoscomVPNSourceJsonSub   = "jsonsub"
	RoscomVPNSourceWhitelist = "whitelist"
	RoscomVPNSourceCustom    = "custom"
)

var roscomVPNSourceURLs = map[string]string{
	RoscomVPNSourceDefault:   "https://raw.githubusercontent.com/hydraponique/roscomvpn-routing/main/HAPP/DEFAULT.DEEPLINK",
	RoscomVPNSourceJsonSub:   "https://raw.githubusercontent.com/hydraponique/roscomvpn-routing/main/HAPP/JSONSUB.DEEPLINK",
	RoscomVPNSourceWhitelist: "https://raw.githubusercontent.com/hydraponique/roscomvpn-routing/main/HAPP/WHITELIST.DEEPLINK",
}

// ResolveRoscomVPNRouting maps an admin-selected routing source to the
// matching .DEEPLINK URL. The "custom" source (and any unknown or empty
// value) falls back to the admin-provided rules string, which is the
// upstream behavior: an inline happ:// deeplink or a permanent HTTPS URL
// of the admin's choice.
func ResolveRoscomVPNRouting(source, custom string) string {
	src := strings.ToLower(strings.TrimSpace(source))
	if src == "" || src == RoscomVPNSourceCustom {
		return custom
	}
	if url, ok := roscomVPNSourceURLs[src]; ok {
		return url
	}
	return custom
}
