// ApiSlingBody represents a ApiSlingBody model.
#[derive(Clone, Debug, Deserialize, Serialize)]
pub struct ApiSlingBody {
    #[serde(rename="attached_bead_id", skip_serializing_if = "Option::is_none")]
    pub attached_bead_id: Option<String>,
    #[serde(rename="bead", skip_serializing_if = "Option::is_none")]
    pub bead: Option<String>,
    #[serde(rename="formula", skip_serializing_if = "Option::is_none")]
    pub formula: Option<String>,
    #[serde(rename="rig", skip_serializing_if = "Option::is_none")]
    pub rig: Option<String>,
    #[serde(rename="scope_kind", skip_serializing_if = "Option::is_none")]
    pub scope_kind: Option<String>,
    #[serde(rename="scope_ref", skip_serializing_if = "Option::is_none")]
    pub scope_ref: Option<String>,
    #[serde(rename="target", skip_serializing_if = "Option::is_none")]
    pub target: Option<String>,
    #[serde(rename="title", skip_serializing_if = "Option::is_none")]
    pub title: Option<String>,
    #[serde(rename="vars", skip_serializing_if = "Option::is_none")]
    pub vars: Option<std::collections::HashMap<String, String>>,
    #[serde(rename="additionalProperties", skip_serializing_if = "Option::is_none")]
    pub additional_properties: Option<std::collections::HashMap<String, serde_json::Value>>,
}
