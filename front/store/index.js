import Vuex from 'vuex';
import github from './github';

export const createStore = () => {
  return new Vuex.Store({
    modules: {
      github
    }
  });
};