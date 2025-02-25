<template>
  <div class="profile flex flex-col items-center p-8">
    <img 
      :src="avatarUrl" 
      alt="Profile Picture" 
      class="w-24 h-24 rounded-full border-2 border-primary-dark" 
    />
    <h1 class="text-2xl font-semibold mt-2">{{ username }}</h1>
    <p class="text-gray-400 mt-1">{{ userBio }}</p>
    <button class="mt-3 px-3 py-1 bg-primary-dark hover:bg-primary text-white rounded-lg focus:outline-none">Edit Profile</button>
    <div class="mt-3 flex space-x-3">
      <Badge :level="userLevel" />
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue';
import { useUserStore } from '@/stores/user';

const { user } = useUserStore();
const github_user = JSON.parse(localStorage.getItem('github_user') || '{}');
console.log('Github User:', github_user);

// Computed property to handle avatar fallback
const avatarUrl = computed(() => {
  return user?.avatar || github_user?.avatar_url  || '/default-avatar.png'; // Fallback to a default avatar
});

// Computed properties for other user data with fallbacks
const userBio = computed(() => {
  return user?.bio || "User's Biography"; // Default message if bio is missing
});

const userLevel = computed(() => {
  return user?.level || 0; // Default level if missing
});

const username = computed(() => {
  return user?.username || github_user?.name || "User"; // Default username if missing
});
</script>
