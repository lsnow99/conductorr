import { createStore } from "vuex";

const store = createStore({
  state() {
    return {
      loggedIn: true,
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
    },
    toasts(state) {
      return state.toasts
    }
  }
});

export default store;
