// ApiSocketMailGetPayload represents a ApiSocketMailGetPayload model.
#[derive(Clone, Debug, Deserialize, Serialize)]
pub struct ApiSocketMailGetPayload {
    #[serde(rename="id", skip_serializing_if = "Option::is_none")]
    pub id: Option<String>,
    #[serde(rename="rig", skip_serializing_if = "Option::is_none")]
    pub rig: Option<String>,
    #[serde(rename="additionalProperties", skip_serializing_if = "Option::is_none")]
    pub additional_properties: Option<std::collections::HashMap<String, serde_json::Value>>,
}
