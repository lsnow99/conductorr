<template>
  <section class="mt-3">
    <o-button variant="primary" @click="showNewMediaServer"
      >New Media Server</o-button
    >
    <config-item
      @edit="editMediaServer(mediaServer, $event)"
      @delete="deleteMediaServer(mediaServer)"
      collapsible
      :title="mediaServer.name"
      :delete-message="`Are you sure you want to delete mediaServer '${mediaServer.name}'?`"
      v-for="mediaServer in mediaServers"
      :key="mediaServer.id"
    >
      <div class="flex flex-col">
        Configuration:
        <code class="bg-gray-800 whitespace-pre p-2">
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
  <edit-plex
    v-model:name="editingName"
    v-model:active="showNewPlexModal"
    @submit="newPlex"
    @close="closeNewMediaServerModal"
  />
  <edit-plex
    :plex="editingMediaServer.config"
    v-model:name="editingName"
    v-model:active="showEditPlexModal"
    @submit="updatePlex"
    @close="closeEditPlexModal"
  />
</template>

<script>
import NewMediaServer from "../components/NewMediaServer.vue";
import EditPlex from "../components/EditPlex.vue";
import ConfigItem from "../components/ConfigItem.vue";
import APIUtil from "../util/APIUtil";
import TabSaver from "../util/TabSaver"

export default {
  data() {
    return {
      _showNewMediaServerModal: false,
      _showEditMediaServerModal: false,
      editingName: "",
      editingMediaServer: {},
      mediaServerType: "",
      mediaServers: [],
    };
  },
  mixins: [TabSaver],
  components: { NewMediaServer, EditPlex, ConfigItem },
  methods: {
    closeNewMediaServerModal() {
      this.showNewMediaServerModal = false;
      this.mediaServerType = "";
      this.restoreFocus();
    },
    showNewMediaServer($event) {
      this.lastButton = $event.currentTarget
      this.mediaServerType = "";
      this.editingName = "";
      this.showNewMediaServerModal = true;
    },
    selectedMediaServer(mediaServerType) {
      this.mediaServerType = mediaServerType;
    },
    loadMediaServers() {
      APIUtil.getMediaServers().then((mediaServers) => {
        this.mediaServers = mediaServers;
      });
    },
    editMediaServer(mediaServer, $event) {
      this.lastButton = $event.currentTarget
      this.showEditMediaServerModal = true;
      this.editingMediaServer = mediaServer;
      this.editingName = mediaServer.name;
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
    newMediaServer(mediaServerType, name, config) {
      APIUtil.newMediaServer(mediaServerType, name, config).then(() => {
        this.$oruga.notification.open({
          duration: 3000,
          message: `Created successfully`,
          position: "bottom-right",
          variant: "success",
          closable: false,
        });
        this.showNewMediaServerModal = false;
        this.loadMediaServers();
      });
    },
    closeEditPlexModal() {
      this.restoreFocus()
      this.showEditPlexModal = false;
    },
    newPlex(config) {
      this.newMediaServer("plex", this.editingName, config);
    },
    updatePlex(config) {
      this.updateMediaServer(
        this.editingMediaServer.id,
        this.editingName,
        config
      );
    },
  },
  mounted() {
    this.loadMediaServers();
  },
  computed: {
    showNewMediaServerModal: {
      get() {
        return this._showNewMediaServerModal && this.mediaServerType == "";
      },
      set(v) {
        this._showNewMediaServerModal = v;
      },
    },
    showEditPlexModal: {
      get() {
        return (
          this.editingMediaServer.media_server_type == "plex" &&
          this._showEditMediaServerModal
        );
      },
      set(v) {
        this._showEditMediaServerModal = v;
      },
    },
    showNewPlexModal: {
      get() {
        return this._showNewMediaServerModal && this.mediaServerType == "plex";
      },
      set(v) {
        this._showNewMediaServerModal = v;
      },
    },
  },
};
</script>
