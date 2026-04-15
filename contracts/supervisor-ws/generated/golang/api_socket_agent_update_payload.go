
package wscontract

type ApiSocketAgentUpdatePayload struct {
  Name string `json:"name,omitempty"`
  Provider string `json:"provider,omitempty"`
  Scope string `json:"scope,omitempty"`
  Suspended *bool `json:"suspended,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-,omitempty"`
}