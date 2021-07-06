<template>
  <header class="modal-card-header">
    <p class="modal-card-title">Results</p>
  </header>
  <section class="modal-card-content min-w-full">
    <div>
    <o-table
      :height="loading ? '400px' : ''"
      :data="releases"
      narrowed
      hoverable
      :loading="loading"
      mobile-cards
    >
      <o-table-column
        field="download_type"
        label="Type"
        v-slot="props"
        position="centered"
      >
        <div
          v-if="props.row.download_type == 'torrent'"
          class="font-bold bg-green-500"
        >
          torrent
        </div>
        <div
          v-if="props.row.download_type == 'nzb'"
          class="font-bold bg-blue-500"
        >
          nzb
        </div>
      </o-table-column>

      <o-table-column field="title" label="Title" v-slot="props">
        {{ props.row.title }}
      </o-table-column>

      <o-table-column field="resolution" label="Resolution" v-slot="props">
        {{ props.row.resolution }}
      </o-table-column>

      <o-table-column field="rip_type" label="Rip Type" v-slot="props">
        {{ props.row.rip_type }}
      </o-table-column>

      <o-table-column field="encoding" label="Encoding" v-slot="props">
        {{ props.row.encoding }}
      </o-table-column>

      <o-table-column sortable field="seeders" label="Seeders" v-slot="props">
        {{ props.row.seeders }}
      </o-table-column>

      <o-table-column sortable field="size" label="Size" v-slot="props">
        {{ niceSize(props.row.size) }}
      </o-table-column>

      <o-table-column
        field="age"
        label="Age"
        sortable
        position="centered"
        v-slot="props"
      >
        {{ `${props.row.age} days` }}
      </o-table-column>

      <o-table-column
        field="warnings"
        label="Warnings"
        position="centered"
        v-slot="props"
      >
        <o-tooltip variant="info" position="bottom">
          <template v-slot:content>
            <div v-for="(warning, index) in props.row.warnings" :key="index">
              {{ warning }}
            </div>
          </template>
          <o-icon
            class="text-red-500"
            v-if="props.row.warnings && props.row.warnings.length > 0"
            icon="exclamation-circle"
          />
        </o-tooltip>
      </o-table-column>

      <o-table-column label="Download" position="centered" v-slot="props">
        <o-icon v-if="props.row.search === 1" spin icon="sync-alt" />
        <o-icon v-if="props.row.search === 2" icon="cloud-download-alt" />
        <o-icon
          v-if="!props.row.search"
          class="cursor-pointer"
          @click="download(props.row)"
          icon="download"
        />
      </o-table-column>
    </o-table>
    </div>
  </section>
  <footer class="modal-card-footer"></footer>
</template>

<script>
import APIUtil from "../util/APIUtil";
import Helpers from "../util/Helpers";

export default {
  props: {
    releases: {
      type: Array,
      default: function () {
        return [];
      },
    },
    mediaID: {
      type: Number,
      default: function () {
        return 0;
      },
    },
    loading: {
      type: Boolean,
      default: function () {
        return false;
      },
    },
  },
  methods: {
    niceSize: Helpers.niceSize,
    download(release) {
      let index = this.releases.findIndex(
        (elem) => elem.download_url == release.download_url
      );
      if (index >= 0) {
        this.releases[index].search = 1;
        APIUtil.downloadMediaRelease(this.mediaID, release)
          .then(() => {})
          .finally(() => {
            let index = this.releases.findIndex(
              (elem) => elem.download_url == release.download_url
            );
            this.releases[index].search = 2;
          });
      }
    },
  },
};
</script>
