<template>
<router-view></router-view>
<toaster />
</template>

<script>
import Toaster from "./components/Toaster.vue"
export default {
  components: {
    Toaster
  }
}
</script>

<style>
body, html {
  @apply dark:text-white;
  @apply dark:bg-gray-800;
}
</style>
