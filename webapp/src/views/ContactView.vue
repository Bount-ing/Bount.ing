<template>
  <section class="container mx-auto px-8 py-14 text-primary">
    <h1 class="text-3xl font-bold text-center mb-16 text-primary">Contact Us</h1>
    <p class="text-lg text-center mb-8">
      We're here to help! If you have any questions, suggestions, or need assistance with Open Bounties, please don't
      hesitate to reach out.
    </p>
    <div class="max-w-md mx-auto p-8 border border-primary rounded-lg shadow-custom bg-secondary">
      <form @submit.prevent="submitForm">
        <div class="mb-6">
          <label for="name" class="block mb-2 text-sm font-medium">Your Name:</label>
          <input v-model="formData.name" type="text" id="name" name="name" required
            class="bg-secondary-light border border-primary text-gray-900 text-sm rounded-lg focus:ring-primary focus:border-primary block w-full p-2.5" />
        </div>
        <div class="mb-6">
          <label for="email" class="block mb-2 text-sm font-medium">Your Email:</label>
          <input v-model="formData.email" type="email" id="email" name="email" required
            class="bg-secondary-light border border-primary text-gray-900 text-sm rounded-lg focus:ring-primary focus:border-primary block w-full p-2.5" />
        </div>
        <div class="mb-16">
          <label for="message" class="block mb-2 text-sm font-medium">Message:</label>
          <textarea v-model="formData.message" id="message" name="message" rows="4" required
            class="bg-secondary-light border border-primary text-gray-900 text-sm rounded-lg focus:ring-primary focus:border-primary block w-full p-2.5"></textarea>
        </div>
        <button type="submit"
          class="w-full border bg-primary border-primary hover:bg-primary-dark focus:ring-4 focus:outline-none focus:ring-primary font-medium rounded-lg text-sm px-5 py-2.5 text-center text-white">
          Send Message
        </button>
      </form>
    </div>

    <!-- Contact Links Section -->
    <div class="mt-12 text-center">
      <h2 class="text-2xl font-semibold mb-4">Other Ways to Reach Us</h2>
      <div class="flex justify-center space-x-6">
        <!-- LinkedIn -->
        <a href="https://www.linkedin.com/company/bount-ing" target="_blank"
          class="text-primary hover:text-primary-light transition">
          <span class="inline-flex items-center space-x-2">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="currentColor" viewBox="0 0 448 512"><path d="M416 32H31.9C14.3 32 0 46.5 0 64.3v383.4C0 465.5 14.3 480 31.9 480H416c17.6 0 32-14.5 32-32.3V64.3c0-17.8-14.4-32.3-32-32.3zM135.4 416H69V202.2h66.5V416zm-33.2-243c-21.3 0-38.5-17.3-38.5-38.5S80.9 96 102.2 96c21.2 0 38.5 17.3 38.5 38.5 0 21.3-17.2 38.5-38.5 38.5zm282.1 243h-66.4V312c0-24.8-.5-56.7-34.5-56.7-34.6 0-39.9 27-39.9 54.9V416h-66.4V202.2h63.7v29.2h.9c8.9-16.8 30.6-34.5 62.9-34.5 67.2 0 79.7 44.3 79.7 101.9V416z"/></svg>
            <span class="font-bold text-base">LinkedIn</span>
          </span>
        </a>

        <!-- Discord -->
        <a href="https://discord.gg/NVV2ww22ek" target="_blank"          
        class="text-primary hover:text-primary-light transition">
          <span class="inline-flex items-center space-x-2 ">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="currentColor" viewBox="0 0 24 24">
              <path
                d="M20.317 4.369a19.791 19.791 0 0 0-4.885-1.621.074.074 0 0 0-.078.037 13.875 13.875 0 0 0-.601 1.233 19.918 19.918 0 0 0-5.506 0 13.875 13.875 0 0 0-.601-1.233.077.077 0 0 0-.078-.037c-1.666.37-3.344.919-4.885 1.621a.069.069 0 0 0-.032.027C.533 9.053.093 13.594.385 18.074a.074.074 0 0 0 .031.051 19.875 19.875 0 0 0 5.989 3.05.073.073 0 0 0 .083-.026 13.564 13.564 0 0 0 1.107-1.789.074.074 0 0 0-.041-.107 12.35 12.35 0 0 1-1.746-.841.073.073 0 0 1-.007-.12c.116-.088.232-.18.344-.271a.075.075 0 0 1 .079-.012c3.684 1.672 7.658 1.672 11.317 0a.075.075 0 0 1 .081.01c.112.092.229.184.346.272a.073.073 0 0 1-.007.12 12.351 12.351 0 0 1-1.747.841.073.073 0 0 0-.04.107c.361.637.753 1.239 1.107 1.789a.073.073 0 0 0 .083.026 19.86 19.86 0 0 0 5.988-3.05.072.072 0 0 0 .031-.051c.307-4.503-.092-9.014-3.884-13.678a.07.07 0 0 0-.033-.027zm-12.33 9.139c-1.052 0-1.912-.967-1.912-2.156 0-1.19.847-2.157 1.912-2.157 1.061 0 1.918.967 1.912 2.157 0 1.19-.847 2.156-1.912 2.156zm7.027 0c-1.052 0-1.912-.967-1.912-2.156 0-1.19.847-2.157 1.912-2.157 1.061 0 1.918.967 1.912 2.157 0 1.19-.851 2.156-1.912 2.156z" />
            </svg>
            <span class="font-bold text-base">Discord</span>
          </span>
        </a>

        <!-- GitHub -->
        <a href="https://github.com/Bount-ing/Bount.ing/issues/new" target="_blank"          
        class="text-primary hover:text-primary-light transition">
          <span class="inline-flex items-center space-x-2">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="currentColor" viewBox="0 0 496 512">
              <path d="M165.9 397.4c0 2-2.3 3.6-5.2 3.6-3.3 .3-5.6-1.3-5.6-3.6 0-2 2.3-3.6 5.2-3.6 3-.3 5.6 1.3 5.6 3.6zm-31.1-4.5c-.7 2 1.3 4.3 4.3 4.9 2.6 1 5.6 0 6.2-2s-1.3-4.3-4.3-5.2c-2.6-.7-5.5 .3-6.2 2.3zm44.2-1.7c-2.9 .7-4.9 2.6-4.6 4.9 .3 2 2.9 3.3 5.9 2.6 2.9-.7 4.9-2.6 4.6-4.6-.3-1.9-3-3.2-5.9-2.9zM244.8 8C106.1 8 0 113.3 0 252c0 110.9 69.8 205.8 169.5 239.2 12.8 2.3 17.3-5.6 17.3-12.1 0-6.2-.3-40.4-.3-61.4 0 0-70 15-84.7-29.8 0 0-11.4-29.1-27.8-36.6 0 0-22.9-15.7 1.6-15.4 0 0 24.9 2 38.6 25.8 21.9 38.6 58.6 27.5 72.9 20.9 2.3-16 8.8-27.1 16-33.7-55.9-6.2-112.3-14.3-112.3-110.5 0-27.5 7.6-41.3 23.6-58.9-2.6-6.5-11.1-33.3 2.6-67.9 20.9-6.5 69 27 69 27 20-5.6 41.5-8.5 62.8-8.5s42.8 2.9 62.8 8.5c0 0 48.1-33.6 69-27 13.7 34.7 5.2 61.4 2.6 67.9 16 17.7 25.8 31.5 25.8 58.9 0 96.5-58.9 104.2-114.8 110.5 9.2 7.9 17 22.9 17 46.4 0 33.7-.3 75.4-.3 83.6 0 6.5 4.6 14.4 17.3 12.1C428.2 457.8 496 362.9 496 252 496 113.3 383.5 8 244.8 8zM97.2 352.9c-1.3 1-1 3.3 .7 5.2 1.6 1.6 3.9 2.3 5.2 1 1.3-1 1-3.3-.7-5.2-1.6-1.6-3.9-2.3-5.2-1zm-10.8-8.1c-.7 1.3 .3 2.9 2.3 3.9 1.6 1 3.6 .7 4.3-.7 .7-1.3-.3-2.9-2.3-3.9-2-.6-3.6-.3-4.3 .7zm32.4 35.6c-1.6 1.3-1 4.3 1.3 6.2 2.3 2.3 5.2 2.6 6.5 1 1.3-1.3 .7-4.3-1.3-6.2-2.2-2.3-5.2-2.6-6.5-1zm-11.4-14.7c-1.6 1-1.6 3.6 0 5.9 1.6 2.3 4.3 3.3 5.6 2.3 1.6-1.3 1.6-3.9 0-6.2-1.4-2.3-4-3.3-5.6-2z"/>
            </svg>
            <span class="font-bold text-base">Github</span>
          </span>
        </a>

      </div>
    </div>
  </section>
</template>
<script lang="ts">
import { defineComponent, reactive } from 'vue';

export default defineComponent({
  name: 'ContactView',
  setup() {
    const formData = reactive({
      name: '',
      email: '',
      message: '',
    });

    const submitForm = () => {
      // Placeholder for form submission logic
      alert('Your message has been sent successfully!');
    };

    return {
      formData,
      submitForm,
    };
  },
});
</script>

<style scoped>

</style>