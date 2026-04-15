
interface ApiHelloEnvelope {
  /**
   * Sorted list of supported action names
   */
  capabilities?: string[] | null;
  /**
   * Protocol version (e.g. 'gc.v1alpha1')
   */
  protocol?: string;
  /**
   * True if mutations are disabled
   */
  readOnly?: boolean;
  /**
   * 'city' or 'supervisor'
   */
  serverRole?: string;
  /**
   * Supported subscription types
   */
  subscriptionKinds?: string[];
  /**
   * Must be 'hello'
   */
  reservedType?: string;
  additionalProperties?: Map<string, any>;
}
export { ApiHelloEnvelope };