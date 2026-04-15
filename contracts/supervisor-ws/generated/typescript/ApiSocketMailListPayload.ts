
interface ApiSocketMailListPayload {
  agent?: string;
  cursor?: string;
  limit?: number | null;
  rig?: string;
  reservedStatus?: string;
  additionalProperties?: Map<string, any>;
}
export { ApiSocketMailListPayload };