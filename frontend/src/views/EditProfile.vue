<template>
  <EditProfileInner v-if="profile" :profile="profile" />
</template>

<script setup lang="ts">
import { ReleaseProfile } from "../types/api/release_profile";
import {
  computed,
  nextTick,
  onMounted,
  ref,
  watch,
  WritableComputedRef,
} from "vue";
import { onBeforeRouteLeave, useRoute } from "vue-router";
import APIUtil from "../util/APIUtil";
import EditProfileInner from "../components/EditProfileInner.vue";

const route = useRoute();

const profile = ref<ReleaseProfile | null>(null);
const profileID = ref<number>(parseInt(route.params.profile_id as string));

const loadProfile = async () => {
  try {
    const loadedProfile = await APIUtil.getProfile(profileID.value);
    profile.value = loadedProfile;
  } catch { }
};

onMounted(() => {
  loadProfile();
});
</script>
