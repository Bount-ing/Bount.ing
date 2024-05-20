<template>
  <li class="flex flex-col md:flex-row items-start bg-gray-100 rounded-lg shadow-lg p-4 justify-between">
    <img :src="issue.image_url || 'default-image.png'" alt="Repo Image" class="w-20 h-20 rounded-full mr-4">
    <div class="flex-grow space-y-3">
      <div class="flex items-center space-x-4">
        <h4 class="text-sm md:text-md font-bold text-blue-500">{{ issue.user_github_login }}</h4>
      </div>
      <div>
        <h5 class="text-lg font-semibold text-blue-500">{{ issue.title }}</h5>
        <p class="text-gray-700">{{ issue.description }}</p>
      </div>
      <div class="space-y-2">
        <p class="text-xs text-gray-500">Repository: {{ issue.repository_name }} (⭐ {{ issue.repository_stars }})</p>
        <p class="text-xs text-gray-500">Created: {{ new Date(issue.created_at).toLocaleDateString() }}</p>
        <p class="text-xs text-gray-500">Last Updated: {{ new Date(issue.updated_at).toLocaleDateString() }}</p>
        <a :href="`https://github.com/issues/${issue.github_id}`" target="_blank" class="text-blue-500 hover:text-blue-600">
          View Issue on GitHub &rarr;
        </a>
      </div>
    </div>
    <div class="flex flex-col items-end space-y-2">
      <span class="bg-blue-500 text-white font-semibold text-md px-4 py-2 rounded-lg">
        <div v-if="issue.bounty" class="">
          {{ bounty.toFixed(2) }} €
        </div>
      </span>
      <div class="flex space-x-2" v-if="username">
        <button @click="$emit('claim', issue)" class="bg-blue-300 hover:bg-blue-400 text-white font-bold py-2 px-4 rounded-lg">
          Claim
        </button>
        <button @click="toggleBountyModal" class="bg-green-500 hover:bg-green-600 text-white font-bold py-2 px-4 rounded-lg">
          Raise
        </button>
        <BountyModal v-if="showBountyModal" :issue="issue" :show="showBountyModal" :username="username"/>
      </div>
    </div>
  </li>
</template>

<script>
export default {
  name: 'IssueItem',
  data() {
    return {
      showBountyModal: false,
    }
  },
  methods: {
    toggleBountyModal() {
  this.showBountyModal = !this.showBountyModal;
  console.log('Bounty Modal toggled:', this.showBountyModal);  // Debugging output
},

  },
  props: {
    issue: {
      type: Object,
      required: true
    },
    username: {
      type: String,
      required: true
    },
    bounty: {
      type: Object,
      required: false
    }
  },
  components: {
    BountyModal: () => import('@/components/BountyModal.vue')
  }
}
</script>

<style scoped>
.rounded-lg {
  border-radius: 0.5rem;
}
</style>
