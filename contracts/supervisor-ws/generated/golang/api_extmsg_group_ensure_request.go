
package wscontract

type ApiExtmsgGroupEnsureRequest struct {
  DefaultHandle string `json:"default_handle,omitempty"`
  Metadata map[string]string `json:"metadata,omitempty"`
  Mode string `json:"mode,omitempty"`
  RootConversation *ExtmsgConversationRef `json:"root_conversation,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-,omitempty"`
}