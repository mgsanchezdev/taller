import { get, post } from './http';
import type { Inspection, WorkshopEntry } from '@workshop-platform/types';

export function listInspectionsByEntry(entryId: string) {
  return get<Inspection[]>(`v1/entries/${entryId}/inspections`);
}

export function getInspection(id: string) {
  return get<Inspection>(`v1/inspections/${id}`);
}

export function createInspection(entryId: string, body: Record<string, unknown>) {
  return post<Inspection>(`v1/entries/${entryId}/inspections`, body);
}

export function listEntries(params?: { limit?: number; offset?: number }) {
  const q = new URLSearchParams();
  if (params?.limit) q.set('limit', String(params.limit));
  if (params?.offset) q.set('offset', String(params.offset));
  const query = q.toString() ? `?${q.toString()}` : '';
  return get<{ items: WorkshopEntry[]; total: number }>(`v1/entries${query}`);
}
