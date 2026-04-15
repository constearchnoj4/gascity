
interface ApiBeadCreateRequest {
  assignee?: string;
  description?: string;
  labels?: string[] | null;
  priority?: number | null;
  rig?: string;
  title?: string;
  reservedType?: string;
  additionalProperties?: Map<string, any>;
}
export { ApiBeadCreateRequest };