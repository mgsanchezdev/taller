export interface Inspection {
  id: string;
  entry_id: string;
  vehicle_id: string;
  template_id: string;
  performed_by: string;
  performed_at: string;
  status: 'draft' | 'completed' | 'signed';
  notes: string;
  created_at: string;
  updated_at: string;
}

export interface InspectionItemResult {
  id: string;
  inspection_id: string;
  item_id: string;
  status: 'ok' | 'damaged' | 'missing' | 'n/a';
  notes?: string;
  photo_url?: string;
}

export interface WorkshopEntry {
  id: string;
  workshop_id: string;
  vehicle_id: string;
  entry_number: string;
  entry_date: string;
  mileage: number;
  fuel_level: string;
  notes: string;
  status: string;
  created_at: string;
  updated_at: string;
}
