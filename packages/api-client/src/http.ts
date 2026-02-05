const defaultBaseURL = typeof process !== 'undefined' && process.env?.NEXT_PUBLIC_API_URL
  ? process.env.NEXT_PUBLIC_API_URL
  : 'http://localhost:8080/api';

export type ApiConfig = {
  baseURL?: string;
  getToken?: () => string | null;
};

let config: ApiConfig = {
  baseURL: defaultBaseURL,
  getToken: () => typeof window !== 'undefined' ? localStorage.getItem('token') : null,
};

export function configureApi(c: Partial<ApiConfig>) {
  config = { ...config, ...c };
}

export async function request<T>(
  path: string,
  options: RequestInit = {}
): Promise<T> {
  const url = `${config.baseURL?.replace(/\/$/, '')}/${path.replace(/^\//, '')}`;
  const token = config.getToken?.() ?? null;
  const headers: HeadersInit = {
    'Content-Type': 'application/json',
    ...(options.headers as Record<string, string>),
  };
  if (token) {
    (headers as Record<string, string>)['Authorization'] = `Bearer ${token}`;
  }
  const res = await fetch(url, { ...options, headers });
  if (!res.ok) {
    const text = await res.text();
    let err: Error & { status?: number } = new Error(text || res.statusText);
    (err as { status?: number }).status = res.status;
    throw err;
  }
  const contentType = res.headers.get('content-type');
  if (contentType?.includes('application/json')) {
    return res.json() as Promise<T>;
  }
  return res.text() as unknown as T;
}

export function get<T>(path: string) {
  return request<T>(path, { method: 'GET' });
}

export function post<T>(path: string, body?: unknown) {
  return request<T>(path, { method: 'POST', body: body ? JSON.stringify(body) : undefined });
}

export function put<T>(path: string, body?: unknown) {
  return request<T>(path, { method: 'PUT', body: body ? JSON.stringify(body) : undefined });
}

export function del<T>(path: string) {
  return request<T>(path, { method: 'DELETE' });
}
