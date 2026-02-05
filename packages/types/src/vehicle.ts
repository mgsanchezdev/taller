export interface Vehicle {
  id: string;
  workshop_id: string;
  plate: string;
  brand: string;
  model: string;
  year: number;
  vin: string;
  color: string;
  owner_name: string;
  owner_phone: string;
  owner_email: string;
  created_at: string;
  updated_at: string;
}

export interface CreateVehicleInput {
  plate: string;
  brand?: string;
  model?: string;
  year?: number;
  vin?: string;
  color?: string;
  owner_name?: string;
  owner_phone?: string;
  owner_email?: string;
}
