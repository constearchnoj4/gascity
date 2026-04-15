// SessionSubmissionCapabilities represents a SessionSubmissionCapabilities model.
#[derive(Clone, Debug, Deserialize, Serialize)]
pub struct SessionSubmissionCapabilities {
    #[serde(rename="supports_follow_up", skip_serializing_if = "Option::is_none")]
    pub supports_follow_up: Option<bool>,
    #[serde(rename="supports_interrupt_now", skip_serializing_if = "Option::is_none")]
    pub supports_interrupt_now: Option<bool>,
    #[serde(rename="additionalProperties", skip_serializing_if = "Option::is_none")]
    pub additional_properties: Option<std::collections::HashMap<String, serde_json::Value>>,
}
