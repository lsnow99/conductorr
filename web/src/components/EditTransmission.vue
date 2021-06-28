<template>
  <header class="modal-card-header">
    <p class="modal-card-title">Configure Transmission</p>
  </header>
  <section class="modal-card-content">
    <div>
      <o-field label="Name">
        <o-input
          type="text"
          v-model="computedTransmission.name"
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
  },
  components: {
    ActionButton,
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
          console.log(resp)
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
      this.computedTransmission.name = this.computedTransmission.name
        ? this.computedTransmission.name.trim()
        : undefined;
      this.computedTransmission.base_url = this.computedTransmission.base_url
        ? this.computedTransmission.base_url.trim()
        : undefined;
    },
    validate() {
      if (!this.computedTransmission.name) {
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
  },
};
</script>
