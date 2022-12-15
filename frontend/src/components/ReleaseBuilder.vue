<template>
  <div>
    <div class="flex flex-row justify-between mt-2 text-2xl">
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
            class="mx-1 text-2xl cursor-pointer"
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
            class="mx-1 text-2xl cursor-pointer"
            icon="dice"
          />
        </o-tooltip>
      </div>
    </div>
    <div class="flex flex-row">
      <div class="flex flex-col flex-1 p-4">
        <o-field label="Title" class="min-w-full">
          <o-input
            type="text"
            placeholder="Release title"
            v-model="computedRelease!.title"
          />
        </o-field>
        <o-field label="Indexer">
          <o-input
            type="text"
            placeholder="Indexer name"
            v-model="computedRelease!.indexer"
          />
        </o-field>
        <o-field label="Seeders">
          <o-input
            type="number"
            @change="numberChanged"
            placeholder="Number of seeders"
            v-model="computedRelease!.seeders"
          />
        </o-field>
        <o-field label="Age">
          <o-input
            type="number"
            @change="numberChanged"
            placeholder="Release age (days)"
            v-model="computedRelease!.age"
          />
        </o-field>
        <o-field label="Runtime (minutes)">
          <o-input
            type="number"
            @change="numberChanged"
            placeholder="Runtime (minutes)"
            v-model="computedRelease!.runtime"
            min="0"
          />
        </o-field>
        <o-field addons grouped label="Size">
          <o-input
            type="number"
            @change="numberChanged"
            placeholder="Size"
            expanded
            v-model="size"
            min="0"
          />
          <o-select v-model="computedRelease!.sizeUnit" placeholder="Units">
            <option value="B">B</option>
            <option value="KB">KB</option>
            <option value="MB">MB</option>
            <option value="GB">GB</option>
            <option value="TB">TB</option>
          </o-select>
        </o-field>
      </div>
      <div class="flex flex-col flex-1 p-4">
        <o-field label="Content Type">
          <o-select
            v-model="computedRelease!.contentType"
            placeholder="Content type"
          >
            <option value="movie">Movie</option>
            <option value="series">TV Series</option>
          </o-select>
        </o-field>
        <o-field label="Download Type">
          <o-select
            v-model="computedRelease!.downloadType"
            placeholder="Download type"
          >
            <option value="torrent">Torrent</option>
            <option value="nzb">NZB</option>
          </o-select>
        </o-field>
        <o-field label="Rip Type">
          <o-select v-model="computedRelease!.ripType" placeholder="Rip type">
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
            v-model="computedRelease!.resolution"
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
          <o-select v-model="computedRelease!.encoding" placeholder="Encoding">
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

<script setup lang="ts">
import {
  computed,
  nextTick,
  onMounted,
  ref,
  watch,
  WritableComputedRef,
} from "vue";
import APIUtil from "../util/APIUtil";
import { RESOLUTION_TYPES, RIP_TYPES, ENCODING_TYPES } from "@/util/Constants";
import { TooltipPosition } from "@/types/tooltip";
import { Release } from "@/types/api/release";
import { Indexer } from "@/types/api/indexer";

const encodingTypes = ref(ENCODING_TYPES);
const ripTypes = ref(RIP_TYPES);
const resolutionTypes = ref(RESOLUTION_TYPES);
const exampleMovies = ref<Release[]>([
  {
    title: "Manos.The.Hands.of.Fate.1966.THEATRiCAL.1080p.BluRay.x264-SADPANDA",
    resolution: "1080p",
    encoding: "x264",
    ripType: "BDRip",
    contentType: "movie",
  },
  {
    title:
      "The.Last.Man.On.Earth.1964.1080p.BluRay.Plus.Comm.DTS.x264-MaG-Obfuscated",
    resolution: "1080p",
    encoding: "x264",
    ripType: "TELESYNC",
    contentType: "movie",
  },
  {
    title:
      "Night.of.the.Living.Dead.1968.1080p.BluRay.CRITERION.Plus.Comms.FLAC.x264-MaG-Obfuscated",
    resolution: "1080p",
    encoding: "x264",
    ripType: "BDRip",
    contentType: "movie",
  },
  {
    title:
      "Santa.Claus.Conquers.the.Martians.1964.1080p.BDRip.DTS.x265.10bit-MarkII",
    resolution: "1080p",
    encoding: "x265",
    ripType: "TELESYNC",
    contentType: "movie",
  },
  {
    title: "The.Terror.1963.720p.BluRay.DTS.x264-DJ",
    resolution: "720p",
    encoding: "x264",
    ripType: "TELESYNC",
    contentType: "movie",
  },
]);

const size = ref<number | null>(null);
const indexers = ref<Indexer[]>([]);
const tooltipPosition = ref<TooltipPosition>("bottom");

const props = defineProps<{
  modelValue: Release | null;
  title: string;
}>();

const emit = defineEmits<{
  (e: "update:modelValue", newVal: Release | null): void;
}>();

const computedRelease: WritableComputedRef<Release | null> = computed({
  get() {
    const release = props.modelValue;
    if (release?.sizeUnit) {
      return release;
    }
    let sizes = ["B", "KB", "MB", "GB", "TB"];
    if (release?.size == 0 || !release?.size || isNaN(release.size)) {
      return release;
    }
    let i = Math.floor(Math.log(release.size) / Math.log(1024));
    size.value = Math.round(release.size / Math.pow(1024, i));
    release.sizeUnit = sizes[i];
    return release;
  },
  set(newVal) {
    emit("update:modelValue", newVal);
  },
});

const randomize = () => {
  const release =
    exampleMovies.value[Math.floor(Math.random() * exampleMovies.value.length)];
  if (release.title == computedRelease.value?.title) {
    randomize();
    return;
  }
  const randSize = Math.floor(Math.random() * 15 * Math.pow(2, 30));
  release.size = randSize;
  release.sizeUnit = undefined;
  release.runtime = Math.floor(Math.random() * 300);
  release.age = Math.floor(Math.random() * 3000);

  if (indexers.value && indexers.value.length > 0) {
    const randomIndexer =
      indexers.value[Math.floor(Math.random() * indexers.value.length)];
    release.indexer = randomIndexer.name;
    release.downloadType = randomIndexer.downloadType;
    if (release.downloadType == "torrent") {
      release.seeders = Math.floor(Math.random() * 50);
    } else {
      release.seeders = 0;
    }
  }

  computedRelease.value = release;
};

const clear = () => {
  size.value = null;
  emit("update:modelValue", null);
};

const updateSize = () => {
  if (!size.value) {
    return;
  }
  let newRelease: Release = computedRelease.value ?? {
    size: 0,
    title: "",
    encoding: "",
    ripType: "",
    contentType: "",
    resolution: "",
  };
  newRelease.size = size.value;
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
  emit("update:modelValue", newRelease);
};

const forceInt = (x: string | number | undefined | null, fallback = -1) => {
  const val = parseInt(`${x}`);
  if (val === NaN) {
    return fallback;
  }
  return val;
};

const numberChanged = () => {
  nextTick(() => {
    if (computedRelease.value) {
      computedRelease.value.seeders = forceInt(computedRelease.value.seeders);
      computedRelease.value.age = forceInt(computedRelease.value.age);
      computedRelease.value.runtime = forceInt(computedRelease.value.runtime);
      computedRelease.value.size = forceInt(computedRelease.value.size);
    }
  });
};

onMounted(async () => {
  try {
    const loadedIndexers = await APIUtil.getIndexers();
    indexers.value = loadedIndexers;
    randomize();
  } catch (error) {
    // TODO: error to user
  }
  const screenWidth = window.innerWidth;
  if (screenWidth < 768) {
    tooltipPosition.value = "left";
  }
});

const sizeUnit = computed(() => {
  return props.modelValue?.sizeUnit;
});

watch(sizeUnit, () => {
  updateSize();
});

watch(size, () => {
  updateSize();
});
</script>
