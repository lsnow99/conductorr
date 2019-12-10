<template>
    <div class="container">
        <b-loading :active.sync="isLoading" :is-full-page="false" :can-cancel="false"></b-loading>
        <div class="setting-header"><img :src="require('@/assets/' + service + '.svg')" class="setting-logo"> {{service.charAt(0).toUpperCase() + service.slice(1)}}</div>
        <div style="margin-top: 20px;" >
            <span style="font-size:20px;">Required Fields</span>
            <setting v-for="x in configObj" v-bind:key="x.property" v-bind:setting="x"></setting>  
        </div>
        <div class="btn-container">
            <b-button type="is-primary" @click="testConnection" :loading="isTesting" :disabled="service == 'filebot'" size="is-medium" outlined style="margin-right: 10px;">Test Connection</b-button>
            <b-button size="is-medium" @click="save" :loading="isSaving" type="is-success">Save</b-button>
        </div>
    </div>
</template>

<script>
import axios from 'axios'
import Setting from '@/components/Setting.vue'
import { SnackbarProgrammatic as Snackbar } from 'buefy'

export default {
    name: 'options',
    components: {
        Setting
    },
    data() {
        return {
            isLoading: true,
            isTesting: false,
            isSaving: false,
            configObj: []
        }
    },
    props: [
        'service'
    ],
    mounted: function() {
        axios.get("/api/config/" + this.service).then(response => {
            this.configObj = response.data
            this.isLoading = false
        })
    },
    methods: {
        save: function() {
            this.isSaving = true
            axios.post("/backend/api/" + this.service, this.configObj)
            .then(() => {
                Snackbar.open({
                    message: 'Successfully saved settings!',
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
        },
        testConnection: function() {
            this.isTesting = true
            axios.post("/api/testConfig/" + this.service, this.configObj)
            .then(() => {
                Snackbar.open({
                    message: 'Successfully connected!',
                    type: 'is-success',
                    indefinite: false,
                    duration: 3000,
                    })
            })
            .catch(() => {
                Snackbar.open({
                    message: 'Test unsuccessful!',
                    type: 'is-danger',
                    indefinite: false,
                    duration: 10000,
                    })
            })
            .finally(() => {
                this.isTesting = false
            })
        }
    }
}
</script>