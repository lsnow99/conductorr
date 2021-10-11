<template>
  <section class="mt-3">
    <o-button
      variant="primary"
      @click="doShowNewIndexerModal($event)"
      >Add Indexer</o-button
    >
    <config-item
      :delete-message="`Are you sure you want to delete indexer '${indexer.name}'?`"
      v-for="indexer in indexers"
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
        <div class="mt-3 flex flex-col">
          <o-switch disabled v-model="indexer.for_movies"
            >Use for Movies</o-switch
          >
          <o-switch disabled v-model="indexer.for_series">Use for TV</o-switch>
        </div>
      </div>
    </config-item>
    <edit-indexer
      title="New Indexer"
      v-model:active="showNewIndexerModal"
      @submit="newIndexer"
      @close="closeModals"
    />
    <edit-indexer
      title="Edit Indexer"
      v-model:active="showEditIndexerModal"
      :indexer="indexerToEdit"
      @submit="updateIndexer"
      @close="closeModals"
    />
  </section>
</template>

<script>
import APIUtil from "../util/APIUtil";
import ConfigItem from "../components/ConfigItem.vue";
import EditIndexer from "../components/EditIndexer.vue";
import TabSaver from "../util/TabSaver";

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
    label: ""
  }
]

export default {
  data() {
    return {
      indexers: [],
      indexerToEdit: null,
      showNewIndexerModal: false,
      showEditIndexerModal: false,
    };
  },
  components: {
    EditIndexer,
    ConfigItem,
  },
  mixins: [TabSaver],
  methods: {
    doShowNewIndexerModal($event) {
      this.lastButton = $event.currentTarget;
      this.showNewIndexerModal = true;
    },
    doShowEditIndexerModal($event) {
      this.lastButton = $event.currentTarget;
      this.showEditIndexerModal = true;
    },
    closeModals() {
      this.showNewIndexerModal = false;
      this.showEditIndexerModal = false;
      this.restoreFocus();
    },
    newIndexer(indexer) {
      APIUtil.newIndexer(
        indexer.name,
        indexer.base_url,
        indexer.api_key,
        indexer.for_movies,
        indexer.for_series,
        indexer.download_type
      ).then(() => {
        this.$oruga.notification.open({
          duration: 3000,
          message: `New indexer saved`,
          position: "bottom-right",
          variant: "success",
          closable: false,
        });
        this.showNewIndexerModal = false;
        this.loadIndexers();
      });
    },
    editIndexer(indexer, $event) {
      this.indexerToEdit = indexer;
      this.showEditIndexerModal = true;
      this.lastButton = $event.currentTarget
    },
    deleteIndexer(indexer) {
      APIUtil.deleteIndexer(indexer.id).then(() => {
        this.loadIndexers();
      });
    },
    loadIndexers() {
      APIUtil.getIndexers().then((indexers) => {
        this.indexers = indexers;
      });
    },
    updateIndexer(indexer) {
      APIUtil.updateIndexer(
        indexer.id,
        indexer.name,
        indexer.base_url,
        indexer.api_key,
        indexer.for_movies,
        indexer.for_series,
        indexer.download_type
      ).then(() => {
        this.$oruga.notification.open({
          duration: 3000,
          message: `Saved successfully`,
          position: "bottom-right",
          variant: "success",
          closable: false,
        });
        this.showEditIndexerModal = false;
        this.loadIndexers();
      });
    },
  },
  mounted() {
    this.loadIndexers();
  },
};
</script>
