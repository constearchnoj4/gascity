import {ExtmsgConversationRef} from './ExtmsgConversationRef';
interface ApiExtmsgUnbindRequest {
  conversation?: ExtmsgConversationRef;
  sessionId?: string;
  additionalProperties?: Map<string, any>;
}
export { ApiExtmsgUnbindRequest };