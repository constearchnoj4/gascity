// ApiExtmsgParticipantUpsertRequest represents a ApiExtmsgParticipantUpsertRequest model.
#[derive(Clone, Debug, Deserialize, Serialize)]
pub struct ApiExtmsgParticipantUpsertRequest {
    #[serde(rename="group_id", skip_serializing_if = "Option::is_none")]
    pub group_id: Option<String>,
    #[serde(rename="handle", skip_serializing_if = "Option::is_none")]
    pub handle: Option<String>,
    #[serde(rename="metadata", skip_serializing_if = "Option::is_none")]
    pub metadata: Option<std::collections::HashMap<String, String>>,
    #[serde(rename="public", skip_serializing_if = "Option::is_none")]
    pub public: Option<bool>,
    #[serde(rename="session_id", skip_serializing_if = "Option::is_none")]
    pub session_id: Option<String>,
    #[serde(rename="additionalProperties", skip_serializing_if = "Option::is_none")]
    pub additional_properties: Option<std::collections::HashMap<String, serde_json::Value>>,
}
