
package wscontract

type ApiSocketEventsListPayload struct {
  Actor string `json:"actor,omitempty"`
  Cursor string `json:"cursor,omitempty"`
  Limit *int `json:"limit,omitempty"`
  Since string `json:"since,omitempty"`
  ReservedType string `json:"type,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-,omitempty"`
}