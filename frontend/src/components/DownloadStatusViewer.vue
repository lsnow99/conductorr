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
          :key="download.identifier"
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
</template>

<script>
import DownloadStatus from "../components/DownloadStatus.vue";
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
        default: function() {
            return ""
        }
    }
  },
  data() {
    return {
      activeDownloads: [],
      finishedDownloads: [],
      refreshInterval: -1,
      loadingDownloads: false,
    };
  },
  components: {
    DownloadStatus,
  },
  methods: {
    refreshDownloads() {
      this.loadingDownloads = true;
      APIUtil.getActiveDownloads(this.mediaID).then((downloads) => {
        this.activeDownloads = downloads;
      });
      APIUtil.getFinishedDownloads(this.mediaID).then((downloads) => {
        this.finishedDownloads = downloads;
      });
    },
    getDownloadInfo(download) {
      if (!download.status) {
        return;
      }

      let fraction = 1;
      if (download.full_size && download.bytes_left) {
        fraction =
          (download.full_size - download.bytes_left) / download.full_size;
      }

      switch (download.status) {
        case "waiting":
          return {
            fraction,
            status_text: "QUEUED",
            background: "bg-purple-600",
          };
        case "paused":
          return {
            fraction,
            status_text: "PAUSED",
            background: "bg-yellow-400",
          };
        case "downloading":
          return {
            fraction,
            status_text: "DOWNLOADING",
            background: "bg-blue-400",
          };
        case "processing":
        case "cprocessing":
          return {
            fraction: 1,
            status_text: "PROCESSING",
            background: "bg-gray-900",
          };
        case "done":
          return {
            fraction: 1,
            status_text: "DONE",
            background: "bg-green-600",
          };
        case "error":
        case "cerror":
          return {
            fraction: 1,
            status_text: "ERROR",
            background: "bg-red-700",
          };
      }
    },
    sortDownloads(a, b) {
      return b.id - a.id;
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
