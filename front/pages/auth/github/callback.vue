<template>
  <div class="p-6 text-center">
    <h1>GitHub Callback</h1>
    <p v-if="$route.query.error">
      Error: {{ $route.query.error_description }}
    </p>
    <p v-else>
      Authentication successful! Redirecting...
    </p>
  </div>
</template>

<script>
export default {
  async asyncData({ app, route }) {

  }
}
</script>
