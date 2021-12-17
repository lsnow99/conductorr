<template>
    Logging you out
</template>

<script>
import APIUtil from '../util/APIUtil'

export default {
    mounted() {
        APIUtil.logout().then(() => {
            this.$store.commit('setLoggedIn', false)
            if (!this.$store.getters.loggedIn) {
                this.$router.push({name: 'home'})
            }
        })
    }
}
</script>