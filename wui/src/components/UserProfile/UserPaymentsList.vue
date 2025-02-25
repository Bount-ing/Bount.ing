<template>
    <div>
      <h2 class="text-2xl font-semibold mb-4">Payments</h2>
      <!-- Payment History Table -->
      <table class="min-w-full table-auto bg-secondary rounded-lg">
        <thead class="text-left text-gray-300">
          <tr>
            <th class="px-4 py-2">Date</th>
            <th class="px-4 py-2">Amount</th>
            <th class="px-4 py-2">Method</th>
            <th class="px-4 py-2">Status</th>
          </tr>
        </thead>
        <tbody>
          <!-- Loop through each payment and display it in a row -->
          <tr v-for="payment in payments" :key="payment.id" class="hover:bg-gray-700">
            <td class="px-4 py-2">{{ payment.date }}</td>
            <td class="px-4 py-2">${{ payment.amount.toFixed(2) }}</td>
            <td class="px-4 py-2">{{ payment.method }}</td>
            <td class="px-4 py-2">
              <span :class="payment.statusClass">{{ payment.status }}</span>
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  </template>
  <script setup>
  import { ref } from 'vue';
  
  // Mock data for payments
  const payments = ref([
  ]);
  
  </script>
  