<template>
  <li class="flex flex-col md:flex-row items-start bg-secondary-dark rounded-lg shadow-lg p-4 justify-between">
    <img :src="issue.image_url || issue.issue_image_url || 'default-image.png'" alt="Repo Image" class="w-20 h-20 rounded-full mr-4">
    <div class="flex-grow space-y-3">
      <div class="flex flex-col">
        <span class="text-sm md:text-md font-bold text-primary">{{ repoOwner }} / {{ repoName }}</span>
        <span class="text-lg md:text-md font-bold text-primary">{{ issue.title }}</span>
        <p class="text-info">{{ issue.description }}</p>
        <a :href="issueGitHubUrl" target="_blank" class="text-info hover:text-info-light">
          {{ $t('View Issue on GitHub') }} &rarr;
        </a>
      </div>
      <div></div>
    </div>
    <div class="flex flex-col items-end space-y-2">
      <span class="text-primary text-mono font-semibold text-md px-4 py-2 rounded-lg" v-if="bounty">
        {{ bounty.toFixed(2) }} €
      </span>
      <div class="flex space-x-2" v-if="username">
        <button @click="toggleClaimModal" class="hover:bg-info text-primary hover:text-secondary-dark font-bold py-2 px-4 rounded-lg">
          {{ $t('Claim') }}
        </button>
        <button @click="openBountySelection" class="hover:bg-success hover:text-secondary-dark text-primary font-bold py-2 px-4 rounded-lg">
          {{ $t('Raise') }}
        </button>
      </div>
      <div v-if="showBountyTypeSelection" class="modal fixed top-0 left-0 inset-0 bg-black h-screen w-screen flex justify-center items-center ">
        <div class="p-4 rounded-lg space-y-4 border border-primary">
          <h3>{{ $t('Select Bounty Type') }}:</h3>
          <button @click="selectBountyType('progressive')" class="w-full border text-success border-success hover:border-success-light font-bold py-2 px-4 rounded-lg">
            {{ $t('Economic') }}
          </button>
          <button @click="selectBountyType('degressive')" class="w-full border text-warning border-warning hover:border-warning-light font-bold py-2 px-4 rounded-lg">
            {{ $t('Fast') }}
          </button>
          <button @click="selectBountyType('single')" class="w-full border text-info border-info hover:border-info-lightfont-bold py-2 px-4 rounded-lg">
            {{ $t('Standard') }}
          </button>
        </div>
      </div>
      <component :is="bountyTypeComponent" v-if="isModalVisible" 
        :isModalVisible="isModalVisible"
        @close-modal="isModalVisible = false"
        :issue="issue"
        :username="username"
      />
    </div>
  </li>
</template>

<script>
import ProgressiveBountySetup from '@/components/ProgressiveBountySetup.vue';
import DegressiveBountySetup from '@/components/DegressiveBountySetup.vue';
import StandardBountySetup from '@/components/StandardBountySetup.vue';

export default {
  name: 'IssueListItem',
  components: {
    ProgressiveBountySetup,
    DegressiveBountySetup,
    StandardBountySetup,
    
  },
  props: {
    issue: Object,
    username: String,
    bounty: Number
  },
  data() {
    return {
      isModalVisible: false,
      isClaimModalVisible: false,
      showBountyTypeSelection: false,
      bountyType: null
    };
  },
  methods: {
    selectBountyType(type) {
      this.bountyType = type;
      this.isModalVisible = true; // Ensure modal visibility is set to true
      this.showBountyTypeSelection = false; // Ensure selection overlay is hidden
    },
    toggleClaimModal() {
      this.isClaimModalVisible = !this.isClaimModalVisible;
    },
    openBountySelection() {
      this.showBountyTypeSelection = true;
      this.isModalVisible = false; // Ensure other modals are closed
    }
  },
  computed: {
    bountyTypeComponent() {
      switch (this.bountyType) {
        case 'progressive': return 'ProgressiveBountySetup';
        case 'degressive': return 'DegressiveBountySetup';
        case 'single': return 'StandardBountySetup';
        default: return null; // Ensuring a null is returned if no match
      }
    },
    issueGitHubUrl() {
      // remove api and repos from url
      try {
        return this.issue.url.replace('api.', '').replace('/repos', '');
      } catch (e) {
        return this.issue.issue_github_url.replace('api.', '').replace('/repos', '');
      }
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
}
</script>
