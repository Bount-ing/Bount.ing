<template>
  <section class="container mx-auto px-4 py-12">
    <!-- Banner Section -->
    <div class="text-center mb-12">
      <h1 class="text-5xl font-bold mb-4">Pricing for Bount.ing</h1>
      <p class="text-xl">
        Contribute to open source projects, earn bounties, and enhance your skills. Fair and simple pricing for everyone.
      </p>
    </div>
    
    <!-- Pricing Section -->
    <div class="flex flex-col md:flex-row justify-center items-center gap-8 mb-16">
      <!-- Public Repositories Pricing Card -->
      <div class="shadow-lg rounded-lg overflow-hidden w-full max-w-md">
        <div class="bg-success p-6 text-center text-white">
          <h2 class="text-3xl font-semibold mb-2">Public Repositories</h2>
          <p class="text-xl">4.2% Service Fee</p>
        </div>
        <div class="p-6 text-center">
          <p class="text-lg mb-6">
            Engage with the community, fix issues, and get rewarded with a minimal fee.
          </p>
          <a href="#" class="bg-success text-white py-2 px-6 rounded-lg text-lg hover:bg-success-light">
            Get Started
          </a>
        </div>
      </div>

      <!-- Private Repositories Pricing Card -->
      <div class="shadow-lg rounded-lg overflow-hidden w-full max-w-md">
        <div class="bg-info p-6 text-center text-white">
          <h2 class="text-3xl font-semibold mb-2">Private Repositories</h2>
          <p class="text-xl">10% Service Fee</p>
        </div>
        <div class="p-6 text-center">
          <p class="text-info text-lg mb-6">
            Keep your projects private while still benefiting from our bounty platform.
          </p>
          <a href="#" class="bg-info text-white py-2 px-6 rounded-lg text-lg hover:bg-info-light">
            Get Started
          </a>
        </div>
      </div>
    </div>

    <!-- Additional Information -->
    <div class="text-center">
      <h2 class="text-4xl font-semibold text-gray-800 mb-4">Why Choose Bount.ing?</h2>
      <p class="text-xl text-gray-600 mb-8">
        By using our platform, you support the open-source community, sharpen your coding skills, and earn rewards for your contributions.
      </p>
      <a href="https://discord.gg/NVV2ww22ek" target="_blank" class="bg-purple-600 text-white py-2 px-6 rounded-lg text-lg hover:bg-purple-700">
        Join Our Discord Community
      </a>
    </div>
  </section>
</template>

<script lang="ts">
import { defineComponent } from 'vue';

export default defineComponent({
  name: 'PricingView'
});
</script>

<style scoped>
/* Add any specific styles here */
</style>
