const BASE = import.meta.env.VITE_API_BASE_URL ?? 'http://localhost:8080'

async function request<T>(path: string, init?: RequestInit, jwt?: string | null): Promise<T> {
  const headers: Record<string, string> = {
    'Content-Type': 'application/json',
    ...(init?.headers as Record<string, string>),
  }
  if (jwt) headers['Authorization'] = `Bearer ${jwt}`

  const res = await fetch(`${BASE}${path}`, { ...init, headers })
  const body = await res.json()
  if (!res.ok || !body.success) {
    throw new Error(body.error ?? `HTTP ${res.status}`)
  }
  return body.data as T
}

export const api = {
  get:  <T>(path: string, jwt?: string | null) => request<T>(path, { method: 'GET' }, jwt),
  post: <T>(path: string, payload: unknown, jwt?: string | null) =>
    request<T>(path, { method: 'POST', body: JSON.stringify(payload) }, jwt),
}
