
package wscontract

type ApiSocketRigCreatePayload struct {
  Name string `json:"name,omitempty"`
  Path string `json:"path,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-,omitempty"`
}