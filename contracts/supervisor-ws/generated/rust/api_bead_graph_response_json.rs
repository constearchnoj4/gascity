// ApiBeadGraphResponseJson represents a ApiBeadGraphResponseJson model.
#[derive(Clone, Debug, Deserialize, Serialize)]
pub struct ApiBeadGraphResponseJson {
    #[serde(rename="beads", skip_serializing_if = "Option::is_none")]
    pub beads: Option<Vec<crate::BeadsBead>>,
    #[serde(rename="deps", skip_serializing_if = "Option::is_none")]
    pub deps: Option<Vec<crate::ApiWorkflowDepResponse>>,
    #[serde(rename="root", skip_serializing_if = "Option::is_none")]
    pub root: Option<Box<crate::BeadsBead>>,
    #[serde(rename="additionalProperties", skip_serializing_if = "Option::is_none")]
    pub additional_properties: Option<std::collections::HashMap<String, serde_json::Value>>,
}
