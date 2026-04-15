
interface ApiSessionCreateRequest {
  alias?: string;
  async?: boolean;
  kind?: string;
  message?: string;
  reservedName?: string;
  options?: Map<string, string>;
  projectId?: string;
  sessionName?: string | null;
  title?: string;
  additionalProperties?: Map<string, any>;
}
export { ApiSessionCreateRequest };