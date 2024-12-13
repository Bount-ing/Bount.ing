<template>
  <div>
    <h2 class="text-2xl font-semibold mb-4">Bounties</h2>
    
    <!-- Loading State -->
    <div v-if="loading" class="text-center text-xl">Loading bounties...</div>

    <!-- Error State -->
    <div v-else-if="error" class="text-center text-red-500 text-lg">
      {{ error }}
    </div>

    <!-- Bounties Display -->
    <div v-else-if="groupedBounties && groupedBounties.length">
      <div
        v-for="issue in groupedBounties"
        :key="issue.id"
        class="mb-6"
      >
        <h3 class="text-xl font-semibold mb-2 text-primary">
          {{ issue.title || 'Untitled Issue' }}
        </h3>
        <ul v-if="issue.bounties && issue.bounties.length" class="space-y-4">
          <li
            v-for="bounty in issue.bounties"
            :key="bounty.id"
            class="p-4 bg-secondary rounded-md shadow-md hover:bg-secondary-light"
          >
            <div class="flex justify-between items-center">
              <span class="text-lg font-semibold text-primary">
                Reward: {{ bounty.amount || 0 }} {{ bounty.currency || 'USD' }}
              </span>
              <button
                @click="confirmDeleteBounty(bounty.id)"
                class="text-red-500 hover:text-red-700 font-semibold"
              >
                Delete
              </button>
            </div>
            <p class="text-gray-400 mt-2">
              {{ bounty.description || 'No description available.' }}
            </p>
            <div class="mt-2 text-sm text-gray-500">
              Valid from {{ formatDate(bounty.startAt) }} to {{ formatDate(bounty.endAt) }}
            </div>
          </li>
        </ul>
      </div>
    </div>

    <!-- No Bounties State -->
    <div v-else class="text-center text-lg text-gray-400">
      No bounties available.
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { api } from '@/stores/api'

// Reactive state with default values
const issues = ref([])
const loading = ref(true)
const error = ref(null)

// Computed property with safe access
const groupedBounties = computed(() => {
  if (!issues.value) return []
  
  return issues.value
    .filter(issue => issue && issue.bounties && issue.bounties.length)
    .map(issue => ({
      id: issue.id || 0,
      title: issue.title || 'Untitled Issue',
      bounties: (issue.bounties || []).map(bounty => ({
        id: bounty.ID || bounty.id || 0,
        amount: bounty.amount || 0,
        currency: bounty.currency || 'USD',
        description: bounty.Description || bounty.description || 'No description available.',
        startAt: bounty.startAt || '',
        endAt: bounty.endAt || ''
      }))
    }))
})

// Fetch bounties and group them by issue
const fetchBounties = async () => {
  loading.value = true
  error.value = null

  try {
    const response = await api.get('/v1/users/me/issues')
    
    // Validate response
    if (!response || !response.data) {
      throw new Error('No data received')
    }

    issues.value = (response.data || []).map(issue => ({
      id: issue?.ID || issue?.id || 0,
      title: issue?.Title || issue?.title || 'Untitled Issue',
      bounties: (issue?.Bounties || issue?.bounties || []).map(bounty => ({
        id: bounty?.ID || bounty?.id || 0,
        amount: bounty?.amount || 0,
        currency: bounty?.currency || 'USD',
        description: bounty?.Description || bounty?.description || 'No description available.',
        startAt: bounty?.startAt || '',
        endAt: bounty?.endAt || ''
      }))
    }))
  } catch (fetchError) {
    console.error('Error fetching issues:', fetchError)
    error.value = 'Failed to load bounties. Please try again later.'
  } finally {
    loading.value = false
  }
}

// Confirm and delete a bounty
const confirmDeleteBounty = async (bountyId) => {
  if (!bountyId) return

  // Optional: Add a confirmation dialog
  const isConfirmed = window.confirm('Are you sure you want to delete this bounty?')
  
  if (isConfirmed) {
    try {
      await api.delete(`/v1/bounties/${bountyId}`)
      
      // Update local state safely
      issues.value = issues.value
        .map(issue => ({
          ...issue,
          bounties: (issue.bounties || []).filter(bounty => bounty.id !== bountyId)
        }))
        .filter(issue => issue.bounties && issue.bounties.length > 0)
    } catch (deleteError) {
      console.error('Error deleting bounty:', deleteError)
      error.value = 'Failed to delete bounty. Please try again.'
    }
  }
}

// Format date for display with fallback
const formatDate = (date) => {
  if (!date) return 'N/A'
  try {
    return new Date(date).toLocaleDateString()
  } catch {
    return 'Invalid Date'
  }
}

// Fetch bounties when the component is mounted
onMounted(fetchBounties)
</script>

<style scoped>
li {
  transition: all 0.3s ease-in-out;
}
button {
  transition: color 0.2s ease-in-out;
}
</style>