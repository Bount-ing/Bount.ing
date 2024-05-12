<template>
  <div class="container mx-auto px-4 py-12 bg-soft"> <!-- Updated to use 'soft' for main background -->
    <h1 class="text-4xl font-bold text-center mb-6 text-highlight font-serif">Welcome {{ username }} !</h1>
    <h2 class="text-2xl font-bold text-center mb-6 text-highlight font-serif">Issues your following:</h2>
    <div class="overflow-auto h-screen">
      <ul class="space-y-4">
        <li v-for="issue in issues" :key="issue.id" class="flex items-start bg-dark rounded-lg shadow-custom p-4 justify-between text-light"> <!-- Updated 'bg-dark' to 'bg-soft' and 'shadow-lg' to 'shadow-custom' -->
          <img :src="issue.image_url" alt="Repo Image" class="w-20 h-20 rounded-full mr-4">
          <div class="flex flex-col justify-between flex-grow">
            <div class="flex flex-row items-center space-x-4">
              <h4 class="text-md font-bold text-primary">{{ issue.repo_owner }}</h4>
              <h3 class="text-xl font-bold text-highlight font-serif">{{ issue.repo_name }}</h3> <!-- 'primary' for primary info -->
              </div>
              <div>
              <h5 class="text-lg font-semibold text-accent font-serif">{{ issue.title }}</h5>
              <p class="text-info">{{ issue.body }}</p> <!-- 'accent' for less important text -->
            </div>
            <div class="flex flex-col">
              <a :href="issue.html_url" target="_blank" class="text-primary hover:text-highlight font-serif"> <!-- 'highlight' for hover -->
                View Issue on GitHub &rarr;
              </a>
            </div>
          </div>
          <div class="flex flex-col items-end space-y-2">
            <span class="text-dark bg-warning font-semibold text-lg px-5 py-2.5 rounded-lg"> <!-- No change needed here as 'warning' and 'dark' are correctly used -->
              {{ issue.bounty || 0 }} €
            </span>
            <div class="flex space-x-2">
              <button @click="claimBounty(issue)" class="bg-info hover:bg-primary text-soft font-bold py-2 px-4 rounded-lg h-12 w-40"> <!-- Updated 'text-light' to 'text-soft' -->
                Claim Bounty
              </button>
              <button @click="raiseBounty(issue)" class="bg-success hover:bg-primary text-dark font-bold py-2 px-4 rounded-lg h-12 w-40"> <!-- No change needed here as 'success' and 'primary' are correctly used -->
                Raise Bounty
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
      userBackground: 'default-image.jpg', // Placeholder for background image
      repoImages: {} // Object to store repository-specific images
    };
  },
  created() {
    this.fetchUserData();
  },
  methods: {
    fetchUserData() {
      if (!this.$auth.loggedIn) return;

      axios.get('https://api.github.com/user', {
        headers: {
          Authorization: `${this.$auth.strategy.token.get()}`
        }
      })
      .then(response => {
        this.username = response.data.login;
        this.userBackground = response.data.avatar_url || 'default-image.jpg';

        // Once user data is fetched, fetch organizations and repos
        this.fetchOrganizationsAndRepos();
      })
      .catch(error => {
        console.error('Error fetching user data:', error);
      });
    },
    async fetchOrganizationsAndRepos() {
      if (!this.$auth.loggedIn) return;

      try {
        const orgsResponse = await axios.get(`https://api.github.com/users/${this.username}/orgs`, {
          headers: {
            Authorization: `${this.$auth.strategy.token.get()}`
          }
        });

        // Collect repositories from user and each organization
        const repoUrls = orgsResponse.data.map(org => `https://api.github.com/orgs/${org.login}/repos`);
        repoUrls.push('https://api.github.com/user/repos'); // User's own repos

        const repoPromises = repoUrls.map(url => axios.get(url, {
          headers: {
            Authorization: `${this.$auth.strategy.token.get()}`
          }
        }));

        const reposResponses = await Promise.all(repoPromises);

        // Process each repo list and fetch issues
        for (const response of reposResponses) {
          if (response && response.data) {
            await this.processRepositories(response.data);
          }
        }
      } catch (error) {
        console.error('Error fetching organizations or repositories:', error);
      }
    },
    async processRepositories(repos) {
  const issuesPromises = repos.map(repo => {
    if (repo) {
      return axios.get(`https://api.github.com/repos/${repo.full_name}/issues`, {
        headers: {
          Authorization: `${this.$auth.strategy.token.get()}`
        }
      }).then(async issueResponse => {
        let issues = [];

        // Process the first page of issues
        issues = issues.concat(issueResponse.data);

        // Check if there are more pages and fetch recursively
        let nextPageUrl = this.getNextPageUrl(issueResponse.headers);
        while (nextPageUrl) {
          // Fetch next page
          const nextPageResponse = await axios.get(nextPageUrl, {
            headers: {
              Authorization: `${this.$auth.strategy.token.get()}`
            }
          });
          issues = issues.concat(nextPageResponse.data);
          // Update next page URL
          nextPageUrl = this.getNextPageUrl(nextPageResponse.headers);
        }

        // Filter issues where the user is involved or mentioned
        const userIssues = issues.filter(issue => {
          return issue.user.login === this.username || issue.body.includes(`@${this.username}`);
        });
        return userIssues.map(issue => ({
          ...issue,
          image_url: issue.user.avatar_url,
  repo_owner: repo.owner.login,
  repo_name: repo.name        }));
      }).catch(error => {
        console.error('Error fetching issues for repository:', error);
        return [];
      });
    }
  });

  const issuesResponses = await Promise.all(issuesPromises.filter(p => p !== undefined));
  this.issues = this.issues.concat(issuesResponses.flatMap(response => response));
},

getNextPageUrl(headers) {
  const linkHeader = headers['link'];
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
