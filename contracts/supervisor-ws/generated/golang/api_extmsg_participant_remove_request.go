
package wscontract

type ApiExtmsgParticipantRemoveRequest struct {
  GroupId string `json:"group_id,omitempty"`
  Handle string `json:"handle,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-,omitempty"`
}