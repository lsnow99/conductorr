<template>
  <div class="flex flex-col h-screen max-h-screen" style="max-width: 100vw">
    <div class="flex flex-col bg-gray-900 bg-opacity-20">
      <div class="p-2 flex flex-1 flex-row justify-between">
        <div class="w-32">
          <o-button @click="$router.go(-1)" icon-left="chevron-left"
            >Back</o-button
          >
        </div>
        <div class="flex text-2xl items-center">
          <span v-if="!editingName"
            >{{ `Editing Profile: ${profile.name}` }}
            <o-icon
              @click="editingName = true"
              class="cursor-pointer ml-4"
              icon="edit" /></span
          ><span v-else class="flex flex-row"
            ><o-input
              @keydown.enter="editingName = false"
              v-model="profile.name"
            /><o-icon
              @click="editingName = false"
              class="cursor-pointer ml-4"
              icon="check"
            />
          </span>
        </div>
        <div class="w-32 hidden sm:flex"></div>
      </div>
      <div class="flex flex-col sm:flex-row justify-between">
        <div class="flex flex-row w-full">
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
        <div class="p-2 flex flex-row justify-center">
          <o-button class="mx-1 my-1 sm:my-0" @click="initSplits(true)"
            >Reset View</o-button
          >
          <o-button
            class="mx-1 my-1 sm:my-0"
            variant="primary"
            @click="updateProfile"
            >Save</o-button
          >
        </div>
      </div>
    </div>
    <div
      class="flex h-full flex-1 w-full"
      style="max-height: calc(100vh - 104px)"
    >
      <div class="flex h-full w-full flex-1 flex-row">
        <div id="split3" class="flex flex-col">
          <div id="split1" class="flex flex-col">
            <div class="titlebar">Generator</div>
          </div>
          <div id="split2" class="flex flex-col">
            <div class="titlebar flex flex-row justify-between">
              <div>Editor</div>
              <div>
                <o-button
                  size="small"
                  class="mx-1"
                  @click="validate"
                  >Validate</o-button
                >
                <o-button
                  variant="primary"
                  size="small"
                  class="mx-1"
                  icon-right="play"
                  @click="run"
                  >Run</o-button
                >
              </div>
            </div>
            <CSLEditor @submit="run" v-model="computedCode" />
          </div>
        </div>
        <div id="split4" class="flex flex-col">
          <div id="split5" class="flex flex-col">
            <div class="titlebar">Test Cases</div>
            <div
              class="px-16 h-full overflow-y-scroll overflow-x-scroll w-full w-94"
            >
              <release-builder title="Release A" v-model="releaseA" />
              <release-builder
                v-show="activeFunction == 'sorter'"
                title="Release B"
                v-model="releaseB"
              />
              <div class="text-xl">
                Rendered code (this is what you can assume is injected
                immediately before running your script):
              </div>
              <div class="p-4">
                <CSLEditor
                  v-if="activeFunction == 'filter'"
                  readonly
                  v-model="releaseACode"
                />
                <CSLEditor
                  v-if="activeFunction == 'sorter'"
                  readonly
                  v-model="releaseABCode"
                />
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
              <LogPane :logs="outputs" />
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
  @apply w-full;
  @apply sm:w-auto;
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
import LogPane from "../components/LogPane.vue";

export default {
  data() {
    return {
      split12: null,
      split34: null,
      split56: null,
      profile: {},
      profileID: 0,
      code: "",
      outputs: [],
      activeFunction: "filter",
      releaseA: {},
      releaseB: {},
      editingName: false,
    };
  },
  components: { CSLEditor, ReleaseBuilder, LogPane },
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

      this.split12 = Split(["#split1", "#split2"], {
        sizes: reset ? undefined : this.loadSizes("split-sizes12"),
        minSize: 200,
        direction: "vertical",
        onDragEnd: function (sizes) {
          localStorage.setItem("split-sizes12", JSON.stringify(sizes));
        },
      });

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

      if (reset) {
        localStorage.setItem("split-sizes12", "");
        localStorage.setItem("split-sizes34", "");
        localStorage.setItem("split-sizes56", "");
      }
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
      let env = {
        a: this.releaseA,
      };
      if (this.activeFunction == "sorter") env.b = this.releaseB;
      Run(this.computedCode, env, (ok, err, result) => {
        if (!ok) {
          this.pushOutput("Execution error: " + err, "danger");
        } else if (result || result === 0 || result === false) {
          console.log(result);
          this.pushOutput(result, "success");
        } else {
          this.pushOutput("Script returned null", "warning");
        }
      });
    },
    scrollOutput() {
      this.$refs.outputScroller.scrollTop =
        this.$refs.outputScroller.scrollHeight;
    },
    renderedCode(release, releaseVar) {
      return `(define ${releaseVar} ("${release.title}" "${release.indexer}" "${release.download_type}" "${release.content_type}" "${release.rip_type}" "${release.resolution}" "${release.encoding}" ${release.seeders} ${release.age} ${release.size} ${release.runtime}))`;
    },
    editName() {
      this.editingName = true;
    },
    updateProfile() {
      APIUtil.updateProfile(
        this.profileID,
        this.profile.name,
        this.profile.filter,
        this.profile.sorter
      ).then(() => {
        this.$oruga.notification.open({
          duration: 3000,
          message: `Saved successfully`,
          position: "bottom-right",
          variant: "success",
          closable: false,
        });
      });
    },
    initCSL() {
      let go = new Go();

      WebAssembly.instantiateStreaming(
        fetch("/api/csl.wasm"),
        go.importObject
      ).then(async (result) => {
        await go.run(result.instance);
        this.initCSL();
      });
    },
  },
  created() {
    this.profileID = parseInt(this.$route.params.profile_id);
  },
  mounted() {
    APIUtil.getProfile(this.profileID).then((profile) => {
      this.profile = profile;
    });
    this.initSplits();
    this.initCSL();
  },
  watch: {
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
      return this.renderedCode(this.releaseA, "a");
    },
    releaseBCode() {
      return this.renderedCode(this.releaseB, "b");
    },
    releaseABCode() {
      return this.releaseACode + "\n" + this.releaseBCode;
    },
    computedCode: {
      get() {
        if (this.activeFunction == "filter") {
          return this.profile.filter;
        } else if (this.activeFunction == "sorter") {
          return this.profile.sorter;
        }
      },
      set(val) {
        if (this.activeFunction == "filter") {
          this.profile.filter = val;
        } else if (this.activeFunction == "sorter") {
          this.profile.sorter = val;
        }
      },
    },
  },
};
</script>
