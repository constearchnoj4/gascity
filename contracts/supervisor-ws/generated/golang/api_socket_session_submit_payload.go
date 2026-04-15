
package wscontract

type ApiSocketSessionSubmitPayload struct {
  Id string `json:"id,omitempty"`
  Intent string `json:"intent,omitempty"`
  Message string `json:"message,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-,omitempty"`
}