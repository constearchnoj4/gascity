// RuntimePendingInteraction represents a RuntimePendingInteraction model.
#[derive(Clone, Debug, Deserialize, Serialize)]
pub struct RuntimePendingInteraction {
    #[serde(rename="kind", skip_serializing_if = "Option::is_none")]
    pub kind: Option<String>,
    #[serde(rename="metadata", skip_serializing_if = "Option::is_none")]
    pub metadata: Option<std::collections::HashMap<String, String>>,
    #[serde(rename="options", skip_serializing_if = "Option::is_none")]
    pub options: Option<Vec<String>>,
    #[serde(rename="prompt", skip_serializing_if = "Option::is_none")]
    pub prompt: Option<String>,
    #[serde(rename="request_id", skip_serializing_if = "Option::is_none")]
    pub request_id: Option<String>,
    #[serde(rename="additionalProperties", skip_serializing_if = "Option::is_none")]
    pub additional_properties: Option<std::collections::HashMap<String, serde_json::Value>>,
}
