<template>
  <div>
    <div class="h-96">
      <CSLEditor v-model="code" />
    </div>
    <LogPane :logs="logs" />
    <button @click="addLog">Click</button>
  </div>
</template>

<script setup lang="ts">
import { DateTime } from "luxon";
import LogPane from "./LogPane.vue";
import CSLEditor from "./CSLEditor.vue";
import { ref } from "vue";
import { LogMessage, Variant } from "./types";

const code = ref(`This is some code`)
const logs = ref<LogMessage[]>([
  {
    msg: "Hi",
    variant: Variant.DANGER,
    timestamp: DateTime.now()
  },
  {
    msg: "Hi",
    variant: Variant.SUCCESS,
    timestamp: DateTime.now()
  }
])

const addLog = () => {
  logs.value.push({
    msg: "Hi",
    variant: Variant.SUCCESS,
    timestamp: DateTime.now()
  })
}
</script>
