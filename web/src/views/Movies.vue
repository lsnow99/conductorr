<template>
    <page-wrapper>
    <section>
      <div class="flex justify-between">
        <o-field label="Search Media" class="w-96">
          <o-input @change="search" v-model="query" type="text" placeholder="Search your library"/>
        </o-field>
      </div>
      <div
        v-if="results && results.length > 0"
        class="
          mt-10
          gap-y-10
          grid
          justify-items-center
          grid-flow-row grid-cols-1
          sm:grid-cols-2
          lg:grid-cols-3
          xl:grid-cols-4
          2xl:grid-cols-5
        "
      >
        <media-card v-for="(media, index) in results" :key="index" :media="media" />
      </div>
      <div v-else class="flex justify-center">
        <div class="flex flex-col mt-24">
          <span class="text-xl text-center">No results in your library</span>
          <div class="flex flex-row justify-center mt-2">
            <o-button variant="primary" @click="addNew">Search and Add</o-button>
          </div>
        </div>
      </div>
      <o-loading v-model:active="loading" :full-page="false" :can-cancel="false"></o-loading>
    </section>
    <o-modal full-screen v-model:active="showNewSearchModal">
        <new-media @close="showNewSearchModal = false"/>
    </o-modal>
    </page-wrapper>
</template>

<script>
import MediaCard from "../components/MediaCard.vue";
import PageWrapper from "../components/PageWrapper.vue";
import NewMedia from "../components/NewMedia.vue";
import APIUtil from '../util/APIUtil';

export default {
  data() {
    return {
      results: [],
      total_results: 0,
      query: '',
      loading: false,
      showNewSearchModal: false,
    };
  },
  components: { PageWrapper, MediaCard, NewMedia },
  methods: {
    search() {
      this.loading = true
      APIUtil.searchLibrary(this.query, 'movies', 1).then((data) => {
        this.total_results = data.total_results
        this.results = data.results
      }).finally(() => {
        this.loading = false
      })
    },
    addNew() {
      this.showNewSearchModal = true
    }
  },
  watch: {
    restrict: function() {
      this.search()
    }
  }
};
</script>
