<template>
  <div class="p-6 text-center">
    <h1>GitHub Callback</h1>
    <div v-if="$route.query.error">
      <p>Error: {{ $route.query.error_description }}</p>
    </div>
    <div v-else>
      <p>Authentication successful! Redirecting...</p>
      <script>setTimeout(() => this.$router.push('/dashboard'), 2000);</script>
    </div>
  </div>


</template>

<script>
export default {
  async asyncData({ app, route }) {
    if (route.query.code) {
      await app.$auth.handleRedirectCallback().then(() => {
        params: { code: route.query.code }
      });
    }
  }
}
</script>
