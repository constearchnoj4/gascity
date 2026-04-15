// ApiSocketBeadsListPayload represents a ApiSocketBeadsListPayload model.
#[derive(Clone, Debug, Deserialize, Serialize)]
pub struct ApiSocketBeadsListPayload {
    #[serde(rename="assignee", skip_serializing_if = "Option::is_none")]
    pub assignee: Option<String>,
    #[serde(rename="cursor", skip_serializing_if = "Option::is_none")]
    pub cursor: Option<String>,
    #[serde(rename="label", skip_serializing_if = "Option::is_none")]
    pub label: Option<String>,
    #[serde(rename="limit", skip_serializing_if = "Option::is_none")]
    pub limit: Option<i32>,
    #[serde(rename="rig", skip_serializing_if = "Option::is_none")]
    pub rig: Option<String>,
    #[serde(rename="status", skip_serializing_if = "Option::is_none")]
    pub status: Option<String>,
    #[serde(rename="type", skip_serializing_if = "Option::is_none")]
    pub reserved_type: Option<String>,
    #[serde(rename="additionalProperties", skip_serializing_if = "Option::is_none")]
    pub additional_properties: Option<std::collections::HashMap<String, serde_json::Value>>,
}
