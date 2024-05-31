<template>
  <div class="container mx-auto px-4 py-12">
    <h1 class="text-4xl font-bold text-center mb-6 text-gray-800">Welcome, {{ username }}!</h1>
    <h2 class="text-2xl font-bold text-center mb-6 text-gray-600">Issues you're following:</h2>
    <div class="mt-8 p-6 rounded-md shadow-md">
      <ul class="space-y-3" v-if="username">
        <li v-for="issue in issues" :key="issue.id" class="issue-item rounded-lg shadow-lg border border-primary">
          <IssueListItem v-bind="issue" :issue="issue" :username="username" :bounty="issue.bounty" />
        </li>
      </ul>
      <div v-else class="text-center">
        <p class="text-lg text-primary-light">Please log in to see your issues.</p>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref, onMounted } from 'vue';
import IssueListItem from '../components/IssueListItem.vue';
import axios from 'axios';
import { useUserStore } from '../stores/user'
import { storeToRefs } from 'pinia'

interface Issue {
  id: number;
  repository_name: string;
  repository_stars: number;
  image_url: string;
  bounty: number;
  [key: string]: any; // For other possible properties
}

const userStore = useUserStore()

const { user, authGithubHeader, authHeader, isLoggedIn } = storeToRefs(userStore)

export default defineComponent({
  components: {
    IssueListItem,
  },
  setup() {
    const username = ref<string>('');
    const issues = ref<Issue[]>([]);
    const userBackground = ref<string>('default-image.jpg');
    const seenRepos = ref<Set<string>>(new Set());

    const initializeData = async () => {
      if (!isLoggedIn.value) return;

      try {
        fetchUserData();
        await fetchOrganizationsAndRepos();
      } catch (error) {
        console.error('Initialization failed:', error);
      }
    };

    const fetchUserData = () => {
      username.value = user.value.username;
      userBackground.value = user.value.avatar || 'default-image.jpg';
    };

    const fetchOrganizationsAndRepos = async () => {
      const orgs = await fetchOrganizations();
      const repos = await fetchReposIncludingUser(orgs);
      await processAllRepositories(repos);
    };

    const fetchOrganizations = async () => {
      let orgs: any[] = [];
      let url = `https://api.github.com/users/${username.value}/orgs`;

      while (url) {
        const response = await axios.get(url, { headers: { Authorization: authGithubHeader.value } });
        orgs = orgs.concat(response.data);
        url = getNextPageUrl(response.headers);
      }
      return orgs;
    };

    const fetchReposIncludingUser = async (orgs: any[]) => {
      let allRepos: any[] = [];
      const repoUrls = orgs.map(org => `https://api.github.com/orgs/${org.login}/repos`);

      repoUrls.push(`https://api.github.com/user/repos`);

      for (const url of repoUrls) {
        let repoUrl = url;
        while (repoUrl) {
          const response = await axios.get(repoUrl, { headers: { Authorization: authGithubHeader.value } });
          response.data.forEach((repo: any) => {
            if (!seenRepos.value.has(repo.full_name)) {
              allRepos.push(repo);
              seenRepos.value.add(repo.full_name);
            }
          });
          repoUrl = getNextPageUrl(response.headers);
        }
      }

      return allRepos;
    };

    const processAllRepositories = async (repos: any[]) => {
      const issuesData = await fetchAndProcessIssues(repos);
      const bounties = await fetchBounties();
      const issuesWithBounties = issuesData.map(issue => ({
        ...issue,
        bounty: bounties[issue.id] || 0,
      }));
      issues.value = issuesWithBounties;
    };

    const fetchAndProcessIssues = async (repos: any[]) => {
      const issuesPromises = repos.map(repo =>
        fetchIssuesForRepo(repo).catch(error => {
          console.error(`Error fetching issues for ${repo.full_name}:`, error);
          return [];
        })
      );

      const issuesResults = await Promise.all(issuesPromises);
      return issuesResults.flat();
    };

    const fetchIssuesForRepo = async (repo: any) => {
      let issues: any[] = [];
      let url = `https://api.github.com/repos/${repo.full_name}/issues`;

      while (url) {
		const response = await axios.get(url, { headers: { Authorization: authGithubHeader.value } });
        issues = issues.concat(response.data);
        url = getNextPageUrl(response.headers);
      }

      issues = issues.map(issue => ({
        ...issue,
        repository_name: repo.name,
        repository_stars: repo.stargazers_count,
        image_url: repo.owner.avatar_url || 'default-image.png',
      }));

      return issues;
    };

    const fetchBounties = async () => {
	  const response = await axios.get('http://0.0.0.0:8080/api/v1/bounties/');
      const currentDate = new Date();

      return response.data.reduce((acc: Record<number, number>, bounty: any) => {
        const startDate = new Date(bounty.start_at);
        const endDate = new Date(bounty.end_at);

        if (currentDate > startDate && currentDate < endDate) {
          if (!acc[bounty.issue_github_id]) {
            acc[bounty.issue_github_id] = 0;
          }
          acc[bounty.issue_github_id] += parseFloat(bounty.amount);
        }
        return acc;
      }, {});
    };

    const getNextPageUrl = (headers: any) => {
      const linkHeader = headers.link;
      if (!linkHeader) return null;

      const links = linkHeader.split(',');
      for (const link of links) {
        const match = link.match(/<([^>]+)>;\s*rel="next"/);
        if (match) {
          return match[1];
        }
      }
      return null;
    };

    onMounted(() => {
      initializeData();
    });

    return {
      username,
      issues,
      userBackground,
    };
  },
});
</script>
