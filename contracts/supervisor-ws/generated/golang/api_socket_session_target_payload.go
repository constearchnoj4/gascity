
package wscontract

type ApiSocketSessionTargetPayload struct {
  Id string `json:"id,omitempty"`
  Peek bool `json:"peek,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-,omitempty"`
}