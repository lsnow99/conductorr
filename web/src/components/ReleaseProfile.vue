<template>
  <div class="my-2 w-full rounded-lg p-2 border-gray-100 border-1 bg-gray-700">
    <div
      class="
        text-3xl
        w-full
        cursor-pointer
        select-none
        border-gray-100 border-b-2
      "
      @click="computedExpanded = !computedExpanded"
    >
      {{ modelValue.name }}
    </div>
    <transition name="fade">
      <div v-if="computedExpanded" class="p-4">
        Filter
        <CSLEditor readonly v-model="modelValue.filter" />
        Sorter
        <CSLEditor readonly v-model="modelValue.sorter" />
        <div class="flex flex-row justify-between mt-4">
          <o-button variant="danger" @click="deleteProfile">
            <action-button :mode="loadingDelete ? 'loading' : ''"
              >Delete</action-button
            ></o-button
          >
          <o-button variant="primary" @click="edit"
            >Edit</o-button
          >
        </div>
      </div>
    </transition>
  </div>
</template>


<script>
import APIUtil from "../util/APIUtil";
import ActionButton from "./ActionButton.vue";
import CSLEditor from './CSLEditor.vue';

export default {
  components: { ActionButton, CSLEditor },
  data() {
    return {
      test: [0, 40],
      code: `(and
  (in (release-resolution a) ('720p', '1080p', '4k', '8k'))
  (in (release-riptype a) ('BDRIP', 'WEBRIP', 'TVRIP', 'SCR', 'CAM', 'TELESYNC', 'DVDRIP', 'WEBDL'))
  (check-res-bitrate a 'BLURAY-480P' 0 4000000)
  (check-res-bitrate a 'BLURAY-480P' 0 4000000)
  (check-res-bitrate a 'BLURAY-480P' 0 4000000)
  (check-res-bitrate a 'BLURAY-480P' 0 4000000)
)`,
      qualityConfigs: [],
      loadingDelete: false,
      loadingSave: false,
    };
  },
  emits: ["reload", "update:expanded"],
  props: {
    expanded: {
      type: Boolean,
      default: function () {
        return false;
      },
    },
    modelValue: {
      type: Object,
      default: function () {
        return {
          name: "",
          filter: "",
          sort: "",
        };
      },
    },
    ripTypes: {
      type: Array,
      default: function () {
        return [];
      },
    },
    qualityTypes: {
      type: Array,
      default: function () {
        return [];
      },
    },
    resolutionTypes: {
      type: Array,
      default: function () {
        return [];
      },
    },
  },
  methods: {
    dragEnded(values) {
      console.log(values);
    },
    save() {
      this.loadingSave = true;
      APIUtil.updateProfile(
        this.modelValue.id,
        this.modelValue.name,
        this.modelValue.filter,
        this.modelValue.sorter
      )
        .then(() => {
          this.$oruga.notification.open({
            duration: 3000,
            message: `Saved profile ${this.modelValue.name}`,
            position: "bottom-right",
            variant: "success",
            closable: false,
          });
        })
        .finally(() => {
          this.loadingSave = false;
        });
    },
    deleteProfile() {
      this.loadingDelete = true;
      APIUtil.deleteProfile(this.modelValue.id)
        .then(() => {
          this.$oruga.notification.open({
            duration: 3000,
            message: `Deleted profile ${this.modelValue.name}`,
            position: "bottom-right",
            variant: "success",
            closable: false,
          });
          this.$emit("reload");
        })
        .finally(() => {
          this.loadingDelete = false;
        });
    },
    edit() {
      this.$router.push({name: 'editProfile', params: {profile_id: this.modelValue.id}})
    }
  },
  computed: {
    computedExpanded: {
      get() {
        return this.expanded;
      },
      set(newVal) {
        this.$emit("update:expanded", newVal);
      },
    },
  },
  watch: {
    qualityTypes: function (newVal) {
      let configs = [];
      if (newVal) {
        for (let i = 0; i < newVal.length; i++) {
          configs.push({
            name: newVal[i],
            bitrate: 0,
            is_enabled: false,
          });
        }
      }
      this.qualityConfigs = configs;
    },
  },
};
</script>
