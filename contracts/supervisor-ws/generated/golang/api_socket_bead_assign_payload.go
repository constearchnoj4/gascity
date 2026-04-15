
package wscontract

type ApiSocketBeadAssignPayload struct {
  Assignee string `json:"assignee,omitempty"`
  Id string `json:"id,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-,omitempty"`
}