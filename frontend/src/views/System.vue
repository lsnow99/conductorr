<template>
  <page-wrapper>
    <section>
      <h2 class="text-xl">System Health</h2>
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
    </section>
    <section>
      <h2 class="text-xl">Jobs</h2>
    </section>
    <section class="mt-10">
      <h2 class="text-xl">Backup & Restore</h2>
      <div class="flex flex-col sm:flex-row justify-evenly">
        <o-button icon-left="download" variant="primary" @click="createBackup">
          <action-button :mode="backupMode">
            Download System Backup
          </action-button>
        </o-button>
        <iframe :src="backupUrl" class="hidden"></iframe>
        <o-button
          class="mt-2 sm:mt-0"
          icon-left="upload"
          variant="primary"
          @click="startRestoreBackup"
        >
          Restore From Backup
        </o-button>
      </div>
    </section>
    <section class="mt-10">
      <div class="flex flex-row justify-between">
        <h2 class="text-xl">Logs</h2>
        <o-button icon-left="download" variant="primary">
          Download Logs
        </o-button>
      </div>
      <div class="h-72 overflow-y-scroll">
        <LogPane :logs="logs" />
      </div>
    </section>
    <Modal
      v-model="showRestoreBackupModal"
      title="Restore from Backup"
      @close="stopRestoreBackup"
    >
      <div class="bg-red-500 bg-opacity-50 rounded-md p-4">
        <o-icon
          class="text-4xl float-left mr-2"
          icon="exclamation-circle"
        /><span
          >WARNING: Restoring from a backup will wipe the database and restore
          it to the state it was when the backup was created. Only proceed if
          you are sure this is what you want. It is recommended to make a backup
          now of your current system in case of a failed restoration. For details on restoring from old versions and other information about backups, please refer to the <a :href="backupDocUrl" target="_blank" class="font-bold">documentation</a>.</span
        >
      </div>
      <o-field class="text-center mt-4">
        <o-upload v-model="restoreFile" drag-drop>
          <div class="border-dashed border-2 border-gray-300 p-4 rounded-lg">
            <p>
              <o-icon icon="upload" size="is-large"> </o-icon>
            </p>
            <p>Drop your backup file here or click to upload</p>
          </div>
        </o-upload>
      </o-field>
      <div class="w-full flex flex-row justify-center">
        <div class="text-bold" v-if="restoreFile">{{ restoreFile.name }}</div>
      </div>
      <template v-slot:footer>
        <div class="flex flex-row w-full justify-between">
          <o-button
            @click="stopRestoreBackup"
            @keydown.space="stopRestoreBackup"
            @keydown.enter="stopRestoreBackup"
            role="button"
            tabindex="0"
            >Cancel</o-button
          >
          <o-button
            @click="doRestore"
            @keydown.space="doRestore"
            @keydown.enter="doRestore"
            :disabled="!restoreFile"
            variant="primary"
            role="button"
            tabindex="0"
            >Restore</o-button
          >
        </div>
      </template>
    </Modal>
  </page-wrapper>
</template>

<script>
import PageWrapper from "../components/PageWrapper.vue";
import SystemStatus from "../components/SystemStatus.vue";
import ActionButton from "../components/ActionButton.vue";
import LogPane from "../components/LogPane.vue";
import APIUtil from "../util/APIUtil";
import { DateTime } from "luxon";
import Modal from "../components/Modal.vue";
import TabSaver from "../util/TabSaver";
import { DOCUMENTATION_URL } from "../util/Constants";

export default {
  components: { PageWrapper, SystemStatus, LogPane, ActionButton, Modal },
  data() {
    return {
      logs: [],
      backupMode: "",
      backupUrl: "",
      tasks: [],
      showRestoreBackupModal: false,
      restoreFile: null,
      backupDocUrl: `${DOCUMENTATION_URL}backups`
    };
  },
  mixins: [TabSaver],
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
    startRestoreBackup($event) {
      this.lastButton = $event.currentTarget;
      this.showRestoreBackupModal = true;
    },
    stopRestoreBackup() {
      this.showRestoreBackupModal = false;
      this.restoreFile = null;
      this.restoreFocus();
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
