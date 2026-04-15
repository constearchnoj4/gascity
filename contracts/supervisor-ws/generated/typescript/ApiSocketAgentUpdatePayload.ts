
interface ApiSocketAgentUpdatePayload {
  reservedName?: string;
  provider?: string;
  scope?: string;
  suspended?: boolean | null;
  additionalProperties?: Map<string, any>;
}
export { ApiSocketAgentUpdatePayload };