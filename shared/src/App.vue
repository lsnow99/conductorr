<template>
  <div>
    <div class="h-96">
      <CSLEditor v-model="code" />
    </div>
    <LogPane :logs="logs" />
    <button @click="addLog">Click</button>
  </div>
</template>

<script>
import { DateTime } from "luxon";
import LogPane from "./LogPane.vue";
import CSLEditor from "./CSLEditor.vue";

export default {
  data() {
    return {
      code: `This is some code`,
      logs: [
        {
          msg: "Hi",
          variant: "danger",
          timestamp: DateTime.now(),
        },
        {
          msg: "Hi",
          variant: "success",
          timestamp: DateTime.now(),
        },
      ],
    };
  },
  components: {
    LogPane,
    CSLEditor,
  },
  methods: {
    addLog() {
      this.logs.push({
        msg: "Hi",
        variant: "success",
        timestamp: DateTime.now(),
      });
    },
  },
};
</script>
