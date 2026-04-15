
interface ApiSocketSessionRespondPayload {
  action?: string;
  id?: string;
  metadata?: Map<string, string>;
  requestId?: string;
  reservedText?: string;
  additionalProperties?: Map<string, any>;
}
export { ApiSocketSessionRespondPayload };