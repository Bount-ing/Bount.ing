export interface LoginCredentials {
  username: string
  password: string
}

export interface TokenResponse {
  accessToken: string
  refreshToken: string
}
