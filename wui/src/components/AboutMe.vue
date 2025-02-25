<template>
    <div class="mb-6 p-6 rounded-lg shadow-md">
      <h2 class="text-2xl font-semibold mb-4">{{ $t('About Me') }}</h2>
      <p class="mt-2">{{ aboutMe }}</p>
    </div>
  </template>
  
  <script>
  export default {
    props: ['aboutMe']
  };
  </script>
  
  <style scoped>
  /* Add any specific styles for AboutMe here */
  </style>
  