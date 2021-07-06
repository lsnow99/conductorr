<template>
  <header class="modal-card-header">
    <p class="modal-card-title">{{ media.title }}</p>
  </header>
  <section class="modal-card-content w-96 min-w-full">
    <o-field label="Profile" :variant="profileVariant" :message="profileVariant?'A profile is required':''">
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
    <o-field label="Root Path" :variant="pathVariant" :message="pathVariant?'A root path is required':''">
      <o-select expanded v-model="pathID" :placeholder="`/media/library/${media.content_type}`">
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
  </section>
  <footer class="modal-card-footer">
    <o-button @click="$emit('close')">Cancel</o-button>
    <div>
      <o-button variant="primary" @click="save">Save</o-button>
    </div>
  </footer>
  <o-modal v-model:active="showNewProfileModal">
    <new-profile @close="showNewProfileModal = false" @submitted="newProfileSubmitted" />
  </o-modal>
  <o-modal v-model:active="showNewPathModal">
    <edit-path @close="showNewPathModal = false" @submitted="newPathSubmitted" />
  </o-modal>
</template>

<script>
import APIUtil from "../util/APIUtil";
import NewProfile from "../components/NewProfile.vue"
import EditPath from "../components/EditPath.vue"

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
  },
  components: {NewProfile, EditPath},
  emits: ["close", "submit"],
  methods: {
    save() {
      this.submittedOnce = true
      if(this.profileID && this.pathID) {
        this.$emit("submit", {profileID: this.profileID, pathID: this.pathID});
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
        if (!this.media.path_id) {
          this.paths.forEach((elem) => {
            if (elem.movies_default && this.media.content_type == "movie") {
              this.pathID = elem.id;
            } else if (
              elem.series_default &&
              this.media.content_type == "series"
            ) {
              this.pathID = elem.id;
            }
          });
        }
      });
    },
    newPath() {
      this.showNewPathModal = true
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
    newProfile() {
      this.showNewProfileModal = true
    },
    newProfileSubmitted() {
      this.loadProfiles()
      this.showNewProfileModal = false
    }
  },
  mounted() {
    this.loadProfiles();
    this.loadPaths();
  },
  computed: {
    profileVariant() {
      if (!this.submittedOnce) {
        return ''
      }
      if (!this.profileID) {
        return 'danger'
      }
    },
    pathVariant() {
      if (!this.submittedOnce) {
        return ''
      }
      if (!this.pathID) {
        return 'danger'
      }
    }
  }
};
</script>
