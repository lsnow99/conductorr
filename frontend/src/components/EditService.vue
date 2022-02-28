<template>
  <Modal :title="title" v-model:active="computedActive" @close="$emit('close')">
    <div>
      <o-field label="Name">
        <o-input type="text" v-model="computedValue.name" placeholder="Name" />
      </o-field>
      <template v-for="field in fields" :key="field.property">
        <template v-if="field.component">
          <component
            :is="field.component"
            v-model="computedValue.config[field.property]"
          />
        </template>
        <o-field v-else :label="field.label">
          <o-input
            :type="field.type"
            v-model="computedValue.config[field.property]"
            :placeholder="field.placeholder"
          />
        </o-field>
      </template>
      <slot />
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
  </Modal>
</template>

<script>
import Modal from "./Modal.vue";
import ActionButton from "./ActionButton.vue";

export default {
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
    title: {
      type: String,
      default: function () {
        return "";
      },
    },
    extraSanitizer: {
      type: Function,
      default() {},
    },
    extraValidator: {
      type: Function,
      default() {},
    },
    testingMode: {
      type: String,
      default: function () {
        return "";
      },
    },
  },
  components: {
    Modal,
    ActionButton,
  },
  emits: ["update:modelValue", "update:active", "save", "close", "test"],
  methods: {
    test() {
      this.sanitize();
      this.$emit("test", this.computedValue);
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
      return this.extraValidator();
    },
    sanitize() {
      this.computedValue.name = this.computedValue.name
        ? this.computedValue.name.trim()
        : "";
      for (const field of this.fields) {
        if (field.trim) {
          this.computedValue.config[field.property] = this.computedValue.config[
            field.property
          ]
            ? this.computedValue.config[field.property].trim()
            : "";
        }
      }
      this.extraSanitizer();
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
