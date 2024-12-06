<template>
  <section class="bg-secondary-dark container mx-auto px-4 my-6 py-6 rounded-lg text-primary">
    <div class="text-center">
      <h1 class="text-4xl font-bold text-center mb-6">Join Bount.ing</h1>
      <div class="flex justify-center mb-6">
        <img src="/bount.ing.banner.png" alt="Bount.ing Banner" class="w-1/2 object-cover rounded-lg" />
      </div>
      <p class="text-lg text-center mb-10">
        Revolutionizing open source development. Contribute, learn, and earn with every code commit while advancing sustainability.
      </p>
    </div>
    <h2 class="text-2xl font-semibold mb-4">How It Works</h2>
    <ul class="list-disc pl-5 mb-8">
      <li>Connect your GitHub account to track and sponsor issues.</li>
      <li>Pick challenges in your favorite projects and set bounties.</li>
      <li>Fix issues to earn bounties and sharpen your skills.</li>
    </ul>
    <h2 class="text-2xl font-semibold mb-4">How to Claim a Bounty</h2>
    <p class="text-lg mb-4">
      Link your pull request to the related GitHub issue. After issue closure, we verify contributions and credit your account automatically. Unclaimed bounties transform into future opportunities or refunds.
    </p>
    <h2 class="text-2xl font-semibold mb-4">Our Mission</h2>
    <p class="text-lg mb-8">
      Founded on May 4, 2024, in Barcelona, Bount.ing is committed to fostering independent work, embracing open-source values, promoting learning while earning, and advancing sustainability in technology.
    </p>
    <h2 class="text-2xl font-semibold mb-4">Join Our Community</h2>
    <p class="mb-8">
      Connect with fellow developers and open-source enthusiasts. Share, learn, and grow together in our vibrant Discord community.
    </p>
    <div class="flex flex-row items-center justify-center mb-8 p-4">
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
    <div class="flex justify-center mx-4">
      <router-link to="/signup" class="text-2xl text-primary hover:text-primary-light transition border rounded-2xl py-2 px-3 border-primary hover:border-primary-light font-bold">Sign Up</router-link>
    </div>
    </div>
  </section>
</template>

<script lang="ts">
import { defineComponent } from 'vue';

export default defineComponent({
  name: 'AboutPage'
});
</script>

<style scoped>
/* Add any specific styles here */
</style>
