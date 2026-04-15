// ApiMailSendRequest represents a ApiMailSendRequest model.
#[derive(Clone, Debug, Deserialize, Serialize)]
pub struct ApiMailSendRequest {
    #[serde(rename="body", skip_serializing_if = "Option::is_none")]
    pub body: Option<String>,
    #[serde(rename="from", skip_serializing_if = "Option::is_none")]
    pub from: Option<String>,
    #[serde(rename="rig", skip_serializing_if = "Option::is_none")]
    pub rig: Option<String>,
    #[serde(rename="subject", skip_serializing_if = "Option::is_none")]
    pub subject: Option<String>,
    #[serde(rename="to", skip_serializing_if = "Option::is_none")]
    pub to: Option<String>,
    #[serde(rename="additionalProperties", skip_serializing_if = "Option::is_none")]
    pub additional_properties: Option<std::collections::HashMap<String, serde_json::Value>>,
}
