// MailMessage represents a MailMessage model.
#[derive(Clone, Debug, Deserialize, Serialize)]
pub struct MailMessage {
    #[serde(rename="body", skip_serializing_if = "Option::is_none")]
    pub body: Option<String>,
    #[serde(rename="cc", skip_serializing_if = "Option::is_none")]
    pub cc: Option<Vec<String>>,
    #[serde(rename="created_at", skip_serializing_if = "Option::is_none")]
    pub created_at: Option<String>,
    #[serde(rename="from", skip_serializing_if = "Option::is_none")]
    pub from: Option<String>,
    #[serde(rename="id", skip_serializing_if = "Option::is_none")]
    pub id: Option<String>,
    #[serde(rename="priority", skip_serializing_if = "Option::is_none")]
    pub priority: Option<i32>,
    #[serde(rename="read", skip_serializing_if = "Option::is_none")]
    pub read: Option<bool>,
    #[serde(rename="reply_to", skip_serializing_if = "Option::is_none")]
    pub reply_to: Option<String>,
    #[serde(rename="rig", skip_serializing_if = "Option::is_none")]
    pub rig: Option<String>,
    #[serde(rename="subject", skip_serializing_if = "Option::is_none")]
    pub subject: Option<String>,
    #[serde(rename="thread_id", skip_serializing_if = "Option::is_none")]
    pub thread_id: Option<String>,
    #[serde(rename="to", skip_serializing_if = "Option::is_none")]
    pub to: Option<String>,
    #[serde(rename="additionalProperties", skip_serializing_if = "Option::is_none")]
    pub additional_properties: Option<std::collections::HashMap<String, serde_json::Value>>,
}
