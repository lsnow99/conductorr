<template>
  <div class="relative my-1 rounded-md p-2 bg-gray-600 overflow-hidden">
    <div
      class="absolute left-0 top-0 bottom-0 opacity-30"
      :class="background"
      :style="`width: ${Math.round(fraction * 100)}%`"
    />
    <span class="font-bold">{{ statusText }}</span>
    {{ truncatedFriendlyName }}
    <span class="float-right">{{
      fraction < 1 ? Math.round(fraction * 100) + "%" : ""
    }}</span>
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
  computed: {
    truncatedFriendlyName() {
      if (this.download.friendly_name.length > 50) {
        return this.download.friendly_name.substring(0, 50) + "...";
      }
      return this.download.friendly_name;
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
