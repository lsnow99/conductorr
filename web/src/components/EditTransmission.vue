<template>
<modal title="Configure Transmission" v-model:active="computedActive" @close="$emit('close')">
    <div>
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
          v-model="computedTransmission.base_url"
          placeholder="Base URL"
        />
      </o-field>
      <o-field label="Username">
        <o-input
          type="text"
          v-model="computedTransmission.username"
          placeholder="transmission"
        />
      </o-field>
      <o-field label="Password">
        <o-input
          type="password"
          v-model="computedTransmission.password"
          placeholder="transmission"
        />
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
</template>

<script>
import APIUtil from '../util/APIUtil';
import ActionButton from "./ActionButton.vue";
import EditServiceUtil from "../util/EditServiceUtil";
import Modal from "./Modal.vue"

export default {
  data() {
    return {
      newTransmission: null,
      testingMode: "",
    };
  },
  props: {
    transmission: {
      type: Object,
      default: function () {
        return {};
      },
    },
    active: {
      type: Boolean,
      default: function() {
        return false
      }
    }
  },
  mixins: [ EditServiceUtil ],
  emits: ["submit", "close", "update:active"],
  components: {
    ActionButton,
    Modal
  },
  methods: {
    save() {
      this.$emit("submit", this.newTransmission);
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
      APIUtil.testDownloader("transmission", this.computedTransmission)
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
      this.computedTransmission.base_url = this.computedTransmission.base_url
        ? this.computedTransmission.base_url.trim()
        : undefined;
    },
    validate() {
      if (!this.computedName) {
        return "Name is required";
      } else if (!this.computedTransmission.base_url) {
        return "Base URL is required";
      } else if (!this.computedTransmission.username) {
        return "Username is required";
      } else if (!this.computedTransmission.password) {
        return "Password is required";
      }
    },
  },
  computed: {
    computedTransmission: {
      get() {
        if (this.newTransmission == null) {
          if (this.transmission) {
            this.newTransmission = this.transmission;
          }
        }
        return this.newTransmission;
      },
      set(newVal) {
        this.newTransmission = newVal;
      },
    },
    computedActive: {
      get() {
        return this.active
      },
      set(v) {
        this.$emit('update:active', v)
      }
    }
  },
};
</script>
