import { post } from './http';

export interface LoginRequest {
  email: string;
  password: string;
  workshop_id: string;
}

export interface User {
  id: string;
  workshop_id: string;
  email: string;
  name: string;
  role: string;
  active: boolean;
  created_at: string;
  updated_at: string;
}

export interface LoginResponse {
  token: string;
  user: User;
}

export function login(body: LoginRequest) {
  return post<LoginResponse>('auth/login', body);
}
