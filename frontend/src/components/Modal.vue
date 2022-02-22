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
    @close="$emit('close')"
  >
    <header class="modal-card-header">
      <slot name="header">
        <h1 class="text-3xl font-light">
          {{ title }}
        </h1>
        <span
          class="cursor-pointer text-2xl"
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
      class="border-t-2 border-b-2 border-gray-300 p-4 flex-1 overflow-y-scroll"
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

<script>
export default {
  props: {
    modelValue: {
      type: Boolean,
      default: function () {
        return false;
      },
    },
    title: {
      type: String,
      default: function () {
        return "";
      },
    },
    fullScreen: {
      type: Boolean,
      default: function () {
        return false;
      },
    },
  },
  emits: ["close", "update:modelValue"],
  methods: {
    doClose() {
      this.$emit("close");
      this.computedValue = false;
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
  },
};
</script>
