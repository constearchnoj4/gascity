
package wscontract

type ApiSocketSessionRespondPayload struct {
  Action string `json:"action,omitempty"`
  Id string `json:"id,omitempty"`
  Metadata map[string]string `json:"metadata,omitempty"`
  RequestId string `json:"request_id,omitempty"`
  Text string `json:"text,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-,omitempty"`
}