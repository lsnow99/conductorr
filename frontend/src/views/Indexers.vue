<template>
  <section class="mt-3">
    <div class="flex flex-col justify-between sm:flex-row">
      <div class="flex justify-center">
        <o-button variant="primary" @click="doShowNewIndexerModal($event)"
          >Add Indexer</o-button
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
    <ConfigItem
      :delete-message="`Are you sure you want to delete indexer '${indexer.name}'?`"
      v-for="indexer in indexers"
      collapsible
      v-model:expanded="indexer.expanded"
      :key="indexer.id"
      :title="indexer.name"
      @edit="editIndexer(indexer, $event)"
      @delete="deleteIndexer(indexer)"
    >
      <div>
        <o-field label="Base URL">
          <o-input type="text" disabled v-model="indexer.base_url" />
        </o-field>
        <o-field label="API Key">
          <o-input type="text" disabled v-model="indexer.api_key" />
        </o-field>
        <div class="flex flex-col mt-3">
          <o-switch disabled v-model="indexer.for_movies"
            >Use for Movies</o-switch
          >
          <o-switch disabled v-model="indexer.for_series">Use for TV</o-switch>
        </div>
      </div>
    </ConfigItem>
    <EditService
      v-model:active="showEditServiceModal"
      v-model="editingIndexer"
      :title="computedTitle"
      :fields="computedFields"
      :testingMode="testingMode"
      @close="closeModal"
      @test="testIndexer"
      @save="onSubmit"
    >
    </EditService>
  </section>
</template>

<script setup lang="ts">
import { ConfigurableIndexer, EditServiceMode } from "@/types/api/service"
import APIUtil from "../util/APIUtil";
import ConfigItem from "../components/ConfigItem.vue";
import EditIndexer from "../components/EditIndexer.vue";
import EditService from "../components/EditService.vue";
import { inject, onMounted, ref } from "vue";
import { TestingMode } from "@/types/testing_mode";
import { useTabSaver, useServiceUtil } from "@/util";
import { Indexer } from "@/types/api/indexer";

const XNAB_FIELDS = [
  {
    type: "text",
    label: "Base URL",
    placeholder: "Base URL",
    property: "base_url",
    required: true,
    trim: true,
  },
  {
    type: "text",
    label: "API Key",
    placeholder: "API Key",
    property: "apiKey",
    required: false,
    trim: true
  },
];

const oruga = inject('oruga')

const editingIndexer = ref<ConfigurableIndexer | null>(null)
const testingMode = ref<TestingMode>(TestingMode.OFF)
const indexers = ref<Indexer[]>([])

const { lastButton, restoreFocus } = useTabSaver()


const loadIndexers = async() => {
  try {
    const indexers = await APIUtil.getIndexers()
    indexers.value = indexers
  } catch (err) {
    
  }
}

const editIndexer = async(indexer: Indexer) => {
  await APIUtil.updateIndexer(indexer.id, indexer.name, indexer.baseUrl, indexer.apiKey, indexer.forMovies, indexer.forSeries, indexer.downloadType)
  oruga.notification.open({
    duration: 3000,
    message: `Saved successfully`,
    position: "bottom-right",
    variant: "success",
    closable: false,
  });
  loadIndexers();
}

const newIndexer = async(indexer: Indexer) => {
  await APIUtil.newIndexer(indexer.name, indexer.baseUrl, indexer.apiKey, indexer.forMovies, indexer.forSeries, indexer.downloadType)
  oruga.notification.open({
    duration: 3000,
    message: `Saved successfully`,
    position: "bottom-right",
    variant: "success",
    closable: false,
  });
  loadIndexers();
}

const testIndexer = async(indexer: Indexer) => {
  testingMode.value = TestingMode.LOADING
  try {
    await APIUtil.testIndexer(
        indexer.name,
        indexer.baseUrl,
        indexer.apiKey,
        indexer.forMovies,
        indexer.forSeries,
        indexer.downloadType
      )
      oruga.notification.open({
            duration: 5000,
            message: `Connected successfully`,
            position: "bottom-right",
            variant: "success",
            closable: false,
          });
      testingMode.value = TestingMode.SUCCESS
  } catch (err) {
    oruga.notification.open({
            duration: 5000,
            message: `Test failed: ${err}`,
            position: "bottom-right",
            variant: "danger",
            closable: false,
          });
          testingMode.value = TestingMode.DANGER
  } finally {
    setTimeout(() => {
      testingMode.value = TestingMode.OFF
    }, 5000)
  }
}

const { showNewServiceModal, showEditServiceModal, closeModal, openNewServiceModal, openEditServiceModal, editService, onSubmit } = useServiceUtil<Indexer>(lastButton, restoreFocus, newIndexer, editIndexer)

const setExpanded = (expanded: boolean) => {
  indexers.value.forEach(elem => {
    elem.expanded = expanded
  })
}

const deleteIndexer = async(indexer: Indexer) => {
  await APIUtil.deleteIndexer(indexer.id)
  loadIndexers();
}

onMounted(() => {
  loadIndexers()
})
</script>
