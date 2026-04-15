
interface ApiEventEnvelope {
  /**
   * Resume cursor for reconnection
   */
  cursor?: string;
  /**
   * Event type (e.g. 'bead.created')
   */
  eventType?: string;
  /**
   * Event sequence number
   */
  index?: number;
  /**
   * Event-specific payload
   */
  payload?: any;
  /**
   * Subscription that produced this event
   */
  subscriptionId?: string;
  /**
   * Must be 'event'
   */
  reservedType?: string;
  additionalProperties?: Map<string, any>;
}
export { ApiEventEnvelope };