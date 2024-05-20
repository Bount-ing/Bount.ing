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
    <div v-if="issues.length > 0" class="w-full mt-6 p-6 bg-white rounded-md shadow">
      <ul class="space-y-3 text-gray-800">
        <li v-for="issue in issues" :key="issue.id" class="flex flex-col md:flex-row items-start bg-gray-100 rounded-lg shadow-lg p-4 justify-between">
          <img :src="issue.image_url || 'default-image.png'" alt="Repo Image" class="w-20 h-20 rounded-full mr-4">
          <div class="flex-grow space-y-3">
            <div class="flex items-center space-x-4">
              <h4 class="text-sm md:text-md font-bold text-blue-500">{{ issue.user_github_login }}</h4>
            </div>
            <div>
              <h5 class="text-lg font-semibold text-blue-500">{{ issue.title }}</h5>
              <p class="text-gray-700">{{ issue.description }}</p>
            </div>
            <div class="space-y-2">
              <p class="text-xs text-gray-500">Repository: {{ issue.repository_name }} (⭐ {{ issue.repository_stars }})</p>
              <p class="text-xs text-gray-500">Created: {{ new Date(issue.created_at).toLocaleDateString() }}</p>
              <p class="text-xs text-gray-500">Last Updated: {{ new Date(issue.updated_at).toLocaleDateString() }}</p>
              <a :href="`https://github.com/issues/${issue.github_id}`" target="_blank" class="text-blue-500 hover:text-blue-600">
                View Issue on GitHub &rarr;
              </a>
            </div>
          </div>
          <div class="flex flex-col items-end space-y-2">
            <span class="bg-blue-500 text-white font-semibold text-md px-4 py-2 rounded-lg">
              {{ issue.amount }} €
            </span>
            <div class="flex space-x-2">
              <button @click="claimBounty(issue)" class="bg-blue-300 hover:bg-blue-400 text-white font-bold py-2 px-4 rounded-lg">
                Claim
              </button>
              <button @click="raiseBounty(issue)" class="bg-green-500 hover:bg-green-600 text-white font-bold py-2 px-4 rounded-lg">
                Raise
              </button>
            </div>
          </div>
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
      this.loading = true;
      try {
        const response = await axios.get('http://0.0.0.0:8080/api/v1/bounties/');
        this.issues = response.data;
        await Promise.all(this.issues.map(issue => this.fetchGitHubIssueData(issue)));
        this.issues.sort((a, b) => b.amount - a.amount); // Sorting by bounty amount in descending order
      } catch (error) {
        console.error('Error fetching bounties:', error);
      } finally {
        this.loading = false;
      }
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
