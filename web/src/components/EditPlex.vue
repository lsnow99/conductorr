<template>
  <modal
    v-model="computedActive"
    @close="$emit('close')"
    title="Configure Plex"
  >
    <div>
      <o-field label="Name">
        <o-input type="text" v-model="computedName" placeholder="Name" />
      </o-field>
      <o-field label="Base URL">
        <o-input
          type="text"
          v-model="computedPlex.base_url"
          placeholder="Base URL"
        />
      </o-field>
      <o-field label="Auth Token">
        <o-input
          expanded
          type="text"
          v-model="computedPlex.token"
          placeholder="Auth Token"
        />
        <o-button @click="showFetchAuthTokenModal = true">Fetch New</o-button>
      </o-field>
    </div>
    <template v-slot:footer>
      <o-button @click="$emit('close')">Cancel</o-button>
      <div>
        <o-button variant="primary" @click="test" class="mr-3">
          <action-button :mode="testingMode"> Test </action-button>
        </o-button>
        <o-button variant="primary" @click="save">Save</o-button>
      </div>
    </template>
  </modal>
  <modal
    v-model="showFetchAuthTokenModal"
    @close="showFetchAuthTokenModal = false"
    title="Fetch Plex Auth Token"
  >
    <div>
      <o-field label="Username">
        <o-input type="text" v-model="username" placeholder="Username" />
      </o-field>
      <o-field label="Password">
        <o-input type="password" v-model="password" placeholder="Password" />
      </o-field>
    </div>
    <template v-slot:footer>
      <o-button @click="showFetchAuthTokenModal = false">Cancel</o-button>
      <o-button variant="primary" @click="fetchAuthToken">
        <action-button :mode="fetchTestingMode"> Fetch </action-button>
      </o-button>
    </template>
  </modal>
</template>

<script>
import APIUtil from "../util/APIUtil";
import ActionButton from "./ActionButton.vue";
import DownloaderUtil from "../util/EditServiceUtil";
import Modal from "../components/Modal.vue";

export default {
  data() {
    return {
      newPlex: null,
      testingMode: "",
      fetchTestingMode: "",
      showFetchAuthTokenModal: false,
      username: "",
      password: "",
    };
  },
  props: {
    plex: {
      type: Object,
      default: function () {
        return {};
      },
    },
    active: {
      type: Boolean,
      default: function () {
        return false;
      },
    },
  },
  mixins: [DownloaderUtil],
  components: {
    ActionButton,
    Modal,
  },
  emits: ["submit", "update:active", "close"],
  methods: {
    save() {
      this.$emit("submit", this.newPlex);
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
      APIUtil.testMediaServer("plex", this.computedPlex)
        .then(() => {
          this.$oruga.notification.open({
            duration: 5000,
            message: `Connected successfully`,
            position: "bottom-right",
            variant: "success",
            closable: false,
          });
          this.testingMode = "success";
        })
        .catch((err) => {
          this.$oruga.notification.open({
            duration: 5000,
            message: `Test failed: ${err.msg}`,
            position: "bottom-right",
            variant: "danger",
            closable: false,
          });
          this.testingMode = "danger";
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
      this.computedPlex.base_url = this.computedPlex.base_url
        ? this.computedPlex.base_url.trim()
        : undefined;
    },
    validate() {
      if (!this.computedName) {
        return "Name is required";
      } else if (!this.computedPlex.base_url) {
        return "Base URL is required";
      } else if (!this.computedPlex.token) {
        return "Token is required";
      }
    },
    fetchAuthToken() {
      this.fetchTestingMode = "loading";
      APIUtil.getPlexAuthToken(this.username, this.password)
        .then((response) => {
          this.computedPlex.token = response.token;
          this.showFetchAuthTokenModal = false;
          this.fetchTestingMode = "";
          this.$oruga.notification.open({
            duration: 5000,
            message: `Successfully fetched Plex auth token`,
            position: "bottom-right",
            variant: "success",
            closable: false,
          });
        })
        .catch((err) => {
          this.fetchTestingMode = "danger";
          this.$oruga.notification.open({
            duration: 5000,
            message: `Plex login failed: ${err}`,
            position: "bottom-right",
            variant: "danger",
            closable: false,
          });
        });
    },
  },
  computed: {
    computedPlex: {
      get() {
        if (this.newPlex == null) {
          if (this.plex) {
            this.newPlex = this.plex;
          }
        }
        return this.newPlex;
      },
      set(v) {
        this.newPlex = v;
      },
    },
    computedActive: {
      get() {
        return this.active;
      },
      set(v) {
        this.$emit("update:active", v);
      },
    },
  },
};
</script>
