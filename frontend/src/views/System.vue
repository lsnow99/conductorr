<template>
  <page-wrapper>
    <section>
      <h2 class="text-xl">System Health</h2>
      <div class="overflow-hidden rounded-md">
        <SystemStatus :status="status['downloader']" system="Downloaders" />
        <SystemStatus :status="status['indexer']" system="Indexers" />
        <SystemStatus :status="status['media_server']" system="Media Servers" />
        <SystemStatus :status="status['path']" system="Paths" />
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
        <iframe v-if="downloadUrl" :src="downloadUrl" class="hidden"></iframe>
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
        <o-button icon-left="download" @click="downloadLogs" variant="primary">
          Download Logs
        </o-button>
      </div>
      <div class="overflow-y-scroll h-72">
        <LogPane :logs="logs" />
      </div>
    </section>
    <Modal
      v-model="showRestoreBackupModal"
      title="Restore from Backup"
      @close="stopRestoreBackup"
    >
      <div class="p-4 bg-red-500 bg-opacity-50 rounded-md">
        <o-icon
          class="float-left mr-2 text-4xl"
          icon="exclamation-circle"
        /><span
          >WARNING: Restoring from a backup will wipe the database and restore
          it to the state it was when the backup was created. Only proceed if
          you are sure this is what you want. It is recommended to make a backup
          now of your current system in case of a failed restoration. For
          details on restoring from old versions and other information about
          backups, please refer to the
          <a :href="backupDocUrl" target="_blank" class="font-bold"
            >documentation</a
          >.</span
        >
      </div>
      <o-field class="mt-4 text-center">
        <o-upload v-model="restoreFile" drag-drop>
          <div class="p-4 border-2 border-gray-300 border-dashed rounded-lg">
            <p>
              <o-icon icon="upload" size="is-large"> </o-icon>
            </p>
            <p>Drop your backup file here or click to upload</p>
          </div>
        </o-upload>
      </o-field>
      <div class="flex flex-row justify-center w-full">
        <div class="text-bold" v-if="restoreFile">{{ restoreFile.name }}</div>
      </div>
      <template v-slot:footer>
        <div class="flex flex-row justify-between w-full">
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

<script setup lang="ts">
import PageWrapper from "../components/PageWrapper.vue";
import SystemStatus from "../components/SystemStatus.vue";
import ActionButton from "../components/ActionButton.vue";
import { LogPane, LogMessage } from "conductorr-lib";
import APIUtil from "../util/APIUtil";
import Modal from "../components/Modal.vue";
import { DOCUMENTATION_URL } from "../util/Constants";
import { storeToRefs } from "pinia";
import { useAppStore } from "@/store";
import { inject, onMounted, ref } from "vue";
import { TestingMode } from "@/types/testing_mode";
import { useTabSaver } from "@/util";
import { TaskStatus } from "@/types/api/task";

const { status } = storeToRefs(useAppStore());
const logs = ref<LogMessage[]>();
const backupMode = ref<TestingMode>(TestingMode.OFF);
const downloadUrl = ref<string | null>(null);
const tasks = ref<TaskStatus[]>([]);
const showRestoreBackupModal = ref(false);
const restoreFile = ref<File | null>(null);
const backupDocUrl = <Readonly<string>>`${DOCUMENTATION_URL}backups`;

const { lastButton, restoreFocus } = useTabSaver();

const oruga = inject("oruga");

const createBackup = async () => {
  backupMode.value = TestingMode.LOADING;
  try {
    const backupData = await APIUtil.createNewBackup();
    backupMode.value = TestingMode.SUCCESS;
    downloadUrl.value = backupData.url;
  } catch (err) {
    oruga.notification.open({
      duration: 5000,
      message: `Failed to create backup: ${err}`,
      position: "bottom-right",
      variant: "danger",
      closable: false,
    });
    backupMode.value = TestingMode.DANGER;
  } finally {
    setTimeout(() => {
      backupMode.value = TestingMode.OFF;
    }, 5000);
  }
};

const downloadLogs = () => {
  downloadUrl.value = `/api/v1/logFile?time=${new Date().getTime()}`;
};

const startRestoreBackup = ($event: Event) => {
  lastButton.value = $event.currentTarget as HTMLElement;
  showRestoreBackupModal.value = true;
};

const stopRestoreBackup = () => {
  showRestoreBackupModal.value = false;
  restoreFile.value = null;
  restoreFocus();
};

const doRestore = () => {
  
}

onMounted(async () => {
  try {
    logs.value = await APIUtil.getLogs();
  } catch {}
  try {
    tasks.value = await APIUtil.getTaskStatuses();
  } catch {}
});
</script>
