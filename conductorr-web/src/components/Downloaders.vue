<template>
  <div class="container">
    <div class="setting-header">
      <img :src="require('@/assets/conductorr.svg')" class="setting-logo" />
      Downloaders
    </div>
    <div class="tile is-parent is-vertical">
      <div class="tile is-child">
        <span class="header-title">
          <b-icon style="vertical-align: middle" icon="wrench" size="is-medium"></b-icon>Configuration
        </span>
      </div>
      <div class="tile is-child subsection">
        <div v-for="(downloader, index) in downloaders" v-bind:key="index">
          <setting :setting="downloader.setting_name"></setting>
          <setting :setting="downloader.setting_download_dir"></setting>
            <div style="display: flex; flex-direction: row; justify-content: space-between">
                <div></div>
                <b-button type="is-danger" @click="deleteDownloader(index)">Delete Downloader {{downloader.setting_name.string_value}}</b-button>
                </div>
        </div>
      </div>
      <b-button size="is-medium" @click="addDownloader('', '')" type="is-primary">Add Downloader</b-button>
    </div>
    <div class="btn-container">
      <b-button size="is-medium" @click="save" :loading="isSaving" type="is-success">Save</b-button>
    </div>
  </div>
</template>

<script>
import { SnackbarProgrammatic as Snackbar } from "buefy";
import axios from "axios";
import Setting from "@/components/Setting.vue";

export default {
  name: "downloaders",
  components: {
    Setting
  },
  data() {
    return {
      isSaving: false,
      downloaders: [],
      fb_logs_enabled: {
        datatype: "bool",
        label: "Record Filebot Logs",
        placeholder: "",
        bool_value: true,
        description: "Save filebot logs in database"
      },
      plex_scanner_logs_enabled: {
        datatype: "bool",
        label: "Record Plex Scanner Logs",
        placeholder: "",
        bool_value: true,
        description: "Save Plex Scanner logs in database"
      }
    };
  },
  methods: {
    save() {
      this.isSaving = true;
      axios
        .post("/api/configuration/downloaders", this.computePostData(), this.axiosAuthConfig())
        .then(() => {
          Snackbar.open({
            message: "Succcessfully saved settings!",
            type: "is-success",
            indefinite: false,
            duration: 3000
          });
        })
        .catch(() => {
          Snackbar.open({
            message: "Error saving settings",
            type: "is-danger",
            indefinite: false,
            duration: 10000
          });
        })
        .finally(() => {
          this.isSaving = false;
        });
    },
    addDownloader(name='', download_dir='', id=-1) {
        let downloader = {
            setting_name: {
                datatype: "string",
                label: "The name of the downloader",
                placeholder: "Downloader_Name",
                description: "This should be the exact case-sensitive string that is sent by your post processing script",
                string_value: name
            },
            setting_download_dir: {
                datatype: "string",
                label: "The completed downloads directory",
                placeholder: "/nfs/downloads/downloader_name/completed",
                description: "This should be the exact download directory for the download (from filebot's perspective)",
                string_value: download_dir
            },
            id: id
        }
        this.downloaders.push(downloader)
    },
    deleteDownloader(index) {
        this.downloaders.splice(index, 1);
    },
    computePostData() {
        let data = []
        for(let i = 0; i < this.downloaders.length; i++) {
            let downloader = {
                name: this.downloaders[i].setting_name.string_value,
                download_dir: this.downloaders[i].setting_download_dir.string_value,
            }
            if(this.downloaders[i].id >= 0) {
                downloader.downloader_configuration_id = this.downloaders[i].id
            }
            data.push(downloader)
        }
        return data
    }
  },
  mounted() {
    axios
      .get("/api/configuration/downloaders", this.axiosAuthConfig())
      .then(response => {
        for(let i = 0; i < response.data.length; i++) {
            this.addDownloader(response.data[i].name, response.data[i].download_dir)
        }
        this.isLoading = false;
      })
      .catch(error => {
        this.checkUnauthorizedToken(error);
      });
  }
};
</script>