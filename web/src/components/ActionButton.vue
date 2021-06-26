<template>
  <div :class="mode ? 'loading' : ''">
    <slot />
  </div>
  <div v-if="mode === 'loading'" class="loader" />
  <div v-if="mode === 'success'" class="success"><o-icon icon="check" /></div>
  <div v-if="mode === 'danger'" class="danger"><o-icon icon="times" /></div>
</template>

<style scoped>
.danger,
.success {
  @apply absolute;
  @apply top-1;
  @apply right-0;
  @apply left-0;
  @apply ml-auto;
  @apply mr-auto;
  @apply text-center;
  @apply text-xl;
}

.danger {
  @apply text-red-500;
}

.success {
  @apply text-green-500;
}

.loading {
  visibility: hidden;
}

.loader {
  content: "";
  position: absolute;
  top: 50%;
  left: 50%;
  margin-left: -10px;
  margin-top: -10px;
  width: 20px;
  height: 20px;
  display: block;
  opacity: 1;
  background-image: url("../assets/loading.svg");
  background-size: cover;
  animation-name: spin;
  animation-duration: 450ms;
  animation-iteration-count: infinite;
  animation-timing-function: linear;
}

@keyframes spin {
  from {
    transform: rotate(0deg);
  }

  to {
    transform: rotate(360deg);
  }
}
</style>

<script>
export default {
  props: {
    mode: {
      type: String,
      default: function () {
        return "";
      },
    },
  },
  computed: {
    contentVisible() {
      return !(
        this.mode === "loading" ||
        this.mode === "success" ||
        this.mode === "danger"
      );
    },
  },
};
</script>
