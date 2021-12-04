<template>
  <page-wrapper>
    <div class="flex flex-col lg:flex-row">
      <section class="flex-1">
        <o-tabs type="boxed" navTabsClass="text-2xl">
          <o-tab-item label="Active Downloads">
            <div class="h-96 overflow-y-scroll p-2">
              <DownloadStatus
                v-for="download in orderedActiveDownloads"
                :key="download.identifier"
                :download="download"
              />
              <div v-if="orderedActiveDownloads.length == 0" class="flex flex-1 justify-center h-full items-center">
                No active downloads
              </div>
            </div>
          </o-tab-item>
          <o-tab-item label="Completed Downloads">
            <div class="text-xl flex flex-row justify-between">
              <o-icon
                @click="refreshDownloads"
                @keydown.enter="refreshDownloads"
                @keydown.space="refreshDownloads"
                tabindex="0"
                role="button"
                aria-label="Refresh downloads"
                class="cursor-pointer"
                icon="sync-alt"
                :spin="loadingDownloads"
              />
            </div>
            <div class="h-96 overflow-y-scroll p-2">
              <DownloadStatus
                v-for="download in orderedFinishedDownloads"
                :key="download.identifier"
                :download="download"
              />
              <div v-if="orderedActiveDownloads.length == 0" class="flex flex-1 justify-center h-full items-center">
                No completed downloads
              </div>
            </div>
          </o-tab-item>
        </o-tabs>
      </section>
    </div>
  </page-wrapper>
</template>

<script>
import PageWrapper from "../components/PageWrapper.vue";
import DownloadStatus from "../components/DownloadStatus.vue";
import APIUtil from "../util/APIUtil";

export default {
  data() {
    return {
      activeDownloads: [],
      finishedDownloads: [],
      refreshInterval: -1,
      loadingDownloads: false,
    };
  },
  components: { PageWrapper, DownloadStatus },
  methods: {
    refreshDownloads() {
      this.loadingDownloads = true;
      let requestsCompleted = 0;
      APIUtil.getActiveDownloads().then((downloads) => {
        this.activeDownloads = downloads;
        requestsCompleted++;
        if (requestsCompleted == 2) {
          this.loadingDownloads = false;
        }
      });
      APIUtil.getFinishedDownloads().then((downloads) => {
        this.finishedDownloads = downloads;
        requestsCompleted++;
        if (requestsCompleted == 2) {
          this.loadingDownloads = false;
        }
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
            fraction,
            status_text: "PROCESSING",
            background: "bg-gray-900",
          };
        case "done":
          return {
            fraction,
            status_text: "DONE",
            background: "bg-green-600",
          };
        case "error":
        case "cerror":
          return {
            fraction,
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
