<template>
  <section class="min-h-screen flex flex-col items-center justify-center p-4">
    <div v-if="loading" class="text-center text-2xl font-semibold text-gray-800">Loading...</div>
    <div v-else class="w-full max-w-6xl rounded-2xl shadow-xl overflow-hidden">
      <div class="relative">
        <img src="/bount.ing.banner.png" alt="Bount.ing Banner" class="w-full h-64 object-cover">
        <div class="absolute inset-0 bg-black opacity-30"></div>
        <div class="absolute inset-0 flex items-center justify-center">
          <h1 class="text-3xl md:text-5xl lg:text-6xl font-bold text-white text-center">
            Empower Open Source Innovation
          </h1>
        </div>
      </div>
      <div class="px-6 md:px-10 lg:px-16 py-8 space-y-6 text-center">
        <p class="text-md md:text-lg lg:text-xl text-primary">
          Enhance your coding skills, foster innovation, and contribute to the growth of open source projects while earning rewards.
        </p>
      </div>
    </div>
    <div v-if="issues.length > 0" class="w-full mt-8 p-6 rounded-md shadow-md">
      <h2 class="text-2xl font-semibold text-primary mb-4">Available Issues</h2>
      <ul class="space-y-3">
        <li v-for="issue in issues" :key="issue.id" class="issue-item rounded-lg shadow-lg border border-primary">
          <PrivateIssueListItem v-if="issue.is_private" :issue="issue" :bounty="issue.amount" />
          <IssueListItem v-else v-bind="issue" :issue="issue" :bounty="issue.amount" />
        </li>
      </ul>
    </div>
  </section>
</template>

<script lang="ts">
import { defineComponent, ref, onMounted } from 'vue';
import axios from 'axios';
import PrivateIssueItem from '../components/PrivateIssueListItem.vue';
import IssueItem from '../components/IssueListItem.vue';

interface Issue {
  id: string;
  amount: number;
  currency: string;
  issue_github_url: string;
  issue_image_url: string;
  user_github_login: string;
  start_at: string;
  end_at: string;
  title?: string;
  description?: string;
  created_at?: string;
  updated_at?: string;
  is_private?: boolean;
}

export default defineComponent({
  name: 'Home',
  components: {
    PrivateIssueItem,
    IssueItem
  },
  setup() {
    const issues = ref<Issue[]>([]);
    const loading = ref(true);

    const fetchBounties = async () => {
      loading.value = true;
      const baseURL = import.meta.env.VITE_API_BASE_URL as string;

      if (!baseURL) {
        console.error("API base URL is not set.");
        loading.value = false;
        return;
      }

      try {
        const response = await axios.get(`${baseURL}/api/v1/bounties/`);
        const currentDate = new Date();

        const activeBounties = response.data.filter((bounty: any) => {
          const startDate = new Date(bounty.start_at);
          const endDate = new Date(bounty.end_at);
          return currentDate > startDate && currentDate < endDate;
        });

        const issueTotals = activeBounties.reduce((acc: Record<string, any>, bounty: any) => {
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

        issues.value = await Promise.all(Object.keys(issueTotals).map(async (id) => {
          const issue: Issue = {
            id,
            amount: issueTotals[id].amount,
            currency: issueTotals[id].currency,
            issue_github_url: issueTotals[id].issueUrl,
            issue_image_url: issueTotals[id].issue_image_url,
            user_github_login: issueTotals[id].user_github_login,
            start_at: issueTotals[id].start_at,
            end_at: issueTotals[id].end_at
          };
          return await fetchGitHubIssueData(issue);
        }));

        issues.value = issues.value.filter((issue) => {
          const startDate = new Date(issue.start_at);
          const endDate = new Date(issue.end_at);
          return currentDate > startDate && currentDate < endDate;
        });

        issues.value = issues.value.sort((a, b) => b.amount - a.amount);
      } catch (error) {
        console.error('Error fetching bounties:', error);
      } finally {
        loading.value = false;
      }
    };

    const fetchGitHubIssueData = async (issue: Issue) => {
      const token = ''; // Replace with the appropriate way to get the auth token
      try {
        const response = await axios.get(issue.issue_github_url, { headers: { Authorization: `Bearer ${token}` } });
        // Merge fetched data with existing issue data
        issue.title = response.data.title;
        issue.description = response.data.body;
        issue.created_at = response.data.created_at;
        issue.updated_at = response.data.updated_at;
        return issue;
      } catch (error) {
        console.error('Error fetching GitHub issue data:', error);
        return { ...issue, is_private: true }; // Fallback for errors or private issues
      }
    };

    onMounted(fetchBounties);

    return {
      issues,
      loading
    };
  }
});
</script>

<style scoped>
/* Add your styles here */
</style>
