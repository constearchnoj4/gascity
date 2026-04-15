// ApiExtmsgBindRequest represents a ApiExtmsgBindRequest model.
#[derive(Clone, Debug, Deserialize, Serialize)]
pub struct ApiExtmsgBindRequest {
    #[serde(rename="conversation", skip_serializing_if = "Option::is_none")]
    pub conversation: Option<Box<crate::ExtmsgConversationRef>>,
    #[serde(rename="metadata", skip_serializing_if = "Option::is_none")]
    pub metadata: Option<std::collections::HashMap<String, String>>,
    #[serde(rename="session_id", skip_serializing_if = "Option::is_none")]
    pub session_id: Option<String>,
    #[serde(rename="additionalProperties", skip_serializing_if = "Option::is_none")]
    pub additional_properties: Option<std::collections::HashMap<String, serde_json::Value>>,
}
