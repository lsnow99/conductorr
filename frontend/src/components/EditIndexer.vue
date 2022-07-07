<template>
  <modal v-model="computedActive" @close="close" :title="title">
    <div>
      Indexer Type
      <div class="flex justify-center">
        <radio-group
          name="indexerType"
          v-model="computedIndexer.download_type"
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
          v-model="computedIndexer.base_url"
          placeholder="URL"
        />
      </o-field>
      <o-field label="API Key">
        <o-input
          type="text"
          v-model="computedIndexer.api_key"
          placeholder="API Key"
        />
      </o-field>
      <!-- <o-field label="Categories">
        <o-inputitems
          v-model="indexer.categories"
          ref="categoryInput"
          icon="filter"
          autocomplete
          :allow-new="false"
          open-on-focus
          :data="filteredTags"
          field="name"
          placeholder="Categories"
          @typing="getFilteredTags"
          @add="resetFilteredTags"
        >
          <template #selected="{ items }">
            <div
              v-for="(item, index) in items"
              :key="index"
              class="inputit-item"
              @click="$refs.categoryInput.removeItem(index, $event)"
            >
              {{ item.name }}
            </div>
          </template>
        </o-inputitems>
        <div class="mt-2">
          <o-switch v-model="allMovies"> Select All Movie Categories </o-switch>
        </div>
        <div class="mt-2">
          <o-switch v-model="allTV"> Select All TV Categories </o-switch>
        </div>
      </o-field> -->
      <div class="flex flex-col mt-3">
        <o-switch v-model="computedIndexer.for_movies">Use for Movies</o-switch>
        <o-switch v-model="computedIndexer.for_series">Use for TV</o-switch>
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
import {TestingMode} from "@/types/testing_mode"
import {ActiveProps, ActiveEmit} from "@/util/ComputedActive"
import { ref } from "vue";
import { useComputedActive } from "@/util";

const selectedType = ref<string | null>(null)
const newIndexer = ref<Indexer | null>(null)
const testingMode = ref<TestingMode>(TestingMode.OFF)

const props = defineProps<{
  indexer: Indexer,
  title: string
} & ActiveProps>()

const emit = defineEmits<{
  (e: "submit", submittedIndexer: Indexer): void
  (e: "close"): void
} & ActiveEmit>()

const computedActive = useComputedActive(props, emit)

const close = () => {
  computedIndexer.value = null
  emit("close")
}

const test = () => {
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
}

export default {
  methods: {
    close() {
      console.log('setting')
      this.computedIndexer = null;
      this.$emit("close");
    },
    test() {
      this.sanitize();
      const validationErr = this.validate();
      if (validationErr) {
        this.$oruga.notification.open({
          duration: 5000,
          message: validationErr,
          position: "bottom-right",
          variant: "danger",
          closable: false,
        });
        return;
      }
      this.testingMode = "loading";
      APIUtil.testIndexer(
        this.computedIndexer.name,
        this.computedIndexer.base_url,
        this.computedIndexer.api_key,
        this.computedIndexer.for_movies,
        this.computedIndexer.for_series,
        this.computedIndexer.download_type
      )
        .then(() => {
          this.$oruga.notification.open({
            duration: 5000,
            message: `Connected successfully`,
            position: "bottom-right",
            variant: "success",
            closable: false,
          });
          this.testingMode = "success";
        })
        .catch((err) => {
          this.$oruga.notification.open({
            duration: 5000,
            message: `Test failed: ${err}`,
            position: "bottom-right",
            variant: "danger",
            closable: false,
          });
          this.testingMode = "danger";
        })
        .finally(() => {
          setTimeout(() => {
            this.testingMode = "";
          }, 5000);
        });
    },
    getFilteredTags(text) {
      this.filteredTags = possibleTags.filter((option) => {
        return (
          option.name.toString().toLowerCase().indexOf(text.toLowerCase()) >= 0
        );
      });
    },
    resetFilteredTags() {
      this.filteredTags = possibleTags;
    },
    sanitize() {
      this.computedIndexer.name = this.computedIndexer.name
        ? this.computedIndexer.name.trim()
        : "";
      this.computedIndexer.base_url = this.computedIndexer.base_url
        ? this.computedIndexer.base_url.trim()
        : "";
      this.computedIndexer.api_key = this.computedIndexer.api_key
        ? this.computedIndexer.api_key.trim()
        : "";
      if (this.computedIndexer.base_url.endsWith("/")) {
        this.computedIndexer.base_url = this.computedIndexer.base_url.slice(
          0,
          -1
        );
      }
    },
    validate() {
      if (!this.computedIndexer.name) {
        return "Name is required";
      } else if (!this.computedIndexer.base_url) {
        return "Base URL is required";
      } else if (!this.computedIndexer.api_key) {
        return "API Key is required";
      }
    },
    submit() {
      this.sanitize();
      const validationErr = this.validate();
      if (validationErr) {
        this.$oruga.notification.open({
          duration: 5000,
          message: validationErr,
          position: "bottom-right",
          variant: "danger",
          closable: false,
        });
        return;
      }
      this.$emit("submit", this.computedIndexer);
    },
  },
  computed: {
    computedIndexer: {
      get() {
        if (!this.newIndexer) {
          this.newIndexer = {}
          Object.assign(this.newIndexer, this.indexer)
        }
        return this.newIndexer;
      },
      set(newVal) {
        this.newIndexer = newVal;
      },
    },
    computedActive: {
      get() {
        return this.active;
      },
      set(v) {
        this.$emit("update:active", v);
      },
    },
  },
};
</script>
