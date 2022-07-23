<template>
  <span
    :aria-disabled="disabled"
    @click.stop="doToggle"
    @keydown.enter.stop="doToggle"
    @keydown.space.stop="doToggle"
    tabindex="0"
    role="button"
    :aria-label="monitoring ? 'Unmonitor media' : 'Monitor media'"
    :class="computedClass"
    ><o-icon v-if="monitoring" icon="bookmark" pack="fas" />
    <o-icon v-else icon="bookmark" pack="far" />
  </span>
</template>

<style scoped>
.disabled {
  @apply opacity-60;
  @apply cursor-default;
}
</style>

<script setup lang="ts">
import { computed } from '@vue/reactivity';

const props = defineProps<{
  class: string,
  monitoring: boolean,
  disabled?: boolean
}>()

const emit = defineEmits<{
  (e: "toggle"): void
}>()

const doToggle = () => {
  if (!props.disabled) {
    emit("toggle")
  }
}

const computedClass = computed(() => {
  let disabledClass = "";
  if (props.disabled) {
    disabledClass = "disabled";
  }
  return props.class + " cursor-pointer " + disabledClass
})
</script>
