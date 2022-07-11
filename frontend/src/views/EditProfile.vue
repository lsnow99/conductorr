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
                Rendered code (you can assume that your script will be run like
                this):
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

<script setup lang="ts">
import { computed, nextTick, onMounted, ref, watch, WritableComputedRef } from "vue";
import APIUtil from "../util/APIUtil";
import { CSLEditor, LogPane, Variant, LogMessage } from "conductorr-lib";
import Split from "split.js";
import "../util/wasm_exec.js";
import { DateTime } from "luxon";
import ReleaseBuilder from "../components/ReleaseBuilder.vue";
import { Release } from "@/types/api/release";
import { onBeforeRouteLeave, useRoute } from "vue-router";

enum ScriptFunction {
  FILTER = "filter",
  SORTER = "sorter",
}

enum TabOptions {
  ONE = 1,
  TWO = 2,
}

const split12 = ref<Split.Instance | null>(null);
const split34 = ref<Split.Instance | null>(null);
const split56 = ref<Split.Instance | null>(null);
const profile = ref<Profile | null>(null);
const uneditedProfile = ref<Profile | null>(null);
const profileID = ref<number | null>(null);
const code = ref("");
const outputs = ref<LogMessage[]>();
const activeFunction = ref<ScriptFunction>(ScriptFunction.FILTER);
const releaseA = ref<Release | null>(null);
const releaseB = ref<Release | null>(null);
const editingName = ref(false);
const oldName = ref("");
const currentTab = ref<TabOptions>(1);
const headerHeight = ref<number>(0);
const outputScroller = ref<HTMLElement | null>(null);
const header = ref<HTMLElement | null>(null);


const route = useRoute();
profileID.value = parseInt(route.params.profile_id as string);

const pushOutput = (msg: string, variant: Variant) => {
  outputs.value?.push({
    msg,
    variant,
    timestamp: DateTime.now(),
  });
};

const loadSizes = (name: string, defaultSizes = [50, 50]) => {
  const sizes = localStorage.getItem(name);
  return sizes ? JSON.parse(sizes) : defaultSizes;
};

const initSplits = (reset = false) => {
  if (split12.value) {
    split12.value.destroy();
    split12.value = null;
  }
  if (split34.value) {
    split34.value.destroy();
    split34.value = null;
  }
  if (split56.value) {
    split56.value.destroy();
    split56.value = null;
  }

  split12.value = Split(["#split1", "#split2"], {
    sizes: reset ? undefined : loadSizes("split-sizes12"),
    minSize: 200,
    direction: "vertical",
    onDragEnd: function (sizes) {
      localStorage.setItem("split-sizes12", JSON.stringify(sizes));
    },
  });

  split34.value = Split(["#split3", "#split4"], {
    sizes: reset ? [40, 60] : loadSizes("split-sizes34", [40, 60]),
    minSize: 200,
    onDragEnd: function (sizes) {
      localStorage.setItem("split-sizes34", JSON.stringify(sizes));
    },
  });

  split56.value = Split(["#split5", "#split6"], {
    direction: "vertical",
    minSize: 200,
    sizes: reset ? [70, 30] : loadSizes("split-sizes56", [70, 30]),
    onDragEnd: function (sizes) {
      localStorage.setItem("split-sizes56", JSON.stringify(sizes));
    },
  });

  if (reset) {
    localStorage.setItem("split-sizes12", "");
    localStorage.setItem("split-sizes34", "");
    localStorage.setItem("split-sizes56", "");
  }
};

const validate = () => {
  Validate(computedCode, (ok: boolean, err: string) => {
    if (!ok) {
      pushOutput(`Validation error: ${err}`, Variant.DANGER);
    } else {
      pushOutput(`Validation succeeded`, Variant.SUCCESS);
    }
  });
};

const run = () => {
  let env: any = {
    a: releaseA.value,
  };
  if (activeFunction.value === ScriptFunction.SORTER) env.b = releaseB.value;
  Run(computedCode, env, (ok: boolean, err: string | null, result: any) => {
    if (!ok) {
      pushOutput(`Execution error: ${err}`, Variant.DANGER);
    } else if (result || result === 0 || result === false) {
      console.log(result);
      pushOutput(result, Variant.SUCCESS);
    } else {
      pushOutput("Script returned null", Variant.WARNING);
    }
  });
};

const scrollOutput = () => {
  if (outputScroller.value)
    outputScroller.value.scrollTop = outputScroller.value.scrollHeight;
};

const renderedReleaseCode = (release: Release | null, indentAmount: number) => {
  let indent = "";
  for (let i = 0; i < indentAmount; i++) {
    indent += "  ";
  }
  let code = `{CSL_INDENT}(
{CSL_INDENT}  "${release?.title}"
{CSL_INDENT}  "${release?.indexer}"
{CSL_INDENT}  "${release?.downloadType}"
{CSL_INDENT}  "${release?.contentType}"
{CSL_INDENT}  "${release?.ripType}"
{CSL_INDENT}  "${release?.resolution}"
{CSL_INDENT}  "${release?.encoding}"
{CSL_INDENT}  ${release?.seeders}
{CSL_INDENT}  ${release?.age}
{CSL_INDENT}  ${release?.size}
{CSL_INDENT}  ${release?.runtime}
{CSL_INDENT})`;
  return code.replaceAll("{CSL_INDENT}", indent);
};

const setUneditedProfile = (name: string, filter: string, sorter: string) => {
  uneditedProfile.name = name;
  uneditedProfile.filter = filter;
  uneditedProfile.sorter = sorter;
};

const updateProfile = async () => {
  try {
    await APIUtil.updateProfile(
      profileID.value,
      profile.value.name,
      profile.value.filter,
      profile.value.sorter
    );
    oruga.notification.open({
      duration: 3000,
      message: `Saved successfully`,
      position: "bottom-right",
      variant: "success",
      closable: false,
    });
    setUneditedProfile(
      profile.value.name,
      profile.value.filter,
      profile.value.sorter
    );
  } catch {}
};

const initCSL = () => {
  let go = new Go();

  WebAssembly.instantiateStreaming(fetch("/api/csl.wasm"), go.importObject)
    .then(async (result) => {
      await go.run(result.instance);
      initCSL();
    })
    .catch((err) => {
      console.log(err);
    });
};

const loadProfile = async () => {
  try {
    const loadedProfile = await APIUtil.getProfile(profileID.value);
    profile.value = loadProfile;
    setUneditedProfile(
      loadedProfile.name,
      loadedProfile.filter,
      loadedProfile.sorter
    );
  } catch {}
};


const onResize = () => (headerHeight.value = header.value?.clientHeight ?? 0);

const startEditingName = () => {
  oldName.value = profile.value.name;
  editingName.value = true;
};

const resetEditingName = () => {
  profile.name = oldName.value;
  editingName.value = false;
};


const isProfileModified = computed(
  () =>
    profile.value.filter != uneditedProfile.value.filter ||
    profile.value.sorter != uneditedProfile.value.sorter ||
    profile.value.name != uneditedProfile.value.name
);

const ARE_YOU_SURE_MSG =
  "Profile has been modified, are you sure you want to leave without saving?";
  
onBeforeRouteLeave((to, from, next) => {
  if (isProfileModified.value) {
    const answer = window.confirm(ARE_YOU_SURE_MSG);
    if (answer) {
      next();
    } else {
      next(false);
    }
  } else {
    next();
  }
});

const beforeWindowUnload = (e: Event) => {
  if (isProfileModified.value) {
    e.preventDefault();
    return ARE_YOU_SURE_MSG;
  }
  return null;
};


onMounted(() => {
  loadProfile();
  initSplits();
  initCSL();
  onResize();
  window.addEventListener("beforeunload", beforeWindowUnload);
  window.addEventListener("resize", onResize);
});

watch(
  outputs,
  () => {
    nextTick(() => {
      scrollOutput();
    });
  },
  { deep: true }
);

watch(currentTab, (newVal) => {
  if (newVal === 2) {
    activeFunction.value = ScriptFunction.SORTER;
  } else {
    activeFunction.value = ScriptFunction.FILTER;
  }
});

const renderedCode = computed(() => {
  let code = `(import "profile:${activeFunction.value}:${profile.value.name}" fn)
`;
  if (activeFunction.value == "filter") {
    code += `
(fn
${renderedReleaseCode(releaseA.value, 1)}
)`;
  } else if (activeFunction.value == "sorter") {
    code += `
(fn
${renderedReleaseCode(releaseA.value, 1)}
${renderedReleaseCode(releaseB.value, 1)}
)`;
  }
  return code;
});

const computedCode: WritableComputedRef<string> = computed({
  get() {
    if (activeFunction.value === ScriptFunction.FILTER) {
      return profile.value.filter
    } else if (activeFunction.value === ScriptFunction.SORTER) {
      return profile.value.sorter
    }
  },
  set(newVal: string) {
    if (activeFunction.value === ScriptFunction.FILTER) {
      profile.value.filter = newVal
    } else if (activeFunction.value === ScriptFunction.SORTER) {
      profile.value.sorter = newVal
    }
  }
})
</script>
