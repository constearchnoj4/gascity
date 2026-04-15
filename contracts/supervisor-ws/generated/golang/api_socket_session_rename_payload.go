
package wscontract

type ApiSocketSessionRenamePayload struct {
  Id string `json:"id,omitempty"`
  Title string `json:"title,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-,omitempty"`
}