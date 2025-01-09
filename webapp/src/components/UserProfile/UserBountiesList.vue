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
      <div v-for="issue in issues" :key="issue.ID" class="mb-6 border border-gray-200 rounded-lg p-4 shadow">
        <!-- Issue Header -->
        <div class="flex justify-between items-center mb-4">
          <h3 class="text-xl font-semibold text-primary">
            {{ issue.Title || 'Untitled Issue' }}
          </h3>
          <div
            v-if="issue.Bounties.some(bounty => bounty.claims && bounty.claims.length)"
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

            <!-- Claims -->
            <div v-if="bounty.claims && bounty.claims.length" class="mt-4">
              <h4 class="text-lg font-medium text-primary mb-2">Claims:</h4>
              <div
                v-for="claim in bounty.claims"
                :key="claim.ID"
                class="border-t border-gray-200 pt-2 mt-2"
              >
                <div class="flex justify-between items-center">
                  <div>
                    <a
                      :href="claim.prUrl"
                      target="_blank"
                      class="text-blue-600 hover:underline"
                    >
                      View Pull Request
                    </a>
                    <span class="ml-2 text-sm text-gray-600">
                      Status: {{ claim.status || 'Pending' }}
                    </span>
                  </div>
                  <!-- Actions -->
                  <div v-if="claim.status === 'pending'" class="space-x-2">
                    <button
                      @click="approveClaim(claim.ID)"
                      class="px-3 py-1 bg-green-500 text-white rounded hover:bg-green-600"
                    >
                      Approve
                    </button>
                    <button
                      @click="rejectClaim(claim.ID)"
                      class="px-3 py-1 bg-red-500 text-white rounded hover:bg-red-600"
                    >
                      Reject
                    </button>
                  </div>
                </div>
                <p v-if="claim.claimDetails" class="mt-2 text-sm text-gray-600">
                  {{ claim.claimDetails }}
                </p>
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

const issues = ref([])
const loading = ref(true)
const error = ref(null)

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
    await api.patch(`/v1/claims/${claimId}`, {
      status: 'approved'
    })
    await fetchBounties() // Refresh data
  } catch (error) {
    console.error('Error approving claim:', error)
    error.value = 'Failed to approve claim. Please try again.'
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
      await api.delete(`/v1/bounties/${bountyId}`)
      issues.value = issues.value
        .map((issue) => ({
          ...issue,
          bounties: (issue.bounties || []).filter((bounty) => bounty.id !== bountyId)
        }))
        .filter((issue) => issue.bounties && issue.bounties.length > 0)
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
