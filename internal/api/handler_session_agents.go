package api

import (
	"encoding/json"
	"errors"

	"github.com/gastownhall/gascity/internal/sessionlog"
)

// handleSessionAgentList returns subagent mappings for a session.
//
//	GET /v0/session/{id}/agents
//	Response: { "agents": [{ "agent_id": "...", "parent_tool_use_id": "..." }] }
// handleSessionAgentGet returns the transcript and status of a subagent.
//
//	GET /v0/session/{id}/agents/{agentId}
//	Response: { "messages": [...], "status": "completed|running|pending|failed" }
// --- Shared methods for WS dispatch ---

func (s *Server) listSessionAgents(target string) (any, error) {
	store := s.state.CityBeadStore()
	if store == nil {
		return nil, httpError{status: 503, code: "unavailable", message: "no bead store configured"}
	}
	id, err := s.resolveSessionIDAllowClosedWithConfig(store, target)
	if err != nil {
		return nil, err
	}
	mgr := s.sessionManager(store)
	logPath, err := mgr.TranscriptPath(id, s.sessionLogPaths())
	if err != nil {
		return nil, err
	}
	if logPath == "" {
		return map[string]any{"agents": []any{}}, nil
	}
	mappings, err := sessionlog.FindAgentMappings(logPath)
	if err != nil {
		return nil, httpError{status: 500, code: "internal", message: "failed to list agents"}
	}
	if mappings == nil {
		mappings = []sessionlog.AgentMapping{}
	}
	return map[string]any{"agents": mappings}, nil
}

func (s *Server) getSessionAgent(target, agentID string) (any, error) {
	store := s.state.CityBeadStore()
	if store == nil {
		return nil, httpError{status: 503, code: "unavailable", message: "no bead store configured"}
	}
	id, err := s.resolveSessionIDAllowClosedWithConfig(store, target)
	if err != nil {
		return nil, err
	}
	if agentID == "" {
		return nil, httpError{status: 400, code: "invalid", message: "agent_id is required"}
	}
	if err := sessionlog.ValidateAgentID(agentID); err != nil {
		return nil, httpError{status: 400, code: "invalid", message: err.Error()}
	}
	mgr := s.sessionManager(store)
	logPath, err := mgr.TranscriptPath(id, s.sessionLogPaths())
	if err != nil {
		return nil, err
	}
	if logPath == "" {
		return nil, httpError{status: 404, code: "not_found", message: "no transcript found for session " + id}
	}
	agentSession, err := sessionlog.ReadAgentSession(logPath, agentID)
	if err != nil {
		if errors.Is(err, sessionlog.ErrAgentNotFound) {
			return nil, httpError{status: 404, code: "not_found", message: "agent not found"}
		}
		return nil, httpError{status: 500, code: "internal", message: "failed to read agent transcript"}
	}
	rawMessages := make([]json.RawMessage, 0, len(agentSession.Messages))
	for _, entry := range agentSession.Messages {
		if len(entry.Raw) > 0 {
			rawMessages = append(rawMessages, entry.Raw)
		}
	}
	return map[string]any{"messages": rawMessages, "status": agentSession.Status}, nil
}
