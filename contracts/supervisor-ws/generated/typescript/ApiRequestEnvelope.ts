import {ApiScope} from './ApiScope';
import {ApiWatchParams} from './ApiWatchParams';
interface ApiRequestEnvelope {
  /**
   * Dotted action name (e.g. 'beads.list')
   */
  action?: string;
  /**
   * Client-assigned correlation ID
   */
  id?: string;
  /**
   * Deduplication key for mutation replay
   */
  idempotencyKey?: string;
  /**
   * Action-specific request payload
   */
  payload?: any;
  scope?: ApiScope;
  /**
   * Must be 'request'
   */
  reservedType?: string;
  watch?: ApiWatchParams;
  additionalProperties?: Map<string, any>;
}
export { ApiRequestEnvelope };