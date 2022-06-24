<template>
  <modal :title="media.title" v-model="computedActive" @close="$emit('close')">
    <o-field
      label="Profile"
      :variant="profileVariant"
      :message="profileVariant ? 'A profile is required' : ''"
    >
      <o-select expanded v-model="profileID" placeholder="Profile">
        <option
          v-for="profileOption in profiles"
          :key="profileOption.id"
          :value="profileOption.id"
        >
          {{ profileOption.name }}
        </option>
      </o-select>
      <o-button @click="newProfile">New</o-button>
    </o-field>
    <o-field
      label="Root Path"
      :variant="pathVariant"
      :message="pathVariant ? 'A root path is required' : ''"
    >
      <o-select
        expanded
        v-model="pathID"
        :placeholder="`/media/library/${media.content_type}`"
      >
        <option
          v-for="pathOption in paths"
          :key="pathOption.id"
          :value="pathOption.id"
        >
          {{ pathOption.path }}
        </option>
      </o-select>
      <o-button @click="newPath">New</o-button>
    </o-field>
    <template v-slot:footer>
      <o-button @click="$emit('close')">Cancel</o-button>
      <div>
        <o-button variant="primary" @click="save">
          <action-button :mode="loading ? 'loading' : ''">Save</action-button>
        </o-button>
      </div>
    </template>
  </modal>
  <new-profile
    v-model:active="showNewProfileModal"
    @close="closeNewProfile"
    @submitted="newProfileSubmitted"
  />
  <edit-path
    v-model:active="showNewPathModal"
    @close="closeNewPath"
    @submitted="newPathSubmitted"
  />
</template>

<script setup lang="ts">
import APIUtil from "@/util/APIUtil";
import NewProfile from "@/components/NewProfile.vue";
import EditPath from "@/components/EditPath.vue";
import ActionButton from "@/components/ActionButton.vue";
import Modal from "@/components/Modal.vue";
import useTabSaver from "@/util/TabSaver";
import { computed, onMounted, ref, watch, WritableComputedRef } from "vue";
import { Media } from "@/types/api/media";
import { inject } from "vue";
import { Path } from "@/types/api/path";

const oruga = inject("oruga");

const { lastButton, restoreFocus } = useTabSaver();

const profiles = ref([]);
const paths = ref([]);
const profileID = ref(null);
const pathID = ref(null);
const showNewProfileModal = ref(false);
const showNewPathModal = ref(false);
const submittedOnce = ref(false);

const props = defineProps<{
  media: Media;
  loading: boolean;
  active: boolean;
}>();

const emit = defineEmits<{
  (e: "close"): void;
  (e: "submit", data: { profileID: number; pathID: number }): void;
  (e: "update:active", active: boolean): void;
}>();

const save = () => {
  submittedOnce.value = true;
  if (profileID.value && pathID.value) {
    emit("submit", {
      profileID: profileID.value,
      pathID: pathID.value,
    });
  }
};

const loadProfiles = async () => {
  profiles.value = await APIUtil.getProfiles();
};

const loadPaths = async () => {
  paths.value = await APIUtil.getPaths();
};

const newPath = ($event: Event) => {
  showNewPathModal.value = true;
  lastButton.value = $event.currentTarget as HTMLElement;
};

const newPathSubmitted = async (path: Path) => {
  try {
    await APIUtil.createNewPath(
      path.path,
      path.moviesDefault,
      path.seriesDefault
    );
    oruga.notification.open({
      duration: 3000,
      message: `Created path ${path.path}`,
      position: "bottom-right",
      variant: "success",
      closable: false,
    });
  } finally {
    showNewPathModal.value = false;
    loadPaths();
  }
};

const newProfile = ($event: Event) => {
  lastButton.value = $event.currentTarget as HTMLElement;
  showNewProfileModal.value = true;
};

const closeNewPath = () => {
  showNewPathModal.value = false;
  restoreFocus();
};

const closeNewProfile = () => {
  showNewProfileModal.value = false;
  restoreFocus();
};

const newProfileSubmitted = () => {
  loadProfiles();
  showNewProfileModal.value = false;
  restoreFocus();
};

onMounted(() => {
  loadProfiles();
  loadPaths();
});

watch(() => props.media, (newVal: Media) => {
  if(newVal) {
    profileID.value = newVal.profile_id

    // Set the default path
    if (!props.media.path_id) {
      paths.value.forEach((elem: Path) => {
        if (
          (elem.movies_default && props.media.content_type == "movie") ||
          (elem.series_default && props.media.content_type == "series")
        ) {
          pathID.value = elem.id;
        }
      })
    } else {
      pathID.value = props.media.path_id
    }
  }
}, {immediate: true})

const computedActive: WritableComputedRef<boolean> = computed({
  get(): boolean {
    return props.active
  },
  set(v): void {
    emit("update:active", v)
  }
})

const profileVariant = computed(() => {
  if(!submittedOnce.value) {
    return ""
  }
  if (!profileID.value) {
    return "danger"
  }
})

const pathVariant = computed(() => {
  if(!submittedOnce.value) {
    return ""
  }
  if (!pathID.value) {
    return "danger"
  }
})
</script>
