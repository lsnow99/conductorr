<template>
  <o-table v-if="load" :data="formattedEpisodes" narrowed hoverable :mobile-cards="true">
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
          <search-actions size="small" :mediaID="props.row.id" />
        </div>
      </div>
    </o-table-column>
  </o-table>
</template>

<script>
import { DateTime } from "luxon";
import SearchActions from "./SearchActions.vue";
import MonitoringToggle from "./MonitoringToggle.vue";
import APIUtil from "../util/APIUtil";

export default {
  props: {
    episodes: {
      type: Array,
      default: function () {
        return [];
      },
    },
    monitoringDisabled: {
      type: Boolean,
      default: function () {
        return false;
      },
    },
    load: {
      type: Boolean,
      default: function() {
        return false;
      }
    }
  },
  emits: ["reload"],
  components: { SearchActions, MonitoringToggle },
  methods: {
    formatDate(timestamp) {
      return DateTime.fromJSDate(new Date(timestamp)).toLocaleString(
        DateTime.DATE_SHORT
      );
    },
    toggleMonitoring(media) {
      APIUtil.setMonitoringMedia(media.id, !media.monitoring).then(() => {
        this.$emit("reload");
      });
    },
  },
  computed: {
    formattedEpisodes() {
      return this.episodes.map((elem) => {
        elem.formattedDate = this.formatDate(elem.released_at)
        return elem
      })
    }
  }
};
</script>
