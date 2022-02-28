<template>
<div class="flex flex-row w-full justify-center mt-16 text-gray-500 ">
    <div class="flex flex-col">
    <span>
    <vue-fontawesome class="opacity-50 text-[16rem]" icon="frown-open" />
    </span>
    <span class="text-center text-2xl mt-4">Page not found</span>
    <span class="text-center text-2xl mt-4">Click <router-link :to="{name: 'home'}">here</router-link> to go home</span>
    </div>
</div>
</template>

<script>
export default {

}
</script>
