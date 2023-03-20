<template>
  <o-tabs type="boxed" navTabsClass="text-lg sm:text-2xl">
    <o-tab-item label="Active Downloads">
      <div :class="wrapperClass" class="p-2 overflow-y-scroll">
        <DownloadStatus
          v-for="download in orderedActiveDownloads"
          :key="download.identifier"
          :download="download"
        />
        <div
          v-if="orderedActiveDownloads.length == 0"
          class="flex items-center justify-center flex-1 h-full"
        >
          No active downloads
        </div>
      </div>
    </o-tab-item>
    <o-tab-item label="Completed Downloads">
      <div :class="wrapperClass" class="p-2 overflow-y-scroll">
        <DownloadStatus
          v-for="download in orderedFinishedDownloads"
          :key="download.id"
          :download="download"
        />
        <div
          v-if="orderedFinishedDownloads.length == 0"
          class="flex items-center justify-center flex-1 h-full"
        >
          No completed downloads
        </div>
      </div>
    </o-tab-item>
  </o-tabs>
  <ConfirmDelete
    v-model:active="showConfirmDeleteModal"
    @delete="doDelete"
    @close="closeDelete"
    delete-message="Are you sure you want to delete this download?"
  />
</template>

<script setup lang="ts">
import DownloadStatus from "../components/DownloadStatus.vue";
import ConfirmDelete from "../components/ConfirmDelete.vue";
import APIUtil from "../util/APIUtil";
import { onMounted, onUnmounted, computed, ref } from "vue";
import { useTabSaver } from "@/util";
import { Download } from "@/types/api/download";

const props = defineProps<{
  mediaID?: number,
  wrapperClass: string
}>()

const activeDownloads = ref<Download[]>([])
const finishedDownloads = ref<Download[]>([])
const refreshInterval = ref(-1)
const loadingDownloads = ref(false)
const deletingIdentifier = ref<string | null>(null)
const showConfirmDeleteModal = ref(false)

const { lastButton, restoreFocus } = useTabSaver()

const refreshDownloads = async() => {
  loadingDownloads.value = true;
  try {
    const downloads = await APIUtil.getActiveDownloads(props.mediaID)
    activeDownloads.value = downloads
  } catch {}
  try {
    const downloads = await APIUtil.getFinishedDownloads(props.mediaID)
    activeDownloads.value = downloads
  } catch {}
}

const sortDownloads = (a: Download, b: Download) => {
  return b.id - a.id;
}

const closeDelete = () => {
  showConfirmDeleteModal.value = false;
  restoreFocus()
}

const doDelete = () => {
  closeDelete();
}

const confirmDelete = ($event: Event, identifier: string) => {
  lastButton.value = ($event.currentTarget as HTMLElement);
  showConfirmDeleteModal.value = false
  deletingIdentifier.value = identifier;
}

onMounted(() => {
  refreshDownloads()
  //refreshInterval.value = setInterval(refreshDownloads, 8000)
})

onUnmounted(() => {
  clearInterval(refreshInterval.value)
})

const orderedFinishedDownloads = computed(() => {
  return finishedDownloads.value.sort(sortDownloads)
})

const orderedActiveDownloads = computed(() => {
  let processing = activeDownloads.value.filter(
    (elem: Download) => elem.status === "cprocessing" || elem.status === "processing"
  )
  let downloading = activeDownloads.value.filter(
    (elem: Download) => elem.status === "downloading"
  );
  let waiting = activeDownloads.value.filter(
    (elem: Download) => elem.status === "waiting"
  );
  let paused = activeDownloads.value.filter(
    (elem: Download) => elem.status === "paused"
  );
  processing = processing.sort(sortDownloads);
    downloading = downloading.sort(sortDownloads);
    waiting = waiting.sort(sortDownloads);
    paused = paused.sort(sortDownloads);

    return [...processing, ...downloading, ...waiting, ...paused];
})
</script>
