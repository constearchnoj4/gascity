// ApiExtmsgParticipantRemoveRequest represents a ApiExtmsgParticipantRemoveRequest model.
#[derive(Clone, Debug, Deserialize, Serialize)]
pub struct ApiExtmsgParticipantRemoveRequest {
    #[serde(rename="group_id", skip_serializing_if = "Option::is_none")]
    pub group_id: Option<String>,
    #[serde(rename="handle", skip_serializing_if = "Option::is_none")]
    pub handle: Option<String>,
    #[serde(rename="additionalProperties", skip_serializing_if = "Option::is_none")]
    pub additional_properties: Option<std::collections::HashMap<String, serde_json::Value>>,
}
