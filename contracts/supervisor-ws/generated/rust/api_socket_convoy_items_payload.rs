// ApiSocketConvoyItemsPayload represents a ApiSocketConvoyItemsPayload model.
#[derive(Clone, Debug, Deserialize, Serialize)]
pub struct ApiSocketConvoyItemsPayload {
    #[serde(rename="id", skip_serializing_if = "Option::is_none")]
    pub id: Option<String>,
    #[serde(rename="items", skip_serializing_if = "Option::is_none")]
    pub items: Option<Vec<String>>,
    #[serde(rename="additionalProperties", skip_serializing_if = "Option::is_none")]
    pub additional_properties: Option<std::collections::HashMap<String, serde_json::Value>>,
}
