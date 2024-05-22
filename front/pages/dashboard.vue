<template>
  <div class="container mx-auto px-4 py-12">
    <h1 class="text-4xl font-bold text-center mb-6 text-gray-800">Welcome, {{ username }}!</h1>
    <h2 class="text-2xl font-bold text-center mb-6 text-gray-600">Issues you're following:</h2>
    <div class="overflow-auto h-screen">
      <ul class="space-y-4" v-if="username">
        <li v-for="issue in issues" :key="issue.id" class="issue-item" >
          <IssueItem v-bind="issue" :issue="issue" :username="username" :bounty="issue.bounty"/>
        </li>
      </ul>
    </div>
  </div>
</template>


<script>
import IssueItem from '@/components/IssueItem.vue';

import axios from 'axios';

export default {
  data() {
    return {
    username: '',
    userId: null,
    issues: [],
    userBackground: 'default-image.jpg',
    repoImages: {},
    currentIssue: null,
    bountyAmount: 0,
    bountyCurrency: 'EUR',
    bountyCurrency: 'EUR',
    bountyStart: '',
    bountyEnd: ''
  };
  },

  created() {
    this.initializeData();
  },

  methods: {
    initializeData() {
      if (!this.$auth.loggedIn) return;

      this.fetchUserData()
        .then(() => this.fetchOrganizationsAndRepos())
        .catch(error => console.error('Initialization failed:', error));
    },

    async fetchUserData() {
      const response = await axios.get('https://api.github.com/user', {
        headers: { Authorization: this.getAuthHeader() }
      });
      this.username = response.data.login;
      this.userBackground = response.data.avatar_url || 'default-image.jpg';
    },

    async fetchOrganizationsAndRepos() {
      const orgs = await this.fetchOrganizations();
      const repos = await this.fetchReposIncludingUser(orgs);
      await this.processAllRepositories(repos);
    },

    async fetchOrganizations() {
      let orgs = [];
      let url = `https://api.github.com/users/${this.username}/orgs`;

      while (url) {
        const response = await axios.get(url, { headers: { Authorization: this.getAuthHeader() } });
        orgs = orgs.concat(response.data);
        url = this.getNextPageUrl(response.headers);
      }
      return orgs;
    },

    async fetchReposIncludingUser(orgs) {
  let allRepos = [];
  const repoUrls = orgs.map(org => `https://api.github.com/orgs/${org.login}/repos?visibility=all`); // Fetch all visibility levels

  // Include user's own repositories, including private ones
  repoUrls.push(`https://api.github.com/user/repos?visibility=all`);

  for (const url of repoUrls) {
    let repoUrl = url;
    while (repoUrl) {
      const response = await axios.get(repoUrl, { headers: { Authorization: this.getAuthHeader() } });
      allRepos = allRepos.concat(response.data);
      repoUrl = this.getNextPageUrl(response.headers);
    }
  }

  return allRepos;
},



    async processAllRepositories(repos) {
      const issues = await this.fetchAndProcessIssues(repos);
      this.issues = this.issues.concat(issues);
  const bounties = await this.fetchBounties(); // Fetch sums of bounties for each issue
  const issuesWithBounties = issues.map(issue => ({
    ...issue,
    bounty: bounties[issue.id] || 0 // Attach the sum of bounties if it exists
  }));
  this.issues = issuesWithBounties;
},

    async fetchAndProcessIssues(repos) {
      const issuesPromises = repos.map(repo =>
        this.fetchIssuesForRepo(repo).catch(error => {
          console.error(`Error fetching issues for ${repo.full_name}:`, error);
          return [];
        })
      );

      const issuesResults = await Promise.all(issuesPromises);
      return issuesResults.flat();
    },

    async fetchIssuesForRepo(repo) {
      let issues = [];
      let url = `https://api.github.com/repos/${repo.full_name}/issues`;

      while (url) {
        const response = await axios.get(url, { headers: { Authorization: this.getAuthHeader() } });
        issues = issues.concat(response.data);
        url = this.getNextPageUrl(response.headers);
      }
      // add data to each issue
      issues = issues.map(issue => ({
        ...issue,
        repository_name: repo.name,
        repository_stars: repo.stargazers_count,
        image_url: repo.owner.avatar_url || 'default-image.png'
      }));
      return issues
    },

    getAuthHeader() {
      return `${this.$auth.strategy.token.get()}`;
    },

    getNextPageUrl(headers) {
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
    },

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
}




  },
  components: {
    IssueItem,
  }
};
</script>




<style scoped>
  .cursor-pointer {
    cursor: pointer;
  }
</style>


