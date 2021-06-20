<template>
  <div class="toasts">
    <transition-group name="fade">
      <div
        class="toast"
        :class="classType(toast.type)"
        v-for="(toast, index) in toasts"
        :key="index"
      >
        {{ toast.msg }}
      </div>
    </transition-group>
  </div>
</template>

<style scoped>
.toast {
  @apply rounded-md;
  @apply text-white;
  @apply mt-4 shadow-lg py-3 px-6 text-lg font-medium max-w-lg;
}
.toasts {
  @apply absolute;
  @apply bottom-3;
  @apply right-3;
}
</style>

<script>
export default {
  methods: {
    classType(type) {
      if (type === "error") {
        return `bg-red-400`;
      } else if (type === "warning") {
        return `bg-yellow-400`;
      } else if (type === "success") {
        return `bg-green-400`;
      }
    },
  },
  computed: {
    toasts() {
      return this.$store.getters.toasts;
    },
  },
};
</script>
