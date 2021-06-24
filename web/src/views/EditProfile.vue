<template>
  <div class="flex flex-col h-screen max-h-screen" style="max-width: 100vw">
    <div class="flex flex-col bg-gray-900 bg-opacity-20">
      <div class="p-2 flex flex-1 flex-row justify-between">
        <div class="w-60">
          <o-button @click="$router.go(-1)" icon-left="chevron-left"
            >Back</o-button
          >
        </div>
        <div class="flex text-2xl items-center">
          Editing Profile: {{ profile.name }}
        </div>
        <div class="w-60"></div>
      </div>
      <div class="flex flex-row justify-between">
        <div class="flex flex-col sm:flex-row">
          <div
            class="tabitem"
            @click="activeFunction = 'filter'"
            :class="activeFunction == 'filter' ? 'tabitem-active' : ''"
          >
            Filter
          </div>
          <div
            class="tabitem"
            @click="activeFunction = 'sorter'"
            :class="activeFunction == 'sorter' ? 'tabitem-active' : ''"
          >
            Sorter
          </div>
        </div>
        <!-- <o-checkbox v-model="showGenerator" variant="primary"
            >Show Generator</o-checkbox
          > -->
        <div class="p-2 flex flex-col sm:flex-row">
          <o-button class="mx-1 my-1 sm:my-0" @click="initSplits(true)">Reset View</o-button>
          <o-button class="mx-1 my-1 sm:my-0" variant="primary" @click="validate"
            >Validate</o-button
          >
          <o-button class="mx-1 my-1 sm:my-0" variant="primary">Save</o-button>
        </div>
      </div>
    </div>
    <div class="flex h-full flex-1 w-full" style="max-height: calc(100vh - 104px)">
      <div class="flex h-full w-full flex-1 flex-row">
        <div id="split3" class="flex flex-col">
          <div v-show="showGenerator" id="split1" class="flex flex-col">
            <div class="titlebar">Generator</div>
          </div>
          <div id="split2" class="flex flex-col">
            <div class="titlebar">Editor</div>
            <CSLEditor @submit="run" v-model="computedCode" />
          </div>
        </div>
        <div id="split4" class="flex flex-col">
          <div id="split5" class="flex flex-col">
            <div class="titlebar">Test Cases</div>
            <div class="px-16 h-full overflow-y-scroll overflow-x-scroll w-full">
              Release A
              <release-builder v-model="releaseA" />
              <div class="text-xl">
                Rendered code (this is what you can assume is injected
                immediately before running your script):
              </div>
              <div class="p-4">
                <CSLEditor v-if="activeFunction=='filter'" readonly v-model="releaseACode" />
                <CSLEditor v-if="activeFunction=='sorter'" readonly v-model="releaseABCode" />
              </div>
              <div class="flex flex-row justify-center p-4">
                <o-button @click="run" variant="primary">Run</o-button>
              </div>
            </div>
          </div>
          <div id="split6" class="flex flex-col">
            <div class="titlebar flex flex-row justify-between">
              <div>Output</div>
              <div>
                <o-icon
                  icon="times-circle"
                  class="cursor-pointer mr-3"
                  @click="outputs = []"
                />
                <o-icon
                  icon="angle-double-down"
                  class="cursor-pointer mr-3"
                  @click="scrollOutput"
                />
              </div>
            </div>
            <div ref="outputScroller" class="h-full overflow-y-scroll">
              <div
                v-for="output in outputs"
                :key="output.timestamp"
                :class="outputClass(output.variant)"
                class="p-2 text-lg font-semibold relative"
              >
                <div
                  v-if="output.variant == 'success'"
                  class="absolute mt-1 ml-1"
                >
                  <o-icon icon="check-circle" />
                </div>
                <div
                  v-if="output.variant == 'danger'"
                  class="absolute mt-1 ml-1"
                >
                  <o-icon icon="exclamation-circle" />
                </div>
                <div
                  v-if="output.variant == 'warning'"
                  class="absolute mt-1 ml-1"
                >
                  <o-icon icon="exclamation-triangle" />
                </div>
                <div v-if="output.variant == ''" class="absolute mt-1 ml-1">
                  <o-icon icon="info-circle" />
                </div>
                <div class="mr-3 ml-8 float-left">
                  {{ formatTime(output.timestamp) }}
                </div>
                <div class="text-gray-100">{{ output.msg }}</div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.titlebar {
  @apply text-lg uppercase font-bold p-2 bg-gray-600 w-full border-b-4 border-gray-800;
}

.tabitem {
  @apply h-full;
  @apply bg-gray-700;
  @apply hover:bg-gray-800;
  @apply flex flex-row;
  @apply items-center;
  @apply px-14;
  @apply uppercase;
  @apply text-2xl;
  @apply cursor-pointer;
}

.tabitem-active {
  @apply bg-yellow-500;
  @apply hover:bg-yellow-500;
}

:deep(.gutter) {
  @apply bg-gray-700;
  background-repeat: no-repeat;
  background-position: 50%;
}

:deep(.gutter.gutter-horizontal) {
  background-image: url("data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAUAAAAeCAYAAADkftS9AAAAIklEQVQoU2M4c+bMfxAGAgYYmwGrIIiDjrELjpo5aiZeMwF+yNnOs5KSvgAAAABJRU5ErkJggg==");
  cursor: col-resize;
}

:deep(.gutter.gutter-vertical) {
  background-image: url("data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAB4AAAAFAQMAAABo7865AAAABlBMVEVHcEzMzMzyAv2sAAAAAXRSTlMAQObYZgAAABBJREFUeF5jOAMEEAIEEFwAn3kMwcB6I2AAAAAASUVORK5CYII=");
  cursor: row-resize;
}
</style>

<script>
import APIUtil from "../util/APIUtil";
import CSLEditor from "../components/CSLEditor.vue";
import Split from "split.js";
import "../util/wasm_exec.js";
import { DateTime } from "luxon";
import ReleaseBuilder from "../components/ReleaseBuilder.vue";

export default {
  data() {
    return {
      split12: null,
      split34: null,
      split56: null,
      profile: {},
      profile_id: 0,
      showGenerator: true,
      code: "",
      outputs: [],
      activeFunction: "filter",
      releaseA: {},
      releaseB: {},
    };
  },
  components: { CSLEditor, ReleaseBuilder },
  methods: {
    pushOutput(msg, variant) {
      const output = {
        msg,
        variant,
        timestamp: DateTime.now(),
      };
      this.outputs.push(output);
    },
    loadSizes(name, defaultSizes = [50, 50]) {
      let sizes = localStorage.getItem(name);

      if (sizes) {
        sizes = JSON.parse(sizes);
      } else {
        sizes = defaultSizes;
      }

      return sizes;
    },
    formatTime(timestamp) {
      return timestamp.toLocaleString(DateTime.TIME_WITH_SECONDS);
    },
    initSplits(reset = false) {
      // Initialize the split panels
      if (this.split12) {
        this.split12.destroy();
        this.split12 = null;
      }
      if (this.split34) {
        this.split34.destroy();
        this.split34 = null;
      }
      if (this.split56) {
        this.split56.destroy();
        this.split56 = null;
      }

      if (this.showGenerator) {
        this.split12 = Split(["#split1", "#split2"], {
          sizes: reset ? undefined : this.loadSizes("split-sizes12"),
          minSize: 200,
          direction: "vertical",
          onDragEnd: function (sizes) {
            localStorage.setItem("split-sizes12", JSON.stringify(sizes));
          },
        });
      }

      this.split34 = Split(["#split3", "#split4"], {
        sizes: reset ? [40, 60] : this.loadSizes("split-sizes34", [40, 60]),
        minSize: 200,
        onDragEnd: function (sizes) {
          localStorage.setItem("split-sizes34", JSON.stringify(sizes));
        },
      });

      this.split56 = Split(["#split5", "#split6"], {
        direction: "vertical",
        minSize: 200,
        sizes: reset ? [70, 30] : this.loadSizes("split-sizes56", [70, 30]),
        onDragEnd: function (sizes) {
          localStorage.setItem("split-sizes56", JSON.stringify(sizes));
        },
      });
    },
    validate() {
      Validate(this.computedCode, (ok, err) => {
        if (!ok) {
          this.pushOutput("Validation error: " + err, "danger");
        } else {
          this.pushOutput("Validation succeeded", "success");
        }
      });
    },
    run() {
      Run(this.computedCode, (ok, err, result) => {
        if (!ok) {
          this.pushOutput("Execution error: " + err, "danger");
        } else if (result) {
          console.log(result)
          this.pushOutput(result, "success");
        } else {
          this.pushOutput("Script returned null", "warning")
        }
      });
    },
    outputClass(variant) {
      if (variant == "success") {
        return `bg-green-600 text-green-300`;
      } else if (variant == "danger") {
        return `bg-red-600 text-red-300`;
      } else if (variant == "warning") {
        return `bg-yellow-600 text-yellow-300`;
      } else {
        return `bg-gray-600 text-gray-300`;
      }
    },
    scrollOutput() {
      this.$refs.outputScroller.scrollTop =
        this.$refs.outputScroller.scrollHeight;
    },
    renderedCode(release) {
      return `(define a ("${release.title}" "${release.indexer}" "${release.download_type}" "${release.content_type}" "${release.rip_type}" "${release.quality}" "${release.resolution}" ${release.encoding} ${release.seeders} ${release.age} ${release.bitrate}))`;
    },
  },
  created() {
    this.profile_id = this.$route.params.profile_id;
  },
  mounted() {
    APIUtil.getProfile(this.profile_id).then((profile) => {
      this.profile = profile;
    });
    let savedPref = localStorage.getItem("show-generator");
    this.showGenerator = savedPref != "";
    this.initSplits();

    let go = new Go();

    WebAssembly.instantiateStreaming(
      fetch("/api/csl.wasm"),
      go.importObject
    ).then((result) => {
      go.run(result.instance);
    });
  },
  watch: {
    showGenerator: function (newVal) {
      this.initSplits();
      if (newVal) {
        localStorage.setItem("show-generator", "true");
      } else {
        localStorage.setItem("show-generator", "");
      }
    },
    outputs: {
      handler() {
        setTimeout(() => {
          this.scrollOutput();
        }, 10);
      },
      deep: true,
    },
  },
  computed: {
    releaseACode() {
      return this.renderedCode(this.releaseA)
    },
    releaseBCode() {
      return this.renderedCode(this.releaseB)
    },
    releaseABCode() {
      return this.releaseACode + '\n' + this.releaseBCode
    },
    computedCode: {
      get() {
        if (this.activeFunction == 'filter') {
          return this.profile.filter
        } else if (this.activeFunction == 'sorter') {
          return this.profile.sorter
        }
      },
      set(val) {
        if (this.activeFunction == 'filter') {
          this.profile.filter = val
        } else if (this.activeFunction == 'sorter') {
          this.profile.sorter = val
        }
      }
    }
  }
};
</script>
