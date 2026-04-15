
package wscontract

type ApiBeadCreateRequest struct {
  Assignee string `json:"assignee,omitempty"`
  Description string `json:"description,omitempty"`
  Labels []string `json:"labels,omitempty"`
  Priority *int `json:"priority,omitempty"`
  Rig string `json:"rig,omitempty"`
  Title string `json:"title,omitempty"`
  ReservedType string `json:"type,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-,omitempty"`
}