<template>
  <section class="container mx-auto px-4 py-12">
    <!-- Banner Section -->
    <div class="text-center mb-12">
      <h1 class="text-5xl font-bold mb-4">Service Fees</h1>
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
          <router-link to="/signup" class="bg-success text-white py-2 px-6 rounded-lg text-lg hover:bg-success-light">
            Get Started
          </router-link>
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
          <router-link to="/signup" class="bg-info text-white py-2 px-6 rounded-lg text-lg hover:bg-info-light">
            Get Started
          </router-link>
        </div>
      </div>
    </div>

    <!-- Additional Information -->
    <div class="text-center">
      <h2 class="text-4xl font-semibold text-secondary-light mb-4">Why Choose Bount.ing?</h2>
      <p class="text-xl text-gray-600 mb-8">
        By using our platform, you support the open-source community, sharpen your coding skills, and earn rewards for your contributions.
      </p>
      <h2 class="text-3xl font-semibold text-secondary-light mb-4">Any doubts ? Ask on Discord !</h2>
      <div class="flex justify-center mx-4">
              <a href="https://discord.gg/NVV2ww22ek" target="_blank"          
        class="text-primary hover:text-primary-light transition border rounded-2xl px-2 border-primary hover:border-primary-light">
          <span class="inline-flex items-center space-x-2 text-xl text-primary hover:text-primary-light transition  rounded-2xl py-2 px-2 font-bold">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="currentColor" viewBox="0 0 24 24">
              <path
                d="M20.317 4.369a19.791 19.791 0 0 0-4.885-1.621.074.074 0 0 0-.078.037 13.875 13.875 0 0 0-.601 1.233 19.918 19.918 0 0 0-5.506 0 13.875 13.875 0 0 0-.601-1.233.077.077 0 0 0-.078-.037c-1.666.37-3.344.919-4.885 1.621a.069.069 0 0 0-.032.027C.533 9.053.093 13.594.385 18.074a.074.074 0 0 0 .031.051 19.875 19.875 0 0 0 5.989 3.05.073.073 0 0 0 .083-.026 13.564 13.564 0 0 0 1.107-1.789.074.074 0 0 0-.041-.107 12.35 12.35 0 0 1-1.746-.841.073.073 0 0 1-.007-.12c.116-.088.232-.18.344-.271a.075.075 0 0 1 .079-.012c3.684 1.672 7.658 1.672 11.317 0a.075.075 0 0 1 .081.01c.112.092.229.184.346.272a.073.073 0 0 1-.007.12 12.351 12.351 0 0 1-1.747.841.073.073 0 0 0-.04.107c.361.637.753 1.239 1.107 1.789a.073.073 0 0 0 .083.026 19.86 19.86 0 0 0 5.988-3.05.072.072 0 0 0 .031-.051c.307-4.503-.092-9.014-3.884-13.678a.07.07 0 0 0-.033-.027zm-12.33 9.139c-1.052 0-1.912-.967-1.912-2.156 0-1.19.847-2.157 1.912-2.157 1.061 0 1.918.967 1.912 2.157 0 1.19-.847 2.156-1.912 2.156zm7.027 0c-1.052 0-1.912-.967-1.912-2.156 0-1.19.847-2.157 1.912-2.157 1.061 0 1.918.967 1.912 2.157 0 1.19-.851 2.156-1.912 2.156z" />
            </svg>
            <span class="font-bold ">Join us on Discord !</span>
          </span>
        </a>
    </div>
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
