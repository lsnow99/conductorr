<template>
      <o-skeleton class="media-card" height="22rem">
      </o-skeleton>
  <!-- <div class="media-card">
  </div> -->
</template>

<style scoped>
.media-card {
  height: 22rem;
  width: 16rem;
  @apply bg-gray-400;
  @apply rounded-lg;
  @apply overflow-hidden;
}
</style>

<script>
export default {
  props: {
    media: {
      type: Object,
      default: function () {
        return {
          title: "",
          poster: "",
        };
      },
    },
  },
};
</script>
