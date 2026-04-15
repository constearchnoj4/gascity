
package wscontract

type ApiErrorEnvelope struct {
  // Machine-readable error code
  Code string `json:"code,omitempty"`
  // Per-field validation errors
  Details []ApiFieldError `json:"details,omitempty"`
  // Correlation ID (empty for connection-level errors)
  Id string `json:"id,omitempty"`
  // Human-readable error message
  Message string `json:"message,omitempty"`
  // Must be 'error'
  ReservedType string `json:"type,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-,omitempty"`
}