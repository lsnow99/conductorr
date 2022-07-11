<template>
  <section class="mt-3">
    <div class="flex flex-col justify-between sm:flex-row">
      <div class="flex justify-center">
        <o-button @click="showNewDownloader" variant="primary"
          >Add Downloader</o-button
        >
      </div>
      <div class="flex justify-center mt-4 sm:mt-0">
        <o-button
          variant="primary"
          icon-left="plus-square"
          @click="setExpanded(true)"
          class="mr-3"
          >Expand All</o-button
        ><o-button
          variant="primary"
          icon-left="minus-square"
          @click="setExpanded(false)"
          >Collapse All</o-button
        >
      </div>
    </div>
    <config-item
      @edit="editDownloader(downloader, $event)"
      @delete="deleteDownloader(downloader)"
      v-model:expanded="downloader.expanded"
      collapsible
      :title="downloader.name"
      :delete-message="`Are you sure you want to delete downloader '${downloader.name}'?`"
      v-for="downloader in downloaders"
      :key="downloader.id"
    >
      <div class="flex flex-col">
        Configuration:
        <code class="p-2 whitespace-pre bg-gray-800">
          {{ JSON.stringify(downloader.config, null, 4) }}
        </code>
      </div>
    </config-item>
    <new-downloader
      v-model:active="showNewDownloaderModal"
      @close="closeNewDownloaderModal"
      @selected="selectedDownloader"
    />
    <EditService
      v-model:active="showEditDownloaderModal"
      v-model="editingDownloader"
      :title="computedTitle"
      :fields="computedFields"
      :downloader-type="downloaderType"
      :extra-validator="validate"
      :testingMode="testingMode"
      @close="closeEditDownloaderModal"
      @test="testDownloader"
      @save="submittedDownloader"
    >
      <o-field label="Post-Processing Action">
        <div class="flex flex-row justify-center mt-1">
          <radio-group
            name="fileAction"
            v-model="editingDownloader.file_action"
            :options="[
              { text: 'MOVE', value: 'move' },
              { text: 'COPY', value: 'copy' },
            ]"
          />
        </div>
      </o-field>
    </EditService>
  </section>
</template>

<script setup lang="ts">
import NewDownloader from "../components/NewDownloader.vue";
import EditService from "../components/EditService.vue";
import APIUtil from "@/util/APIUtil";
import ConfigItem from "../components/ConfigItem.vue";
import TabSaver from "../util/TabSaver";
import RadioGroup from "../components/RadioGroup.vue";

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
      testingMode: "",
    };
  },
  components: {
    NewDownloader,
    ConfigItem,
    EditService,
    RadioGroup,
  },
  mixins: [TabSaver],
  methods: {
    setExpanded(expanded) {
      this.downloaders.forEach((element) => {
        element.expanded = expanded;
      });
    },
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
      this.mode = "";
      this.downloaderType = "";
      this.editingDownloader = {
        config: {},
      };
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
        });
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
          this.closeEditDownloaderModal();
          this.loadDownloaders();
        });
      }
    },
    testDownloader(downloader) {
      this.testingMode = "loading";
      APIUtil.testDownloader(this.downloaderType, downloader.config)
        .then((resp) => {
          this.$oruga.notification.open({
            duration: 5000,
            message: `Connected successfully`,
            position: "bottom-right",
            variant: "success",
            closable: false,
          });
          this.testingMode = "success";
        })
        .catch((err) => {
          this.$oruga.notification.open({
            duration: 5000,
            message: `Test failedf: ${err}`,
            position: "bottom-right",
            variant: "danger",
            closable: false,
          });
          this.testingMode = "danger";
        })
        .finally(() => {
          setTimeout(() => {
            this.testingMode = "";
          }, 5000);
        });
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
      Object.assign(this.editingDownloader, downloader);
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
    validate() {
      if (!this.editingDownloader.file_action) {
        return "File Action is required";
      }
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
