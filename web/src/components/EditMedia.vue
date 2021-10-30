<template>
  <modal :title="media.title" v-model="computedActive" @close="$emit('close')">
    <o-field
      label="Profile"
      :variant="profileVariant"
      :message="profileVariant ? 'A profile is required' : ''"
    >
      <o-select expanded v-model="profileID" placeholder="Profile">
        <option
          v-for="profileOption in profiles"
          :key="profileOption.id"
          :value="profileOption.id"
        >
          {{ profileOption.name }}
        </option>
      </o-select>
      <o-button @click="newProfile">New</o-button>
    </o-field>
    <o-field
      label="Root Path"
      :variant="pathVariant"
      :message="pathVariant ? 'A root path is required' : ''"
    >
      <o-select
        expanded
        v-model="pathID"
        :placeholder="`/media/library/${media.content_type}`"
      >
        <option
          v-for="pathOption in paths"
          :key="pathOption.id"
          :value="pathOption.id"
        >
          {{ pathOption.path }}
        </option>
      </o-select>
      <o-button @click="newPath">New</o-button>
    </o-field>
    <template v-slot:footer>
      <o-button @click="$emit('close')">Cancel</o-button>
      <div>
        <o-button variant="primary" @click="save">
          <action-button :mode="loading ? 'loading' : ''">Save</action-button>
        </o-button>
      </div>
    </template>
  </modal>
  <new-profile
    v-model:active="showNewProfileModal"
    @close="closeNewProfile"
    @submitted="newProfileSubmitted"
  />
  <edit-path
    v-model:active="showNewPathModal"
    @close="closeNewPath"
    @submitted="newPathSubmitted"
  />
</template>

<script>
import APIUtil from "../util/APIUtil";
import NewProfile from "../components/NewProfile.vue";
import EditPath from "../components/EditPath.vue";
import ActionButton from "../components/ActionButton.vue";
import Modal from "../components/Modal.vue";
import TabSaver from "../util/TabSaver";

export default {
  data() {
    return {
      profiles: [],
      paths: [],
      profileID: null,
      pathID: null,
      showNewProfileModal: false,
      showNewPathModal: false,
      submittedOnce: false,
    };
  },
  props: {
    media: {
      type: Object,
      default: function () {
        return {};
      },
    },
    loading: {
      type: Boolean,
      default: function () {
        return false;
      },
    },
    active: {
      type: Boolean,
      default: function () {
        return false;
      },
    },
  },
  components: { NewProfile, EditPath, ActionButton, Modal },
  emits: ["close", "submit", "update:active"],
  mixins: [TabSaver],
  methods: {
    save() {
      this.submittedOnce = true;
      if (this.profileID && this.pathID) {
        this.$emit("submit", {
          profileID: this.profileID,
          pathID: this.pathID,
        });
      }
    },
    loadProfiles() {
      APIUtil.getProfiles().then((profiles) => {
        this.profiles = profiles;
      });
    },
    loadPaths() {
      APIUtil.getPaths().then((paths) => {
        this.paths = paths;
      });
    },
    newPath($event) {
      this.showNewPathModal = true;
      this.lastButton = $event.currentTarget;
    },
    newPathSubmitted(path) {
      APIUtil.createNewPath(path.path, path.moviesDefault, path.seriesDefault)
        .then(() => {
          this.$oruga.notification.open({
            duration: 3000,
            message: `Created path ${path.path}`,
            position: "bottom-right",
            variant: "success",
            closable: false,
          });
        })
        .finally(() => {
          this.showNewPathModal = false;
          this.loadPaths();
        });
    },
    newProfile($event) {
      this.lastButton = $event.currentTarget;
      this.showNewProfileModal = true;
    },
    closeNewPath() {
      this.showNewPathModal = false;
      this.restoreFocus();
    },
    closeNewProfile() {
      this.showNewProfileModal = false;
      this.restoreFocus();
    },
    newProfileSubmitted() {
      this.loadProfiles();
      this.showNewProfileModal = false;
      this.restoreFocus();
    },
  },
  mounted() {
    this.loadProfiles();
    this.loadPaths();
  },
  watch: {
    media: {
      handler: function (newVal) {
        if (newVal) {
          this.profileID = newVal.profile_id;
          if (!this.media.path_id) {
            this.paths.forEach((elem) => {
              console.log('checking ', elem, this.media)
              if (
                (elem.movies_default && this.media.content_type == "movie") ||
                (elem.series_default && this.media.content_type == "series")
              ) {
                this.pathID = elem.id;
              }
            });
          } else {
            this.pathID = this.media.path_id;
          }
        }
      },
      immediate: true,
    },
  },
  computed: {
    computedActive: {
      get() {
        return this.active;
      },
      set(v) {
        this.$emit("update:active", v);
      },
    },
    profileVariant() {
      if (!this.submittedOnce) {
        return "";
      }
      if (!this.profileID) {
        return "danger";
      }
    },
    pathVariant() {
      if (!this.submittedOnce) {
        return "";
      }
      if (!this.pathID) {
        return "danger";
      }
    },
  },
};
</script>
