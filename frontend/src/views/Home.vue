<template>
  <page-wrapper>
    <section>
      <h1 class="text-2xl">Recently Added to Library</h1>
      <div class="flex flex-wrap p-2 justify-between">
        <template v-for="media in recentMedia" :key="media.id">
          <MediaCard :media="media" @click="gotoMedia"></MediaCard>
        </template>
      </div>
    </section>
    <section class="flex-1">
      <DownloadStatusViewer wrapperClass="h-96" />
    </section>
  </page-wrapper>
</template>

<script>
import PageWrapper from "../components/PageWrapper.vue";
import DownloadStatusViewer from "../components/DownloadStatusViewer.vue";
import APIUtil from "../util/APIUtil";
import MediaCard from "../components/MediaCard.vue";

export default {
  data() {
    return {
      activeDownloads: [],
      finishedDownloads: [],
      refreshInterval: -1,
      loadingDownloads: false,
      recentMedia: [],
    };
  },
  components: { PageWrapper, DownloadStatusViewer, MediaCard },
  methods: {
    gotoMedia(media) {
      this.$router.push({ name: 'media', params: { media_id: media.id } });
    },
  },
  mounted() {
    APIUtil.getRecentMedia().then((media) => {
      this.recentMedia = media;
    });
  },
};
</script>
