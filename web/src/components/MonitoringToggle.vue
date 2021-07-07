<template>
  <span :class="computedClass"
    ><o-icon
      v-if="monitoring"
      icon="bookmark"
      pack="fas"
      :aria-disabled="disabled"
      @click="doToggle"
    />
    <o-icon
      v-if="!monitoring"
      icon="bookmark"
      pack="far"
      :aria-disabled="disabled"
      @click="doToggle"
    />
  </span>
</template>

<style scoped>
.disabled {
  @apply opacity-60;
  @apply cursor-default;
}
</style>

<script>
export default {
  props: {
    class: {
      type: String,
      default: function () {
        return "";
      },
    },
    monitoring: {
      type: Boolean,
      default: function () {
        return false;
      },
    },
    disabled: {
      type: Boolean,
      default: function () {
        return false;
      },
    },
  },
  emits: ["toggle"],
  methods: {
    doToggle() {
      if (!this.disabled) {
        this.$emit("toggle");
      }
    },
  },
  computed: {
    computedClass() {
      let disabledClass = "";
      if (this.disabled) {
        disabledClass = "disabled";
      }
      return this.class + " cursor-pointer " + disabledClass;
    },
  },
};
</script>
