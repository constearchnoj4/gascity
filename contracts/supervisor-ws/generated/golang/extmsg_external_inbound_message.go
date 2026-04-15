
package wscontract

type ExtmsgExternalInboundMessage struct {
  Actor *ExtmsgExternalActor `json:"actor,omitempty"`
  Attachments []ExtmsgExternalAttachment `json:"attachments,omitempty"`
  Conversation *ExtmsgConversationRef `json:"conversation,omitempty"`
  DedupKey string `json:"dedup_key,omitempty"`
  ExplicitTarget string `json:"explicit_target,omitempty"`
  ProviderMessageId string `json:"provider_message_id,omitempty"`
  ReceivedAt string `json:"received_at,omitempty"`
  ReplyToMessageId string `json:"reply_to_message_id,omitempty"`
  Text string `json:"text,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-,omitempty"`
}