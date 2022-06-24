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
          @fda="tes"
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

<script>
import DownloadStatus from "../components/DownloadStatus.vue";
import ConfirmDelete from "../components/ConfirmDelete.vue";
import APIUtil from "../util/APIUtil";

export default {
  props: {
    mediaID: {
      type: Number,
      default: function () {
        return 0;
      },
    },
    wrapperClass: {
      type: String,
      default: function () {
        return "";
      },
    },
  },
  data() {
    return {
      activeDownloads: [],
      finishedDownloads: [],
      refreshInterval: -1,
      loadingDownloads: false,
      deletingIdentifier: null,
      showConfirmDeleteModal: false
    };
  },
  components: {
    DownloadStatus,
    ConfirmDelete
  },
  methods: {
    tes() {
      console.log('tes')
    },
    refreshDownloads() {
      this.loadingDownloads = true;
      APIUtil.getActiveDownloads(this.mediaID).then((downloads) => {
        this.activeDownloads = downloads;
      });
      APIUtil.getFinishedDownloads(this.mediaID).then((downloads) => {
        this.finishedDownloads = downloads;
      });
    },
    sortDownloads(a, b) {
      return b.id - a.id;
    },
    doDelete() {
      alert('deleted' + this.deletingIdentifier)
      this.closeDelete();
    },
    closeDelete() {
      this.showConfirmDeleteModal = false;
      this.restoreFocus();
    },
    confirmDelete($event, identifier) {
      console.log('test')
      this.lastButton = $event.currentTarget;
      this.showConfirmDeleteModal = true;
      this.deletingIdentifier = identifier;
    },
  },
  mounted() {
    this.refreshDownloads();
    this.refreshInterval = setInterval(this.refreshDownloads, 3000);
  },
  unmounted() {
    clearInterval(this.refreshInterval);
  },
  computed: {
    orderedFinishedDownloads() {
      return this.finishedDownloads.sort(this.sortDownloads);
    },
    orderedActiveDownloads() {
      let processing = this.activeDownloads.filter(
        (elem) => elem.status == "cprocessing" || elem.status == "processing"
      );
      let downloading = this.activeDownloads.filter(
        (elem) => elem.status == "downloading"
      );
      let waiting = this.activeDownloads.filter(
        (elem) => elem.status == "waiting"
      );
      let paused = this.activeDownloads.filter(
        (elem) => elem.status == "paused"
      );

      processing = processing.sort(this.sortDownloads);
      downloading = downloading.sort(this.sortDownloads);
      waiting = waiting.sort(this.sortDownloads);
      paused = paused.sort(this.sortDownloads);

      return [...processing, ...downloading, ...waiting, ...paused];
    },
  },
};
</script>
