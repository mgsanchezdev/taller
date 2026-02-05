import { get } from './http';
import type { Report } from '@workshop-platform/types';

export function listReportsByEntry(entryId: string) {
  return get<Report[]>(`v1/entries/${entryId}/reports`);
}

export function generateReport(entryId: string, type: 'inspection' | 'repair' | 'full') {
  return get<Report>(`v1/entries/${entryId}/reports/generate?type=${type}`);
}
