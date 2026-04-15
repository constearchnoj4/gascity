// ApiCityPatchRequest represents a ApiCityPatchRequest model.
#[derive(Clone, Debug, Deserialize, Serialize)]
pub struct ApiCityPatchRequest {
    #[serde(rename="suspended", skip_serializing_if = "Option::is_none")]
    pub suspended: Option<bool>,
    #[serde(rename="additionalProperties", skip_serializing_if = "Option::is_none")]
    pub additional_properties: Option<std::collections::HashMap<String, serde_json::Value>>,
}
