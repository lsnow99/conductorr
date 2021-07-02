<template>
  <div v-if="initialized">
    <router-view></router-view>
  </div>
</template>

<script>
import APIUtil from "./util/APIUtil";
import EventBus from "./util/EventBus";

export default {
  data() {
    return {
      initialized: false,
      interval: 0
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
    this.interval = setInterval(() => {
      APIUtil.getStatus().then((status) => {
        this.$store.commit("setStatus", status)
      })
    }, 5000)
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
  created() {
    EventBus.on('notification', (data) => {
      this.$oruga.notification.open(data);
    })
  },
  beforeUnmount() {
    clearInterval(this.interval)
  }
};
</script>

<style>
body,
html {
  @apply dark:text-white;
  @apply dark:bg-gray-800;
}
</style>
