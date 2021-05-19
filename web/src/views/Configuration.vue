<template>
  <page-wrapper>
    <o-tabs type="boxed">
      <o-tab-item>
        <template v-slot:header>
          <span class="text-2xl flex flex-row items-center">
            <vue-fontawesome icon="server"></vue-fontawesome
            ><span style="white-space: nowrap" class="flex ml-2"
              >System</span
            >
          </span>
        </template>
        <section></section>
      </o-tab-item>
      <o-tab-item>
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
      </o-tab-item>
      <o-tab-item>
        <template v-slot:header>
          <span class="text-2xl flex flex-row items-center">
            <vue-fontawesome icon="search"></vue-fontawesome
            ><span style="white-space: nowrap" class="flex ml-2"
              >Indexers</span
            >
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
        <section></section>
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
import Indexers from '../components/Indexers.vue';
import PageWrapper from "../components/PageWrapper.vue";

export default {
  components: {
    PageWrapper,
    Indexers
  },
};
</script>
