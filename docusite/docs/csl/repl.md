# REPL

Enter a CSL expression and press enter to evaluate it

<div class="h-64 overflow-y-scroll" ref="logWrapper">
<LogPane :logs="logs" />
</div>
<input 
    type="text" 
    placeholder="Type something here" 
    v-model="replInput" 
    class="bg-gray-700 w-full p-1 text-white text-lg border-0 outline-none" 
    @keydown.enter="execute"
    @keydown.up="upHistory"
    @keydown.down="downHistory" />

<style>
@import 'conductorr-lib/dist/style.css';
</style>

<script>
import { DateTime } from "luxon";
import { LogPane } from 'conductorr-lib'
// import "/src/util/wasm_exec.js";

export default {
    data() {
        return {
            replInput: '',
            logs: [],
            history: [],
            historyIndex: 0
        }
    },
    components: {
        LogPane
    },
    methods: {
        initCSL() {
            let go = new Go();

            WebAssembly.instantiateStreaming(
                fetch("/csl.wasm"),
                go.importObject
            ).then(async (result) => {
                await go.run(result.instance);
                this.initCSL();
            });
        },
        upHistory() {
            if(this.historyIndex > 0) {
                this.historyIndex --;
            }
            this.replInput = this.history[this.historyIndex];
        },
        downHistory() {
            if(this.historyIndex <= this.history.length - 1) {
                this.historyIndex ++;
            } 
            if (this.historyIndex == this.history.length) {
                this.replInput = ''
            } else {
                this.replInput = this.history[this.historyIndex];
            }
        },
        execute() {
            if(this.history[this.history.length - 1] != this.replInput) {
                this.history.push(this.replInput);
            }
            this.historyIndex = this.history.length;
            this.pushOutput(this.replInput, "default", "italic")
            Execute(this.replInput, (ok, err, result) => {
                this.replInput = '';
                if (!ok) {
                    this.pushOutput("Execution error: " + err, "danger", "bold");
                } else if (result || result === 0 || result === false) {
                    console.log(result);
                    this.pushOutput(result, "success", "bold");
                } else {
                    this.pushOutput("null", "warning", "bold");
                }
            });
        },
        pushOutput(msg, variant, decoration) {
            const output = {
                msg,
                variant,
                decoration,
                timestamp: DateTime.now(),
            };
            this.logs.push(output);
        },
    },
    mounted() {
        this.initCSL();
    },
    watch: {
        logs: {
            handler: function(){
                this.$nextTick(() => {
                    this.$refs.logWrapper.scrollTop = this.$refs.logWrapper.scrollHeight;
                })
            },
            deep: true
        }
    }
}
</script>
