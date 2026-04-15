// ExtmsgExternalAttachment represents a ExtmsgExternalAttachment model.
#[derive(Clone, Debug, Deserialize, Serialize)]
pub struct ExtmsgExternalAttachment {
    #[serde(rename="mime_type", skip_serializing_if = "Option::is_none")]
    pub mime_type: Option<String>,
    #[serde(rename="provider_id", skip_serializing_if = "Option::is_none")]
    pub provider_id: Option<String>,
    #[serde(rename="url", skip_serializing_if = "Option::is_none")]
    pub url: Option<String>,
    #[serde(rename="additionalProperties", skip_serializing_if = "Option::is_none")]
    pub additional_properties: Option<std::collections::HashMap<String, serde_json::Value>>,
}
