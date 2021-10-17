# REPL

Use the REPL
<CSLEditor v-model="code"/>
<LogPane :logs="logs" />

<style>
@import 'conductorr-lib/dist/style.css';
</style>

<script>
import { DateTime } from "luxon";
import {CSLEditor, LogPane} from 'conductorr-lib'

export default {
    data() {
        return {
            code: `(if
    this
    then
    that
)`,
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
        CSLEditor: CSLEditor,
        LogPane: LogPane
    }
}
</script>
