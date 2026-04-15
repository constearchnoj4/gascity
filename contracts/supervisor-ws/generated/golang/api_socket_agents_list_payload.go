
package wscontract

type ApiSocketAgentsListPayload struct {
  Peek bool `json:"peek,omitempty"`
  Pool string `json:"pool,omitempty"`
  Rig string `json:"rig,omitempty"`
  Running string `json:"running,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-,omitempty"`
}