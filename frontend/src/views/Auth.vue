<template>
  <div class="flex flex-row justify-center">
    <section class="max-w-lg mt-40">
      <img src="/logo-rect-transparent.svg" width="100" class="pb-16 m-auto" />
      <div v-if="firstTime" class="text-2xl">Create User</div>
      <div v-else class="text-2xl">Login</div>
      <o-field label="Username">
        <o-input type="text" @keydown.enter="submit" v-model="username" />
      </o-field>
      <o-field label="Password">
        <o-input type="password" @keydown.enter="submit" v-model="password" />
      </o-field>
      <o-field v-if="firstTime" label="Confirm Password">
        <o-input
          type="password"
          @keydown.enter="submit"
          v-model="confirmPassword"
        />
      </o-field>
      <div class="flex flex-row justify-between mt-2">
        <div />
        <o-button variant="primary" v-if="firstTime" @click="submit"
          >Register</o-button
        >
        <o-button variant="primary" v-else @click="submit">Login</o-button>
      </div>
    </section>
    <o-loading
      :full-page="true"
      v-model:active="loading"
      :can-cancel="false"
    ></o-loading>
  </div>
</template>

<script setup lang="ts">
import { inject, onMounted, ref } from "vue";
import APIUtil from "../util/APIUtil";
import { useRouter } from "vue-router";
import { storeToRefs } from "pinia";
import { useAppStore } from "@/store";

const username = ref("");
const password = ref("");
const confirmPassword = ref("");
const loading = ref(true);
const firstTime = ref(false);

const router = useRouter();
const oruga = inject("oruga");

const { loggedIn } = storeToRefs(useAppStore())

const submitLogin = async () => {
  loading.value = true;
  try {
    await APIUtil.signIn(username.value, password.value);
    loggedIn.value = true
    router.push({ name: "home" });
  } catch (err) {
    oruga.notification.open({
      duration: 5000,
      message: `Error logging in: ${err}`,
      position: "bottom-right",
      variant: "danger",
      closable: false,
    });
    loggedIn.value = false
  } finally {
    loading.value = false;
  }
};

const submitRegister = async () => {
  if (confirmPassword.value != password.value) {
    oruga.notification.open({
      duration: 5000,
      message: `Passwords do not match`,
      position: "bottom-right",
      variant: "danger",
      closable: false,
    });
    return;
  }
  loading.value = true;
  try {
    await APIUtil.signUp(username.value, password.value);
    loggedIn.value = true
    router.push({ name: "home" });
  } catch (err) {
    oruga.notification.open({
      duration: 5000,
      message: `Error signing up: ${err}`,
      position: "bottom-right",
      variant: "danger",
      closable: false,
    });
  } finally {
    loading.value = false;
  }
};

const submit = async () => {
  if (firstTime.value) {
    submitRegister();
  } else {
    submitLogin();
  }
};

onMounted(async () => {
  if (loggedIn.value) {
    router.push({ name: "home" });
  } else {
    try {
      await APIUtil.isFirstTime();
      firstTime.value = true;
    } catch (err) {
      firstTime.value = false;
    } finally {
      loading.value = false;
    }
  }
});
</script>
