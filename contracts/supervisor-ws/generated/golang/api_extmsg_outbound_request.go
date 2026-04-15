
package wscontract

type ApiExtmsgOutboundRequest struct {
  Conversation *ExtmsgConversationRef `json:"conversation,omitempty"`
  IdempotencyKey string `json:"idempotency_key,omitempty"`
  ReplyToMessageId string `json:"reply_to_message_id,omitempty"`
  SessionId string `json:"session_id,omitempty"`
  Text string `json:"text,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-,omitempty"`
}