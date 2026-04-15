
package wscontract

type ApiResponseEnvelope struct {
  // Correlation ID matching the request
  Id string `json:"id,omitempty"`
  // Server event index for watch semantics
  Index int `json:"index,omitempty"`
  // Action-specific response payload
  Result interface{} `json:"result,omitempty"`
  // Must be 'response'
  ReservedType string `json:"type,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-,omitempty"`
}