# REPL

<CSLEditor v-model="code"/>
<LogPane :logs="logs" />

<style scoped>
@import 'shared/dist/style.css'
</style>

<script>
import { DateTime } from "luxon";
import C from 'shared'
import "vue-prism-editor/dist/prismeditor.min.css";

console.log(C);

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
        LogPane: C.LogPane
    }
}
</script>
