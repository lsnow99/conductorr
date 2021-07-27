<template>
  <page-wrapper>
    <section class="flex flex-row">
      <img class="hidden md:block rounded-md" :src="media.poster" />
      <div class="ml-4 flex flex-1 flex-col" v-show="media.id">
        <h1 class="text-4xl lg:text-6xl">
          <monitoring-toggle
            class="text-3xl lg:text-5xl"
            :monitoring="media.monitoring"
            @toggle="toggleMonitoring(media)"
          />
          {{ media.title }}
        </h1>
        <div
          class="
            py-4
            flex flex-col
            sm:flex-row sm:items-center
            justify-between
            px-4
          "
        >
          <div class="flex flex-row justify-evenly md:items-center">
            <div class="text-2xl md:mr-4 text-gray-300">
              {{ mediaYear(media) }}
            </div>
            <div class="text-2xl md:mx-4 text-gray-300">
              <o-icon class="text-lg" icon="star" />
              {{ media.imdb_rating }}%
            </div>
            <a
              :href="`https://www.imdb.com/title/${media.imdb_id}`"
              target="_blank"
              class="inline-block pt-1 md:mx-4"
            >
              <o-icon
                class="text-4xl text-gray-300 hover:text-yellow-300"
                pack="fab"
                icon="imdb"
              />
            </a>
          </div>
          <div class="flex flex-row justify-evenly md:self-end">
            <o-tooltip
              variant="info"
              :position="tooltipPosition"
              label="Delete Media"
            >
              <div class="text-2xl md:mx-2 text-gray-300">
                <div @click="showConfirmDeleteModal = true">
                  <o-icon class="cursor-pointer" icon="trash" />
                </div>
              </div>
            </o-tooltip>
            <o-tooltip
              variant="info"
              :position="tooltipPosition"
              label="Edit Media"
            >
              <div class="text-2xl md:mx-2 text-gray-300">
                <div @click="editMedia">
                  <o-icon class="cursor-pointer" icon="wrench" />
                </div>
              </div>
            </o-tooltip>
            <search-actions :mediaID="mediaID" size="large" />
          </div>
        </div>
        <p class="text-lg">{{ media.description }}</p>
      </div>
    </section>
    <section class="mt-4">
      <div
        class="p-5 bg-gray-600 cursor-pointer rounded-md my-4"
        v-for="season in media.children"
        :key="season.id"
        @click="expandedCfg[season.id] = !expandedCfg[season.id]"
      >
        <div class="text-2xl">
          <monitoring-toggle
            @click.stop
            class="text-xl"
            :monitoring="season.monitoring"
            :disabled="!media.monitoring"
            @toggle="toggleMonitoring(season)"
          />
          {{ season.title }}
        </div>
        <transition name="fade">
          <div
            @click.prevent
            @click.stop
            class="bg-gray-800 rounded-md p-1 cursor-default"
            v-show="expandedCfg[season.id]"
          >
            <episode-list @reload="loadMedia" :monitoring-disabled="!season.monitoring || !media.monitoring" :episodes="season.children" />
          </div>
        </transition>
        <div class="text-center mt-2">
          <o-icon
            v-if="expandedCfg[season.id]"
            size="large"
            icon="chevron-up"
          />
          <o-icon v-else size="large" icon="chevron-down" />
        </div>
      </div>
    </section>
    <o-modal
      v-model:active="showEditMediaModal"
      @close="showEditMediaModal = false"
    >
      <edit-media
        @submit="updateMedia"
        :media="media"
        @close="showEditMediaModal = false"
      />
    </o-modal>
    <o-modal
      v-model:active="showConfirmDeleteModal"
      @close="showConfirmDeleteModal = false"
    >
      <ConfirmDelete
        @delete="doDelete"
        @close="showConfirmDeleteModal = false"
        :delete-message="`Are you sure you want to delete '${media.title}'?`"
      />
    </o-modal>
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

export default {
  data() {
    return {
      media: {},
      mediaID: 0,
      releases: [],
      loadingManualSearch: false,
      loadingAutoSearch: false,
      loading: true,
      tooltipPosition: "bottom",
      showEditMediaModal: false,
      showConfirmDeleteModal: false,
      showManualReleasesModal: false,
      expandedCfg: {},
    };
  },
  mixins: [MediaUtil],
  components: {
    PageWrapper,
    EditMedia,
    ConfirmDelete,
    ManualSearchResults,
    EpisodeList,
    SearchActions,
    MonitoringToggle,
  },
  methods: {
    toggleMonitoring(media) {
      APIUtil.setMonitoringMedia(media.id, !media.monitoring).then(() => {
        this.loadMedia();
      });
    },
    editMedia() {
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
    loadMedia() {
      APIUtil.getMedia(this.mediaID)
        .then((media) => {
          this.media = media;
        })
        .finally(() => {
          this.loading = false;
        });
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
