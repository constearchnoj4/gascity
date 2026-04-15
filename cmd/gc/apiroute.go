package main

import (
	"fmt"
	"net"
	"path/filepath"
	"strconv"

	"github.com/gastownhall/gascity/internal/api"
	"github.com/gastownhall/gascity/internal/config"
	"github.com/gastownhall/gascity/internal/fsys"
)

// apiClient returns an API client if a city controller or supervisor API is
// available for the target city. When a controller is running with an API
// port, CLI mutations must use the API and fail closed on transport or
// read-only errors instead of silently mutating local files behind its back.
func apiClient(cityPath string) *api.Client {
	// Check if controller is alive.
	if controllerAlive(cityPath) != 0 {
		// Load config to find API port.
		tomlPath := filepath.Join(cityPath, "city.toml")
		cfg, err := config.Load(fsys.OSFS{}, tomlPath)
		if err != nil {
			return nil
		}
		if cfg.API.Port <= 0 {
			return nil
		}
		return api.NewClient(apiBaseURL(cfg.API.BindOrDefault(), cfg.API.Port))
	}
	return supervisorCityAPIClient(cityPath)
}

func apiBaseURL(bind string, port int) string {
	switch bind {
	case "0.0.0.0":
		bind = "127.0.0.1"
	case "::", "[::]":
		bind = "::1"
	}
	return fmt.Sprintf("http://%s", net.JoinHostPort(bind, strconv.Itoa(port)))
}

// resolveAgentForAPI resolves a bare agent name (e.g., "worker") to its
// qualified form (e.g., "myrig/worker") using the current rig context, so
// the API server can find the agent. If already qualified or resolution
// fails, the original name is returned.
func resolveAgentForAPI(cityPath, name string) string {
	cfg, err := loadCityConfig(cityPath)
	if err != nil {
		return name
	}
	resolved, ok := resolveAgentIdentity(cfg, name, currentRigContext(cfg))
	if !ok {
		return name
	}
	return resolved.QualifiedName()
}
