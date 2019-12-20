<template>
  <div class="container">
    <section>
      <b-field label="Search..." label-position="on-border">
        <b-input
          @keyup.enter.native="loadAsyncData"
          placeholder="Search..."
          v-model="filter"
          type="search"
        ></b-input>
        <p class="control">
          <b-button @click="loadAsyncData" class="button is-primary">Search</b-button>
        </p>
      </b-field>

      <b-table
        :opened-detailed="defaultOpenedDetails"
        detailed
        detail-key="job_id"
        :show-detail-icon="showDetailIcon"
        :data="data"
        :loading="loading"
        paginated
        backend-pagination
        :total="total"
        :per-page="perPage"
        @page-change="onPageChange"
        aria-next-label="Next page"
        aria-previous-label="Previous page"
        aria-page-label="Page"
        aria-current-label="Current page"
        backend-sorting
        :default-sort-direction="defaultSortOrder"
        :default-sort="[sortField, sortOrder]"
        @sort="onSort"
      >
        <template slot-scope="props">
          <b-table-column
            field="job_id"
            label="Job ID"
            width="80"
            numeric
            sortable
          >{{ props.row.job_id }}</b-table-column>

          <b-table-column field="title" label="Title" sortable>
            <template v-if="showDetailIcon">
              <a target="_blank" :href="'https://www.imdb.com/title/' + props.row.imdb_id">
                {{ props.row.title }}
                <b-icon size="is-small" icon="external-link-alt"></b-icon>
              </a>
            </template>
            <template v-else>
              <a @click="toggle(props.row)">{{ props.row.title }}</a>
            </template>
          </b-table-column>

          <b-table-column field="grabbed_quality" label="Quality" sortable>
            <b-tag size="is-medium" type="is-success" rounded>{{ props.row.grabbed_quality }}</b-tag>
          </b-table-column>

          <b-table-column field="grabbed_size" label="Size" sortable>
            <b-tag size="is-medium" rounded>{{ props.row.grabbed_size }}</b-tag>
          </b-table-column>

          <b-table-column
            field="time_grabbed"
            label="Created"
            sortable
            centered
          >
            <timeago :datetime="props.row.time_grabbed"></timeago>
          </b-table-column>

          <b-table-column field="status" label="Status" sortable centered>
            <span class="tag is-info">{{ getStatus(props.row.status) }}</span>
          </b-table-column>
        </template>

        <template slot="detail" slot-scope="props">
          <extended-data :jobId="props.row.job_id"></extended-data>
        </template>
      </b-table>
    </section>
  </div>
</template>

<script>
import ExtendedData from "@/components/ExtendedData.vue";
import axios from "axios";
import { SnackbarProgrammatic as Snackbar } from "buefy";

export default {
  name: "overview",
  components: {
    ExtendedData
  },
  props: {},
  data() {
    return {
      data: [],
      total: 0,
      loading: false,
      sortField: "job_id",
      sortOrder: "desc",
      defaultSortOrder: "desc",
      page: 1,
      perPage: 20,
      defaultOpenedDetails: [],
      showDetailIcon: true,
      activeTab: 0,
      filter: ""
    };
  },
  methods: {
    toggle(row) {
      this.$refs.table.toggleDetails(row);
    },
    loadAsyncData() {
      const params = [
        `sort_column=${this.sortField}`,
        `sort_order=${this.sortOrder}`,
        `page=${this.page}`,
        `filter=${this.filter}`
      ].join("&");
      this.loading = true;
      axios
        .get(`/api/jobs?${params}`, this.axiosAuthConfig())
        .then(response => {
          this.data = [];
          this.total = response.data.total;
          response.data.data.forEach(element => {
            // Process data if necessary
            this.data.push(element);
          });
        })
        .catch(error => {
          this.checkUnauthorizedToken(error);
          Snackbar.open({
            message: "Error loading rows",
            type: "is-danger",
            indefinite: false,
            duration: 3000
          });
          this.data = [];
          this.total = 0;
          throw error;
        })
        .finally(() => {
          this.loading = false;
        });
    },
    onPageChange(page) {
      this.page = page;
      this.loadAsyncData();
    },
    onSort(field, order) {
      this.sortField = field;
      this.sortOrder = order;
      this.loadAsyncData();
    },
    getStatus(status) {
      switch(status) {
          case 'DOWNLOADING':
              return "DOWNLOADING"
          case 'FILEBOT':
              return "FILEBOT RUNNING"
          case 'PLEX':
              return "PLEX SCAN RUNNING"
          case 'DONE':
              return "DONE"
          default:
              return "GRABBED"
      }
    }
  },
  mounted() {
    this.loadAsyncData();
  }
};
</script>

<style scoped>
.icon {
  vertical-align: middle !important;
}
</style>
