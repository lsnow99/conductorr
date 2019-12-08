<template>
    <b-menu class="menu">
        <b-menu-list label="Main Configuration">
            <b-menu-item :active="isActive" icon="server" label="System" v-on:click="navigate('system')"></b-menu-item>
            <b-menu-item icon="lock" label="Security" v-on:click="navigate('security')"></b-menu-item>
            <!-- <b-menu-item
                label="Support"
                icon="link"
                tag="router-link"
                target="_blank"
                to="/expo">
            </b-menu-item> -->
        </b-menu-list>
        <b-menu-list label="Connections">
            <b-menu-item icon="tv" label="Sonarr" v-on:click="navigate('sonarr')"></b-menu-item>
            <b-menu-item icon="film" label="Radarr" v-on:click="navigate('radarr')"></b-menu-item>
            <b-menu-item icon="robot" label="Filebot" v-on:click="navigate('filebot')"></b-menu-item>
            <b-menu-item icon="play" label="Plex" v-on:click="navigate('plex')"></b-menu-item>
        </b-menu-list>
    </b-menu>
</template>

<script>
export default {
    name: 'ConfigurationMenu',
    data() {
        return {
            isActive: true
        }
    },
    methods: {
        navigate: function(navTo) {
            this.$emit("selectionChanged", navTo)
        }
    }
}
</script>

<style>
.icon.is-small {
    vertical-align: middle !important
}
</style>

<style scoped>
@media screen and (min-width: 775px) {
    .menu {
        position: fixed;
    }
}
</style>