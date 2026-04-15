
package wscontract

type ExtmsgExternalActor struct {
  DisplayName string `json:"display_name,omitempty"`
  Id string `json:"id,omitempty"`
  IsBot bool `json:"is_bot,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-,omitempty"`
}