<template>
  <section class="mt-3">
    <div class="flex flex-col sm:flex-row justify-between">
      <div class="flex justify-center">
        <o-button variant="primary" @click="openNewProfileModal($event)"
          >New Profile</o-button
        >
      </div>
      <div class="flex justify-center mt-4 sm:mt-0">
        <o-button
          variant="primary"
          icon-left="plus-square"
          @click="setExpanded(true)"
          class="mr-3"
          >Expand All</o-button
        ><o-button
          variant="primary"
          icon-left="minus-square"
          @click="setExpanded(false)"
          >Collapse All</o-button
        >
      </div>
    </div>
    <config-item
      :delete-message="`Are you sure you want to delete profile '${profile.name}'`"
      collapsible
      v-model:expanded="profile.expanded"
      v-for="profile in profiles"
      :key="profile.id"
      @delete="deleteProfile(profile)"
      @edit="edit(profile)"
      :title="profile.name"
    >
      <div class="p-4">
        Filter
        <CSLEditor readonly v-model="profile.filter" />
        Sorter
        <CSLEditor readonly v-model="profile.sorter" />
      </div>
    </config-item>
    <new-profile
      v-model:active="showNewProfileModal"
      @close="closeNewProfileModal"
      @submitted="newProfileSubmitted"
    />
  </section>
</template>

<script setup lang="ts">
import APIUtil from "../util/APIUtil";
import NewProfile from "../components/NewProfile.vue";
import ConfigItem from "../components/ConfigItem.vue";
import { CSLEditor } from "conductorr-lib";
import TabSaver from "../util/TabSaver";
import { ReleaseProfile } from "@/types/api/release_profile";
import { inject, onMounted, ref } from "vue";
import { useTabSaver } from "@/util";
import { useRouter } from "vue-router";

type ExpandableReleaseProfile = ReleaseProfile & {
  expanded?: boolean;
};

const showNewProfileModal = ref(false);
const profiles = ref<ExpandableReleaseProfile[]>([]);
const loadingDelete = ref(false);

const oruga = inject("oruga");
const { lastButton, restoreFocus } = useTabSaver();
const router = useRouter();

const setExpanded = (expanded: boolean | undefined) =>
  profiles.value.forEach((profile) => (profile.expanded = expanded));

const loadProfiles = async () => {
  const data = await APIUtil.getProfiles();
  profiles.value = data;
};

const newProfileSubmitted = () => {
  showNewProfileModal.value = false;
  loadProfiles();
};

const edit = (profile: ReleaseProfile) =>
  router.push({
    name: "editProfile",
    params: { profile_id: profile.id },
  });

const deleteProfile = async (profile: ReleaseProfile) => {
  loadingDelete.value = true;
  try {
    void (await APIUtil.deleteProfile(profile.id));
    oruga.notification.open({
      duration: 3000,
      message: `Deleted profile ${profile.name}`,
      position: "bottom-right",
      variant: "success",
      closable: false,
    });
  } catch {
  } finally {
    loadingDelete.value = false;
  }
};

const openNewProfileModal = ($event: Event) => {
  lastButton.value = $event.currentTarget as HTMLElement;
  showNewProfileModal.value = true;
};

const closeNewProfileModal = () => {
  showNewProfileModal.value = false;
  restoreFocus();
};

onMounted(loadProfiles);
</script>
