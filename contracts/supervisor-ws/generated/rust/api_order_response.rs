// ApiOrderResponse represents a ApiOrderResponse model.
#[derive(Clone, Debug, Deserialize, Serialize)]
pub struct ApiOrderResponse {
    #[serde(rename="capture_output", skip_serializing_if = "Option::is_none")]
    pub capture_output: Option<bool>,
    #[serde(rename="check", skip_serializing_if = "Option::is_none")]
    pub check: Option<String>,
    #[serde(rename="description", skip_serializing_if = "Option::is_none")]
    pub description: Option<String>,
    #[serde(rename="enabled", skip_serializing_if = "Option::is_none")]
    pub enabled: Option<bool>,
    #[serde(rename="exec", skip_serializing_if = "Option::is_none")]
    pub exec: Option<String>,
    #[serde(rename="formula", skip_serializing_if = "Option::is_none")]
    pub formula: Option<String>,
    #[serde(rename="gate", skip_serializing_if = "Option::is_none")]
    pub gate: Option<String>,
    #[serde(rename="interval", skip_serializing_if = "Option::is_none")]
    pub interval: Option<String>,
    #[serde(rename="name", skip_serializing_if = "Option::is_none")]
    pub name: Option<String>,
    #[serde(rename="on", skip_serializing_if = "Option::is_none")]
    pub on: Option<String>,
    #[serde(rename="pool", skip_serializing_if = "Option::is_none")]
    pub pool: Option<String>,
    #[serde(rename="rig", skip_serializing_if = "Option::is_none")]
    pub rig: Option<String>,
    #[serde(rename="schedule", skip_serializing_if = "Option::is_none")]
    pub schedule: Option<String>,
    #[serde(rename="scoped_name", skip_serializing_if = "Option::is_none")]
    pub scoped_name: Option<String>,
    #[serde(rename="timeout", skip_serializing_if = "Option::is_none")]
    pub timeout: Option<String>,
    #[serde(rename="timeout_ms", skip_serializing_if = "Option::is_none")]
    pub timeout_ms: Option<i32>,
    #[serde(rename="type", skip_serializing_if = "Option::is_none")]
    pub reserved_type: Option<String>,
    #[serde(rename="additionalProperties", skip_serializing_if = "Option::is_none")]
    pub additional_properties: Option<std::collections::HashMap<String, serde_json::Value>>,
}
