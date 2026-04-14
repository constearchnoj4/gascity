package api

// Types used by WS dispatch handlers in dispatch_agents.go and websocket.go.

// agentCreateRequest is the JSON body for POST /v0/agents.
type agentCreateRequest struct {
	Name     string `json:"name"`
	Dir      string `json:"dir,omitempty"`
	Provider string `json:"provider"`
	Scope    string `json:"scope,omitempty"`
}

// agentUpdateRequest is the JSON body for PATCH /v0/agent/{name}.
type agentUpdateRequest struct {
	Provider  string `json:"provider,omitempty"`
	Scope     string `json:"scope,omitempty"`
	Suspended *bool  `json:"suspended,omitempty"`
}
