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
      interval: 0,
    };
  },
  methods: {
    getStatus() {
      if (this.loggedIn) {
        APIUtil.getStatus().then((status) => {
          this.$store.commit("setStatus", status);
        });
      }
    },
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
    this.getStatus();
    this.interval = setInterval(() => {
      this.getStatus();
    }, 5000);
    // Whenever the user explicitly chooses dark mode
    localStorage.theme = "dark";
    if (
      localStorage.theme === "dark" ||
      (!("theme" in localStorage) &&
        window.matchMedia("(prefers-color-scheme: dark)").matches)
    ) {
      document.documentElement.classList.add("dark");
    } else {
      document.documentElement.classList.remove("dark");
    }
  },
  computed: {
    loggedIn() {
      return this.$store.getters.loggedIn;
    },
  },
  watch: {
    loggedIn: function (newVal) {
      if (!newVal) {
        this.$router.push({ name: "home" });
      }
    },
  },
  created() {
    EventBus.on("notification", (data) => {
      this.$oruga.notification.open(data);
    });
    EventBus.on("forceLogout", () => {
      this.$router.push({name: "auth"})
    });
  },
  beforeUnmount() {
    clearInterval(this.interval);
  },
};
</script>

<style>
body,
html {
  @apply dark:text-white;
  @apply dark:bg-darkgray;
}
</style>
