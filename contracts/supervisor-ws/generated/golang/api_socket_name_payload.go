
package wscontract

type ApiSocketNamePayload struct {
  Name string `json:"name,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-,omitempty"`
}