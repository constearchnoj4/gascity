// ApiResponseEnvelope represents a ApiResponseEnvelope model.
#[derive(Clone, Debug, Deserialize, Serialize)]
pub struct ApiResponseEnvelope {
    #[serde(rename="id", skip_serializing_if = "Option::is_none")]
    pub id: Option<String>,
    #[serde(rename="index", skip_serializing_if = "Option::is_none")]
    pub index: Option<i32>,
    #[serde(rename="result", skip_serializing_if = "Option::is_none")]
    pub result: Option<serde_json::Value>,
    #[serde(rename="type", skip_serializing_if = "Option::is_none")]
    pub reserved_type: Option<String>,
    #[serde(rename="additionalProperties", skip_serializing_if = "Option::is_none")]
    pub additional_properties: Option<std::collections::HashMap<String, serde_json::Value>>,
}
