<template>
  <page-wrapper>
    <section class="flex flex-row">
      <img class="hidden md:block" :src="media.poster" />
      <div class="ml-4">
        <h1 class="text-4xl lg:text-6xl">{{ media.title }}</h1>
        <div class="py-4 flex flex-row items-center">
          <div class="text-2xl mx-4 text-gray-300">
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
        <p class="text-lg">{{ media.description }}</p>
      </div>
      <o-tooltip variant="info" position="bottom" label="Search Manually">
        <div @click="searchManual">
          <o-icon v-if="!loadingManualSearch" icon="search" />
          <o-icon v-else class="text-4xl text-gray-300" icon="sync-alt" spin />
        </div>
      </o-tooltip>
    </section>
    <o-table
      :data="releases"
      striped
      narrowed
      hoverable
      :loading="isLoading"
      mobile-cards
    >
      <o-table-column field="id" label="ID" width="40" numeric v-slot="props">
        {{ props.row.id }}
      </o-table-column>

      <o-table-column field="first_name" label="First Name" v-slot="props">
        {{ props.row.first_name }}
      </o-table-column>

      <o-table-column field="last_name" label="Last Name" v-slot="props">
        {{ props.row.last_name }}
      </o-table-column>

      <o-table-column
        field="date"
        label="Date"
        position="centered"
        v-slot="props"
      >
        {{ new Date(props.row.date).toLocaleDateString() }}
      </o-table-column>

      <o-table-column label="Gender" v-slot="props">
        <span>
          <o-icon
            pack="fas"
            :icon="props.row.gender === 'Male' ? 'mars' : 'venus'"
          >
          </o-icon>
          {{ props.row.gender }}
        </span>
      </o-table-column>
    </o-table>
    <div v-for="release in releases" :key="release.title">
      {{ release }}
    </div>
  </page-wrapper>
</template>

<script>
import PageWrapper from "../components/PageWrapper.vue";
import APIUtil from "../util/APIUtil";
import MediaUtil from "../util/MediaUtil";

const STATUS_TYPES = ["wanted", ""];

export default {
  data() {
    return {
      media: {},
      mediaID: 0,
      releases: [],
      loadingManualSearch: false,
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
  },
  created() {
    this.mediaID = parseInt(this.$route.params.media_id);
  },
  mounted() {
    APIUtil.getMedia(this.mediaID).then((media) => {
      this.media = media;
    });
  },
};
</script>
