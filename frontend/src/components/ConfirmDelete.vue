<template>
  <Modal title="Confirm Delete" v-model="computedActive">
    <p class="text-2xl">{{ deleteMessage }}</p>
    <template v-slot:footer>
      <o-button @click="emit('close')">Cancel</o-button>
      <o-button @click="emit('delete')" variant="danger">Delete</o-button>
    </template>
  </Modal>
</template>

<script setup lang="ts">
import { useComputedActive } from "@/util";
import Modal from "./Modal.vue"

const props = withDefaults(defineProps<{
  deleteMessage: string,
  active: boolean
}>(), {
  deleteMessage: "You sure?"
})

const emit = defineEmits<{
  (e: "update:active", newVal: boolean): void,
  (e: "close"): void
  (e: "delete"): void
}>()

const computedActive = useComputedActive(props, emit);
</script>
