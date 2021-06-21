<template>
  <header class="modal-card-header">
    <p class="modal-card-title">New Downloader</p>
  </header>
  <section class="modal-card-content">
    <div v-if="!showNext" class="grid md:grid-cols-2">
      <div
        v-for="downloader in downloaderTypes"
        :key="downloader.value"
        class="downloader"
        :class="selectedDownloader == downloader.value ? 'active' : ''"
        @click="selectedDownloader = downloader.value"
      >
        {{ downloader.name }}
      </div>
    </div>
    <div v-else>
        <div v-if="selectedDownloader == 'transmission'">
            <edit-transmission />
        </div>
    </div>
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

<style scoped>
.downloader {
  @apply border-dashed border-4 border-gray-300 py-8 px-16 m-2 text-2xl cursor-pointer hover:opacity-50 flex justify-center font-bold select-none;
}

.downloader.active {
  @apply border-green-500 text-green-500;
}
</style>

<script>
import EditTransmission from "./EditTransmission.vue"

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
  props: {
    active: {
      type: Boolean,
      default: function () {
        return false;
      },
    },
  },
  components: {
      EditTransmission
  },
  emits: ['close', 'selected'],
  methods: {
    next() {
      this.$emit('selected', this.selectedDownloader)
    },
  },
};
</script>
