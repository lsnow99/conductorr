<template>
  <section class="mt-3">
    <div
      v-for="(pathConfig, index) in pathConfigs"
      :key="index"
      class="flex flex-col md:flex-row items-center my-6"
    >
      <div class="flex flex-1">
        <o-field addons grouped label="Library Path" class="w-full">
          <o-input
            v-model="pathConfig.path"
            expanded
            placeholder="/library/movies"
          />
          <o-button @click="testPath(pathConfig)"
            ><action-button>Test</action-button></o-button
          >
        </o-field>
      </div>
      <div class="flex flex-row items-center mt-6">
        <o-switch
          @input="moviesDefaultChanged(index)"
          v-model="pathConfig.movies_default"
          class="mx-3"
          >Default for Movies</o-switch
        >
        <o-switch
          @input="seriesDefaultChanged(index)"
          v-model="pathConfig.series_default"
          class="mx-3"
          >Default for TV</o-switch
        >
        <o-button
          @click="deletePath(pathConfig, index)"
          class="mr-3"
          variant="danger"
          icon-right="trash"
        />
      </div>
    </div>
    <div class="flex flex-row justify-between">
      <div />
      <div>
        <o-button class="mr-3" @click="pathConfigs.push({})">New Path</o-button
        ><o-button variant="primary" @click="savePaths">Save Paths</o-button>
      </div>
    </div>
  </section>
</template>

<script>
import ActionButton from "../components/ActionButton.vue";
import APIUtil from "../util/APIUtil";
export default {
  components: { ActionButton },
  data() {
    return {
      pathConfigs: [],
    };
  },
  methods: {
    testPath(pathConfig) {
      APIUtil.testPath(pathConfig.path).then((resp) => {
        if (resp.success) {
          this.$oruga.notification.open({
            duration: 3000,
            message: `Path is OK`,
            position: "bottom-right",
            variant: "success",
            closable: false,
          });
        } else {
          this.$oruga.notification.open({
            duration: 5000,
            message: `Test failed: ${resp.msg}`,
            position: "bottom-right",
            variant: "danger",
            closable: false,
          });
        }
      });
    },
    deletePath(pathConfig, index) {
      if (pathConfig && pathConfig.id) {
        APIUtil.deletePath(pathConfig.id).then(() => {
          this.loadPaths();
        });
      } else {
        this.pathConfigs.splice(index, 1);
      }
    },
    moviesDefaultChanged(index) {
      for (let i = 0; i < this.pathConfigs.length; i++) {
        this.pathConfigs[i].movies_default = false;
      }
        this.pathConfigs[index].movies_default = true;
    },
    seriesDefaultChanged(index) {
      for (let i = 0; i < this.pathConfigs.length; i++) {
        this.pathConfigs[i].series_default = false;
      }
      this.pathConfigs[index].series_default = true;
    },
    loadPaths() {
      APIUtil.getPaths().then((paths) => {
        this.pathConfigs = paths;
      });
    },
    savePaths() {
      APIUtil.updatePaths(this.pathConfigs).then(() => {
        this.loadPaths();
        this.$oruga.notification.open({
          duration: 3000,
          message: `Successfully saved paths`,
          position: "bottom-right",
          variant: "success",
          closable: false,
        });
      });
    },
  },
  mounted() {
    this.loadPaths();
  },
};
</script>
