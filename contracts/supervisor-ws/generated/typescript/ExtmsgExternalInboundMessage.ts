import {ExtmsgExternalActor} from './ExtmsgExternalActor';
import {ExtmsgExternalAttachment} from './ExtmsgExternalAttachment';
import {ExtmsgConversationRef} from './ExtmsgConversationRef';
interface ExtmsgExternalInboundMessage {
  actor?: ExtmsgExternalActor;
  attachments?: ExtmsgExternalAttachment[];
  conversation?: ExtmsgConversationRef;
  dedupKey?: string;
  explicitTarget?: string;
  providerMessageId?: string;
  receivedAt?: string;
  replyToMessageId?: string;
  reservedText?: string;
  additionalProperties?: Map<string, any>;
}
export { ExtmsgExternalInboundMessage };