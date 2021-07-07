<template>
  <section class="mt-3">
    <o-button @click="showNewDownloader" variant="primary"
      >Add Downloader</o-button
    >
    <config-item
      @edit="editDownloader(downloader)"
      @delete="deleteDownloader(downloader)"
      collapsible
      :title="downloader.name"
      :delete-message="`Are you sure you want to delete downloader '${downloader.name}'?`"
      v-for="downloader in downloaders"
      :key="downloader.id"
    >
      <div class="flex flex-col">
        Configuration:
        <code class="bg-gray-800 whitespace-pre p-2">
          {{ JSON.stringify(downloader.config, null, 4) }}
        </code>
      </div>
    </config-item>
    <o-modal v-model:active="showNewDownloaderModal" @close="closeNewDownloaderModal">
      <new-downloader
        v-if="downloaderType == ''"
        @close="closeNewDownloaderModal"
        @selected="selectedDownloader"
      />
      <edit-transmission
        v-if="downloaderType == 'transmission'"
        v-model:name="editingName"
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
        v-model:name="editingName"
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
import ConfigItem from "../components/ConfigItem.vue";

export default {
  data() {
    return {
      showNewDownloaderModal: false,
      showEditDownloaderModal: false,
      downloaderType: "",
      downloaders: [],
      editingDownloader: {},
      editingName: "",
    };
  },
  components: {
    NewDownloader,
    EditTransmission,
    ConfigItem,
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
      this.editingName = downloader.name;
    },
    deleteDownloader(downloader) {
      APIUtil.deleteDownloader(downloader.id).then(() => {
        this.$oruga.notification.open({
          duration: 3000,
          message: `Deleted successfully`,
          position: "bottom-right",
          variant: "success",
          closable: false,
        });
        this.loadDownloaders()
      })
    },
    updateDownloader(id, name, config) {
      APIUtil.updateDownloader(id, name, config).then(() => {
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
    newDownloader(downloaderType, name, config) {
      APIUtil.newDownloader(downloaderType, name, config).then(() => {
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
      this.newDownloader("transmission", this.editingName, config);
    },
    updateTransmission(config) {
      this.updateDownloader(
        this.editingDownloader.id,
        this.editingName,
        config
      );
    },
    showNewDownloader() {
      this.downloaderType = ""
      this.editingName = ""
      this.showNewDownloaderModal = true
    }
  },
  mounted() {
    this.loadDownloaders();
  },
};
</script>
