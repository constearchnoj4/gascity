// ExtmsgExternalActor represents a ExtmsgExternalActor model.
#[derive(Clone, Debug, Deserialize, Serialize)]
pub struct ExtmsgExternalActor {
    #[serde(rename="display_name", skip_serializing_if = "Option::is_none")]
    pub display_name: Option<String>,
    #[serde(rename="id", skip_serializing_if = "Option::is_none")]
    pub id: Option<String>,
    #[serde(rename="is_bot", skip_serializing_if = "Option::is_none")]
    pub is_bot: Option<bool>,
    #[serde(rename="additionalProperties", skip_serializing_if = "Option::is_none")]
    pub additional_properties: Option<std::collections::HashMap<String, serde_json::Value>>,
}
