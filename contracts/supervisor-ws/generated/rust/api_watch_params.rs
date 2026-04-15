// ApiWatchParams represents a ApiWatchParams model.
#[derive(Clone, Debug, Deserialize, Serialize)]
pub struct ApiWatchParams {
    #[serde(rename="index", skip_serializing_if = "Option::is_none")]
    pub index: Option<i32>,
    #[serde(rename="wait", skip_serializing_if = "Option::is_none")]
    pub wait: Option<String>,
    #[serde(rename="additionalProperties", skip_serializing_if = "Option::is_none")]
    pub additional_properties: Option<std::collections::HashMap<String, serde_json::Value>>,
}
