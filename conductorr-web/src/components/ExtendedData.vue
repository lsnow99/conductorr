<template>
  <div class="container">
    <b-loading
      style="z-index: 1000"
      :active.sync="loading"
      :is-full-page="false"
      :can-cancel="false"
    ></b-loading>
    <b-tabs>
      <b-tab-item label="Timeline">
        <b-steps v-model="activeStep" type="is-info" :has-navigation="false">
          <b-step-item label="Grabbed" icon="search" :clickable="false"></b-step-item>
          <b-step-item label="Downloading" icon="download" :clickable="false"></b-step-item>
          <b-step-item label="Filebot" icon="robot" :clickable="false"></b-step-item>
          <b-step-item label="Plex Scan" icon="play" :clickable="false"></b-step-item>
          <b-step-item label="Done" icon="check" :clickable="false"></b-step-item>
          <div class="tile is-ancestor">
            <div class="tile" style="justify-content: center">
              <timeago :datetime="job.time_grabbed"></timeago>
            </div>
            <div class="tile" style="justify-content: center">
              <timeago v-if="activeStep > 1" :datetime="job.time_filebot_started"></timeago>
              <span v-if="activeStep == 1">now</span>
            </div>
            <div class="tile" style="justify-content: center">
              <timeago v-if="activeStep > 2" :datetime="job.time_filebot_done"></timeago>
              <span v-if="activeStep == 2">now</span>
            </div>
            <div class="tile" style="justify-content: center">
              <timeago v-if="activeStep > 3" :datetime="job.time_scan_started"></timeago>
              <span v-if="activeStep == 3">now</span>
            </div>
            <div class="tile" style="justify-content: center">
              <timeago v-if="activeStep > 3" :datetime="job.time_scan_done"></timeago>
            </div>
          </div>
        </b-steps>
      </b-tab-item>

      <b-tab-item label="Filebot Logs">
        <code-viewer
          :codeText="job.filebot_logs"
        ></code-viewer>
      </b-tab-item>
      
      <b-tab-item label="Metadata">
        <code-viewer
          :codeText="getMetadata(job)"
        ></code-viewer>
      </b-tab-item>
    </b-tabs>
  </div>
</template>

<script>
import CodeViewer from "@/components/CodeViewer.vue";
import axios from "axios";
import { SnackbarProgrammatic as Snackbar } from "buefy";

export default {
  name: "extended-data",
  components: {
    CodeViewer
  },
  props: ["jobId"],
  mounted: function() {
    axios
      .get(`/api/job/${this.jobId}`, this.axiosAuthConfig())
      .then(response => {
        this.job = response.data
        switch(this.job.status) {
            case 'DOWNLOADING':
                this.activeStep = 1
                break
            case 'FILEBOT':
                this.activeStep = 2
                break
            case 'PLEX':
                this.activeStep = 3
                break
            case 'DONE':
                this.activeStep = 4
                break
            default:
                this.activeStep = 0
        }
      })
      .catch(error => {
        this.checkUnauthorizedToken(error);
        Snackbar.open({
          message: "Error loading job",
          type: "is-danger",
          indefinite: false,
          duration: 3000
        });
        throw error;
      })
      .finally(() => {
        this.loading = false;
      });
  },
  data() {
    return {
      loading: true,
      activeStep: 2,
      job: {}
    };
  },
  methods: {
      getMetadata(job) {
          let clone = JSON.parse(JSON.stringify(job))
          delete clone.filebot_logs
          delete clone.scanner_logs
          return clone
      }
  }
};
</script>
