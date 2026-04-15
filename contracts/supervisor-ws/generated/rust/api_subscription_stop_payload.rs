// ApiSubscriptionStopPayload represents a ApiSubscriptionStopPayload model.
#[derive(Clone, Debug, Deserialize, Serialize)]
pub struct ApiSubscriptionStopPayload {
    #[serde(rename="subscription_id", skip_serializing_if = "Option::is_none")]
    pub subscription_id: Option<String>,
    #[serde(rename="additionalProperties", skip_serializing_if = "Option::is_none")]
    pub additional_properties: Option<std::collections::HashMap<String, serde_json::Value>>,
}
