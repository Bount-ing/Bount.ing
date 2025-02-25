<template>
  <div class="repository-container">
    <h2 class="text-2xl font-semibold mb-4">Repositories</h2>

    <!-- Display message if no repositories are found -->
    <div v-if="repositories.length === 0" class="text-center text-gray-400">
      No repositories found.
    </div>

    <!-- Display list of repositories -->
    <ul v-else class="space-y-4">
      <li 
        v-for="repo in repositories" 
        :key="repo.id" 
        class="bg-secondary p-4 rounded-lg flex items-start space-x-4"
      >
        <!-- Repository sponsor avatar -->
        <img 
          :src="repo.sponsor.avatar_url" 
          alt="Organization Logo" 
          class="w-16 h-16 rounded-full"
        />

        <!-- Repository details -->
        <div class="flex-1">
          <div class="text-lg font-medium">
            <a 
              :href="repo.html_url" 
              target="_blank" 
              class="text-blue-400 hover:underline"
            >
              {{ repo.name }}
            </a>
          </div>
          <p class="text-gray-300 text-sm mb-2">
            {{ repo.description || 'No description available' }}
          </p>

          <!-- Repository language -->
          <div class="text-sm text-gray-500">
            <span v-if="repo.language">{{ repo.language }}</span>
          </div>
        </div>
      </li>
    </ul>
  </div>
</template>


<script setup>
import { ref, onMounted } from 'vue';
import { useUserStore } from '@/stores/user';

// Create a reactive variable to hold repositories
const repositories = ref([]);

// Access the user store
const userStore = useUserStore();

// Fetch repositories on component mount
onMounted(() => {
  // Assume userStore.user.repos is either pre-fetched or fetched on demand
  repositories.value = userStore.repos || [];
  console.log('Repositories:', repositories.value);
});
</script>

<style scoped>
/* Additional styling for the list items */
li {
  transition: background-color 0.2s ease;
}

li:hover {
  background-color: #2d3748;
}
</style>
