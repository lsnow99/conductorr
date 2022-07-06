<template>
  <header class="modal-card-header">
    <p class="modal-card-title">{{ title }}</p>
  </header>
  <section class="min-w-full modal-card-content w-96">
    <div>
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
    </div>
  </section>
  <footer class="modal-card-footer">
    <o-button @click="$emit('close')">Cancel</o-button>
    <div>
      <o-button variant="primary" @click="test" class="mr-3">
        <action-button :mode="testingMode"> Test </action-button>
      </o-button>
      <o-button variant="primary" @click="submit">Save</o-button>
    </div>
  </footer>
</template>

<script setup lang="ts">
import { Media } from '@/types/api/media';

const props = defineProps<{
  media: Media
}>()

</script>
