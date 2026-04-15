// ApiSocketSessionRespondPayload represents a ApiSocketSessionRespondPayload model.
#[derive(Clone, Debug, Deserialize, Serialize)]
pub struct ApiSocketSessionRespondPayload {
    #[serde(rename="action", skip_serializing_if = "Option::is_none")]
    pub action: Option<String>,
    #[serde(rename="id", skip_serializing_if = "Option::is_none")]
    pub id: Option<String>,
    #[serde(rename="metadata", skip_serializing_if = "Option::is_none")]
    pub metadata: Option<std::collections::HashMap<String, String>>,
    #[serde(rename="request_id", skip_serializing_if = "Option::is_none")]
    pub request_id: Option<String>,
    #[serde(rename="text", skip_serializing_if = "Option::is_none")]
    pub text: Option<String>,
    #[serde(rename="additionalProperties", skip_serializing_if = "Option::is_none")]
    pub additional_properties: Option<std::collections::HashMap<String, serde_json::Value>>,
}
