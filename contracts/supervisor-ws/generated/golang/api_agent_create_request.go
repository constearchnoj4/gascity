
package wscontract

type ApiAgentCreateRequest struct {
  Dir string `json:"dir,omitempty"`
  Name string `json:"name,omitempty"`
  Provider string `json:"provider,omitempty"`
  Scope string `json:"scope,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-,omitempty"`
}