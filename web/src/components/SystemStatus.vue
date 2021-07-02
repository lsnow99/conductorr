<template>
  <div class="text-2xl">
    <o-icon icon="times" class="text-red-500" />
    <o-icon icon="check" class="text-green-500" />
    {{system}}
    <ul>
        <li v-for="statusMsg in status.statuses" :key="statusMsg.msg">statusMsg.msg</li>
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
        default: function() {
            return ""
        }
    }
  },
};
</script>
