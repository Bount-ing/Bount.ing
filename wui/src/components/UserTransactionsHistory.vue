<template>
    <div>
      <div v-if="transactions.length">
        <table class="min-w-full bg-gray-700">
          <thead>
            <tr>
              <th class="py-2">{{ $t('Type') }}</th>
              <th class="py-2">{{ $t('Amount') }}</th>
              <th class="py-2">{{ $t('Date') }}</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="transaction in transactions" :key="transaction.id" class="bg-gray-600">
              <td class="py-2 px-4">{{ transaction.type }}</td>
              <td class="py-2 px-4">{{ transaction.amount }}</td>
              <td class="py-2 px-4">{{ transaction.date }}</td>
            </tr>
          </tbody>
        </table>
      </div>
      <div v-else>
        <p class="text-gray-400">{{ $t('No transactions found.') }}</p>
      </div>
    </div>
  </template>
  
  <script>
  export default {
    props: {
      transactions: {
        type: Array,
        required: true,
      },
    },
  };
  </script>
  
  <style scoped>
  /* Add styles if needed */
  </style>
  