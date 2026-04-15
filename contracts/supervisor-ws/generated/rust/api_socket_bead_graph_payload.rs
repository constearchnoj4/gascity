// ApiSocketBeadGraphPayload represents a ApiSocketBeadGraphPayload model.
#[derive(Clone, Debug, Deserialize, Serialize)]
pub struct ApiSocketBeadGraphPayload {
    #[serde(rename="root_id", skip_serializing_if = "Option::is_none")]
    pub root_id: Option<String>,
    #[serde(rename="additionalProperties", skip_serializing_if = "Option::is_none")]
    pub additional_properties: Option<std::collections::HashMap<String, serde_json::Value>>,
}
