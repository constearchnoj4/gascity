
interface MailMessage {
  body?: string;
  cc?: string[];
  createdAt?: string;
  reservedFrom?: string;
  id?: string;
  priority?: number;
  read?: boolean;
  replyTo?: string;
  rig?: string;
  subject?: string;
  threadId?: string;
  to?: string;
  additionalProperties?: Map<string, any>;
}
export { MailMessage };