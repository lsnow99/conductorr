<template>
  <header class="modal-card-header">
    <p class="modal-card-title">New Downloader</p>
  </header>
  <section class="modal-card-content">
    <service-options :services="downloaderTypes" v-model="selectedDownloader" />
  </section>
  <footer class="modal-card-footer">
    <o-button @click="$emit('close')">Cancel</o-button>
    <div>
      <o-button variant="primary" :disabled="selectedDownloader == ''" @click="next"
        >Next</o-button
      >
    </div>
  </footer>
</template>

<script>
import ServiceOptions from "./ServiceOptions.vue"

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
      selectedDownloader: ''
    };
  },
  components: {
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
