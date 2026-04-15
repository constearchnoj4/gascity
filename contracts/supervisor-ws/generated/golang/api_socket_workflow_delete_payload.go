
package wscontract

type ApiSocketWorkflowDeletePayload struct {
  Delete bool `json:"delete,omitempty"`
  Id string `json:"id,omitempty"`
  ScopeKind string `json:"scope_kind,omitempty"`
  ScopeRef string `json:"scope_ref,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-,omitempty"`
}