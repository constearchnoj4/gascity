
package wscontract

type ApiSocketOrdersHistoryPayload struct {
  Before string `json:"before,omitempty"`
  Limit int `json:"limit,omitempty"`
  ScopedName string `json:"scoped_name,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-,omitempty"`
}