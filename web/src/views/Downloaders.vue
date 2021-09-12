<template>
  <section class="mt-3">
    <o-button @click="showNewDownloader" variant="primary"
      >Add Downloader</o-button
    >
    <config-item
      @edit="editDownloader(downloader, $event)"
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
    <new-downloader
      v-model:active="showNewDownloaderModal"
      @close="closeNewDownloaderModal"
      @selected="selectedDownloader"
    />
    <edit-transmission
      v-model:active="showNewTransmissionModal"
      @close="closeNewTransmissionModal"
      @submit="newTransmission"
      v-model:name="editingName"
    />
    <EditNZBGet
      v-model:active="showNewNZBGetModal"
      @close="closeNewNZBGetModal"
      @submit="newNZBGet"
      v-model:name="editingName"
    />
    <edit-transmission
      v-model:active="showEditTransmissionModal"
      @close="closeEditTransmissionModal"
      @submit="updateTransmission"
      v-model:name="editingName"
      :transmission="editingDownloader.config"
    />
    <EditNZBGet
      v-model:active="showEditNZBGetModal"
      @close="closeEditNZBGetModal"
      @submit="updateNZBGet"
      v-model:name="editingName"
      :nzbget="editingDownloader.config"
    />
  </section>
</template>

<script>
import NewDownloader from "../components/NewDownloader.vue";
import EditTransmission from "../components/EditTransmission.vue";
import EditNZBGet from "../components/EditNZBGet.vue";
import APIUtil from "../util/APIUtil";
import ConfigItem from "../components/ConfigItem.vue";
import TabSaver from "../util/TabSaver";

export default {
  data() {
    return {
      showNewDownloaderModal: false,
      showEditTransmissionModal: false,
      showEditNZBGetModal: false,
      showNewTransmissionModal: false,
      showNewNZBGetModal: false,
      downloaderType: "",
      downloaders: [],
      editingDownloader: {},
      editingName: "",
    };
  },
  components: {
    NewDownloader,
    EditTransmission,
    EditNZBGet,
    ConfigItem,
  },
  mixins: [TabSaver],
  methods: {
    selectedDownloader(downloaderType) {
      if (downloaderType == "transmission") {
        this.showNewTransmissionModal = true;
      } else if (downloaderType == "nzbget") {
        this.showNewNZBGetModal = true;
      }
      this.showNewDownloaderModal = false;
    },
    closeNewDownloaderModal() {
      this.showNewDownloaderModal = false;
      this.restoreFocus();
    },
    closeNewTransmissionModal() {
      this.showNewTransmissionModal = false;
      this.restoreFocus();
    },
    closeNewNZBGetModal() {
      this.showNewNZBGetModal = false;
      this.restoreFocus();
    },
    closeEditTransmissionModal() {
      this.showEditTransmissionModal = false;
      this.restoreFocus();
    },
    closeEditNZBGetModal() {
      this.showEditNZBGetModal = false;
      this.restoreFocus();
    },
    loadDownloaders() {
      APIUtil.getDownloaders().then((downloaders) => {
        this.downloaders = downloaders;
      });
    },
    editDownloader(downloader, $event) {
      if (downloader.downloader_type == "transmission") {
        this.showEditTransmissionModal = true;
      } else if (downloader.downloader_type == "nzbget") {
        this.showEditNZBGetModal = true;
      }
      this.editingDownloader = downloader;
      this.editingName = downloader.name;
      this.lastButton = $event.currentTarget;
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
        this.loadDownloaders();
      });
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
    newNZBGet(config) {
      this.newDownloader("nzbget", this.editingName, config);
    },
    updateTransmission(config) {
      this.updateDownloader(
        this.editingDownloader.id,
        this.editingName,
        config
      );
    },
    updateNZBGet(config) {
      this.updateNZBGet(this.editingDownloader.id, this.editingName, config);
    },
    showNewDownloader($event) {
      this.lastButton = $event.currentTarget;
      this.editingName = "";
      this.showNewDownloaderModal = true;
    },
  },
  mounted() {
    this.loadDownloaders();
  },
};
</script>
