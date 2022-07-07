<template>
  <Modal :title="title" v-model:active="computedActive" @close="emit('close')">
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
import { useComputedActive, useComputedValue } from "@/util";
import { inject } from "vue";

const oruga = inject("oruga")

const props = withDefaults(defineProps<{
  modelValue: ConfigurableService,
  active: boolean,
  fields: any[],
  title: string,
  extraSanitizer: () => void,
  extraValidator: () => string | null,
  testingMode: string
}>(), {
  extraSanitizer: () => {},
  extraValidator: () => null
})

const emit = defineEmits<{
  (e: "update:modelValue", newVal: ConfigurableService): void,
  (e: "update:active", newVal: boolean): void,
  (e: "save", savedVal: ConfigurableService): Promise<void>,
  (e: "test", savedVal: ConfigurableService): Promise<void>,
  (e: "close"): void
}>()

const computedValue = useComputedValue<ConfigurableService>(props, emit)
const computedActive = useComputedActive(props, emit)

const sanitize = () => {
  computedValue.value.name = computedValue.value.name?.trim() ?? ""
  for (const field of props.fields) {
    if (field.trim) {
      computedValue.value.config[field.property] = computedValue.value.config[
        field.property
      ]
        ? computedValue.value.config[field.property].trim()
        : "";
    }
  }
  props.extraSanitizer();
}

const validate = (configOnly: boolean = false) => {
  if(!computedValue.value.name && !configOnly) {
    return "Name is required"
  }
  for (const field of props.fields) {
    if (field.required && !computedValue.value.config[field.property]) {
      return `${field.label} is required`;
    }
  }
  return props.extraValidator();
}

const test = () => {
  sanitize();
  emit("test", computedValue.value)
}

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
    return
  }
  emit("save", computedValue.value)
}

</script>
