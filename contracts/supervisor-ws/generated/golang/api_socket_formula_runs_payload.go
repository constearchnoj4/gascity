
package wscontract

type ApiSocketFormulaRunsPayload struct {
  Limit int `json:"limit,omitempty"`
  Name string `json:"name,omitempty"`
  ScopeKind string `json:"scope_kind,omitempty"`
  ScopeRef string `json:"scope_ref,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-,omitempty"`
}