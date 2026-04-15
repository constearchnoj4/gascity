// ApiSocketNamePayload represents a ApiSocketNamePayload model.
#[derive(Clone, Debug, Deserialize, Serialize)]
pub struct ApiSocketNamePayload {
    #[serde(rename="name", skip_serializing_if = "Option::is_none")]
    pub name: Option<String>,
    #[serde(rename="additionalProperties", skip_serializing_if = "Option::is_none")]
    pub additional_properties: Option<std::collections::HashMap<String, serde_json::Value>>,
}
