<template>
  <div class="issue-container">
    <GitHubIssueImport />
    <h2 class="text-2xl font-semibold mb-4">Issues</h2>

    <!-- Display message if no issues are found -->
    <div v-if="issues.length === 0" class="text-center text-gray-400">No issues found.</div>

    <!-- Display list of issues -->
    <ul v-else class="space-y-4">
      <li
        v-for="issue in issues"
        :key="issue.id"
        class="bg-secondary p-4 rounded-lg flex items-start space-x-4 shadow-lg hover:shadow-xl transition-shadow duration-300"
      >
        <div v-if="issue && issue.Title">
          <!-- Double coin avatar container -->
          <div class="avatar-container relative flex space-x-2 mr-4">
            <!-- Repository owner avatar, overlaps on the right half of the first avatar -->
            <img
              :src="issue.avatarUrl"
              alt="Repository Owner Avatar"
              class="w-16 h-16 rounded-full border-gray-700 shadow-md"
            />

            <img
              v-if="issue?.user?.avatar_url && issue?.user?.avatar_url !== issue?.repo_avatar"
              :src="issue.user.avatar_url"
              alt="Issue Creator Avatar"
              class="w-10 h-10 rounded-full border-gray-700 absolute -right-4 bottom-0 shadow-md"
            />
          </div>

          <!-- Issue details -->
          <div class="flex-1">
            <div class="text-lg font-medium">
              <a :href="issue.html_url" target="_blank" class="text-blue-400 hover:underline">
                {{ issue.Title }}
              </a>
            </div>
            <p class="text-gray-300 text-sm mb-2">
              {{ issue.body || 'No description available' }}
            </p>

            <!-- Issue status (e.g., open/closed) -->
            <div class="text-sm text-gray-500">
              <span
                v-if="issue?.status"
                :class="{
                  'text-secondary font-semibold bg-primary py-1 px-3 rounded-full':
                    issue.state === 'open',
                  'text-secondary font-semibold bg-error-light py-1 px-3 rounded-full':
                    issue.state === 'closed'
                }"
              >
                {{ issue?.state?.charAt(0).toUpperCase() + issue?.state?.slice(1) || 'Unknown' }}
              </span>
            </div>
          </div>

          <!-- Button to set bounty -->
          <button
            @click="openBountyModal(issue)"
            class="mt-2 px-4 py-2 bg-primary text-secondary font-semibold rounded-full hover:bg-primary-dark transition-all"
          >
            Set Bounty
          </button>
        </div>
      </li>
    </ul>

    <BountyModal
      :is-open="isBountyModalOpen"
      :selected-issue="selectedIssue"
      @close="closeBountyModal"
      @submit="submitBounty"
    />
  </div>
</template>

<script setup>
import GitHubIssueImport from '../GitHubIssueImport.vue'
import BountyModal from '../BountyModal.vue'

import { ref, onMounted, watch } from 'vue'
import { useUserStore } from '@/stores/user'
import { useNotificationStore } from '@/stores/notification'
import { api } from '@/stores/api'

const userStore = useUserStore()
const notificationStore = useNotificationStore()

const issues = ref([])
const isBountyModalOpen = ref(false)
const amount = ref(0)
const startAt = ref('')
const endAt = ref('')
const selectedIssue = ref(null)
const variables = ref([]) // Start with an empty array for variables

onMounted(() => {
  issues.value = userStore.importedIssues || []
})

const openBountyModal = (issue) => {
  if (issue) {
    selectedIssue.value = issue
  } else {
    createIssue()
  }
  isBountyModalOpen.value = true
}

const closeBountyModal = () => {
  isBountyModalOpen.value = false
  amount.value = 0
  startAt.value = ''
  endAt.value = ''
  variables.value = [] // Reset variables
}

const addVariable = () => {
  variables.value.push({ amount: 0, startAt: '', endAt: '', direction: 'increase' })
}

const removeVariable = (index) => {
  variables.value.splice(index, 1)
}

const createIssue = async () => {
  const issueData = {
    title: 'Issue Title',
    body: 'Issue Body'
  }

  try {
    const response = await api.post('/v1/issues', issueData)
    if (response.ok) {
      const createdIssue = await response.json()
      selectedIssue.value = createdIssue
      isBountyModalOpen.value = true
    } else {
      console.error('Failed to create issue', response.statusText)
    }
  } catch (error) {
    console.error('Error creating issue:', error)
  }
}

const submitBounty = async (bountyData) => {
  try {
    if (bountyData.amount < 10) {
      console.log('❌ Bounty amount too low')
      notificationStore.showNotification({
        type: 'error',
        message: 'Bounty amount must be at least $10'
      })
      return
    }

    await api.post('/v1/bounties', bountyData)

    closeBountyModal()
  } catch (error) {
  console.log('❌ Caught error in catch block:', error); // Debugging

  // Check if it's an Axios or fetch-style error
  if (error.response) {
    // If error response is present, check the status code
    if (error.response.status === 500) {
      if (error.response.data && error.response.data.error) {
        const errorMessage = error.response.data.error.toLowerCase();
        
        if (errorMessage.includes('user does not have a stripe account')) {
          notificationStore.showNotification({
            type: 'error',
            message:
              "It seems you don't have any payment methods set up. Please add a payment method to your account. Dashboard > Hosts > Stripe Connect"
          });
        } else {
          notificationStore.showNotification({
            type: 'error',
            message: error.response.data.error
          });
        }
      } else {
        notificationStore.showNotification({
          type: 'error',
          message: 'An unexpected error occurred.'
        });
      }
    } else {
      notificationStore.showNotification({
        type: 'error',
        message: error.response.data.error
      });
    }
  } else if (error.message.toLowerCase().includes('stripe')) {
    notificationStore.showNotification({
      type: 'error',
      message:
        "It seems you don't have any payment methods set up. Please add a payment method to your account. Dashboard > Hosts > Stripe Connect"
    });
  } else {
    notificationStore.showNotification({
      type: 'error',
      message: error.message
    });
  }
}
}

watch(
  () => userStore.importedIssues,
  (newIssues) => {
    console.log('Updated issues:', newIssues)
    issues.value = newIssues || []
  },
  { immediate: true } // Ensures it runs on component mount
)
</script>

<style scoped>
.avatar-container {
  position: relative;
}
.avatar-container img {
  transition: transform 0.3s ease-in-out;
}
</style>
