
package wscontract

type ApiConvoyCreateRequest struct {
  Items []string `json:"items,omitempty"`
  Rig string `json:"rig,omitempty"`
  Title string `json:"title,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-,omitempty"`
}