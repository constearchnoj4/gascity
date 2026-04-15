// BeadsBead represents a BeadsBead model.
#[derive(Clone, Debug, Deserialize, Serialize)]
pub struct BeadsBead {
    #[serde(rename="assignee", skip_serializing_if = "Option::is_none")]
    pub assignee: Option<String>,
    #[serde(rename="created_at", skip_serializing_if = "Option::is_none")]
    pub created_at: Option<String>,
    #[serde(rename="dependencies", skip_serializing_if = "Option::is_none")]
    pub dependencies: Option<Vec<crate::BeadsDep>>,
    #[serde(rename="description", skip_serializing_if = "Option::is_none")]
    pub description: Option<String>,
    #[serde(rename="from", skip_serializing_if = "Option::is_none")]
    pub from: Option<String>,
    #[serde(rename="id", skip_serializing_if = "Option::is_none")]
    pub id: Option<String>,
    #[serde(rename="issue_type", skip_serializing_if = "Option::is_none")]
    pub issue_type: Option<String>,
    #[serde(rename="labels", skip_serializing_if = "Option::is_none")]
    pub labels: Option<Vec<String>>,
    #[serde(rename="metadata", skip_serializing_if = "Option::is_none")]
    pub metadata: Option<std::collections::HashMap<String, String>>,
    #[serde(rename="needs", skip_serializing_if = "Option::is_none")]
    pub needs: Option<Vec<String>>,
    #[serde(rename="parent", skip_serializing_if = "Option::is_none")]
    pub parent: Option<String>,
    #[serde(rename="priority", skip_serializing_if = "Option::is_none")]
    pub priority: Option<i32>,
    #[serde(rename="ref", skip_serializing_if = "Option::is_none")]
    pub reserved_ref: Option<String>,
    #[serde(rename="status", skip_serializing_if = "Option::is_none")]
    pub status: Option<String>,
    #[serde(rename="title", skip_serializing_if = "Option::is_none")]
    pub title: Option<String>,
    #[serde(rename="additionalProperties", skip_serializing_if = "Option::is_none")]
    pub additional_properties: Option<std::collections::HashMap<String, serde_json::Value>>,
}
