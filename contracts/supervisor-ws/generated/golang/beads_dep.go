
package wscontract

type BeadsDep struct {
  DependsOnId string `json:"depends_on_id,omitempty"`
  IssueId string `json:"issue_id,omitempty"`
  ReservedType string `json:"type,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-,omitempty"`
}