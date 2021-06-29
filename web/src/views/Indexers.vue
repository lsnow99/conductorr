<template>
  <section class="mt-3">
    <o-button variant="primary" @click="showNewIndexerModal = true"
      >Add Indexer</o-button
    >
    <config-item :delete-message="`Are you sure you want to delete indexer '${indexer.name}'?`" v-for="indexer in indexers" :key="indexer.id" :title="indexer.name" @edit="editIndexer(indexer)" @delete="deleteIndexer(indexer)">
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
    <o-modal v-model:active="showNewIndexerModal">
      <edit-indexer
        title="New Indexer"
        @submit="newIndexer"
        @close="showNewIndexerModal = false"
      />
    </o-modal>
    <o-modal v-model:active="showEditIndexerModal">
      <edit-indexer
        title="Edit Indexer"
        :indexer="indexerToEdit"
        @submit="updateIndexer"
        @close="showEditIndexerModal = false"
      />
    </o-modal>
  </section>
</template>

<script>
import APIUtil from "../util/APIUtil";
import ConfigItem from '../components/ConfigItem.vue';
import EditIndexer from "../components/EditIndexer.vue";

export default {
  data() {
    return {
      indexers: [],
      indexerToEdit: {},
      showNewIndexerModal: false,
      showEditIndexerModal: false,
    };
  },
  components: {
    EditIndexer,
    ConfigItem
  },
  methods: {
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
    editIndexer(indexer) {
      this.indexerToEdit = indexer;
      this.showEditIndexerModal = true;
    },
    deleteIndexer(indexer) {
      APIUtil.deleteIndexer(indexer.id).then(() => {
        this.loadIndexers()
      })
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
