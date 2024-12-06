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
        <li
          v-for="repo in repositories"
          :key="repo.id"
          class="bg-gray-700 p-4 rounded-lg hover:bg-gray-600 cursor-pointer"
        >
          <div class="text-lg font-semibold">{{ repo.name }}</div>
          <div class="text-sm text-gray-400">{{ repo.description || 'No description available' }}</div>
          <div class="text-xs text-gray-500 mt-2">
            <span class="mr-4">⭐ {{ repo.stars }} Stars</span>
            <span class="mr-4">🍴 {{ repo.forks }} Forks</span>
            <span>🔗 <a :href="repo.url" target="_blank" class="text-blue-400">View Repository</a></span>
          </div>
        </li>
      </ul>
    </div>
  </template>
  
  <script setup>
  import { ref, onMounted } from 'vue';
  
  // Dummy data for repositories (in practice, this would be fetched from an API or store)
  const repositories = ref([]);
  const loading = ref(true);
  const error = ref(null);
  
  // Function to fetch repositories (this would typically make an API call)
  const fetchRepositories = async () => {
    try {
      // Simulating an API call with a timeout
      setTimeout(() => {
        // Example dummy repositories data
        repositories.value = [
          { id: 1, name: 'Repo 1', description: 'A great repo', stars: 120, forks: 45, url: 'https://github.com/repo1' },
          { id: 2, name: 'Repo 2', description: 'Another awesome repo', stars: 65, forks: 20, url: 'https://github.com/repo2' },
          { id: 3, name: 'Repo 3', description: 'Yet another repo', stars: 25, forks: 8, url: 'https://github.com/repo3' },
        ];
        loading.value = false;
      }, 1500);
    } catch (err) {
      error.value = err.message || 'Failed to fetch repositories';
      loading.value = false;
    }
  };
  
  // Fetch repositories when component is mounted
  onMounted(() => {
    fetchRepositories();
  });
  </script>
  