<template>
  <page-wrapper>
    <span class="text-xl">System Health</span>
    <div class="rounded-md overflow-hidden">
      <SystemStatus
        :status="$store.getters.status['downloader']"
        system="Downloaders"
      />
      <SystemStatus
        :status="$store.getters.status['indexer']"
        system="Indexers"
      />
      <SystemStatus
        :status="$store.getters.status['media_server']"
        system="Media Servers"
      />
      <SystemStatus :status="$store.getters.status['path']" system="Paths" />
    </div>
    <div class="text-xl mt-10">Logs</div>
    <div class="h-72 overflow-y-scroll">
      <LogPane :logs="logs" />
    </div>
    <div>
      <o-button icon-left="download" variant="primary" @click="createBackup">
        <action-button :mode="backupMode"> Download System Backup </action-button>
      </o-button>
      <o-button icon-left="download" variant="primary">
        Download Logs
      </o-button>
      <iframe :src="backupUrl" class="hidden"></iframe>
    </div>
  </page-wrapper>
</template>

<script>
import PageWrapper from "../components/PageWrapper.vue";
import SystemStatus from "../components/SystemStatus.vue";
import ActionButton from "../components/ActionButton.vue";
import LogPane from "../components/LogPane.vue";
import APIUtil from "../util/APIUtil";
import { DateTime } from "luxon";

export default {
  components: { PageWrapper, SystemStatus, LogPane, ActionButton },
  data() {
    return {
      logs: [],
      backupMode: "",
      backupUrl: "",
      tasks: [],
    };
  },
  methods: {
    createBackup() {
      this.backupMode = "loading";
      APIUtil.createNewBackup()
        .then((backupData) => {
          this.backupMode = "success";
          this.backupUrl = backupData.url;
        })
        .catch((err) => {
          this.$oruga.notification.open({
            duration: 5000,
            message: `Failed to create backup: ${err}`,
            position: "bottom-right",
            variant: "danger",
            closable: false,
          });
          this.backupMode = "danger";
        })
        .finally(() => {
          setTimeout(() => {
            this.backupMode = "";
          }, 5000);
        });
    },
  },
  mounted() {
    APIUtil.getLogs().then((logs) => {
      for (let i = 0; i < logs.length; i++) {
        logs[i].timestamp = DateTime.fromJSDate(new Date(logs[i].timestamp));
      }
      this.logs = logs;
    });
    APIUtil.getTaskStatuses().then((tasks) => {
      this.tasks = tasks;
    });
  },
};
</script>
