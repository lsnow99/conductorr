<template>
  <modal v-model="computedActive" @close="close" :title="title">
    <div v-if="computedIndexer">
      Indexer Type
      <div class="flex justify-center">
        <radio-group
          name="indexerType"
          v-model="computedIndexer.downloadType"
          :options="[
            { text: 'Torznab', value: 'torrent' },
            { text: 'Newznab', value: 'nzb' },
          ]"
        />
      </div>
      <o-field label="Name">
        <o-input
          type="text"
          v-model="computedIndexer.name"
          placeholder="Name"
        />
      </o-field>
      <o-field label="URL">
        <o-input
          type="text"
          v-model="computedIndexer.baseUrl"
          placeholder="URL"
        />
      </o-field>
      <o-field label="API Key">
        <o-input
          type="text"
          v-model="computedIndexer.apiKey"
          placeholder="API Key"
        />
      </o-field>
      <div class="flex flex-col mt-3">
        <o-switch v-model="computedIndexer.forMovies">Use for Movies</o-switch>
        <o-switch v-model="computedIndexer.forSeries">Use for TV</o-switch>
      </div>
    </div>
    <template v-slot:footer>
      <o-button @click="$emit('close')">Cancel</o-button>
      <div>
        <o-button variant="primary" @click="test" class="mr-3">
          <action-button :mode="testingMode"> Test </action-button>
        </o-button>
        <o-button variant="primary" @click="submit">Save</o-button>
      </div>
    </template>
  </modal>
</template>

<script setup lang="ts">
import APIUtil from "../util/APIUtil";
import ActionButton from "./ActionButton.vue";
import RadioGroup from "../components/RadioGroup.vue";
import Modal from "../components/Modal.vue";
import { TestingMode } from "@/types/testing_mode";
import { computed, inject, ref, WritableComputedRef } from "vue";
import { useComputedActive } from "@/util";
import { Indexer } from "@/types/api/indexer";
import { DownloadType } from "@/types/api/download";

const oruga = inject("oruga");

const selectedType = ref<string | null>(null);
const newIndexer = ref<Indexer | null>(null);
const testingMode = ref<TestingMode>(TestingMode.OFF);

const computedIndexer: WritableComputedRef<Indexer | null> = computed({
  get() {
    if (!newIndexer.value) {
      newIndexer.value = {
        id: 0,
        name: "",
        baseUrl: "",
        apiKey: "",
        userID: 0,
        forMovies: false,
        forSeries: false,
        downloadType: DownloadType.NZB
      }
    }
    return newIndexer.value;
  },
  set(newVal: Indexer | null) {
    newIndexer.value = newVal;
  },
});

const props = defineProps<{
  indexer: Indexer;
  title: string;
  active: boolean;
}>();

const emit = defineEmits<{
  (e: "submit", submittedIndexer: Indexer): void;
  (e: "close"): void;
  (e: "update:active", active: boolean): void;
}>();

const computedActive = useComputedActive(props, emit);

const close = () => {
  computedIndexer.value = null;
  emit("close");
};

const test = async () => {
  sanitize();
  const validationErr = validate();
  if (validationErr) {
    oruga.notification.open({
      duration: 5000,
      message: validationErr,
      position: "bottom-right",
      variant: "danger",
      closable: false,
    });
    return;
  }
  testingMode.value = TestingMode.LOADING;
  try {
    await APIUtil.testIndexer(
      computedIndexer.value!.name,
      computedIndexer.value!.baseUrl,
      computedIndexer.value!.apiKey,
      computedIndexer.value!.forMovies,
      computedIndexer.value!.forSeries,
      computedIndexer.value!.downloadType
    );
    oruga.notification.open({
      duration: 5000,
      message: `Connected successfully`,
      position: "bottom-right",
      variant: "success",
      closable: false,
    });
    testingMode.value = TestingMode.SUCCESS;
  } catch (err) {
    oruga.notification.open({
      duration: 5000,
      message: `Test failed: ${err}`,
      position: "bottom-right",
      variant: "danger",
      closable: false,
    });
    testingMode.value = TestingMode.DANGER;
  } finally {
    setTimeout(() => {
      testingMode.value = TestingMode.OFF;
    }, 5000);
  }
};

const sanitize = () => {
  if (computedIndexer.value) {
    computedIndexer.value.name = computedIndexer.value.name.trim() ?? "";
    computedIndexer.value.baseUrl = computedIndexer.value.baseUrl.trim() ?? "";
    computedIndexer.value.apiKey = computedIndexer.value.apiKey.trim() ?? "";
    if (computedIndexer.value.baseUrl.endsWith("/")) {
      computedIndexer.value.baseUrl = computedIndexer.value.baseUrl.slice(
        0,
        -1
      );
    }
  }
};

const validate = () => {
  if (computedIndexer.value?.name) {
    return "Name is required";
  } else if (computedIndexer.value?.baseUrl) {
    return "Base URL is required";
  } else if (computedIndexer.value?.apiKey) {
    return "API Key is required";
  }
};

const submit = () => {
  sanitize();
  const validationErr = validate();
  if (validationErr) {
    oruga.notification.open({
      duration: 5000,
      message: validationErr,
      position: "bottom-right",
      variant: "danger",
      closable: false,
    });
    return;
  }
  emit("submit", computedIndexer.value!);
};
</script>
