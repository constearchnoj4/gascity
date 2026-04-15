// ApiSocketProviderUpdatePayload represents a ApiSocketProviderUpdatePayload model.
#[derive(Clone, Debug, Deserialize, Serialize)]
pub struct ApiSocketProviderUpdatePayload {
    #[serde(rename="name", skip_serializing_if = "Option::is_none")]
    pub name: Option<String>,
    #[serde(rename="update", skip_serializing_if = "Option::is_none")]
    pub update: Option<std::collections::HashMap<String, serde_json::Value>>,
    #[serde(rename="additionalProperties", skip_serializing_if = "Option::is_none")]
    pub additional_properties: Option<std::collections::HashMap<String, serde_json::Value>>,
}
