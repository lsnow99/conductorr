<template>
  <div>
    <div class="flex flex-row">
      <div class="flex flex-1 flex-col p-4">
        <o-field label="Title">
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
        <o-field label="Quality">
          <o-select v-model="computedRelease.quality" placeholder="Quality">
            <option
              v-for="quality in QUALITY_TYPES"
              :key="quality"
              :value="quality"
            >
              {{ quality }}
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
        <o-field label="Bitrate">
          <o-slider rounded v-model="computedRelease.bitrate" />
        </o-field>
      </div>
    </div>
  </div>
</template>

<script>
import {
  RESOLUTION_TYPES,
  QUALITY_TYPES,
  RIP_TYPES,
  ENCODING_TYPES,
} from "../util/Constants";

export default {
  data() {
    return {
      RESOLUTION_TYPES,
      QUALITY_TYPES,
      RIP_TYPES,
      ENCODING_TYPES,
    };
  },
  props: {
    modelValue: {
      type: Object,
      default: function () {
        return {};
      },
    },
  },
  emits: ["update:modelValue"],
  computed: {
    computedRelease: {
      get() {
        return this.modelValue;
      },
      set(newVal) {
        this.$emit("update:modelValue", newVal);
      },
    },
  },
};
</script>
