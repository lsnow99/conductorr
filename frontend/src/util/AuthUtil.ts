import { useAppStore } from "../store";

const logout = () => {
  const appStore = useAppStore()
  appStore.loggedIn = false
};

export default {
  logout
};
