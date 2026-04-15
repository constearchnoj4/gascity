// ApiSocketWorkflowDeletePayload represents a ApiSocketWorkflowDeletePayload model.
#[derive(Clone, Debug, Deserialize, Serialize)]
pub struct ApiSocketWorkflowDeletePayload {
    #[serde(rename="delete", skip_serializing_if = "Option::is_none")]
    pub delete: Option<bool>,
    #[serde(rename="id", skip_serializing_if = "Option::is_none")]
    pub id: Option<String>,
    #[serde(rename="scope_kind", skip_serializing_if = "Option::is_none")]
    pub scope_kind: Option<String>,
    #[serde(rename="scope_ref", skip_serializing_if = "Option::is_none")]
    pub scope_ref: Option<String>,
    #[serde(rename="additionalProperties", skip_serializing_if = "Option::is_none")]
    pub additional_properties: Option<std::collections::HashMap<String, serde_json::Value>>,
}
