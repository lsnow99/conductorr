<template>
  <page-wrapper>
    <section class="flex flex-row">
      <img class="hidden md:block" :src="media.poster" />
      <div class="ml-4">
        <h1 class="text-4xl lg:text-6xl">{{ media.title }}</h1>
        <div class="py-4 flex flex-row items-center">
          <div class="text-2xl mx-4 text-gray-300">
            {{ mediaYear(media) }}
          </div>
          <div class="text-2xl mx-4 text-gray-300">
            <o-icon class="text-lg" icon="star" />
            {{ media.imdb_rating }}%
          </div>
          <a :href="`https://www.imdb.com/title/${media.imdb_id}`" target="_blank" class="inline-block pt-1 mx-4">
              <o-icon
                class="text-4xl text-gray-300 hover:text-yellow-300"
                pack="fab"
                icon="imdb"
              />
            </a>
        </div>
        <p class="text-lg">{{ media.description }}</p>
      </div>
    </section>
  </page-wrapper>
</template>

<script>
import PageWrapper from "../components/PageWrapper.vue";
import APIUtil from "../util/APIUtil";
import MediaUtil from "../util/MediaUtil";

const STATUS_TYPES = ["wanted", ""];

export default {
  data() {
    return {
      media: {},
      mediaID: 0,
    };
  },
  mixins: [MediaUtil],
  components: { PageWrapper },
  created() {
    this.mediaID = parseInt(this.$route.params.media_id);
  },
  mounted() {
    APIUtil.getMedia(this.mediaID).then((media) => {
      this.media = media;
    });
  },
};
</script>
