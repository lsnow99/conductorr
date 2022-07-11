<template>
    Logging you out
</template>

<script setup lang="ts">
import { useAppStore } from '@/store';
import { storeToRefs } from 'pinia';
import { onMounted } from 'vue';
import { useRouter } from 'vue-router';
import APIUtil from '../util/APIUtil'

const { loggedIn } = storeToRefs(useAppStore())
const router = useRouter()

onMounted(async() => {
    await APIUtil.logout()
    loggedIn.value = false
    router.push({name: 'home'})
})
</script>