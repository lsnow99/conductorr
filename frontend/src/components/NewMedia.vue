<template>
  <Modal
    title="New Media"
    v-model="computedActive"
    @close="emit('close')"
    full-screen
  >
    <SearchMedia
      v-model:query="query"
      v-model:currentPage="currentPage"
      v-model:contentType="contentType"
      :results="results"
      :total-results="totalResults"
      :loading="loading"
      :per-page="perPage"
      results-wrapper-class="flex flex-row flex-wrap justify-center mt-10 gap-y-10 gap-x-10"
      @search="search"
    >
      <template v-slot:empty class="flex justify-center">
        <div class="flex flex-col mt-24">
          <span class="text-xl text-center">No results</span>
        </div>
      </template>
      <template v-slot:result="{ media }">
        <MediaCard
          :media="media"
          @click="(m, $e) => selectedMedia(m as MediaSearchResult, $e)"
        />
      </template>
    </SearchMedia>
  </Modal>
  <EditMedia
    v-if="media"
    v-model:active="showNewMediaModal"
    :loading="loadingNewMedia"
    @submit="addMedia"
    :media="media"
    @close="closeNewMedia"
  />
</template>

<script setup lang="ts">
import APIUtil from "../util/APIUtil";
import MediaCard from "../components/MediaCard.vue";
import SearchMedia from "./SearchMedia.vue";
import EditMedia from "./EditMedia.vue";
import useTabSaver from "../util/TabSaver";
import Modal from "./Modal.vue";
import { ContentType, MediaSearchResult } from "@/types/api/media";
import { computed, ref, watch, WritableComputedRef } from "vue";
import { useRouter } from "vue-router";

const results = ref([]);
const totalResults = ref(0);
const perPage = ref(0);
const contentType = ref<ContentType>(ContentType.ALL);
const currentPage = ref(1);
const loading = ref(false);
const query = ref("");
const media = ref<MediaSearchResult | null>(null);
const showNewMediaModal = ref(false);
const loadingNewMedia = ref(false);

const props = defineProps<{
  defaultQuery: string;
  active: boolean;
}>();

const { lastButton, restoreFocus } = useTabSaver();
const router = useRouter();

const emit = defineEmits<{
  (e: "close"): void;
  (e: "update:active", active: boolean): void;
}>();

const search = async (q: string, ct: ContentType, p: number) => {
  console.log('searching')
  loading.value = true;
  try {
    const data = await APIUtil.searchNew(q, ct, p);
    totalResults.value = data.total_results;
    results.value = data.results;
    perPage.value = data.per_page;
  } catch {
    totalResults.value = 0;
    results.value = [];
  } finally {
    loading.value = false;
  }
};

const selectedMedia = (m: MediaSearchResult, $event: Event) => {
  showNewMediaModal.value = true;
  media.value = m;
  lastButton.value = $event.currentTarget as HTMLElement;
};

const addMedia = async ({
  profileID,
  pathID,
}: {
  profileID: number;
  pathID: number;
}) => {
  if (!media.value) {
    return;
  }
  loadingNewMedia.value = true;
  try {
    const id = await APIUtil.addMedia(media.value.searchID, profileID, pathID);
    router.push({ name: "media", params: { media_id: id } });
  } catch (err) {
    console.log(err);
  } finally {
    loadingNewMedia.value = false;
  }
};

const closeNewMedia = () => {
  showNewMediaModal.value = false;
  restoreFocus();
};

watch(
  () => props.defaultQuery,
  (v: string) => {
    query.value = v;
  },
  { immediate: true }
);

const computedActive: WritableComputedRef<boolean> = computed({
  get(): boolean {
    return props.active;
  },
  set(v: boolean): void {
    emit("update:active", v);
  },
});
</script>
