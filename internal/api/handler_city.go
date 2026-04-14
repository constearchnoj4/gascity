package api

import (
	"net/http"
	"time"
)

type cityGetResponse struct {
	Name            string `json:"name"`
	Path            string `json:"path"`
	Version         string `json:"version,omitempty"`
	Suspended       bool   `json:"suspended"`
	Provider        string `json:"provider,omitempty"`
	SessionTemplate string `json:"session_template,omitempty"`
	UptimeSec       int    `json:"uptime_sec"`
	AgentCount      int    `json:"agent_count"`
	RigCount        int    `json:"rig_count"`
}

// cityPatchRequest is the JSON body for PATCH /v0/city.
type cityPatchRequest struct {
	Suspended *bool `json:"suspended,omitempty"`
}

func (s *Server) cityGet() cityGetResponse {
	cfg := s.state.Config()
	return cityGetResponse{
		Name:            s.state.CityName(),
		Path:            s.state.CityPath(),
		Version:         s.state.Version(),
		Suspended:       cfg.Workspace.Suspended,
		Provider:        cfg.Workspace.Provider,
		SessionTemplate: cfg.Workspace.SessionTemplate,
		UptimeSec:       int(time.Since(s.state.StartedAt()).Seconds()),
		AgentCount:      len(cfg.Agents),
		RigCount:        len(cfg.Rigs),
	}
}

func (s *Server) patchCitySuspended(suspend bool) error {
	sm, ok := s.state.(StateMutator)
	if !ok {
		return httpError{status: http.StatusNotImplemented, code: "internal", message: "mutations not supported"}
	}
	if suspend {
		return sm.SuspendCity()
	}
	return sm.ResumeCity()
}
