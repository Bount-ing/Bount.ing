<template>
  <div v-if="isClaimModalVisible" class="fixed inset-0 bg-gray-500 bg-opacity-75 flex items-center justify-center">
    <div class="bg-white p-6 rounded-lg shadow-lg max-w-sm mx-auto">
      <h3 class="text-lg font-semibold mb-4">Claim Bounty</h3>
      <p class="text-sm text-gray-600 mb-4">Are you sure you want to claim this bounty?</p>
      <p class="text-sm mb-6">
        Until further notice to claim a bounty you must write to claim@issues.ninja, specifying the issue and PR IDs from GitHub.
        The email must be sent from the same email as your public GitHub email.
      </p>
      <div class="flex justify-between space-x-4" v-if="username">
        <button @click="submit" class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded transition duration-150 ease-in-out">
          Claim
        </button>
        <button @click="close" class="bg-red-500 hover:bg-red-700 text-white font-bold py-2 px-4 rounded transition duration-150 ease-in-out">
          Cancel
        </button>
      </div>
    </div>
  </div>
</template>


<script>
import axios from 'axios';

export default {
  name: 'ClaimModal',
  props: {
    isClaimModalVisible: {
      type: Boolean,
      required: true
    },
    issue: {
      type: Object,
      required: true
    },
    username: {
      type: String,
      required: true
    }
  },
  data() {
    return {
      amount: 0.0,
      bountyEnd: '',
      bountyStart: '',
      bountyCurrency: 'EUR',
      currentIssue: null,
    };
  },
  methods: {
    close() {
      this.isClaimModalVisible = false;
      this.$emit('update:isClaimModalVisible', false);
    },
    submit() {
      this.submitClaim();
    },
  }
};
</script>

<style scoped>
.modal {
  display: flex;
  align-items: center;
  justify-content: center;
}
</style>
