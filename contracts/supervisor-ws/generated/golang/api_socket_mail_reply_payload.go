
package wscontract

type ApiSocketMailReplyPayload struct {
  Body string `json:"body,omitempty"`
  From string `json:"from,omitempty"`
  Id string `json:"id,omitempty"`
  Rig string `json:"rig,omitempty"`
  Subject string `json:"subject,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-,omitempty"`
}