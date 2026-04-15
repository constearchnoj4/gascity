
package wscontract

type RuntimePendingInteraction struct {
  Kind string `json:"kind,omitempty"`
  Metadata map[string]string `json:"metadata,omitempty"`
  Options []string `json:"options,omitempty"`
  Prompt string `json:"prompt,omitempty"`
  RequestId string `json:"request_id,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-,omitempty"`
}