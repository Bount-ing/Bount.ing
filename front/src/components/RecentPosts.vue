<template>
    <div class="mb-6 p-6 rounded-lg shadow-md">
      <h2 class="text-2xl font-semibold mb-4">Recent Posts</h2>
      <ul>
        <li v-for="post in recentPosts" :key="post.id" class="mb-4">
          <p class="text-gray-800">{{ post.content }}</p>
          <p class="text-gray-500 text-sm">{{ post.timestamp }}</p>
        </li>
      </ul>
    </div>
  </template>
  
  <script>
  export default {
    props: ['recentPosts']
  };
  </script>
  
  <style scoped>
  /* Add any specific styles for RecentPosts here */
  </style>
  