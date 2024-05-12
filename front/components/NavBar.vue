<template>
  <nav class="bg-gray-900 text-white">
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
      <div class="flex items-center justify-between h-16">
        <div class="flex items-center">
          <a href="/" class="flex items-center">
            <img src="https://flowbite.com/docs/images/logo.svg" class="h-8" alt="Open Bounties Logo" />
            <span class="ml-3 text-xl font-bold tracking-tight">Open Bounties</span>
          </a>
          <div class="hidden md:block">
  <div class="ml-10 flex items-baseline space-x-4">
    <NuxtLink to="/" class="px-3 py-2 rounded-md text-sm font-medium text-blue-400 bg-gray-700 hover:bg-gray-600">Home</NuxtLink>
    <NuxtLink to="/about" class="px-3 py-2 rounded-md text-sm font-medium hover:text-white hover:bg-gray-700">About</NuxtLink>
    <NuxtLink to="/dashboard" class="px-3 py-2 rounded-md text-sm font-medium hover:text-white hover:bg-gray-700">Dashboard</NuxtLink>
    <NuxtLink to="/contact" class="px-3 py-2 rounded-md text-sm font-medium hover:text-white hover:bg-gray-700">Contact</NuxtLink>
  </div>
</div>

        </div>
        <div class="hidden md:block">
          <button v-if="!$auth.loggedIn"  @click="$auth.login()" class="text-white bg-blue-600 hover:bg-blue-500 px-3 py-2 rounded-md text-sm font-medium">Login</button>
          <button v-else @click="$auth.logout()" class="text-white bg-red-600 hover:bg-red-500 px-3 py-2 rounded-md text-sm font-medium">Logout</button>
        </div>
        <div class="-mr-2 flex md:hidden">
          <button @click="isOpen = !isOpen" class="bg-gray-800 inline-flex items-center justify-center p-2 rounded-md text-gray-400 hover:text-white hover:bg-gray-700">
            <span class="sr-only">Open main menu</span>
            <!-- Icon when menu is closed. -->
            <!--
              Heroicon name: menu

              Menu open: "hidden", Menu closed: "block"
            -->
            <svg class="block h-6 w-6" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor" aria-hidden="true">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16M4 18h16"/>
            </svg>
            <!-- Icon when menu is open. -->
            <!--
              Heroicon name: x

              Menu open: "block", Menu closed: "hidden"
            -->
            <svg class="hidden h-6 w-6" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor" aria-hidden="true">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"/>
            </svg>
          </button>
        </div>
      </div>
    </div>

    <!--
      Mobile menu, toggle classes based on menu state.

      Menu open: "block", Menu closed: "hidden"
    -->
    <div class="md:hidden">
      <div class="px-2 pt-2 pb-3 space-y-1 sm:px-3">
        <a href="#" class="bg-gray-900 text-white block px-3 py-2 rounded-md text-base font-medium">Home</a>
        <a href="#" class="text-gray-300 hover:bg-gray-700 hover:text-white block px-3 py-2 rounded-md text-base font-medium">About</a>
        <a href="#" class="text-gray-300 hover:bg-gray-700 hover:text-white block px-3 py-2 rounded-md text-base font-medium">Services</a>
        <a href="#" class="text-gray-300 hover:bg-gray-700 hover:text-white block px-3 py-2 rounded-md text-base font-medium">Pricing</a>
        <a href="#" class="text-gray-300 hover:bg-gray-700 hover:text-white block px-3 py-2 rounded-md text-base font-medium">Contact</a>
      </div>
    </div>
  </nav>
  </template>



<script>
export default {
  name: 'NavBar'
  // Component setup and methods can be added here if needed
}
</script>

<style scoped>

</style>
