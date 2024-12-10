<template>
  <div>
    <h2 class="text-2xl font-semibold mb-4">Issues</h2>

    <!-- If there are no issues, display a message -->
    <div v-if="issues.length === 0" class="text-gray-400">
      No issues found.
    </div>

    <!-- List of issues -->
    <ul class="space-y-4">
      <li v-for="issue in issues" :key="issue.id" class="p-4 bg-gray-700 rounded-lg hover:bg-gray-600 transition">
        <div class="flex justify-between">
          <h3 class="text-xl font-semibold">{{ issue.title }}</h3>
          <span v-if="issue.status === 'open'" class="bg-green-500 text-white px-2 py-1 rounded-full">
            Open
          </span>
          <span v-else class="bg-red-500 text-white px-2 py-1 rounded-full">
            Closed
          </span>
        </div>
        <p class="text-gray-300">{{ issue.description }}</p>
        <div class="text-sm text-gray-500 mt-2">
          Created on: {{ new Date(issue.createdAt).toLocaleDateString() }}
        </div>
      </li>
    </ul>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { api } from '@/stores/api'; // Import the API client

// Reactive state variables
const issues = ref([]);

// Fetch issues data on component mount
const fetchIssues = async () => {
  try {
    const response = await api.get('/v1/issues'); // Replace with your API endpoint
    issues.value = response.data; // Assuming the API returns an array of issues
  } catch (error) {
    console.error('Error fetching issues:', error);
  }
};

// Use the lifecycle hook to fetch data
onMounted(fetchIssues);
</script>
