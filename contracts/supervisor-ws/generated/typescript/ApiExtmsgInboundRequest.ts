import {ExtmsgExternalInboundMessage} from './ExtmsgExternalInboundMessage';
interface ApiExtmsgInboundRequest {
  accountId?: string;
  message?: ExtmsgExternalInboundMessage;
  payload?: string;
  provider?: string;
  additionalProperties?: Map<string, any>;
}
export { ApiExtmsgInboundRequest };