<template>
  <div class="flex flex-col items-center justify-center min-h-screen bg-white text-primary">
    <h1>Login with GitHub</h1>
    <button class="btn mt-4" @click="loginWithGithub">Login with GitHub</button>
    <p v-if="loading" class="text-gray-500">Logging in...</p>
  </div>
</template>

<script>
export default {
  methods: {
    loginWithGithub() {
      this.$auth.loginWith('github');
    }
  }
}
</script>
