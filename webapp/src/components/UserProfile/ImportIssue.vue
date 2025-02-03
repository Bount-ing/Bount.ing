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

    <!-- Bounty Modal -->
    <div
      v-if="isBountyModalOpen"
      class="fixed inset-0 bg-black bg-opacity-50 flex justify-center items-center mt-16"
    >
      <div class="bg-secondary p-6 rounded-lg w-full max-w-lg max-h-[90vh] overflow-y-auto">
        <h3 class="text-2xl font-semibold mb-6 text-center">Set Bounty for Issue</h3>

        <form @submit.prevent="submitBounty">
          <!-- Bounty Amount -->
          <div class="mb-5">
            <label for="bounty" class="block text-sm font-medium mb-2"
              >Bounty Amount (in euros)</label
            >
            <input
              v-model="amount"
              id="bounty"
              type="number"
              min="10"
              required
              class="border px-4 py-2 w-full rounded-lg bg-secondary focus:ring-2 focus:ring-green-500"
              placeholder="Enter bounty amount"
              aria-label="Bounty Amount"
            />
          </div>

          <!-- Start Date -->
          <div class="mb-5">
            <label for="start-date" class="block text-sm font-medium mb-2">Start Date</label>
            <input
              v-model="startAt"
              id="start-date"
              type="date"
              required
              class="border px-4 py-2 w-full rounded-lg bg-secondary focus:ring-2 focus:ring-green-500"
              aria-label="Start Date"
            />
          </div>

          <!-- End Date -->
          <div class="mb-5">
            <label for="end-date" class="block text-sm font-medium mb-2">End Date</label>
            <input
              v-model="endAt"
              id="end-date"
              type="date"
              required
              min="2024-01-01"
              class="border px-4 py-2 w-full rounded-lg bg-secondary focus:ring-2 focus:ring-green-500"
              aria-label="End Date"
            />
          </div>

          <!-- Variables for bounty adjustments -->
          <div
            v-if="variables.length > 0"
            v-for="(variable, index) in variables"
            :key="index"
            class="mb-5"
          >
            <h4 class="text-sm font-semibold mb-2">Variable {{ index + 1 }}</h4>

            <div class="mb-3">
              <label class="block text-sm">Direction</label>
              <select
                v-model="variable.direction"
                required
                class="border px-4 py-2 w-full rounded-lg bg-secondary focus:ring-2 focus:ring-green-500"
                aria-label="Variable Direction"
              >
                <option value="increase">Increase</option>
                <option value="decrease">Decrease</option>
              </select>
            </div>

            <div class="mb-3">
              <label for="variable-amount-{{ index }}" class="block text-sm"
                >Max Amount (in euros)</label
              >
              <input
                v-model="variable.amount"
                id="variable-amount-{{ index }}"
                type="number"
                min="1"
                required
                class="border px-4 py-2 w-full rounded-lg bg-secondary focus:ring-2 focus:ring-green-500"
                placeholder="Amount"
                aria-label="Variable Amount"
              />
            </div>

            <div class="mb-3">
              <label for="variable-start-date-{{ index }}" class="block text-sm">Start Date</label>
              <input
                v-model="variable.startAt"
                id="variable-start-date-{{ index }}"
                type="date"
                required
                class="border px-4 py-2 w-full rounded-lg bg-secondary focus:ring-2 focus:ring-green-500"
                placeholder="Start Date"
                aria-label="Variable Start Date"
              />
            </div>

            <div class="mb-3">
              <label for="variable-end-date-{{ index }}" class="block text-sm">End Date</label>
              <input
                v-model="variable.endAt"
                id="variable-end-date-{{ index }}"
                type="date"
                required
                class="border px-4 py-2 w-full rounded-lg bg-secondary focus:ring-2 focus:ring-green-500"
                placeholder="End Date"
                aria-label="Variable End Date"
              />
            </div>

            <!-- Remove Variable Button -->
            <button
              type="button"
              @click="removeVariable(index)"
              class="text-red-500 mt-2 hover:text-red-700 transition-all"
            >
              Remove Variable
            </button>
          </div>

          <!-- Add Variable Button -->
          <div class="mb-5">
            <button
              type="button"
              @click="addVariable"
              class="px-6 py-3 bg-secondary-light text-white rounded-lg hover:text-secondary-dark transition-all"
            >
              Add Variable
            </button>
          </div>

          <!-- Action Buttons -->
          <div class="flex justify-between mt-6">
            <button
              type="submit"
              class="px-6 py-3 bg-primary text-white rounded-lg hover:bg-primary-light transition-all focus:outline-none focus:ring-2 focus:ring-green-500"
            >
              Set Bounty
            </button>
            <button
              @click="closeBountyModal"
              class="px-6 py-3 bg-secondary-light text-white rounded-lg hover:text-secondary-dark transition-all focus:outline-none focus:ring-2 focus:ring-gray-500"
            >
              Cancel
            </button>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, watch } from 'vue'
import { useUserStore } from '@/stores/user'
import { api } from '@/stores/api'
import GitHubIssueImport from '../GitHubIssueImport.vue'

const issues = ref([])
const userStore = useUserStore()
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

const submitBounty = async () => {
  if (!selectedIssue.value) {
    console.error('No issue selected')
    return
  }

  console.log('Selected issue:', selectedIssue.value)

  try {
    // Proceed with bounty creation if the amount is valid
    if (amount.value >= 10) {
      const formattedStartDate = startAt.value ? new Date(startAt.value).toISOString() : null
      const formattedEndDate = endAt.value ? new Date(endAt.value).toISOString() : null

      const formattedVariables = variables.value.map((variable) => ({
        ...variable,
        startAt: variable.startAt ? new Date(variable.startAt).toISOString() : null,
        endAt: variable.endAt ? new Date(variable.endAt).toISOString() : null
      }))

      const bountyData = {
        amount: amount.value,
        currency: 'EUR',
        issue_id: selectedIssue.value.ID,
        issueUrl: selectedIssue.value.issueUrl,
        issueTitle: selectedIssue.value.title,
        issueBody: selectedIssue.value.body,
        startAt: formattedStartDate,
        endAt: formattedEndDate,
        variables: formattedVariables
      }

      const bountyResponse = await api.post('/v1/bounties', bountyData)

      console.log('Bounty set successfully!')
      closeBountyModal()
    } else {
      console.error('Bounty amount must be at least 10')
    }
  } catch (error) {
    console.error('Error during bounty submission process:', error)
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
