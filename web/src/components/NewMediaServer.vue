<template>
  <header class="modal-card-header">
    <p class="modal-card-title">New Media Server</p>
  </header>
  <section class="modal-card-content">
    <service-options :services="mediaServerTypes" v-model="selectedMediaServer" v-if="!selectedMediaServer" />
    <div v-else>
        <div v-if="selectedMediaServer == 'transmission'">
            <edit-plex />
        </div>
    </div>
  </section>
  <footer class="modal-card-footer">
    <o-button @click="$emit('close')">Cancel</o-button>
    <div>
      <o-button variant="primary" :disabled="selectedMediaServer == ''" @click="next"
        >Next</o-button
      >
    </div>
  </footer>
</template>

<script>
import EditPlex from "./EditPlex.vue"
import ServiceOptions from "./ServiceOptions.vue"

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
      selectedMediaServer: '',
      showNext: false
    };
  },
  components: {
      EditPlex,
      ServiceOptions
  },
  emits: ['close', 'selected'],
  methods: {
    next() {
      this.$emit('selected', this.selectedDownloader)
    },
  },
};
</script>
