<template>
  <page-wrapper>
    <o-tabs type="boxed">
      <!-- <o-tab-item>
        <template v-slot:header>
          <span class="text-2xl flex flex-row items-center">
            <vue-fontawesome icon="film"></vue-fontawesome
            ><span style="white-space: nowrap" class="flex ml-2"
              >Movies</span
            >
          </span>
        </template>
        <section></section>
      </o-tab-item>
      <o-tab-item>
        <template v-slot:header>
          <span class="text-2xl flex flex-row items-center">
            <vue-fontawesome icon="tv"></vue-fontawesome
            ><span style="white-space: nowrap" class="flex ml-2"
              >TV Shows</span
            >
          </span>
        </template>
        <section></section>
      </o-tab-item> -->
      <o-tab-item>
        <template v-slot:header>
          <span class="text-2xl flex flex-row items-center">
            <vue-fontawesome icon="filter"></vue-fontawesome
            ><span style="white-space: nowrap" class="flex ml-2"
              >Release Profiles</span
            >
          </span>
        </template>
        <release-profiles @reload="loadConfiguration" v-model="configuration.releaseProfiles" />
      </o-tab-item>
      <o-tab-item>
        <template v-slot:header>
          <span class="text-2xl flex flex-row items-center">
            <vue-fontawesome icon="search"></vue-fontawesome
            ><span style="white-space: nowrap" class="flex ml-2">Indexers</span>
          </span>
        </template>
        <indexers />
      </o-tab-item>
      <o-tab-item>
        <template v-slot:header>
          <span class="text-2xl flex flex-row items-center">
            <vue-fontawesome icon="download"></vue-fontawesome
            ><span style="white-space: nowrap" class="flex ml-2"
              >Downloaders</span
            >
          </span>
        </template>
        <downloaders />
      </o-tab-item>
      <o-tab-item>
        <template v-slot:header>
          <span class="text-2xl flex flex-row items-center">
            <vue-fontawesome icon="play-circle"></vue-fontawesome
            ><span style="white-space: nowrap" class="flex ml-2"
              >Media Server</span
            >
          </span>
        </template>
        <section></section>
      </o-tab-item>
    </o-tabs>
  </page-wrapper>
</template>

<script>
import Indexers from "../components/Indexers.vue";
import NewDownloader from "../components/NewDownloader.vue";
import PageWrapper from "../components/PageWrapper.vue";
import ReleaseProfiles from "./ReleaseProfiles.vue";
import Downloaders from "../components/Downloaders.vue";

import APIUtil from "../util/APIUtil";

export default {
  components: {
    PageWrapper,
    Indexers,
    ReleaseProfiles,
    NewDownloader,
    Downloaders,
  },
  data() {
    return {
      configuration: {
        releaseProfiles: [{ name: "Logan", sort: "code", filter: "morecode" }],
      },
    };
  },
  methods: {
    loadConfiguration() {
      APIUtil.getConfiguration().then((re) => {
        this.configuration = re.data
      })
    }
  },
  mounted() {
    this.loadConfiguration()
  },
};
</script>
