
interface RuntimePendingInteraction {
  kind?: string;
  metadata?: Map<string, string>;
  options?: string[];
  reservedPrompt?: string;
  requestId?: string;
  additionalProperties?: Map<string, any>;
}
export { RuntimePendingInteraction };