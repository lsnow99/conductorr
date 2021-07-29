<template>
  <o-tooltip
    variant="info"
    :position="tooltipPosition"
    label="Search/Download Automatically"
  >
    <div :class="fontSize" class="mx-2 text-gray-300">
      <div v-if="!loadingAutoSearch" @click="searchManual" role="button" aria-label="Search automatically">
        <o-icon class="cursor-pointer" icon="bolt" />
      </div>
      <o-icon v-else icon="sync-alt" spin />
    </div>
  </o-tooltip>
  <o-tooltip variant="info" :position="tooltipPosition" label="Search Manually">
    <div :class="fontSize" class="mx-2 text-gray-300">
      <div v-if="!loadingManualSearch" @click="searchManual" role="button" aria-label="Search manually">
        <o-icon class="cursor-pointer" icon="search" />
      </div>
      <o-icon v-else icon="sync-alt" spin />
    </div>
  </o-tooltip>
    <o-modal
      full-screen
      v-model:active="showManualReleasesModal"
      @close="showManualReleasesModal = false"
    >
      <manual-search-results
        @close="showManualReleasesModal = false"
        :releases="releases"
        :loading="loadingManualSearch"
        :mediaID="mediaID"
      />
    </o-modal>
</template>

<script>
import APIUtil from "../util/APIUtil"
import ManualSearchResults from "./ManualSearchResults.vue"

export default {
  data() {
    return {
      releases: [],
      showManualReleasesModal: false,
      tooltipPosition: "bottom",
      loadingManualSearch: false,
      loadingAutoSearch: false,
    };
  },
  props: {
    mediaID: {
      type: Number,
      default: function () {
        return undefined;
      },
    },
    size: {
      type: String,
      default: function () {
        return "";
      },
    },
  },
  components: { ManualSearchResults },
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
  },
  mounted() {
    const screenWidth = window.innerWidth;
    if (screenWidth < 768) {
      this.tooltipPosition = "left";
    }
  },
  computed: {
    fontSize() {
      switch (this.size) {
        case "small":
          return "text-md";
        case "large":
          return "text-2xl";
        default:
          return "text-lg";
      }
    },
  },
};
</script>
