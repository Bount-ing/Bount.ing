import { createStore } from 'vuex'

export default createStore({
  state: {
    bounties: []
  },
  mutations: {
    SET_BOUNTIES(state, bounties) {
      state.bounties = bounties;
    }
  },
  actions: {
    fetchBounties({ commit }) {
      // Here you would fetch the bounties from the backend
      // This is just a placeholder
      const bounties = [
        { id: 1, title: "Fix bug in repo", amount: 100 },
        { id: 2, title: "Add new feature", amount: 150 }
      ];
      commit('SET_BOUNTIES', bounties);
    }
  },
  modules: {}
})

