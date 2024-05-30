<template>
    <section class="container mx-auto px-8 py-14 text-primary">
      <h1 class="text-3xl font-bold text-center mb-16 text-primary">Contact Us</h1>
      <p class="text-lg text-center mb-8">
        We're here to help! If you have any questions, suggestions, or need assistance with Open Bounties, please don't hesitate to reach out.
      </p>
      <div class="max-w-md mx-auto p-8 border border-primary rounded-lg shadow-custom">
        <form @submit.prevent="submitForm">
          <div class="mb-16">
            <label for="name" class="block mb-2 text-sm font-medium">Your Name:</label>
            <input
              v-model="formData.name"
              type="text"
              id="name"
              name="name"
              required
              class="bg-black border border-primary text-gray-900 text-sm rounded-lg focus:ring-primary focus:border-primary block w-full p-2.5"
            />
          </div>
          <div class="mb-6">
            <label for="email" class="block mb-2 text-sm font-medium">Your Email:</label>
            <input
              v-model="formData.email"
              type="email"
              id="email"
              name="email"
              required
              class="bg-black border border-primary text-gray-900 text-sm rounded-lg focus:ring-primary focus:border-primary block w-full p-2.5"
            />
          </div>
          <div class="mb-16">
            <label for="message" class="block mb-2 text-sm font-medium">Message:</label>
            <textarea
              v-model="formData.message"
              id="message"
              name="message"
              rows="4"
              required
              class="bg-black border border-primary text-gray-900 text-sm rounded-lg focus:ring-primary focus:border-primary block w-full p-2.5"
            ></textarea>
          </div>
          <button
            type="submit"
            class="w-full border border-primary hover:bg-primary focus:ring-4 focus:outline-none focus:ring-primary font-medium rounded-lg text-sm px-5 py-2.5 text-center text-white"
          >
            Send Message
          </button>
        </form>
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
  /* Additional scoped styles if necessary */
  </style>
  