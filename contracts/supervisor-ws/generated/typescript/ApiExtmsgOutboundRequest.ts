import {ExtmsgConversationRef} from './ExtmsgConversationRef';
interface ApiExtmsgOutboundRequest {
  conversation?: ExtmsgConversationRef;
  idempotencyKey?: string;
  replyToMessageId?: string;
  sessionId?: string;
  reservedText?: string;
  additionalProperties?: Map<string, any>;
}
export { ApiExtmsgOutboundRequest };