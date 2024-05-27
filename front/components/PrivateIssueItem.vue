<template>
  <li class="flex flex-col md:flex-row items-start bg-secondary-dark rounded-lg shadow-lg p-4 justify-between">
    <div class="flex items-start">
      <img :src="issue.image_url || issue.issue_image_url || 'default-image.png'" alt="Repo Image" class="w-20 h-20 rounded-full mr-4">
      <div class="flex flex-col">
        <span class="text-sm md:text-md font-bold text-primary">{{ repoOwner }} / {{ repoName }}</span>
        <span class="text-lg md:text-md font-bold text-primary">Private Issue</span>
        <p class="text-info">Details are confidential.</p>
        <p class="text-info">You need to be logged in and granted to see it.</p>
      </div>
    </div>
    <div class="flex flex-col justify-center items-center md:items-end space-y-2 mt-4 md:mt-0">
      <span class="text-primary text-mono font-semibold text-md px-4 py-2 rounded-lg" v-if="bounty">
        {{ bounty.toFixed(2) }} €
      </span>
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
    },
    repoOwner() {
      const urlParts = this.issueGitHubUrl.split('/');
      return urlParts[3];
    },
    repoName() {
      const urlParts = this.issueGitHubUrl.split('/');
      return urlParts[4];
    },
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
