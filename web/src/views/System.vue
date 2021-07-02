<template>
  <page-wrapper>
      System Health
      <SystemStatus :status="$store.getters.status['downloader']" :system="'Downloaders'" />
      {{$store.getters.status}}
  </page-wrapper>
</template>

<script>
import PageWrapper from "../components/PageWrapper.vue";
import SystemStatus from "../components/SystemStatus.vue"

export default {
  components: { PageWrapper, SystemStatus },
};
</script>
