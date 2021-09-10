<template>
  <o-modal
    v-model:active="computedValue"
    role="dialog"
    :aria-label="title"
    aria-modal
    @close="$emit('close')"
  >
    <header class="modal-card-header">
      <slot name="header">
        <p class="modal-card-title">
          {{ title }}
        </p>
      </slot>
    </header>
    <section class="modal-card-content">
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
