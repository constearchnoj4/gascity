
package wscontract

type ApiCityPatchRequest struct {
  Suspended *bool `json:"suspended,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-,omitempty"`
}