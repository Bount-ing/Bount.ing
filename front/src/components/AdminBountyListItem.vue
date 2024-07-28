<template>
    <div class="mb-4 p-4 rounded-lg bg-gray-700">
      <h3 class="text-xl font-semibold">{{ bounty.title }}</h3>
      <p class="text-gray-300">{{ bounty.description }}</p>
      <p class="text-gray-400">{{ $t('Amount') }}: {{ bounty.amount }}</p>
      <p class="text-gray-400">{{ $t('Status') }}: {{ bounty.status }}</p>
      <button @click="$emit('finalize-bounty', bounty.ID)" class="mt-2 p-2 bg-green-600 text-white rounded">
        {{ $t('Finalize') }}
        </button>
      <button @click="$emit('cancel-bounty', bounty.ID)" class="mt-2 p-2 bg-red-600 text-white rounded">
        {{ $t('Cancel') }}
        </button>
    </div>
  </template>

  <script>
  export default {
    props: {
      bounty: Object
    },
  };
  </script>

