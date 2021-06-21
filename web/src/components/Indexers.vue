<template>
    <section>
        <o-button variant="primary" @click="showNewIndexerModal = true">Add Indexer</o-button>
        <o-modal v-model:active="showNewIndexerModal">
            <new-indexer @close="showNewIndexerModal = false"/>
        </o-modal>
    </section>
</template>

<script>
import NewIndexer from "./EditIndexer.vue"
export default {
    data() {
        return {
            indexers: [],
            showNewIndexerModal: false,
        }
    },
    components: {
        NewIndexer
    },
    mounted() {
        this.indexers = [
            {
                type: 'torznab',
                name: '',
                priority: 1,
                url: '',
                apiKey: '',
                categories: []
            }
        ]
    }
}
</script>