<template>
  <section class="mt-3">
    <div class="flex flex-col justify-between sm:flex-row">
      <div class="flex justify-center">
        <o-button variant="primary" @click="openNewMediaServerModal"
          >New Media Server</o-button
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
      @edit="editMediaServer(mediaServer, $event)"
      @delete="deleteMediaServer(mediaServer)"
      v-model:expanded="mediaServer.expanded"
      collapsible
      :title="mediaServer.name"
      :delete-message="`Are you sure you want to delete mediaServer '${mediaServer.name}'?`"
      v-for="mediaServer in mediaServers"
      :key="mediaServer.id"
    >
      <div class="flex flex-col">
        Configuration:
        <code class="p-2 whitespace-pre bg-gray-800">
          {{ JSON.stringify(mediaServer.config, null, 4) }}
        </code>
      </div>
    </config-item>
  </section>
  <new-media-server
    v-model:active="showNewMediaServerModal"
    @close="closeNewMediaServerModal"
    @selected="selectedMediaServer"
  />
  <EditService
    v-model="editingMediaServer"
    v-model:active="showEditMediaServerModal"
    :title="computedTitle"
    :fields="computedFields"
    :testingMode="testingMode"
    @close="closeEditMediaServerModal"
    @test="testMediaServer"
    @save="submittedMediaServer"
  />
</template>

<script setup lang="ts">
import NewMediaServer from "../components/NewMediaServer.vue";
import ConfigItem from "../components/ConfigItem.vue";
import APIUtil from "../util/APIUtil";
import TabSaver from "../util/TabSaver";
import Modal from "../components/Modal.vue";
import ActionButton from "../components/ActionButton.vue";
import EditService from "../components/EditService.vue";
import { ref } from "vue"

const JELLYFIN_FIELDS = [
  {
    label: "Base URL",
    type: "text",
    property: "base_url",
    placeholder: "Base URL",
  },
  {
    label: "API Key",
    type: "text",
    property: "api_key",
    placeholder: "API Key",
  },
];

const PLEX_FIELDS = [
  {
    label: "Base URL",
    type: "text",
    property: "base_url",
    placeholder: "Base URL",
  },
  {
    label: "Auth Token",
    type: "custom",
    property: "token",
    component: "PlexAuthTokenInput",
  },
];

const showNewMediaServerModal = ref(false)
const showEditMediaServerModal = ref(false)
const editingMediaServer = ref<>({config: {}})

export default {
  data() {
    return {
      showNewMediaServerModal: false,
      showEditMediaServerModal: false,
      editingMediaServer: {
        config: {},
      },
      mode: "",
      mediaServerType: "",
      mediaServers: [],
      testingMode: "",
    };
  },
  mixins: [TabSaver],
  components: {
    NewMediaServer,
    ConfigItem,
    Modal,
    ActionButton,
    EditService,
  },
  methods: {
    setExpanded(expanded) {
      this.mediaServers.forEach((element) => {
        element.expanded = expanded;
      });
    },
    closeNewMediaServerModal() {
      this.showNewMediaServerModal = false;
      this.mediaServerType = "";
      this.editingMediaServer = {
        config: {}
      }
      this.restoreFocus();
    },
    openNewMediaServerModal($event) {
      this.lastButton = $event.currentTarget;
      this.mediaServerType = "";
      this.mode = "new"
      this.showNewMediaServerModal = true;
    },
    selectedMediaServer(mediaServerType) {
      this.mediaServerType = mediaServerType;
      this.showEditMediaServerModal = true;
      this.showNewMediaServerModal = false;
    },
    loadMediaServers() {
      APIUtil.getMediaServers().then((mediaServers) => {
        this.mediaServers = mediaServers;
      });
    },
    editMediaServer(mediaServer, $event) {
      this.lastButton = $event.currentTarget;
      this.showEditMediaServerModal = true;
      this.mode = "edit"
      this.mediaServerType = mediaServer.media_server_type
      Object.assign(this.editingMediaServer, mediaServer);
    },
    deleteMediaServer(mediaServer) {
      APIUtil.deleteMediaServer(mediaServer.id).then(() => {
        this.$oruga.notification.open({
          duration: 3000,
          message: `Deleted successfully`,
          position: "bottom-right",
          variant: "success",
          closable: false,
        });
        this.loadMediaServers();
      });
    },
    updateMediaServer(id, name, config) {
      APIUtil.updateMediaServer(id, name, config).then(() => {
        this.$oruga.notification.open({
          duration: 3000,
          message: `Updated successfully`,
          position: "bottom-right",
          variant: "success",
          closable: false,
        });
        this.showEditMediaServerModal = false;
        this.loadMediaServers();
      });
    },
    closeEditMediaServerModal() {
      this.restoreFocus();
      this.mode = ""
      this.mediaServerType = ""
      this.editingMediaServer = {
        config: {}
      }
      this.showEditMediaServerModal = false;
    },
    submittedMediaServer(mediaServer) {
      if (this.mode == "new") {
        APIUtil.newMediaServer(
          this.mediaServerType,
          mediaServer.name,
          mediaServer.config
        ).then(() => {
          this.$oruga.notification.open({
            duration: 3000,
            message: `Created successfully`,
            position: "bottom-right",
            variant: "success",
            closable: false,
          });
          this.closeEditMediaServerModal();
          this.loadMediaServers();
        });
      } else if (this.mode == "edit") {
        APIUtil.updateMediaServer(
          mediaServer.id,
          mediaServer.name,
          mediaServer.config
        ).then(() => {
          this.$oruga.notification.open({
            duration: 3000,
            message: `Updated successfully`,
            position: "bottom-right",
            variant: "success",
            closable: false,
          });
          this.closeEditMediaServerModal();
          this.loadMediaServers();
        });
      }
    },
    testMediaServer(mediaServer) {
      this.testingMode = "loading";
      APIUtil.testMediaServer(this.mediaServerType, mediaServer.config)
        .then(() => {
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
            message: `Test failed: ${err}`,
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
  },
  mounted() {
    this.loadMediaServers();
  },
  computed: {
    computedFields() {
      switch (this.mediaServerType) {
        case "plex":
          return PLEX_FIELDS;
        case "jellyfin":
          return JELLYFIN_FIELDS;
        default:
          return [];
      }
    },
    computedTitle() {
      switch (this.mediaServerType) {
        case "plex":
          return "Configure Plex";
        case "jellyfin":
          return "Configure Jellyfin";
      }
    },
  },
};
</script>
