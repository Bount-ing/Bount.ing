<template>
    <div>
      <div v-if="bounties.length">
        <div v-for="bounty in bounties" :key="bounty.id" class="mb-4 p-4 rounded-lg bg-gray-700">
          <h3 class="text-xl font-semibold">{{ bounty.title }}</h3>
          <p class="text-gray-300">{{ bounty.description }}</p>
          <p class="text-gray-400">{{ $t('Amount') }}: {{ bounty.amount }}</p>
          <p class="text-gray-400">{{ $t('Status') }}: {{ bounty.status }}</p>
        </div>
      </div>
      <div v-else>
        <p class="text-gray-400">{{ $t('No bounties found.') }}</p>
      </div>
    </div>
  </template>
  
  <script>
  export default {
    props: {
      bounties: {
        type: Array,
        required: true,
      },
    },
  };
  </script>
  
  <style scoped>
  /* Add styles if needed */
  </style>
  