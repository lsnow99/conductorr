# REPL

Enter a CSL expression and press enter to evaluate it

<div class="relative">
    <div v-if="loadingMsg" class="absolute top-0 bottom-0 left-0 right-0 flex items-center justify-center text-2xl text-white bg-gray-700 bg-opacity-75">
        {{ loadingMsg }}
    </div>
    <div style="height: 16em; overflow-y: auto" ref="logWrapper">
        <LogPane :logs="logs" />
    </div>
    <input 
        type="text" 
        placeholder="Type something here" 
        v-model="replInput" 
        class="prompt" 
        @keydown.enter="execute"
        @keydown.up="upHistory"
        @keydown.down="downHistory" />
</div>

<style>
@import 'conductorr-lib/dist/style.css';

.prompt {
    @apply w-full p-1 text-lg text-white border-4 border-white !important;
    border-bottom: 4px solid #bbb !important;
}
.prompt::placeholder {
    @apply text-white;
}
</style>

<script>
import { DateTime } from "luxon";
import { LogPane } from 'conductorr-lib'

export default {
    data() {
        return {
            replInput: '',
            logs: [],
            history: [],
            historyIndex: 0,
            loadingMsg: "Loading CSL WebAssembly module...",
            lastRetryTime: new Date().getTime(),
        }
    },
    components: {
        LogPane
    },
    methods: {
        initCSL() {
            let go = new Go();

            if (typeof WebAssembly === "object"
             && typeof WebAssembly.instantiateStreaming === "function") {
                WebAssembly.instantiateStreaming(
                    fetch("/csl.wasm"),
                    go.importObject
                ).then(async (result) => {
                    this.loadingMsg = ""
                    await go.run(result.instance);
                    const now = new Date().getTime()
                    if (now - this.lastRetryTime > 2000) {
                        this.initCSL();
                    } else {
                        this.loadingMsg = "Time since last WebAssembly initialization attempt is less than the minimum of 20 seconds. Refresh the page to reattempt initialization"
                    }
                }).catch(() => {
                    this.loadingMsg = "Error downloading WebAssembly module"
                })
            } else {
                this.loadingMsg = "WebAssembly is not supported by your browser"
            }
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
        import('../public/wasm_exec').then(() => {
            this.initCSL();
        })
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
