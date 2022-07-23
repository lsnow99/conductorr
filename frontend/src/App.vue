<template>
  <div v-if="initialized">
    <router-view></router-view>
  </div>
</template>

<script setup lang="ts">
import { inject, onBeforeUnmount, onMounted, ref, watch } from "vue";
import { useRouter } from "vue-router";
import APIUtil from "./util/APIUtil";
import EventBus from "./util/EventBus";
import { useAppStore } from "@/store"
import { storeToRefs } from "pinia";

const initialized = ref(false)
const interval = ref(0)

const router = useRouter();
const oruga = inject("oruga")

const { loggedIn, status } = storeToRefs(useAppStore())

const getStatus = async() => {
  try {
    status.value = await APIUtil.getStatus()
  } catch {}
}

onMounted(async() => {
  try {
    loggedIn.value = await APIUtil.checkAuth()
    if (!loggedIn.value) {
      router.push({name: "auth"})
    }
  } catch(err) {
    loggedIn.value = false
    console.log("error communicating with backend:", err)
    router.push({name: "auth"})
  }

  initialized.value = true

  getStatus()

  localStorage.theme = "dark";
  if (localStorage.theme === "dark" || (!("theme" in localStorage) && window.matchMedia("(prefers-color-scheme: dark)").matches)) {
    document.documentElement.classList.add("dark")
  } else {
    document.documentElement.classList.remove("dark");
  }
})

watch(() => loggedIn, (newVal) => {
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
