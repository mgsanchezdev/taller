export interface Report {
  id: string;
  workshop_id: string;
  entry_id: string;
  vehicle_id: string;
  type: 'inspection' | 'repair' | 'full';
  pdf_url: string;
  generated_at: string;
  created_at: string;
}
