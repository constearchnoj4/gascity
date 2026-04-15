// ApiEventEmitRequest represents a ApiEventEmitRequest model.
#[derive(Clone, Debug, Deserialize, Serialize)]
pub struct ApiEventEmitRequest {
    #[serde(rename="actor", skip_serializing_if = "Option::is_none")]
    pub actor: Option<String>,
    #[serde(rename="message", skip_serializing_if = "Option::is_none")]
    pub message: Option<String>,
    #[serde(rename="subject", skip_serializing_if = "Option::is_none")]
    pub subject: Option<String>,
    #[serde(rename="type", skip_serializing_if = "Option::is_none")]
    pub reserved_type: Option<String>,
    #[serde(rename="additionalProperties", skip_serializing_if = "Option::is_none")]
    pub additional_properties: Option<std::collections::HashMap<String, serde_json::Value>>,
}
