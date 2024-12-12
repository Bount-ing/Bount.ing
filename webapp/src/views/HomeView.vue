<template>
  <div>
    <h2 class="text-3xl font-semibold mb-4 text-gray-900">Bounties</h2>
    <div v-if="groupedBounties.length === 0" class="text-center text-gray-400">
      No bounties found.
    </div>
    <ul v-else class="space-y-4 p-4">
      <!-- Display each group -->
      <li 
        v-for="(group, index) in groupedBounties" 
        :key="index" 
        class="bg-secondary p-4 rounded-lg flex items-center space-x-4 shadow-md hover:shadow-lg transition-shadow duration-300 ease-in-out"
      >
        <!-- Avatar Section with Soft Background -->
        <div v-if="group.avatarUrl" class="flex-shrink-0">
          <img :src="group.avatarUrl" alt="Avatar" class="rounded-full w-14 h-14 border-2 border-primary" />
        </div>

        <!-- Content Section -->
        <div class="flex-1">
          <!-- Title and Description Section -->
          <div class="mb-2">
            <h3 class="text-xl font-semibold text-primary-light">{{ group.Title }}</h3>
            <p v-if="group.Description" class="text-gray-600 mt-1 text-sm">{{ group.Description }}</p>
            <p v-else class="text-gray-500 mt-1 text-sm italic">No description provided.</p>
          </div>

          <!-- Issue URL Section -->
          <div class="mb-2">
            <a 
              :href="group.issueUrl" 
              target="_blank" 
              rel="noopener noreferrer" 
              class="text-primary hover:text-primary-dark text-sm"
            >
              View Issue
            </a>
          </div>
        </div>

        <!-- Right-aligned Section for Bounty Amount and Timing -->
        <div class="ml-auto flex flex-col items-end space-y-2">
          <!-- Bounty Amount as a Button -->
          <div v-if="group.Bounties.length" class="flex items-center justify-between mb-2">
            <button class="bg-secondary text-primary-light border-primary-light border py-2 px-6 rounded-full text-sm font-semibold shadow-md transform transition-transform duration-200 hover:scale-105">
              {{ calculateCurrentAmount(group).toFixed(2) }} €
            </button>
          </div>

          <!-- Bounty Status and Timing -->
          <div v-if="group.Bounties.length" class="text-sm text-gray-500 space-y-1">
            <div v-for="(bounty, index) in group.Bounties" :key="index" class="flex justify-between items-center">
              <span>Expires: {{ new Date(bounty.endAt).toLocaleDateString() }}</span>
            </div>
          </div>
        </div>
      </li>
    </ul>
  </div>
</template>


<script setup>
import { ref, onMounted, onUnmounted } from 'vue';
import { api } from '@/stores/api';

const groupedBounties = ref([]); // Reactive state to hold grouped bounties
let intervalId = null; // Variable to store the interval ID for cleanup


async function fetchBounties() {
  try {
    const response = await api.get('/v1/public-bounties-issue');
    const issues = response.data;

    // Group bounties by issueUrl, filtering out issues without bounties
    const grouped = issues
      .filter(issue => issue.Bounties && issue.Bounties.length > 0)
      .map(issue => ({
        ...issue,
        currentAmount: calculateCurrentAmount(issue)
      }));

    groupedBounties.value = grouped;
    
    console.log('Grouped Bounties:', JSON.stringify(groupedBounties.value, null, 2));
  } catch (error) {
    console.error('Error fetching bounties:', error);
  }
}

function calculateCurrentAmount(issue) {
  // Check if the issue has any bounties
  if (!issue.Bounties || issue.Bounties.length === 0) {
    console.log(`No bounties for issue: ${issue.issueUrl}`);
    return 0;
  }

  // Calculate total amount for all bounties in this issue
  let totalAmount = 0;

  issue.Bounties.forEach(bounty => {
    totalAmount += calculateTotalAmount(bounty);
  });

  return totalAmount;
}

function calculateTotalAmount(bounty) {
  if (!bounty) {
    console.error('Bounty is undefined or null');
    return 0;
  }

  // Ensure amount is a valid number
  const amount = bounty.amount !== undefined && bounty.amount !== null 
    ? Number(bounty.amount) 
    : 0;
  
  if (isNaN(amount)) {
    console.error(`Invalid amount for bounty: ${bounty.issueUrl}`, 
      `Raw amount: ${bounty.amount}`, 
      `Parsed amount: ${amount}`,
      'Full bounty object:', 
      JSON.stringify(bounty, null, 2)
    );
    return 0;
  }

  // Check if the bounty itself is active
  if (!isActive(bounty.startAt, bounty.endAt)) {
    console.log(`Bounty not active: ${bounty.issueUrl}`);
    return 0;  // Return 0 if not active
  }

  let baseAmount = amount;

  // Process variables if they exist
  if (bounty.variables && Array.isArray(bounty.variables)) {
    const now = new Date();

    bounty.variables.forEach(variable => {
      // Validate variable amount
      const variableAmount = variable.amount !== undefined && variable.amount !== null 
        ? Number(variable.amount) 
        : 0;
      
      if (isNaN(variableAmount)) {
        console.error(`Invalid variable amount for bounty: ${bounty.issueUrl}`, 
          `Raw variable amount: ${variable.amount}`, 
          `Parsed amount: ${variableAmount}`
        );
        return;
      }

      // Correct date parsing and handling
      const start = variable.startAt ? new Date(variable.startAt) : null;
      const end = variable.endAt ? new Date(variable.endAt) : null;

      // Check if the variable is within the valid time range
      if (start && end && now >= start && now <= end) {
        const totalDuration = end.getTime() - start.getTime();
        const elapsedDuration = now.getTime() - start.getTime();

        const proportion = totalDuration > 0 ? Math.min(1, Math.max(0, elapsedDuration / totalDuration)) : 0;

        const currentAmount = variableAmount * proportion;

        if (variable.direction === 'increase') {
          baseAmount += currentAmount;
        } else if (variable.direction === 'decrease') {
          baseAmount -= currentAmount;
        }
      }
    });
  }

  // Ensure the amount is never negative
  const finalAmount = Math.max(0, baseAmount);
  
  console.log(`Bounty ${bounty.issueUrl} - Base: ${amount}, Final: ${finalAmount}`);

  return finalAmount;
}

// Helper function to check if the current date is within a given range
function isActive(startAt, endAt) {
  const now = new Date();
  const start = startAt ? new Date(startAt) : null;
  const end = endAt ? new Date(endAt) : null;

  if (start && now < start) return false; // Not started yet
  if (end && now > end) return false;    // Already ended
  return true;                           // Active
}
// Set up an interval to recalculate every second
onMounted(() => {
  fetchBounties(); // Initial fetch

  intervalId = setInterval(() => {
    groupedBounties.value = [...groupedBounties.value]; // Force reactivity update
  }, 1000);
});

// Clean up interval when the component is unmounted
onUnmounted(() => {
  if (intervalId) {
    clearInterval(intervalId); // Prevent memory leaks
  }
});
</script>
