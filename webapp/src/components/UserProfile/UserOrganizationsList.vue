<template>
    <div>
      <h2 class="text-2xl font-semibold mb-4">Organizations</h2>
      <div v-if="organizations.length === 0" class="text-center text-gray-400">
        No organizations found.
      </div>
      <ul v-else class="space-y-2">
        <!-- Display each organization -->
        <li v-for="org in organizations" :key="org.id" class="bg-gray-700 p-4 rounded-lg flex items-center">
          <img :src="org.avatar_url" alt="Organization Logo" class="w-12 h-12 rounded-full mr-4" />
          <div>
            <div class="text-lg font-medium">{{ org.name }}</div>
            <div class="text-gray-300 text-sm">{{ org.description || 'No description available' }}</div>
          </div>
        </li>
      </ul>
    </div>
  </template>
  
  <script setup>
  import { ref } from 'vue';
  
  // Sample data for organizations (replace with actual data from your API or store)
  const organizations = ref([
    {
      id: 1,
      name: 'Vue.js',
      description: 'The Progressive JavaScript Framework',
      avatar_url: 'https://vuejs.org/images/logo.png',
    },
    {
      id: 2,
      name: 'GitHub',
      description: 'Where the world builds software',
      avatar_url: 'https://github.githubassets.com/images/modules/logos_page/GitHub-Mark.png',
    },
    {
      id: 3,
      name: 'OpenAI',
      description: 'Creating Artificial General Intelligence',
      avatar_url: 'https://openai.com/favicon.ico',
    },
  ]);
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
  