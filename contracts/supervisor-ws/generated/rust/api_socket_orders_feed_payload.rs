// ApiSocketOrdersFeedPayload represents a ApiSocketOrdersFeedPayload model.
#[derive(Clone, Debug, Deserialize, Serialize)]
pub struct ApiSocketOrdersFeedPayload {
    #[serde(rename="limit", skip_serializing_if = "Option::is_none")]
    pub limit: Option<i32>,
    #[serde(rename="scope_kind", skip_serializing_if = "Option::is_none")]
    pub scope_kind: Option<String>,
    #[serde(rename="scope_ref", skip_serializing_if = "Option::is_none")]
    pub scope_ref: Option<String>,
    #[serde(rename="additionalProperties", skip_serializing_if = "Option::is_none")]
    pub additional_properties: Option<std::collections::HashMap<String, serde_json::Value>>,
}
