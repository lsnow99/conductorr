<template>
    <div class="container">
        <div class="setting-header">
            <b-icon
                icon="lock"
                type="is-success">
            </b-icon>
             Security
        </div>
        <div class="tile is-parent is-vertical">
            <div class="tile is-child">
                <span class="header-title">
                    <b-icon
                        style="vertical-align: middle"
                        icon="shield-alt"
                        size="is-medium">
                    </b-icon>
                    User Account
                </span>
                <setting v-bind:setting="username"></setting>
                <setting v-bind:setting="password"></setting>
            </div>
        </div>
        <div class="btn-container">
            <b-button size="is-medium" @click="save" :loading="isSaving" type="is-success">Save</b-button>
        </div>
    </div>
</template>

<script>
import { SnackbarProgrammatic as Snackbar } from 'buefy'
import axios from 'axios'
import Setting from '@/components/Setting.vue'

export default {
    name: 'security',
    components: {
        Setting,
    },
    data() {
        return {
            username: {
                datatype: 'string',
                label: 'Username',
                placeholder: 'admin',
                string_value: 'admin',
                description: 'Username for Conductorr web account'
            },
            password: {
                datatype: 'password',
                label: 'Password',
                string_value: 'supersecurepwd',
                description: 'Set a secure password for your Conductorr web account'
            },
            isSaving: false
        }
    },
    methods: {
        save() {
            this.isSaving = true
            axios.post("/backend/config/system", {})
            .then(() => {
                Snackbar.open({
                    message: 'Succcessfully saved settings!',
                    type: 'is-success',
                    indefinite: false,
                    duration: 3000,
                    })
            })
            .catch(() => {
                Snackbar.open({
                    message: 'Error saving settings',
                    type: 'is-danger',
                    indefinite: false,
                    duration: 10000,
                    })
            })
            .finally(() => {
                this.isSaving = false
            })
        }
    }
}
</script>

<style scoped>

</style>