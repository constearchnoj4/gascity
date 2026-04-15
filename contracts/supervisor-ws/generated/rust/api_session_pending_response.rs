// ApiSessionPendingResponse represents a ApiSessionPendingResponse model.
#[derive(Clone, Debug, Deserialize, Serialize)]
pub struct ApiSessionPendingResponse {
    #[serde(rename="pending", skip_serializing_if = "Option::is_none")]
    pub pending: Option<Box<crate::RuntimePendingInteraction>>,
    #[serde(rename="supported", skip_serializing_if = "Option::is_none")]
    pub supported: Option<bool>,
    #[serde(rename="additionalProperties", skip_serializing_if = "Option::is_none")]
    pub additional_properties: Option<std::collections::HashMap<String, serde_json::Value>>,
}
