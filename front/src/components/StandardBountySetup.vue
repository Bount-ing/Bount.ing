<template>
  <div v-if="isModalVisible" class="modal fixed top-0 left-0 inset-0 bg-black h-screen w-screen flex justify-center items-center">
    <div class="modal-content p-6 rounded-lg shadow-lg max-w-md w-full border-primary border">
      <h3 class="text-xl font-semibold mb-6 text-gray-800">Create a Single Bounty</h3>
      <form @submit.prevent="handleSubmit">
        <div class="form-group">
          <label class="block mb-4">
            <span class="text-gray-700">Bounty Amount (EUR):</span>
            <input
              type="number"
              v-model.number="individualAmount"
              placeholder="Enter amount"
              min="1"
              class="mt-1 block w-full rounded-md shadow-sm focus:border-indigo-300 focus:ring focus:ring-indigo-200 focus:ring-opacity-50 bg-black border-primary text-primary"
            />
          </label>
        </div>
        <div class="form-group">
          <label class="block mb-4">
          <span class="text-gray-700">Start Date:</span>
          <input type="date" v-model="bountyStart" class="mt-1 block w-full rounded-md shadow-sm focus:border-indigo-300 focus:ring focus:ring-indigo-200 focus:ring-opacity-50 border-primary text-primary bg-grey-500" >
        </label>
        <label class="block mb-4">
          <span class="text-gray-700">End Date:</span>
          <input type="date" v-model="bountyEnd" class="mt-1 block w-full rounded-md shadow-sm focus:border-indigo-300 focus:ring focus:ring-indigo-200 focus:ring-opacity-50 border-primary text-primary bg-grey-500">
        </label>
        </div>
        <div class="flex justify-between mt-6">
          <button
            type="submit"
            class="btn border border-success text-success font-bold py-2 px-6 rounded-lg focus:outline-none focus:shadow-outline"
          >
            Submit
          </button>
          <button
            type="button"
            @click="close"
            class="btn border border-error text-error font-bold py-2 px-6 rounded-lg focus:outline-none focus:shadow-outline"
          >
            Cancel
          </button>
        </div>
      </form>
    </div>
  </div>
</template>

<script>
import axios from 'axios';
import { useUserStore } from '@/stores/user';


export default {
  name: 'StandardBountySetup',
  components: {

  },
  props: {
    isModalVisible: Boolean,
    issue: Object,
    username: String
  },
  data() {

    const today = new Date().toISOString().substring(0, 10);
    const oneyearfromnow = new Date(new Date().setFullYear(new Date().getFullYear() + 1)).toISOString().substring(0, 10);
    return {
      individualAmount: 0.0,
      bountyCurrency: 'EUR',
      authToken: '',
      bountyStart: today,
      bountyEnd: oneyearfromnow,
      bountyType: 'flat'
    };
  },
  methods: {
    close() {
      this.$emit('close-modal');
    },
    async handleSubmit() {
      if (this.validateDates() && this.individualAmount > 0) {
        try {
          const userStore = useUserStore();
          if (!userStore.isLoggedIn) {
            await userStore.login();
          }
          await this.submitBounty(this.individualAmount, this.bountyStart, this.bountyEnd, userStore.authHeader);
          this.close();
        } catch (error) {
          console.error(error);
          alert('Submission failed: ' + error.message);
        }
      } else {
        alert('Please ensure the end date is after the start date and all inputs are valid.');
      }
    },
    validateDates() {
      return new Date(this.bountyEnd) >= new Date(this.bountyStart);
    },
    async submitBounty(amount, start_at, end_at, authHeader) {
      const baseURL = import.meta.env.VITE_API_BASE_URL;
      try {
        const response = await axios.post(
          `${baseURL}/api/v1/bounties/`,
          {
            amount: parseFloat(amount),
            currency: this.bountyCurrency,
            bounty_type: this.bountyType,
            issue_github_id: this.issue.id,
            issue_github_url: this.issue.url,
            issue_image_url: this.issue.image_url,
            user_github_login: this.username,
            start_at: new Date(start_at).toISOString(),
            end_at: new Date(end_at).toISOString()
          },
          {
            headers: {
              Authorization: authHeader
            }
          }
        );
        console.log('Bounty submitted:', response);
      } catch (error) {
        console.error('Error submitting bounty:', error);
        throw new Error('Failed to submit bounty. Please check your network and try again.');
      }
    }
  }
};
</script>
