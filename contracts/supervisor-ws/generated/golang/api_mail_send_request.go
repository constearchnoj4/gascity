
package wscontract

type ApiMailSendRequest struct {
  Body string `json:"body,omitempty"`
  From string `json:"from,omitempty"`
  Rig string `json:"rig,omitempty"`
  Subject string `json:"subject,omitempty"`
  To string `json:"to,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-,omitempty"`
}