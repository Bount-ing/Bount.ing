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
    { id: 1, date: '2024-12-01', amount: 50.0, method: 'Credit Card', status: 'Completed', statusClass: 'text-green-500' },
    { id: 2, date: '2024-11-25', amount: 120.0, method: 'PayPal', status: 'Pending', statusClass: 'text-yellow-500' },
    { id: 3, date: '2024-11-10', amount: 75.5, method: 'Bank Transfer', status: 'Failed', statusClass: 'text-red-500' },
    { id: 4, date: '2024-10-30', amount: 200.0, method: 'Credit Card', status: 'Completed', statusClass: 'text-green-500' },
  ]);
  
  </script>
  