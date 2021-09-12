<template>
  <o-modal
    v-model:active="computedValue"
    role="dialog"
    :aria-label="title"
    :full-screen="fullScreen"
    aria-modal
    @close="$emit('close')"
  >
    <header class="modal-card-header">
      <slot name="header">
        <p class="text-3xl font-light">
          {{ title }}
        </p>
      </slot>
    </header>
    <section class="border-t-2 border-b-2 border-gray-300 p-4 flex-1 overflow-y-scroll">
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
      default: function() {
        return false
      }
    }
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
