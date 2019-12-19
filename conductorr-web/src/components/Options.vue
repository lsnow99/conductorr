<template>
  <div class="container">
    <b-loading :active.sync="isLoading" :is-full-page="false" :can-cancel="false"></b-loading>
    <div class="setting-header">
      <img :src="require('@/assets/' + service + '.svg')" class="setting-logo" />
      {{ service.charAt(0).toUpperCase() + service.slice(1) }}
    </div>
    <div style="margin-top: 20px;">
      <span style="font-size:20px;">Required Fields</span>
      <setting v-for="x in configObj" v-bind:key="x.property" v-bind:setting="x"></setting>
    </div>
    <div class="btn-container">
      <b-button
        type="is-primary"
        @click="testConnection"
        :loading="isTesting"
        :disabled="service == 'filebot'"
        size="is-medium"
        outlined
        style="margin-right: 10px;"
      >Test Connection</b-button>
      <b-button size="is-medium" @click="save" :loading="isSaving" type="is-success">Save</b-button>
    </div>
  </div>
</template>

<script>
import axios from "axios";
import Setting from "@/components/Setting.vue";
import { SnackbarProgrammatic as Snackbar } from "buefy";

function createConfigData(configObj) {
  let configurationData = {}
  for (let setting of configObj) {
    switch (setting.datatype) {
      case "bool":
        configurationData[setting.property] = setting.bool_value;
        break;
      case "string":
        configurationData[setting.property] = setting.string_value;
        break;
      case "int":
        configurationData[setting.property] = setting.int_value;
        break;
    }
  }
  return configurationData
}

export default {
  name: "options",
  components: {
    Setting
  },
  data() {
    return {
      isLoading: true,
      isTesting: false,
      isSaving: false,
      configObj: []
    };
  },
  props: ["service"],
  mounted: function() {
    axios
      .get("/api/config/" + this.service, this.axiosAuthConfig())
      .then(response => {
        this.configObj = response.data;
        this.isLoading = false;
      })
      .catch(error => {
        this.checkUnauthorizedToken(error);
      });
  },
  methods: {
    save: function() {
      this.isSaving = true;
      let configurationData = createConfigData(this.configObj);
      axios
        .post(
          "/api/config/" + this.service,
          configurationData,
          this.axiosAuthConfig()
        )
        .then(() => {
          Snackbar.open({
            message: "Successfully saved settings!",
            type: "is-success",
            indefinite: false,
            duration: 2000
          });
        })
        .catch(error => {
          this.checkUnauthorizedToken(error);
          Snackbar.open({
            message: "Error saving settings",
            type: "is-danger",
            indefinite: false,
            duration: 3000
          });
        })
        .finally(() => {
          this.isSaving = false;
        });
    },
    testConnection: function() {
      this.isTesting = true;
      let configurationData = createConfigData(this.configObj);
      axios
        .post(
          "/api/testConfig/" + this.service,
          configurationData,
          this.axiosAuthConfig()
        )
        .then(() => {
          Snackbar.open({
            message: "Test successful",
            type: "is-success",
            indefinite: false,
            duration: 2000
          });
        })
        .catch((error) => {
          this.checkUnauthorizedToken(error);
          Snackbar.open({
            message: "Test unsuccessful",
            type: "is-danger",
            indefinite: false,
            duration: 3000
          });
        })
        .finally(() => {
          this.isTesting = false;
        });
    }
  }
};
</script>
