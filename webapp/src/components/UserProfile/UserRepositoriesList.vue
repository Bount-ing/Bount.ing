<template>
  <div>
    <h2 class="text-2xl font-semibold mb-4">Repositories</h2>
    
    <!-- Loading state -->
    <div v-if="loading" class="text-center text-xl">Loading repositories...</div>

    <!-- Error state -->
    <div v-if="error" class="text-center text-xl text-red-500">
      Error loading repositories: {{ error }}
    </div>

    <!-- Repository list -->
    <ul v-if="!loading && !error" class="space-y-4">
      <RepoLine
        v-for="repo in repositories"
        :key="repo.ID"
        :repo="transformRepository(repo)"
      />
    </ul>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { api } from '@/stores/api'; // Import the API client
import RepoLine from './RepoLine.vue'; // Import the RepoLine component

// Reactive state variables
const repositories = ref([]);
const loading = ref(true);
const error = ref(null);

// Transform the API data to match RepoLine's expected structure
const transformRepository = (repo) => ({
  id: repo.ID,
  name: repo.Name || 'Unnamed Repository',
  description: 'No description available', // Adjust when description data is available
  stars: 0, // Replace with actual stars count if API provides it
  forks: 0, // Replace with actual forks count if API provides it
  url: repo.GithubURL || '#',
});

// Fetch repositories data on component mount
const fetchRepositories = async () => {
  try {
    const response = await api.get('/v1/repositories'); // Replace with your API endpoint
    repositories.value = response.data; // Assuming the API returns an array of repositories
  } catch (error) {
    console.error('Error fetching repositories:', error);
    error.value = error.message;
  } finally {
    loading.value = false;
  }
};

// Use the lifecycle hook to fetch data
onMounted(fetchRepositories);
</script>

<style scoped>
/* Add styles specific to the main component if needed */
</style>
