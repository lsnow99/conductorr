<template>
  <div @click="$emit('click', media)" class="media-card relative">
    <img :src="media.poster" :alt="`Banner image for movie ${media.title}`" />
    <div class="absolute top-0 bottom-0 left-0 right-0 overlay">
      <div class="flex flex-row h-full">
        <div class="p-2 text-2xl self-end font-bold z-10">
          {{ `${media.title} (${year})` }}
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.media-card {
  height: 22rem;
  width: 16rem;
  transition: all 0.1s linear;

  @apply bg-gray-400;
  @apply rounded-lg;
  @apply overflow-hidden;
  @apply cursor-pointer;
}

.media-card:hover {
  box-shadow: 0 12px 18px -4px rgba(41, 41, 41, 0.75);
  -webkit-box-shadow: 0 12px 18px -4px rgba(41, 41, 41, 0.75);
  -moz-box-shadow: 0 12px 18px -4px rgba(41, 41, 41, 0.75);
}

.overlay {
  background: rgb(0, 0, 0);
  background: linear-gradient(
    180deg,
    rgba(0, 0, 0, 0) 0%,
    rgba(9, 9, 121, 0) 50%,
    rgba(0, 0, 0, 1) 100%
  );
  z-index: 2;
}

.overlay::before {
  position: absolute;
  content: "";
  top: 0;
  right: 0;
  bottom: 0;
  left: 0;
  background: rgb(0, 0, 0);
  background: linear-gradient(
    180deg,
    rgba(0, 0, 0, 0.7) 0%,
    rgba(9, 9, 121, 0) 50%,
    rgba(0, 0, 0, 0.7) 100%
  );
  z-index: 1;
  transition: opacity 0.1s linear;
  opacity: 0;
}

.overlay:hover::before {
  opacity: 1;
}
</style>

<script>
import MediaUtil from "../util/MediaUtil"

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
  mixins: [MediaUtil],
  emits: ["click"],
  computed: {
    year() {
      return new Date(this.media.released_at).getUTCFullYear();
    },
  },
};
</script>
