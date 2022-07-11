<template>
  <div :class="mode !== TestingMode.OFF ? 'loading' : ''">
    <slot />
  </div>
  <div v-if="mode === TestingMode.LOADING" class="loader" />
  <div v-if="mode === TestingMode.SUCCESS" class="success"><o-icon icon="check" /></div>
  <div v-if="mode === TestingMode.DANGER" class="danger"><o-icon icon="times" /></div>
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

<script setup lang="ts">
import { TestingMode } from '@/types/testing_mode';

defineProps<{
  mode: TestingMode
}>()
</script>
