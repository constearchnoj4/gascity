// ApiListResponse represents a ApiListResponse model.
#[derive(Clone, Debug, Deserialize, Serialize)]
pub struct ApiListResponse {
    #[serde(rename="items", skip_serializing_if = "Option::is_none")]
    pub items: Option<serde_json::Value>,
    #[serde(rename="next_cursor", skip_serializing_if = "Option::is_none")]
    pub next_cursor: Option<String>,
    #[serde(rename="total", skip_serializing_if = "Option::is_none")]
    pub total: Option<i32>,
    #[serde(rename="additionalProperties", skip_serializing_if = "Option::is_none")]
    pub additional_properties: Option<std::collections::HashMap<String, serde_json::Value>>,
}
