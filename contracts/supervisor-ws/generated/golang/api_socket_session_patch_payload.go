
package wscontract

type ApiSocketSessionPatchPayload struct {
  Alias *string `json:"alias,omitempty"`
  Id string `json:"id,omitempty"`
  Title *string `json:"title,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-,omitempty"`
}