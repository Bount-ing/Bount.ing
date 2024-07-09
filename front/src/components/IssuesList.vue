<template>
  <div class="mb-6 bg-white p-6 rounded-lg shadow-md">
    <h2 class="text-2xl font-semibold mb-4">{{ $t('Issues') }}</h2>
    <ul>
      <li v-for="issue in issues" :key="issue.id" class="mb-4">
        <p class="text-gray-800">{{ issue.title }}</p>
        <p class="text-gray-500 text-sm">{{ issue.body }}</p>
      </li>
    </ul>
  </div>
</template>

<script>
export default {
  props: ['issues']
};
</script>

<style scoped>
/* Add any specific styles for IssuesList here */
</style>
