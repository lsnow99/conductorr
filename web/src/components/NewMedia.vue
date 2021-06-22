<template>
  <header class="modal-card-header">
    <p class="modal-card-title">New Media</p>
  </header>
  <section class="modal-card-content">
    <search-media
      v-model:query="query"
      :results="results"
      :total-results="totalResults"
      :loading="loading"
      @search="search"
      @selected-media="selectedMedia"
      results-wrapper-class="
        mt-10
        flex
        flex-row
        flex-wrap
        justify-center
        gap-y-10
        gap-x-10
      "
    >
      <template v-slot:empty class="flex justify-center">
        <div class="flex flex-col mt-24">
          <span class="text-xl text-center">No results</span>
        </div>
      </template>
      <template v-slot:result="{ media }">
        <media-card :media="media" @click="selectedMedia" />
      </template>
    </search-media>
  </section>
  <footer class="modal-card-footer">
    <o-button @click="$emit('close')">Cancel</o-button>
    <o-button variant="primary" @click="save">Save</o-button>
  </footer>
</template>

<script>
import APIUtil from "../util/APIUtil";
import MediaCard from "../components/MediaCard.vue";
import SearchMedia from "./SearchMedia.vue";

export default {
  data() {
    return {
      results: [],
      totalResults: 0,
      loading: false,
      query: "",
    };
  },
  props: {
    mediaType: {
      type: String,
      default: function () {
        return "";
      },
    },
    defaultQuery: {
      type: String,
      default: function () {
        return "";
      },
    },
  },
  components: { MediaCard, SearchMedia },
  emits: ["close"],
  methods: {
    search(query, contentType, page) {
      this.loading = true;
      APIUtil.searchNew(query, contentType, page)
        .then((data) => {
          this.totalResults = data.total_results;
          this.results = data.results;
        })
        .catch(() => {
          this.totalResults = 0;
          this.results = [];
        })
        .finally(() => {
          this.loading = false;
        });
    },
    selectedMedia(media) {
      this.loading = true;
      console.log(media);
      APIUtil.addMedia(media.imdb_id)
        .then((id) => {
          this.$oruga.notification.open({
            duration: 3000,
            message: `Successfully added ${media.title}`,
            position: "bottom-right",
            variant: "success",
            closable: false,
          });
          this.$emit("close");
          this.$router.push({ name: "media", params: { media_id: id } });
        })
        .finally(() => {
          this.loading = false;
        });
    },
  },
  created() {
    this.query = this.defaultQuery;
  },
};
</script>
