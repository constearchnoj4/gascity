// ApiSessionResponse represents a ApiSessionResponse model.
#[derive(Clone, Debug, Deserialize, Serialize)]
pub struct ApiSessionResponse {
    #[serde(rename="active_bead", skip_serializing_if = "Option::is_none")]
    pub active_bead: Option<String>,
    #[serde(rename="activity", skip_serializing_if = "Option::is_none")]
    pub activity: Option<String>,
    #[serde(rename="alias", skip_serializing_if = "Option::is_none")]
    pub alias: Option<String>,
    #[serde(rename="attached", skip_serializing_if = "Option::is_none")]
    pub attached: Option<bool>,
    #[serde(rename="configured_named_session", skip_serializing_if = "Option::is_none")]
    pub configured_named_session: Option<bool>,
    #[serde(rename="context_pct", skip_serializing_if = "Option::is_none")]
    pub context_pct: Option<i32>,
    #[serde(rename="context_window", skip_serializing_if = "Option::is_none")]
    pub context_window: Option<i32>,
    #[serde(rename="created_at", skip_serializing_if = "Option::is_none")]
    pub created_at: Option<String>,
    #[serde(rename="display_name", skip_serializing_if = "Option::is_none")]
    pub display_name: Option<String>,
    #[serde(rename="id", skip_serializing_if = "Option::is_none")]
    pub id: Option<String>,
    #[serde(rename="kind", skip_serializing_if = "Option::is_none")]
    pub kind: Option<String>,
    #[serde(rename="last_active", skip_serializing_if = "Option::is_none")]
    pub last_active: Option<String>,
    #[serde(rename="last_output", skip_serializing_if = "Option::is_none")]
    pub last_output: Option<String>,
    #[serde(rename="metadata", skip_serializing_if = "Option::is_none")]
    pub metadata: Option<std::collections::HashMap<String, String>>,
    #[serde(rename="model", skip_serializing_if = "Option::is_none")]
    pub model: Option<String>,
    #[serde(rename="options", skip_serializing_if = "Option::is_none")]
    pub options: Option<std::collections::HashMap<String, String>>,
    #[serde(rename="pool", skip_serializing_if = "Option::is_none")]
    pub pool: Option<String>,
    #[serde(rename="provider", skip_serializing_if = "Option::is_none")]
    pub provider: Option<String>,
    #[serde(rename="reason", skip_serializing_if = "Option::is_none")]
    pub reason: Option<String>,
    #[serde(rename="rig", skip_serializing_if = "Option::is_none")]
    pub rig: Option<String>,
    #[serde(rename="running", skip_serializing_if = "Option::is_none")]
    pub running: Option<bool>,
    #[serde(rename="session_name", skip_serializing_if = "Option::is_none")]
    pub session_name: Option<String>,
    #[serde(rename="state", skip_serializing_if = "Option::is_none")]
    pub state: Option<String>,
    #[serde(rename="submission_capabilities", skip_serializing_if = "Option::is_none")]
    pub submission_capabilities: Option<Box<crate::SessionSubmissionCapabilities>>,
    #[serde(rename="template", skip_serializing_if = "Option::is_none")]
    pub template: Option<String>,
    #[serde(rename="title", skip_serializing_if = "Option::is_none")]
    pub title: Option<String>,
    #[serde(rename="additionalProperties", skip_serializing_if = "Option::is_none")]
    pub additional_properties: Option<std::collections::HashMap<String, serde_json::Value>>,
}
