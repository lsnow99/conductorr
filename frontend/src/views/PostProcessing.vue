<template>
  <section class="mt-3">
    <div class="flex flex-col justify-between sm:flex-row">
      <div class="flex justify-center">
        <o-button @click="openNewServiceModal($event)" variant="primary">Add Path</o-button>
      </div>
      <div class="flex justify-center mt-4 sm:mt-0">
        <o-button variant="primary" icon-left="plus-square" @click="setExpanded(true)" class="mr-3">Expand
          All</o-button><o-button variant="primary" icon-left="minus-square" @click="setExpanded(false)">Collapse
          All</o-button>
      </div>
    </div>
    <config-item @edit="editService(path, $event)" @delete="deletePath(path)" v-model:expanded="path.expanded"
      collapsible :title="path.path" :delete-message="`Are you sure you want to delete path '${path.path}'?`"
      v-for="path in paths" :key="path.id">
      <div class="flex flex-col">
        Configuration:
        <code class="p-2 whitespace-pre bg-gray-800">
          {{ JSON.stringify(path) }}
        </code>
      </div>
    </config-item>
    <EditService v-model:active="showModal" v-model="editingService" :title="computedTitle" :fields="FIELDS"
      :testingMode="testingMode" @close="closeModal" @test="() => {}" @save="onSubmit" />
  </section>
</template>

<script setup lang="ts">
import EditService from "../components/EditService.vue";
import APIUtil from "@/util/APIUtil";
import {
  Path,
  NewPath as NewPath
} from "../types/api/path";
import ConfigItem from "../components/ConfigItem.vue";
import RadioGroup from "../components/RadioGroup.vue";
import { ref, inject, onMounted, computed, warn } from "vue";
import { TestingMode } from "@/types/testing_mode";
import { useTabSaver, useServiceUtil } from "@/util";
import { ConfigurableService } from "@/types/api/service";


const FIELDS = [
  {
    type: "text",
    label: "Path",
    property: "path",
    required: true,
    trim: true,
  },
];

const oruga = inject("oruga");

const testingMode = ref<TestingMode>(TestingMode.OFF);
const paths = ref<(Path & { expanded?: boolean })[]>([]);

const { lastButton, restoreFocus } = useTabSaver();

const editPath = async (path: Path) => {
  void (await APIUtil.updatePath(
    path.id,
    path.path,
    path.moviesDefault,
    path.seriesDefault
  ));
  oruga.notification.open({
    duration: 3000,
    message: "Saved successfully",
    position: "bottom-right",
    variant: "success",
    closable: false,
  });
  loadPaths()
};

const createPath = async (path: NewPath) => {
  console.log('path')
  void (await APIUtil.createNewPath(
    path.path,
    path.moviesDefault,
    path.seriesDefault
  ));
  oruga.notification.open({
    duration: 3000,
    message: `Saved successfully`,
    position: "bottom-right",
    variant: "success",
    closable: false,
  });
  loadPaths()
};

const {
  showModal,
  closeModal,
  openNewServiceModal,
  editService,
  editingService,
  onSubmit,
  mode,
} = useServiceUtil<Path>(
  lastButton,
  restoreFocus,
  createPath,
  editPath
);

const loadPaths = async () => {
console.log('loading paths')
  const data = await APIUtil.getPaths();
  paths.value = data;
};

const setExpanded = (expanded: boolean) => {
  paths.value.forEach((elem) => {
    elem.expanded = expanded;
  });
};

const deletePath = async (path: Path) => {
  await APIUtil.deletePath(path.id);
  loadPaths();
};

const computedTitle = computed(() => mode.value === "edit" ? "Edit Path" : "New Path")

onMounted(() => loadPaths())
</script>
