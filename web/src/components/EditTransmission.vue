<template>
  <modal
    title="Configure Transmission"
    v-model:active="computedActive"
    @close="$emit('close')"
  >
    <div>
      <o-field label="Name">
        <o-input type="text" v-model="computedValue.name" placeholder="Name" />
      </o-field>
      <o-field label="Base URL">
        <o-input
          type="text"
          v-model="computedValue.config.base_url"
          placeholder="Base URL"
        />
      </o-field>
      <o-field label="Username">
        <o-input
          type="text"
          v-model="computedValue.config.username"
          placeholder="transmission"
        />
      </o-field>
      <o-field label="Password">
        <o-input
          type="password"
          v-model="computedValue.config.password"
          placeholder="transmission"
        />
      </o-field>
      <o-field label="Post-Processing Action">
        <div class="flex flex-row justify-center mt-1">
          <radio-group
            name="fileAction"
            v-model="computedValue.file_action"
            :options="[
              { text: 'MOVE', value: 'move' },
              { text: 'COPY', value: 'copy' },
            ]"
          />
        </div>
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
import APIUtil from "../util/APIUtil";
import ActionButton from "./ActionButton.vue";
import EditServiceUtil from "../util/EditServiceUtil";
import Modal from "./Modal.vue";
import RadioGroup from "./RadioGroup.vue";

export default {
  data() {
    return {
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
      default: function () {
        return false;
      },
    },
    modelValue: {
      type: Object,
      default: function () {
        return {};
      },
    },
  },
  mixins: [EditServiceUtil],
  emits: ["submit", "close", "update:active", "update:modelValue"],
  components: {
    ActionButton,
    Modal,
    RadioGroup,
  },
  methods: {
    save() {
      this.$emit("submit", this.computedValue);
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
      APIUtil.testDownloader("transmission", this.computedValue)
        .then((resp) => {
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
      this.computedValue.name = this.computedValue.name
        ? this.computedValue.name.trim()
        : "";
      this.computedValue.config.base_url = this.computedValue.config.base_url
        ? this.computedValue.config.base_url.trim()
        : "";
    },
    validate() {
      if (!this.computedValue.name) {
        return "Name is required";
      } else if (!this.computedValue.config.base_url) {
        return "Base URL is required";
      } else if (!this.computedValue.config.username) {
        return "Username is required";
      } else if (!this.computedValue.config.password) {
        return "Password is required";
      }
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
