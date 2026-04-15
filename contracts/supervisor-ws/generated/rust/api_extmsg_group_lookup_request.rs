// ApiExtmsgGroupLookupRequest represents a ApiExtmsgGroupLookupRequest model.
#[derive(Clone, Debug, Deserialize, Serialize)]
pub struct ApiExtmsgGroupLookupRequest {
    #[serde(rename="account_id", skip_serializing_if = "Option::is_none")]
    pub account_id: Option<String>,
    #[serde(rename="conversation_id", skip_serializing_if = "Option::is_none")]
    pub conversation_id: Option<String>,
    #[serde(rename="kind", skip_serializing_if = "Option::is_none")]
    pub kind: Option<String>,
    #[serde(rename="provider", skip_serializing_if = "Option::is_none")]
    pub provider: Option<String>,
    #[serde(rename="scope_id", skip_serializing_if = "Option::is_none")]
    pub scope_id: Option<String>,
    #[serde(rename="additionalProperties", skip_serializing_if = "Option::is_none")]
    pub additional_properties: Option<std::collections::HashMap<String, serde_json::Value>>,
}
