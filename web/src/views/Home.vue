<template>
  <page-wrapper>
    <div>
      <section class="">
        <div v-for="download in downloads" :key="download.identifier">{{download}}</div>
      </section>
    </div>
  </page-wrapper>
</template>

<script>
import PageWrapper from "../components/PageWrapper.vue";
import APIUtil from "../util/APIUtil";

export default {
  data() {
    return {
      downloads: [],
    };
  },
  components: { PageWrapper },
  methods: {
    refreshDownloads() {
      APIUtil.getDownloads().then((downloads) => {
        this.downloads = downloads;
      });
    },
  },
  mounted() {
    this.refreshDownloads();
  },
};
</script>
