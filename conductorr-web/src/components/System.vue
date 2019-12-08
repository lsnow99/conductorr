<template>
    <div class="container">
        <div class="setting-header">
            <img
                :src="require('@/assets/conductorr.svg')"
                class="setting-logo"
            />
            System
        </div>
        <div class="tile is-parent is-vertical">
            <div class="tile is-child">
                <span class="header-title">
                    <b-icon
                        style="vertical-align: middle"
                        icon="code-branch"
                        size="is-medium"
                    >
                    </b-icon>
                    Version
                </span>
            </div>
            <div class="tile is-child subsection">
                <div class="subsection-title">
                    Running version v1.0
                </div>
                <button @click="checkUpdates" class="button is-primary">
                    <b-icon
                        pack="fas"
                        icon="sync-alt"
                        :custom-class="checkingUpdates ? 'fa-spin' : ''"
                    >
                    </b-icon>
                    <span>Check for Updates</span>
                </button>
            </div>
            <div class="tile is-child">
                <span class="header-title">
                    <b-icon
                        style="vertical-align: middle"
                        icon="heartbeat"
                        size="is-medium"
                        type="is-health"
                    >
                    </b-icon>
                    System Health
                </span>
            </div>
            <div class="tile is-child subsection">
                <span class="subsection-title">
                    Overall Health
                    <b-icon
                        style="vertical-align: middle"
                        :icon="allSuccess() ? 'check' : 'times'"
                        size="is-medium"
                        :type="allSuccess() ? 'is-success' : 'is-danger'"
                    >
                    </b-icon>
                </span>
                <ul style="padding-left: 7px">
                    <li v-for="(value, name) in health" v-bind:key="name">
                        <span class="subsection-subtitles">
                            {{ name.charAt(0).toUpperCase() + name.substr(1) }}
                            <b-icon
                                style="vertical-align: middle"
                                :icon="value ? 'check' : 'times'"
                                size="is-medium"
                                :type="value ? 'is-success' : 'is-danger'"
                            >
                            </b-icon>
                        </span>
                    </li>
                </ul>
            </div>
            <div class="tile is-child">
                <span class="header-title">
                    <b-icon
                        style="vertical-align: middle"
                        icon="wrench"
                        size="is-medium"
                    >
                    </b-icon>
                    Configuration
                </span>
            </div>
            <div class="tile is-child subsection">
                <setting :setting="filebot_logs"></setting>
                <setting :setting="plex_scanner_logs"></setting>
            </div>
        </div>
        <div class="btn-container">
            <b-button
                size="is-medium"
                @click="save"
                :loading="isSaving"
                type="is-success"
                >Save</b-button
            >
        </div>
    </div>
</template>

<script>
import { SnackbarProgrammatic as Snackbar } from "buefy";
import axios from "axios";
import Setting from "@/components/Setting.vue";

export default {
    name: "system",
    components: {
        Setting
    },
    data() {
        return {
            checkingUpdates: false,
            isSaving: false,
            health: {
                sonarr: true,
                radarr: false,
                filebot: true,
                plex: true
            },
            filebot_logs: {
                datatype: "bool",
                label: "Record Filebot Logs",
                placeholder: "",
                bool_value: true,
                description: "Save filebot logs in database"
            },
            plex_scanner_logs: {
                datatype: "bool",
                label: "Record Plex Scanner Logs",
                placeholder: "",
                bool_value: true,
                description: "Save Plex Scanner logs in database"
            }
        };
    },
    methods: {
        checkUpdates() {
            /* eslint-disable */
            this.checkingUpdates = true;
            axios
                .post("/backend/checkUpdates/")
                .then(() => {
                    Snackbar.open({
                        message: "New version available",
                        type: "is-success",
                        indefinite: false,
                        duration: 3000
                    });
                })
                .catch(() => {
                    Snackbar.open({
                        message: "Error checking for new versions",
                        type: "is-danger",
                        indefinite: false,
                        duration: 10000
                    });
                })
                .finally(() => {
                    this.checkingUpdates = false;
                });
        },
        allSuccess() {
            let success = true;
            for (let key of Object.keys(this.health)) {
                success = success & this.health[key];
            }
            return success;
        },
        save() {
            this.isSaving = true;
            axios
                .post("/backend/config/system", {})
                .then(() => {
                    Snackbar.open({
                        message: "Succcessfully saved settings!",
                        type: "is-success",
                        indefinite: false,
                        duration: 3000
                    });
                })
                .catch(() => {
                    Snackbar.open({
                        message: "Error saving settings",
                        type: "is-danger",
                        indefinite: false,
                        duration: 10000
                    });
                })
                .finally(() => {
                    this.isSaving = false;
                });
        }
    }
};
</script>

<style scoped></style>
