<template>
  <div class="container mx-auto p-6">
    <div class="rounded-lg shadow-lg p-6">
      <!-- Page Header -->
      <h1 class="text-2xl font-bold text-primary mb-4">Claim Issue's Resolution</h1>

      <!-- Loading State -->
      <div v-if="loading" class="text-center py-4">Loading issue details...</div>

      <!-- Content when data is loaded -->
      <template v-else>
        <!-- Issue Details -->
        <div class="bg-secondary p-4 rounded-md shadow-md mb-6">
          <div class="flex items-start gap-4">
            <img
              v-if="issue?.avatarUrl"
              :src="issue.avatarUrl"
              :alt="issue.Title"
              class="w-12 h-12 rounded-full"
            />
            <div>
              <h2 class="text-lg font-semibold text-primary-light">{{ issue?.Title }}</h2>
              <p class="text-gray-300 mt-2" v-if="issue?.Description">
                {{ issue.Description || 'No description provided.' }}
              </p>
              <a
                v-if="issue?.issueUrl"
                :href="issue.issueUrl"
                target="_blank"
                rel="noopener noreferrer"
                class="text-primary hover:text-primary-dark mt-2 inline-block"
              >
                View Original Issue
              </a>
            </div>
          </div>
        </div>

        <!-- Displaying Bounties -->
        <div v-if="bounties?.length > 0" class="mt-6">
          <h2 class="text-lg font-semibold text-primary-light mb-4">Available Bounties</h2>
          <div class="space-y-4">
            <div
              v-for="bounty in bounties"
              :key="bounty.ID"
              class="bg-secondary rounded-lg p-4 shadow-md"
            >
              <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
                <div>
                  <div class="flex items-baseline gap-2 mb-2">
                    <span class="text-2xl font-bold text-primary">
                      {{ bounty.amount }} {{ bounty.currency }}
                    </span>
                  </div>
                  <div class="text-gray-300 space-y-1">
                    <div class="flex items-center gap-2">
                      <span class="font-medium">Start Date:</span>
                      {{ formatDate(bounty.startAt) }}
                    </div>
                    <div class="flex items-center gap-2">
                      <span class="font-medium">End Date:</span>
                      {{ formatDate(bounty.endAt) }}
                    </div>
                    <div class="flex items-center gap-2">
                      <span class="font-medium">Status:</span>
                      <span class="capitalize">{{ bounty.status || 'Open' }}</span>
                    </div>
                  </div>
                </div>
                <div class="text-gray-300">
                  <div class="mb-2">
                    <span class="font-medium">Bounty ID:</span> #{{ bounty.ID }}
                  </div>
                  <div class="mb-2">
                    <span class="font-medium">Created:</span> {{ formatDate(bounty.CreatedAt) }}
                  </div>
                  <div v-if="bounty.finalized_at !== '0001-01-01T00:00:00Z'">
                    <span class="font-medium">Finalized:</span>
                    {{ formatDate(bounty.finalized_at) }}
                  </div>
                </div>
              </div>

              <!-- Bounty Claimer and Owner Status -->
              <div class="mt-4">
                <h3 class="text-lg font-semibold text-primary-light">
                  Bounty Claimer and Owner Status
                </h3>
                <div class="grid grid-cols-1 md:grid-cols-2 gap-6 mt-4">
                  <!-- Bounty Claimer Status -->
                  <div>
                    <h4 class="text-md font-semibold text-primary">Bounty Claimer</h4>
                    <div class="space-y-2">
                      <div class="flex items-center gap-2">
                        <span class="font-medium">PR has been found:</span>
                        <span v-if="bounty.claimer?.prFound" class="text-green-500">✔</span>
                        <span v-else class="text-red-500">❌</span>
                      </div>
                      <div class="flex items-center gap-2">
                        <span class="font-medium">PR links to the issue:</span>
                        <span v-if="bounty.claimer?.prLinked" class="text-green-500">✔</span>
                        <span v-else class="text-red-500">❌</span>
                      </div>
                      <div class="flex items-center gap-2">
                        <span class="font-medium">PR has been accepted:</span>
                        <span v-if="bounty.claimer?.prAccepted" class="text-green-500">✔</span>
                        <span v-else class="text-red-500">❌</span>
                      </div>
                      <div class="flex items-center gap-2">
                        <span class="font-medium">Issue has been closed:</span>
                        <span v-if="bounty.claimer?.issueClosed" class="text-green-500">✔</span>
                        <span v-else class="text-red-500">❌</span>
                      </div>
                    </div>
                  </div>

                  <!-- Bounty Owner Status -->
                  <div>
                    <h4 class="text-md font-semibold text-primary">Bounty Owner</h4>
                    <div class="space-y-2">
                      <div class="flex items-center gap-2">
                        <span class="font-medium">PR has been found:</span>
                        <span v-if="bounty.owner?.prFound" class="text-green-500">✔</span>
                        <span v-else class="text-red-500">❌</span>
                      </div>
                      <div class="flex items-center gap-2">
                        <span class="font-medium">PR links to the issue:</span>
                        <span v-if="bounty.owner?.prLinked" class="text-green-500">✔</span>
                        <span v-else class="text-red-500">❌</span>
                      </div>
                      <div class="flex items-center gap-2">
                        <span class="font-medium">PR has been accepted:</span>
                        <span v-if="bounty.owner?.prAccepted" class="text-green-500">✔</span>
                        <span v-else class="text-red-500">❌</span>
                      </div>
                      <div class="flex items-center gap-2">
                        <span class="font-medium">Issue has been closed:</span>
                        <span v-if="bounty.owner?.issueClosed" class="text-green-500">✔</span>
                        <span v-else class="text-red-500">❌</span>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- Claim Form -->
        <form @submit.prevent="submitClaim" class="mt-6">
          <!-- PR URL Field -->
          <div class="mb-4">
            <label for="prUrl" class="block text-gray-300 font-semibold mb-2">PR URL</label>
            <div class="flex items-center gap-4">
              <input
                id="prUrl"
                v-model="prUrl"
                type="url"
                class="w-full bg-secondary-dark border border-gray-600 rounded-lg p-3 focus:ring-primary focus:border-primary text-gray-200"
                placeholder="Enter the PR URL"
              />
              <button
                @click="checkPrStatus"
                type="button"
                class="bg-primary text-white px-4 py-2 rounded-lg shadow-md hover:bg-primary-dark transition duration-200"
              >
                Verify
              </button>
            </div>

            <!-- PR Status Section -->
            <div v-if="prStatusChecked" class="mt-4 space-y-4">
              <h3 class="text-lg font-semibold text-primary-light">PR Status</h3>

              <div v-if="prStatus" class="space-y-4">
                <!-- Add PR Author Information -->
                <div class="flex items-center justify-between">
                  <span class="font-medium">PR Author:</span>
                  <div class="flex items-center gap-2">
                    <img
                      v-if="prStatus.authorAvatar"
                      :src="prStatus.authorAvatar"
                      :alt="prStatus.authorUsername"
                      class="w-6 h-6 rounded-full"
                    />
                    <span class="text-primary">{{ prStatus.authorUsername }}</span>
                  </div>
                </div>
                <!-- PR Found -->
                <div class="flex items-center justify-between">
                  <span class="font-medium">PR has been found:</span>
                  <span v-if="prStatus.found" class="text-green-500 text-lg">✔</span>
                </div>

                <!-- PR Linked to Issue -->
                <div class="flex items-center justify-between">
                  <span class="font-medium">PR links to the issue:</span>
                  <span v-if="prStatus.linked" class="text-green-500 text-lg">✔</span>
                  <span v-else class="text-red-500 text-lg">❌</span>
                </div>

                <!-- PR Accepted -->
                <div class="flex items-center justify-between">
                  <span class="font-medium">PR has been accepted:</span>
                  <span v-if="prStatus.accepted" class="text-green-500 text-lg">✔</span>
                  <span v-else class="text-red-500 text-lg">❌</span>
                </div>

                <!-- Issue Closed -->
                <div class="flex items-center justify-between">
                  <span class="font-medium">Issue has been closed:</span>
                  <span v-if="prStatus.closed" class="text-green-500 text-lg">✔</span>
                  <span v-else class="text-red-500 text-lg">❌</span>
                </div>
              </div>

              <!-- PR Not Found Error -->
              <div v-else-if="error" class="text-red-500 flex items-center gap-2">
                <span class="font-medium">PR has not been found:</span>
                <span class="text-red-500 text-lg">❌</span>
              </div>
            </div>
          </div>

          <div class="mb-4">
            <label for="claimDetails" class="block text-gray-300 font-semibold mb-2">
              Claim Details
            </label>
            <textarea
              id="claimDetails"
              v-model="claimDetails"
              class="w-full bg-secondary-dark border border-gray-600 rounded-lg p-3 focus:ring-primary focus:border-primary text-gray-200"
              placeholder="Provide details about your claim..."
              rows="5"
            ></textarea>
          </div>

          <div class="flex justify-end">
            <button
              type="submit"
              :disabled="submitting"
              class="bg-success text-white py-2 px-4 rounded-lg shadow-md hover:bg-success-dark transition duration-200 disabled:opacity-50"
            >
              {{ submitting ? 'Submitting...' : 'Submit Claim' }}
            </button>
          </div>
        </form>
      </template>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, watch } from 'vue'
import { useRoute } from 'vue-router'
import { api } from '@/stores/api'
import axios from 'axios' // Ensure axios is available

const route = useRoute()

// Add debug data ref
const debugData = ref({})

// Refs for issue and bounty
const issue = ref(null)
const bounties = ref([])
const claimDetails = ref('')
const prUrl = ref('') // PR URL ref
const prStatus = ref(null) // PR status information
const prStatusChecked = ref(false) // Flag to track if PR has been checked
const loading = ref(true)
const error = ref(null)
const submitting = ref(false)
const submitError = ref(null)

// Format date helper
const formatDate = (date) => {
  if (!date) return 'N/A'
  return new Date(date).toLocaleString()
}

// Fetching issue and bounty data
const fetchIssueDataAndBounties = async () => {
  const id = route.params.id

  if (!id) {
    error.value = 'Missing issue ID'
    loading.value = false
    return
  }

  try {
    loading.value = true
    error.value = null

    const response = await api.get(`/v1/issues/${id}/bounties`)

    console.log('Raw API Response:', response)
    console.log('Response Data:', response.data)

    // Store debug data
    debugData.value = {
      responseData: response.data,
      responseStatus: response.status,
      endpoint: `/v1/issues/${id}/bounties`
    }

    if (!response?.data) {
      throw new Error('Invalid response from server')
    }

    issue.value = response.data
    if (response.data.Bounties) {
      bounties.value = response.data.Bounties
    } else if (Array.isArray(response.data)) {
      bounties.value = response.data
    } else {
      console.warn('Unexpected bounties data structure:', response.data)
    }

    console.log('Processed Data:', {
      issue: issue.value,
      bounties: bounties.value
    })
  } catch (err) {
    console.error('Error details:', {
      message: err.message,
      response: err.response,
      stack: err.stack
    })
    error.value = 'Failed to load issue details. Please try again later.'
  } finally {
    loading.value = false
  }
}

const checkPrStatus = async () => {
  console.log('Checking PR status for URL:', prUrl.value)

  if (!prUrl.value) {
    prStatusChecked.value = false
    return
  }

  try {
    const prUrlValue = new URL(prUrl.value)
    const prParts = prUrlValue.pathname.split('/')

    if (prParts.length < 5 || prParts[3] !== 'pull') {
      console.warn('Invalid PR URL format')
      prStatusChecked.value = false
      return
    }

    const repoOwner = prParts[1]
    const repoName = prParts[2]
    const prNumber = prParts[4]

    console.log(`Fetching PR status for ${repoOwner}/${repoName}#${prNumber}`)

    const response = await axios.get(
      `https://api.github.com/repos/${repoOwner}/${repoName}/pulls/${prNumber}`
    )

    if (response.status === 404) {
      prStatus.value = null
      error.value = `PR #${prNumber} not found in repository ${repoOwner}/${repoName}. Please verify the PR URL.`
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
      authorAvatar: prData.user?.avatar_url || null
    }

    prStatusChecked.value = true
    
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

const submitClaim = async () => {
  if (!prStatus || !prStatusChecked) {
    submitError.value = 'Please verify the PR before submitting the claim.'
    return
  }

  submitting.value = true
  submitError.value = null

  try {
    // Extract PR details from URL
    const prUrlValue = new URL(prUrl.value)
    const prParts = prUrlValue.pathname.split('/')
    const repoOwner = prParts[1]
    const repoName = prParts[2]
    const prNumber = prParts[4]

    // Prepare claim data
    const claimData = {
      bountyId: route.params.id, // Assuming this is the bounty ID
      prUrl: prUrl.value,
      claimDetails: claimDetails.value,
      prVerification: {
        found: prStatus.value.found,
        linked: prStatus.value.linked,
        accepted: prStatus.value.accepted,
        closed: prStatus.value.closed,
        authorUsername: prStatus.value.authorUsername,
        repoOwner,
        repoName,
        prNumber
      }
    }

    console.log('Submitting claim with data:', claimData)

    // Make API request to submit claim
    const response = await api.post('/v1/claims', claimData)

    console.log('Claim submission response:', response)

    // Reset form after successful submission
    prUrl.value = ''
    claimDetails.value = ''
    prStatus.value = null
    prStatusChecked.value = false

    // Show success message (you might want to handle this differently)
    alert('Claim submitted successfully!')

  } catch (err) {
    console.error('Error submitting claim:', err)
    submitError.value = err.response?.data?.message || 'Failed to submit claim. Please try again.'
  } finally {
    submitting.value = false
  }
}


onMounted(() => {
  fetchIssueDataAndBounties()
})
</script>
