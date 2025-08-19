import { type Place as PlaceType } from '../types/placeTypes';

export async function postJSON<T>(url: string, body: unknown, init?: RequestInit): Promise<T> {
  const res = await fetch(url, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
      ...(init?.headers || {}),
    },
    body: JSON.stringify(body),
    ...init,
  });
  if (!res.ok) {
    const text = await res.text().catch(() => '');
    throw new Error(text || `Request failed with status ${res.status}`);
  }
  return res.json() as Promise<T>;
}

export async function putJSON<T>(url: string, body: unknown, init?: RequestInit): Promise<T> {
  const res = await fetch(url, {
    method: 'PUT',
    headers: {
      'Content-Type': 'application/json',
      ...(init?.headers || {}),
    },
    body: JSON.stringify(body),
    ...init,
  });
  if (!res.ok) {
    const text = await res.text().catch(() => '');
    throw new Error(text || `Request failed with status ${res.status}`);
  }
  return res.json() as Promise<T>;
}

export async function getJSON<T>(url: string, init?: RequestInit): Promise<T> {
  const res = await fetch(url, {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json',
      ...(init?.headers || {}),
    },
    ...init,
  });
  if (!res.ok) {
    const text = await res.text().catch(() => '');
    throw new Error(text || `Request failed with status ${res.status}`);
  }
  return res.json() as Promise<T>;
}

export function createPlace(payload: { name: string; tags: string[] }): Promise<PlaceType> {
  // Backend base path aligns with existing GET '/api/v1/places/'
  return postJSON<PlaceType>('/api/v1/places/', payload);
}

export function updatePlace(id: number, payload: { name: string; tags: string[] }): Promise<PlaceType> {
  return putJSON<PlaceType>(`/api/v1/places/${id}`, payload);
}

export function getPlaces(): Promise<PlaceType[]> {
  return getJSON<PlaceType[]>('/api/v1/places/');
}