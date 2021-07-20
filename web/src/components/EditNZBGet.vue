<template>
  <header class="modal-card-header">
    <p class="modal-card-title">Configure NZBGet</p>
  </header>
  <section class="modal-card-content">
    <div>
      {{computedTest}}
      <o-field label="Name">
        <o-input
          type="text"
          v-model="computedName"
          placeholder="Name"
        />
      </o-field>
      <o-field label="Base URL">
        <o-input
          type="text"
          v-model="computedNZBGet.base_url"
          placeholder="Base URL"
        />
      </o-field>
      <o-field label="Username">
        <o-input
          type="text"
          v-model="computedNZBGet.username"
          placeholder="nzbget"
        />
      </o-field>
      <o-field label="Password">
        <o-input
          type="password"
          v-model="computedNZBGet.password"
          placeholder="tegbzn6789"
        />
      </o-field>
    </div>
  </section>
  <footer class="modal-card-footer">
    <o-button @click="$emit('close')">Cancel</o-button>
    <div>
      <o-button variant="primary" @click="test" class="mr-3">
        <action-button :mode="testingMode"> Test </action-button>
      </o-button>
      <o-button variant="primary" @click="save">Save</o-button>
    </div>
  </footer>
</template>

<script>
import APIUtil from '../util/APIUtil';
import ActionButton from "./ActionButton.vue";
import DownloaderUtil from "../util/DownloaderUtil";

export default {
  data() {
    return {
      newNZBGet: null,
      testingMode: "",
    };
  },
  props: {
    nzbget: {
      type: Object,
      default: function () {
        return {};
      },
    },
  },
  mixins: [ DownloaderUtil ],
  components: {
    ActionButton,
  },
  methods: {
    save() {
      this.$emit("submit", this.newNZBGet);
    },
    test() {
      this.sanitize();
      const validationErr = this.validate();
      if (validationErr) {
        this.$oruga.notification.open({
          duration: 5000,
          message: validationErr,
          position: "bottom-right",
          variant: "danger",
          closable: false,
        });
        return;
      }
      this.testingMode = "loading";
      APIUtil.testDownloader("nzbget", this.computedNZBGet)
        .then((resp) => {
          if (resp.success) {
            this.$oruga.notification.open({
              duration: 5000,
              message: `Connected successfully`,
              position: "bottom-right",
              variant: "success",
              closable: false,
            });
            this.testingMode = "success";
          } else {
            this.$oruga.notification.open({
              duration: 5000,
              message: `Test failed: ${resp.msg}`,
              position: "bottom-right",
              variant: "danger",
              closable: false,
            });
            this.testingMode = "danger";
          }
        })
        .finally(() => {
          setTimeout(() => {
            this.testingMode = "";
          }, 5000);
        });
    },
    sanitize() {
      this.computedName = this.computedName
        ? this.computedName.trim()
        : undefined;
      this.computedNZBGet.base_url = this.computedNZBGet.base_url
        ? this.computedNZBGet.base_url.trim()
        : undefined;
    },
    validate() {
      if (!this.computedName) {
        return "Name is required";
      } else if (!this.computedNZBGet.base_url) {
        return "Base URL is required";
      } else if (!this.computedNZBGet.username) {
        return "Username is required";
      } else if (!this.computedNZBGet.password) {
        return "Password is required";
      }
    },
  },
  computed: {
    computedNZBGet: {
      get() {
        if (this.newNZBGet == null) {
          if (this.nzbget) {
            this.newNZBGet = this.nzbget;
          }
        }
        return this.newNZBGet;
      },
      set(newVal) {
        this.newNZBGet = newVal;
      },
    },
  },
};
</script>
