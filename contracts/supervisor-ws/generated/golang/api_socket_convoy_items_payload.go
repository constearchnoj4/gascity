
package wscontract

type ApiSocketConvoyItemsPayload struct {
  Id string `json:"id,omitempty"`
  Items []string `json:"items,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-,omitempty"`
}