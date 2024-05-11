import axios from 'axios';

export const state = () => ({
  issues: []
});

export const mutations = {
  SET_ISSUES(state, issues) {
    state.issues = issues;
  }
};

export const actions = {
  async fetchIssues({ commit }) {
    try {
      const response = await axios.get('https://api.github.com/repos/your-username/your-repo/issues');
      commit('SET_ISSUES', response.data);
    } catch (error) {
      console.error('Error fetching issues from GitHub:', error);
    }
  }
};

export const getters = {
  getIssues(state) {
    return state.issues;
  }
};