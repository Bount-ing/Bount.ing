<template>
    <div class="mb-6  p-6 rounded-lg shadow-md">
      <h2 class="text-2xl font-semibold mb-4">Interests</h2>
      <div class="mt-2 flex flex-wrap">
        <span v-for="interest in interests" :key="interest" class="bg-blue-100 text-blue-800 text-sm font-semibold mr-2 px-2.5 py-0.5 rounded">{{ interest }}</span>
      </div>
    </div>
  </template>
  
  <script>
  export default {
    props: ['interests']
  };
  </script>
  
  <style scoped>
  /* Add any specific styles for Interests here */
  </style>
  