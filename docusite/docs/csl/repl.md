# REPL

<CSLEditor v-model="code"/>
<LogPane :logs="logs" />
<DownloadStatus :download="{status: 'waiting', friendly_name: 'Download', fraction: 0.5}" />

<style scoped>
@import 'conductorr-lib/dist/style.css'
</style>

<script>
import { DateTime } from "luxon";
import C from 'conductorr-lib'
import "vue-prism-editor/dist/prismeditor.min.css";

export default {
    data() {
        return {
            code: `this
            is
            code`,
            logs: [
                {
                    msg: "Hi",
                    variant: "danger",
                    timestamp: DateTime.now()
                }
            ]
        }
    },
    components: {
        CSLEditor: C.CSLEditor,
        LogPane: C.LogPane,
        DownloadStatus: C.DownloadStatus
    }
}
</script>
