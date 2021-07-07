<template>
  <section class="mt-3">
    <o-button class="mr-3" variant="primary" @click="newPath">New Path</o-button>
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
          @click="editPath(pathConfig)"
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
  <o-modal v-model:active="showNewPathModal">
    <edit-path
      @close="showNewPathModal = false"
      @submitted="newPathSubmitted"
    />
  </o-modal>
  <o-modal v-model:active="showEditPathModal">
    <edit-path
      :path="editingPath"
      @close="showEditPathModal = false"
      @submitted="editPathSubmitted"
    />
  </o-modal>
</template>

<script>
import ActionButton from "../components/ActionButton.vue";
import EditPath from "../components/EditPath.vue";
import APIUtil from "../util/APIUtil";

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
    editPath(pathConfig) {
      this.editingPath = pathConfig;
      this.showEditPathModal = true;
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
    newPath() {
      this.showNewPathModal = true;
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
