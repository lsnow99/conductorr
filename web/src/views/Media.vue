<template>
  <page-wrapper>
    <section class="flex flex-row">
      <img class="hidden md:block" :src="media.poster" />
      <div class="ml-4">
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
    <section class="mt-4">
      <o-table
        :height="loading ? '400px' : ''"
        :data="releases"
        narrowed
        hoverable
        :loading="loadingManualSearch"
        mobile-cards
      >
        <o-table-column
          field="download_type"
          label="Type"
          v-slot="props"
          position="centered"
        >
          <div
            v-if="props.row.download_type == 'torrent'"
            class="font-bold bg-green-500"
          >
            torrent
          </div>
          <div
            v-if="props.row.download_type == 'nzb'"
            class="font-bold bg-blue-500"
          >
            nzb
          </div>
        </o-table-column>

        <o-table-column field="title" label="Title" v-slot="props">
          {{ props.row.title }}
        </o-table-column>

        <o-table-column field="resolution" label="Resolution" v-slot="props">
          {{ props.row.resolution }}
        </o-table-column>

        <o-table-column field="rip_type" label="Rip Type" v-slot="props">
          {{ props.row.rip_type }}
        </o-table-column>

        <o-table-column field="encoding" label="Encoding" v-slot="props">
          {{ props.row.encoding }}
        </o-table-column>

        <o-table-column sortable field="size" label="Size" v-slot="props">
          {{ niceSize(props.row.size) }}
        </o-table-column>

        <o-table-column
          field="age"
          label="Age"
          sortable
          position="centered"
          v-slot="props"
        >
          {{ `${props.row.age} days` }}
        </o-table-column>

        <o-table-column
          field="warnings"
          label="Warnings"
          position="centered"
          v-slot="props"
        >
          <o-tooltip variant="info" position="bottom">
            <template v-slot:content>
              <div v-for="(warning, index) in props.row.warnings" :key="index">
                {{ warning }}
              </div>
            </template>
            <o-icon
              class="text-red-500"
              v-if="props.row.warnings && props.row.warnings.length > 0"
              icon="exclamation-circle"
            />
          </o-tooltip>
        </o-table-column>

        <o-table-column label="Download" position="centered" v-slot="props">
          <o-icon v-if="props.row.search" spin icon="sync-alt" />
          <o-icon
            v-else
            class="cursor-pointer"
            @click="download(props.row)"
            icon="download"
          />
        </o-table-column>
      </o-table>
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
  </page-wrapper>
</template>

<script>
import PageWrapper from "../components/PageWrapper.vue";
import APIUtil from "../util/APIUtil";
import MediaUtil from "../util/MediaUtil";
import Helpers from "../util/Helpers";
import EditMedia from "../components/EditMedia.vue";

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
    };
  },
  mixins: [MediaUtil],
  components: { PageWrapper, EditMedia },
  methods: {
    searchManual() {
      this.loadingManualSearch = true;
      APIUtil.searchReleasesManual(this.mediaID)
        .then((releases) => {
          this.releases = releases;
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
    download(release) {
      let index = this.releases.findIndex((elem) => elem.download_url == release.download_url);
      if (index >= 0) {
        this.releases[index].search = true;
        APIUtil.downloadMediaRelease(this.mediaID, release)
          .then(() => {})
          .finally(() => {
            let index = this.releases.findIndex(
              (elem) => elem.download_url == release.download_url
            );
            this.releases[index].search = false;
          });
      }
    },
    niceSize: Helpers.niceSize,
    editMedia() {
      this.showEditMediaModal = true;
    },
    updateMedia(profileID, pathID) {
      APIUtil.updateMedia(this.mediaID, profileID, pathID).then(() => {
        this.showEditMediaModal = false;
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
