<template>
  <modal title="Results" v-model="computedActive" @close="$emit('close')" full-screen>
    <o-table
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

      <o-table-column
        sortable
        numeric
        field="seeders"
        label="Seeders"
        v-slot="props"
      >
        {{ props.row.seeders }}
      </o-table-column>

      <o-table-column sortable numeric field="size" label="Size" v-slot="props">
        {{ niceSize(props.row.size) }}
      </o-table-column>

      <o-table-column
        field="age"
        label="Age"
        sortable
        numeric
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
  </modal>
</template>

<script setup lang="ts">
import { Release } from "@/types/api/release";
import { niceSize, useComputedActive } from "@/util";
import APIUtil from "../util/APIUtil";
import Modal from "./Modal.vue";

interface ReleaseWithSearch extends Release {
  search?: "searching" | "done"
}

const props = defineProps<{
  releases: ReleaseWithSearch[],
  mediaID: number,
  loading: boolean,
  active: boolean
}>()

const emit = defineEmits<{
  (e: "close"): void,
  (e: "update:active"): void
}>()

const computedActive = useComputedActive(props, emit)

const download = async(release: ReleaseWithSearch) => {
  let index = props.releases.findIndex(elem => elem.downloadUrl === release.downloadUrl)
  if (index >= 0) {
    props.releases[index].search = "searching";
    try {
      await APIUtil.downloadMediaRelease(props.mediaID, release)
    } finally {
      index = props.releases.findIndex(
        (elem) => elem.downloadUrl === release.downloadUrl
      )
      props.releases[index].search = "done";
    }
  }
}
</script>
