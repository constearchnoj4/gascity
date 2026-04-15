// ApiSocketSessionRenamePayload represents a ApiSocketSessionRenamePayload model.
#[derive(Clone, Debug, Deserialize, Serialize)]
pub struct ApiSocketSessionRenamePayload {
    #[serde(rename="id", skip_serializing_if = "Option::is_none")]
    pub id: Option<String>,
    #[serde(rename="title", skip_serializing_if = "Option::is_none")]
    pub title: Option<String>,
    #[serde(rename="additionalProperties", skip_serializing_if = "Option::is_none")]
    pub additional_properties: Option<std::collections::HashMap<String, serde_json::Value>>,
}
