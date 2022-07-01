<template>
  <div v-if="initialized">
    <router-view></router-view>
  </div>
</template>

<script setup lang="ts">
import { computed, inject, onBeforeUnmount, onMounted, ref, watch } from "vue";
import { useRouter } from "vue-router";
import store from "./store";
import APIUtil from "./util/APIUtil";
import EventBus from "./util/EventBus";

const initialized = ref(false)
const interval = ref(0)

const router = useRouter();
const oruga = inject("oruga")

const loggedIn = computed(() => {
  return store.getters.loggedIn
})

const getStatus = async() => {
  if (loggedIn) {
    const status = await APIUtil.getStatus()
    store.commit("setStatus", status)
  }
}

onMounted(async() => {
  const isLoggedIn = await APIUtil.checkAuth()
  store.commit("setLoggedIn", isLoggedIn)
  if (!isLoggedIn) {
    router.push({name: "auth"})
  }

  getStatus()

  localStorage.theme = "dark";
  if (localStorage.theme === "dark" || (!("theme" in localStorage) && window.matchMedia("(prefers-color-scheme: dark)").matches)) {
    document.documentElement.classList.add("dark")
  } else {
      document.documentElement.classList.remove("dark");
  }
})

watch(loggedIn, (newVal) => {
  if (!newVal) {
    router.push({ name: "home" });
  }
})

EventBus.on("notification", (data) => {
  oruga.notification.open(data);
});
EventBus.on("forceLogout", () => {
  router.push({name: "auth"})
});

onBeforeUnmount(() => {
  clearInterval(interval.value)
})
</script>

<style>
body,
html {
  @apply dark:text-white;
  @apply dark:bg-darkgray;
}
</style>
