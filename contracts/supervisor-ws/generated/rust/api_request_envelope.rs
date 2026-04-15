// ApiRequestEnvelope represents a ApiRequestEnvelope model.
#[derive(Clone, Debug, Deserialize, Serialize)]
pub struct ApiRequestEnvelope {
    #[serde(rename="action", skip_serializing_if = "Option::is_none")]
    pub action: Option<String>,
    #[serde(rename="id", skip_serializing_if = "Option::is_none")]
    pub id: Option<String>,
    #[serde(rename="idempotency_key", skip_serializing_if = "Option::is_none")]
    pub idempotency_key: Option<String>,
    #[serde(rename="payload", skip_serializing_if = "Option::is_none")]
    pub payload: Option<serde_json::Value>,
    #[serde(rename="scope", skip_serializing_if = "Option::is_none")]
    pub scope: Option<Box<crate::ApiScope>>,
    #[serde(rename="type", skip_serializing_if = "Option::is_none")]
    pub reserved_type: Option<String>,
    #[serde(rename="watch", skip_serializing_if = "Option::is_none")]
    pub watch: Option<Box<crate::ApiWatchParams>>,
    #[serde(rename="additionalProperties", skip_serializing_if = "Option::is_none")]
    pub additional_properties: Option<std::collections::HashMap<String, serde_json::Value>>,
}
