
interface ApiEventEmitRequest {
  actor?: string;
  message?: string;
  subject?: string;
  reservedType?: string;
  additionalProperties?: Map<string, any>;
}
export { ApiEventEmitRequest };