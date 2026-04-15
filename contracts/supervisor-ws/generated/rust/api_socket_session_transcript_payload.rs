// ApiSocketSessionTranscriptPayload represents a ApiSocketSessionTranscriptPayload model.
#[derive(Clone, Debug, Deserialize, Serialize)]
pub struct ApiSocketSessionTranscriptPayload {
    #[serde(rename="before", skip_serializing_if = "Option::is_none")]
    pub before: Option<String>,
    #[serde(rename="format", skip_serializing_if = "Option::is_none")]
    pub format: Option<String>,
    #[serde(rename="id", skip_serializing_if = "Option::is_none")]
    pub id: Option<String>,
    #[serde(rename="turns", skip_serializing_if = "Option::is_none")]
    pub turns: Option<i32>,
    #[serde(rename="additionalProperties", skip_serializing_if = "Option::is_none")]
    pub additional_properties: Option<std::collections::HashMap<String, serde_json::Value>>,
}
