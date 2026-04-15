
package wscontract

type ApiExtmsgAdapterUnregisterRequest struct {
  AccountId string `json:"account_id,omitempty"`
  Provider string `json:"provider,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-,omitempty"`
}