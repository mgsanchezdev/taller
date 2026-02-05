import { get, post } from './http';
import type { Repair } from '@workshop-platform/types';

export function listRepairsByEntry(entryId: string) {
  return get<Repair[]>(`v1/entries/${entryId}/repairs`);
}

export function getRepair(id: string) {
  return get<Repair>(`v1/repairs/${id}`);
}

export function createRepair(entryId: string, body: Record<string, unknown>) {
  return post<Repair>(`v1/entries/${entryId}/repairs`, body);
}
