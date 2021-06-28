<template>
  <section class="mt-3">
    <div v-for="downloader in downloaders" :key="downloader.id">
      {{ downloader
      }}<o-button variant="primary" @click="editDownloader(downloader)"
        >Edit</o-button
      >
    </div>
    <o-button @click="showNewDownloaderModal = true" variant="primary"
      >Add Downloader</o-button
    >
    <o-modal v-model:active="showNewDownloaderModal" @close="close">
      <new-downloader
        v-if="downloaderType == ''"
        @close="closeNewDownloaderModal"
        @selected="selectedDownloader"
      />
      <edit-transmission
        v-if="downloaderType == 'transmission'"
        @submit="newTransmission"
        @close="closeNewDownloaderModal"
      />
    </o-modal>
    <o-modal
      v-model:active="showEditDownloaderModal"
      @close="showEditDownloaderModal = false"
    >
      <edit-transmission
        :transmission="editingDownloader.config"
        @submit="updateTransmission"
        @close="showEditDownloaderModal = false"
      />
    </o-modal>
  </section>
</template>

<script>
import NewDownloader from "../components/NewDownloader.vue";
import EditTransmission from "../components/EditTransmission.vue";
import APIUtil from "../util/APIUtil";

export default {
  data() {
    return {
      showNewDownloaderModal: false,
      showEditDownloaderModal: false,
      downloaderType: "",
      downloaders: [],
      editingDownloader: {},
    };
  },
  components: {
    NewDownloader,
    EditTransmission,
  },
  methods: {
    selectedDownloader(downloaderType) {
      this.downloaderType = downloaderType;
    },
    closeNewDownloaderModal() {
      this.showNewDownloaderModal = false;
      setTimeout(() => {
        this.downloaderType = "";
      }, 200);
    },
    loadDownloaders() {
      APIUtil.getDownloaders().then((downloaders) => {
        this.downloaders = downloaders;
      });
    },
    editDownloader(downloader) {
      this.showEditDownloaderModal = true;
      this.editingDownloader = downloader;
    },
    updateDownloader(id, config) {
      APIUtil.updateDownloader(id, config).then(() => {
        this.$oruga.notification.open({
          duration: 3000,
          message: `Updated successfully`,
          position: "bottom-right",
          variant: "success",
          closable: false,
        });
        this.showEditDownloaderModal = false;
        this.loadDownloaders();
      });
      },
    newDownloader(downloaderType, config) {
      APIUtil.newDownloader(downloaderType, config).then(() => {
        this.$oruga.notification.open({
          duration: 3000,
          message: `Created successfully`,
          position: "bottom-right",
          variant: "success",
          closable: false,
        });
        this.showNewDownloaderModal = false;
        this.loadDownloaders();
      });
    },
    newTransmission(config) {
      this.newDownloader("transmission", config);
    },
    updateTransmission(config) {
        this.updateDownloader(this.editingDownloader.id, config)
    },
  },
  mounted() {
    this.loadDownloaders();
  },
};
</script>
