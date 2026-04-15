
package wscontract

type ApiSessionCreateRequest struct {
  Alias string `json:"alias,omitempty"`
  Async bool `json:"async,omitempty"`
  Kind string `json:"kind,omitempty"`
  Message string `json:"message,omitempty"`
  Name string `json:"name,omitempty"`
  Options map[string]string `json:"options,omitempty"`
  ProjectId string `json:"project_id,omitempty"`
  SessionName *string `json:"session_name,omitempty"`
  Title string `json:"title,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-,omitempty"`
}