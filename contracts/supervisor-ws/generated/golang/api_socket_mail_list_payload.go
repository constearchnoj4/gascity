
package wscontract

type ApiSocketMailListPayload struct {
  Agent string `json:"agent,omitempty"`
  Cursor string `json:"cursor,omitempty"`
  Limit *int `json:"limit,omitempty"`
  Rig string `json:"rig,omitempty"`
  Status string `json:"status,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-,omitempty"`
}