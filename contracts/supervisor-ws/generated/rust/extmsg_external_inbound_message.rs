// ExtmsgExternalInboundMessage represents a ExtmsgExternalInboundMessage model.
#[derive(Clone, Debug, Deserialize, Serialize)]
pub struct ExtmsgExternalInboundMessage {
    #[serde(rename="actor", skip_serializing_if = "Option::is_none")]
    pub actor: Option<Box<crate::ExtmsgExternalActor>>,
    #[serde(rename="attachments", skip_serializing_if = "Option::is_none")]
    pub attachments: Option<Vec<crate::ExtmsgExternalAttachment>>,
    #[serde(rename="conversation", skip_serializing_if = "Option::is_none")]
    pub conversation: Option<Box<crate::ExtmsgConversationRef>>,
    #[serde(rename="dedup_key", skip_serializing_if = "Option::is_none")]
    pub dedup_key: Option<String>,
    #[serde(rename="explicit_target", skip_serializing_if = "Option::is_none")]
    pub explicit_target: Option<String>,
    #[serde(rename="provider_message_id", skip_serializing_if = "Option::is_none")]
    pub provider_message_id: Option<String>,
    #[serde(rename="received_at", skip_serializing_if = "Option::is_none")]
    pub received_at: Option<String>,
    #[serde(rename="reply_to_message_id", skip_serializing_if = "Option::is_none")]
    pub reply_to_message_id: Option<String>,
    #[serde(rename="text", skip_serializing_if = "Option::is_none")]
    pub text: Option<String>,
    #[serde(rename="additionalProperties", skip_serializing_if = "Option::is_none")]
    pub additional_properties: Option<std::collections::HashMap<String, serde_json::Value>>,
}
