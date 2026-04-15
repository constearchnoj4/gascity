
package wscontract

type ApiExtmsgTranscriptAckRequest struct {
  Conversation *ExtmsgConversationRef `json:"conversation,omitempty"`
  Sequence int `json:"sequence,omitempty"`
  SessionId string `json:"session_id,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-,omitempty"`
}