// ApiBeadCreateRequest represents a ApiBeadCreateRequest model.
#[derive(Clone, Debug, Deserialize, Serialize)]
pub struct ApiBeadCreateRequest {
    #[serde(rename="assignee", skip_serializing_if = "Option::is_none")]
    pub assignee: Option<String>,
    #[serde(rename="description", skip_serializing_if = "Option::is_none")]
    pub description: Option<String>,
    #[serde(rename="labels", skip_serializing_if = "Option::is_none")]
    pub labels: Option<Vec<String>>,
    #[serde(rename="priority", skip_serializing_if = "Option::is_none")]
    pub priority: Option<i32>,
    #[serde(rename="rig", skip_serializing_if = "Option::is_none")]
    pub rig: Option<String>,
    #[serde(rename="title", skip_serializing_if = "Option::is_none")]
    pub title: Option<String>,
    #[serde(rename="type", skip_serializing_if = "Option::is_none")]
    pub reserved_type: Option<String>,
    #[serde(rename="additionalProperties", skip_serializing_if = "Option::is_none")]
    pub additional_properties: Option<std::collections::HashMap<String, serde_json::Value>>,
}
