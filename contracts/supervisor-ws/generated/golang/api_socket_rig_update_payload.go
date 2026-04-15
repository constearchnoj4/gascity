
package wscontract

type ApiSocketRigUpdatePayload struct {
  Name string `json:"name,omitempty"`
  Path string `json:"path,omitempty"`
  Prefix string `json:"prefix,omitempty"`
  Suspended *bool `json:"suspended,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-,omitempty"`
}