<template>
    <div class="p-4">
      <h2 class="text-2xl font-semibold mb-4">Badges</h2>
      <!-- Display badges if available -->
      <div v-if="badges.length">
        <ul class="space-y-4">
          <li v-for="badge in badges" :key="badge.id" class="flex items-center space-x-3 bg-gray-700 p-4 rounded-md">
            <!-- Badge Icon -->
            <img :src="badge.icon" alt="badge icon" class="w-8 h-8 rounded-full" />
            <!-- Badge Name and Description -->
            <div>
              <h3 class="text-lg font-semibold">{{ badge.name }}</h3>
              <p class="text-sm text-gray-400">{{ badge.description }}</p>
            </div>
          </li>
        </ul>
      </div>
      <!-- Show message if no badges are available -->
      <div v-else>
        <p class="text-gray-400">No badges available</p>
      </div>
    </div>
  </template>
  
  <script setup>
  import { ref } from 'vue';
  
  // Sample badge data (replace with real data from API or store)
  const badges = ref([
    {
      id: 1,
      name: 'Contributor',
      description: 'Awarded for contributing to open source.',
      icon: 'https://via.placeholder.com/150', // Example icon, replace with actual icon URL
    },
    {
      id: 2,
      name: 'Top Contributor',
      description: 'Awarded to top contributors based on activity.',
      icon: 'https://via.placeholder.com/150', // Example icon, replace with actual icon URL
    },
    {
      id: 3,
      name: 'Verified',
      description: 'Badge for verified users.',
      icon: 'https://via.placeholder.com/150', // Example icon, replace with actual icon URL
    },
  ]);
  
  </script>
  
  <style scoped>
  /* Custom styles for badge list (optional) */
  </style>
  