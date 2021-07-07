<template>
  <page-wrapper>
    <o-tabs type="boxed" ref="tabs" v-model="currentTab">
      <o-tab-item>
        <template v-slot:header>
          <span class="text-2xl flex flex-row items-center">
            <vue-fontawesome icon="filter"></vue-fontawesome
            ><span style="white-space: nowrap" class="flex ml-2"
              >Release Profiles</span
            >
          </span>
        </template>
        <release-profiles />
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
            <vue-fontawesome icon="cogs"></vue-fontawesome
            ><span style="white-space: nowrap" class="flex ml-2"
              >Post-Processing</span
            >
          </span>
        </template>
        <post-processing />
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
        <media-servers />
      </o-tab-item>
    </o-tabs>
  </page-wrapper>
</template>

<script>
import Indexers from "./Indexers.vue";
import NewDownloader from "../components/NewDownloader.vue";
import PageWrapper from "../components/PageWrapper.vue";
import ReleaseProfiles from "./ReleaseProfiles.vue";
import Downloaders from "./Downloaders.vue";
import PostProcessing from "./PostProcessing.vue";
import MediaServers from "./MediaServers.vue";

export default {
  data() {
    return {
      currentTab: 1
    }
  },
  components: {
    PageWrapper,
    Indexers,
    ReleaseProfiles,
    NewDownloader,
    Downloaders,
    PostProcessing,
    MediaServers
  },
  methods: {
    tabsChanged(newTab) {
      this.$refs.tabs.childItems[newTab-1].scrollIntoView
    }
  },
  watch: {
    currentTab(newVal, oldVal) {
      this.tabsChanged(newVal)
    }
  }
};
</script>
