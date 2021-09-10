<template>
  <div
    class="overflow-hidden rounded-lg inline-block mt-4 sm:mt-0 radio-wrapper"
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

<script>
export default {
  props: {
    modelValue: {
      type: String,
      default: function () {
        return "";
      },
    },
    name: {
      type: String,
      default: function () {
        return "";
      },
    },
    options: {
      type: Array,
      default: function () {
        return [];
      },
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
