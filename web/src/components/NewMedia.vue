<template>
  <header class="modal-card-header">
    <p class="modal-card-title">New Media</p>
  </header>
  <section class="modal-card-content">
    <div class="flex justify-between">
      <o-field label="Search" class="flex-1">
        <o-input
          @change="search"
          type="text"
          v-model="query"
          placeholder="Title or IMDB id"
        />
      </o-field>
      <div class="self-end ml-4">
        <o-button variant="primary" @click="search">Search</o-button>
      </div>
    </div>
    <div
      v-if="results && results.length > 0"
      class="
        mt-10
        gap-y-10
        grid
        justify-items-center
        grid-flow-row grid-cols-1
        sm:grid-cols-2
        xl:grid-cols-3
        2xl:grid-cols-4
        3xl:grid-cols-5
      "
    >
      <media-card
        v-for="(media, index) in results"
        @click="selectedMedia"
        :key="index"
        :media="media"
      />
    </div>
    <o-loading
      :full-screen="false"
      :can-cancel="false"
      v-model:active="loading"
    />
  </section>
  <footer class="modal-card-footer">
    <o-button @click="$emit('close')">Cancel</o-button>
    <o-button variant="primary" @click="save">Save</o-button>
  </footer>
</template>

<script>
import APIUtil from "../util/APIUtil";
import MediaCard from "../components/MediaCard.vue";

export default {
  data() {
    return {
      results: [],
      total_results: 0,
      query: "",
      loading: false,
    };
  },
  props: {
    mediaType: {
      type: String,
      default: function() {
        return ''
      }
    }
  },
  components: { MediaCard },
  emits: ["close"],
  methods: {
    search() {
      this.loading = true;
      APIUtil.searchNew(this.query, this.mediaType, 1)
        .then((data) => {
          this.total_results = data.total_results;
          this.results = data.results;
        })
        .finally(() => {
          this.loading = false;
        });
    },
    selectedMedia(media) {
      this.loading = true;
      console.log(media)
      APIUtil.addMedia(media.imdb_id)
        .then(() => {
          this.$oruga.notification.open({
            duration: 300000,
            message: `Successfully added ${media.title}`,
            position: "bottom-right",
            variant: "success",
            closable: false,
          });
        })
        .finally(() => {
          this.loading = false;
        });
    },
  },
};
</script>
