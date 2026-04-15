
package wscontract

type ApiSocketFormulaScopePayload struct {
  ScopeKind string `json:"scope_kind,omitempty"`
  ScopeRef string `json:"scope_ref,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-,omitempty"`
}