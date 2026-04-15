
package wscontract

type ApiEventEmitRequest struct {
  Actor string `json:"actor,omitempty"`
  Message string `json:"message,omitempty"`
  Subject string `json:"subject,omitempty"`
  ReservedType string `json:"type,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-,omitempty"`
}