<template>
    <div class="container">
            <b-loading :active.sync="isLoading" :is-full-page="false" :can-cancel="false"></b-loading>
        <div class="setting-header"><img src="@/assets/filebot.logo.svg" class="setting-logo"> Filebot</div>
        <div style="margin-top: 20px;" >
            <span style="font-size:20px;">Required Fields</span>
            <setting v-for="x in configObj" v-bind:key="x.property" v-bind:setting="x"></setting>  
        </div>
        <b-button style="float:right" size="is-medium" @click="save" :loading="isSaving" type="is-success">Save</b-button>
    </div>
</template>

<script>
import axios from 'axios'
import Setting from '@/components/Setting.vue'
import { SnackbarProgrammatic as Snackbar } from 'buefy'

export default {
    name: 'filebot',
    components: {
        Setting
    },
    data() {
        return {
            isLoading: true,
            isSaving: false,
            name: 'alan',
            configObj: []
        }
    },
    mounted: function() {
        axios.get("/backend/config/filebot").then(response => {
            this.configObj = response.data
            this.isLoading = false
        })
    },
    methods: {
        save: function() {
            this.isSaving = true
            axios.post("/backend/config/filebot", this.configObj)
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
                    message: 'Error saving settings!',
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