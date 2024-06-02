<template>
  <div v-if="isModalVisible" class="fixed inset-0 bg-black bg-opacity-50 flex justify-center items-center">
    <div class=" p-6 rounded-lg shadow-lg max-w-md w-full border border-primary">
      <h3 class="text-xl font-semibold mb-6 text-gray-800">Create a Progressive Bounty</h3>
      <form @submit.prevent="submit">
        <label class="block mb-4">
          <span class="text-gray-700">Bounty Amount (EUR):</span>
          <input type="number" v-model.number="individualAmount" placeholder="Enter amount" min="1" class="mt-1 block w-full rounded-md shadow-sm focus:border-indigo-300 focus:ring focus:ring-indigo-200 focus:ring-opacity-50 border-primary text-primary bg-black">
        </label>
        <label class="block mb-4">
          <span class="text-gray-700">Frequency (Days):</span>
          <input type="number" v-model.number="stepFrequency" placeholder="Enter frequency in days" min="1" class="mt-1 block w-full rounded-md shadow-sm focus:border-indigo-300 focus:ring focus:ring-indigo-200 focus:ring-opacity-50 border-primary text-primary bg-black">
        </label>
        <label class="block mb-4">
          <span class="text-gray-700">Start Date:</span>
          <input type="date" v-model="bountyStart" class="mt-1 block w-full rounded-md shadow-sm focus:border-indigo-300 focus:ring focus:ring-indigo-200 focus:ring-opacity-50 border-primary text-primary bg-grey-500" >
        </label>
        <label class="block mb-4">
          <span class="text-gray-700">End Date:</span>
          <input type="date" v-model="bountyEnd" class="mt-1 block w-full rounded-md shadow-sm focus:border-indigo-300 focus:ring focus:ring-indigo-200 focus:ring-opacity-50 border-primary text-primary bg-grey-500">
        </label>
        <div class="stats mt-4">
          <p class="text-gray-800">Total Bounties Created: <strong>{{ formattedTotalBounties }}</strong></p>
          <p class="text-gray-800">Total Amount: <strong>{{ formattedTotalAmount }} EUR</strong></p>
        </div>
        <div class="flex justify-between mt-6">
          <button type="submit" class="border border-success text-success font-bold py-2 px-6 rounded-lg focus:outline-none focus:shadow-outline">
            Submit
          </button>
          <button type="button" @click="close" class="border border-error text-error font-bold py-2 px-6 rounded-lg focus:outline-none focus:shadow-outline">
            Cancel
          </button>
        </div>
      </form>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import type { PropType } from 'vue';
import axios from 'axios';

export default defineComponent({
  name: 'ProgressiveBountySetup',
  props: {
    isModalVisible: {
      type: Boolean,
      required: true
    },
    issue: {
      type: Object as PropType<{ id: number; url: string; image_url: string }>,
      required: true
    },
    username: {
      type: String,
      required: true
    },
    bountyType: {
      type: String,
      default: 'progressive'
    }
  },
  data() {
    const today = new Date().toISOString().split('T')[0];
    const nextYear = new Date();
    nextYear.setFullYear(nextYear.getFullYear() + 1);
    const nextYearDate = nextYear.toISOString().split('T')[0];
    
    return {
      individualAmount: 0.0,
      stepFrequency: 1,
      bountyStart: today,
      bountyEnd: nextYearDate,
      totalBounties: 0,
      totalAmount: 0.0
    };
  },
  computed: {
    formattedTotalBounties(): string {
      return this.calculateTotalBounties().toString().replace(/\d(?=(\d{3})+$)/g, '$&,');
    },
    formattedTotalAmount(): string {
      return this.calculateTotalAmount().toFixed(2).replace(/\d(?=(\d{3})+\.)/g, '$&,');
    }
  },
  methods: {
    close() {
      this.$emit('close-modal');
    },
    async submit() {
      if (new Date(this.bountyEnd) > new Date(this.bountyStart) && this.individualAmount > 0) {
        await this.submitBounty();
        this.close();
      } else {
        alert('Please ensure the end date is after the start date and all inputs are valid.');
      }
    },
    async submitBounty() {
      const baseURL = import.meta.env.VITE_API_BASE_URL;
      if (!baseURL) {
        console.error('API base URL is not set.');
        alert('API base URL is not set.');
        return;
      }
      let startDate = new Date(this.bountyStart);
      const endDate = new Date(this.bountyEnd);

      if (startDate >= endDate) {
        alert('The start date must be before the end date.');
        return;
      }

      try {
        while (startDate < endDate) {
          const response = await axios.post(`${baseURL}/api/v1/bounties/`, {
            amount: this.individualAmount,
            currency: 'EUR',
            issue_github_id: this.issue.id,
            issue_github_url: this.issue.url,
            issue_image_url: this.issue.image_url,
            user_github_login: this.username,
            start_at: startDate.toISOString(),
            end_at: endDate.toISOString(),
          });
          console.log('Bounty submitted:', response.data);

          // Increment the start date by the frequency
          startDate.setDate(startDate.getDate() + this.stepFrequency);
        }
      } catch (error) {
        console.error('Error submitting bounty:', error);
        alert('Failed to submit multiple bounties. Please check your network connection and try again.');
      }
    },
    calculateTotalBounties(): number {
      const days = (new Date(this.bountyEnd).getTime() - new Date(this.bountyStart).getTime()) / (1000 * 60 * 60 * 24);
      return Math.floor(days / this.stepFrequency) + 1;
    },
    calculateTotalAmount(): number {
      return this.calculateTotalBounties() * this.individualAmount;
    }
  }
});
</script>

<style scoped>
/* Custom styles if needed */
</style>
