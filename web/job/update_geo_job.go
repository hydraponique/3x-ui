package job

import (
	"strings"
	"time"

	"github.com/mhsanaei/3x-ui/v2/logger"
	"github.com/mhsanaei/3x-ui/v2/web/service"
)

// sourceFilenames maps an abstract source key (as stored in the
// geoAutoUpdateSources setting) to the concrete .dat file names managed
// by ServerService.UpdateGeofile. Keep this in sync with x-ui.sh and
// service/server.go.
var sourceFilenames = map[string][]string{
	"main":   {"geoip.dat", "geosite.dat"},
	"IR":     {"geoip_IR.dat", "geosite_IR.dat"},
	"RU":     {"geoip_RU.dat", "geosite_RU.dat"},
	"ROSCOM": {"geoip_ROSCOM.dat", "geosite_ROSCOM.dat"},
}

// UpdateGeoJob refreshes the geo data files configured by the admin.
// Scheduling lives in web.go; the job re-reads all settings on every
// tick so toggles and source-list changes apply without a restart.
type UpdateGeoJob struct {
	serverService  service.ServerService
	settingService service.SettingService
}

// NewUpdateGeoJob creates a geo auto-update job.
func NewUpdateGeoJob() *UpdateGeoJob {
	return &UpdateGeoJob{}
}

// Run is the cron entrypoint.
func (j *UpdateGeoJob) Run() {
	enabled, err := j.settingService.GetGeoAutoUpdate()
	if err != nil {
		logger.Warning("UpdateGeoJob: failed to read geoAutoUpdate flag:", err)
		return
	}
	if !enabled {
		return
	}

	sources, err := j.settingService.GetGeoAutoUpdateSources()
	if err != nil {
		logger.Warning("UpdateGeoJob: failed to read geoAutoUpdateSources:", err)
		return
	}

	// Empty list ⇒ refresh every file in the allowlist.
	var files []string
	if len(sources) == 0 {
		files = nil // signals "update all" below
	} else {
		for _, src := range sources {
			names, ok := sourceFilenames[src]
			if !ok {
				logger.Warningf("UpdateGeoJob: unknown source key %q — skipping", src)
				continue
			}
			files = append(files, names...)
		}
	}

	// Batch update so Xray is restarted at most once per tick, regardless
	// of how many sources the admin selected.
	var errs []string
	if err := j.serverService.UpdateGeofiles(files); err != nil {
		errs = append(errs, err.Error())
	}

	now := time.Now().Unix()
	status := "ok"
	if len(errs) > 0 {
		status = "error: " + strings.Join(errs, "; ")
		logger.Warning("UpdateGeoJob:", status)
	} else {
		logger.Info("UpdateGeoJob: all configured geofiles refreshed")
	}
	if err := j.settingService.SetGeoLastAutoUpdateAt(now); err != nil {
		logger.Warning("UpdateGeoJob: failed to persist last-run timestamp:", err)
	}
	if err := j.settingService.SetGeoLastAutoUpdateStatus(status); err != nil {
		logger.Warning("UpdateGeoJob: failed to persist last-run status:", err)
	}
}
