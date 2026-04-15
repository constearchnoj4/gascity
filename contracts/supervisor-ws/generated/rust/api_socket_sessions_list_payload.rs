// ApiSocketSessionsListPayload represents a ApiSocketSessionsListPayload model.
#[derive(Clone, Debug, Deserialize, Serialize)]
pub struct ApiSocketSessionsListPayload {
    #[serde(rename="cursor", skip_serializing_if = "Option::is_none")]
    pub cursor: Option<String>,
    #[serde(rename="limit", skip_serializing_if = "Option::is_none")]
    pub limit: Option<i32>,
    #[serde(rename="peek", skip_serializing_if = "Option::is_none")]
    pub peek: Option<bool>,
    #[serde(rename="state", skip_serializing_if = "Option::is_none")]
    pub state: Option<String>,
    #[serde(rename="template", skip_serializing_if = "Option::is_none")]
    pub template: Option<String>,
    #[serde(rename="additionalProperties", skip_serializing_if = "Option::is_none")]
    pub additional_properties: Option<std::collections::HashMap<String, serde_json::Value>>,
}
