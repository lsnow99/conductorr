<template>
  <div :class="mode?'loading':''">
    <slot />
  </div>
  <div v-if="mode === 'loading'" class="loader" />
  <div v-if="mode === 'success'" class="success" />
  <div v-if="mode === 'error'" class="error" />
</template>

<style scoped>
.loading {
  visibility: hidden;
}

.loader{
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
        return '';
      },
    },
  },
  computed: {
    contentVisible() {
      return !(this.mode === "loading" || this.mode === "success" || this.mode === "error")
    }
  }
};
</script>
