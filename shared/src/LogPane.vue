<template>
    <div
      v-for="output in logs"
      :key="output.timestamp"
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
        {{ output.msg }}
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
        return `success`;
      } else if (variant == "danger") {
        return `danger`;
      } else if (variant == "warning") {
        return `warning`;
      } else {
        return `default`;
      }
    },
    msgClass(decoration) {
      if (decoration == "bold") {
        return `font-semibold`;
      } else if (decoration == "italic") {
        return `italic`;
      }
    },
    formatTime(timestamp) {
      return timestamp.toLocaleString(DateTime.TIME_WITH_SECONDS);
    },
  },
};
</script>
