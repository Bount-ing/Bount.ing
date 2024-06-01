<template>
    <div class="mb-6  p-6 rounded-lg shadow-md">
      <h2 class="text-2xl font-semibold mb-4">Personal Information</h2>
      <p class="mt-2"><strong>Full Name:</strong> {{ user.fullName }}</p>
      <p class="mt-2"><strong>Email Address:</strong> {{ user.email }}</p>
      <p class="mt-2"><strong>Phone Number:</strong> {{ user.phoneNumber }}</p>
      <p class="mt-2"><strong>Location:</strong> {{ user.location }}</p>
    </div>
  </template>
  
  <script>
  export default {
    props: ['user']
  };
  </script>
  
  <style scoped>
  /* Add any specific styles for PersonalInfo here */
  </style>
  