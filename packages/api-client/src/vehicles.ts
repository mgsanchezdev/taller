import { get, post, put, del } from './http';
import type { Vehicle, CreateVehicleInput } from '@workshop-platform/types';

const base = 'v1/vehicles';

export function listVehicles() {
  return get<Vehicle[]>(base + '/');
}

export function getVehicle(id: string) {
  return get<Vehicle>(`${base}/${id}`);
}

export function createVehicle(data: CreateVehicleInput) {
  return post<Vehicle>(base + '/', data);
}

export function updateVehicle(id: string, data: Partial<CreateVehicleInput>) {
  return put<Vehicle>(`${base}/${id}`, data);
}

export function deleteVehicle(id: string) {
  return del<void>(`${base}/${id}`);
}
