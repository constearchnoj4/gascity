// ApiExtmsgAdapterRegisterRequest represents a ApiExtmsgAdapterRegisterRequest model.
#[derive(Clone, Debug, Deserialize, Serialize)]
pub struct ApiExtmsgAdapterRegisterRequest {
    #[serde(rename="account_id", skip_serializing_if = "Option::is_none")]
    pub account_id: Option<String>,
    #[serde(rename="callback_url", skip_serializing_if = "Option::is_none")]
    pub callback_url: Option<String>,
    #[serde(rename="capabilities", skip_serializing_if = "Option::is_none")]
    pub capabilities: Option<std::collections::HashMap<String, serde_json::Value>>,
    #[serde(rename="name", skip_serializing_if = "Option::is_none")]
    pub name: Option<String>,
    #[serde(rename="provider", skip_serializing_if = "Option::is_none")]
    pub provider: Option<String>,
    #[serde(rename="additionalProperties", skip_serializing_if = "Option::is_none")]
    pub additional_properties: Option<std::collections::HashMap<String, serde_json::Value>>,
}
