<template>
  <section>
    <o-button @click="showDownloaderModal = true">Add Downloader</o-button>
    <o-modal v-model:active="showDownloaderModal" @close="close">
      <new-downloader v-if="downloader == ''" @close="close" @selected="selectedDownloader"/>
      <edit-transmission v-if="downloader == 'transmission'" @close="close"/>
    </o-modal>
  </section>
</template>

<script>
import NewDownloader from "./NewDownloader.vue"
import EditTransmission from "./EditTransmission.vue"

export default {
    data() {
        return {
            showDownloaderModal: false,
            downloader: ''
        }
    },
    components: {
        NewDownloader,
        EditTransmission
    },
    methods: {
        selectedDownloader(downloader) {
            this.downloader = downloader
        },
        close() {
            this.showDownloaderModal = false
            setTimeout(() => {
                this.downloader = ''
            }, 200)
        }
    }
}
</script>