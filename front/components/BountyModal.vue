<template>
  <div v-if="isModalVisible" class="modal fixed inset-0 bg-gray-500 bg-opacity-75 flex justify-center items-center">
    <div class="modal-content bg-white p-4 rounded-lg shadow-lg">
      <h3 class="text-lg font-semibold mb-4">Set Bounty Amount</h3>
      <input type="number" v-model="amount" placeholder="Enter amount in EUR" class="input border border-gray-300 p-2 rounded">
      <div class="flex space-x-4 mt-4" v-if="username">
        <button @click="submit" class="btn bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded">Submit</button>
        <button @click="close" class="btn bg-red-500 hover:bg-red-700 text-white font-bold py-2 px-4 rounded">Cancel</button>
      </div>
    </div>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  name: 'BountyModal',
  props: {
    isModalVisible: {
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
      this.isModalVisible = false;
      this.$emit('update:isModalVisible', false);
    },
    submit() {
      this.submitBounty();
    },
    validateDates() {
      return new Date(this.bountyEnd) > new Date(this.bountyStart);
    },
    async submitBounty() {
      try {
        const response = await axios.post('http://0.0.0.0:8080/api/v1/bounties/', {
          amount: parseFloat(this.amount),
          currency: this.bountyCurrency,
          issue_github_id: this.issue.id,
          issue_github_url: this.issue.url,
          user_github_login: this.username,
          created_at: new Date().toISOString(),
          start_at: new Date().toISOString(),
          end_at: new Date().toISOString(),
        });
        this.show = false;
        this.$emit('submit-bounty', this.amount);
        this.close();
      } catch (error) {
        console.error('Error submitting bounty:', error);
      }
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
