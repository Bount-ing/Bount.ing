<template>
  <div v-if="isModalVisible" class="modal fixed top-0 left-0 inset-0 bg-black h-screen w-screen flex justify-center items-center">
    <div class="modal-content p-6 rounded-lg shadow-lg max-w-md w-full  border-primary border">
      <h3 class="text-xl font-semibold mb-6 text-gray-800">Create Multiple Degressive Bounties</h3>
      <form @submit.prevent="submitAllBounties">
        <label class="block mb-4">
          <span class="text-gray-700">Total Bounty Amount (EUR):</span>
          <input
            type="number"
            v-model.number="totalAmount"
            placeholder="Enter total amount"
            min="1"
            class="mt-1 block w-full rounded-md shadow-sm focus:border-indigo-300 focus:ring focus:ring-indigo-200 focus:ring-opacity-50 bg-black border-primary text-primary"
          />
        </label>
        <label class="block mb-4">
          <span class="text-gray-700">Step Frequency (Days):</span>
          <input
            type="number"
            v-model.number="stepFrequency"
            placeholder="Frequency in days"
            min="1"
            class="mt-1 block w-full rounded-md shadow-sm focus:border-indigo-300 focus:ring focus:ring-indigo-200 focus:ring-opacity-50 bg-black border-primary text-primary"          />
        </label>
        <label class="block mb-4">
          <span class="text-gray-700">Start Date:</span>
          <input
            type="date"
            v-model="bountyStart"
            class="mt-1 block w-full rounded-md shadow-sm focus:border-indigo-300 focus:ring focus:ring-indigo-200 focus:ring-opacity-50 bg-black border-primary text-primary"          />
        </label>
        <label class="block mb-4">
          <span class="text-gray-700">End Date:</span>
          <input
            type="date"
            v-model="bountyEnd"
            class="mt-1 block w-full rounded-md shadow-sm focus:border-indigo-300 focus:ring focus:ring-indigo-200 focus:ring-opacity-50 bg-black border-primary text-primary"          />
        </label>
        <div v-if="isValidInput" class="mt-4 mb-6">
          <p class="text-gray-700">Number of Bounties to Create: {{ numberOfBounties }}</p>
          <p class="text-gray-700">Value per Bounty: {{ bountyValue.toFixed(2) }} EUR</p>
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
            class="btn border  border-error text-error font-bold py-2 px-6 rounded-lg focus:outline-none focus:shadow-outline"
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
  name: 'DegressiveBountyModal',
  props: {
    isModalVisible: Boolean,
    issue: Object,
    username: String
  },
  data() {
    const today = new Date().toISOString().substring(0, 10);
    return {
      totalAmount: 0,
      stepFrequency: 1,
      bountyStart: today,
      bountyEnd: today
    };
  },
  computed: {
    isValidInput() {
      return this.totalAmount > 0 && this.stepFrequency > 0 && this.bountyStart && this.bountyEnd;
    },
    numberOfBounties() {
      return Math.floor(this.durationDays / this.stepFrequency) + 1;
    },
    durationDays() {
      return (new Date(this.bountyEnd) - new Date(this.bountyStart)) / (1000 * 60 * 60 * 24);
    },
    bountyValue() {
      return this.totalAmount / this.numberOfBounties;
    }
  },
  methods: {
    close() {
      this.$emit('close-modal');
    },
    async submitBounty(start_at, end_at, amount, authHeader) {
      const baseURL = import.meta.env.VITE_API_BASE_URL;
      try {
        const response = await axios.post(
          `${baseURL}/api/v1/bounties/`,
          {
            amount: parseFloat(amount),
            currency: 'EUR',
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
        console.log('Bounty submitted:', response.data);
      } catch (error) {
        console.error('Error submitting bounty:', error);
        throw new Error('Failed to submit bounty. Please check your network and try again.');
      }
    },
    async submitAllBounties() {
      if (this.isValidInput) {
        const userStore = useUserStore();
        if (!userStore.isLoggedIn) {
          await userStore.login();
        }
        const authHeader = userStore.authHeader;
        const startDate = new Date(this.bountyStart);
        let currentDate = new Date(this.bountyStart);
        const endDate = new Date(this.bountyEnd);

        if (startDate >= endDate) {
          alert('The start date must be before the end date.');
          return;
        }

        if (isNaN(startDate) || isNaN(endDate)) {
          alert('Invalid date inputs. Please check the dates.');
          return;
        }

        try {
          while (currentDate <= endDate) {
            await this.submitBounty(
              startDate.toISOString(),
              currentDate.toISOString(),
              this.bountyValue.toFixed(2),
              authHeader
            );
            currentDate.setDate(currentDate.getDate() + this.stepFrequency);
          }
          this.close();
        } catch (error) {
          alert('Failed to submit multiple bounties. Please check your network connection and try again.');
        }
      } else {
        alert('Please ensure all inputs are valid.');
      }
    }
  }
};
</script>

