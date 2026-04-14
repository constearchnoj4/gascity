package api

import "context"

func init() {
	RegisterVoidAction("packs.list", ActionDef{
		Description:       "List packs",
		RequiresCityScope: true,
	}, func(_ context.Context, s *Server) (map[string]any, error) {
		return map[string]any{"packs": s.listPacks()}, nil
	})
}
