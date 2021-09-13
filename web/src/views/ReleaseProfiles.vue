<template>
  <section class="mt-3">
    <div class="flex flex-col sm:flex-row justify-between">
      <div class="flex justify-center">
        <o-button variant="primary" @click="newProfile">New Profile</o-button>
      </div>
      <div class="flex justify-center mt-4 sm:mt-0">
        <o-button
          variant="primary"
          icon-left="plus-square"
          @click="expandAll"
          class="mr-3"
          >Expand All</o-button
        ><o-button
          variant="primary"
          icon-left="minus-square"
          @click="collapseAll"
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
    <!-- <release-profile
      v-for="(profile, index) in profiles"
      v-model="profiles[index]"
      v-model:expanded="profile.expanded"
      @reload="loadProfiles"
      :key="index"
    /> -->
    <new-profile
      v-model:active="showNewProfileModal"
      @close="closeNewProfileModal"
      @submitted="newProfileSubmitted"
    />
  </section>
</template>

<script>
import APIUtil from "../util/APIUtil";
import NewProfile from "../components/NewProfile.vue";
import ConfigItem from "../components/ConfigItem.vue";
import CSLEditor from "../components/CSLEditor.vue";
import TabSaver from "../util/TabSaver";

export default {
  data() {
    return {
      ripTypes: [],
      qualityTypes: [],
      resolutionTypes: [],
      showNewProfileModal: false,
      profiles: [],
    };
  },
  components: {
    NewProfile,
    ConfigItem,
    CSLEditor,
  },
  mixins: [TabSaver],
  methods: {
    expandAll() {
      this.profiles.forEach((element) => {
        element.expanded = true;
      });
    },
    collapseAll() {
      this.profiles.forEach((element) => {
        element.expanded = false;
      });
    },
    newProfileSubmitted() {
      this.showNewProfileModal = false;
      this.loadProfiles();
    },
    loadProfiles() {
      APIUtil.getProfiles().then((data) => {
        this.profiles = data;
      });
    },
    edit(profile) {
      this.$router.push({
        name: "editProfile",
        params: { profile_id: profile.id },
      });
    },
    deleteProfile(profile) {
      this.loadingDelete = true;
      APIUtil.deleteProfile(profile.id)
        .then(() => {
          this.$oruga.notification.open({
            duration: 3000,
            message: `Deleted profile ${profile.name}`,
            position: "bottom-right",
            variant: "success",
            closable: false,
          });
          this.loadProfiles();
        })
        .finally(() => {
          this.loadingDelete = false;
        });
    },
    newProfile($event) {
      this.lastButton = $event.currentTarget;
      this.showNewProfileModal = true;
    },
    closeNewProfileModal() {
      this.showNewProfileModal = false;
      this.restoreFocus();
    },
  },
  mounted() {
    this.loadProfiles();

    APIUtil.getReleaseProfileCfg().then((data) => {
      this.ripTypes = data.rip_types;
      this.qualityTypes = data.quality_types;
      this.resolutionTypes = data.resolutionTypes;
    });
  },
};
</script>
