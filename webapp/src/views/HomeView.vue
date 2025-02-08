<template>
  <div class="py-4 my-4">
    <h2 class="text-3xl font-semibold mx-4 my-2">Open Bounties</h2>
    <h3 class="text-xl font-semibold mx-4 my-2">
      Bounties are rewards offered for solving specific issues.
    </h3>
    <div v-if="groupedBounties.length === 0" class="text-center text-gray-400">
      No bounties found.
    </div>
    <ul v-else class="space-y-4 p-4">
      <BountyItem
        v-for="(group, index) in groupedBounties"
        :key="index"
        :group="group"
        :index="index"
        :calculateCurrentAmount="calculateCurrentAmount"
        @raise-bounty-modal="raiseBountyModal"
      />
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
import { ref, onMounted, onUnmounted } from 'vue'
import { api } from '@/stores/api'
import { useErrorStore } from '@/stores/errors'
import BountyItem from '@/components/BountyItem.vue'
import BountyModal from '@/components/BountyModal.vue'

const errorStore = useErrorStore()

const isBountyModalOpen = ref(false)
const selectedIssue = ref(null)

const groupedBounties = ref([]) // Reactive state to hold grouped bounties
let intervalId = null // Variable to store the interval ID for cleanup

async function fetchBounties() {
  try {
    const response = await api.get('/v1/public-bounties-issue')
    const issues = response.data

    // Group bounties by issueUrl, filtering out issues without bounties
    const grouped = issues
      .filter((issue) => issue.Bounties && issue.Bounties.length > 0)
      .map((issue) => ({
        ...issue,
        currentAmount: calculateCurrentAmount(issue)
      }))

    groupedBounties.value = grouped

    console.log('Grouped Bounties:', JSON.stringify(groupedBounties.value, null, 2))
  } catch (error) {
    console.error('Error fetching bounties:', error)
  }
}

function calculateCurrentAmount(issue) {
  // Check if the issue has any bounties
  if (!issue.Bounties || issue.Bounties.length === 0) {
    console.log(`No bounties for issue: ${issue.issueUrl}`)
    return 0
  }

  // Calculate total amount for all bounties in this issue
  let totalAmount = 0

  issue.Bounties.forEach((bounty) => {
    totalAmount += calculateTotalAmount(bounty)
  })

  return totalAmount
}

function calculateTotalAmount(bounty) {
  if (!bounty) {
    console.error('Bounty is undefined or null')
    return 0
  }

  // Ensure amount is a valid number
  const amount = bounty.amount !== undefined && bounty.amount !== null ? Number(bounty.amount) : 0

  if (isNaN(amount)) {
    console.error(
      `Invalid amount for bounty: ${bounty.issueUrl}`,
      `Raw amount: ${bounty.amount}`,
      `Parsed amount: ${amount}`,
      'Full bounty object:',
      JSON.stringify(bounty, null, 2)
    )
    return 0
  }

  // Check if the bounty itself is active
  if (!isActive(bounty.startAt, bounty.endAt)) {
    console.log(`Bounty not active: ${bounty.issueUrl}`)
    return 0 // Return 0 if not active
  }

  let baseAmount = amount

  // Process variables if they exist
  if (bounty.variables && Array.isArray(bounty.variables)) {
    const now = new Date()

    bounty.variables.forEach((variable) => {
      // Validate variable amount
      const variableAmount =
        variable.amount !== undefined && variable.amount !== null ? Number(variable.amount) : 0

      if (isNaN(variableAmount)) {
        console.error(
          `Invalid variable amount for bounty: ${bounty.issueUrl}`,
          `Raw variable amount: ${variable.amount}`,
          `Parsed amount: ${variableAmount}`
        )
        return
      }

      // Correct date parsing and handling
      const start = variable.startAt ? new Date(variable.startAt) : null
      const end = variable.endAt ? new Date(variable.endAt) : null

      // Check if the variable is within the valid time range
      if (start && end && now >= start && now <= end) {
        const totalDuration = end.getTime() - start.getTime()
        const elapsedDuration = now.getTime() - start.getTime()

        const proportion =
          totalDuration > 0 ? Math.min(1, Math.max(0, elapsedDuration / totalDuration)) : 0

        const currentAmount = variableAmount * proportion

        if (variable.direction === 'increase') {
          baseAmount += currentAmount
        } else if (variable.direction === 'decrease') {
          baseAmount += variableAmount - currentAmount
        }
      }
    })
  }

  // Ensure the amount is never negative
  const finalAmount = Math.max(0, baseAmount)

  console.log(`Bounty ${bounty.issueUrl} - Base: ${amount}, Final: ${finalAmount}`)

  return finalAmount
}

// Helper function to check if the current date is within a given range
function isActive(startAt, endAt) {
  const now = new Date()
  const start = startAt ? new Date(startAt) : null
  const end = endAt ? new Date(endAt) : null

  if (start && now < start) return false // Not started yet
  if (end && now > end) return false // Already ended
  return true // Active
}

const raiseBountyModal = (issue) => {
  selectedIssue.value = issue
  isBountyModalOpen.value = true
}

const closeBountyModal = () => {
  isBountyModalOpen.value = false
  selectedIssue.value = null
  //refresh bounties
  fetchBounties()
}

const submitBounty = async (bountyData) => {
  try {
    if (bountyData.amount < 10) {
      console.log('❌ Bounty amount too low')
      errorStore.showError('Bounty amount must be at least 10')
      return
    }

    await api.post('/v1/bounties', bountyData)

    closeBountyModal()
  } catch (error) {
    console.log('❌ Caught error in catch block:', error) // Debugging

    // Check if it's an Axios or fetch-style error
    if (error.response) {
      // If error response is present, check the status code
      if (error.response.status === 500) {
        if (error.response.data && error.response.data.error) {
          const errorMessage = error.response.data.error.toLowerCase()

          if (errorMessage.includes('user does not have a stripe account')) {
            errorStore.showError(
              "It seems you don't have a Stripe account connected. Please create one or link your account to proceed. Dashboard > Hosts > Connect Stripe. Then go to payment methods to add one."
            )
          } else {
            errorStore.showError('There was a server issue. Please try again later.')
          }
        } else {
          errorStore.showError('An unexpected server error occurred.')
        }
      } else {
        // If other HTTP status codes, handle them appropriately
        errorStore.showError(error.response.data.message || 'An unexpected error occurred.')
      }
    } else if (error.message.toLowerCase().includes('stripe')) {
      errorStore.showError(
        "It seems you don't have any payment methods set up. Please add a payment method to your account."
      )
    } else if (error.response.status === 401 || error.response.status === 403) {
      errorStore.showError(
        'You are not authorized to perform this action. Make sure you are logged in.'
      )
    } else {
      errorStore.showError(error.message || 'An unexpected error occurred')
    }
  }
}

// Set up an interval to recalculate every second
onMounted(() => {
  fetchBounties() // Initial fetch

  groupedBounties.value = [...groupedBounties.value] // Force reactivity update
})

// Clean up interval when the component is unmounted
onUnmounted(() => {
  if (intervalId) {
    clearInterval(intervalId) // Prevent memory leaks
  }
})
</script>
