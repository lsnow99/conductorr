<template>
    <section class="mt-3">
        <o-button variant="primary" @click="showNewMediaServer">New Media Server</o-button>
    </section>
    <o-modal v-model:active="showNewMediaServerModal" @close="closeNewMediaServerModal">
        <new-media-server
        v-if="mediaServerType == ''"
        @close="closeNewMediaServerModal"
        @selected="selectedMediaServer"
      />
    </o-modal>
</template>

<script>
import NewMediaServer from "../components/NewMediaServer.vue"

export default {
    data() {
        return {
            showNewMediaServerModal: false,
            editingName: '',
            mediaServerType: ''
        }
    },
    components: {NewMediaServer},
    methods: {
        closeNewMediaServerModal() {
            this.showNewDownloaderModal = false;
            setTimeout(() => {
                this.mediaServerType = "";
            }, 200);
        },
        showNewMediaServer() {
            this.mediaServerType = ""
            this.editingName = ""
            this.showNewMediaServerModal = true
        },
        selectedMediaServer(mediaServerType) {
            this.mediaServerType = mediaServerType
        }
    }
}
</script>