<template>
  <page-wrapper>
    <section class="flex flex-row">
      <img class="hidden md:block" :src="media.poster" />
      <div class="ml-4">
        <h1 class="text-4xl lg:text-6xl">{{ media.title }}</h1>
        <div class="py-4">
          <o-dropdown aria-role="list">
            <template v-slot:trigger="{ active }">
              <o-button variant="primary">
                <span>{{media.status}}</span>
                <o-icon v-if="active" icon="caret-up" />
                <o-icon v-else icon="caret-down" />
              </o-button>
            </template>

            <o-dropdown-item aria-role="listitem">Action</o-dropdown-item>
            <o-dropdown-item aria-role="listitem"
              >Another action</o-dropdown-item
            >
            <o-dropdown-item aria-role="listitem"
              >Something else</o-dropdown-item
            >
          </o-dropdown>
          <div>
              {{media.imdb_rating}}%
          </div>
        </div>
        <p class="text-lg">{{ media.description }}</p>
      </div>
    </section>
  </page-wrapper>
</template>

<script>
import PageWrapper from "../components/PageWrapper.vue";
import APIUtil from "../util/APIUtil";

const STATUS_TYPES = [
    "wanted",
    ""
]

export default {
  data() {
    return {
      media: {},
      mediaID: 0,
    };
  },
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
