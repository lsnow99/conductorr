<template>
  <Modal
    title="New Downloader"
    v-model="computedActive"
    @close="emit('close')"
  >
    <ServiceOptions :services="downloaderTypes" v-model="selectedDownloader" />
    <template v-slot:footer>
      <o-button @click="emit('close')">Cancel</o-button>
      <div>
        <o-button
          variant="primary"
          :disabled="selectedDownloader === null"
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
import { ref } from "vue";
import { useComputedActive } from "@/util";
import { DownloaderType } from "@/types/api/downloader";

const downloaderTypes = ref([
  {
    name: "Transmission",
    value: "transmission",
  },
  {
    name: "NZBGet",
    value: "nzbget",
  },
])

const selectedDownloader = ref<DownloaderType | null>(null)

const props = defineProps<{
  active: boolean
}>()

const emit = defineEmits<{
  (e: "close"): void,
  (e: "selected", downloaderType: DownloaderType): void,
  (e: "update:active", active: boolean): void
}>()

const next = () => {
  emit("selected", selectedDownloader.value!)
  selectedDownloader.value = null
}

const computedActive = useComputedActive(props, emit)
</script>
