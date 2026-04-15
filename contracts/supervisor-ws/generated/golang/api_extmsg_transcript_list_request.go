
package wscontract

type ApiExtmsgTranscriptListRequest struct {
  AccountId string `json:"account_id,omitempty"`
  ConversationId string `json:"conversation_id,omitempty"`
  Kind string `json:"kind,omitempty"`
  ParentConversationId string `json:"parent_conversation_id,omitempty"`
  Provider string `json:"provider,omitempty"`
  ScopeId string `json:"scope_id,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-,omitempty"`
}