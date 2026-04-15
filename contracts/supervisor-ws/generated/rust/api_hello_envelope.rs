// ApiHelloEnvelope represents a ApiHelloEnvelope model.
#[derive(Clone, Debug, Deserialize, Serialize)]
pub struct ApiHelloEnvelope {
    #[serde(rename="capabilities", skip_serializing_if = "Option::is_none")]
    pub capabilities: Option<Vec<String>>,
    #[serde(rename="protocol", skip_serializing_if = "Option::is_none")]
    pub protocol: Option<String>,
    #[serde(rename="read_only", skip_serializing_if = "Option::is_none")]
    pub read_only: Option<bool>,
    #[serde(rename="server_role", skip_serializing_if = "Option::is_none")]
    pub server_role: Option<String>,
    #[serde(rename="subscription_kinds", skip_serializing_if = "Option::is_none")]
    pub subscription_kinds: Option<Vec<String>>,
    #[serde(rename="type", skip_serializing_if = "Option::is_none")]
    pub reserved_type: Option<String>,
    #[serde(rename="additionalProperties", skip_serializing_if = "Option::is_none")]
    pub additional_properties: Option<std::collections::HashMap<String, serde_json::Value>>,
}
