<template>
  <div v-if="isModalVisible" class="modal fixed inset-0 bg-gray-800 bg-opacity-50 flex justify-center items-center">
    <div class="modal-content bg-white p-6 rounded-lg shadow-lg max-w-md w-full">
      <h3 class="text-xl font-semibold mb-6 text-gray-800">Create a Single Bounty</h3>
      <label class="block mb-4">
        <span class="text-gray-700">Bounty Amount (EUR):</span>
        <input type="number" v-model="individualAmount" placeholder="Enter amount" min="1" class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-300 focus:ring focus:ring-indigo-200 focus:ring-opacity-50">
      </label>
      <label class="block mb-4">
        <span class="text-gray-700">Start Date:</span>
        <input type="date" v-model="bountyStart" class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-300 focus:ring focus:ring-indigo-200 focus:ring-opacity-50">
      </label>
      <label class="block mb-4">
        <span class="text-gray-700">End Date:</span>
        <input type="date" v-model="bountyEnd" class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-300 focus:ring focus:ring-indigo-200 focus:ring-opacity-50">
      </label>
      <div class="flex justify-between mt-6">
        <button @click="submit" class="btn bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-6 rounded-lg focus:outline-none focus:shadow-outline">Submit</button>
        <button @click="close" class="btn bg-red-500 hover:bg-red-700 text-white font-bold py-2 px-6 rounded-lg focus:outline-none focus:shadow-outline">Cancel</button>
      </div>
    </div>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  name: 'SingleBountyModal',
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
      individualAmount: 0.0,
      bountyStart: '',
      bountyEnd: '',
      bountyCurrency: 'EUR',
      authToken: '',
    };
  },
  methods: {
    close() {
      this.isModalVisible = false;
      this.$emit('update:isModalVisible', false);
    },
    async submit() {
      if (this.validateDates()) {
        await this.loginAndSetToken();
        await this.submitBounty(this.individualAmount, this.bountyStart, this.bountyEnd);
        this.close();
      } else {
        alert('Please ensure the end date is after the start date and all inputs are valid.');
      }
    },
    validateDates() {
      return new Date(this.bountyEnd) > new Date(this.bountyStart) && this.individualAmount > 0;
    },
    async loginAndSetToken() {
      try {
        const baseURL = process.env.API_BASE_URL;
        const githubToken = process.env.GITHUB_TOKEN;

        const response = await axios.post(`${baseURL}/api/v1/login`, {
          github_token: githubToken,
        });

        this.authToken = response.data.jwt;
      } catch (error) {
        console.error('Error logging in:', error);
        alert('Failed to login. Please check your network and try again.');
      }
    },
    async submitBounty(amount, start_at, end_at) {
      try {
        const baseURL = process.env.API_BASE_URL;

        const response = await axios.post(`${baseURL}/api/v1/bounties/`, {
          amount: parseFloat(amount),
          currency: this.bountyCurrency,
          auth_token: this.authToken,
          issue_github_id: this.issue.id,
          issue_github_url: this.issue.url,
          issue_image_url: this.issue.image_url,
          user_github_login: this.username,
          start_at: new Date(start_at).toISOString(),
          end_at: new Date(end_at).toISOString(),
        });
        console.log('Bounty submitted:', response);
      } catch (error) {
        console.error('Error submitting bounty:', error);
        alert('Failed to submit bounty. Please check your network and try again.');
      }
    }
  }
};
</script>
