
package wscontract

type ApiExtmsgAdapterRegisterRequest struct {
  AccountId string `json:"account_id,omitempty"`
  CallbackUrl string `json:"callback_url,omitempty"`
  Capabilities map[string]interface{} `json:"capabilities,omitempty"`
  Name string `json:"name,omitempty"`
  Provider string `json:"provider,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-,omitempty"`
}