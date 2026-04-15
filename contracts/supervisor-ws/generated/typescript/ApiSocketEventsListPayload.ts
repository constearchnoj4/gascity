
interface ApiSocketEventsListPayload {
  actor?: string;
  cursor?: string;
  limit?: number | null;
  since?: string;
  reservedType?: string;
  additionalProperties?: Map<string, any>;
}
export { ApiSocketEventsListPayload };