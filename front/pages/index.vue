<template>
  <section class="min-h-screen flex flex-col items-center justify-center p-4">
    <div v-if="loading" class="text-center">Loading...</div>
    <div v-else class="w-full max-w-5xl bg-white rounded-2xl shadow-xl overflow-hidden">
      <div class="px-6 md:px-10 lg:px-16 py-8 space-y-6">
        <h1 class="text-2xl md:text-4xl lg:text-5xl font-bold text-gray-900 text-center">
          Empower Open Source Innovation
        </h1>
        <p class="text-sm md:text-lg lg:text-xl text-gray-700 text-center">
          Enhance your coding skills, foster innovation, and contribute to the growth of open source projects while earning rewards.
        </p>
      </div>
    </div>
    <div v-if="issues.length > 0" class="w-full mt-6 p-6 rounded-md shadow">
      <ul class="space-y-3 text-gray-800">
        <li v-for="issue in issues" :key="issue.id" class="issue-item">
          <PrivateIssueItem v-if="issue.is_private" :issue="issue" :bounty="issue.amount"/>
          <IssueItem v-else v-bind="issue" :issue="issue" :bounty="issue.amount"/>
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
      const baseURL = process.env.API_BASE_URL;
      if (!baseURL) {
        console.error("API base URL is not set.");
        this.loading = false;
        return;
      }
      try {
        const response = await axios.get(`${baseURL}/api/v1/bounties/`);
        const currentDate = new Date();

        const activeBounties = response.data.filter(bounty => {
          const startDate = new Date(bounty.start_at);
          const endDate = new Date(bounty.end_at);
          return currentDate > startDate && currentDate < endDate;
        });

        const issueTotals = activeBounties.reduce((acc, bounty) => {
          const id = bounty.issue_github_id;
          if (!acc[id]) {
            acc[id] = {
              amount: 0,
              issueUrl: bounty.issue_github_url,
              currency: bounty.currency,
              issue_image_url: bounty.issue_image_url,
              user_github_login: bounty.user_github_login,
              start_at: bounty.start_at,
              end_at: bounty.end_at
            };
          }
          acc[id].amount += parseFloat(bounty.amount);
          return acc;
        }, {});

        this.issues = await Promise.all(Object.keys(issueTotals).map(async id => {
          const issue = {
            id,
            amount: issueTotals[id].amount,
            currency: issueTotals[id].currency,
            issue_github_url: issueTotals[id].issueUrl,
            issue_image_url: issueTotals[id].issue_image_url,
            user_github_login: issueTotals[id].user_github_login,
            start_at: issueTotals[id].start_at,
            end_at: issueTotals[id].end_at
          };
          return await this.fetchGitHubIssueData(issue);
        }));

        this.issues = this.issues.filter(issue => {
          const startDate = new Date(issue.start_at);
          const endDate = new Date(issue.end_at);
          return currentDate > startDate && currentDate < endDate;
        });

        this.issues = this.issues.sort((a, b) => b.amount - a.amount);
      } catch (error) {
        console.error('Error fetching bounties:', error);
      } finally {
        this.loading = false;
      }
    },

    async fetchGitHubIssueData(issue) {
      const token = this.$auth.strategy.token.get();
      if (!token) {
        console.error("GitHub token is not set.");
        return { ...issue, is_private: true }; // Indicate inaccessible issue
      }
      try {
        const response = await axios.get(issue.issue_github_url, { headers: { Authorization: `${token}` } });
        // Merge fetched data with existing issue data
        this.$set(issue, 'title', response.data.title);
        this.$set(issue, 'description', response.data.body);
        this.$set(issue, 'created_at', response.data.created_at);
        this.$set(issue, 'updated_at', response.data.updated_at);
        return issue;
      } catch (error) {
        console.error('Error fetching GitHub issue data:', error);
        return { ...issue, is_private: true }; // Fallback for errors or private issues
      }
    }
  },
  created() {
    this.fetchBounties();
  }
}
</script>
