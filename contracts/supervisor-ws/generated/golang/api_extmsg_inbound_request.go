
package wscontract

type ApiExtmsgInboundRequest struct {
  AccountId string `json:"account_id,omitempty"`
  Message *ExtmsgExternalInboundMessage `json:"message,omitempty"`
  Payload string `json:"payload,omitempty"`
  Provider string `json:"provider,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-,omitempty"`
}