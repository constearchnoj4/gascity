
package wscontract

type ApiBeadGraphResponseJson struct {
  Beads []BeadsBead `json:"beads,omitempty"`
  Deps []ApiWorkflowDepResponse `json:"deps,omitempty"`
  Root *BeadsBead `json:"root,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-,omitempty"`
}