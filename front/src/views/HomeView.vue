<template>
  <section class="min-h-screen flex flex-col items-center justify-center p-4">
    <div v-if="issues.length > 0" class="w-full mt-8 p-6 rounded-md shadow-md">
      <h2 class="text-2xl font-semibold text-primary mb-4">{{ $t('Available Issues') }}</h2>
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
import { defineComponent, ref, onMounted, onUnmounted } from 'vue';
import axios from 'axios';
import PrivateIssueListItem from '../components/PrivateIssueListItem.vue';
import IssueListItem from '../components/IssueListItem.vue';
import { useUserStore } from '../stores/user'
import { storeToRefs } from 'pinia'

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
  state?: string;
}

export default defineComponent({
  name: 'HomeView',
  components: {
    PrivateIssueListItem,
    IssueListItem
  },
  setup() {
    const issues = ref<Issue[]>([]);
    const loading = ref(true);
    const userStore = useUserStore();

    const fetchBounties = async () => {
      loading.value = true;
      const baseURL = import.meta.env.VITE_API_BASE_URL as string;

      if (!baseURL) {
        console.error("API base URL is not set.");
        loading.value = false;
        return;
      }

      try {
        const response = await axios.get(`${baseURL}/api/v1/bounties/`, { headers: {
          Authorization: userStore.authHeader,
          RefererPolicy: 'origin-when-cross-origin'
        } });
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
              bounty_type: bounty.bounty_type,
              issueUrl: bounty.issue_github_url,
              currency: bounty.currency,
              issue_image_url: bounty.issue_image_url,
              user_github_login: bounty.user_github_login,
              start_at: bounty.start_at,
              end_at: bounty.end_at
            };
          }
          const startDate = new Date(bounty.start_at);
          const endDate = new Date(bounty.end_at);
          const timeElapsed = (currentDate.getTime() - startDate.getTime()) / (endDate.getTime() - startDate.getTime());

          let adjustedAmount = parseFloat(bounty.amount);
          if (bounty.bounty_type === 'crescendo') {
            adjustedAmount *= timeElapsed;
          } else if (bounty.bounty_type === 'decrescendo') {
            adjustedAmount *= (1 - timeElapsed);
          }

          acc[id].amount += adjustedAmount;
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
            end_at: issueTotals[id].end_at,
            state: issueTotals[id].state,
          };
          return await fetchGitHubIssueData(issue);
        }));

        issues.value = issues.value.filter((issue) => {
          const startDate = new Date(issue.start_at);
          const endDate = new Date(issue.end_at);
          return currentDate > startDate && currentDate < endDate && issue.state === 'open';
        });

        issues.value = issues.value.sort((a, b) => b.amount - a.amount);
      } catch (error) {
        console.error('Error fetching bounties:', error);
      } finally {
        loading.value = false;
      }
    };

    const fetchGitHubIssueData = async (issue: Issue) => {
      const issueUrlParts = issue.issue_github_url.split('/');
      const owner = issueUrlParts[issueUrlParts.length - 4];
      const repo = issueUrlParts[issueUrlParts.length - 3];
      const issueNumber = issueUrlParts[issueUrlParts.length - 1];

      try {
        const response = await axios.get(`https://api.github.com/repos/${owner}/${repo}/issues/${issueNumber}`);
        // Merge fetched data with existing issue data
        issue.title = response.data.title;
        issue.description = response.data.body;
        issue.created_at = response.data.created_at;
        issue.updated_at = response.data.updated_at;
        issue.state = response.data.state;
        return issue;
      } catch (error) {
        console.error('Error fetching GitHub issue data without token:', error);
        
        // If unauthenticated request fails, try with token (assuming it is a private issue)
        const token = ''; // Replace with the appropriate way to get the auth token
        try {
          const response = await axios.get(`https://api.github.com/repos/${owner}/${repo}/issues/${issueNumber}`, { headers:  { Authorization: userStore.authGithubHeader } });
          // Merge fetched data with existing issue data
          issue.title = response.data.title;
          issue.description = response.data.body;
          issue.created_at = response.data.created_at;
          issue.updated_at = response.data.updated_at;
          issue.state = response.data.state;
          return issue;
        } catch (authError) {
          console.error('Error fetching GitHub issue data with token:', authError);
          return { ...issue, is_private: true }; // Fallback for errors or private issues
        }
      }
    };

    onMounted(() => {
      fetchBounties();
      const interval = setInterval(fetchBounties, 60000); // Refresh every 60 seconds
      onUnmounted(() => clearInterval(interval)); // Clear interval when component is unmounted
    });

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
