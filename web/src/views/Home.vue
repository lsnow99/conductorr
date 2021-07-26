<template>
  <page-wrapper>
    <div class="flex flex-col lg:flex-row">
      <section class="flex-1">
        <div class="relative my-1 rounded-md p-2 bg-gray-600 overflow-hidden" v-for="download in downloads" :key="download.identifier">
          <div class="absolute left-0 top-0 bottom-0 opacity-30" :class="getDownloadInfo(download).background" :style="`width: ${getDownloadInfo(download).fraction * 100}%`" />
          <span class="font-bold">{{getDownloadInfo(download).status_text}}</span>
          {{download.friendly_name}}
          <span class="float-right">{{(getDownloadInfo(download).fraction < 1)?getDownloadInfo(download).fraction * 100 + '%':''}}</span>
        </div>
      </section>
      <div class="flex-1">asdf</div>
    </div>
  </page-wrapper>
</template>

<script>
import PageWrapper from "../components/PageWrapper.vue";
import APIUtil from "../util/APIUtil";

export default {
  data() {
    return {
      downloads: [],
    };
  },
  components: { PageWrapper },
  methods: {
    refreshDownloads() {
      APIUtil.getDownloads().then((downloads) => {
        this.downloads = downloads;
        this.downloads = [
          {
            status: "waiting",
            friendly_name: "Sopranos s01e01"
          },
          {
            status: "paused",
            friendly_name: "Rick and Morty s02e04",
          },
          {
            status: "downloading",
            friendly_name: "Bloodline s02e09",
            full_size: 10,
            bytes_left: 3,
          },
          {
            status: "downloading",
            friendly_name: "Lord of the Rings s02e09",
            full_size: 10,
            bytes_left: 4,
          },
          {
            status: "downloading",
            friendly_name: "Tarzan s02e09",
            full_size: 10,
            bytes_left: 9,
          },
          {
            status: "downloading",
            friendly_name: "Moana s02e09",
            full_size: 10,
            bytes_left: 8,
          },
          {
            status: "downloading",
            friendly_name: "Frozen s02e09",
            full_size: 10,
            bytes_left: 1,
          },
          {
            status: "processing",
            friendly_name: "Kevin can Fuck Himself",
          },
          {
            status: "done",
            friendly_name: "The Hobbit"
          },
          {
            status: "error",
            friendly_name: "The Desolation of Smaug"
          }
        ]
      });
    },
    getDownloadInfo(download) {
      if (!download.status) {
        return
      }

      let fraction = 1
      if (download.full_size && download.bytes_left) {
        fraction = (download.full_size - download.bytes_left) / download.full_size
      }

      switch (download.status) {
        case "waiting":
          return {
            fraction,
            status_text: "QUEUED",
            background: "bg-purple-600",
          }
        case "paused":
          return {
            fraction,
            status_text: "PAUSED",
            background: "bg-yellow-400",
          }
        case "downloading":
          return {
            fraction,
            status_text: "DOWNLOADING",
            background: "bg-blue-400",
          }
        case "processing":
        case "cprocessing":
          return {
            fraction,
            status_text: "PROCESSING",
            background: "bg-gray-900",
          }
        case "done":
          return {
            fraction,
            status_text: "DONE",
            background: "bg-green-600",
          }
        case "error":
        case "cerror":
          return {
            fraction,
            status_text: "ERROR",
            background: "bg-red-700"
          }
      }
    }
  },
  mounted() {
    this.refreshDownloads();
  },
};
</script>
