<template>
  <div class="w-full h-screen max-h-screen overflow-hidden">
    <div class="flex flex-col bg-gray-900 bg-opacity-20">
      <div
        ref="header"
        class="flex flex-col items-center justify-between flex-1 p-2 lg:flex-row"
      >
        <div class="flex flex-row">
          <div class="p-2">
            <o-button @click="$router.go(-1)" icon-left="chevron-left"
              >Back</o-button
            >
          </div>
          <o-tabs type="boxed" ref="tabs" v-model="currentTab" class="p-2">
            <o-tab-item>
              <template v-slot:header>
                <span class="flex flex-row items-center text-2xl">
                  <vue-fontawesome icon="filter"></vue-fontawesome
                  ><span style="white-space: nowrap" class="flex ml-2"
                    >Filter</span
                  >
                </span>
              </template>
              <div></div>
            </o-tab-item>
            <o-tab-item>
              <template v-slot:header>
                <span class="flex flex-row items-center text-2xl">
                  <vue-fontawesome icon="sort"></vue-fontawesome
                  ><span style="white-space: nowrap" class="flex ml-2"
                    >Sorter</span
                  >
                </span>
              </template>
              <div></div>
            </o-tab-item>
          </o-tabs>
        </div>
        <div class="flex items-center text-2xl">
          <span v-if="!editingName"
            >{{ `Editing Profile: ${profile.name}` }}
            <o-icon
              @click="startEditingName"
              @keydown.enter="startEditingName"
              @keydown.space="startEditingName"
              tabindex="0"
              role="button"
              aria-label="Edit name"
              class="ml-4 cursor-pointer"
              icon="edit" /></span
          ><span v-else class="flex flex-row">
            <o-input
              @keydown.enter="editingName = false"
              v-model="profile.name"
            />
            <o-icon
              @click="editingName = false"
              @keydown.enter="editingName = false"
              @keydown.space="editingName = false"
              tabindex="0"
              role="button"
              aria-label="Edit name"
              class="ml-4 cursor-pointer"
              icon="check"
            />
            <o-icon
              @click="resetEditingName"
              @keydown.enter="resetEditingName"
              @keydown.space="resetEditingName"
              tabindex="0"
              role="button"
              aria-label="Edit name"
              class="ml-4 cursor-pointer"
              icon="times"
            />
          </span>
        </div>

        <div class="flex flex-col justify-between sm:flex-row">
          <div class="flex flex-row justify-center p-2">
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
    </div>
    <div
      id="splitWrapper"
      class="flex flex-row"
      :style="`height: calc(100% - ${headerHeight}px);`"
    >
      <div id="split3" class="flex flex-col">
        <div id="split1" class="flex flex-col">
          <div class="titlebar">Generator</div>
        </div>
        <div id="split2" class="flex flex-col">
          <div class="flex flex-row justify-between titlebar">
            <div>Editor</div>
            <div>
              <o-button size="small" class="mx-1" @click="validate"
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
          <div class="overflow-x-auto">
            <div
              class="h-full px-2 lg:px-16 overflow-y-scroll min-w-[500px]"
              style="width: 100%"
            >
              <release-builder title="Release A" v-model="releaseA" />
              <release-builder
                v-show="activeFunction == 'sorter'"
                title="Release B"
                v-model="releaseB"
              />
              <div class="text-xl">
                Rendered code (you can assume that your script will be run like this):
              </div>
              <div class="p-4">
                <CSLEditor readonly v-model="renderedCode" />
              </div>
              <div class="flex flex-row justify-center p-4">
                <o-button @click="run" variant="primary">Run</o-button>
              </div>
            </div>
          </div>
        </div>
        <div id="split6" class="flex flex-col">
          <div class="flex flex-row justify-between titlebar">
            <div>Output</div>
            <div>
              <o-icon
                icon="times-circle"
                class="mr-3 cursor-pointer"
                @click="outputs = []"
                @keydown.enter="outputs = []"
                @keydown.space="outputs = []"
                tabindex="0"
                role="button"
                aria-label="Clear output"
              />
              <o-icon
                icon="angle-double-down"
                class="mr-3 cursor-pointer"
                @click="scrollOutput"
                @keydown.enter="scrollOutput"
                @keydown.space="scrollOutput"
                tabindex="0"
                role="button"
                aria-label="Scroll output to bottom"
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
import { nextTick } from "vue";
import APIUtil from "../util/APIUtil";
import { CSLEditor, LogPane } from "conductorr-lib";
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
      uneditedProfile: {},
      profileID: 0,
      code: "",
      outputs: [],
      activeFunction: "filter",
      releaseA: {},
      releaseB: {},
      editingName: false,
      oldName: "",
      currentTab: 1,
      headerHeight: 0,
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
    renderedReleaseCode(release, indentAmount) {
      let indent = "";
      for (let i = 0; i < indentAmount; i++) {
        indent += "  ";
      }
      let code = `{CSL_INDENT}(
{CSL_INDENT}  "${release.title}" 
{CSL_INDENT}  "${release.indexer}" 
{CSL_INDENT}  "${release.download_type}" 
{CSL_INDENT}  "${release.content_type}" 
{CSL_INDENT}  "${release.rip_type}" 
{CSL_INDENT}  "${release.resolution}" 
{CSL_INDENT}  "${release.encoding}" 
{CSL_INDENT}  ${release.seeders} 
{CSL_INDENT}  ${release.age} 
{CSL_INDENT}  ${release.size} 
{CSL_INDENT}  ${release.runtime}
{CSL_INDENT})`;
      return code.replaceAll("{CSL_INDENT}", indent);
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
        this.setUneditedProfile(
          this.profile.name,
          this.profile.filter,
          this.profile.sorter
        );
      });
    },
    initCSL() {
      let go = new Go();

      WebAssembly.instantiateStreaming(fetch("/api/csl.wasm"), go.importObject)
        .then(async (result) => {
          await go.run(result.instance);
          this.initCSL();
        })
        .catch((err) => {
          console.log(err);
        });
    },
    loadProfile() {
      APIUtil.getProfile(this.profileID).then((profile) => {
        this.profile = profile;
        this.setUneditedProfile(profile.name, profile.filter, profile.sorter);
      });
    },
    setUneditedProfile(name, filter, sorter) {
      this.uneditedProfile.name = name;
      this.uneditedProfile.filter = filter;
      this.uneditedProfile.sorter = sorter;
    },
    beforeWindowUnload(e) {
      if (this.isProfileModified) {
        e.preventDefault();
        return "Profile has been modified, are you sure you want to leave without saving?";
      }
      return null;
    },
    onResize() {
      this.headerHeight = this.$refs.header.clientHeight;
    },
    startEditingName() {
      this.oldName = this.profile.name;
      this.editingName = true;
    },
    resetEditingName() {
      this.profile.name = this.oldName;
      this.editingName = false;
    },
  },
  created() {
    this.profileID = parseInt(this.$route.params.profile_id);
  },
  beforeRouteLeave(to, from, next) {
    if (this.isProfileModified) {
      const answer = window.confirm(
        "Profile has been modified, are you sure you want to leave without saving?"
      );
      if (answer) {
        next();
      } else {
        next(false);
      }
    } else {
      next();
    }
  },
  mounted() {
    this.loadProfile();
    this.initSplits();
    this.initCSL();
    this.onResize();
    window.addEventListener("beforeunload", this.beforeWindowUnload);
    window.addEventListener("resize", this.onResize);
  },
  beforeUnmount() {
    window.removeEventListener("resize", this.onResize);
    window.removeEventListener("beforeunload", this.beforeWindowUnload);
  },
  watch: {
    outputs: {
      handler() {
        nextTick(() => {
          this.scrollOutput();
        });
      },
      deep: true,
    },
    currentTab(newVal) {
      console.log(newVal);
      if (newVal === 2) {
        this.activeFunction = "sorter";
      } else {
        this.activeFunction = "filter";
      }
    },
  },
  computed: {
    renderedCode() {
      let code = `(import "profile:${this.activeFunction}:${this.profile.name}" fn)
`;
      if (this.activeFunction == "filter") {
        code += `
(fn
${this.renderedReleaseCode(this.releaseA, 1)}
)`;
      } else if (this.activeFunction == "sorter") {
        code += `
(fn
${this.renderedReleaseCode(this.releaseA, 1)}
${this.renderedReleaseCode(this.releaseB, 1)}
)`;
      }
      return code;
    },
    isProfileModified() {
      return (
        this.profile.filter != this.uneditedProfile.filter ||
        this.profile.sorter != this.uneditedProfile.sorter ||
        this.profile.name != this.uneditedProfile.name
      );
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
