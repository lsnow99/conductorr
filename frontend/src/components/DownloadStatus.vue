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

<script setup lang="ts">
import { computed } from "@vue/reactivity";

const props = defineProps<{
  download: Download;
}>();

const truncatedFriendlyName = computed(() => {
  if (props.download.friendlyName.length > 50) {
    return props.download.friendlyName.substring(0, 50) + "...";
  }
  return props.download.friendlyName;
});

const deletable = computed(() => {
  switch (props.download.status) {
    case "waiting":
    case "paused":
    case "downloading":
    case "processing":
      return true;
    default:
      return false;
  }
});

const background = computed(() => {
  switch (props.download.status) {
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
});

const statusText = computed(() => {
  switch (props.download.status) {
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
});

const fraction = computed(() => {
  let fraction = 1;
  if (props.download.download.fullSize && props.download.download.bytesSeft) {
    fraction =
      (props.download.download.fullSize - props.download.download.bytesSeft) /
      props.download.download.fullSize;
  }
  fraction = Math.round(fraction * 100) / 100;
  return fraction;
});
</script>
