
package wscontract

type ApiExtmsgParticipantUpsertRequest struct {
  GroupId string `json:"group_id,omitempty"`
  Handle string `json:"handle,omitempty"`
  Metadata map[string]string `json:"metadata,omitempty"`
  Public bool `json:"public,omitempty"`
  SessionId string `json:"session_id,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-,omitempty"`
}