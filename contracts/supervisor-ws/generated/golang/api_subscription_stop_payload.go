
package wscontract

type ApiSubscriptionStopPayload struct {
  // Subscription to stop
  SubscriptionId string `json:"subscription_id,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-,omitempty"`
}