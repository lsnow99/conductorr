<template>
    <div class="container">
        <b-loading
            style="z-index: 1000"
            :active.sync="isLoading"
            :is-full-page="false"
            :can-cancel="false"
        ></b-loading>
        <b-tabs>
                        <b-tab-item label="Timeline">
                <b-steps
                    v-model="activeStep"
                    type="is-info"
                    :has-navigation="false"
                >
                    <b-step-item label="Grabbed" icon="search"></b-step-item>
                    <b-step-item
                        label="Downloading"
                        icon="download"
                    ></b-step-item>
                    <b-step-item label="Filebot" icon="robot"></b-step-item>
                    <b-step-item label="Plex Scan" icon="play"></b-step-item>
                    <b-step-item label="Done" icon="check"></b-step-item>
                    <div class="tile is-ancestor">
                        <div class="tile" style="justify-content: center">
                            <vue-moments-ago
                                prefix=""
                                suffix="ago"
                                date="2018-05-02T20:22:22.285Z"
                                lang="en"
                            ></vue-moments-ago>
                        </div>
                        <div class="tile" style="justify-content: center">
                            <vue-moments-ago
                                prefix=""
                                suffix="ago"
                                date="2018-05-02T20:22:22.285Z"
                                lang="en"
                            ></vue-moments-ago>
                        </div>
                        <div class="tile" style="justify-content: center">
                            <vue-moments-ago
                                prefix=""
                                suffix="ago"
                                date="2018-05-02T20:22:22.285Z"
                                lang="en"
                            ></vue-moments-ago>
                        </div>
                        <div class="tile" style="justify-content: center">
                            <vue-moments-ago
                                prefix=""
                                suffix="ago"
                                date="2018-05-02T20:22:22.285Z"
                                lang="en"
                            ></vue-moments-ago>
                        </div>
                        <div class="tile" style="justify-content: center">
                            <vue-moments-ago
                                prefix=""
                                suffix="ago"
                                date="2018-05-02T20:22:22.285Z"
                                lang="en"
                            ></vue-moments-ago>
                        </div>
                    </div>
                </b-steps>
            </b-tab-item>
            
            <b-tab-item label="Filebot Logs">
                <code-viewer
                    codeText="Filebot exited normally with output: Locking /var/lib/plexmediaserver/.filebot/logs/amc.log\nRun script [fn:amc] at [Sat Dec 07 17:09:13 EST 2019]\nParameter: unsorted = y\nParameter: music = n\nParameter: subtitles = en\nParameter: artwork = y\nParameter: excludeList = .excludes\nArgument[0]: /valinor/downloads/nzbget/completed/Series/The.Kitchen.S23E04.Home.Cooked.for.the.Holidays.1080p.WEB.x264.1-CAFFEiNE-Obfuscated\nArgument[1]: /valinor/downloads/nzbget/completed/Series/The.Kitchen.S23E04.Home.Cooked.for.the.Holidays.1080p.WEB.x264.1-CAFFEiNE-Obfuscated\nUse excludes: /valinor/plex/.excludes (3481)\nInput: /valinor/downloads/nzbget/completed/Series/The.Kitchen.S23E04.Home.Cooked.for.the.Holidays.1080p.WEB.x264.1-CAFFEiNE-Obfuscated/9feb118fb26742018930ff21ca6eda61.mkv\nInput: /valinor/downloads/nzbget/completed/Series/The.Kitchen.S23E04.Home.Cooked.for.the.Holidays.1080p.WEB.x264.1-CAFFEiNE-Obfuscated/9feb118fb26742018930ff21ca6eda61.mkv\nGroup: [Series:the kitchen] =\u003e [9feb118fb26742018930ff21ca6eda61.mkv]\nGet [English] subtitles for 1 files\nLooking up subtitles by hash via OpenSubtitles\nRename episodes using [TheTVDB]\nAuto-detected query: [The Kitchen]\nFetching episode data for [The Kitchen]\nFetching episode data for [The Kitchen ]\nFetching episode data for [The Kitchen (UK)]\nFetching episode data for [The Trial Kitchen]\nFetching episode data for [The Chef's Kitchen]\n[COPY] from [/valinor/downloads/nzbget/completed/Series/The.Kitchen.S23E04.Home.Cooked.for.the.Holidays.1080p.WEB.x264.1-CAFFEiNE-Obfuscated/9feb118fb26742018930ff21ca6eda61.mkv] to [/valinor/plex/TV Shows/The Kitchen/Season 23/The Kitchen - S23E04 - Home Cooked for the Holidays.mkv]\nProcessed 1 files\nFetching series artwork for [The Kitchen / Season 23] to [/valinor/plex/TV Shows/The Kitchen/Season 23]\nGenerate Series NFO: null [278664]\nBanner already exists: /valinor/plex/TV Shows/The Kitchen/poster.jpg\nBanner already exists: /valinor/plex/TV Shows/The Kitchen/banner.jpg\nBanner already exists: /valinor/plex/TV Shows/The Kitchen/fanart.jpg\nBanner not found: /valinor/plex/TV Shows/The Kitchen/Season 23/poster.jpg / season:season\nBanner not found: /valinor/plex/TV Shows/The Kitchen/Season 23/banner.jpg / seasonwide:seasonwide\nFanart not found: /valinor/plex/TV Shows/The Kitchen/clearart.png / hdclearart\nFanart not found: /valinor/plex/TV Shows/The Kitchen/clearart.png / clearart\nFanart not found: /valinor/plex/TV Shows/The Kitchen/logo.png / hdtvlogo\nFanart not found: /valinor/plex/TV Shows/The Kitchen/logo.png / clearlogo\nFanart already exists: /valinor/plex/TV Shows/The Kitchen/landscape.jpg\nFanart not found: /valinor/plex/TV Shows/The Kitchen/Season 23/landscape.jpg / seasonthumb\nDone ヾ(＠⌒ー⌒＠)ノ\n"
                    >gdf</code-viewer
                >
            </b-tab-item>

            <b-tab-item label="Plex Scanner Logs">
                What light is light, if Silvia be not seen? <br />
                What joy is joy, if Silvia be not by— <br />
                Unless it be to think that she is by <br />
                And feed upon the shadow of perfection? <br />
                Except I be by Silvia in the night, <br />
                There is no music in the nightingale.
            </b-tab-item>

            <b-tab-item label="Metadata">
                What light is light, if Silvia be not seen? <br />
                What joy is joy, if Silvia be not by— <br />
                Unless it be to think that she is by <br />
                And feed upon the shadow of perfection? <br />
                Except I be by Silvia in the night, <br />
                There is no music in the nightingale.
            </b-tab-item>
        </b-tabs>
    </div>
</template>

<script>
import VueMomentsAgo from "vue-moments-ago";
import CodeViewer from "@/components/CodeViewer.vue";

export default {
    name: "extended-data",
    components: {
        CodeViewer,
        VueMomentsAgo
    },
    props: ["jobId"],
    mounted: function() {
        // TODO: do
        this.isLoading = false;
    },
    data() {
        return {
            isLoading: true,
            activeStep: 2
        };
    }
};
</script>
