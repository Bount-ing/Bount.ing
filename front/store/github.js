// store/index.js
import axios from 'axios';

export const state = () => ({
  issues: [],
  isLoading: false,
  token: null,
  username: null,
  error: null
});

export const getters = {
  getIssues: state => state.issues,
  isLoading: state => state.isLoading,
  error: state => state.error,
  token: state => state.token,
  username: state => state.username,
};

export const mutations = {
  setIssues(state, issues) {
    state.issues = issues;
  },
  setToken(state, token) {
    state.token = token;
  },
  setUsername(state, username) {
    state.username = username;
  },
  setLoading(state, isLoading) {
    state.isLoading = isLoading;
  },
  setError(state, error) {
    state.error = error;
  }
};

export const actions = {
  async fetchUsername({ commit, state }) {
    if (!state.token) {
      commit('setError', 'Authentication token is missing');
      return;
    }
    const headers = {
      'Authorization': `token ${state.token}`
    };
    try {
      const response = await axios.get('https://api.github.com/user', { headers });
      commit('setUsername', response.data.login);  // 'login' is the field where the username is stored
    } catch (error) {
      commit('setError', error.response ? error.response.data : error.message);
    }
  },
  async fetchIssues({ state, commit }, repo) {
    if (!state.username) {
      commit('setError', 'Username is undefined');
      commit('setLoading', false);
      return;
    }

    commit('setLoading', true);
    commit('setError', null);

    const headers = {
      'Authorization': `token ${state.token}`
    };
    const url = `https://api.github.com/repos/${state.username}/${repo}/issues`;
    const params = {
      creator: state.username
    };

    try {
      const response = await axios.get(url, { headers, params });
      commit('setIssues', response.data);
    } catch (error) {
      commit('setError', error.response ? error.response.data : error.message);
    } finally {
      commit('setLoading', false);
    }
  },
  loginWithGithub() {
    console.log("Attempting to log in with GitHub");
    this.$auth.loginWith('github').then(() => {
      console.log("Logged in with GitHub");
      // Assuming the token is set in the store during the login process
      this.$store.dispatch('github/fetchUsername').then(() => {
        // Replace 'repoName' with actual repository name you want to fetch issues from
        this.$store.dispatch('github/fetchIssues', 'repoName');
      });
    });
  }


};

export default {
  namespaced: true,
  state,
  mutations,
  actions,
};
