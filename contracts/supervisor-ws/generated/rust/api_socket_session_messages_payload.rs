// ApiSocketSessionMessagesPayload represents a ApiSocketSessionMessagesPayload model.
#[derive(Clone, Debug, Deserialize, Serialize)]
pub struct ApiSocketSessionMessagesPayload {
    #[serde(rename="id", skip_serializing_if = "Option::is_none")]
    pub id: Option<String>,
    #[serde(rename="message", skip_serializing_if = "Option::is_none")]
    pub message: Option<String>,
    #[serde(rename="additionalProperties", skip_serializing_if = "Option::is_none")]
    pub additional_properties: Option<std::collections::HashMap<String, serde_json::Value>>,
}
