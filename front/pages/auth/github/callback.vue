<template>
  <div class="p-6 text-center">
    <h1>GitHub Callback</h1>
    <div v-if="$route.query.error">
      <p>Error: {{ $route.query.error_description }}</p>
    </div>
    <div v-else>
      <p>Authentication successful! Redirecting...</p>
      <script>setTimeout(() => this.$router.push('/'), 2000);</script>
    </div>
  </div>


</template>

<script>
export default {
}
</script>
