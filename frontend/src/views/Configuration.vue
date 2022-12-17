<template>
  <page-wrapper>
    <o-tabs type="boxed" ref="tabs" v-model="currentTab">
      <o-tab-item>
        <template v-slot:header>
          <span class="flex flex-row items-center text-2xl">
            <vue-fontawesome icon="filter"></vue-fontawesome
            ><span style="white-space: nowrap" class="flex ml-2"
              >Release Profiles</span
            >
          </span>
        </template>
        <ReleaseProfiles />
      </o-tab-item>
      <o-tab-item>
        <template v-slot:header>
          <span class="flex flex-row items-center text-2xl">
            <vue-fontawesome icon="search"></vue-fontawesome
            ><span style="white-space: nowrap" class="flex ml-2">Indexers</span>
          </span>
        </template>
        <Indexers />
      </o-tab-item>
      <o-tab-item>
        <template v-slot:header>
          <span class="flex flex-row items-center text-2xl">
            <vue-fontawesome icon="download"></vue-fontawesome
            ><span style="white-space: nowrap" class="flex ml-2"
              >Downloaders</span
            >
          </span>
        </template>
        <Downloaders />
        <div></div>
      </o-tab-item>
      <o-tab-item>
        <template v-slot:header>
          <span class="flex flex-row items-center text-2xl">
            <vue-fontawesome icon="cogs"></vue-fontawesome
            ><span style="white-space: nowrap" class="flex ml-2"
              >Post-Processing</span
            >
          </span>
        </template>
        <!--<post-processing />-->
        <div></div>
      </o-tab-item>
      <o-tab-item>
        <template v-slot:header>
          <span class="flex flex-row items-center text-2xl">
            <vue-fontawesome icon="play-circle"></vue-fontawesome
            ><span style="white-space: nowrap" class="flex ml-2"
              >Media Server</span
            >
          </span>
        </template>
        <!--<media-servers />-->
        <div></div>
      </o-tab-item>
    </o-tabs>
  </page-wrapper>
</template>

<style scoped>
:deep(.o-tab-item__content) {
  min-height: 70vh;
}
</style>

<script setup lang="ts">
import Indexers from "./Indexers.vue";
import PageWrapper from "../components/PageWrapper.vue";
import ReleaseProfiles from "./ReleaseProfiles.vue";
import Downloaders from "./Downloaders.vue";
// import PostProcessing from "./PostProcessing.vue";
// import MediaServers from "./MediaServers.vue";
import { onMounted, ref, watch } from "vue";
import { useRoute, useRouter } from "vue-router";

enum TabOptions {
  PROFILES = 1,
  INDEXERS = 2,
  DOWNLOADERS = 3,
  POST_PROCESSING = 4,
  MEDIA_SERVERS = 5,
}

const SUBPATHS: { [P in TabOptions]: string } = {
  [TabOptions.PROFILES]: "profiles",
  [TabOptions.INDEXERS]: "indexers",
  [TabOptions.DOWNLOADERS]: "downloaders",
  [TabOptions.POST_PROCESSING]: "postProcessing",
  [TabOptions.MEDIA_SERVERS]: "mediaServers",
};

const currentTab = ref<TabOptions>(1);
const mounted = ref(false);
const tabs = ref<any | null>(null);

const tabsChanged = (newTab: TabOptions) => {
  if (mounted.value) {
    tabs.value?.childItems[newTab - 1].$el.scrollIntoView();
    mounted.value = false;
  }
};

const route = useRoute();
const router = useRouter();

const urlSubpath = route.params.subpath;
for (const [tabIndex, subpath] of Object.entries(SUBPATHS)) {
  if (subpath === urlSubpath) {
    currentTab.value = parseInt(tabIndex);
  }
}

onMounted(() => {
  mounted.value = true;
});

watch(currentTab, (newVal) => {
  router.replace({
    name: "configuration",
    params: { subpath: SUBPATHS[newVal] },
  });
  tabsChanged(newVal);
});
</script>
