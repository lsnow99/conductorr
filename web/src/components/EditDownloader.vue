<template>
  <modal :title="title" v-model:active="computedActive" @close="$emit('close')">
    <div>
      <o-field label="Name">
        <o-input type="text" v-model="computedValue.name" placeholder="Name" />
      </o-field>
      <o-field
        :label="field.label"
        v-for="field in fields"
        :key="field.property"
      >
        <o-input
          :type="field.type"
          v-model="computedValue.config[field.property]"
          :placeholder="field.placeholder"
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
import Modal from "./Modal.vue";
import ActionButton from "./ActionButton.vue";
import RadioGroup from "./RadioGroup.vue";
import APIUtil from "../util/APIUtil";

export default {
  data() {
    return {
      testingMode: "",
    };
  },
  props: {
    modelValue: {
      type: Object,
      default: function () {
        return {
          config: {},
        };
      },
    },
    active: {
      type: Boolean,
      default: function () {
        return false;
      },
    },
    fields: {
      type: Array,
      default: function () {
        return [];
      },
    },
    downloaderType: {
      type: String,
      default: function () {
        return "";
      },
    },
    title: {
      type: String,
      default: function () {
        return "";
      },
    },
  },
  components: {
    Modal,
    ActionButton,
    RadioGroup,
  },
  emits: ["update:modelValue", "update:active", "save", "close"],
  methods: {
    test() {
      this.sanitize();
      const validationErr = this.validate(true);
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
      APIUtil.testDownloader(this.downloaderType, this.computedValue.config)
        .then((resp) => {
            console.log(resp)
            this.$oruga.notification.open({
              duration: 5000,
              message: `Connected successfully`,
              position: "bottom-right",
              variant: "success",
              closable: false,
            });
            this.testingMode = "success";
        })
        .catch(err => {
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
    save() {
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
      this.$emit("save", this.computedValue);
    },
    validate(configOnly) {
      if (!this.computedValue.name && !configOnly) {
        return "Name is required";
      }
      for (const field of this.fields) {
        if (field.required && !this.computedValue.config[field.property]) {
          return `${field.label} is required`;
        }
      }
      if (!this.computedValue.file_action && !configOnly) {
        return "File Action is required";
      }
    },
    sanitize() {
      this.computedValue.name = this.computedValue.name
        ? this.computedValue.name.trim()
        : "";
      for (const field of this.fields) {
          console.log(field)
        if (field.trim) {
          this.computedValue.config[field.property] = this.computedValue.config[
            field.property
          ]
            ? this.computedValue.config[field.property].trim()
            : "";
        }
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
