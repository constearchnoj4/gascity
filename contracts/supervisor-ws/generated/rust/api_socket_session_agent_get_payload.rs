// ApiSocketSessionAgentGetPayload represents a ApiSocketSessionAgentGetPayload model.
#[derive(Clone, Debug, Deserialize, Serialize)]
pub struct ApiSocketSessionAgentGetPayload {
    #[serde(rename="agent_id", skip_serializing_if = "Option::is_none")]
    pub agent_id: Option<String>,
    #[serde(rename="id", skip_serializing_if = "Option::is_none")]
    pub id: Option<String>,
    #[serde(rename="additionalProperties", skip_serializing_if = "Option::is_none")]
    pub additional_properties: Option<std::collections::HashMap<String, serde_json::Value>>,
}
