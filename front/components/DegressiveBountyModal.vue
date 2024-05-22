<template>
  <div v-if="isModalVisible" class="modal fixed inset-0 bg-gray-800 bg-opacity-50 flex justify-center items-center">
    <div class="modal-content bg-white p-6 rounded-lg shadow-lg max-w-md w-full">
      <h3 class="text-xl font-semibold mb-6 text-gray-800">Create Multiple Degressive Bounties</h3>
      <label class="block mb-4">
        <span class="text-gray-700">Total Bounty Amount (EUR):</span>
        <input type="number" v-model="totalAmount" placeholder="Enter total amount" min="1" class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-300 focus:ring focus:ring-indigo-200 focus:ring-opacity-50">
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
      <div v-if="isValidInput" class="mt-4 mb-6">
        <p class="text-gray-700">Number of Bounties to Create: {{ numberOfBounties }}</p>
        <p class="text-gray-700">Value per Bounty: {{ bountyValue.toFixed(2) }} EUR</p>
      </div>
      <div class="flex justify-between mt-6">
        <button @click="submitAllBounties" class="btn bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-6 rounded-lg focus:outline-none focus:shadow-outline">Submit All Bounties</button>
        <button @click="close" class="btn bg-red-500 hover:bg-red-700 text-white font-bold py-2 px-6 rounded-lg focus:outline-none focus:shadow-outline">Cancel</button>
      </div>
    </div>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  name: 'DegressiveBountyModal',
  props: {
    isModalVisible: Boolean,
    issue: Object,
    username: String
  },
  data() {
    return {
      totalAmount: 0,
      stepFrequency: 1,
      bountyStart: '',
      bountyEnd: '',
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
      this.isModalVisible = false;
      this.$emit('update:isModalVisible', false);
    },
    async submitBounty() {
  const startDate = new Date(this.bountyStart);
  let currentDate = new Date(this.bountyStart);
  const endDate = new Date(this.bountyEnd);

  if (startDate >= endDate) {
    console.error('The start date must be before the end date.');
    alert('The start date must be before the end date.');
    return;
  }

  if (isNaN(startDate) || isNaN(endDate)) {
    console.error('Invalid date inputs:', this.bountyStart, this.bountyEnd);
    alert('Invalid date inputs. Please check the dates.');
    return;
  }

  try {
    while (currentDate <= endDate) {
      const response = await axios.post('http://0.0.0.0:8080/api/v1/bounties/', {
        amount: parseFloat(this.bountyValue.toFixed(2)),
        currency: 'EUR',
        issue_github_id: this.issue.id,
        issue_github_url: this.issue.url,
        user_github_login: this.username,
        start_at: startDate.toISOString(),
        end_at: currentDate.toISOString(),
      });
      console.log('Bounty submitted:', response.data);

      // Increment the current date by the step frequency for the next end date
      currentDate.setDate(currentDate.getDate() + this.stepFrequency);
    }
  } catch (error) {
    console.error('Error submitting bounty:', error);
    alert('Failed to submit multiple bounties. Please check your network connection and try again.');
  }
},

    submitAllBounties() {
      if (this.isValidInput) {
        this.submitBounty();
        this.close();
      } else {
        alert('Please ensure all inputs are valid.');
      }
    }
  }
};
</script>
