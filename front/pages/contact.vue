<template>
  <section class="container mx-auto px-4 py-8">
    <h1 class="text-3xl font-bold text-center mb-6">Contact Us</h1>
    <p class="text-lg text-gray-700 text-center mb-6">
      We're here to help! If you have any questions, suggestions, or need assistance with Open Bounties, please don't hesitate to reach out.
    </p>
    <div class="max-w-md mx-auto bg-white p-8 border border-gray-200 rounded-lg shadow-lg">
      <form @submit.prevent="submitForm">
        <div class="mb-6">
          <label for="name" class="block mb-2 text-sm font-medium text-gray-900">Your Name:</label>
          <input type="text" id="name" name="name" required class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5">
        </div>
        <div class="mb-6">
          <label for="email" class="block mb-2 text-sm font-medium text-gray-900">Your Email:</label>
          <input type="email" id="email" name="email" required class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5">
        </div>
        <div class="mb-6">
          <label for="message" class="block mb-2 text-sm font-medium text-gray-900">Message:</label>
          <textarea id="message" name="message" rows="4" required class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5"></textarea>
        </div>
        <button type="submit" class="w-full text-white bg-blue-500 hover:bg-blue-700 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center">
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
.container {
  background-color: #f0f9ff; /* light blue background to differentiate the contact area */
  border-radius: 16px; /* rounded corners for a softer look */
  box-shadow: 0 2px 8px rgba(0,0,0,0.1); /* subtle shadow for depth */
}
.label,
.input,
.textarea {
  display: block;
  width: 100%;
  margin-top: 0.5rem;
}
.input,
.textarea {
  font-size: 0.875rem; /* consistent font-size for input and textarea */
  padding: 0.875rem; /* comfortable padding */
  border: 1px solid #d1d5db; /* subtle border */
  border-radius: 0.375rem; /* matching border-radius */
}
.button {
  background-color: #1d4ed8; /* vibrant blue for visibility and appeal */
  color: white;
  padding: 0.625rem 1.25rem; /* balanced padding for button */
  border: none;
  border-radius: 0.375rem; /* rounded corners for the button */
  cursor: pointer; /* cursor changes to pointer to indicate clickable */
}
.button:hover {
  background-color: #2563eb; /* slightly lighter blue on hover for effect */
}
</style>
