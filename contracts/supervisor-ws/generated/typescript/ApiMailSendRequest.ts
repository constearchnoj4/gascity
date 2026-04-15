
interface ApiMailSendRequest {
  body?: string;
  reservedFrom?: string;
  rig?: string;
  subject?: string;
  to?: string;
  additionalProperties?: Map<string, any>;
}
export { ApiMailSendRequest };