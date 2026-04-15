
package wscontract

type ApiExtmsgBindRequest struct {
  Conversation *ExtmsgConversationRef `json:"conversation,omitempty"`
  Metadata map[string]string `json:"metadata,omitempty"`
  SessionId string `json:"session_id,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-,omitempty"`
}