
interface ApiOrderResponse {
  captureOutput?: boolean;
  check?: string;
  description?: string;
  enabled?: boolean;
  exec?: string;
  formula?: string;
  gate?: string;
  interval?: string;
  reservedName?: string;
  on?: string;
  pool?: string;
  rig?: string;
  schedule?: string;
  scopedName?: string;
  timeout?: string;
  timeoutMs?: number;
  reservedType?: string;
  additionalProperties?: Map<string, any>;
}
export { ApiOrderResponse };