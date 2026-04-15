import {ApiFieldError} from './ApiFieldError';
interface ApiErrorEnvelope {
  /**
   * Machine-readable error code
   */
  code?: string;
  /**
   * Per-field validation errors
   */
  details?: ApiFieldError[];
  /**
   * Correlation ID (empty for connection-level errors)
   */
  id?: string;
  /**
   * Human-readable error message
   */
  message?: string;
  /**
   * Must be 'error'
   */
  reservedType?: string;
  additionalProperties?: Map<string, any>;
}
export { ApiErrorEnvelope };