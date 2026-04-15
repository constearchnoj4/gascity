
package wscontract

type ApiSocketProviderCreatePayload struct {
  Name string `json:"name,omitempty"`
  Spec map[string]interface{} `json:"spec,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-,omitempty"`
}