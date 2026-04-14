package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gastownhall/gascity/internal/config"
	"github.com/gastownhall/gascity/internal/sessionlog"
	workdirutil "github.com/gastownhall/gascity/internal/workdir"
)

// outputTurn is a single conversation turn in the unified output response.
type outputTurn struct {
	Role      string `json:"role"`
	Text      string `json:"text"`
	Timestamp string `json:"timestamp,omitempty"`
}

// agentOutputResponse is the response for GET /v0/agent/{name}/output.
type agentOutputResponse struct {
	Agent      string                     `json:"agent"`
	Format     string                     `json:"format"` // "conversation" or "text"
	Turns      []outputTurn               `json:"turns"`
	Pagination *sessionlog.PaginationInfo `json:"pagination,omitempty"`
}

// handleAgentOutput returns unified conversation output for an agent.
// Tries structured session logs first, falls back to Peek().
//
// Query params:
//   - tail: number of compaction segments to return (default 1, 0 = all)
//   - before: message UUID cursor for loading older messages
func (s *Server) handleAgentOutput(w http.ResponseWriter, r *http.Request, name string) {
	cfg := s.state.Config()
	agentCfg, ok := findAgent(cfg, name)
	if !ok {
		writeError(w, http.StatusNotFound, "not_found", "agent "+name+" not found")
		return
	}

	// Try structured session log first.
	resp, err := s.trySessionLogOutput(r, name, agentCfg)
	if err != nil {
		// Session file exists but failed to read — surface the error.
		writeError(w, http.StatusInternalServerError, "internal", "reading session log: "+err.Error())
		return
	}
	if resp != nil {
		writeJSON(w, http.StatusOK, resp)
		return
	}

	// No session file found — fall back to Peek() (raw terminal text).
	s.peekFallbackOutput(w, name, cfg)
}

// trySessionLogOutput attempts to read structured conversation data from
// a Claude JSONL session file. Returns (nil, nil) if no session file is
// found (expected — triggers fallback). Returns (nil, err) if the file
// exists but cannot be read (unexpected — surface to caller).
func (s *Server) trySessionLogOutput(r *http.Request, name string, agentCfg config.Agent) (*agentOutputResponse, error) {
	cfg := s.state.Config()
	workDir := s.resolveAgentWorkDir(agentCfg, name)
	if workDir == "" {
		return nil, nil
	}
	provider := strings.TrimSpace(agentCfg.Provider)
	if provider == "" && cfg != nil {
		provider = strings.TrimSpace(cfg.Workspace.Provider)
	}

	searchPaths := s.sessionLogSearchPaths
	if searchPaths == nil {
		searchPaths = sessionlog.MergeSearchPaths(cfg.Daemon.ObservePaths)
	}
	path := sessionlog.FindSessionFileForProvider(searchPaths, provider, workDir)
	if path == "" {
		return nil, nil
	}

	tail := 1
	if v := r.URL.Query().Get("tail"); v != "" {
		if n, err := strconv.Atoi(v); err == nil && n >= 0 {
			tail = n
		}
	}
	before := r.URL.Query().Get("before")

	var sess *sessionlog.Session
	var err error
	if before != "" {
		sess, err = sessionlog.ReadProviderFileOlder(provider, path, tail, before)
	} else {
		sess, err = sessionlog.ReadProviderFile(provider, path, tail)
	}
	if err != nil {
		return nil, err
	}

	turns := make([]outputTurn, 0, len(sess.Messages))
	for _, e := range sess.Messages {
		turn := entryToTurn(e)
		if turn.Text == "" {
			continue
		}
		turns = append(turns, turn)
	}

	return &agentOutputResponse{
		Agent:      name,
		Format:     "conversation",
		Turns:      turns,
		Pagination: sess.Pagination,
	}, nil
}

// peekFallbackOutput returns raw terminal text wrapped as a single turn.
func (s *Server) peekFallbackOutput(w http.ResponseWriter, name string, cfg *config.City) {
	sp := s.state.SessionProvider()
	sessionName := agentSessionName(s.state.CityName(), name, cfg.Workspace.SessionTemplate)

	if !sp.IsRunning(sessionName) {
		writeError(w, http.StatusNotFound, "not_found", "agent "+name+" not running")
		return
	}

	output, err := sp.Peek(sessionName, 100)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "internal", err.Error())
		return
	}

	turns := []outputTurn{}
	if output != "" {
		turns = append(turns, outputTurn{Role: "output", Text: output})
	}

	writeJSON(w, http.StatusOK, agentOutputResponse{
		Agent:  name,
		Format: "text",
		Turns:  turns,
	})
}

// resolveAgentWorkDir returns the absolute working directory for an agent,
// honoring work_dir template expansion.
func (s *Server) resolveAgentWorkDir(a config.Agent, qualifiedName string) string {
	cfg := s.state.Config()
	return workdirutil.ResolveWorkDirPath(
		s.state.CityPath(),
		workdirutil.CityName(s.state.CityPath(), cfg),
		qualifiedName,
		a,
		cfg.Rigs,
	)
}

// entryToTurn converts a sessionlog Entry to a human-readable outputTurn.
func entryToTurn(e *sessionlog.Entry) outputTurn {
	turn := outputTurn{
		Role: e.Type,
	}
	if !e.Timestamp.IsZero() {
		turn.Timestamp = e.Timestamp.Format("2006-01-02T15:04:05Z07:00")
	}

	// Try plain string content (message is a JSON object with string content).
	if text := e.TextContent(); text != "" {
		turn.Text = text
		return turn
	}

	// Try structured content blocks — extract human-readable text.
	if blocks := e.ContentBlocks(); len(blocks) > 0 {
		var parts []string
		for _, b := range blocks {
			switch b.Type {
			case "text":
				if b.Text != "" {
					parts = append(parts, b.Text)
				}
			case "tool_use":
				if b.Name != "" {
					parts = append(parts, "["+b.Name+"]")
				}
			case "tool_result":
				text := extractToolResultText(b.Content)
				if text != "" {
					if len(text) > 500 {
						text = text[:500] + "…"
					}
					parts = append(parts, "[result] "+text)
				}
			case "thinking":
				// Redact thinking blocks — internal model reasoning
				// should not be surfaced to the UI.
				parts = append(parts, "[thinking]")
			}
		}
		turn.Text = strings.Join(parts, "\n")
		return turn
	}

	// Claude JSONL double-encodes the message field as a JSON string
	// containing JSON. Unwrap and try again.
	turn.Text = unwrapDoubleEncoded(e.Message)
	return turn
}

// extractToolResultText extracts human-readable text from a tool_result
// Content field (json.RawMessage). The content can be a plain string or
// an array of content blocks (e.g., [{type:"text", text:"..."}]).
func extractToolResultText(raw json.RawMessage) string {
	if len(raw) == 0 {
		return ""
	}
	// Try plain string.
	var s string
	if json.Unmarshal(raw, &s) == nil && s != "" {
		return s
	}
	// Try array of content blocks.
	var blocks []sessionlog.ContentBlock
	if json.Unmarshal(raw, &blocks) == nil {
		var parts []string
		for _, b := range blocks {
			if b.Type == "text" && b.Text != "" {
				parts = append(parts, b.Text)
			}
		}
		return strings.Join(parts, "\n")
	}
	return ""
}

// outputStreamPollInterval controls how often stream implementations check for
// new output. Used by the log file watcher and session stream emitters.
const outputStreamPollInterval = 2 * time.Second

// unwrapDoubleEncoded handles Claude's double-encoded message format
// where the "message" field is a JSON string containing a JSON object.
// Returns the human-readable content text, or "" if not parseable.
func unwrapDoubleEncoded(raw []byte) string {
	if len(raw) == 0 {
		return ""
	}
	// Try to unwrap: raw might be a JSON string like "{\"role\":...}"
	var inner string
	if err := json.Unmarshal(raw, &inner); err != nil {
		return ""
	}
	// Now inner is the JSON object as a string. Parse it.
	var mc sessionlog.MessageContent
	if err := json.Unmarshal([]byte(inner), &mc); err != nil {
		return ""
	}
	// Try string content.
	var s string
	if err := json.Unmarshal(mc.Content, &s); err == nil && s != "" {
		return s
	}
	// Try array of content blocks.
	var blocks []sessionlog.ContentBlock
	if err := json.Unmarshal(mc.Content, &blocks); err == nil {
		var parts []string
		for _, b := range blocks {
			if b.Type == "text" && b.Text != "" {
				parts = append(parts, b.Text)
			}
		}
		return strings.Join(parts, "\n")
	}
	return ""
}
