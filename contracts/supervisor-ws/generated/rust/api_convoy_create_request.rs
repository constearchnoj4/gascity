// ApiConvoyCreateRequest represents a ApiConvoyCreateRequest model.
#[derive(Clone, Debug, Deserialize, Serialize)]
pub struct ApiConvoyCreateRequest {
    #[serde(rename="items", skip_serializing_if = "Option::is_none")]
    pub items: Option<Vec<String>>,
    #[serde(rename="rig", skip_serializing_if = "Option::is_none")]
    pub rig: Option<String>,
    #[serde(rename="title", skip_serializing_if = "Option::is_none")]
    pub title: Option<String>,
    #[serde(rename="additionalProperties", skip_serializing_if = "Option::is_none")]
    pub additional_properties: Option<std::collections::HashMap<String, serde_json::Value>>,
}
