
package wscontract

type ApiSocketBeadsListPayload struct {
  Assignee string `json:"assignee,omitempty"`
  Cursor string `json:"cursor,omitempty"`
  Label string `json:"label,omitempty"`
  Limit *int `json:"limit,omitempty"`
  Rig string `json:"rig,omitempty"`
  Status string `json:"status,omitempty"`
  ReservedType string `json:"type,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-,omitempty"`
}