import { createStore } from "vuex";

const store = createStore({
  state() {
    return {
      loggedIn: false,
    };
  },
  mutations: {
    setLoggedIn(state, loggedIn) {
      state.loggedIn = loggedIn;
    },
  },
  getters: {
    loggedIn (state) {
      return state.loggedIn
    }
  }
});

export default store;
