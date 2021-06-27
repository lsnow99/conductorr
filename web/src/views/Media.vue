<template>
  <page-wrapper>
    <section class="flex flex-row">
      <img class="hidden md:block" :src="media.poster" />
      <div class="ml-4">
        <h1 class="text-4xl lg:text-6xl">{{ media.title }}</h1>
        <div class="py-4 flex flex-row items-center justify-between px-4">
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
          <div>
            <o-tooltip variant="info" :position="tooltipPosition" label="Search/Download Automatically">
              <div class="text-2xl mx-2 text-gray-300">
                <div v-if="!loadingAutoSearch" @click="searchManual">
                  <o-icon
                    class="cursor-pointer"
                    icon="bolt"
                  />
                </div>
                <o-icon v-else icon="sync-alt" spin />
              </div>
            </o-tooltip>
            <o-tooltip variant="info" :position="tooltipPosition" label="Search Manually">
              <div class="text-2xl mx-2 text-gray-300">
                <div v-if="!loadingManualSearch" @click="searchManual">
                  <o-icon
                    class="cursor-pointer"
                    icon="search"
                  />
                </div>
                <o-icon v-else icon="sync-alt" spin />
              </div>
            </o-tooltip>
          </div>
        </div>
        <p class="text-lg">{{ media.description }}</p>
      </div>
    </section>
    <o-table
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
        <div v-if="props.row.download_type == 'torrent'" class="bg-green-500">
          torrent
        </div>
        <div v-if="props.row.download_type == 'nzb'" class="bg-blue-500">
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
        <o-tooltip variant="info" position="bottom" label="Search Manually">
          <o-icon
            class="bg-red-500"
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
  </page-wrapper>
</template>

<script>
import PageWrapper from "../components/PageWrapper.vue";
import APIUtil from "../util/APIUtil";
import MediaUtil from "../util/MediaUtil";
import Helpers from "../util/Helpers";

const STATUS_TYPES = ["wanted", ""];

export default {
  data() {
    return {
      media: {},
      mediaID: 0,
      releases: [],
      loadingManualSearch: false,
      loadingAutoSearch: false,
      tooltipPosition: 'bottom'
    };
  },
  mixins: [MediaUtil],
  components: { PageWrapper },
  methods: {
    searchManual() {
      this.loadingManualSearch = true;
      APIUtil.searchReleasesManual(this.mediaID)
        .then((releases) => {
          this.releases = releases;
        })
        .finally(() => {
          this.loadingManualSearch = false;
        });
    },
    download(release) {
      const index = this.releases.findIndex(elem => elem.id == release.id)
      if (index >= 0) {
        this.releases[index].search = true
      }
    },
    niceSize: Helpers.niceSize,
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
