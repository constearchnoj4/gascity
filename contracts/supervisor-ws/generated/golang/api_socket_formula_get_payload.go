
package wscontract

type ApiSocketFormulaGetPayload struct {
  Name string `json:"name,omitempty"`
  ScopeKind string `json:"scope_kind,omitempty"`
  ScopeRef string `json:"scope_ref,omitempty"`
  Target string `json:"target,omitempty"`
  Vars map[string]string `json:"vars,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-,omitempty"`
}