<template>
  <o-table
    v-if="load"
    :data="formattedEpisodes"
    narrowed
    hoverable
    :mobile-cards="true"
  >
    <o-table-column field="monitoring" label="" v-slot="props" position="left">
      <monitoring-toggle
        :monitoring="props.row.monitoring"
        :disabled="monitoringDisabled"
        @toggle="toggleMonitoring(props.row)"
        class="text-xl"
      />
    </o-table-column>

    <o-table-column
      sortable
      numeric
      field="number"
      label="#"
      v-slot="props"
      position="left"
    >
      {{ props.row.number }}
    </o-table-column>

    <o-table-column field="title" label="Title" v-slot="props" width="60%">
      {{ props.row.title }}
    </o-table-column>

    <o-table-column
      field="released_at"
      label="Air Date"
      v-slot="props"
      position="right"
    >
      {{ props.row.formattedDate }}
    </o-table-column>

    <o-table-column label="File" position="right" width="5%" v-slot="props">
      <div class="flex flex-row">
        <o-icon
          v-if="props.row.path_ok"
          icon="check-circle"
          class="text-green-600"
        />
      </div>
    </o-table-column>

    <o-table-column label="Actions" position="right" v-slot="props">
      <div class="flex flex-row justify-between">
        <div />
        <div>
          <SearchActions size="small" :mediaID="props.row.id" />
        </div>
      </div>
    </o-table-column>
  </o-table>
</template>

<script setup lang="ts">
import { DateTime } from "luxon";
import SearchActions from "./SearchActions.vue";
import MonitoringToggle from "./MonitoringToggle.vue";
import APIUtil from "../util/APIUtil";
import { Media } from "@/types/api/media";
import { computed } from "vue";

const props = defineProps<{
  episodes: Media[];
  monitoringDisabled: boolean;
  load: boolean;
}>();

const emit = defineEmits<{
  (e: "reload"): void;
}>();

const formatDate = (timestamp: string) => {
  return DateTime.fromJSDate(new Date(timestamp)).toLocaleString(
    DateTime.DATE_SHORT
  );
};

const formattedEpisodes = computed(() => {
  return props.episodes.map((elem) => ({
    ...elem,
    formattedDate: formatDate(elem.released_at),
  }));
});

const toggleMonitoring = async (media: Media) => {
  try {
    await APIUtil.setMonitoringMedia(media.id, !media.monitoring);
    emit("reload");
  } catch {}
};
</script>
