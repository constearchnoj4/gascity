
package wscontract

type BeadsBead struct {
  Assignee string `json:"assignee,omitempty"`
  CreatedAt string `json:"created_at,omitempty"`
  Dependencies []BeadsDep `json:"dependencies,omitempty"`
  Description string `json:"description,omitempty"`
  From string `json:"from,omitempty"`
  Id string `json:"id,omitempty"`
  IssueType string `json:"issue_type,omitempty"`
  Labels []string `json:"labels,omitempty"`
  Metadata map[string]string `json:"metadata,omitempty"`
  Needs []string `json:"needs,omitempty"`
  Parent string `json:"parent,omitempty"`
  Priority *int `json:"priority,omitempty"`
  Ref string `json:"ref,omitempty"`
  Status string `json:"status,omitempty"`
  Title string `json:"title,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-,omitempty"`
}