
package wscontract

type ApiSocketSessionTranscriptPayload struct {
  Before string `json:"before,omitempty"`
  Format string `json:"format,omitempty"`
  Id string `json:"id,omitempty"`
  Turns int `json:"turns,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-,omitempty"`
}