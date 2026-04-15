
interface ApiResponseEnvelope {
  /**
   * Correlation ID matching the request
   */
  id?: string;
  /**
   * Server event index for watch semantics
   */
  index?: number;
  /**
   * Action-specific response payload
   */
  result?: any;
  /**
   * Must be 'response'
   */
  reservedType?: string;
  additionalProperties?: Map<string, any>;
}
export { ApiResponseEnvelope };