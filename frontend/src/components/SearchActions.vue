<template>
  <o-tooltip
    variant="info"
    :position="tooltipPosition"
    :active="tooltipActive"
    label="Search/Download Automatically"
  >
    <div :class="fontSize" class="mx-2 text-gray-300">
      <div
        v-if="!loadingAutoSearch"
        @click="searchAuto"
        @keydown.enter="searchAuto"
        @keydown.space="searchAuto"
        tabindex="0"
        role="button"
        aria-label="Search automatically"
      >
        <o-icon class="cursor-pointer" icon="bolt" />
      </div>
      <o-icon v-else icon="sync-alt" spin />
    </div>
  </o-tooltip>
  <o-tooltip
    variant="info"
    :position="tooltipPosition"
    label="Search Manually"
    :active="tooltipActive"
  >
    <div :class="fontSize" class="mx-2 text-gray-300">
      <div
        v-if="!loadingManualSearch"
        @click="searchManual"
        @keydown.space="searchManual"
        @keydown.enter="searchManual"
        tabindex="0"
        role="button"
        aria-label="Search manually"
      >
        <o-icon class="cursor-pointer" icon="search" />
      </div>
      <o-icon v-else icon="sync-alt" spin />
    </div>
  </o-tooltip>
    <manual-search-results
      v-model:active="showManualReleasesModal"
      @close="closeManualReleases"
      :releases="releases"
      :loading="loadingManualSearch"
      :mediaID="mediaID"
    />
</template>

<script setup lang="ts">
import APIUtil from "../util/APIUtil";
import ManualSearchResults from "./ManualSearchResults.vue";
import { computed, inject, onMounted, ref } from "vue";
import { TooltipPosition } from "@/types/tooltip"
import { useTabSaver, useTooltip } from "@/util"

const oruga = inject('oruga')

const props = defineProps<{
  mediaID: number,
  size: string
}>()

const releases = ref([])
const showManualReleasesModal = ref(false)
const tooltipPosition = ref<TooltipPosition>("bottom")
const loadingManualSearch = ref(false)
const loadingAutoSearch = ref(false)

const { restoreFocus } = useTabSaver()
const { tooltipActive, resetTooltips } = useTooltip()

const searchManual = async() => {
  loadingManualSearch.value = true
  try {
    const releases = await APIUtil.searchReleasesManual(props.mediaID)
  } catch (re) {
    oruga.notification.open({
      duration: 3000,
      message: `Error searching: ${re.msg}`,
      variant: "danger",
      closable: false,
      position: "bottom-right",
    });
  } finally {
    loadingManualSearch.value = false
    resetTooltips()
  }
}

const searchAuto = async() => {
  loadingAutoSearch.value = true;
  try {
    const num = await APIUtil.searchReleasesAuto(props.mediaID)
    if (num > 0) {
      oruga.notification.open({
        duration: 3000,
        message: `Successfully queued ${num} releases for auto downloading`,
        variant: "success",
        closable: false,
        position: "bottom-right",
      });
    } else {
      oruga.notification.open({
        duration: 3000,
        message: `Could not find any releases to queue`,
        variant: "danger",
        closable: false,
        position: "bottom-right",
      });
    }
  } catch (re) {
    oruga.notification.open({
      duration: 3000,
      message: `Error searching: ${re.msg}`,
      variant: "danger",
      closable: false,
      position: "bottom-right",
    });
  } finally {
    loadingAutoSearch.value = false;
    resetTooltips();
  }
}

const closeManualReleases = () => {
  showManualReleasesModal.value = false;
  restoreFocus()
}

onMounted(() => {
  const screenWidth = window.innerWidth;
  if (screenWidth < 768) {
    tooltipPosition.value = "left";
  }
})

const fontSize = computed(() => {
  switch(props.size) {
    case "small":
      return "text-md"
    case "large":
      return "text-2xl"
    default:
      return "text-lg"
  }
})
</script>
