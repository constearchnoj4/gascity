
package wscontract

type ApiFieldError struct {
  Field string `json:"field,omitempty"`
  Message string `json:"message,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-,omitempty"`
}