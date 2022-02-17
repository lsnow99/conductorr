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
    <div class="md:my-0 md:ml-4 my-2 flex justify-center self-end">
      <radio-group
        v-model="computedContentType"
        name="contentType"
        :options="[
          {
            text: 'All',
            value: '',
          },
          {
            text: 'Movies',
            value: 'movie',
          },
          {
            text: 'TV Series',
            value: 'series',
          },
        ]"
      />
    </div>
  </div>
  <div class="mt-4">
    <o-pagination
      v-model:current="computedCurrentPage"
      :total="totalResults"
      :per-page="perPage"
      :range-before="3"
      :range-after="3"
      order="default"
      size="medium"
      aria-next-label="Next page"
      aria-previous-label="Previous page"
      aria-page-label="Page"
      aria-current-label="Current page"
      @change="pageChanged"
    />
  </div>
  <div v-if="results && results.length > 0" :class="resultsWrapperClass">
    <div v-for="(media, index) in results" :key="index">
      <slot name="result" :media="media" />
    </div>
  </div>
  <slot name="empty" v-else />
</template>

<script>
import { nextTick } from "vue";
import RadioGroup from "./RadioGroup.vue";

export default {
  data() {
    return {
      page: 1,
      lastSearchTime: null,
    };
  },
  props: {
    currentPage: {
      type: Number,
      default: function() {
        return 1
      }
    },
    contentType: {
      type: String,
      default: function() {
        return ''
      }
    },
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
    perPage: {
      type: Number,
      default: function() {
        return 0;
      }
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
  components: {
    RadioGroup,
  },
  emits: ["close", "search", "selected-media", "update:query", "update:currentPage", "update:contentType"],
  methods: {
    search(disableDebounce) {
      const now = new Date();
      if (
        !this.lastSearchTime ||
        (now.getTime() - this.lastSearchTime.getTime() > 300 ||
          disableDebounce)
      ) {
        this.$emit("search", this.computedQuery, this.computedContentType, this.page);
        this.lastSearchTime = now;
      }
    },
    clear() {
      this.computedQuery = "";
      this.search();
    },
    pageChanged(nPage) {
      this.page = nPage
      this.search();
    }
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
    computedContentType: function () {
      this.page = 1;
      this.computedCurrentPage = 1;
      this.search(true);
    },
    computedQuery: function (newVal, oldVal) {
      this.page = 1;
      this.computedCurrentPage = 1;
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
      set(v) {
        this.$emit("update:query", v);
      },
    },
    computedCurrentPage: {
      get() {
        return this.currentPage;
      },
      set(v) {
        this.$emit("update:currentPage", v)
      }
    },
    computedContentType: {
      get() {
        return this.contentType;
      },
      set(v) {
        this.$emit("update:contentType", v)
      }
    }
  },
};
</script>
