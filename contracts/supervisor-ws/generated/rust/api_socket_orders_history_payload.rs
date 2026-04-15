// ApiSocketOrdersHistoryPayload represents a ApiSocketOrdersHistoryPayload model.
#[derive(Clone, Debug, Deserialize, Serialize)]
pub struct ApiSocketOrdersHistoryPayload {
    #[serde(rename="before", skip_serializing_if = "Option::is_none")]
    pub before: Option<String>,
    #[serde(rename="limit", skip_serializing_if = "Option::is_none")]
    pub limit: Option<i32>,
    #[serde(rename="scoped_name", skip_serializing_if = "Option::is_none")]
    pub scoped_name: Option<String>,
    #[serde(rename="additionalProperties", skip_serializing_if = "Option::is_none")]
    pub additional_properties: Option<std::collections::HashMap<String, serde_json::Value>>,
}
