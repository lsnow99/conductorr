import store from "../store";

const logout = () => {
  store.commit('setLoggedIn', false);
};

export default {
  logout
};
