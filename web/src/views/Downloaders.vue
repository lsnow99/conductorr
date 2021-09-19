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
    <EditDownloader
      v-model:active="showEditDownloaderModal"
      v-model="editingDownloader"
      :title="computedTitle"
      :fields="computedFields"
      :downloader-type="downloaderType"
      @close="closeEditDownloaderModal"
      @save="submittedDownloader"
    />
  </section>
</template>

<script>
import NewDownloader from "../components/NewDownloader.vue";
import EditTransmission from "../components/EditTransmission.vue";
import EditNZBGet from "../components/EditNZBGet.vue";
import EditDownloader from "../components/EditDownloader.vue";
import APIUtil from "../util/APIUtil";
import ConfigItem from "../components/ConfigItem.vue";
import TabSaver from "../util/TabSaver";

const NZBGET_FIELDS = [
  {
    type: "text",
    label: "Base URL",
    placeholder: "http://localhost:6789",
    property: "base_url",
    required: true,
    trim: true,
  },
  {
    type: "text",
    label: "Username",
    placeholder: "nzbget",
    property: "username",
    required: true,
    trim: true,
  },
  {
    type: "password",
    label: "Password",
    placeholder: "tegbzn6789",
    property: "password",
    required: true,
    trim: true,
  },
];

const TRANSMISSION_FIELDS = [
  {
    type: "text",
    label: "Base URL",
    placeholder: "http://localhost:9091",
    property: "base_url",
    required: true,
    trim: true,
  },
  {
    type: "text",
    label: "Username",
    placeholder: "transmission",
    property: "username",
    required: true,
    trim: true,
  },
  {
    type: "password",
    label: "Password",
    placeholder: "transmission",
    property: "password",
    required: true,
    trim: true,
  },
];

export default {
  data() {
    return {
      showNewDownloaderModal: false,
      showEditDownloaderModal: false,
      downloaders: [],
      downloaderType: "",
      mode: "edit",
      editingDownloader: {
        config: {},
      },
      editingFileAction: "",
    };
  },
  components: {
    NewDownloader,
    EditTransmission,
    EditNZBGet,
    ConfigItem,
    EditDownloader,
  },
  mixins: [TabSaver],
  methods: {
    selectedDownloader(downloaderType) {
      this.mode = "new";
      this.downloaderType = downloaderType;
      this.showEditDownloaderModal = true;
      this.showNewDownloaderModal = false;
    },
    closeNewDownloaderModal() {
      this.showNewDownloaderModal = false;
      this.restoreFocus();
    },
    closeEditDownloaderModal() {
      this.showEditDownloaderModal = false;
      this.mode = ''
      this.downloaderType = ''
      this.editingDownloader = {
        config: {}
      }
      this.restoreFocus();
    },
    submittedDownloader(downloader) {
      if (this.mode == "new") {
        APIUtil.newDownloader(
          this.downloaderType,
          downloader.name,
          downloader.file_action,
          downloader.config
        ).then(() => {
          this.$oruga.notification.open({
            duration: 3000,
            message: `Created successfully`,
            position: "bottom-right",
            variant: "success",
            closable: false,
          });
          this.closeEditDownloaderModal();
          this.loadDownloaders();
        })
      } else if (this.mode == "edit") {
        APIUtil.updateDownloader(
        downloader.id,
        downloader.name,
        downloader.file_action,
        downloader.config
      ).then(() => {
        this.$oruga.notification.open({
          duration: 3000,
          message: `Updated successfully`,
          position: "bottom-right",
          variant: "success",
          closable: false,
        });
          this.closeEditDownloaderModal()
          this.loadDownloaders();
        });
      }
    },
    loadDownloaders() {
      APIUtil.getDownloaders().then((downloaders) => {
        this.downloaders = downloaders;
      });
    },
    editDownloader(downloader, $event) {
      this.mode = "edit";
      this.showEditDownloaderModal = true;
      this.downloaderType = downloader.downloader_type;
      this.editingDownloader = downloader;
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
    showNewDownloader($event) {
      this.lastButton = $event.currentTarget;
      this.editingName = "";
      this.showNewDownloaderModal = true;
    },
  },
  mounted() {
    this.loadDownloaders();
  },
  computed: {
    computedFields() {
      switch (this.downloaderType) {
        case "transmission":
          return TRANSMISSION_FIELDS;
        case "nzbget":
          return NZBGET_FIELDS;
        default:
          return [];
      }
    },
    computedTitle() {
      switch (this.downloaderType) {
        case "transmission":
          return "Configure Transmission";
        case "nzbget":
          return "Configure NZBGet";
        default:
          return "";
      }
    },
  },
};
</script>
