<template>
  <nav class="bg-gray-800 text-white shadow-lg">
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
      <div class="flex items-center justify-between h-16">
        <!-- Logo and navigation links -->
        <div class="flex items-center">
          <a href="/" class="flex items-center">
            <img src="https://flowbite.com/docs/images/logo.svg" class="h-8" alt="Open Bounties Logo" />
            <span class="ml-3 text-xl font-bold tracking-tight">Bount.ing</span>
          </a>
          <div class="hidden md:flex space-x-4 ml-10">
            <NuxtLink to="/" class="px-3 py-2 rounded-md text-sm font-medium bg-gray-900 hover:bg-gray-700">Home</NuxtLink>
            <NuxtLink to="/dashboard" class="px-3 py-2 rounded-md text-sm font-medium hover:bg-gray-700">Dashboard</NuxtLink>
            <NuxtLink to="/about" class="px-3 py-2 rounded-md text-sm font-medium hover:bg-gray-700">About</NuxtLink>
            <NuxtLink to="/contact" class="px-3 py-2 rounded-md text-sm font-medium hover:bg-gray-700">Contact</NuxtLink>
          </div>
        </div>

        <!-- Authentication buttons -->
        <div class="hidden md:block">
          <button v-if="!$auth.loggedIn" @click="$auth.login()" class="px-3 py-2 rounded-md text-sm font-medium bg-blue-500 hover:bg-blue-600">Login</button>
          <button v-else @click="$auth.logout()" class="px-3 py-2 rounded-md text-sm font-medium bg-red-600 hover:bg-red-700">Logout</button>
        </div>

        <!-- Mobile menu button -->
        <div class="md:hidden bg-gray-700 p-2 rounded-lg">
          <button @click="isOpen = !isOpen" class="inline-flex items-center justify-center p-2 rounded-md text-gray-400 hover:text-white hover:bg-gray-600">
            <span class="sr-only">Open main menu</span>
            <svg v-show="!isOpen" class="block h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor" aria-hidden="true">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16m-16 6h16"/>
            </svg>
            <svg v-show="isOpen" class="block h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor" aria-hidden="true">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"/>
            </svg>
          </button>
        </div>
      </div>
    </div>

    <!-- Mobile Menu -->
    <div v-show="isOpen" class="md:hidden">
      <div class="px-2 pt-2 pb-3 space-y-1 sm:px-3">
        <NuxtLink to="/" class="block px-3 py-2 rounded-md text-base font-medium bg-gray-900 hover:bg-gray-700">Home</NuxtLink>
        <NuxtLink to="/dashboard" class="block px-3 py-2 rounded-md text-base font-medium hover:bg-gray-700">Dashboard</NuxtLink>
        <NuxtLink to="/about" class="block px-3 py-2 rounded-md text-base font-medium hover:bg-gray-700">About</NuxtLink>
        <NuxtLink to="/contact" class="block px-3 py-2 rounded-md text-base font-medium hover:bg-gray-700">Contact</NuxtLink>
      </div>
    </div>
  </nav>
</template>

<script>
export default {
  data() {
    return {
      isOpen: false,
    };
  },
  name: 'NavBar',
  methods: {
    loginWithGithub() {
    console.log("Attempting to log in with GitHub");
    this.$auth.loginWith('github').then(() => {
      console.log("Logged in successfully!");
    }).catch(err => {
      console.error("Error logging in:", err);
    });
  }
  }
}
</script>

<style scoped>

</style>
