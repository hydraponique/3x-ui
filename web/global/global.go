// Package global provides global variables and interfaces for accessing web and subscription servers.
package global

import (
	"context"
	_ "unsafe"

	"github.com/robfig/cron/v3"
)

var (
	webServer WebServer
	subServer SubServer
)

// WebServer interface defines methods for accessing the web server instance.
type WebServer interface {
	GetCron() *cron.Cron     // Get the cron scheduler
	GetCtx() context.Context // Get the server context
	GetWSHub() any           // Get the WebSocket hub (using any to avoid circular dependency)
}

// SubServer interface defines methods for accessing the subscription server instance.
type SubServer interface {
	GetCtx() context.Context // Get the server context
}

// SetWebServer sets the global web server instance.
func SetWebServer(s WebServer) {
	webServer = s
}

// GetWebServer returns the global web server instance.
func GetWebServer() WebServer {
	return webServer
}

// SetSubServer sets the global subscription server instance.
func SetSubServer(s SubServer) {
	subServer = s
}

// GetSubServer returns the global subscription server instance.
func GetSubServer() SubServer {
	return subServer
}

// geoAutoUpdateReloader is set by the web server so other packages (e.g. the
// controllers) can re-register the geo auto-update cron entry when the admin
// changes the cron expression, without restarting the panel.
var geoAutoUpdateReloader func(string) error

// SetGeoAutoUpdateReloader registers the reloader implementation.
func SetGeoAutoUpdateReloader(f func(string) error) {
	geoAutoUpdateReloader = f
}

// ReloadGeoAutoUpdate re-schedules the geo auto-update job with the given
// cron expression. It is a no-op if no reloader has been registered yet.
func ReloadGeoAutoUpdate(expr string) error {
	if geoAutoUpdateReloader == nil {
		return nil
	}
	return geoAutoUpdateReloader(expr)
}
