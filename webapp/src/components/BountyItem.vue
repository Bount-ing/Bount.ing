<template>
    <li
      class="bg-secondary p-4 rounded-lg flex items-center space-x-4 shadow-md hover:shadow-lg transition-shadow duration-300 ease-in-out"
    >
      <!-- Avatar Section -->
      <div v-if="group.avatarUrl" class="flex-shrink-0">
        <img :src="group.avatarUrl" alt="Avatar" class="rounded-full w-14 h-14 border-primary" />
      </div>
  
      <!-- Content Section -->
      <div class="flex-1">
        <!-- Title and Description -->
        <div class="mb-2">
          <h3 class="text-xl font-semibold text-primary-light">{{ group.Title }}</h3>
          <p v-if="group.Description" class="text-gray-600 mt-1 text-sm">{{ group.Description }}</p>
          <p v-else class="text-gray-500 mt-1 text-sm italic">No description provided.</p>
        </div>
  
        <!-- Issue URL -->
        <div class="mb-2">
          <a
            :href="group.issueUrl"
            target="_blank"
            rel="noopener noreferrer"
            class="text-primary hover:text-primary-dark text-sm"
          >
            View Issue
          </a>
        </div>
      </div>
  
      <!-- Right-aligned Section -->
      <div class="ml-auto flex items-center space-y-2">
        <!-- Price Button Section -->
        <div class="mr-4 flex justify-center">
          <button
            v-if="group.Bounties.length"
            class="bg-secondary text-primary-light pt-2 text-xl font-semibold transform transition-transform duration-200 hover:scale-105 rounded-md "
          >
            {{ calculateCurrentAmount(group).toFixed(2) }} €
          </button>
        </div>
  
        <!-- Raise and Claim Buttons Section -->
        <div class="flex flex-col items-center space-y-2">
          <!-- Raise Button -->
          <button
            class="bg-secondary text-primary-light py-1 px-4  text-sm font-semibold transform transition-transform duration-200 hover:scale-105 hover:bg-warning-light hover:text-secondary border-primary-light border rounded-md shadow-md"
          >
            Raise
          </button>
  
          <!-- Claim Button -->
          <button
            class="bg-secondary text-primary-light py-1 px-4 text-sm font-semibold transform transition-transform duration-200 hover:scale-105 hover:bg-success-light hover:text-secondary border-primary-light border rounded-md shadow-md"
          >
            Claim
          </button>
        </div>
      </div>
    </li>
  </template>
  
  <script setup>
  defineProps({
    group: Object, // The group object passed as a prop
    calculateCurrentAmount: Function // The calculation function passed as a prop
  })
  </script>
  