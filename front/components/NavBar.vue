<template>
  <nav class="bg-white shadow-md">
    <div class="container mx-auto px-6 py-3 flex justify-between items-center">
      <a class="text-xl font-bold text-gray-800" href="/">OpenBounties</a>
      <div>
        <a class="px-4 hover:text-blue-800" href="/about">About</a>
        <a class="px-4 hover:text-blue-800" href="/contact">Contact</a>
        <button v-if="!$auth.loggedIn" @click="$auth.login()" class="bg-blue-500 text-white px-4 py-2 rounded hover:bg-blue-600">Login</button>
        <button v-else @click="$auth.logout()" class="bg-red-500 text-white px-4 py-2 rounded hover:bg-red-600">Logout</button>
      </div>
    </div>
  </nav>
</template>

<script>
export default {
  name: 'NavBar'
}
</script>

<style scoped>
/* Additional styling can be added here */
  .nav-link {
    margin-right: 1rem;
  }
</style>
