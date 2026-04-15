
package wscontract

type SessionSubmissionCapabilities struct {
  SupportsFollowUp bool `json:"supports_follow_up,omitempty"`
  SupportsInterruptNow bool `json:"supports_interrupt_now,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-,omitempty"`
}