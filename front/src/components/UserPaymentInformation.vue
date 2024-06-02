<template>
    <div class="bg-gray-700 p-4 rounded-lg">
      <h3 class="text-xl font-semibold mb-3">Card Information</h3>
      <p class="text-gray-300">Card Number: {{ paymentInfo.cardNumber }}</p>
      <p class="text-gray-300">Expiry Date: {{ paymentInfo.expiryDate }}</p>
      <p class="text-gray-300">Card Holder Name: {{ paymentInfo.cardHolderName }}</p>
    </div>
  </template>
  
  <script>
  export default {
    props: {
      paymentInfo: {
        type: Object,
        required: true,
      },
    },
  };
  </script>
  
  <style scoped>
  /* Add styles if needed */
  </style>
  