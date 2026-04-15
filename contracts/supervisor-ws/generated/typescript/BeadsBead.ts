import {BeadsDep} from './BeadsDep';
interface BeadsBead {
  assignee?: string;
  createdAt?: string;
  dependencies?: BeadsDep[];
  description?: string;
  reservedFrom?: string;
  id?: string;
  issueType?: string;
  labels?: string[];
  metadata?: Map<string, string>;
  needs?: string[];
  reservedParent?: string;
  priority?: number | null;
  ref?: string;
  reservedStatus?: string;
  title?: string;
  additionalProperties?: Map<string, any>;
}
export { BeadsBead };