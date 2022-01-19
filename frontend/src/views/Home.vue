<template>
  <page-wrapper>
    <section>
      <h1 class="text-2xl">Recently Added to Library</h1>
      <div class="flex flex-wrap p-2 justify-between">
        <o-carousel ref="carousel" repeat>
          <template v-slot:arrow="props">
            <div
              v-if="props.hasPrev"
              @click="$refs.carousel.prev"
              class="absolute top-1/2 -translate-y-1/2 left-2 h-10 w-10 rounded-full bg-gray-50 shadow-md cursor-pointer"
            >
              <div class="relative w-full h-full">
                <o-icon
                  class="text-yellow-500 text-2xl absolute top-1/2 -translate-y-1/2 left-1/2 -translate-x-1/2"
                  icon="chevron-left"
                ></o-icon>
              </div>
            </div>
            <div
              v-if="props.hasNext"
              @click="$refs.carousel.next"
              class="absolute top-1/2 -translate-y-1/2 right-2 h-10 w-10 rounded-full bg-gray-50 shadow-md cursor-pointer"
            >
              <div class="relative w-full h-full">
                <o-icon
                  class="text-yellow-500 text-2xl absolute top-1/2 -translate-y-1/2 left-1/2 -translate-x-1/2"
                  icon="chevron-right"
                ></o-icon>
              </div>
            </div>
          </template>
          <o-carousel-item v-for="media in recentMedia" :key="media.id">
            <div class="flex flex-row justify-center md:justify-start md:px-16 py-4">
              <div class="flex flex-1 flex-row justify-center">
                <MediaCard :media="media" @click="gotoMedia"></MediaCard>
              </div>
              <div class="hidden md:flex flex-col p-4">
                <span class="text-3xl">{{media.title}}</span>
                <span class="text-lg mt-8">{{media.description}}</span>
                <span class="self-end"><router-link :to="{name: 'media', params: {media_id: media.id}}"><o-button>Go to Media</o-button></router-link></span>
              </div>
            </div>
          </o-carousel-item>
        </o-carousel>
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
      this.$router.push({ name: "media", params: { media_id: media.id } });
    },
  },
  mounted() {
    APIUtil.getRecentMedia().then((media) => {
      this.recentMedia = media;
    });
  },
};
</script>
