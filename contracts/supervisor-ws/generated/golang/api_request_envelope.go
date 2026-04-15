
package wscontract

type ApiRequestEnvelope struct {
  // Dotted action name (e.g. 'beads.list')
  Action string `json:"action,omitempty"`
  // Client-assigned correlation ID
  Id string `json:"id,omitempty"`
  // Deduplication key for mutation replay
  IdempotencyKey string `json:"idempotency_key,omitempty"`
  // Action-specific request payload
  Payload interface{} `json:"payload,omitempty"`
  Scope *ApiScope `json:"scope,omitempty"`
  // Must be 'request'
  ReservedType string `json:"type,omitempty"`
  Watch *ApiWatchParams `json:"watch,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-,omitempty"`
}