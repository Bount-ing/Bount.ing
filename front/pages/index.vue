<template>

  <section class="container mx-auto px-4 sm:px-6 lg:px-8 py-6 bg-white rounded-lg shadow">
    <NavBar />
    <div v-if="$auth.loggedIn" class="bg-white shadow-md rounded-lg p-4 md:p-6 transition duration-300 ease-in-out hover:shadow-lg">
      <h1 class="text-3xl md:text-4xl text-gray-900 font-semibold mb-4">Welcome back, {{ $auth.user.name }}</h1>
      <p class="text-sm md:text-base text-gray-600">You are logged in as {{ $auth.user.username }}.</p>
      <p class="text-sm md:text-base text-gray-600">Your email is {{ $auth.user.email }}.</p>
    </div>
    <div v-else class="text-center">
      <h1 class="text-3xl md:text-4xl text-gray-900 font-semibold mb-4">Welcome to OpenBounties</h1>
      <p class="text-sm md:text-base text-gray-600">Please login to see your profile information.</p>
      <button class="mt-4 px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-700 focus:ring-opacity-50" @click="$auth.login()">Login with GitHub</button>
    </div>
    <Footer />
  </section>
</template>

<script>
export default {
  name: 'HomePage',
  // You might want to add some script logic here if needed
}
</script>
</script>

<style>
/* Add any global styles or module-specific styles here */
</style>
