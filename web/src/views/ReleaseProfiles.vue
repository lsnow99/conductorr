<template>
  <section>
    <div class="flex flex-row justify-between mt-3">
      <div>
        <o-button variant="primary" icon-left="plus" @click="showNewProfileModal = true">New Profile</o-button>
      </div>
      <div>
        <o-button variant="primary" icon-left="plus-square" @click="expandAll" class="mr-3"
          >Expand All</o-button
        ><o-button variant="primary" icon-left="minus-square" @click="collapseAll"
          >Collapse All</o-button
        >
      </div>
    </div>
    <release-profile
      v-for="(profile, index) in profiles"
      :ripTypes="ripTypes"
      :qualityTypes="qualityTypes"
      :resolutionTypes="resolutionTypes"
      v-model="profiles[index]"
      v-model:expanded="profile.expanded"
      @reload="loadProfiles"
      :key="index"
    />
    <o-modal v-model:active="showNewProfileModal">
      <new-profile @close="showNewProfileModal = false" @submitted="newProfileSubmitted" />
    </o-modal>
  </section>
</template>

<script>
import APIUtil from "../util/APIUtil";
import ReleaseProfile from "../components/ReleaseProfile.vue";
import NewProfile from "../components/NewProfile.vue";

export default {
  data() {
    return {
      ripTypes: [],
      qualityTypes: [],
      resolutionTypes: [],
      showNewProfileModal: false,
      profiles: []
    };
  },
  components: {
    ReleaseProfile,
    NewProfile,
  },
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
      this.loadProfiles()
    },
    loadProfiles() {
      APIUtil.getProfiles().then((data) => {
        this.profiles = data
      })
    }
  },
  mounted() {
    this.loadProfiles()

    APIUtil.getReleaseProfileCfg().then((data) => {
      this.ripTypes = data.rip_types;
      this.qualityTypes = data.quality_types;
      this.resolutionTypes = data.resolutionTypes;
    });
  },
};
</script>
