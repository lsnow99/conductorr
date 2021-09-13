<template>
  <modal
    title="New Media"
    v-model="computedActive"
    @close="$emit('close')"
    full-screen
  >
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
  </modal>
  <edit-media
    v-model:active="showNewMediaModal"
    :loading="loadingNewMedia"
    @submit="addMedia"
    :media="media"
    @close="closeNewMedia"
  />
</template>

<script>
import APIUtil from "../util/APIUtil";
import MediaCard from "../components/MediaCard.vue";
import SearchMedia from "./SearchMedia.vue";
import EditMedia from "./EditMedia.vue";
import TabSaver from "../util/TabSaver";
import Modal from "./Modal.vue";

export default {
  data() {
    return {
      results: [],
      totalResults: 0,
      loading: false,
      query: "",
      media: {},
      showNewMediaModal: false,
      loadingNewMedia: false,
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
    active: {
      type: Boolean,
      default: function () {
        return false;
      },
    },
  },
  mixins: [TabSaver],
  components: { MediaCard, SearchMedia, EditMedia, Modal },
  emits: ["close", "update:active"],
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
    selectedMedia(media, $event) {
      this.showNewMediaModal = true;
      this.media = media;
      this.lastButton = $event.currentTarget;
    },
    addMedia({ profileID, pathID }) {
      this.loadingNewMedia = true;
      APIUtil.addMedia(this.media.imdb_id, profileID, pathID)
        .then((id) => {
          this.$router.push({ name: "media", params: { media_id: id } });
        })
        .catch((err) => {
          console.log(err);
        })
        .finally(() => {
          this.loadingNewMedia = false;
        });
    },
    closeNewMedia() {
      this.showNewMediaModal = false;
      this.restoreFocus();
    },
  },
  watch: {
    defaultQuery: {
      handler(v) {
        if (v) {
          this.query = v;
        }
      },
      immediate: true
    }
  },
  computed: {
    computedActive: {
      get() {
        return this.active;
      },
      set(v) {
        this.$emit("update:active", v);
      },
    },
  },
};
</script>
