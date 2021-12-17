<template>
  <section class="mt-3">
    <o-button variant="primary" @click="newPath">New Path</o-button>
    <div
      v-for="(pathConfig, index) in pathConfigs"
      :key="index"
      class="flex flex-col md:flex-row items-center my-3"
    >
      <div class="flex flex-1">
        <o-field addons grouped label="Library Path" class="w-full">
          <o-input
            disabled
            v-model="pathConfig.path"
            expanded
            placeholder="/library/movies"
          />
        </o-field>
      </div>
      <div class="flex flex-row items-center mt-6">
        <o-switch disabled v-model="pathConfig.movies_default" class="mx-3"
          >Default for Movies</o-switch
        >
        <o-switch disabled v-model="pathConfig.series_default" class="mx-3"
          >Default for TV</o-switch
        >
        <o-button
          @click="editPath(pathConfig, $event)"
          class="mr-3"
          variant="primary"
          icon-right="edit"
        />
        <o-button
          @click="deletePath(pathConfig)"
          class="mr-3"
          variant="danger"
          icon-right="trash"
        />
      </div>
    </div>
  </section>
  <edit-path
    v-model:active="showNewPathModal"
    @close="closeNewPathModal"
    @submitted="newPathSubmitted"
  />
  <edit-path
    v-model:active="showEditPathModal"
    :path="editingPath"
    @close="closeEditPathModal"
    @submitted="editPathSubmitted"
  />
</template>

<script>
import ActionButton from "../components/ActionButton.vue";
import EditPath from "../components/EditPath.vue";
import APIUtil from "../util/APIUtil";
import TabSaver from "../util/TabSaver";

export default {
  components: { ActionButton, EditPath },
  data() {
    return {
      pathConfigs: [],
      editingPath: {},
      showNewPathModal: false,
      showEditPathModal: false,
    };
  },
  mixins: [TabSaver],
  methods: {
    loadPaths() {
      APIUtil.getPaths().then((paths) => {
        this.pathConfigs = paths;
      });
    },
    deletePath(pathConfig) {
      APIUtil.deletePath(pathConfig.id).then(() => {
        this.loadPaths();
      });
    },
    editPath(pathConfig, $event) {
      this.editingPath = pathConfig;
      this.showEditPathModal = true;
      this.lastButton = $event.currentTarget
    },
    editPathSubmitted(path) {
      APIUtil.updatePath(
        this.editingPath.id,
        path.path,
        path.moviesDefault,
        path.seriesDefault
      )
        .then(() => {
          this.$oruga.notification.open({
            duration: 3000,
            message: `Updated path ${path.path}`,
            position: "bottom-right",
            variant: "success",
            closable: false,
          });
        })
        .finally(() => {
          this.showEditPathModal = false;
          this.loadPaths();
        });
    },
    closeNewPathModal() {
      this.showNewPathModal = false;
      this.restoreFocus();
    },
    closeEditPathModal() {
      this.showEditPathModal = false;
      this.restoreFocus();
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
  },
  mounted() {
    this.loadPaths();
  },
};
</script>
