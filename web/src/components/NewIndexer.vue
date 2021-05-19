<template>
  <header class="modal-card-header">
    <p class="modal-card-title">New Indexer</p>
    <o-icon
      clickable
      native-type="button"
      icon="times"
      @click="$emit('close')"
    />
  </header>
  <section class="modal-card-content">
    <div v-if="selectedType === ''">
      <div class="indexer-option" @click="selectedType = 'torznab'">Torrents (Torznab)</div>
      <div class="indexer-option" @click="selectedType = 'newznab'">Usenet (Newznab)</div>
    </div>
    <div v-if="selectedType === 'torznab'">
        <o-field label="Name">
            <o-input type="text" placeholder="Name" />
        </o-field>
    </div>
    <div v-if="selectedType === 'newznab'">

    </div>
  </section>
  <footer class="modal-card-footer">
    <o-button @click="$emit('close')">Cancel</o-button>
  </footer>
</template>

<style scoped>
.indexer-option {
  @apply py-10;
  @apply flex;
  @apply justify-center;
  @apply border-4;
  @apply rounded-lg;
  @apply m-4;
  @apply cursor-pointer;
  @apply text-2xl;
  @apply border-green-500;
  @apply hover:opacity-60;
}
</style>

<script>
export default {
  data() {
    return {
      selectedType: "",
    };
  },
};
</script>
