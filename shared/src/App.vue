<template>
    <div>
        <CSLEditor v-model="code" />
        <!-- <LogPane :logs="logs"/> -->
    </div>
</template>

<script>
import { DateTime } from "luxon";
// import LogPane from "./LogPane.vue"
import CSLEditor from "./CSLEditor.vue"

export default {
    data() {
        return {
            code: `This is some code`,
            logs: [
                {
                    msg: "Hi",
                    variant: "danger",
                    timestamp: DateTime.now()
                },
                {
                    msg: "Hi",
                    variant: "success",
                    timestamp: DateTime.now()
                }
            ]
        }
    },
    components: {
        // LogPane,
        CSLEditor
    }
}
</script>