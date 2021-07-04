<template>
  <header class="modal-card-header">
    <p class="modal-card-title">{{ media.title }}</p>
  </header>
  <section class="modal-card-content w-96 min-w-full">
    <o-field label="Profile">
      <o-select v-model="profileID" placeholder="Profile">
        <option
          v-for="profileOption in profiles"
          :key="profileOption.id"
          :value="profileOption.id"
        >
          {{ profileOption.name }}
        </option>
      </o-select>
    </o-field>
    <o-field label="Root Path">
      <o-select v-model="pathID" placeholder="Path">
        <option
          v-for="pathOption in paths"
          :key="pathOption.id"
          :value="pathOption.id"
        >
          {{ pathOption.path }}
        </option>
      </o-select>
    </o-field>
  </section>
  <footer class="modal-card-footer">
    <o-button @click="$emit('close')">Cancel</o-button>
    <div>
      <o-button variant="primary" @click="save">Save</o-button>
    </div>
  </footer>
</template>

<script>
import APIUtil from "../util/APIUtil";
export default {
  data() {
    return {
      profiles: [],
      paths: [],
      profileID: 0,
      pathID: 0,
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
  emits: ["close", "submit"],
  methods: {
    save() {
      this.$emit("submit", this.profileID, this.pathID);
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
  },
  mounted() {
    this.loadProfiles();
    this.loadPaths();
  },
};
</script>
