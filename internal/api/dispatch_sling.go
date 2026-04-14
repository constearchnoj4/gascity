package api

import "context"

func init() {
	RegisterAction("sling.run", ActionDef{
		Description:       "Run sling dispatch",
		IsMutation:        true,
		RequiresCityScope: true,
	}, func(ctx context.Context, s *Server, payload slingBody) (any, error) {
		result, _, err := s.runSling(ctx, payload)
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}
