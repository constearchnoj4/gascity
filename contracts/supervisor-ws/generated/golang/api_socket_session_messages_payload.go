
package wscontract

type ApiSocketSessionMessagesPayload struct {
  Id string `json:"id,omitempty"`
  Message string `json:"message,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-,omitempty"`
}