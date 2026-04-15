// ApiWorkflowDepResponse represents a ApiWorkflowDepResponse model.
#[derive(Clone, Debug, Deserialize, Serialize)]
pub struct ApiWorkflowDepResponse {
    #[serde(rename="from", skip_serializing_if = "Option::is_none")]
    pub from: Option<String>,
    #[serde(rename="kind", skip_serializing_if = "Option::is_none")]
    pub kind: Option<String>,
    #[serde(rename="to", skip_serializing_if = "Option::is_none")]
    pub to: Option<String>,
    #[serde(rename="additionalProperties", skip_serializing_if = "Option::is_none")]
    pub additional_properties: Option<std::collections::HashMap<String, serde_json::Value>>,
}
