// ApiSocketMailCountPayload represents a ApiSocketMailCountPayload model.
#[derive(Clone, Debug, Deserialize, Serialize)]
pub struct ApiSocketMailCountPayload {
    #[serde(rename="agent", skip_serializing_if = "Option::is_none")]
    pub agent: Option<String>,
    #[serde(rename="rig", skip_serializing_if = "Option::is_none")]
    pub rig: Option<String>,
    #[serde(rename="additionalProperties", skip_serializing_if = "Option::is_none")]
    pub additional_properties: Option<std::collections::HashMap<String, serde_json::Value>>,
}
