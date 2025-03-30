const TOKEN_NAME = "userToken";
function _getToken(): string | null {
  return localStorage.getItem(TOKEN_NAME)
}

function _setToken(token: string): void {
  localStorage.setItem(TOKEN_NAME, token)
}

export async function getToken(): Promise<string> {
  const token = _getToken()
  if (token) return token
  const data = await apiUnauthFetch<{token: string}>("/token")
  if (data.token) {
    _setToken(data.token)
    return data.token
  }
  throw new Error("No token found")
}

export function apiUrl(path: string): string {
  const baseUrl = import.meta.env.VITE_API_URL || "http://localhost:3000"
  return `${baseUrl}${path}`
}

export async function apiUnauthFetch<R>(
  path: string,
  options?: RequestInit
): Promise<R> {
  const response = await fetch(apiUrl(path), {
    ...options,
    headers: {
      "Content-Type": "application/json",
      "Accept": "application/json",
      ...options?.headers,
    },
  })
  return response.json()
}


export async function apiFetch<R>(
  path: string,
  options?: RequestInit
): Promise<R> {
  return apiUnauthFetch<R>(path, {
    ...options,
    headers: {
      "Authorization": `Token ${await getToken()}`,
      ...options?.headers,
    },
  })
}