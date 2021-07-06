<template>
  <page-wrapper>
    <section class="flex flex-row">
      <img class="hidden md:block" :src="media.poster" />
      <div class="ml-4 flex flex-1 flex-col">
        <h1 class="text-4xl lg:text-6xl">{{ media.title }}</h1>
        <div
          class="
            py-4
            flex flex-col
            sm:flex-row sm:items-center
            justify-between
            px-4
          "
        >
          <div class="flex flex-row items-center">
            <div class="text-2xl mr-4 text-gray-300">
              {{ mediaYear(media) }}
            </div>
            <div class="text-2xl mx-4 text-gray-300">
              <o-icon class="text-lg" icon="star" />
              {{ media.imdb_rating }}%
            </div>
            <a
              :href="`https://www.imdb.com/title/${media.imdb_id}`"
              target="_blank"
              class="inline-block pt-1 mx-4"
            >
              <o-icon
                class="text-4xl text-gray-300 hover:text-yellow-300"
                pack="fab"
                icon="imdb"
              />
            </a>
          </div>
          <div class="flex self-end">
            <o-tooltip
              variant="info"
              :position="tooltipPosition"
              label="Delete Media"
            >
              <div class="text-2xl mx-2 text-gray-300">
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
              <div class="text-2xl mx-2 text-gray-300">
                <div @click="editMedia">
                  <o-icon class="cursor-pointer" icon="wrench" />
                </div>
              </div>
            </o-tooltip>
            <o-tooltip
              variant="info"
              :position="tooltipPosition"
              label="Search/Download Automatically"
            >
              <div class="text-2xl mx-2 text-gray-300">
                <div v-if="!loadingAutoSearch" @click="searchManual">
                  <o-icon class="cursor-pointer" icon="bolt" />
                </div>
                <o-icon v-else icon="sync-alt" spin />
              </div>
            </o-tooltip>
            <o-tooltip
              variant="info"
              :position="tooltipPosition"
              label="Search Manually"
            >
              <div class="text-2xl mx-2 text-gray-300">
                <div v-if="!loadingManualSearch" @click="searchManual">
                  <o-icon class="cursor-pointer" icon="search" />
                </div>
                <o-icon v-else icon="sync-alt" spin />
              </div>
            </o-tooltip>
          </div>
        </div>
        <p class="text-lg">{{ media.description }}</p>
      </div>
    </section>
    <section class="mt-4"></section>
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

    <o-modal
    full-screen 
      v-model:active="showManualReleasesModal"
      @close="showManualReleasesModal = false"
    >
      <manual-search-results
        :releases="releases"
        :loading="loadingManualSearch"
        :mediaID="mediaID"
      />
    </o-modal>
  </page-wrapper>
</template>

<script>
import PageWrapper from "../components/PageWrapper.vue";
import APIUtil from "../util/APIUtil";
import MediaUtil from "../util/MediaUtil";
import EditMedia from "../components/EditMedia.vue";
import ConfirmDelete from "../components/ConfirmDelete.vue";
import ManualSearchResults from "../components/ManualSearchResults.vue";

export default {
  data() {
    return {
      media: {},
      mediaID: 0,
      releases: [],
      loadingManualSearch: false,
      loadingAutoSearch: false,
      tooltipPosition: "bottom",
      showEditMediaModal: false,
      showConfirmDeleteModal: false,
      showManualReleasesModal: false,
    };
  },
  mixins: [MediaUtil],
  components: { PageWrapper, EditMedia, ConfirmDelete, ManualSearchResults },
  methods: {
    searchManual() {
      this.loadingManualSearch = true;
      APIUtil.searchReleasesManual(this.mediaID)
        .then((releases) => {
          this.releases = releases;
          this.showManualReleasesModal = true;
        })
        .catch((err) => {
          this.$oruga.notification.open({
            duration: 3000,
            message: `Error searching: ${err}`,
            variant: "danger",
            closable: false,
            position: "bottom-right",
          });
        })
        .finally(() => {
          this.loadingManualSearch = false;
        });
    },
    editMedia() {
      this.showEditMediaModal = true;
    },
    updateMedia({ profileID, pathID }) {
      APIUtil.updateMedia(this.mediaID, profileID, pathID).then(() => {
        this.showEditMediaModal = false;
      });
    },
    doDelete() {
      APIUtil.deleteMedia(this.mediaID).then(() => {
        this.$router.push({ name: "library" });
      });
    },
  },
  created() {
    this.mediaID = parseInt(this.$route.params.media_id);
  },
  mounted() {
    APIUtil.getMedia(this.mediaID).then((media) => {
      this.media = media;
    });
    const screenWidth = window.innerWidth;
    if (screenWidth < 768) {
      this.tooltipPosition = "left";
    }
  },
};
</script>
