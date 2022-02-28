<template>
  <page-wrapper>
    <section class="flex flex-row">
      <img
        v-if="media.poster"
        class="hidden md:block rounded-md w-72"
        :src="media.poster"
        :alt="`Poster for ${media.title}`"
      />
      <div class="ml-4 flex flex-1 flex-col" v-show="media.id">
        <h1 class="text-4xl lg:text-6xl">
          <monitoring-toggle
            class="text-3xl lg:text-5xl"
            :monitoring="media.monitoring"
            @toggle="toggleMonitoring(media)"
          />
          {{ media.title }}
        </h1>
        <div class="text-xl flex flex-row">
          <div
            class="w-1 h-full"
            :class="media.path_ok ? 'bg-green-600' : 'bg-red-600'"
          />
          <span class="ml-2">{{ media.path }}</span>
        </div>
        <div
          class="
            py-4
            flex flex-col
            lg:flex-row lg:items-center
            justify-between
            px-4
          "
        >
          <div class="flex flex-row justify-evenly lg:items-center">
            <div class="text-2xl lg:mr-4 text-gray-300">
              {{ mediaYear(media) }}
            </div>
            <div class="text-2xl lg:mx-4 text-gray-300">
              <o-icon class="text-lg" icon="star" />
              {{ media.imdb_rating }}%
            </div>
            <a
              :href="`https://www.imdb.com/title/${media.imdb_id}`"
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
              <div class="text-2xl lg:mx-2 text-gray-300">
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
              <div class="text-2xl lg:mx-2 text-gray-300">
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
              <div class="text-2xl lg:mx-2 text-gray-300">
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
            <search-actions :mediaID="mediaID" size="large" />
          </div>
        </div>
        <p class="text-lg">{{ media.description }}</p>
      </div>
    </section>
    <section class="mt-4">
      <DownloadStatusViewer wrapperClass="h-48" :mediaID="mediaID" />
    </section>
    <section class="mt-4">
      <div
        class="p-5 bg-gray-600 cursor-pointer rounded-md my-4"
        v-for="season in media.children"
        :key="season.id"
        @click="expand(season.id)"
        @keydown.enter="expand(season.id)"
        @keydown.space="expand(season.id)"
        tabindex="0"
        role="button"
        :aria-label="`Expand ${season.title}`"
      >
        <div class="text-2xl flex flex-row justify-between">
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
            <search-actions :mediaID="season.id" size="large" />
          </div>
        </div>
        <transition name="fade">
          <div
            @click.prevent
            @click.stop
            class="bg-gray-800 rounded-md p-1 cursor-default"
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
        <div class="text-center mt-2">
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

<script>
import PageWrapper from "../components/PageWrapper.vue";
import APIUtil from "../util/APIUtil";
import MediaUtil from "../util/MediaUtil";
import EditMedia from "../components/EditMedia.vue";
import ConfirmDelete from "../components/ConfirmDelete.vue";
import ManualSearchResults from "../components/ManualSearchResults.vue";
import EpisodeList from "../components/EpisodeList.vue";
import SearchActions from "../components/SearchActions.vue";
import MonitoringToggle from "../components/MonitoringToggle.vue";
import TabSaver from "../util/TabSaver";
import DownloadStatusViewer from "../components/DownloadStatusViewer.vue";

export default {
  data() {
    return {
      media: {},
      mediaID: 0,
      releases: [],
      loadingManualSearch: false,
      loadingAutoSearch: false,
      loadingRefreshMetadata: false,
      loading: true,
      tooltipPosition: "bottom",
      showEditMediaModal: false,
      showConfirmDeleteModal: false,
      showManualReleasesModal: false,
      expandedCfg: {},
      loadedCfg: {},
      loadingCfg: {},
    };
  },
  mixins: [MediaUtil, TabSaver],
  components: {
    PageWrapper,
    EditMedia,
    ConfirmDelete,
    ManualSearchResults,
    EpisodeList,
    SearchActions,
    MonitoringToggle,
    DownloadStatusViewer,
  },
  methods: {
    closeDelete() {
      this.restoreFocus();
      this.showConfirmDeleteModal = false;
    },
    closeEditMedia() {
      this.restoreFocus();
      this.showEditMediaModal = false;
    },
    toggleMonitoring(media) {
      APIUtil.setMonitoringMedia(media.id, !media.monitoring).then(() => {
        this.loadMedia();
      });
    },
    editMedia($event) {
      this.lastButton = $event.currentTarget;
      this.showEditMediaModal = true;
    },
    updateMedia({ profileID, pathID }) {
      APIUtil.updateMedia(this.mediaID, profileID, pathID).then(() => {
        this.loadMedia();
        this.showEditMediaModal = false;
      });
    },
    doDelete() {
      APIUtil.deleteMedia(this.mediaID).then(() => {
        this.$router.push({ name: "library" });
      });
    },
    refreshMediaMetadata() {
      this.loadingRefreshMetadata = true;
      APIUtil.refreshMediaMetadata(this.mediaID)
        .then(() => {
          this.loadMedia();
        })
        .finally(() => {
          this.loadingRefreshMetadata = false;
        });
    },
    loadMedia() {
      APIUtil.getMedia(this.mediaID)
        .then((media) => {
          this.media = media;
        })
        .finally(() => {
          this.loading = false;
        });
    },
    expand(id) {
      this.expandedCfg[id] = !this.expandedCfg[id];
      this.loadingCfg[id] = true;
      setTimeout(() => {
        this.loadedCfg[id] = true;
        this.loadingCfg[id] = false;
      }, 0);
    },
  },
  created() {
    this.mediaID = parseInt(this.$route.params.media_id);
  },
  mounted() {
    this.loadMedia();
    const screenWidth = window.innerWidth;
    if (screenWidth < 768) {
      this.tooltipPosition = "left";
    }
  },
};
</script>
