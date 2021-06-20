import { createStore } from "vuex";

const store = createStore({
  state() {
    return {
      loggedIn: false,
      toasts: [],
    };
  },
  mutations: {
    setLoggedIn(state, loggedIn) {
      state.loggedIn = loggedIn;
    },
    addToast(state, toast) {
      state.toasts.push(toast) - 1
      setTimeout(() => {
        const index = state.toasts.indexOf(toast)
        state.toasts.splice(index, 1)
      }, 4000)
    }
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
