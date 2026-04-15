
package wscontract

type ApiWatchParams struct {
  // Block until server index exceeds this value
  Index int `json:"index,omitempty"`
  // Maximum wait duration (e.g. '30s')
  Wait string `json:"wait,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-,omitempty"`
}