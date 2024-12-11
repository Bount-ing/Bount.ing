<template>
  <div>
    <h2 class="text-2xl font-semibold mb-4">Organizations</h2>
    <div v-if="organizations.length === 0" class="text-center text-gray-400">
      No organizations found.
    </div>
    <ul v-else class="space-y-2">
      <!-- Display each organization -->
      <li 
        v-for="org in organizations" 
        :key="org.id" 
        class="bg-secondary p-4 rounded-lg flex items-center"
      >
        <img 
          :src="org.avatar_url" 
          alt="Organization Logo" 
          class="w-12 h-12 rounded-full mr-4" 
        />
        <div>
          <div class="text-lg font-medium">{{ org.login }}</div>
          <div class="text-gray-300 text-sm">{{ org.description || 'No description available' }}</div>
        </div>
      </li>
    </ul>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { useUserStore } from '@/stores/user';

// Create a reactive variable to hold organizations
const organizations = ref([]);

// Access the user store
const userStore = useUserStore();

// Fetch organizations on component mount
onMounted(() => {
  // Assume userStore.user.orgs is either pre-fetched or fetched on demand
  organizations.value = userStore.orgs || [];
  console.log('Organizations:', organizations.value);
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
