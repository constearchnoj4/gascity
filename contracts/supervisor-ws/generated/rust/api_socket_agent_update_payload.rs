// ApiSocketAgentUpdatePayload represents a ApiSocketAgentUpdatePayload model.
#[derive(Clone, Debug, Deserialize, Serialize)]
pub struct ApiSocketAgentUpdatePayload {
    #[serde(rename="name", skip_serializing_if = "Option::is_none")]
    pub name: Option<String>,
    #[serde(rename="provider", skip_serializing_if = "Option::is_none")]
    pub provider: Option<String>,
    #[serde(rename="scope", skip_serializing_if = "Option::is_none")]
    pub scope: Option<String>,
    #[serde(rename="suspended", skip_serializing_if = "Option::is_none")]
    pub suspended: Option<bool>,
    #[serde(rename="additionalProperties", skip_serializing_if = "Option::is_none")]
    pub additional_properties: Option<std::collections::HashMap<String, serde_json::Value>>,
}
