<template>
  <section class="min-h-screen flex flex-col items-center justify-center p-4">
    <div v-if="loading" class="text-center">Loading...</div>
    <div v-else class="w-full max-w-5xl bg-white rounded-2xl shadow-xl overflow-hidden">
      <div class="px-6 md:px-10 lg:px-16 py-8 space-y-6">
        <h1 class="text-2xl md:text-4xl lg:text-5xl font-bold text-gray-900 text-center">Empower Open Source Innovation</h1>
        <p class="text-sm md:text-lg lg:text-xl text-gray-700 text-center">
          Enhance your coding skills, foster innovation, and contribute to the growth of open source projects while earning rewards.
        </p>
      </div>
    </div>
    <div v-if="issues.length > 0" class="w-full mt-6 p-6  rounded-md shadow">
      <ul class="space-y-3 text-gray-800">
        <li v-for="issue in issues" :key="issue.id" class="issue-item">
          <IssueItem v-bind="issue" :issue="issue" :bounty="issue.amount"/>
        </li>
      </ul>
    </div>
  </section>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      issues: [],
      loading: false
    };
  },
  name: 'Home',
  methods: {
    async fetchBounties() {
  const response = await axios.get('http://0.0.0.0:8080/api/v1/bounties/');
  const currentDate = new Date(); // Current date for comparison

  return response.data.reduce((acc, bounty) => {
    const startDate = new Date(bounty.start_at);
    const endDate = new Date(bounty.end_at);

    // Check if the current date is between the start and end dates
    if (currentDate > startDate && currentDate < endDate) {
      // Initialize the sum for the issue if it hasn't been added yet
      if (!acc[bounty.issue_github_id]) {
        acc[bounty.issue_github_id] = 0; // Use zero directly, no need to parse it
      }
      // Add the current bounty's amount to the sum
      acc[bounty.issue_github_id] += parseFloat(bounty.amount);
    }
    return acc;
  }, {});
},

    async fetchGitHubIssueData(issue) {
      try {
        const response = await axios.get(issue.issue_github_url, {
          headers: { 'Authorization': 'token YOUR_GITHUB_PERSONAL_ACCESS_TOKEN' }
        });
        Object.assign(issue, response.data);
      } catch (error) {
        console.error('Error fetching GitHub issue data:', error);
      }
    },
    claimBounty(issue) {
      console.log(`Claiming bounty for issue ${issue.github_id}`);
      // Logic to claim bounty
    },
    raiseBounty(issue) {
      console.log(`Raising bounty for issue ${issue.github_id}`);
      // Logic to raise bounty
    }
  },
  created() {
    this.fetchBounties();
  }
}
</script>

<style>
/* Add your CSS styling here */
</style>
