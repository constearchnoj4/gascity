// ApiSocketFormulaGetPayload represents a ApiSocketFormulaGetPayload model.
#[derive(Clone, Debug, Deserialize, Serialize)]
pub struct ApiSocketFormulaGetPayload {
    #[serde(rename="name", skip_serializing_if = "Option::is_none")]
    pub name: Option<String>,
    #[serde(rename="scope_kind", skip_serializing_if = "Option::is_none")]
    pub scope_kind: Option<String>,
    #[serde(rename="scope_ref", skip_serializing_if = "Option::is_none")]
    pub scope_ref: Option<String>,
    #[serde(rename="target", skip_serializing_if = "Option::is_none")]
    pub target: Option<String>,
    #[serde(rename="vars", skip_serializing_if = "Option::is_none")]
    pub vars: Option<std::collections::HashMap<String, String>>,
    #[serde(rename="additionalProperties", skip_serializing_if = "Option::is_none")]
    pub additional_properties: Option<std::collections::HashMap<String, serde_json::Value>>,
}
