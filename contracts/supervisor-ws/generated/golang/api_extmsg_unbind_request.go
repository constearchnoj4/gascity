
package wscontract

type ApiExtmsgUnbindRequest struct {
  Conversation *ExtmsgConversationRef `json:"conversation,omitempty"`
  SessionId string `json:"session_id,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-,omitempty"`
}