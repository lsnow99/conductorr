<template>
  <o-modal
    v-model:active="computedValue"
    role="dialog"
    :full-screen="fullScreen"
    :aria-label="title"
    aria-modal
    auto-focus
    trap-focus
    :can-cancel="['escape', 'outside']"
    @close="emit('close')"
  >
    <header class="modal-card-header">
      <slot name="header">
        <h1 class="text-3xl font-light">
          {{ title }}
        </h1>
        <span
          class="ml-3 text-2xl cursor-pointer"
          role="button"
          tabindex="0"
          @click="doClose"
          @keydown.enter="doClose"
          @keydown.space="doClose"
          ><o-icon icon="times"></o-icon
        ></span>
      </slot>
    </header>
    <section
      class="flex-1 p-4 overflow-y-scroll border-t-2 border-b-2 border-gray-300"
    >
      <slot />
    </section>
    <footer class="modal-card-footer">
      <slot name="footer">
        <o-button
          @click="doClose"
          @keydown.space="doClose"
          @keydown.enter="doClose"
          role="button"
          tabindex="0"
          >Close</o-button
        >
      </slot>
    </footer>
  </o-modal>
</template>

<style scoped>
.modal-card-header,
.modal-card-footer {
  @apply flex;
  @apply flex-row;
  @apply justify-between;
  @apply px-4;
  @apply py-3;
  @apply items-center;
}
</style>

<script setup lang="ts">
import { useComputedValue } from '@/util';

const props = defineProps<{
  modelValue: boolean,
  title: string,
  fullScreen: boolean
}>()

const emit = defineEmits<{
  (e: "close"): void,
  (e: "update:modelValue", newVal: boolean): void
}>()

const computedValue = useComputedValue<boolean>(props, emit)

const doClose = () => {
  emit("close")
  computedValue.value = false;
}
</script>
