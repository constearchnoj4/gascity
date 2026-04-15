
interface ApiSocketSessionsListPayload {
  cursor?: string;
  limit?: number | null;
  peek?: boolean;
  state?: string;
  template?: string;
  additionalProperties?: Map<string, any>;
}
export { ApiSocketSessionsListPayload };