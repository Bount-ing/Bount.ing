<template>
  <div>
    <h2 class="text-2xl font-semibold mb-4">Issues and Bounties</h2>

    <!-- Loading State -->
    <div v-if="loading" class="text-center text-xl">Loading issues and bounties...</div>

    <!-- Error State -->
    <div v-else-if="error" class="text-center text-red-500 text-lg">
      {{ error }}
    </div>

    <!-- Issues with Bounties -->
    <div v-else-if="issues && issues.length">
      <div
        v-for="issue in issues"
        :key="issue.ID"
        class="mb-6 border border-gray-200 rounded-lg p-4 shadow"
      >
        <!-- Issue Header -->
        <div class="flex justify-between items-center mb-4">
          <h3 class="text-xl font-semibold text-primary">
            {{ issue.Title || 'Untitled Issue' }}
          </h3>
          <div
            v-if="issue.Bounties.some((bounty) => bounty.claims && bounty.claims.length)"
            class="px-3 py-1 bg-yellow-100 text-yellow-800 rounded-full text-sm"
          >
            Claimed
          </div>
        </div>

        <!-- Bounties -->
        <ul v-if="issue.Bounties && issue.Bounties.length" class="space-y-4">
          <li
            v-for="bounty in issue.Bounties"
            :key="bounty.ID"
            class="p-4 bg-secondary rounded-md shadow-md hover:bg-secondary-light"
          >
            <div class="flex justify-between items-center">
              <!-- Bounty Details -->
              <div>
                <span class="text-lg font-semibold text-primary">
                  Reward: {{ bounty.amount || 0 }} {{ bounty.currency || 'USD' }}
                </span>
                <p class="text-gray-400 mt-2">
                  {{ bounty.description || 'No description available.' }}
                </p>
                <div class="mt-2 text-sm text-gray-500">
                  Valid from {{ formatDate(bounty.startAt) }} to {{ formatDate(bounty.endAt) }}
                </div>
              </div>
              <!-- Delete Button -->
              <button
                @click="confirmDeleteBounty(bounty.ID)"
                class="text-red-500 hover:text-red-700 font-semibold"
                :disabled="bounty.claims && bounty.claims.length"
              >
                Delete
              </button>
            </div>

            <!-- Claims Approval Section -->
            <div v-if="bounty.claims && bounty.claims.length" class="mt-6">
              <h2 class="text-lg font-semibold text-primary-light mb-4">Claims for Verification</h2>
              <div
                v-for="claim in bounty.claims"
                :key="claim.id"
                class="bg-secondary rounded-lg p-4 shadow-md mb-4"
              >
                <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
                  <div>
                    <h3 class="text-md font-semibold text-primary">Claim Details</h3>
                    <p class="text-gray-300">{{ claim.claimDetails || 'No details provided' }}</p>
                    <a
                      :href="claim.prUrl"
                      target="_blank"
                      class="text-blue-500 hover:underline mt-2 inline-block"
                    >
                      View Pull Request
                    </a>
                  </div>
                  <div>
                    <h3 class="text-md font-semibold text-primary">PR Verification</h3>
                    <table class="w-full bg-secondary-dark rounded-lg">
                      <tr>
                        <td class="px-4 py-2 text-gray-300">PR Found</td>
                        <td class="px-4 py-2">
                          <span
                            :class="claim.prVerification?.found ? 'text-green-500' : 'text-red-500'"
                          >
                            {{ claim.prVerification?.found ? '✔' : '❌' }}
                          </span>
                        </td>
                      </tr>
                      <tr>
                        <td class="px-4 py-2 text-gray-300">PR Linked</td>
                        <td class="px-4 py-2">
                          <span
                            :class="
                              claim.prVerification?.linked ? 'text-green-500' : 'text-red-500'
                            "
                          >
                            {{ claim.prVerification?.linked ? '✔' : '❌' }}
                          </span>
                        </td>
                      </tr>
                      <tr>
                        <td class="px-4 py-2 text-gray-300">PR Accepted</td>
                        <td class="px-4 py-2">
                          <span
                            :class="
                              claim.prVerification?.accepted ? 'text-green-500' : 'text-red-500'
                            "
                          >
                            {{ claim.prVerification?.accepted ? '✔' : '❌' }}
                          </span>
                        </td>
                      </tr>
                      <tr>
                        <td class="px-4 py-2 text-gray-300">Issue Closed</td>
                        <td class="px-4 py-2">
                          <span
                            :class="
                              claim.prVerification?.closed ? 'text-green-500' : 'text-red-500'
                            "
                          >
                            {{ claim.prVerification?.closed ? '✔' : '❌' }}
                          </span>
                        </td>
                      </tr>
                    </table>
                  </div>
                </div>

                <!-- Claim Actions -->
                <div
                  class="mt-4 flex justify-between items-center space-x-4"
                  v-if="claim.status === 'pending'"
                >
                  <button
                    @click="recheckPR(claim.prUrl, issue)"
                    class="bg-blue-500 text-white px-4 py-2 rounded hover:bg-blue-600 flex items-center"
                  >
                    <svg
                      v-if="claim.recheckLoading"
                      class="animate-spin h-5 w-5 mr-2"
                      xmlns="http://www.w3.org/2000/svg"
                      fill="none"
                      viewBox="0 0 24 24"
                    >
                      <circle
                        class="opacity-25"
                        cx="12"
                        cy="12"
                        r="10"
                        stroke="currentColor"
                        stroke-width="4"
                      ></circle>
                      <path
                        class="opacity-75"
                        fill="currentColor"
                        d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"
                      ></path>
                    </svg>
                    Recheck PR
                  </button>
                  <div class="flex space-x-4">
                    <button
                      @click="approveClaim(claim.ID)"
                      class="bg-green-500 text-white px-4 py-2 rounded hover:bg-green-600"
                    >
                      Approve Claim
                    </button>
                    <button
                      @click="rejectClaim(claim.id)"
                      class="bg-red-500 text-white px-4 py-2 rounded hover:bg-red-600"
                    >
                      Reject Claim
                    </button>
                  </div>
                </div>
              </div>
            </div>
          </li>
        </ul>
      </div>
    </div>

    <!-- No Issues State -->
    <div v-else class="text-center text-lg text-gray-400">No issues or bounties available.</div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { api } from '@/stores/api'
import axios from 'axios' // Ensure axios is available
import { useUserStore } from '../../stores/user'

const userStore = useUserStore()

const issues = ref([])
const loading = ref(true)
const error = ref(null)
const prStatus = ref(null) // PR status information
const prStatusChecked = ref(false) // Flag to track if PR has been checked

const groupedBounties = computed(() => {
  if (!issues.value) return []

  return issues.value
    .filter((issue) => issue && (issue.Bounties?.length || issue.bounties?.length))
    .map((issue) => ({
      id: issue.ID || issue.id || 0,
      title: issue.Title || issue.title || 'Untitled Issue',
      bounties: (issue.Bounties || issue.bounties || []).map((bounty) => ({
        id: bounty.ID || bounty.id || 0,
        amount: bounty.Amount || bounty.amount || 0,
        currency: bounty.Currency || bounty.currency || 'USD',
        description: bounty.Description || bounty.description || 'No description available.',
        startAt: bounty.StartAt || bounty.startAt || '',
        endAt: bounty.EndAt || bounty.endAt || '',
        claims: (bounty.Claims || bounty.claims || []).map((claim) => ({
          id: claim.ID || claim.id || 0,
          pullRequestURL: claim.PullRequestURL || claim.pullRequestURL || '',
          claimDetails: claim.ClaimDetails || claim.claimDetails || '',
          status: claim.Status || claim.status || 'pending',
          prVerification: claim.PRVerification || claim.prVerification || null
        }))
      }))
    }))
})

const fetchBounties = async () => {
  loading.value = true
  error.value = null

  try {
    const response = await api.get('/v1/users/me/issues')

    if (!response || !response.data) {
      throw new Error('No data received')
    }

    // Store the raw response data
    issues.value = response.data

    console.log('Fetched issues:', issues.value) // Debug log
  } catch (fetchError) {
    console.error('Error fetching issues:', fetchError)
    error.value = 'Failed to load bounties. Please try again later.'
  } finally {
    loading.value = false
  }
}

const approveClaim = async (claimId) => {
  try {
    // Find the specific claim to get its details
    let claimDetails = null
    for (const issue of issues.value) {
      for (const bounty of issue.Bounties || []) {
        for (const claim of bounty.claims || []) {
          if (claim.ID === claimId) {
            claimDetails = claim
            break
          }
        }
        if (claimDetails) break
      }
      if (claimDetails) break
    }

    if (!claimDetails) {
      throw new Error('Claim not found')
    }

    if (!claimDetails) {
      throw new Error('Claim not found')
    }

    // Parse PR URL to get repository details
    const prUrlValue = new URL(claimDetails.prUrl)
    const prParts = prUrlValue.pathname.split('/')
    const repoSponsor = prParts[1]
    const repoName = prParts[2]
    const prNumber = prParts[4]

    const approvalData = {
      status: 'approved',
      bountyClaimerCheck: {
        found: claimDetails.prVerification?.found,
        linked: claimDetails.prVerification?.linked,
        accepted: claimDetails.prVerification?.accepted,
        closed: claimDetails.prVerification?.closed,
        authorUsername: claimDetails.authorUsername,
        authorExternalId: claimDetails.authorExternalId,
        repoSponsor: String(repoSponsor),
        repoName: String(repoName),
        prNumber: String(prNumber),
        checkerType: 'OWNER',
        checkerId: userStore.user.id
      }
    }

    console.log('Raw outgoing request:', JSON.stringify(approvalData, null, 2))

    const response = await api.post(`/v1/claims/${claimId}/approve`, approvalData)

    console.log('Claim approval response:', response)
    await fetchBounties() // Refresh data
  } catch (error) {
    console.error('Error approving claim:', error)
    error.value = error.response?.data?.message || 'Failed to approve claim. Please try again.'
  }
}
const rejectClaim = async (claimId) => {
  const reason = window.prompt('Please provide a reason for rejection:')
  if (reason === null) return // User cancelled

  try {
    await api.patch(`/v1/claims/${claimId}`, {
      status: 'rejected',
      rejectReason: reason
    })
    await fetchBounties() // Refresh data
  } catch (error) {
    console.error('Error rejecting claim:', error)
    error.value = 'Failed to reject claim. Please try again.'
  }
}

const confirmDeleteBounty = async (bountyId) => {
  if (!bountyId) return

  const isConfirmed = window.confirm('Are you sure you want to delete this bounty?')

  if (isConfirmed) {
    try {
      // Call API to delete bounty
      await api.delete(`/v1/bounties/${bountyId}`)

      // After successful deletion, re-fetch the bounties to update the view
      await fetchBounties()

    } catch (deleteError) {
      console.error('Error deleting bounty:', deleteError)
      error.value = 'Failed to delete bounty. Please try again.'
    }
  }
}


const formatDate = (date) => {
  if (!date) return 'N/A'
  try {
    return new Date(date).toLocaleDateString()
  } catch {
    return 'Invalid Date'
  }
}

const recheckPR = async (prUrl, issue) => {
  console.log('Checking PR status for URL:', prUrl)

  if (!prUrl) {
    prStatusChecked.value = false
    return
  }

  try {
    const prUrlValue = new URL(prUrl)
    const prParts = prUrlValue.pathname.split('/')

    if (prParts.length < 5 || prParts[3] !== 'pull') {
      console.warn('Invalid PR URL format')
      prStatusChecked.value = false
      return
    }

    const repoSponsor = prParts[1]
    const repoName = prParts[2]
    const prNumber = prParts[4]

    console.log(`Fetching PR status for ${repoSponsor}/${repoName}#${prNumber}`)

    const response = await axios.get(
      `https://api.github.com/repos/${repoSponsor}/${repoName}/pulls/${prNumber}`
    )

    if (response.status === 404) {
      prStatus.value = null
      error.value = `PR #${prNumber} not found in repository ${repoSponsor}/${repoName}. Please verify the PR URL.`
      prStatusChecked.value = true
      return
    }

    const prData = response.data

    prStatus.value = {
      found: prData ? true : false,
      linked: prData && prData.issue_url === issue.value?.issueUrl,
      accepted: prData && prData.merged_at !== null,
      closed: prData && prData.state === 'closed',
      // Add author information
      authorUsername: prData.user?.login || 'Unknown',
      authorExternalId: prData.user?.id || null,
      authorAvatar: prData.user?.avatar_url || null
    }

    prStatusChecked.value = true

    issues.value.forEach((issueItem) => {
      issueItem.Bounties?.forEach((bounty) => {
        bounty.claims?.forEach((claim) => {
          if (claim.prUrl === prUrl) {
            claim.prVerification = {
              found: prData ? true : false,
              linked: prData && prData.issue_url === issue?.issueUrl,
              accepted: prData && prData.merged_at !== null,
              closed: prData && prData.state === 'closed'
            }
          }
        })
      })
    })
    issues.value = [...issues.value]

    console.log('PR Author:', prStatus.value.authorUsername)
    console.log('PR Status:', prStatus.value)
  } catch (err) {
    console.error('Error fetching PR data from GitHub:', err)
    prStatus.value = null
    if (err.response && err.response.status === 404) {
      error.value = 'PR not found. Please check the PR URL.'
    } else {
      error.value = 'Failed to fetch PR status. Please try again later.'
    }
    prStatusChecked.value = true
  }
}

onMounted(fetchBounties)
</script>

<style scoped>
li {
  transition: all 0.3s ease-in-out;
}
button {
  transition: color 0.2s ease-in-out;
}
button:disabled {
  opacity: 0.5;
  cursor: not-allowed;
}
</style>
