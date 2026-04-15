import {ExtmsgConversationRef} from './ExtmsgConversationRef';
interface ApiExtmsgTranscriptAckRequest {
  conversation?: ExtmsgConversationRef;
  sequence?: number;
  sessionId?: string;
  additionalProperties?: Map<string, any>;
}
export { ApiExtmsgTranscriptAckRequest };