export interface Repair {
  id: string;
  entry_id: string;
  vehicle_id: string;
  fault_id?: string;
  description: string;
  labor_hours: number;
  parts_cost: number;
  status: 'pending' | 'in_progress' | 'completed';
  completed_at?: string;
  technician_id?: string;
  created_at: string;
  updated_at: string;
}
