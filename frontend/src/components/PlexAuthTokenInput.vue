<template>
  <o-field label="Auth Token">
    <o-input
      expanded
      type="text"
      v-model="computedValue"
      placeholder="Auth Token"
    />
    <o-button @click="showFetchAuthTokenModal = true">Fetch New</o-button>
  </o-field>
  <modal
    v-model="showFetchAuthTokenModal"
    @close="showFetchAuthTokenModal = false"
    title="Fetch Plex Auth Token"
  >
    <div>
      <o-field label="Username">
        <o-input type="text" v-model="plexUsername" placeholder="Username" />
      </o-field>
      <o-field label="Password">
        <o-input
          type="password"
          v-model="plexPassword"
          placeholder="Password"
        />
      </o-field>
    </div>
    <template v-slot:footer>
      <o-button @click="showFetchAuthTokenModal = false">Cancel</o-button>
      <o-button variant="primary" @click="fetchPlexAuthToken">
        <action-button :mode="fetchTestingMode"> Fetch </action-button>
      </o-button>
    </template>
  </modal>
</template>

<script>
import APIUtil from "../util/APIUtil";
import ActionButton from "./ActionButton.vue";
import Modal from "./Modal.vue";

export default {
  data() {
    return {
      plexUsername: "",
      plexPassword: "",
      fetchTestingMode: "",
      showFetchAuthTokenModal: false,
    };
  },
  props: {
    modelValue: {
      type: String,
      default: function () {
        return "";
      },
    },
  },
  emits: ["update:modelValue"],
  components: { ActionButton, Modal },
  methods: {
    fetchPlexAuthToken() {
      this.fetchTestingMode = "loading";
      APIUtil.getPlexAuthToken(this.plexUsername, this.plexPassword)
        .then((response) => {
          this.computedValue = response.token;
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
    computedValue: {
      get() {
        return this.modelValue;
      },
      set(v) {
        this.$emit("update:modelValue", v);
      },
    },
  },
};
</script>
