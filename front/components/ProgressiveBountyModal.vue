<template>
  <div v-if="isModalVisible && bountyType === 'progressive'" class="modal fixed inset-0 bg-gray-800 bg-opacity-50 flex justify-center items-center">
    <div class="modal-content bg-white p-6 rounded-lg shadow-lg max-w-md w-full">
      <h3 class="text-xl font-semibold mb-6 text-gray-800">Create a Progressive Bounty</h3>

      <label class="block mb-4">
        <span class="text-gray-700">Individual Bounty Amount (EUR):</span>
        <input type="number" v-model="individualAmount" placeholder="Enter initial amount" min="1" class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-300 focus:ring focus:ring-indigo-200 focus:ring-opacity-50">
      </label>
      <label class="block mb-4">
        <span class="text-gray-700">Total Bounty Amount (EUR):</span>
        <input type="number" v-model="totalAmount" placeholder="Enter total amount" min="1" class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-300 focus:ring focus:ring-indigo-200 focus:ring-opacity-50">
      </label>
      <label class="block mb-4">
        <span class="text-gray-700">Step Amount (EUR):</span>
        <input type="number" v-model="stepAmount" placeholder="Minimum 1 EUR" min="1" class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-300 focus:ring focus:ring-indigo-200 focus:ring-opacity-50">
      </label>
      <label class="block mb-4">
        <span class="text-gray-700">Step Frequency (Days):</span>
        <input type="number" v-model="stepFrequency" placeholder="Frequency in days" min="1" class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-300 focus:ring focus:ring-indigo-200 focus:ring-opacity-50">
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
  name: 'ProgressiveBountyModal',
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
    },
    bountyType: {
      type: String,
      default: 'progressive'  // Ensures this component only handles progressive bounties
    }
  },
  data() {
    return {
      individualAmount: 0.0,
      totalAmount: 0.0,
      stepAmount: 1.0,
      stepFrequency: 1,
      bountyStart: '',
      bountyEnd: '',
      bountyCurrency: 'EUR',
      ignoreWatch: false
    };
  },
  methods: {
    close() {
      this.isModalVisible = false;
      this.$emit('update:isModalVisible', false);
    },
    async submit() {
      if (this.validateDates()) {
        await this.submitMultipleBounties();
        this.close();
      } else {
        alert('Please ensure the end date is after the start date and all inputs are valid.');
      }
    },
    validateDates() {
      return new Date(this.bountyEnd) > new Date(this.bountyStart) && (this.individualAmount > 0 && this.totalAmount > 0);
    },
    async submitMultipleBounties() {
      let startDate = new Date(this.bountyStart);
      let currentAmount = parseFloat(this.individualAmount);

      while (startDate < new Date(this.bountyEnd) && currentAmount < this.totalAmount) {
        let endDate = new Date(startDate.getTime() + this.stepFrequency * 86400000);
        endDate = new Date(endDate - 1); // Ensuring end date is just before the next period starts

        if (endDate > new Date(this.bountyEnd)) {
          endDate = new Date(this.bountyEnd);
        }

        await this.submitBounty(currentAmount, startDate, endDate);

        // Update startDate for next interval
        startDate = new Date(endDate.getTime() + 1); // Start next period right after current period ends

        // Update the amount
        currentAmount += this.stepAmount;
        currentAmount = Math.min(currentAmount, this.totalAmount); // Ensure not to exceed total amount
      }
    },
    async submitBounty(amount, start_at, end_at) {
      try {
        const response = await axios.post('http://0.0.0.0:8080/api/v1/bounties/', {
          amount: parseFloat(amount),
          currency: this.bountyCurrency,
          issue_github_id: this.issue.id,
          issue_github_url: this.issue.url,
          user_github_login: this.username,
          start_at: new Date(start_at).toISOString(),
          end_at: new Date(end_at).toISOString(),
        });
        console.log('Bounty submitted:', response);
      } catch (error) {
        console.error('Error submitting bounty:', error);
      }
    }
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
