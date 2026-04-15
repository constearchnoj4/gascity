import {RuntimePendingInteraction} from './RuntimePendingInteraction';
interface ApiSessionPendingResponse {
  pending?: RuntimePendingInteraction;
  supported?: boolean;
  additionalProperties?: Map<string, any>;
}
export { ApiSessionPendingResponse };