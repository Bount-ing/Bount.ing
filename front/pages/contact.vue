<template>
  <section class="container mx-auto px-8 py-12">
    <h1 class="text-3xl font-bold text-center mb-16 text-dark">Contact Us</h1>
    <p class="text-lg text-gray-700 text-center mb-6">
      We're here to help! If you have any questions, suggestions, or need assistance with Open Bounties, please don't hesitate to reach out.
    </p>
    <div class="max-w-md mx-auto bg-white p-8 border border-gray-200 rounded-lg shadow-custom">
      <form @submit.prevent="submitForm">
        <div class="mb-16">
          <label for="name" class="block mb-2 text-sm font-medium text-gray-900">Your Name:</label>
          <input type="text" id="name" name="name" required class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-primary focus:border-primary block w-full p-2.5">
        </div>
        <div class="mb-6">
          <label for="email" class="block mb-2 text-sm font-medium text-gray-900">Your Email:</label>
          <input type="email" id="email" name="email" required class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-primary focus:border-primary block w-full p-2.5">
        </div>
        <div class="mb-16">
          <label for="message" class="block mb-2 text-sm font-medium text-gray-900">Message:</label>
          <textarea id="message" name="message" rows="4" required class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-primary focus:border-primary block w-full p-2.5"></textarea>
        </div>
        <button type="submit" class="w-full bg-primary hover:bg-dark focus:ring-4 focus:outline-none focus:ring-primary font-medium rounded-lg text-sm px-5 py-2.5 text-center text-white">
          Send Message
        </button>
      </form>
    </div>
  </section>
</template>

<script>
export default {
  name: 'ContactPage',
  methods: {
    submitForm() {
      // Placeholder for form submission logic
      alert('Your message has been sent successfully!');
    }
  }
}
</script>

<style scoped>

  /* Additional scoped styles if necessary */

</style>
