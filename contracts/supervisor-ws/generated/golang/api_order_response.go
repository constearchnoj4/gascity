
package wscontract

type ApiOrderResponse struct {
  CaptureOutput bool `json:"capture_output,omitempty"`
  Check string `json:"check,omitempty"`
  Description string `json:"description,omitempty"`
  Enabled bool `json:"enabled,omitempty"`
  Exec string `json:"exec,omitempty"`
  Formula string `json:"formula,omitempty"`
  Gate string `json:"gate,omitempty"`
  Interval string `json:"interval,omitempty"`
  Name string `json:"name,omitempty"`
  On string `json:"on,omitempty"`
  Pool string `json:"pool,omitempty"`
  Rig string `json:"rig,omitempty"`
  Schedule string `json:"schedule,omitempty"`
  ScopedName string `json:"scoped_name,omitempty"`
  Timeout string `json:"timeout,omitempty"`
  TimeoutMs int `json:"timeout_ms,omitempty"`
  ReservedType string `json:"type,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-,omitempty"`
}