<template>
  <li class="flex flex-col md:flex-row items-start bg-secondary-dark rounded-lg shadow-lg p-4 justify-between">
    <img :src="issue.image_url || issue.issue_image_url || 'default-image.png'" alt="Repo Image" class="w-20 h-20 rounded-full mr-4">
    <div class="flex-grow space-y-3">
      <div class="flex items-center space-x-4">
        <h4 class="text-sm md:text-md font-bold text-primary-light">{{ issue.user_github_login }}</h4>
      </div>
      <div>
        <h5 class="text-lg font-semibold text-primary-light">{{ issue.title }}</h5>
        <p class="text-info">{{ issue.description }}</p>
      </div>
      <div class="space-y-2">
        <p class="text-xs text-info">Repository: {{ issue.repository_name }} (⭐ {{ issue.repository_stars }})</p>
        <p class="text-xs text-info">Created: {{ new Date(issue.created_at).toLocaleDateString() }}</p>
        <p class="text-xs text-info">Last Updated: {{ new Date(issue.updated_at).toLocaleDateString() }}</p>
        <a :href="issueGitHubUrl" target="_blank" class="text-info hover:text-info-light">
          View Issue on GitHub &rarr;
        </a>
      </div>
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
      <div v-if="showBountyTypeSelection" class="absolute top-0 left-0 right-0 bottom-0 bg-black bg-opacity-50 flex justify-center items-center">
        <div class="bg-white p-4 rounded-lg space-y-4">
          <h3>Select Bounty Type:</h3>
          <button @click="selectBountyType('progressive')" class="w-full bg-green-500 hover:bg-green-600 text-white font-bold py-2 px-4 rounded-lg">
            Economic
          </button>
          <button @click="selectBountyType('degressive')" class="w-full bg-orange-500 hover:bg-orange-600 text-white font-bold py-2 px-4 rounded-lg">
            Fast
          </button>
          <button @click="selectBountyType('single')" class="w-full bg-blue-500 hover:bg-blue-600 text-white font-bold py-2 px-4 rounded-lg">
            Standard
          </button>
        </div>
      </div>
      <!-- Dynamically inject the correct modal component -->
      <component :is="bountyTypeComponent" v-if="isModalVisible" :isModalVisible="isModalVisible" :issue="issue" :username="username"/>
      <ClaimModal v-if="isClaimModalVisible" :issue="issue" :username="username"/>
    </div>
  </li>
</template>

<script>
import ProgressiveBountyModal from '@/components/ProgressiveBountyModal.vue';
import DegressiveBountyModal from '@/components/DegressiveBountyModal.vue';
import SingleBountyModal from '@/components/SingleBountyModal.vue';
import ClaimModal from '@/components/ClaimModal.vue';

export default {
  name: 'IssueItem',
  components: {
    ProgressiveBountyModal,
    DegressiveBountyModal,
    SingleBountyModal,
    ClaimModal
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
        case 'progressive': return 'ProgressiveBountyModal';
        case 'degressive': return 'DegressiveBountyModal';
        case 'single': return 'SingleBountyModal';
        default: return null; // Ensuring a null is returned if no match
      }
    },
    issueGitHubUrl() {
      // remove api and repos from url
      try{
        return this.issue.url.replace('api.', '').replace('/repos', '');
      }catch(e){
        return this.issue.issue_github_url.replace('api.', '').replace('/repos', '');
      }
    }
  }
}
</script>
