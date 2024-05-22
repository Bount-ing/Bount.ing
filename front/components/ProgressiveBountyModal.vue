<template>
  <div v-if="isModalVisible && bountyType === 'progressive'" class="modal fixed inset-0 bg-gray-800 bg-opacity-50 flex justify-center items-center">
    <div class="modal-content bg-white p-6 rounded-lg shadow-lg max-w-md w-full">
      <h3 class="text-xl font-semibold mb-6 text-gray-800">Create a Progressive Bounty</h3>
      <form @submit.prevent="submit">
        <label class="block mb-4">
          <span class="text-gray-700">Bounty Amount (EUR):</span>
          <input type="number" v-model.number="individualAmount" placeholder="Enter amount" min="1" class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-300 focus:ring focus:ring-indigo-200 focus:ring-opacity-50">
        </label>
        <label class="block mb-4">
          <span class="text-gray-700">Frequency (Days):</span>
          <input type="number" v-model.number="stepFrequency" placeholder="Enter frequency in days" min="1" class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-300 focus:ring focus:ring-indigo-200 focus:ring-opacity-50">
        </label>
        <label class="block mb-4">
          <span class="text-gray-700">Start Date:</span>
          <input type="date" v-model="bountyStart" class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-300 focus:ring focus:ring-indigo-200 focus:ring-opacity-50">
        </label>
        <label class="block mb-4">
          <span class="text-gray-700">End Date:</span>
          <input type="date" v-model="bountyEnd" class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-300 focus:ring focus:ring-indigo-200 focus:ring-opacity-50">
        </label>
        <label class="block mb-4">
          <div class="stats mt-4">
  <p class="text-gray-800">Total Bounties Created: <strong>{{ formattedTotalBounties }}</strong></p>
  <p class="text-gray-800">Total Amount: <strong>{{ formattedTotalAmount }} EUR</strong></p>
</div>

        </label>
        <div class="flex justify-between mt-6">
          <button type="submit" class="btn bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-6 rounded-lg focus:outline-none focus:shadow-outline">Submit</button>
          <button @click="close" type="button" class="btn bg-red-500 hover:bg-red-700 text-white font-bold py-2 px-6 rounded-lg focus:outline-none focus:shadow-outline">Cancel</button>
        </div>
      </form>
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
      default: 'progressive'
    }
  },
  data() {
    return {
      individualAmount: 0.0,
      stepFrequency: 1,
      bountyStart: '',
      bountyEnd: '',
      totalBounties: 0,     // To keep track of the number of bounties
      totalAmount: 0.0      // To keep track of the total amount of bounties in EUR
    };
  },

  computed: {
    formattedTotalBounties() {
      return this.calculateTotalBounties().toString().replace(/\d(?=(\d{3})+$)/g, '$&,');
    },
    formattedTotalAmount() {
      return this.calculateTotalAmount().toFixed(2).replace(/\d(?=(\d{3})+\.)/g, '$&,');
    }
  },

  methods: {
    close() {
      this.isModalVisible = false;
      this.$emit('update:isModalVisible', false);
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
      let startDate = new Date(this.bountyStart);
      const endDate = new Date(this.bountyEnd);

      if (startDate >= endDate) {
        alert('The start date must be before the end date.');
        return;
      }

      try {
        while (startDate < endDate) {
          const response = await axios.post('http://0.0.0.0:8080/api/v1/bounties/', {
            amount: parseFloat(this.individualAmount),
            currency: 'EUR',
            issue_github_id: this.issue.id,
            issue_github_url: this.issue.url,
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
    calculateTotalBounties() {
      const days = (new Date(this.bountyEnd) - new Date(this.bountyStart)) / (1000 * 60 * 60 * 24);
      return Math.floor(days / this.stepFrequency) + 1;
    },
    calculateTotalAmount() {
      return this.calculateTotalBounties() * this.individualAmount;
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
