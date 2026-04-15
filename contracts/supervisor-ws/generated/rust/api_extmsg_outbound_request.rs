// ApiExtmsgOutboundRequest represents a ApiExtmsgOutboundRequest model.
#[derive(Clone, Debug, Deserialize, Serialize)]
pub struct ApiExtmsgOutboundRequest {
    #[serde(rename="conversation", skip_serializing_if = "Option::is_none")]
    pub conversation: Option<Box<crate::ExtmsgConversationRef>>,
    #[serde(rename="idempotency_key", skip_serializing_if = "Option::is_none")]
    pub idempotency_key: Option<String>,
    #[serde(rename="reply_to_message_id", skip_serializing_if = "Option::is_none")]
    pub reply_to_message_id: Option<String>,
    #[serde(rename="session_id", skip_serializing_if = "Option::is_none")]
    pub session_id: Option<String>,
    #[serde(rename="text", skip_serializing_if = "Option::is_none")]
    pub text: Option<String>,
    #[serde(rename="additionalProperties", skip_serializing_if = "Option::is_none")]
    pub additional_properties: Option<std::collections::HashMap<String, serde_json::Value>>,
}
