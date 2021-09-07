<template>
  <div class="md:flex flex-col md:flex-row justify-between items-center">
    <o-field label="Search" class="flex-1">
      <o-input
        @change="search"
        type="text"
        v-model="computedQuery"
        icon-right="times-circle"
        icon-right-clickable
        @icon-right-click="clear"
        placeholder="Title or IMDB id"
        ref="searchbar"
      />
    </o-field>
    <div class="flex my-2 md:my-0 self-end md:ml-4 justify-center">
      <o-button variant="primary" @click="search">Search</o-button>
    </div>
    <div class="md:my-0 my-2 flex justify-center self-end">
      <div class="md:ml-4 overflow-hidden rounded-lg inline-block">
        <o-radio v-model="contentType" name="contentType" native-value=""
          >All</o-radio
        >
        <o-radio v-model="contentType" name="contentType" native-value="movie"
          >Movies</o-radio
        >
        <o-radio v-model="contentType" name="contentType" native-value="series"
          >TV Series</o-radio
        >
      </div>
    </div>
  </div>
  <div class="mt-4">
    <o-pagination
      :total="totalResults"
      v-model:current="current"
      :range-before="3"
      :range-after="3"
      order="default"
      size="medium"
      :per-page="10"
      aria-next-label="Next page"
      aria-previous-label="Previous page"
      aria-page-label="Page"
      aria-current-label="Current page"
    />
  </div>
  <div v-if="results && results.length > 0" :class="resultsWrapperClass">
    <div v-for="(media, index) in results" :key="index">
      <slot name="result" :media="media" />
    </div>
  </div>
  <slot name="empty" v-else />
  <!-- <o-loading
    :full-screen="false"
    :can-cancel="false"
    v-model:active="loading"
  /> -->
</template>

<script>
import { nextTick } from "vue";

export default {
  data() {
    return {
      contentType: "",
      page: 1,
      current: 1,
      lastSearchTime: null,
    };
  },
  props: {
    mediaType: {
      type: String,
      default: function () {
        return "";
      },
    },
    results: {
      type: Array,
      default: function () {
        return [];
      },
    },
    loading: {
      type: Boolean,
      default: function () {
        return false;
      },
    },
    totalResults: {
      type: Number,
      default: function () {
        return 0;
      },
    },
    resultsWrapperClass: {
      type: String,
      default: function () {
        return "";
      },
    },
    query: {
      type: String,
      default: function () {
        return "";
      },
    },
  },
  emits: ["close", "search", "selected-media", "update:query"],
  methods: {
    search() {
      const now = new Date();
      if (
        !this.lastSearchTime ||
        now.getTime() - this.lastSearchTime.getTime() > 300
      ) {
        this.$emit("search", this.computedQuery, this.contentType, this.page);
        this.lastSearchTime = now;
      }
    },
    clear() {
      this.computedQuery = "";
      this.search();
    },
  },
  mounted() {
    const screenWidth = window.innerWidth;
    if (screenWidth >= 768) {
      nextTick(() => {
        this.$refs.searchbar.$el.firstChild.focus();
      });
    }
    this.search();
  },
  watch: {
    contentType: function () {
      this.page = 0;
      this.search();
    },
    current: function (newVal) {
      this.page = newVal;
      this.search();
    },
    computedQuery: function (newVal, oldVal) {
      this.page = 1;
      if (newVal === "" && oldVal != "") {
        this.search();
      }
    },
  },
  computed: {
    computedQuery: {
      get() {
        return this.query;
      },
      set(q) {
        this.$emit("update:query", q);
      },
    },
  },
};
</script>
