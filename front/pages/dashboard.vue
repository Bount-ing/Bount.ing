<template>
  <div class="p-6">
    <h1>GitHub Issues Dashboard</h1>
    <div v-if="issues.length">
      <ul>
        <li v-for="issue in issues" :key="issue.id">
          <h3>{{ issue.title }}</h3>
          <p>Number: {{ issue.number }}</p>
          <p>Status: {{ issue.state }}</p>
          <a :href="issue.html_url" class="text-primary underline" target="_blank">View on GitHub</a>
        </li>
      </ul>
    </div>
    <div v-else>
      <p>No issues found.</p>
    </div>
  </div>
</template>

<script>
export default {
  async mounted() {
    await this.$store.dispatch('github/fetchIssues');
  },
  computed: {
    issues() {
      return this.$store.getters['github/getIssues'];
    },
  },
};
</script>
