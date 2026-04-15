// ApiFieldError represents a ApiFieldError model.
#[derive(Clone, Debug, Deserialize, Serialize)]
pub struct ApiFieldError {
    #[serde(rename="field", skip_serializing_if = "Option::is_none")]
    pub field: Option<String>,
    #[serde(rename="message", skip_serializing_if = "Option::is_none")]
    pub message: Option<String>,
    #[serde(rename="additionalProperties", skip_serializing_if = "Option::is_none")]
    pub additional_properties: Option<std::collections::HashMap<String, serde_json::Value>>,
}
