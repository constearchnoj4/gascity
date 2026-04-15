
package wscontract

type ApiSocketFormulaFeedPayload struct {
  Limit int `json:"limit,omitempty"`
  ScopeKind string `json:"scope_kind,omitempty"`
  ScopeRef string `json:"scope_ref,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-,omitempty"`
}