// ApiExtmsgInboundRequest represents a ApiExtmsgInboundRequest model.
#[derive(Clone, Debug, Deserialize, Serialize)]
pub struct ApiExtmsgInboundRequest {
    #[serde(rename="account_id", skip_serializing_if = "Option::is_none")]
    pub account_id: Option<String>,
    #[serde(rename="message", skip_serializing_if = "Option::is_none")]
    pub message: Option<Box<crate::ExtmsgExternalInboundMessage>>,
    #[serde(rename="payload", skip_serializing_if = "Option::is_none")]
    pub payload: Option<String>,
    #[serde(rename="provider", skip_serializing_if = "Option::is_none")]
    pub provider: Option<String>,
    #[serde(rename="additionalProperties", skip_serializing_if = "Option::is_none")]
    pub additional_properties: Option<std::collections::HashMap<String, serde_json::Value>>,
}
