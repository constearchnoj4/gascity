// ApiExtmsgGroupEnsureRequest represents a ApiExtmsgGroupEnsureRequest model.
#[derive(Clone, Debug, Deserialize, Serialize)]
pub struct ApiExtmsgGroupEnsureRequest {
    #[serde(rename="default_handle", skip_serializing_if = "Option::is_none")]
    pub default_handle: Option<String>,
    #[serde(rename="metadata", skip_serializing_if = "Option::is_none")]
    pub metadata: Option<std::collections::HashMap<String, String>>,
    #[serde(rename="mode", skip_serializing_if = "Option::is_none")]
    pub mode: Option<String>,
    #[serde(rename="root_conversation", skip_serializing_if = "Option::is_none")]
    pub root_conversation: Option<Box<crate::ExtmsgConversationRef>>,
    #[serde(rename="additionalProperties", skip_serializing_if = "Option::is_none")]
    pub additional_properties: Option<std::collections::HashMap<String, serde_json::Value>>,
}
