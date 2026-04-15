
package wscontract

type ApiScope struct {
  // City name for supervisor-scoped requests
  City string `json:"city,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-,omitempty"`
}