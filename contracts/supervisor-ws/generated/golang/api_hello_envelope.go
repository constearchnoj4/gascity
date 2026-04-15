
package wscontract

type ApiHelloEnvelope struct {
  // Sorted list of supported action names
  Capabilities []string `json:"capabilities,omitempty"`
  // Protocol version (e.g. 'gc.v1alpha1')
  Protocol string `json:"protocol,omitempty"`
  // True if mutations are disabled
  ReadOnly bool `json:"read_only,omitempty"`
  // 'city' or 'supervisor'
  ServerRole string `json:"server_role,omitempty"`
  // Supported subscription types
  SubscriptionKinds []string `json:"subscription_kinds,omitempty"`
  // Must be 'hello'
  ReservedType string `json:"type,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-,omitempty"`
}