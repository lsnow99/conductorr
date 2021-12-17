<template>
  <div
    v-for="output in logs"
    :key="output.timestamp"
    :class="outputClass(output.variant)"
    class="p-2 text-lg font-semibold relative"
  >
    <div v-if="output.variant == 'success'" class="absolute mt-1 ml-1">
      <!-- <OIcon icon="check-circle" /> -->
    </div>
    <div v-if="output.variant == 'danger'" class="absolute mt-1 ml-1">
      <!-- <OIcon icon="exclamation-circle" /> -->
    </div>
    <div v-if="output.variant == 'warning'" class="absolute mt-1 ml-1">
      <!-- <OIcon icon="exclamation-triangle" /> -->
    </div>
    <div v-if="output.variant == ''" class="absolute mt-1 ml-1">
      <!-- <OIcon icon="info-circle" /> -->
    </div>
    <div class="mr-3 ml-8 float-left">
      {{ formatTime(output.timestamp) }}
    </div>
    <div class="text-gray-100">{{ output.msg }}</div>
  </div>
</template>

<script>
import { DateTime } from "luxon";
import OIcon from "@oruga-ui/oruga-next"

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
    OIcon,
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
