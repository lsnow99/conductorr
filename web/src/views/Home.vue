<template>
  <page-wrapper>
    <div class="flex flex-col lg:flex-row">
      <section class="flex-1">
        <div class="text-xl flex flex-row justify-between">
          Active Downloads<o-icon
            @click="refreshDownloads"
            icon="sync-alt cursor-pointer"
          />
        </div>
        <div class="h-96 overflow-y-scroll p-2">
          <DownloadStatus
            v-for="download in orderedActiveDownloads"
            :key="download.identifier"
            :download="download"
          />
        </div>
      </section>
      <section class="flex-1">
        <div class="text-xl flex flex-row justify-between">
          Recently Completed Downloads<o-icon
            @click="refreshDownloads"
            icon="sync-alt cursor-pointer"
          />
        </div>
        <div class="h-96 overflow-y-scroll p-2">
          <DownloadStatus
            v-for="download in orderedDoneDownloads"
            :key="download.identifier"
            :download="download"
          />
        </div>
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
      doneDownloads: [],
    };
  },
  components: { PageWrapper, DownloadStatus },
  methods: {
    refreshDownloads() {
      APIUtil.getActiveDownloads().then((downloads) => {
        this.activeDownloads = downloads;
      });
      APIUtil.getDoneDownloads().then((downloads) => {
        this.doneDownloads = downloads;
      })
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
      return a.id - b.id
    }
  },
  mounted() {
    this.refreshDownloads();
  },
  computed: {
    orderedDoneDownloads() {
      return this.doneDownloads.sort(this.sortDownloads)
    },
    orderedActiveDownloads() {
      let processing = this.activeDownloads.filter((elem) => (elem.status == 'cprocessing' || elem.status == 'processing'))
      let downloading = this.activeDownloads.filter((elem) => elem.status == 'downloading')
      let waiting = this.activeDownloads.filter((elem) => elem.status == 'waiting')
      let paused = this.activeDownloads.filter((elem) => elem.status == 'paused')

      processing = processing.sort(this.sortDownloads)
      downloading = downloading.sort(this.sortDownloads)
      waiting = waiting.sort(this.sortDownloads)
      paused = paused.sort(this.sortDownloads)

      return [...processing, ...downloading, ...waiting, ...paused]
    }
  }
};
</script>
