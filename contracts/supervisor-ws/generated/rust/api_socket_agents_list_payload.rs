// ApiSocketAgentsListPayload represents a ApiSocketAgentsListPayload model.
#[derive(Clone, Debug, Deserialize, Serialize)]
pub struct ApiSocketAgentsListPayload {
    #[serde(rename="peek", skip_serializing_if = "Option::is_none")]
    pub peek: Option<bool>,
    #[serde(rename="pool", skip_serializing_if = "Option::is_none")]
    pub pool: Option<String>,
    #[serde(rename="rig", skip_serializing_if = "Option::is_none")]
    pub rig: Option<String>,
    #[serde(rename="running", skip_serializing_if = "Option::is_none")]
    pub running: Option<String>,
    #[serde(rename="additionalProperties", skip_serializing_if = "Option::is_none")]
    pub additional_properties: Option<std::collections::HashMap<String, serde_json::Value>>,
}
