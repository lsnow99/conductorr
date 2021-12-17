import { createStore } from "vuex";

const store = createStore({
  state() {
    return {
      loggedIn: true,
      status: {}
    };
  },
  mutations: {
    setLoggedIn(state, loggedIn) {
      state.loggedIn = loggedIn;
    },
    setStatus(state, status) {
      state.status = status
    }
  },
  getters: {
    loggedIn (state) {
      return state.loggedIn
    },
    status (state) {
      return state.status
    }
  }
});

export default store;
