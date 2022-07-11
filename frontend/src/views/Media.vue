<template>
  <page-wrapper v-if="media">
    <section class="flex flex-row">
      <img
        v-if="media.poster"
        class="hidden rounded-md md:block w-72"
        :src="media.poster"
        :alt="`Poster for ${media.title}`"
      />
      <div class="flex flex-col flex-1 ml-4" v-show="media.id">
        <h1 class="text-4xl lg:text-6xl">
          <monitoring-toggle
            class="text-3xl lg:text-5xl"
            :monitoring="media.monitoring"
            @toggle="toggleMonitoring(media!)"
          />
          {{ media.title }}
        </h1>
        <div class="flex flex-row text-xl">
          <div
            class="w-1 h-full"
            :class="media.pathOk ? 'bg-green-600' : 'bg-red-600'"
          />
          <span class="ml-2">{{ media.path }}</span>
        </div>
        <div
          class="flex flex-col justify-between px-4 py-4 lg:flex-row lg:items-center"
        >
          <div class="flex flex-row justify-evenly lg:items-center">
            <div class="text-2xl text-gray-300 lg:mr-4">
              {{ mediaYear(media) }}
            </div>
            <div class="text-2xl text-gray-300 lg:mx-4">
              <o-icon class="text-lg" icon="star" />
              {{`${media.imdbRating}%`}}
            </div>
            <a
              :href="`https://www.imdb.com/title/${media.imdbId}`"
              target="_blank"
              class="inline-block pt-1 lg:mx-4"
              :aria-label="`Link to IMDB page for ${media.title}`"
            >
              <o-icon
                class="text-4xl text-gray-300 hover:text-yellow-300"
                pack="fab"
                icon="imdb"
              />
            </a>
          </div>
          <div class="flex flex-row justify-evenly lg:self-end">
            <o-tooltip
              variant="info"
              :position="tooltipPosition"
              label="Delete Media"
            >
              <div class="text-2xl text-gray-300 lg:mx-2">
                <div
                  @click="showConfirmDeleteModal = true"
                  @keydown.space="showConfirmDeleteModal = true"
                  @keydown.enter="showConfirmDeleteModal = true"
                  role="button"
                  aria-label="Delete media"
                  tabindex="0"
                >
                  <o-icon class="cursor-pointer" icon="trash" />
                </div>
              </div>
            </o-tooltip>
            <o-tooltip
              variant="info"
              :position="tooltipPosition"
              label="Edit Media"
            >
              <div class="text-2xl text-gray-300 lg:mx-2">
                <div
                  @click="editMedia"
                  @keydown.space="editMedia"
                  @keydown.enter="editMedia"
                  tabindex="0"
                  role="button"
                  aria-label="Edit media"
                >
                  <o-icon class="cursor-pointer" icon="wrench" />
                </div>
              </div>
            </o-tooltip>
            <o-tooltip
              variant="info"
              :position="tooltipPosition"
              label="Refresh Metadata"
            >
              <div class="text-2xl text-gray-300 lg:mx-2">
                <div
                  v-if="!loadingRefreshMetadata"
                  @click="refreshMediaMetadata"
                  @keydown.space="refreshMediaMetadata"
                  @keydown.enter="refreshMediaMetadata"
                  tabindex="0"
                  role="button"
                  aria-label="Refresh Metadata"
                >
                  <o-icon class="cursor-pointer" icon="sync-alt" />
                </div>
                <o-icon v-else icon="sync-alt" spin />
              </div>
            </o-tooltip>
            <search-actions :mediaId="mediaId" size="large" />
          </div>
        </div>
        <p class="text-lg">{{ media.description }}</p>
      </div>
    </section>
    <section class="mt-4">
      <DownloadStatusViewer wrapperClass="h-48" :mediaId="mediaId" />
    </section>
    <section class="mt-4">
      <div
        class="p-5 my-4 bg-gray-600 rounded-md cursor-pointer"
        v-for="season in media.children"
        :key="season.id"
        @click="expand(season.id)"
        @keydown.enter="expand(season.id)"
        @keydown.space="expand(season.id)"
        tabindex="0"
        role="button"
        :aria-label="`Expand ${season.title}`"
      >
        <div class="flex flex-row justify-between text-2xl">
          <div>
            <monitoring-toggle
              class="text-xl"
              :monitoring="season.monitoring"
              :disabled="!media.monitoring"
              @toggle="toggleMonitoring(season)"
            />
            {{ season.title }}
          </div>
          <div class="text-base" @click.prevent @click.stop>
            <search-actions :mediaId="season.id" size="large" />
          </div>
        </div>
        <transition name="fade">
          <div
            @click.prevent
            @click.stop
            class="p-1 bg-gray-800 rounded-md cursor-default"
            v-show="expandedCfg[season.id]"
          >
            <episode-list
              :load="loadedCfg[season.id]"
              @reload="loadMedia"
              :monitoring-disabled="!media.monitoring"
              :episodes="season.children"
            />
          </div>
        </transition>
        <div class="mt-2 text-center">
          <o-icon
            v-if="loadingCfg[season.id]"
            size="large"
            icon="hourglass-start"
          />
          <o-icon
            v-if="!loadingCfg[season.id] && expandedCfg[season.id]"
            size="large"
            icon="chevron-up"
          />
          <o-icon
            v-if="!loadingCfg[season.id] && !expandedCfg[season.id]"
            size="large"
            icon="chevron-down"
          />
        </div>
      </div>
    </section>
    <edit-media
      v-model:active="showEditMediaModal"
      @submit="updateMedia"
      :media="media"
      @close="closeEditMedia"
    />
    <ConfirmDelete
      v-model:active="showConfirmDeleteModal"
      @delete="doDelete"
      @close="closeDelete"
      :delete-message="`Are you sure you want to delete '${media.title}'?`"
    />
    <o-loading v-model:active="loading" is-full-page />
  </page-wrapper>
</template>

<script setup lang="ts">
import PageWrapper from "../components/PageWrapper.vue";
import APIUtil from "../util/APIUtil";
import EditMedia from "../components/EditMedia.vue";
import ConfirmDelete from "../components/ConfirmDelete.vue";
import EpisodeList from "../components/EpisodeList.vue";
import SearchActions from "../components/SearchActions.vue";
import MonitoringToggle from "../components/MonitoringToggle.vue";
import DownloadStatusViewer from "../components/DownloadStatusViewer.vue";
import { nextTick, onMounted, ref } from "vue";
import { Media } from "@/types/api/media";
import useTabSaver from "@/util/TabSaver";
import useMediaUtil from "@/util/MediaUtil";
import { useRouter, useRoute } from "vue-router";

const route = useRoute();

const media = ref<Media | null>(null);
const mediaId = ref(parseInt(route.params.media_id as string));
const loadingRefreshMetadata = ref(false);
const loading = ref(true);
const tooltipPosition = ref("bottom");
const showEditMediaModal = ref(false);
const showConfirmDeleteModal = ref(false);
const expandedCfg = ref<{ [key: number]: boolean }>({});
const loadedCfg = ref<{ [key: number]: boolean }>({});
const loadingCfg = ref<{ [key: number]: boolean }>({});

const { lastButton, restoreFocus } = useTabSaver();
const { mediaYear } = useMediaUtil();

const router = useRouter();

const closeDelete = () => {
  restoreFocus();
  showConfirmDeleteModal.value = false;
};
const closeEditMedia = () => {
  restoreFocus();
  showEditMediaModal.value = false;
};
const loadMedia = async() => {
  try {
    media.value = await APIUtil.getMedia(mediaId.value)
  } finally {
    loading.value  = false
  }
};
const toggleMonitoring = (media: Media) => {
  APIUtil.setMonitoringMedia(media.id, !media.monitoring).then(() => {
    loadMedia();
  });
};
const editMedia = ($event: Event) => {
  lastButton.value = $event.currentTarget as HTMLElement;
  showEditMediaModal.value = true;
};
const updateMedia = ({
  profileID,
  pathID,
}: {
  profileID: number;
  pathID: number;
}) => {
  APIUtil.updateMedia(mediaId.value, profileID, pathID).then(() => {
    loadMedia();
    showEditMediaModal.value = false;
  });
};
const doDelete = () => {
  APIUtil.deleteMedia(mediaId.value).then(() => {
    router.push({ name: "library" });
  });
};
const refreshMediaMetadata = () => {
  loadingRefreshMetadata.value = true;
  APIUtil.refreshMediaMetadata(mediaId)
    .then(() => {
      loadMedia();
    })
    .finally(() => {
      loadingRefreshMetadata.value = false;
    });
};
const expand = (id: number) => {
  expandedCfg.value[id] = !expandedCfg.value[id];
  loadingCfg.value[id] = true;
  nextTick(() => {
    loadedCfg.value[id] = true;
    loadingCfg.value[id] = false;
  });
};

onMounted(() => {
  loadMedia();
  if (window.innerWidth < 768) {
    tooltipPosition.value = "left";
  }
});
</script>
