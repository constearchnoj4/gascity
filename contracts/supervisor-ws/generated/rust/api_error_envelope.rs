// ApiErrorEnvelope represents a ApiErrorEnvelope model.
#[derive(Clone, Debug, Deserialize, Serialize)]
pub struct ApiErrorEnvelope {
    #[serde(rename="code", skip_serializing_if = "Option::is_none")]
    pub code: Option<String>,
    #[serde(rename="details", skip_serializing_if = "Option::is_none")]
    pub details: Option<Vec<crate::ApiFieldError>>,
    #[serde(rename="id", skip_serializing_if = "Option::is_none")]
    pub id: Option<String>,
    #[serde(rename="message", skip_serializing_if = "Option::is_none")]
    pub message: Option<String>,
    #[serde(rename="type", skip_serializing_if = "Option::is_none")]
    pub reserved_type: Option<String>,
    #[serde(rename="additionalProperties", skip_serializing_if = "Option::is_none")]
    pub additional_properties: Option<std::collections::HashMap<String, serde_json::Value>>,
}
