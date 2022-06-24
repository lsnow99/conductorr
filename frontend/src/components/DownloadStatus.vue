<template>
  <div class="relative p-2 my-1 overflow-hidden bg-gray-600 rounded-md">
    <div
      class="absolute top-0 bottom-0 left-0 z-10 opacity-30"
      :class="background"
      :style="`width: ${Math.round(fraction * 100)}%`"
    ></div>
    <span
      v-if="deletable"
      role="button"
      tabindex="0"
      :aria-label="`delete download ${truncatedFriendlyName}`"
      @click="te"
      @keydown.enter="$emit('delete', $event)"
      @keydown.space="$emit('delete', $event)"
      class="absolute top-0 bottom-0 right-0 z-20 flex flex-row justify-center w-8 cursor-pointer hover:bg-red-600 focus:bg-red-600 opacity-80"
      >{{ fraction < 1 ? Math.round(fraction * 100) + "%" : "" }}
      <o-icon icon="times"></o-icon
    ></span>
    <span class="font-bold">{{ statusText }}</span>
    {{ truncatedFriendlyName }}
  </div>
</template>

<script>
export default {
  props: {
    download: {
      type: Object,
      default: function () {
        return {};
      },
    },
  },
  emits: ["fda"],
  methods: {
    te($event) {
      console.log(this.$emit('fda'))
      console.log('wtf')
    }
  },
  computed: {
    truncatedFriendlyName() {
      if (this.download.friendly_name.length > 50) {
        return this.download.friendly_name.substring(0, 50) + "...";
      }
      return this.download.friendly_name;
    },
    deletable() {
      switch (this.download.status) {
        case "waiting":
        case "paused":
        case "downloading":
        case "processing":
          return true;
        default:
          return false;
      }
    },
    background() {
      switch (this.download.status) {
        case "waiting":
          return "bg-purple-600";
        case "paused":
          return "bg-yellow-400";
        case "downloading":
          return "bg-blue-400";
        case "processing":
        case "cprocessing":
          return "bg-gray-900";
        case "done":
          return "bg-green-600";
        case "error":
        case "cerror":
          return "bg-red-700";
      }
    },
    statusText() {
      switch (this.download.status) {
        case "waiting":
          return "QUEUED";
        case "paused":
          return "PAUSED";
        case "downloading":
          return "DOWNLOADING";
        case "processing":
        case "cprocessing":
          return "PROCESSING";
        case "done":
          return "DONE";
        case "error":
        case "cerror":
          return "ERROR";
      }
    },
    fraction() {
      let fraction = 1;
      if (this.download.full_size && this.download.bytes_left) {
        fraction =
          (this.download.full_size - this.download.bytes_left) /
          this.download.full_size;
      }
      fraction = Math.round(fraction * 100) / 100;
      return fraction;
    },
  },
};
</script>
