package api

import "context"

func init() {
	// cities.list is supervisor-only (handled inline in SupervisorMux.handleSocketRequest),
	// but needs a table entry for capabilities and spec generation.
	RegisterMeta("cities.list", ActionDef{
		Description: "List managed cities (supervisor)",
	})

	RegisterVoidAction("health.get", ActionDef{
		Description: "Health check",
	}, func(_ context.Context, s *Server) (map[string]any, error) {
		return s.healthResponse(), nil
	})

	RegisterVoidAction("status.get", ActionDef{
		Description:       "City status snapshot",
		RequiresCityScope: true,
		SupportsWatch:     true,
	}, func(_ context.Context, s *Server) (any, error) {
		return s.statusSnapshot(), nil
	})
}
