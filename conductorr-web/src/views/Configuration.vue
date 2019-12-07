<template>
<div class="container">
    <div class="tile is-ancestor">
        <div class="tile is-3">
            <configuration-menu @selectionChanged="switchSettingView"></configuration-menu>
        </div>
            <config-wrapper class="tile">
                <transition name="fade" mode="out-in">
                    <component :is="selection" transition="fade" transition-mode="out-in"></component>
                </transition>
            </config-wrapper>
    </div>
</div>
</template>

<script>
import ConfigurationMenu from '@/components/ConfigurationMenu.vue'
import ConfigWrapper from '@/components/ConfigWrapper.vue'
import System from '@/components/System.vue'
import Plex from '@/components/Plex.vue'
import Sonarr from '@/components/Sonarr.vue'
import Radarr from '@/components/Radarr.vue'
import Filebot from '@/components/Filebot.vue'
import Security from '@/components/Security.vue'

export default {
    name: 'configuration',
    components: {
        ConfigurationMenu,
        ConfigWrapper,
        System,
        Plex,
        Sonarr,
        Radarr,
        Filebot,
        Security
    },
    data() {
        return {
            selection: 'system'
        }
    },
    methods: {
        switchSettingView(selection) {
            this.selection = selection
        }
    }
}
</script>

<style scoped>
.fade-transition {
   transition: opacity 0.2s ease;
 }

.fade-enter, .fade-leave {
  opacity: 0;
}
</style>