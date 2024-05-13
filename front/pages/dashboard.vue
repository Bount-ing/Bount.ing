<template>
  <div class="container mx-auto px-4 py-12 bg-secondary"> <!-- 'soft' used for a subtle main background -->
    <h1 class="text-4xl font-bold text-center mb-6 text-text font-serif">Welcome {{ username }} !</h1>
    <h2 class="text-2xl font-bold text-center mb-6 text-secondary font-serif">Issues you're following:</h2>
    <div class="overflow-auto h-screen">
      <ul class="space-y-4">
        <li v-for="issue in issues" :key="issue.id" class="flex items-start bg-secondary rounded-lg shadow-custom p-4 justify-between text-soft"> <!-- 'soft' for text to reduce harshness -->
          <img :src="issue.image_url" alt="Repo Image" class="w-20 h-20 rounded-full mr-4">
          <div class="flex flex-col justify-between flex-grow">
            <div class="flex flex-row items-center space-x-4">
              <h4 class="text-md font-bold text-accent">{{ issue.repo_owner }}</h4>
              <h3 class="text-xl font-bold text-secondary font-serif">{{ issue.repo_name }}</h3> <!-- 'secondary' for secondary information to keep it consistent -->
              </div>
              <div>
              <h5 class="text-lg font-semibold text-accent font-serif">{{ issue.title }}</h5>
              <p class="text-info">{{ issue.body }}</p> <!-- 'info' used here for a slightly softer contrast -->
            </div>
            <div class="flex flex-col">
              <a :href="issue.html_url" target="_blank" class="text-text hover:text-secondary font-serif"> <!-- 'secondary' for hover to maintain elegance -->
                View Issue on GitHub &rarr;
              </a>
            </div>
          </div>
          <div class="flex flex-col items-end space-y-2">
            <span class="bg-primary font-semibold text-lg px-5 py-2.5 rounded-lg"> <!-- 'soft' text for more readability against 'warning' background -->
              {{ issue.bounty || 0 }} €
            </span>
            <div class="flex space-x-2">
              <button @click="claimBounty(issue)" class="bg-info hover:bg-secondary text-dark font-bold py-2 px-4 rounded-lg h-12 w-20"> <!-- 'secondary' hover for a subtle transition -->
                Claim
              </button>
              <button @click="raiseBounty(issue)" class="bg-success hover:bg-accent text-dark font-bold py-2 px-4 rounded-lg h-12 w-20"> <!-- 'accent' hover for a mild emphasis -->
                Raise
              </button>
            </div>
          </div>
        </li>
      </ul>
    </div>
  </div>
</template>


<script>
import axios from 'axios';

export default {
  data() {
    return {
      username: '',
      issues: [],
      userBackground: 'default-image.jpg',
      repoImages: {}
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

      // Filter and map as needed
      return issues.filter(issue => issue.user.login === this.username || issue.body.includes(`@${this.username}`))
                   .map(issue => ({ ...issue, image_url: issue.user.avatar_url, repo_owner: repo.owner.login, repo_name: repo.name }));
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
    }
  },
};
</script>




<style scoped>
  .cursor-pointer {
    cursor: pointer;
  }
</style>
