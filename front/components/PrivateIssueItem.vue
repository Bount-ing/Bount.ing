<template>
  <li class="flex flex-col md:flex-row items-start bg-secondary-dark rounded-lg shadow-lg p-4 justify-between">
    <img :src="issueImage" alt="Issue Image" class="w-20 h-20 rounded-full mr-4">
    <div class="flex-grow space-y-3">
      <h4 class="text-sm md:text-md font-bold text-primary-light">{{ issue.user_github_login }}</h4>
      <div>
        <h5 class="text-lg font-semibold text-primary-light">Private Issue</h5>
        <p class="text-info">Details are confidential.</p>
      </div>
      <div class="space-y-2">
        <p class="text-xs text-info">Repository: {{ issue.repository_name }} (⭐ {{ issue.repository_stars }})</p>
        <p class="text-xs text-info">Created: {{ formattedDate(issue.created_at) }}</p>
        <p class="text-xs text-info">Last Updated: {{ formattedDate(issue.updated_at) }}</p>
        <a :href="issueGitHubUrl" target="_blank" class="text-info hover:text-info-light">
  View Issue Details &rarr;
</a>

      </div>
      <div class="flex flex-col items-end space-y-2">
        <span class="bg-info text-white font-semibold text-md px-4 py-2 rounded-lg" v-if="bounty">
          {{ bounty.toFixed(2) }} €
        </span>
        <div class="flex space-x-2" v-if="username">
          <button @click="toggleClaimModal" class="bg-info hover:bg-info-light text-white font-bold py-2 px-4 rounded-lg">
            Claim
          </button>
          <button @click="openBountySelection" class="bg-success hover:bg-success-light text-white font-bold py-2 px-4 rounded-lg">
            Raise
          </button>
        </div>
      </div>
    </div>
  </li>
</template>
<script>
export default {
  name: 'PrivateIssueItem',
  props: {
    issue: {
      type: Object,
      required: true
    },
    bounty: Number,
    username: String
  },
  computed: {
    issueImage() {
      return this.issue.issue_image_url || 'default-image.png';
    },
    issueGitHubUrl() {
      // remove api and repos from url
      return this.issue.issue_github_url.replace('api.', '').replace('/repos', '');
    }
  },
  methods: {
    formattedDate(date) {
      return new Date(date).toLocaleDateString();
    },
    toggleClaimModal() {
      // Implementation needed here
    },
    openBountySelection() {
      // Implementation needed here
    }
  }
}
</script>
