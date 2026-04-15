
package wscontract

type ApiSlingBody struct {
  AttachedBeadId string `json:"attached_bead_id,omitempty"`
  Bead string `json:"bead,omitempty"`
  Formula string `json:"formula,omitempty"`
  Rig string `json:"rig,omitempty"`
  ScopeKind string `json:"scope_kind,omitempty"`
  ScopeRef string `json:"scope_ref,omitempty"`
  Target string `json:"target,omitempty"`
  Title string `json:"title,omitempty"`
  Vars map[string]string `json:"vars,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-,omitempty"`
}