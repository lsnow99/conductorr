<template>
  <modal
    :title="path ? 'Editing Path' : 'New Path'"
    v-model="computedActive"
    @close="$emit('close')"
  >
    <o-field label="Path">
      <o-input
        expanded
        @keydown.enter="submit"
        placeholder="/media/library/name"
        v-model="pathDir"
      />
      <o-button @click="testPath">Test</o-button>
    </o-field>
    <o-field label="Defaults">
      <div class="flex flex-col">
        <o-switch v-model="moviesDefault">Default for Movies</o-switch>
        <o-switch v-model="seriesDefault">Default for Series</o-switch>
      </div>
    </o-field>
    <template v-slot:footer>
      <o-button @click="$emit('close')">Cancel</o-button>
      <div>
        <o-button variant="primary" @click="submit"
          ><action-button :mode="loading ? 'loading' : ''"
            >Submit</action-button
          ></o-button
        >
      </div>
    </template>
  </modal>
</template>

<script>
import APIUtil from "../util/APIUtil";
import ActionButton from "./ActionButton.vue";
import Modal from "./Modal.vue";

export default {
  components: { ActionButton, Modal },
  data() {
    return {
      pathDir: "",
      moviesDefault: false,
      seriesDefault: false,
      loading: false,
    };
  },
  props: {
    path: {
      type: Object,
      default: function () {
        return null;
      },
    },
    active: {
      type: Boolean,
      default: function () {
        return false;
      },
    },
  },
  emits: ["close", "submitted", "update:active"],
  methods: {
    submit() {
      this.$emit("submitted", {
        path: this.pathDir,
        moviesDefault: this.moviesDefault,
        seriesDefault: this.seriesDefault,
      });
    },
    testPath() {
      this.loading = true;
      APIUtil.testPath(this.pathDir).then((resp) => {
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
      }).finally(() => {
        this.loading = false
      });
    },
  },
  mounted() {
    if (this.path) {
      this.pathDir = this.path.path;
      this.moviesDefault = this.path.movies_default;
      this.seriesDefault = this.path.series_default;
    }
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
  },
};
</script>
