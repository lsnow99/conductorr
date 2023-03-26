<template>
  <div class="flex-col items-center justify-between md:flex md:flex-row">
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
    <div class="flex self-end justify-center my-2 md:my-0 md:ml-4">
      <o-button variant="primary" @click="search">Search</o-button>
    </div>
    <div class="flex self-end justify-center my-2 md:my-0 md:ml-4">
      <radio-group
        v-model="computedContentType"
        name="contentType"
        :options="[
          {
            text: 'All',
            value: ContentType.ALL,
          },
          {
            text: 'Movies',
            value: ContentType.MOVIE,
          },
          {
            text: 'TV Series',
            value: ContentType.SERIES,
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

<script setup lang="ts">
import {
  computed,
  nextTick,
  onMounted,
  ref,
  watch,
  WritableComputedRef,
} from "vue";
import RadioGroup from "./RadioGroup.vue";
import { ContentType, MediaSearchResult } from "@/types/api/media";

const page = ref(1);
const lastSearchTime = ref<Date | null>(null);
const searchbar = ref<Element | null>(null);
const lastSearchId = ref<string>("");

const props = withDefaults(
  defineProps<{
    currentPage: number;
    contentType: ContentType;
    results: MediaSearchResult[];
    loading: boolean;
    totalResults: number;
    perPage: number;
    resultsWrapperClass: string;
    query: string;
  }>(),
  {
    currentPage: 1,
  }
);

const emit = defineEmits<{
  (e: "close"): void;
  (e: "search", query: string, contentType: ContentType, page: number): void;
  (e: "update:query", query: string): void;
  (e: "update:currentPage", currentPage: number): void;
  (e: "update:contentType", contentType: ContentType | null): void;
}>();

const computedQuery: WritableComputedRef<string> = computed({
  get(): string {
    return props.query;
  },
  set(v: string) {
    emit("update:query", v);
  },
});

const computedCurrentPage: WritableComputedRef<number> = computed({
  get() {
    return props.currentPage;
  },
  set(v) {
    emit("update:currentPage", v);
  },
});

const computedContentType: WritableComputedRef<ContentType> = computed({
  get() {
    return props.contentType;
  },
  set(v: ContentType) {
    emit("update:contentType", v);
  },
});

const search = (disableDebounce = false) => {
  const now = new Date();
  if (
    (!lastSearchTime.value ||
      now.getTime() - lastSearchTime.value.getTime() > 300 ||
      disableDebounce) &&
    computedContentType.value
  ) {
    emit("search", computedQuery.value, computedContentType.value, page.value);
    lastSearchTime.value = now;
  } else if !disableDebounce {
    const currentId = crypto.randomUUID()
    lastSearchId.value = currentId
    setTimeout(() => {
      if(lastSearchId.value === currentId) {
        search(true)
      }
    }, 300)
  }
};

const clear = () => {
  computedQuery.value = "";
  search();
};

const pageChanged = (nPage: number) => {
  page.value = nPage;
  search();
};

onMounted(() => {
  const screenWidth = window.innerWidth;
  if (screenWidth >= 768) {
    nextTick(() => {
      (searchbar.value?.firstChild as HTMLInputElement | undefined)?.focus();
    });
  }
  search();
});

watch(computedContentType, () => {
  page.value = 1;
  computedCurrentPage.value = 1;
  search(true);
});

watch(computedQuery, (newVal, oldVal) => {
  page.value = 1;
  computedCurrentPage.value = 1;
  if (newVal === "" && oldVal !== "") {
    search();
  }
});
</script>
