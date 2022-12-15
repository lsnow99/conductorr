<template>
  <Modal :title="title" v-model="computedActive" @close="emit('close')">
    <div>
      <o-field label="Name">
        <o-input type="text" v-model="computedValue.name" placeholder="Name" />
      </o-field>
      <template v-for="field in fields" :key="field.property">
        <template v-if="field.component">
          <component
            :is="field.component"
            v-model="computedValue[field.property]"
          />
        </template>
        <o-field v-else :label="field.label">
          <o-input
            :type="field.type"
            v-model="computedValue[field.property]"
            :placeholder="field.placeholder"
          />
        </o-field>
      </template>
      <slot />
    </div>
    <template v-slot:footer>
      <o-button @click="emit('close')">Cancel</o-button>
      <div>
        <o-button variant="primary" @click="test" class="mr-3">
          <action-button :mode="testingMode"> Test </action-button>
        </o-button>
        <o-button variant="primary" @click="save">Save</o-button>
      </div>
    </template>
  </Modal>
</template>

<script setup lang="ts">
import Modal from "./Modal.vue";
import ActionButton from "./ActionButton.vue";
import { ConfigurableService } from "@/types/api/service";
import { useComputedActive } from "@/util";
import { inject, onMounted, computed } from "vue";
import { TestingMode } from "@/types/testing_mode";
import { useComputedValue } from "conductorr-lib"

const oruga = inject("oruga");

const props = withDefaults(
  defineProps<{
    modelValue: ConfigurableService;
    active: boolean;
    fields: any[];
    title: string;
    extraSanitizer?: () => void;
    extraValidator?: () => string | null;
    testingMode: TestingMode;
  }>(),
  {
    extraSanitizer: () => {},
    extraValidator: () => null,
  }
);

const emit = defineEmits<{
  (e: "update:modelValue", newVal: ConfigurableService): void;
  (e: "update:active", newVal: boolean): void;
  (e: "save", savedVal: ConfigurableService): Promise<void>;
  (e: "test", savedVal: ConfigurableService): Promise<void>;
  (e: "close"): void;
}>();

const computedActive = useComputedActive(props, emit);

const computedValue: WritableComputedRef<ConfigurableService> = computed({
  get() {
    return props.modelValue ?? {}
  },
  set(v) {
    emit("update:modelValue", v)
    }
    })

const sanitize = () => {
  for (const field of props.fields) {
    if (field.trim) {
      computedValue.value[field.property] =
        (computedValue.value[field.property] as (string | undefined | null))?.trim() ?? "";
    }
  }
  props.extraSanitizer();
};

const validate = (configOnly: boolean = false) => {
  if (!computedValue.value.name && !configOnly) {
    return "Name is required";
  }
  for (const field of props.fields) {
    if (field.required && !computedValue.value[field.property]) {
      return `${field.label} is required`;
    }
  }
  return props.extraValidator();
};

const test = () => {
  sanitize();
  emit("test", computedValue.value);
};

const save = () => {
  sanitize();
  const validationErr = validate();
  if (validationErr) {
    oruga.notification.open({
      duration: 5000,
      message: validationErr,
      position: "bottom-right",
      variant: "danger",
      closable: false,
    });
    return;
  }
  emit("save", computedValue.value);
};
</script>
