// ApiSocketRigUpdatePayload represents a ApiSocketRigUpdatePayload model.
#[derive(Clone, Debug, Deserialize, Serialize)]
pub struct ApiSocketRigUpdatePayload {
    #[serde(rename="name", skip_serializing_if = "Option::is_none")]
    pub name: Option<String>,
    #[serde(rename="path", skip_serializing_if = "Option::is_none")]
    pub path: Option<String>,
    #[serde(rename="prefix", skip_serializing_if = "Option::is_none")]
    pub prefix: Option<String>,
    #[serde(rename="suspended", skip_serializing_if = "Option::is_none")]
    pub suspended: Option<bool>,
    #[serde(rename="additionalProperties", skip_serializing_if = "Option::is_none")]
    pub additional_properties: Option<std::collections::HashMap<String, serde_json::Value>>,
}
