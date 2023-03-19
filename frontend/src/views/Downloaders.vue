<template>
  <section class="mt-3">
    <div class="flex flex-col justify-between sm:flex-row">
      <div class="flex justify-center">
        <o-button @click="showNewDownloader" variant="primary"
          >Add Downloader</o-button
        >
      </div>
      <div class="flex justify-center mt-4 sm:mt-0">
        <o-button
          variant="primary"
          icon-left="plus-square"
          @click="setExpanded(true)"
          class="mr-3"
          >Expand All</o-button
        ><o-button
          variant="primary"
          icon-left="minus-square"
          @click="setExpanded(false)"
          >Collapse All</o-button
        >
      </div>
    </div>
    <config-item
      @edit="editService(downloader, $event)"
      @delete="deleteDownloader(downloader)"
      v-model:expanded="downloader.expanded"
      collapsible
      :title="downloader.name"
      :delete-message="`Are you sure you want to delete downloader '${downloader.name}'?`"
      v-for="downloader in downloaders"
      :key="downloader.id"
    >
      <div class="flex flex-col">
        Configuration:
        <code class="p-2 whitespace-pre bg-gray-800">
          {{ JSON.stringify(getConfig(downloader).config, null, 4) }}
        </code>
      </div>
    </config-item>
    <NewDownloader
      v-model:active="showNewDownloaderModal"
      @close="closeNewDownloaderModal"
      @selected="selectDownloaderType"
    />
    <EditService
      v-model:active="showModal"
      v-model="editingService"
      :title="computedTitle"
      :fields="computedFields"
      :testingMode="testingMode"
      @close="closeModalWrapper"
      @test="testDownloader"
      @save="onSubmit"
    />
  </section>
</template>

<script setup lang="ts">
import NewDownloader from "../components/NewDownloader.vue";
import EditService from "../components/EditService.vue";
import APIUtil from "@/util/APIUtil";
import {
  Downloader,
  FlatDownloader,
  DownloaderType,
} from "@/types/api/downloader";
import ConfigItem from "../components/ConfigItem.vue";
import RadioGroup from "../components/RadioGroup.vue";
import { ref, inject, onMounted, computed, warn } from "vue";
import { TestingMode } from "@/types/testing_mode";
import { useTabSaver, useServiceUtil } from "@/util";
import { ConfigurableService } from "@/types/api/service";

const NZBGET_FIELDS = [
  {
    type: "text",
    label: "Base URL",
    placeholder: "http://localhost:6789",
    property: "baseUrl",
    required: true,
    trim: true,
  },
  {
    type: "text",
    label: "Username",
    placeholder: "nzbget",
    property: "username",
    required: true,
    trim: true,
  },
  {
    type: "password",
    label: "Password",
    placeholder: "tegbzn6789",
    property: "password",
    required: true,
    trim: true,
  },
];

const QBITTORRENT_FIELDS = [
  {
    type: "text",
    label: "Base URL",
    placeholder: "http://localhost:6881",
    property: "baseUrl",
    required: true,
    trim: true,
  },
  {
    type: "text",
    label: "Username",
    placeholder: "admin",
    property: "username",
    required: true,
    trim: true,
  },
  {
    type: "password",
    label: "Password",
    placeholder: "adminadmin",
    property: "password",
    required: true,
    trim: true,
  },
];

const TRANSMISSION_FIELDS = [
  {
    type: "text",
    label: "Base URL",
    placeholder: "http://localhost:9091",
    property: "baseUrl",
    required: true,
    trim: true,
  },
  {
    type: "text",
    label: "Username",
    placeholder: "transmission",
    property: "username",
    required: true,
    trim: true,
  },
  {
    type: "password",
    label: "Password",
    placeholder: "transmission",
    property: "password",
    required: true,
    trim: true,
  },
];

const COMMON_FIELDS_BEGINNING = [
  {
    type: "text",
    label: "Name",
    placeholder: "Name",
    property: "name",
    required: true,
    trim: true
  },

]

const COMMON_FIELDS_END = [
  {
    type: "radio",
    label: "Post-Processing Action",
    property: "fileAction",
    options: [
      { text: "MOVE", value: "move" },
      { text: "COPY", value: "copy" },
    ],
  },
];

const oruga = inject("oruga");

const testingMode = ref<TestingMode>(TestingMode.OFF);
const downloaders = ref<(FlatDownloader & { expanded?: boolean })[]>([]);
const showNewDownloaderModal = ref(false);
const selectedDownloaderType = ref<DownloaderType | null>(null);

const { lastButton, restoreFocus } = useTabSaver();

const getConfig = (downloader: FlatDownloader) => {
  const downloaderType =
    selectedDownloaderType.value ?? downloader.downloaderType;
  if (!downloaderType) {
    throw "downloader type not supplied"
  }
  let config;
  if (downloaderType === DownloaderType.nzbget || downloaderType === DownloaderType.transmission || downloaderType === DownloaderType.qbittorrent) {
    config = {
      baseUrl: downloader.baseUrl,
      username: downloader.username,
      password: downloader.password,
    };
  }
  return {config, downloaderType};
};

const editDownloader = async (downloader: FlatDownloader) => {
  const {config} = getConfig(downloader)
  void (await APIUtil.updateDownloader(
    downloader.id,
    downloader.name,
    downloader.fileAction,
    config
  ));
  oruga.notification.open({
    duration: 3000,
    message: "Saved successfully",
    position: "bottom-right",
    variant: "success",
    closable: false,
  });
  loadDownloaders();
};

const createDownloader = async (downloader: FlatDownloader) => {
  const {config, downloaderType} = getConfig(downloader)
  void (await APIUtil.newDownloader(
    downloaderType,
    downloader.name,
    downloader.fileAction,
    config
  ));
  oruga.notification.open({
    duration: 3000,
    message: `Saved successfully`,
    position: "bottom-right",
    variant: "success",
    closable: false,
  });
  loadDownloaders();
};

const {
  showModal,
  closeModal,
  openNewServiceModal,
  editService,
  editingService,
  onSubmit,
  mode,
} = useServiceUtil<FlatDownloader>(
  lastButton,
  restoreFocus,
  createDownloader,
  editDownloader
);

const closeModalWrapper = () => {
  selectedDownloaderType.value = null;
  closeModal();
};

const selectDownloaderType = (downloaderType: DownloaderType) => {
  selectedDownloaderType.value = downloaderType;
  closeNewDownloaderModal();
  openNewServiceModal();
};

const closeNewDownloaderModal = () => {
  showNewDownloaderModal.value = false;
  restoreFocus();
};

const showNewDownloader = ($event: Event) => {
  lastButton.value = $event.target as HTMLElement;
  showNewDownloaderModal.value = true;
  selectedDownloaderType.value = null;
};

const loadDownloaders = async () => {
  const data = await APIUtil.getDownloaders();
  const flatDownloaders: FlatDownloader[] = data.map((downloader) => ({...downloader, ...downloader.config}))
  downloaders.value = flatDownloaders;
};

const testDownloader = async (downloaderService: ConfigurableService) => {
  const downloader = downloaderService as FlatDownloader;
  testingMode.value = TestingMode.LOADING;
  let testResult;
  const {config, downloaderType} = getConfig(downloader)
  try {
    testResult = await APIUtil.testDownloader(
      downloaderType,
      config
    );
  } catch (err) {
    testResult = {
      success: false,
      msg: `${err}`,
    };
  } finally {
    if (testResult.success) {
      oruga.notification.open({
        duration: 5000,
        message: `Connected successfully`,
        position: "bottom-right",
        variant: "success",
        closable: false,
      });
      testingMode.value = TestingMode.SUCCESS;
    } else {
      oruga.notification.open({
        duration: 5000,
        message: `Test failed: ${testResult.msg}`,
        position: "bottom-right",
        variant: "danger",
        closable: false,
      });
      testingMode.value = TestingMode.DANGER;
    }
    setTimeout(() => {
      testingMode.value = TestingMode.OFF;
    }, 5000);
  }
};

const setExpanded = (expanded: boolean) => {
  downloaders.value.forEach((elem) => {
    elem.expanded = expanded;
  });
};

const deleteDownloader = async (downloader: FlatDownloader) => {
  await APIUtil.deleteDownloader(downloader.id);
  loadDownloaders();
};

const computedTitle = computed(() =>
  mode.value === "edit" ? "Edit Downloader" : "New Downloader"
);

const computedFields = computed(() => {
  const downloaderType = editingService.value?.downloaderType ?? selectedDownloaderType.value
  if (
    downloaderType === DownloaderType.nzbget
  ) {
    return [...COMMON_FIELDS_BEGINNING, ...NZBGET_FIELDS, ...COMMON_FIELDS_END];
  } else if (
    downloaderType === DownloaderType.transmission
  ) {
    return [...COMMON_FIELDS_BEGINNING, ...TRANSMISSION_FIELDS, ...COMMON_FIELDS_END];
  } else if (downloaderType === DownloaderType.qbittorrent) {
    return [...COMMON_FIELDS_BEGINNING, ...QBITTORRENT_FIELDS, ...COMMON_FIELDS_END];
  } else if (!downloaderType) {
    return [];
  }
  else throw `unknown downloader type ${downloaderType}`
});

onMounted(loadDownloaders);
</script>
