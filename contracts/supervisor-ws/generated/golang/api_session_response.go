
package wscontract

type ApiSessionResponse struct {
  ActiveBead string `json:"active_bead,omitempty"`
  Activity string `json:"activity,omitempty"`
  Alias string `json:"alias,omitempty"`
  Attached bool `json:"attached,omitempty"`
  ConfiguredNamedSession bool `json:"configured_named_session,omitempty"`
  ContextPct *int `json:"context_pct,omitempty"`
  ContextWindow *int `json:"context_window,omitempty"`
  CreatedAt string `json:"created_at,omitempty"`
  DisplayName string `json:"display_name,omitempty"`
  Id string `json:"id,omitempty"`
  Kind string `json:"kind,omitempty"`
  LastActive string `json:"last_active,omitempty"`
  LastOutput string `json:"last_output,omitempty"`
  Metadata map[string]string `json:"metadata,omitempty"`
  Model string `json:"model,omitempty"`
  Options map[string]string `json:"options,omitempty"`
  Pool string `json:"pool,omitempty"`
  Provider string `json:"provider,omitempty"`
  Reason string `json:"reason,omitempty"`
  Rig string `json:"rig,omitempty"`
  Running bool `json:"running,omitempty"`
  SessionName string `json:"session_name,omitempty"`
  State string `json:"state,omitempty"`
  SubmissionCapabilities *SessionSubmissionCapabilities `json:"submission_capabilities,omitempty"`
  Template string `json:"template,omitempty"`
  Title string `json:"title,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-,omitempty"`
}