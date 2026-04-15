
package wscontract

type ApiSocketProviderUpdatePayload struct {
  Name string `json:"name,omitempty"`
  Update map[string]interface{} `json:"update,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-,omitempty"`
}