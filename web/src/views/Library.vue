<template>
  <page-wrapper>
    <section>
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
          pb-10
        "
      >
        <template v-slot:empty class="flex justify-center">
          <div class="flex flex-col mt-24">
            <span class="text-xl text-center">No results in your library</span>
          </div>
        </template>
        <template v-slot:result="{ media }">
          <media-card :media="media" @click="selectedMedia" />
        </template>
      </search-media>
      <div class="flex flex-row justify-center mt-2">
        <o-button variant="primary" @click="addNew"
          >Search for New Media</o-button
        >
      </div>
    </section>
    <new-media
      v-model:active="showNewSearchModal"
      :default-query="query"
      @close="closeNewSearchModal"
    />
  </page-wrapper>
</template>

<script>
import MediaCard from "../components/MediaCard.vue";
import PageWrapper from "../components/PageWrapper.vue";
import NewMedia from "../components/NewMedia.vue";
import APIUtil from "../util/APIUtil";
import SearchMedia from "../components/SearchMedia.vue";
import TabSaver from "../util/TabSaver";

export default {
  data() {
    return {
      query: "",
      results: [],
      totalResults: 0,
      loading: false,
      showNewSearchModal: false,
    };
  },
  mixins: [TabSaver],
  components: { PageWrapper, MediaCard, NewMedia, SearchMedia },
  methods: {
    search(query, contentType, page) {
      this.loading = true;
      APIUtil.searchLibrary(query, contentType, page)
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
      this.$router.push({ name: "media", params: { media_id: media.id } });
    },
    addNew($event) {
      this.showNewSearchModal = true;
      this.lastButton = $event.currentTarget;
    },
    closeNewSearchModal() {
      this.showNewSearchModal = false;
      this.restoreFocus();
    },
  },
};
</script>
