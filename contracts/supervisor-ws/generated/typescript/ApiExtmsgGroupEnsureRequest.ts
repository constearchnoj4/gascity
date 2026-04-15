import {ExtmsgConversationRef} from './ExtmsgConversationRef';
interface ApiExtmsgGroupEnsureRequest {
  defaultHandle?: string;
  metadata?: Map<string, string>;
  mode?: string;
  rootConversation?: ExtmsgConversationRef;
  additionalProperties?: Map<string, any>;
}
export { ApiExtmsgGroupEnsureRequest };