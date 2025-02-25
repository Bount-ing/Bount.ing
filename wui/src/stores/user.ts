import { ref, computed } from 'vue'
import { defineStore } from 'pinia'
import axios from 'axios'
import { api } from './api'
import { useAuthStore } from './auth'
import type { User, EditableUserFields, DEFAULT_USER_VALUES } from '../types/user'
import type { GithubUser, HostData, Issue, Repo } from '../types/github'

export const useUserStore = defineStore('user', () => {
  const user = ref<User | null>(null)
  const githubUser = ref<GithubUser | null>(null)
  const hosts = ref<{ [host: string]: HostData }>({})
  const issues = ref<Issue[]>([])
  const importedIssues = ref<Issue[]>([])
  const repos = ref<Repo[]>([])
  const orgs = ref<any[]>([])

  const authGithubHeader = computed<string>(() => {
    const storedGithubToken = localStorage.getItem('githubToken')
    return storedGithubToken ? `Bearer ${storedGithubToken}` : ''
  })

  async function getUserInfo(): Promise<void> {
    const authStore = useAuthStore()
    try {
      const response = await api.get<User>('/v1/users/me', {
        headers: { Authorization: authStore.authHeader }
      })

      user.value = response.data
      localStorage.setItem('user', JSON.stringify(response.data))
      console.log('User information:', user.value)
    } catch (error: any) {
      console.error('Failed to fetch user info:', error)

      if (error.response && error.response.status === 401) {
        console.warn('Unauthorized while fetching user info. Attempting to refresh token...')
        const refreshSuccessful = await authStore.refreshJwt()
        if (!refreshSuccessful) {
          clearUserData()
          await authStore.logout()
        }
      }
    }
  }

  async function syncGithubData(token: string): Promise<boolean> {
    try {
      // Fetch the authenticated user's details
      const userResponse = await axios.get<GithubUser>('https://api.github.com/user', {
        headers: {
          Authorization: `Bearer ${token}`,
          Accept: 'application/vnd.github.v3+json'
        }
      })
      const githubUserData = userResponse.data

      console.log('GitHub User Response:', githubUserData)

      // Update profile
      updateProfilFromGithub(githubUserData)

      // Fetch organizations of the user
      const orgsResponse = await axios.get<any[]>('https://api.github.com/user/orgs', {
        headers: {
          Authorization: `Bearer ${token}`,
          Accept: 'application/vnd.github.v3+json'
        }
      })
      const organizations = orgsResponse.data
      orgs.value = orgsResponse.data

      // Create a function to fetch repositories for a given login (user or org)
      const fetchRepositories = async (login: string, org: boolean): Promise<Repo[]> => {
        const endpoint = org
          ? `https://api.github.com/orgs/${login}/repos`
          : `https://api.github.com/users/${login}/repos`

        const reposResponse = await axios.get<Repo[]>(endpoint, {
          headers: {
            Authorization: `Bearer ${token}`,
            Accept: 'application/vnd.github.v3+json'
          }
        })
        return reposResponse.data
      }

      // Create a function to fetch issues for a given repository
      const fetchIssues = async (sponsor: string, repo: string): Promise<Issue[]> => {
        const issuesResponse = await axios.get<Issue[]>(
          `https://api.github.com/repos/${sponsor}/${repo}/issues`,
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
      const allLogins = [githubUserData.login, ...organizations.map((org: any) => org.login)]
      const allRepos: Repo[] = []
      const allIssues: Issue[] = []

      for (const login of allLogins) {
        console.log(`Fetching repositories for: ${login}`)
        const repositories = await fetchRepositories(login, login !== githubUserData.login)
        allRepos.push(...repositories)

        for (const repo of repositories) {
          console.log(`Fetching issues for repo: ${repo.name} of ${login}`)
          const fetchedIssues = await fetchIssues(login, repo.name)

          for (const issue of fetchedIssues) {
            issue.repo_avatar = repo.sponsor.avatar_url
          }

          allIssues.push(...fetchedIssues)
          console.log(`Issues for repo ${repo.name}:`, fetchedIssues)
        }
      }

      repos.value = allRepos
      issues.value = allIssues

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

  function updateProfileInfo(userprofileInfo: EditableUserFields): void {
    api.put<User>('/v1/users/me/profile', userprofileInfo, {
      headers: {
        Authorization: authGithubHeader.value
      }
    }).then((response) => {
      if (response.status >= 200 && response.status < 300) {
        getUserInfo()
      }
    })
  }

  function updateProfilFromGithub(githubData: any): void {
    githubUser.value = githubData
    localStorage.setItem('github_user', JSON.stringify(githubData))
  }

  function clearUserData(): void {

    user.value = null
    githubUser.value = null
    hosts.value = {}
    issues.value = []
    importedIssues.value = []
    repos.value = []
    orgs.value = []
    localStorage.removeItem('user')
    localStorage.removeItem('github_user')
  }

  return {
    user,
    githubUser,
    hosts,
    issues,
    importedIssues,
    repos,
    orgs,
    authGithubHeader,
    getUserInfo,
    updateProfileInfo,
    syncGithubData,
    clearUserData
  }
})
