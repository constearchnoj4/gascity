
package wscontract

type ApiWorkflowDepResponse struct {
  From string `json:"from,omitempty"`
  Kind string `json:"kind,omitempty"`
  To string `json:"to,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-,omitempty"`
}