<template>
  <modal
    title="New Downloader"
    v-model="computedActive"
    @close="$emit('close')"
  >
    <service-options :services="downloaderTypes" v-model="selectedDownloader" />
    <template v-slot:footer>
      <o-button @click="$emit('close')">Cancel</o-button>
      <div>
        <o-button
          variant="primary"
          :disabled="selectedDownloader == ''"
          @click="next"
          >Next</o-button
        >
      </div>
    </template>
  </modal>
</template>

<script>
import ServiceOptions from "./ServiceOptions.vue";
import Modal from "./Modal.vue";

const downloaderTypes = [
  {
    name: "Transmission",
    value: "transmission",
  },
  {
    name: "NZBGet",
    value: "nzbget",
  },
];

export default {
  data() {
    return {
      downloaderTypes,
      selectedDownloader: "",
    };
  },
  props: {
    active: {
      type: Boolean,
      default: function () {
        return false;
      },
    },
  },
  components: {
    ServiceOptions,
    Modal
  },
  emits: ["close", "selected", "update:active"],
  methods: {
    next() {
      this.$emit("selected", this.selectedDownloader);
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
  },
};
</script>
