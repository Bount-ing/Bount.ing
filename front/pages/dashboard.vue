<template>
  <div class="pt-16 relative bg-cover bg-no-repeat bg-center h-[60vh] md:h-[90vh] lg:h-[100vh]" :style="{ backgroundImage: `url(${userBackground})` }">
    <div class="absolute inset-0 bg-gradient-to-r from-cyan-500 via-purple-500 to-pink-500 opacity-90 mix-blend-multiply"></div>
    <div class="container mx-auto px-4 py-12 flex flex-col items-center justify-center h-full">
      <h1 class="text-6xl md:text-9xl lg:text-[10rem] font-black text-transparent bg-clip-text bg-gradient-to-br from-yellow-300 via-red-500 to-pink-600 text-center animate__animated animate__bounceIn">
        {{ username }}'s GitHub Playground
      </h1>
      <p class="text-xl md:text-3xl lg:text-4xl text-gray-100 text-center mt-6 animate__animated animate__fadeInUpBig shadow-lg">
        Dive into a universe of code, collaboration, and boundless creativity.
      </p>
      <div class="flex mt-6 w-full max-w-6xl bg-white bg-opacity-30 backdrop-filter backdrop-blur-lg p-6 rounded-xl shadow-xl overflow-hidden">
        <div class="flex-1 overflow-auto h-96 mr-4">
          <!-- Scrollable container for repos, with fixed height to manage overflow -->
          <ul class="space-y-4 animate__animated animate__fadeIn">
            <li v-for="repo in repos" :key="repo.id" @click="selectRepo(repo.name)" class="cursor-pointer flex items-center justify-between p-4 bg-gradient-to-r from-gray-700 to-gray-900 rounded-lg shadow hover:shadow-lg transition duration-300 ease-in-out">
              <div>
                <h3 class="text-xl text-white font-bold">{{ repo.name }}</h3>
                <p class="text-gray-300">Updated {{ repo.updated_at | formatDate }}</p>
              </div>
            </li>
          </ul>
        </div>
        <div class="flex-1 overflow-auto h-96">
          <!-- Scrollable container for issues of the selected repo -->
          <ul class="space-y-4 animate__animated animate__fadeIn">
            <li v-for="issue in issues" :key="issue.id" class="flex items-center justify-between p-4 bg-gradient-to-r from-gray-700 to-gray-900 rounded-lg shadow hover:shadow-lg transition duration-300 ease-in-out">
              <div class="flex-1">
                <h3 class="text-xl text-white font-bold">{{ issue.title }}</h3>
                <p class="text-gray-300">{{ issue.body }}</p>
              </div>
              <a :href="issue.html_url" target="_blank" class="text-teal-300 hover:text-teal-100 transition duration-150 ease-in-out">
                View &rarr;
              </a>
            </li>
          </ul>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      username: '',
      repos: [],
      issues: [],
      userBackground: 'default-image.jpg', // Placeholder for background image
      selectedRepo: ''
    };
  },
  created() {
    this.fetchUserData();
    this.fetchRepositories();
  },
  filters: {
    formatDate(value) {
      if (value) {
        return new Date(value).toLocaleDateString(undefined, { day: 'numeric', month: 'short', year: 'numeric' });
      }
      return '';
    }
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
        this.username = response.data.login; // 'login' is the username
      })
      .catch(error => {
        console.error('Error fetching user data:', error);
      });
    },
    fetchRepositories() {
      if (!this.$auth.loggedIn) return;

      axios.get('https://api.github.com/user/repos', {
        headers: {
          Authorization: `${this.$auth.strategy.token.get()}`
        }
      })
      .then(response => {
        this.repos = response.data;
      })
      .catch(error => {
        console.error('Error fetching repos:', error);
      });
    },
    selectRepo(repoName) {
      this.selectedRepo = repoName;
      this.fetchIssues(repoName);
    },
    fetchIssues(repoName) {
      if (!this.$auth.loggedIn) return;

      axios.get(`https://api.github.com/repos/${this.username}/${repoName}/issues`, {
        headers: {
          Authorization: `${this.$auth.strategy.token.get()}`
        }
      })
      .then(response => {
        this.issues = response.data;
      })
      .catch(error => {
        console.error('Error fetching issues:', error);
        this.issues = []; // Clear issues if there's an error or none found
      });
    }
  }
}
</script>

<style scoped>
  .cursor-pointer {
    cursor: pointer;
  }
</style>
