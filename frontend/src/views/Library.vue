<template>
  <page-wrapper>
    <section>
      <search-media
        v-model:query="query"
        v-model:currentPage="currentPage"
        v-model:contentType="contentType"
        :results="results"
        :total-results="totalResults"
        :per-page="perPage"
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
      <div class="flex flex-col md:flex-row text-lg mt-2">
        <div class="flex flex-row items-center mx-4">
          <div class="w-6 h-2 bg-green-600 mr-1" />
          Complete/All Monitored Available
        </div>
        <div class="flex flex-row items-center mx-4">
          <div class="w-6 h-2 bg-purple-600 mr-1" />
          Continuing/All Monitored Available
        </div>
        <div class="flex flex-row items-center mx-4">
          <div class="w-6 h-2 bg-gray-500 mr-1" />
          Unmonitored
        </div>
        <div class="flex flex-row items-center mx-4">
          <div class="w-6 h-2 bg-red-600 mr-1" />
          Missing/Monitored
        </div>
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
      perPage: 0,
      currentPage: 1,
      contentType: "",
      loading: false,
      showNewSearchModal: false,
    };
  },
  mixins: [TabSaver],
  components: { PageWrapper, MediaCard, NewMedia, SearchMedia },
  methods: {
    search(query, contentType, page) {
      this.$router.replace({
        name: "library",
        query: { q: query, content_type: contentType, page: page },
      });
      this.loading = true;
      APIUtil.searchLibrary(query, contentType, page)
        .then((data) => {
          this.totalResults = data.total_results;
          this.results = data.results;
          this.perPage = data.per_page;
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
  mounted() {
    let q = this.$route.query.q;
    let contentType = this.$route.query.content_type;
    let page = parseInt(this.$route.query.page);

    q = q?q:'';
    contentType = contentType?contentType:'';
    page = page?page:1;
    
    this.query = q;
    this.contentType = contentType;
    this.currentPage = page;

    console.log('searching with', q, contentType, page)
    this.search(q, contentType, page);
  },
};
</script>
