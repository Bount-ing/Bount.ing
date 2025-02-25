import { ref, computed } from 'vue'
import { defineStore } from 'pinia'
import { api } from './api'
import router from '../router'
import type { LoginCredentials, TokenResponse } from '../types/auth'
import { useUserStore } from './user'
import { useNotificationStore } from './notification'

export const useAuthStore = defineStore('auth', () => {
  const token = ref<string>('')
  const loggedIn = ref<boolean>(false)
  const refreshToken = ref<string>('')

  const isLoggedIn = computed<boolean>(() => {
    return loggedIn.value || !!localStorage.getItem('token')
  })

  const authHeader = computed<string>(() => {
    const storedToken = localStorage.getItem('token') || token.value
    return `Bearer ${storedToken}`
  })

  function parseJwt(token: string): Record<string, any> | null {
    try {
      const base64Url = token.split('.')[1]
      const base64 = base64Url.replace(/-/g, '+').replace(/_/g, '/')
      const jsonPayload = atob(base64)
        .split('')
        .map((c) => `%${('00' + c.charCodeAt(0).toString(16)).slice(-2)}`)
        .join('')
      return JSON.parse(decodeURIComponent(jsonPayload))
    } catch (error) {
      console.error('Failed to parse JWT:', error)
      return null
    }
  }

  async function login(creds: LoginCredentials): Promise<void> {
    const notificationStore = useNotificationStore()
    const userStore = useUserStore()

    try {
      const response = await api.post<TokenResponse>('/v1/signin', creds)
      const { accessToken, refreshToken: newRefreshToken } = response.data

      if (!accessToken || !newRefreshToken) {
        throw new Error('Missing tokens in response')
      }

      localStorage.setItem('token', accessToken)
      localStorage.setItem('refreshToken', newRefreshToken)
      token.value = accessToken
      refreshToken.value = newRefreshToken
      loggedIn.value = true

      const decodedJwt = parseJwt(accessToken)
      if (!decodedJwt?.UID) {
        throw new Error('Invalid token: missing UID')
      }
      localStorage.setItem('userId', decodedJwt.UID.toString())

      await userStore.getUserInfo()
      router.push('/profile')
    } catch (error: any) {
      console.error('Login failed:', error)
      notificationStore.showNotification(
        error.response?.data?.reason || 'Login failed. Please try again.',
        'error'
      )
      throw error
    }
  }

  async function refreshJwt(): Promise<boolean> {
    const currentRefreshToken = localStorage.getItem('refreshToken') || refreshToken.value
    console.log('Refreshing token...')
    if (!currentRefreshToken) {
      await logout()
      return false
    }

    console.log('Current refresh token:', currentRefreshToken)
    try {
      const response = await api.post<TokenResponse>(
        '/v1/refresh',
        {},
        {
          headers: {
            Authorization: `Bearer ${currentRefreshToken}`
          },
          withCredentials: true // Important for cookie handling
        }
      )

      const { accessToken, refreshToken: newRefreshToken } = response.data
      if (!accessToken || !newRefreshToken) {
        await logout()
        return false
      }

      // Update tokens
      localStorage.setItem('token', accessToken)
      localStorage.setItem('refreshToken', newRefreshToken)
      token.value = accessToken
      refreshToken.value = newRefreshToken
      loggedIn.value = true

      // Verify the new token is valid
      const decodedJwt = parseJwt(accessToken)
      if (!decodedJwt?.UID) {
        await logout()
        return false
      }

      // Update user information with new token
      const userStore = useUserStore()
      await userStore.getUserInfo()
      return true
    } catch (error) {
      console.error('Token refresh failed:', error)
      router.push('/signin')
      return false
    }
  }

  async function logout(): Promise<void> {
    console.log('Logging out...')
    token.value = ''
    refreshToken.value = ''
    loggedIn.value = false

    localStorage.removeItem('token')
    localStorage.removeItem('refreshToken')

    const userStore = useUserStore()
    userStore.clearUserData()

    // Reset login status
    router.push({ path: '/signin' })
  }

  return {
    token,
    refreshToken,
    loggedIn,
    isLoggedIn,
    authHeader,
    login,
    logout,
    refreshJwt,
    parseJwt
  }
})
