<template>
    <li class="flex flex-col md:flex-row items-start bg-secondary-dark rounded-lg shadow-lg p-4 justify-between">
      <div class="flex items-start">
        <img :src="issue.image_url || issueImage" alt="Repo Image" class="w-20 h-20 rounded-full mr-4">
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
  
  <script lang="ts">
  import { defineComponent, computed } from 'vue';
  
  export default defineComponent({
    name: 'PrivateIssueListItem',
    props: {
      issue: {
        type: Object as () => { issue_image_url?: string; issue_github_url: string; image_url?: string },
        required: true
      },
      bounty: Number,
      username: String
    },
    setup(props) {
      const issueImage = computed(() => props.issue.issue_image_url || 'default-image.png');
  
      const issueGitHubUrl = computed(() => {
        return props.issue.issue_github_url.replace('api.', '').replace('/repos', '');
      });
  
      const repoOwner = computed(() => {
        const urlParts = issueGitHubUrl.value.split('/');
        return urlParts[3];
      });
  
      const repoName = computed(() => {
        const urlParts = issueGitHubUrl.value.split('/');
        return urlParts[4];
      });
  
      const formattedDate = (date: string) => {
        return new Date(date).toLocaleDateString();
      };
  
      const toggleClaimModal = () => {
        // Implementation needed here
      };
  
      const openBountySelection = () => {
        // Implementation needed here
      };
  
      return {
        issueImage,
        issueGitHubUrl,
        repoOwner,
        repoName,
        formattedDate,
        toggleClaimModal,
        openBountySelection
      };
    }
  });
  </script>
  