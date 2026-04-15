// ApiSocketEventsListPayload represents a ApiSocketEventsListPayload model.
#[derive(Clone, Debug, Deserialize, Serialize)]
pub struct ApiSocketEventsListPayload {
    #[serde(rename="actor", skip_serializing_if = "Option::is_none")]
    pub actor: Option<String>,
    #[serde(rename="cursor", skip_serializing_if = "Option::is_none")]
    pub cursor: Option<String>,
    #[serde(rename="limit", skip_serializing_if = "Option::is_none")]
    pub limit: Option<i32>,
    #[serde(rename="since", skip_serializing_if = "Option::is_none")]
    pub since: Option<String>,
    #[serde(rename="type", skip_serializing_if = "Option::is_none")]
    pub reserved_type: Option<String>,
    #[serde(rename="additionalProperties", skip_serializing_if = "Option::is_none")]
    pub additional_properties: Option<std::collections::HashMap<String, serde_json::Value>>,
}
