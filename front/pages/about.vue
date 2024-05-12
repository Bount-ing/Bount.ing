<template>
  <section class="container mx-auto px-4 py-6">
    <h1 class="text-3xl font-bold text-center mb-4">Open Bounties</h1>
    <p class="text-lg text-gray-700 text-center mb-6">
      Open Bounties is a revolutionary platform where open source enthusiasts can sponsor development by putting bounties on GitHub issues. Enhance the progress of your favorite projects by financially supporting specific development tasks.
    </p>
    <div class="flex justify-center">
      <img src="https://via.placeholder.com/350x200.png?text=Open+Bounties" alt="Open Bounties Image" class="shadow-lg rounded-lg mb-6">
    </div>
    <h2 class="text-2xl font-semibold mb-3">How It Works</h2>
    <ul class="list-disc pl-5 mb-6">
      <li>Connect your GitHub account to get started.</li>
      <li>Choose issues you want to sponsor in your favorite repositories.</li>
      <li>Set a bounty and watch the community tackle it with enthusiasm.</li>
    </ul>
    <h2 class="text-2xl font-semibold mb-3">Join Us</h2>
    <p class="mb-6">Become a part of the Open Bounties community today and help accelerate the development of open source projects. Your contribution makes a difference!</p>
    <div class="flex justify-center">
      <button class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded">
        Sign Up Now
      </button>
    </div>
  </section>
</template>

<script>
export default {
  name: 'AboutPage'
}
</script>

<style scoped>
.container {
  background-color: #f9fafb; /* light gray background */
  border: 1px solid #e2e8f0; /* subtle border */
  border-radius: 12px; /* rounded corners */
  padding: 20px; /* padding around content */
}
.text-lg {
  line-height: 1.6; /* improved line spacing for readability */
}
.list-disc {
  margin-top: 10px; /* space above list */
}
.bg-blue-500 {
  transition: background-color 0.3s; /* smooth transition for button hover */
}
</style>
