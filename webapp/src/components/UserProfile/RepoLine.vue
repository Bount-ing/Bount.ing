<template>
    <li class="bg-gray-700 p-4 rounded-lg hover:bg-gray-600 cursor-pointer">
      <div class="text-lg font-semibold">{{ repo.name || 'Unnamed Repository' }}</div>
      <div class="text-sm text-gray-400">
        {{ repo.description || 'No description available' }}
      </div>
      <div class="text-xs text-gray-500 mt-2">
        <span class="mr-4">⭐ {{ 0 || 0 }} Stars</span>
        <span class="mr-4">🍴 {{ 0 || 0 }} Forks</span>
        <span>
          🔗 
          <a 
            :href="repo.url || '#'"
            target="_blank" 
            class="text-blue-400">
            View Repository
          </a>
        </span>
      </div>
    </li>
  </template>
  
  <script setup>
  defineProps({
    repo: {
      type: Object,
      required: true,
    },
  });
  </script>
  
  <style scoped>
  /* Add styles specific to RepoLine if needed */
  </style>
  