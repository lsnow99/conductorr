<template>
  <div
    v-for="output in logs"
    :key="output.timestamp"
    :class="outputClass(output.variant)"
    class="p-2 text-lg font-semibold relative"
  >
    <div v-if="output.variant == 'success'" class="absolute">
      <vue-fontawesome icon="check-circle" />
    </div>
    <div v-if="output.variant == 'danger'" class="absolute">
      <vue-fontawesome icon="exclamation-circle" />
    </div>
    <div v-if="output.variant == 'warning'" class="absolute">
      <vue-fontawesome icon="exclamation-triangle" />
    </div>
    <div v-if="output.variant == ''" class="absolute">
      <vue-fontawesome icon="info-circle" />
    </div>
    <div class="mr-3 ml-8 float-left">
      {{ formatTime(output.timestamp) }}
    </div>
    <div class="text-gray-100">{{ output.msg }}</div>
  </div>
</template>

<script>
import { DateTime } from "luxon";
import { library } from "@fortawesome/fontawesome-svg-core";
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";
import {
  faCheckCircle,
  faExclamationCircle,
  faExclamationTriangle,
  faInfoCircle,
} from "@fortawesome/free-solid-svg-icons";

library.add(faCheckCircle);
library.add(faExclamationCircle);
library.add(faExclamationTriangle);
library.add(faInfoCircle);

export default {
  props: {
    logs: {
      type: Array,
      default: function () {
        return [];
      },
    },
  },
  components: {
    "vue-fontawesome": FontAwesomeIcon,
  },
  methods: {
    outputClass(variant) {
      if (variant == "success") {
        return `bg-green-600 text-green-300`;
      } else if (variant == "danger") {
        return `bg-red-600 text-red-300`;
      } else if (variant == "warning") {
        return `bg-yellow-600 text-yellow-300`;
      } else {
        return `bg-gray-600 text-gray-300`;
      }
    },
    formatTime(timestamp) {
      return timestamp.toLocaleString(DateTime.TIME_WITH_SECONDS);
    },
  },
};
</script>
