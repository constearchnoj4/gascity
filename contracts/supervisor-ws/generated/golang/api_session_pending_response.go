
package wscontract

type ApiSessionPendingResponse struct {
  Pending *RuntimePendingInteraction `json:"pending,omitempty"`
  Supported bool `json:"supported,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-,omitempty"`
}