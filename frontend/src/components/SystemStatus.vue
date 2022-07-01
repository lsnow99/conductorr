<template>
  <div
    v-if="status"
    class="p-2"
    :class="status.healthy === false ? 'bg-red-500' : 'bg-green-500'"
  >
    <div class="flex flex-col justify-between text-xl md:flex-row">
      <div>
        <o-icon v-if="status.healthy === false" icon="times" />
        <o-icon v-else icon="check" />
        {{ system }}
      </div>
      <div>
        {{ status.msg }}
      </div>
    </div>
    <ul>
      <li v-for="statusMsg in status.statuses" :key="statusMsg.msg">
        <b>{{ statusMsg.system_name }}</b
        >: {{ statusMsg.msg }}
      </li>
    </ul>
  </div>
</template>

<script setup lang="ts">
import { Status } from "@/types/api/status";

const props = defineProps<{
  status: Status | null;
  system: string;
}>();
</script>
