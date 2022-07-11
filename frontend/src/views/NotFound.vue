<template>
<div class="flex flex-row justify-center w-full mt-16 text-gray-500 ">
    <div class="flex flex-col">
    <span>
    <vue-fontawesome class="opacity-50 text-[16rem]" icon="frown-open" />
    </span>
    <span class="mt-4 text-2xl text-center">Page not found</span>
    <span class="mt-4 text-2xl text-center">Click <router-link :to="{name: 'home'}">here</router-link> to go home</span>
    </div>
</div>
</template>

<script setup lang="ts">
</script>
