<template>
  <section class="container mx-auto px-4 sm:px-6 lg:px-8 py-6 bg-gradient-to-r from-primary to-accent rounded-lg shadow-lg">
    <div v-if="$auth.loggedIn" class="bg-gradient-to-br from-leather-brown to-desert-orange shadow-md rounded-lg p-4 md:p-6 transition duration-300 ease-in-out hover:shadow-xl space-y-4">
        <h1 class="text-3xl md:text-4xl text-dark font-bold mb-4">Welcome back, {{ $auth.user.name }}</h1>
       <p class="text-sm md:text-base text-treasure-gold">You are logged in as {{ $store.state.github.token ? 'a GitHub user' : 'a guest' }}.</p>
       <p class="text-sm md:text-base text-treasure-gold">Your email address is {{ $store.state.github.email || 'unknown' }}.</p>
    </div>
    <div v-else class="bg-white shadow-md rounded-lg p-4 md:p-6 transition duration-300 ease-in-out hover:shadow-xl space-y-4">
      <h1 class="text-3xl md:text-4xl text-midnight-black font-bold mb-4">Welcome to Open Bounties</h1>
      <p class="text-sm md:text-base text-secondary">Please log in to access your dashboard.</p>
       <button @click="$auth.login()" class="bg-primary hover:bg-secondary text-light font-bold py-2 px-4 rounded transition duration-300 ease-in-out">Login with GitHub</button>
    </div>
  </section>
</template>

<script>
export default {
  name: 'Home',
  methods: {
    loginWithGithub() {
      console.log("Attempting to log in with GitHub");
      this.$auth.loginWith('github').then(() => {
        console.log("Logged in with GitHub");
      });
    }
  }
}
</script>
