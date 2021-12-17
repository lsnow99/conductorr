<template>
  <page-wrapper>
    <div class="flex flex-col lg:flex-row">
      <section class="flex-1">
        <DownloadStatusViewer wrapperClass="h-96" />
      </section>
    </div>
  </page-wrapper>
</template>

<script>
import PageWrapper from "../components/PageWrapper.vue";
import DownloadStatusViewer from "../components/DownloadStatusViewer.vue";

export default {
  data() {
    return {
      activeDownloads: [],
      finishedDownloads: [],
      refreshInterval: -1,
      loadingDownloads: false,
    };
  },
  components: { PageWrapper, DownloadStatusViewer },
};
</script>
