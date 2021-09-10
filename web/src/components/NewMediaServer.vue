<template>
  <modal
    title="New Media Server"
    v-model="computedActive"
    @close="$emit('close')"
  >
    <service-options
      :services="mediaServerTypes"
      v-model="selectedMediaServer"
    />
    <template v-slot:footer>
      <o-button @click="$emit('close')">Cancel</o-button>
      <div>
        <o-button
          variant="primary"
          :disabled="selectedMediaServer == ''"
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

const mediaServerTypes = [
  {
    name: "Plex",
    value: "plex",
  },
];

export default {
  data() {
    return {
      mediaServerTypes,
      selectedMediaServer: "",
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
    Modal,
  },
  emits: ["close", "selected", "update:active"],
  methods: {
    next() {
      this.$emit("selected", this.selectedMediaServer);
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
