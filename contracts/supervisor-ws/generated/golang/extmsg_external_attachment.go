
package wscontract

type ExtmsgExternalAttachment struct {
  MimeType string `json:"mime_type,omitempty"`
  ProviderId string `json:"provider_id,omitempty"`
  Url string `json:"url,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-,omitempty"`
}