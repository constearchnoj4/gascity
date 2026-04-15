// ApiAgentCreateRequest represents a ApiAgentCreateRequest model.
#[derive(Clone, Debug, Deserialize, Serialize)]
pub struct ApiAgentCreateRequest {
    #[serde(rename="dir", skip_serializing_if = "Option::is_none")]
    pub dir: Option<String>,
    #[serde(rename="name", skip_serializing_if = "Option::is_none")]
    pub name: Option<String>,
    #[serde(rename="provider", skip_serializing_if = "Option::is_none")]
    pub provider: Option<String>,
    #[serde(rename="scope", skip_serializing_if = "Option::is_none")]
    pub scope: Option<String>,
    #[serde(rename="additionalProperties", skip_serializing_if = "Option::is_none")]
    pub additional_properties: Option<std::collections::HashMap<String, serde_json::Value>>,
}
