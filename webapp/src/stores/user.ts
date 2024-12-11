import axios from 'axios'
import { ref, computed } from 'vue'
import { defineStore } from 'pinia'
import { api } from './api'
import router from '../router'

interface Issue {
  id: number
  number: number
  title: string
  body: string
  state: string
  created_at: string
  updated_at: string
  closed_at: string
  labels: {
    id: number
    name: string
    color: string
  }[]
  assignees: {
    login: string
  }[]
}
interface Repo {
  id: number
  name: string
  full_name: string
  owner: {
    login: string
  }
  html_url: string
  description: string
  stargazers_count: number
  watchers_count: number
  forks_count: number
  open_issues_count: number
  license: {
    name: string
  }
  created_at: string
  updated_at: string
  pushed_at: string
  issues: Issue[]
}

interface HostData {
  connected: boolean
  token: string | null
  userInfo: any | null
  repos: Repo[]
}
interface User {
  userid: string
  username: string
  avatar: string
  userBio?: string
  fullName?: string
  email?: string
  phoneNumber?: string
  location?: string
  bearerToken?: string
  refreshJwt?: string
  aboutMe?: string
  interests?: string[]
  recentPosts?: any[]
  level?: number
  achievements?: any[]
  activities?: any[]
  bounties?: any[]
  transactions?: any[]
  paymentInfo?: any
  isLoggedIn?: boolean
  authGithubHeader?: string
  githubUser?: GithubUser
  repos?: Repo[]
  hosts?: { [host: string]: HostData }
  issues?: Issue[]
  orgs?: any[]
}

interface GithubUser {
  id: number
  login: string
  avatar_url: string
}

interface LoginCredentials {
  username: string
  password: string
}

export const useUserStore = defineStore('user', () => {
  const user = ref<User | null>(null)
  const loggedIn = ref<boolean>(false)
  const token = ref<string>('')
  const githubUser = ref<GithubUser | null>(null)
  const hosts = ref<{ [host: string]: HostData }>({})
  const issues = ref<Issue[]>([])
  const repos = ref<Repo[]>([])
  const orgs = ref<any[]>([])

  const isLoggedIn = computed<boolean>(() => {
    return loggedIn.value || !!localStorage.getItem('token')
  })

  const authHeader = computed<string>(() => {
    const storedToken = localStorage.getItem('token')
    return `Bearer ${storedToken || ''}`
  })

  const authGithubHeader = computed<string>(() => {
    const storedGithubToken = localStorage.getItem('githubToken')
    return storedGithubToken ? `Bearer ${storedGithubToken}` : ''
  })

  async function syncGithubData(token: string): Promise<boolean> {
    try {
      // Fetch the authenticated user's details
      const userResponse = await axios.get('https://api.github.com/user', {
        headers: {
          Authorization: `Bearer ${token}`,
          Accept: 'application/vnd.github.v3+json'
        }
      })
      const user = userResponse.data

      console.log('GitHub User Response:', user)

      // Update profile
      updateProfilFromGithub(user)

      // Fetch organizations of the user
      const orgsResponse = await axios.get('https://api.github.com/user/orgs', {
        headers: {
          Authorization: `Bearer ${token}`,
          Accept: 'application/vnd.github.v3+json'
        }
      })
      const organizations = orgsResponse.data
      orgs.value.push(...orgsResponse.data)

      // Create a function to fetch repositories for a given login (user or org)
      const fetchRepositories = async (login: string, org: boolean): Promise<any[]> => {
        if (org) {
          const reposResponse = await axios.get(`https://api.github.com/orgs/${login}/repos`, {
            headers: {
              Authorization: `Bearer ${token}`,
              Accept: 'application/vnd.github.v3+json'
            }
          })
          return reposResponse.data
        } else {
          const reposResponse = await axios.get(`https://api.github.com/users/${login}/repos`, {
            headers: {
              Authorization: `Bearer ${token}`,
              Accept: 'application/vnd.github.v3+json'
            }
          })
          return reposResponse.data
        }
      }

      // Create a function to fetch issues for a given repository
      const fetchIssues = async (owner: string, repo: string): Promise<any[]> => {
        const issuesResponse = await axios.get(
          `https://api.github.com/repos/${owner}/${repo}/issues`,
          {
            headers: {
              Authorization: `Bearer ${token}`,
              Accept: 'application/vnd.github.v3+json'
            }
          }
        )
        return issuesResponse.data
      }

      // Fetch repositories and issues for user and organizations
      const allLogins = [user.login, ...organizations.map((org: any) => org.login)]
      for (const login of allLogins) {
        console.log(`Fetching repositories for: ${login}`)
        const repositories = await fetchRepositories(login, login !== user.login)
        repos.value.push(...repositories)
        for (const repo of repositories) {
          console.log(`Fetching issues for repo: ${repo.name} of ${login}`)
          const fetchedIssues = await fetchIssues(login, repo.name)
          for (const issue of fetchedIssues) {
            issue.repo_avatar = repo.owner.avatar_url
          }
          issues.value.push(...fetchedIssues)
          console.log(`Issues for repo ${repo.name}:`, issues)
          //append issues to userStore
        }
      }

      return true
    } catch (error: any) {
      // Comprehensive error logging
      if (axios.isAxiosError(error)) {
        console.error('GitHub API Error:', {
          status: error.response?.status,
          data: error.response?.data,
          headers: error.response?.headers
        })
      } else {
        console.error('Unexpected error:', error)
      }

      return false
    }
  }

  async function refreshJwt(): Promise<boolean> {
    const storedToken = localStorage.getItem('token')
    if (!storedToken) {
      logout()
      return false
    }

    try {
      const response = await axios.post(
        'https://api.example.com/auth/refresh',
        {},
        {
          headers: { Authorization: `Bearer ${storedToken}` }
        }
      )

      const newToken = response.data.token
      if (!newToken) {
        logout()
        return false
      }

      localStorage.setItem('token', newToken)
      token.value = newToken
      loggedIn.value = true
      return true
    } catch (error) {
      console.error('Token refresh failed:', error)
      logout()
      return false
    }
  }

  async function login(creds: LoginCredentials): Promise<void> {
    try {
      const response = await api.post('/v1/signin', creds)
      const { accessToken, refreshToken } = response.data
      console.log(response.data)

      if (accessToken && refreshToken) {
        localStorage.setItem('token', accessToken)
        localStorage.setItem('refreshToken', refreshToken)
        token.value = accessToken
        loggedIn.value = true

        //push to dashboard
        router.push({ name: 'Profile' })
      } else {
        throw new Error('Missing tokens in response')
      }
    } catch (error) {
      console.error('Login failed:', error)
      throw new Error('Invalid credentials or login error')
    }
  }

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

  async function logout(): Promise<void> {
    try {
      // Clear tokens from localStorage
      token.value = ''
      localStorage.removeItem('token')

      localStorage.removeItem('refreshToken')

      // Reset the token state and loggedIn status
      loggedIn.value = false

      // Redirect to the home page
      router.push({ path: '/' })
    } catch (error) {
      console.error('Logout failed:', error)
      throw new Error('Logout error')
    }
  }

  async function getUserInfo(): Promise<void> {
    // Mock implementation; replace with actual API call.
    try {
      const response = await axios.get('/auth/user', {
        headers: { Authorization: authHeader.value }
      })
      user.value = response.data
      localStorage.setItem('user', JSON.stringify(response.data))
    } catch (error) {
      console.error('Failed to fetch user info:', error)
    }
  }

  async function updateProfilFromGithub(githubData: any): Promise<void> {
    localStorage.setItem('github_user', JSON.stringify(githubData))
  }

  return {
    user,
    loggedIn,
    isLoggedIn,
    authHeader,
    token,
    authGithubHeader,
    githubUser,
    hosts,
    issues,
    repos,
    orgs,
    syncGithubData,
    refreshJwt,
    login,
    logout,
    getUserInfo
  }
})
