<template>
    <div class="container mx-auto px-4 py-12">
      <h1 class="text-4xl font-bold text-center mb-6 text-gray-800">Login with GitHub</h1>
      <div class="flex justify-center">
        <button
          @click="loginWithGitHub"
          class="bg-gray-800 text-white font-bold py-2 px-4 rounded shadow-lg hover:bg-gray-700"
        >
          Login with GitHub
        </button>
      </div>
    </div>
  </template>
  
  <script lang="ts">
  import { defineComponent } from 'vue';
  import axios from 'axios';
  
  export default defineComponent({
    setup() {
      const loginWithGitHub = async () => {
        try {
          const clientId = 'YOUR_GITHUB_CLIENT_ID';
          const redirectUri = 'YOUR_REDIRECT_URI';
          const scope = 'read:user repo';
          const authUrl = `https://github.com/login/oauth/authorize?client_id=${clientId}&redirect_uri=${redirectUri}&scope=${scope}`;
  
          window.location.href = authUrl;
        } catch (error) {
          console.error('GitHub login failed:', error);
        }
      };
  
      return {
        loginWithGitHub,
      };
    },
  });
  </script>
  
  <style scoped>
  .cursor-pointer {
    cursor: pointer;
  }
  </style>
  