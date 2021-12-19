<template>
  <div class="p-2" :class="status.healthy === false ? 'bg-red-500' : 'bg-green-500'">
    <div class="text-xl flex flex-col md:flex-row justify-between">
      <div>
      <o-icon v-if="status.healthy === false" icon="times" />
      <o-icon v-else icon="check" />
      {{ system }}
      </div>
      <div>
      {{status.msg}}
      </div>
    </div>
    <ul>
      <li v-for="statusMsg in status.statuses" :key="statusMsg.msg">
        <b>{{statusMsg.system_name}}</b>: {{ statusMsg.msg }}
      </li>
    </ul>
  </div>
</template>

<script>
export default {
  props: {
    status: {
      type: Object,
      default: function () {
        return {};
      },
    },
    system: {
      type: String,
      default: function () {
        return "";
      },
    },
  },
};
</script>
