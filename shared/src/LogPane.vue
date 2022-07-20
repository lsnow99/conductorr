<template>
  <div
    v-for="(output, index) in logs"
    :key="index"
    :class="outputClass(output.variant)"
    class="relative p-2 text-lg"
  >
    <div v-if="output.variant == 'success'" class="absolute">
      <vue-fontawesome icon="check-circle" />
    </div>
    <div v-else-if="output.variant == 'danger'" class="absolute">
      <vue-fontawesome icon="exclamation-circle" />
    </div>
    <div v-else-if="output.variant == 'warning'" class="absolute">
      <vue-fontawesome icon="exclamation-triangle" />
    </div>
    <div v-else class="absolute">
      <vue-fontawesome icon="info-circle" />
    </div>
    <div class="float-left ml-8 mr-3">
      {{ formatTime(output.timestamp) }}
    </div>
    <div class="text-gray-100" :class="msgClass(output.decoration)">
      {{ output.msg? output.msg : "<empty>" }}
    </div>
  </div>
</template>

<style scoped>
.success {
  @apply bg-green-600 text-green-300;
}
.danger {
  @apply bg-red-600 text-red-300;
}
.warning {
  @apply bg-yellow-600 text-yellow-300;
}
.default {
  @apply bg-gray-600 text-gray-300;
}
</style>

<script setup lang="ts">
import { DateTime } from "luxon";
import { library } from "@fortawesome/fontawesome-svg-core";
import { FontAwesomeIcon as VueFontawesome } from "@fortawesome/vue-fontawesome";
import {
  faCheckCircle,
  faExclamationCircle,
  faExclamationTriangle,
  faInfoCircle,
} from "@fortawesome/free-solid-svg-icons";
import { LogMessage, Variant } from "@/types";

library.add(faCheckCircle);
library.add(faExclamationCircle);
library.add(faExclamationTriangle);
library.add(faInfoCircle);

const props = withDefaults(
  defineProps<{
    logs: LogMessage[];
  }>(),
  {
    logs: () => [],
  }
);

const outputClass = (variant: Variant) => {
  switch (variant) {
    case Variant.SUCCESS:
      return `success`;
    case Variant.WARNING:
      return `warning`;
    case Variant.DANGER:
      return `danger`;
    default:
      return `default`;
  }
};

const msgClass = (decoration?: string) => {
  if (decoration === "bold") {
    return `font-semibold`;
  } else if (decoration === "italic") {
    return `italic`;
  }
};

const formatTime = (timestamp: DateTime) =>
  timestamp.toLocaleString(DateTime.TIME_WITH_SECONDS);
</script>
