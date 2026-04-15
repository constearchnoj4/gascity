
package wscontract

type MailMessage struct {
  Body string `json:"body,omitempty"`
  Cc []string `json:"cc,omitempty"`
  CreatedAt string `json:"created_at,omitempty"`
  From string `json:"from,omitempty"`
  Id string `json:"id,omitempty"`
  Priority int `json:"priority,omitempty"`
  Read bool `json:"read,omitempty"`
  ReplyTo string `json:"reply_to,omitempty"`
  Rig string `json:"rig,omitempty"`
  Subject string `json:"subject,omitempty"`
  ThreadId string `json:"thread_id,omitempty"`
  To string `json:"to,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-,omitempty"`
}