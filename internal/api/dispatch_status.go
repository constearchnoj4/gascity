package api

import "context"

func init() {
	RegisterMeta("cities.list", ActionDef{
		Description: "List managed cities (supervisor)",
		ServerRoles: actionServerRoleSupervisor,
	})

	RegisterVoidAction("status.get", ActionDef{
		Description:       "City status snapshot",
		RequiresCityScope: true,
		SupportsWatch:     true,
	}, func(_ context.Context, s *Server) (any, error) {
		return s.statusSnapshot(), nil
	})
}
