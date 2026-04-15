
package wscontract

type ApiSocketSessionsListPayload struct {
  Cursor string `json:"cursor,omitempty"`
  Limit *int `json:"limit,omitempty"`
  Peek bool `json:"peek,omitempty"`
  State string `json:"state,omitempty"`
  Template string `json:"template,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-,omitempty"`
}