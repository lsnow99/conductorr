<template>
  <div>
    <div class="flex flex-row justify-between text-2xl mt-2">
      <div>{{ title }}</div>
      <div>
        <o-tooltip
          class="text-base"
          variant="info"
          :position="tooltipPosition"
          label="Clear"
        >
          <o-icon
            @click="clear"
            class="text-2xl cursor-pointer mx-1"
            icon="times-circle"
          />
        </o-tooltip>
        <o-tooltip
          class="text-base"
          variant="info"
          :position="tooltipPosition"
          label="Randomize"
        >
          <o-icon
            @click="randomize"
            class="text-2xl cursor-pointer mx-1"
            icon="dice"
          />
        </o-tooltip>
      </div>
    </div>
    <div class="flex flex-row">
      <div class="flex flex-1 flex-col p-4">
        <o-field label="Title" class="min-w-full">
          <o-input
            type="text"
            placeholder="Release title"
            v-model="computedRelease.title"
          />
        </o-field>
        <o-field label="Indexer">
          <o-input
            type="text"
            placeholder="Indexer name"
            v-model="computedRelease.indexer"
          />
        </o-field>
        <o-field label="Seeders">
          <o-input
            type="number"
            placeholder="Number of seeders"
            v-model="computedRelease.seeders"
          />
        </o-field>
        <o-field label="Age">
          <o-input
            type="number"
            placeholder="Release age (days)"
            v-model="computedRelease.age"
          />
        </o-field>
        <o-field label="Runtime (minutes)">
          <o-input
            type="number"
            placeholder="Runtime (minutes)"
            v-model="computedRelease.runtime"
            min="0"
          />
        </o-field>
        <o-field addons grouped label="Size">
          <o-input
            type="number"
            placeholder="Size"
            expanded
            v-model="size"
            min="0"
          />
          <o-select v-model="computedRelease.sizeUnit" placeholder="Units">
            <option value="B">B</option>
            <option value="KB">KB</option>
            <option value="MB">MB</option>
            <option value="GB">GB</option>
            <option value="TB">TB</option>
          </o-select>
        </o-field>
      </div>
      <div class="flex flex-1 flex-col p-4">
        <o-field label="Content Type">
          <o-select
            v-model="computedRelease.content_type"
            placeholder="Content type"
          >
            <option value="movie">Movie</option>
            <option value="series">TV Series</option>
          </o-select>
        </o-field>
        <o-field label="Download Type">
          <o-select
            v-model="computedRelease.download_type"
            placeholder="Download type"
          >
            <option value="torrent">Torrent</option>
            <option value="nzb">NZB</option>
          </o-select>
        </o-field>
        <o-field label="Rip Type">
          <o-select v-model="computedRelease.rip_type" placeholder="Rip type">
            <option
              v-for="ripType in RIP_TYPES"
              :key="ripType"
              :value="ripType"
            >
              {{ ripType }}
            </option>
          </o-select>
        </o-field>
        <o-field label="Resolution">
          <o-select
            v-model="computedRelease.resolution"
            placeholder="Resolution"
          >
            <option
              v-for="resolution in RESOLUTION_TYPES"
              :key="resolution"
              :value="resolution"
            >
              {{ resolution }}
            </option>
          </o-select>
        </o-field>
        <o-field label="Encoding">
          <o-select v-model="computedRelease.encoding" placeholder="Encoding">
            <option
              v-for="encoding in ENCODING_TYPES"
              :key="encoding"
              :value="encoding"
            >
              {{ encoding }}
            </option>
          </o-select>
        </o-field>
      </div>
    </div>
  </div>
</template>

<style scoped>
:deep(.o-field) {
  @apply flex-grow-0;
}
</style>

<script>
import APIUtil from "../util/APIUtil";
import { RESOLUTION_TYPES, RIP_TYPES, ENCODING_TYPES } from "../util/Constants";

const EXAMPLE_MOVIES = [
  {
    title: "Manos.The.Hands.of.Fate.1966.THEATRiCAL.1080p.BluRay.x264-SADPANDA",
    resolution: "1080p",
    encoding: "x264",
    rip_type: "BDRip",
    content_type: "movie",
  },
  {
    title:
      "The.Last.Man.On.Earth.1964.1080p.BluRay.Plus.Comm.DTS.x264-MaG-Obfuscated",
    resolution: "1080p",
    encoding: "x264",
    rip_type: "TELESYNC",
    content_type: "movie",
  },
  {
    title:
      "Night.of.the.Living.Dead.1968.1080p.BluRay.CRITERION.Plus.Comms.FLAC.x264-MaG-Obfuscated",
    resolution: "1080p",
    encoding: "x264",
    rip_type: "BDRip",
    content_type: "movie",
  },
  {
    title:
      "Santa.Claus.Conquers.the.Martians.1964.1080p.BDRip.DTS.x265.10bit-MarkII",
    resolution: "1080p",
    encoding: "x265",
    rip_type: "TELESYNC",
    content_type: "movie",
  },
  {
    title: "The.Terror.1963.720p.BluRay.DTS.x264-DJ",
    resolution: "720p",
    encoding: "x264",
    rip_type: "TELESYNC",
    content_type: "movie",
  },
];

export default {
  data() {
    return {
      RESOLUTION_TYPES,
      RIP_TYPES,
      ENCODING_TYPES,
      size: null,
      indexers: [],
      tooltipPosition: "bottom",
    };
  },
  props: {
    modelValue: {
      type: Object,
      default: function () {
        return {};
      },
    },
    title: {
      type: String,
      default: function () {
        return "";
      },
    },
  },
  emits: ["update:modelValue"],
  methods: {
    randomize() {
      const release =
        EXAMPLE_MOVIES[Math.floor(Math.random() * EXAMPLE_MOVIES.length)];
      if (release.title == this.computedRelease.title) {
        this.randomize();
        return;
      }
      const size = Math.floor(Math.random() * 15 * Math.pow(2, 30));
      release.size = size;
      release.sizeUnit = undefined;
      release.runtime = Math.floor(Math.random() * 300);
      release.age = Math.floor(Math.random() * 3000);

      if (this.indexers && this.indexers.length > 0) {
        const randomIndexer =
          this.indexers[Math.floor(Math.random() * this.indexers.length)];
        release.indexer = randomIndexer.name;
        release.download_type = randomIndexer.download_type;
        if (release.download_type == "torrent") {
          release.seeders = Math.floor(Math.random() * 50);
        } else {
          release.seeders = 0;
        }
      }

      this.computedRelease = release;
    },
    clear() {
      this.size = null
      this.$emit('update:modelValue', {})
    },
    updateSize() {
      if(this.size == null) {
        return
      }
      let newRelease = this.computedRelease;
      newRelease.size = this.size;
      switch (newRelease.sizeUnit) {
        case "B":
          break;
        case "KB":
          newRelease.size *= Math.pow(2, 10);
          break;
        case "MB":
          newRelease.size *= Math.pow(2, 20);
          break;
        case "GB":
          newRelease.size *= Math.pow(2, 30);
          break;
        case "TB":
          newRelease.size *= Math.pow(2, 40);
          break;
        default:
          break;
      }
      this.$emit("update:modelValue", newRelease);
    },
  },
  mounted() {
    APIUtil.getIndexers().then((indexers) => {
      this.indexers = indexers;
      this.randomize()
    });
    const screenWidth = window.innerWidth;
    if (screenWidth < 768) {
      this.tooltipPosition = "left";
    }
  },
  computed: {
    computedRelease: {
      get() {
        const release = this.modelValue;
        if (release.sizeUnit) {
          return release;
        }
        let sizes = ["B", "KB", "MB", "GB", "TB"];
        if (release.size == 0 || isNaN(release.size)) {
          return release;
        }
        let i = parseInt(Math.floor(Math.log(release.size) / Math.log(1024)));
        this.size = Math.round(release.size / Math.pow(1024, i), 2);
        release.sizeUnit = sizes[i];
        return release;
      },
      set(newVal) {
        this.$emit("update:modelValue", newVal);
      },
    },
    sizeUnit() {
      return this.modelValue.sizeUnit;
    },
  },
  watch: {
    sizeUnit() {
      this.updateSize();
    },
    size() {
      this.updateSize();
    },
  },
};
</script>
