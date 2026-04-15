
package wscontract

type ApiSocketMailCountPayload struct {
  Agent string `json:"agent,omitempty"`
  Rig string `json:"rig,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-,omitempty"`
}