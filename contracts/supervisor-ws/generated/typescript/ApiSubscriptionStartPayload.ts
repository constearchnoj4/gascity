
interface ApiSubscriptionStartPayload {
  /**
   * Resume from this cursor
   */
  afterCursor?: string;
  /**
   * Resume from this event sequence
   */
  afterSeq?: number;
  /**
   * Stream format: 'text', 'raw', 'jsonl'
   */
  format?: string;
  /**
   * Subscription type: 'events' or 'session.stream'
   */
  kind?: string;
  /**
   * Session ID or name (for session.stream)
   */
  target?: string;
  /**
   * Most recent N turns (0=all)
   */
  turns?: number;
  additionalProperties?: Map<string, any>;
}
export { ApiSubscriptionStartPayload };