<template>
  <div
    class="inline-block mt-4 overflow-hidden rounded-lg sm:mt-0 radio-wrapper"
  >
    <template v-for="option in options" :key="option.value">
      <o-radio
        v-model="computedValue"
        :name="name"
        :native-value="option.value"
      >
        {{ option.text }}
      </o-radio>
    </template>
  </div>
</template>

<script setup lang="ts">
import { useComputedValue } from "conductorr-lib";

const props = defineProps<
  {
    name: string;
    options: {
      text: string;
      value: unknown;
    }[];
    modelValue: string | null
  }
>();

const emit = defineEmits<{
  (e: "update:modelValue", newVal: unknown): void
}>();

const computedValue = useComputedValue<string | null>(props, emit);
</script>
