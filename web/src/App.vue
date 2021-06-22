<template>
  <div v-if="initialized">
    <router-view></router-view>
  </div>
</template>

<script>
import APIUtil from "./util/APIUtil";

export default {
  data() {
    return {
      initialized: false,
    }
  },
  mounted() {
    APIUtil.checkAuth().then((isLoggedIn) => {
      this.$store.commit("setLoggedIn", isLoggedIn);
      if (isLoggedIn) {
        // this.$router.push({ name: "home" });
      } else {
        this.$router.push({ name: "auth" });
      }
      this.initialized = true;
    });
  },
  computed: {
    loggedIn() {
      return this.$store.getters.loggedIn;
    },
  },
  watch: {
    loggedIn: function (newVal) {
      if (!newVal) {
        this.$router.push({ name: "logout" });
      }
    },
  },
};
</script>

<style>
body,
html {
  @apply dark:text-white;
  @apply dark:bg-gray-800;
}
</style>
