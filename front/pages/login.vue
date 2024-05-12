<template>
  <div class="flex flex-col items-center justify-center min-h-screen bg-light text-primary">
    <h1 class="font-serif text-xl text-secondary">Login with GitHub</h1>
    <button class="px-8 py-4 bg-primary text-light rounded-lg shadow-custom mt-4 transition duration-300 ease-in-out hover:bg-accent focus:outline-none focus:ring-2 focus:ring-primary focus:ring-opacity-50" @click="loginWithGithub">
      Login with GitHub
    </button>
    <p v-if="loading" class="text-info">Logging in...</p>
  </div>
</template>

<script>
export default {
  data() {
    return {
      loading: false
    };
  },
  methods: {
    loginWithGithub() {
      this.loading = true;
      this.$auth.loginWith('github').then((response) => {
        this.$store.commit('github/setToken', response.data.access_token);
        this.loading = false;
       });      setTimeout(() => {
        this.loading = false;
        alert('Logged in successfully!');
      }, 2000);
    }
  }
}
</script>
