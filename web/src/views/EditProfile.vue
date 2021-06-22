<template>
  <div class="flex flex-col h-screen max-h-screen">
    <div class="flex flex-col bg-gray-900 bg-opacity-20">
      <div class="p-2 flex flex-1 flex-row justify-between">
        <div class="w-60">
          <o-button icon-left="chevron-left">Back</o-button>
        </div>
        <div class="flex text-2xl items-center">
          Editing Profile: {{ profile.name }}
        </div>
        <div class="w-60"></div>
      </div>
      <div class="flex flex-row justify-between">
        <div>
          <o-checkbox v-model="showGenerator" variant="primary"
            >Show Generator</o-checkbox
          >
        </div>
        <div>
          <o-button @click="initSplits(true)">Reset View</o-button>
          <o-button variant="primary" @click="validate">Validate</o-button>
          <o-button variant="primary">Save</o-button>
        </div>
      </div>
    </div>
    <div class="flex h-full flex-1" style="max-height: calc(100vh - 104px)">
      <div class="flex h-full flex-1 flex-row">
        <div id="split3" class="flex flex-col">
          <div v-show="showGenerator" id="split1" class="flex">Test1</div>
          <div id="split2" class="flex h-full">
            <CSLEditor v-model="code" />
          </div>
        </div>
        <div id="split4" class="flex flex-col">
          <div id="split5" class="flex">Test2</div>
          <div id="split6" class="flex">Test3</div>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
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

export default {
  data() {
    return {
      split12: null,
      split34: null,
      split56: null,
      profile: {},
      profile_id: 0,
      showGenerator: true,
      code: `(and
  (in (release-resolution a) ('720p', '1080p', '4k', '8k'))
  (in (release-riptype a) ('BDRIP', 'WEBRIP', 'TVRIP', 'SCR', 'CAM', 'TELESYNC', 'DVDRIP', 'WEBDL'))
  (check-res-bitrate a 'BLURAY-480P' 0 4000000)
  (check-res-bitrate a 'BLURAY-480P' 0 4000000)
  (check-res-bitrate a 'BLURAY-480P' 0 4000000)
  (check-res-bitrate a 'BLURAY-480P' 0 4000000)
)
(and
  (in (release-resolution a) ('720p', '1080p', '4k', '8k'))
  (in (release-riptype a) ('BDRIP', 'WEBRIP', 'TVRIP', 'SCR', 'CAM', 'TELESYNC', 'DVDRIP', 'WEBDL'))
  (check-res-bitrate a 'BLURAY-480P' 0 4000000)
  (check-res-bitrate a 'BLURAY-480P' 0 4000000)
  (check-res-bitrate a 'BLURAY-480P' 0 4000000)
  (check-res-bitrate a 'BLURAY-480P' 0 4000000)
)
(and
  (in (release-resolution a) ('720p', '1080p', '4k', '8k'))
  (in (release-riptype a) ('BDRIP', 'WEBRIP', 'TVRIP', 'SCR', 'CAM', 'TELESYNC', 'DVDRIP', 'WEBDL'))
  (check-res-bitrate a 'BLURAY-480P' 0 4000000)
  (check-res-bitrate a 'BLURAY-480P' 0 4000000)
  (check-res-bitrate a 'BLURAY-480P' 0 4000000)
  (check-res-bitrate a 'BLURAY-480P' 0 4000000)
)
(and
  (in (release-resolution a) ('720p', '1080p', '4k', '8k'))
  (in (release-riptype a) ('BDRIP', 'WEBRIP', 'TVRIP', 'SCR', 'CAM', 'TELESYNC', 'DVDRIP', 'WEBDL'))
  (check-res-bitrate a 'BLURAY-480P' 0 4000000)
  (check-res-bitrate a 'BLURAY-480P' 0 4000000)
  (check-res-bitrate a 'BLURAY-480P' 0 4000000)
  (check-res-bitrate a 'BLURAY-480P' 0 4000000)
)
(and
  (in (release-resolution a) ('720p', '1080p', '4k', '8k'))
  (in (release-riptype a) ('BDRIP', 'WEBRIP', 'TVRIP', 'SCR', 'CAM', 'TELESYNC', 'DVDRIP', 'WEBDL'))
  (check-res-bitrate a 'BLURAY-480P' 0 4000000)
  (check-res-bitrate a 'BLURAY-480P' 0 4000000)
  (check-res-bitrate a 'BLURAY-480P' 0 4000000)
  (check-res-bitrate a 'BLURAY-480P' 0 4000000)
)
`,
    };
  },
  components: { CSLEditor },
  methods: {
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

      if (this.showGenerator) {
        this.split12 = Split(["#split1", "#split2"], {
          sizes: reset?undefined:this.loadSizes("split-sizes12"),
          minSize: 0,
          direction: "vertical",
          onDragEnd: function (sizes) {
            localStorage.setItem("split-sizes12", JSON.stringify(sizes));
          },
        });
      }

      this.split34 = Split(["#split3", "#split4"], {
        sizes: reset?undefined:this.loadSizes("split-sizes34"),
        minSize: 0,
        onDragEnd: function (sizes) {
          localStorage.setItem("split-sizes34", JSON.stringify(sizes));
        },
      });

      this.split56 = Split(["#split5", "#split6"], {
        direction: "vertical",
        minSize: 0,
        sizes: reset?undefined:this.loadSizes("split-sizes56"),
        onDragEnd: function (sizes) {
          localStorage.setItem("split-sizes56", JSON.stringify(sizes));
        },
      });
    },
    validate() {
      let err = Validate(this.code)
      if (err) {
        this.$oruga.notification.open({
          duration: 3000,
          message: `Validation error: ${err}`,
          position: "bottom-right",
          variant: "danger",
          closable: false,
        });
      } else {
        this.$oruga.notification.open({
          duration: 3000,
          message: `Parsed successfully!`,
          position: "bottom-right",
          variant: "success",
          closable: false,
        });
      }
    }
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

    let go = new Go()

    WebAssembly.instantiateStreaming(fetch("/api/csl.wasm"), go.importObject).then(
      (result) => {
        go.run(result.instance)
      }
    );
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
  },
};
</script>
