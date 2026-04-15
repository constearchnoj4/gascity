// ApiSocketProviderCreatePayload represents a ApiSocketProviderCreatePayload model.
#[derive(Clone, Debug, Deserialize, Serialize)]
pub struct ApiSocketProviderCreatePayload {
    #[serde(rename="name", skip_serializing_if = "Option::is_none")]
    pub name: Option<String>,
    #[serde(rename="spec", skip_serializing_if = "Option::is_none")]
    pub spec: Option<std::collections::HashMap<String, serde_json::Value>>,
    #[serde(rename="additionalProperties", skip_serializing_if = "Option::is_none")]
    pub additional_properties: Option<std::collections::HashMap<String, serde_json::Value>>,
}
