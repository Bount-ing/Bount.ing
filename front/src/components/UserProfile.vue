<template>
    <div class="profile flex flex-col items-center p-8 bg-gray-900">
      <img :src="user.avatar" alt="Profile Picture" class="w-24 h-24 rounded-full border-2 border-gray-500" />
      <h1 class="text-2xl font-semibold mt-2">{{ user.username }}</h1>
      <p class="text-gray-400 mt-1">{{ user.userBio }}</p>
      <button class="mt-3 px-3 py-1 bg-blue-600 text-white rounded-lg focus:outline-none">Edit Profile</button>
      <div class="mt-3 flex space-x-3">
        <Badge :level="user.level" />
        <Achievement :achievements="user.achievements" />
      </div>
    </div>
  </template>
  
  <script>
  export default {
    props: {
      user: Object
    }
  };
  </script>
  