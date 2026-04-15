import {ExtmsgConversationRef} from './ExtmsgConversationRef';
interface ApiExtmsgBindRequest {
  conversation?: ExtmsgConversationRef;
  metadata?: Map<string, string>;
  sessionId?: string;
  additionalProperties?: Map<string, any>;
}
export { ApiExtmsgBindRequest };