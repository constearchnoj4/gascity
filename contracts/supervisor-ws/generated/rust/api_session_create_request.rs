// ApiSessionCreateRequest represents a ApiSessionCreateRequest model.
#[derive(Clone, Debug, Deserialize, Serialize)]
pub struct ApiSessionCreateRequest {
    #[serde(rename="alias", skip_serializing_if = "Option::is_none")]
    pub alias: Option<String>,
    #[serde(rename="async", skip_serializing_if = "Option::is_none")]
    pub reserved_async: Option<bool>,
    #[serde(rename="kind", skip_serializing_if = "Option::is_none")]
    pub kind: Option<String>,
    #[serde(rename="message", skip_serializing_if = "Option::is_none")]
    pub message: Option<String>,
    #[serde(rename="name", skip_serializing_if = "Option::is_none")]
    pub name: Option<String>,
    #[serde(rename="options", skip_serializing_if = "Option::is_none")]
    pub options: Option<std::collections::HashMap<String, String>>,
    #[serde(rename="project_id", skip_serializing_if = "Option::is_none")]
    pub project_id: Option<String>,
    #[serde(rename="session_name", skip_serializing_if = "Option::is_none")]
    pub session_name: Option<String>,
    #[serde(rename="title", skip_serializing_if = "Option::is_none")]
    pub title: Option<String>,
    #[serde(rename="additionalProperties", skip_serializing_if = "Option::is_none")]
    pub additional_properties: Option<std::collections::HashMap<String, serde_json::Value>>,
}
