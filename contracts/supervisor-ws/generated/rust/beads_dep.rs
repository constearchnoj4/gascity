// BeadsDep represents a BeadsDep model.
#[derive(Clone, Debug, Deserialize, Serialize)]
pub struct BeadsDep {
    #[serde(rename="depends_on_id", skip_serializing_if = "Option::is_none")]
    pub depends_on_id: Option<String>,
    #[serde(rename="issue_id", skip_serializing_if = "Option::is_none")]
    pub issue_id: Option<String>,
    #[serde(rename="type", skip_serializing_if = "Option::is_none")]
    pub reserved_type: Option<String>,
    #[serde(rename="additionalProperties", skip_serializing_if = "Option::is_none")]
    pub additional_properties: Option<std::collections::HashMap<String, serde_json::Value>>,
}
