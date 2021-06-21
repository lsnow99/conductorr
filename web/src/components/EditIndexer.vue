<template>
  <header class="modal-card-header">
    <p class="modal-card-title">New Indexer</p>
  </header>
  <section class="modal-card-content">
    <div>
      Indexer Type
      <div class="flex justify-center">
        <div class="overflow-hidden rounded-lg inline-block">
          <o-radio
            v-model="computedIndexer.indexer_type"
            name="indexerType"
            native-value="torznab"
            >Torznab</o-radio
          >
          <o-radio
            v-model="computedIndexer.indexer_type"
            name="indexerType"
            native-value="newznab"
            >Newznab</o-radio
          >
        </div>
      </div>
      <o-field label="Name">
        <o-input
          type="text"
          v-model="computedIndexer.name"
          placeholder="Name"
        />
      </o-field>
      <o-field label="URL">
        <o-input type="text" v-model="computedIndexer.url" placeholder="URL" />
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
      <div class="mt-3 flex flex-col">
        <o-switch v-model="computedIndexer.use_for_movies"
          >Use for Movies</o-switch
        >
        <o-switch v-model="computedIndexer.use_for_tv">Use for TV</o-switch>
      </div>
    </div>
  </section>
  <footer class="modal-card-footer">
    <o-button @click="$emit('close')">Cancel</o-button>
    <div>
      <o-button variant="primary" @click="test" class="mr-3">
        <action-button :mode="testingMode"> Test </action-button>
      </o-button>
      <o-button variant="primary" @click="save">Save</o-button>
    </div>
  </footer>
</template>

<script>
import ActionButton from "./ActionButton.vue";

const possibleTags = [
  { name: "TV", code: 5000 },
  { name: "WEB-DL", code: 5010 },
  { name: "Foreign", code: 5020 },
  { name: "SD", code: 5030 },
  { name: "HD", code: 5040 },
  { name: "UHD", code: 5045 },
  { name: "Other", code: 5050 },
  { name: "Sport", code: 5060 },
  { name: "Anime", code: 5070 },
  { name: "Documentary", code: 5080 },
];

export default {
  data() {
    return {
      selectedType: "",
      filteredTags: possibleTags,
      newIndexer: null,
      testingMode: "",
    };
  },
  props: {
    indexer: {
      type: Object,
      default: function () {
        return {
          indexer_type: "torznab",
        };
      },
    },
  },
  components: {
    ActionButton,
  },
  methods: {
    save() {},
    test() {
      if (this.testingMode === "loading") {
        this.testingMode = "success";
      } else if (this.testingMode === "success") {
        this.testingMode = "error";
      } else if (this.testingMode === "error") {
        this.testingMode = "";
      } else {
        this.testingMode = "loading";
      }
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
  },
  computed: {
    computedIndexer: {
      get() {
        if (this.newIndexer == null) {
          if (this.indexer) {
            this.newIndexer = this.indexer;
          }
        }
        return this.newIndexer;
      },
      set(newVal) {
        this.newIndexer = newVal;
      },
    },
  },
};
</script>
