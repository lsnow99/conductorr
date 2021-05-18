<template>
<div class="flex flex-row justify-center">
    <section class="max-w-lg mt-72">
        <o-field label="Username">
            <o-input type="text" v-model="username" />
        </o-field>
        <o-field label="Password">
            <o-input type="password" v-model="password" />
        </o-field>
        <div class="flex flex-row justify-between mt-2">
            <div />
            <o-button :loading="true" @click="submit">Login</o-button>
        </div>
    </section>
    <o-loading :full-page="true" v-model:active="loading" :can-cancel="false"></o-loading>
</div>
</template>

<script>
import APIUtil from "../util/APIUtil"

export default {
    data() {
        return {
            username: "",
            password: "",
            loading: false
        }
    },
    methods: {
        submit() {
            this.loading = true;
            APIUtil.signIn().then(() => {
                console.log('success')
                this.$router.push({ name: 'home' })
            }).catch((err) => {
                console.log(err)
            }).finally(() => {
                this.loading = false;
            })
        }
    }
}
</script>