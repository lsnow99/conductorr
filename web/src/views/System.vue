<template>
  <page-wrapper>
    <span class="text-xl">System Health</span>
    <div class="rounded-md overflow-hidden">
      <SystemStatus
        :status="$store.getters.status['downloader']"
        :system="'Downloaders'"
      />
      <SystemStatus
        :status="$store.getters.status['indexer']"
        :system="'Indexers'"
      />
      <SystemStatus :status="$store.getters.status['path']" :system="'Paths'" />
    </div>
    <div class="text-xl mt-10">Logs</div>
    <div class="h-72 overflow-y-scroll">
      <LogPane :logs="logs" />
    </div>
  </page-wrapper>
</template>

<script>
import PageWrapper from "../components/PageWrapper.vue";
import SystemStatus from "../components/SystemStatus.vue";
import LogPane from "../components/LogPane.vue";
import APIUtil from "../util/APIUtil";
import { DateTime } from "luxon";

export default {
  components: { PageWrapper, SystemStatus, LogPane },
  data() {
    return {
      logs: [],
    };
  },
  mounted() {
    APIUtil.getLogs().then((logs) => {
      for (let i = 0; i < logs.length; i++) {
        logs[i].timestamp = DateTime.fromJSDate(new Date(logs[i].timestamp));
      }
      this.logs = logs;
    });
  },
};
</script>
