
package wscontract

type ApiListResponse struct {
  Items interface{} `json:"items,omitempty"`
  NextCursor string `json:"next_cursor,omitempty"`
  Total int `json:"total,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-,omitempty"`
}