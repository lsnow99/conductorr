<template>
  <Modal title="New Media Server" v-model="computedActive" @close="close">
    <ServiceOptions
      :services="mediaServerTypes"
      v-model="selectedMediaServer"
    />
    <template v-slot:footer>
      <o-button @click="close">Cancel</o-button>
      <div>
        <o-button
          variant="primary"
          :disabled="!!selectedMediaServer"
          @click="next"
          >Next</o-button
        >
      </div>
    </template>
  </Modal>
</template>

<script setup lang="ts">
import ServiceOptions from "./ServiceOptions.vue";
import Modal from "./Modal.vue";
import { MediaServerType, MediaServer } from "@/types/api/media_server";
import { ref, watch } from "vue";
import { useComputedActive } from "@/util";

const selectedMediaServer = ref<MediaServerType | null>(null)
const mediaServerTypes = ref<{name: string, value: MediaServerType}[]>([
  {
    name: "Plex",
    value: MediaServerType.PLEX,
  },
  {
    name: "Jellyfin",
    value: MediaServerType.JELLYFIN,
  },
]);

const props = defineProps<{
  active: boolean
}>()

const emit = defineEmits<{
  (e: "close"): void
  (e: "selected", selectedVal: MediaServer): void
  (e: "update:active", active: boolean): void
}>()

const next = () => selectedMediaServer.value && emit("selected", selectedMediaServer.value)

const close = () => {
  selectedMediaServer.value = null
  emit("close")
}

watch(() => props.active, (newVal) => {
  if (!newVal) {
    selectedMediaServer.value = null
  }
})

const computedActive = useComputedActive(props, emit)
</script>
