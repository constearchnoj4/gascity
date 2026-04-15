// ApiSocketMailListPayload represents a ApiSocketMailListPayload model.
#[derive(Clone, Debug, Deserialize, Serialize)]
pub struct ApiSocketMailListPayload {
    #[serde(rename="agent", skip_serializing_if = "Option::is_none")]
    pub agent: Option<String>,
    #[serde(rename="cursor", skip_serializing_if = "Option::is_none")]
    pub cursor: Option<String>,
    #[serde(rename="limit", skip_serializing_if = "Option::is_none")]
    pub limit: Option<i32>,
    #[serde(rename="rig", skip_serializing_if = "Option::is_none")]
    pub rig: Option<String>,
    #[serde(rename="status", skip_serializing_if = "Option::is_none")]
    pub status: Option<String>,
    #[serde(rename="additionalProperties", skip_serializing_if = "Option::is_none")]
    pub additional_properties: Option<std::collections::HashMap<String, serde_json::Value>>,
}
