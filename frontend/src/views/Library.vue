<template>
  <page-wrapper>
    <section>
      <SearchMedia
        v-model:query="query"
        v-model:currentPage="currentPage"
        v-model:contentType="contentType"
        :results="results"
        :total-results="totalResults"
        :per-page="perPage"
        :loading="loading"
        @search="search"
        @selected-media="selectedMedia"
        results-wrapper-class="flex flex-row flex-wrap justify-center pb-10 mt-10 gap-y-10 gap-x-10"
      >
        <template v-slot:empty class="flex justify-center">
          <div class="flex flex-col mt-24">
            <span class="text-xl text-center">No results in your library</span>
          </div>
        </template>
        <template v-slot:result="{ media }">
          <media-card :media="media" @click="selectedMedia" />
        </template>
      </SearchMedia>
      <div class="flex flex-row justify-center mt-2">
        <o-button variant="primary" @click="addNew"
          >Search for New Media</o-button
        >
      </div>
      <div class="flex flex-col mt-2 text-lg md:flex-row">
        <div class="flex flex-row items-center mx-4">
          <div class="w-6 h-2 mr-1 bg-green-600" />
          Complete/All Monitored Available
        </div>
        <div class="flex flex-row items-center mx-4">
          <div class="w-6 h-2 mr-1 bg-purple-600" />
          Continuing/All Monitored Available
        </div>
        <div class="flex flex-row items-center mx-4">
          <div class="w-6 h-2 mr-1 bg-gray-500" />
          Unmonitored
        </div>
        <div class="flex flex-row items-center mx-4">
          <div class="w-6 h-2 mr-1 bg-red-600" />
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

<script setup lang="ts">
import MediaCard from "../components/MediaCard.vue";
import PageWrapper from "../components/PageWrapper.vue";
import NewMedia from "../components/NewMedia.vue";
import APIUtil from "../util/APIUtil";
import SearchMedia from "../components/SearchMedia.vue";
import { onMounted, ref } from "vue";
import { ContentType, Media, MediaSearchResult } from "@/types/api/media";
import { safeParseInt, useTabSaver } from "@/util";
import { LocationQueryValue, useRoute, useRouter } from "vue-router";

const query = ref("")
const results = ref<MediaSearchResult[]>([])
const totalResults = ref(0)
const perPage = ref(0)
const currentPage = ref(1)
const contentType = ref<ContentType>(ContentType.ALL)
const loading = ref(false)
const showNewSearchModal = ref(false)

const { lastButton, restoreFocus } = useTabSaver()
const route = useRoute()
const router = useRouter()

const search = async(q: string, contentType: ContentType, page: number) => {
  router.replace({
    name: "library",
    query: {q, content_type: contentType, page}
  })
  loading.value = true
  try {
    const data = await APIUtil.searchLibrary(q, contentType, page)
    totalResults.value = data.totalResults
    results.value = data.results
    perPage.value = data.per_page
  } catch (error) {
    totalResults.value = 0
    results.value = []
  } finally {
    loading.value = false
  }
}

const selectedMedia = (media: Media) => {
  router.push({name: "media", params: {media_id: media.id}})
}

const addNew = ($event: Event) => {
  showNewSearchModal.value = true;
  lastButton.value = $event.currentTarget as HTMLElement
}

const closeNewSearchModal = () => {
  showNewSearchModal.value = false;
  restoreFocus()
}

onMounted(() => {
  query.value = route.query.q as LocationQueryValue ?? '';
  contentType.value = route.query.content_type === "" ? ContentType.ALL : route.query.content_type as ContentType;
  currentPage.value = safeParseInt(`${route.query.page}`) ?? 1;

  search(query.value, contentType.value, currentPage.value);
})
</script>
