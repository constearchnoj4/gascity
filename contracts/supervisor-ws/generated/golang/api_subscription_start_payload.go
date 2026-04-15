
package wscontract

type ApiSubscriptionStartPayload struct {
  // Resume from this cursor
  AfterCursor string `json:"after_cursor,omitempty"`
  // Resume from this event sequence
  AfterSeq int `json:"after_seq,omitempty"`
  // Stream format: 'text', 'raw', 'jsonl'
  Format string `json:"format,omitempty"`
  // Subscription type: 'events' or 'session.stream'
  Kind string `json:"kind,omitempty"`
  // Session ID or name (for session.stream)
  Target string `json:"target,omitempty"`
  // Most recent N turns (0=all)
  Turns int `json:"turns,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-,omitempty"`
}