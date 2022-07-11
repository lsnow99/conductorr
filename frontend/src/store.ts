import { defineStore } from "pinia";

const useAppStore = defineStore('app', {
  state: () => {
    return {
      loggedIn: true,
      status: {}
    };
  },
});

export {
  useAppStore
};
