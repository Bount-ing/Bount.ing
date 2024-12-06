<template>
    <div>
      <h2 class="text-2xl font-semibold mb-4">Bounties</h2>
      <!-- Display loading text if bounties are still being fetched -->
      <div v-if="loading" class="text-center text-xl">Loading bounties...</div>
      
      <!-- Display bounties list -->
      <ul v-else class="space-y-4">
        <li
          v-for="bounty in bounties"
          :key="bounty.id"
          class="p-4 bg-secondary rounded-md shadow-md hover:bg-secondary-light"
        >
          <div class="flex justify-between items-center">
            <span class="text-lg font-semibold text-primary">{{ bounty.title }}</span>
            <span :class="bounty.statusClass" class="text-sm px-2 py-1 rounded-md">
              {{ bounty.status }}
            </span>
          </div>
          <p class="text-gray-400 mt-2">{{ bounty.description }}</p>
          <div class="mt-2 text-right">
            <span class="font-semibold text-primary-light">{{ bounty.reward }} USD</span>
          </div>
        </li>
      </ul>
  
      <!-- If no bounties are available -->
      <div v-if="!bounties.length && !loading" class="text-center text-lg text-gray-400">
        No bounties available.
      </div>
    </div>
  </template>
  
  <script setup>
  import { ref, onMounted } from 'vue';
  
  // Sample state to simulate fetching bounties
  const bounties = ref([]);
  const loading = ref(true);
  
  // Simulate an API call to fetch bounties
  const fetchBounties = async () => {
    // Simulate a delay for loading
    setTimeout(() => {
      bounties.value = [
        {
          id: 1,
          title: 'Fix critical bug in API',
          description: 'Resolve the performance issue in the API causing delays.',
          status: 'Open',
          statusClass: 'bg-info', // Classes for different statuses
          reward: 100,
        },
        {
          id: 2,
          title: 'Create new feature for web app',
          description: 'Implement a new user authentication system for the web app.',
          status: 'In Progress',
          statusClass: 'bg-yellow-500',
          reward: 150,
        },
        {
          id: 3,
          title: 'Refactor UI for mobile app',
          description: 'Improve the UI/UX of the mobile app for better user experience.',
          status: 'Completed',
          statusClass: 'bg-green-500',
          reward: 200,
        },
      ];
      loading.value = false;
    }, 1000); // Simulate a delay of 1 second
  };
  
  // Fetch bounties when the component is mounted
  onMounted(fetchBounties);
  </script>
  
  <style scoped>
  /* Additional styles for better presentation */
  li {
    transition: all 0.3s ease-in-out;
  }

  </style>
  