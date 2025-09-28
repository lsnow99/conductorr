<template>
  <Modal
    :title="path ? 'Editing Path' : 'New Path'"
    v-model="computedActive"
    @close="$emit('close')"
  >
    <o-field label="Path">
      <o-input
        expanded
        @keydown.enter="submit"
        placeholder="/media/library/name"
        v-model="pathDir"
      />
      <o-button @click="testPath">Test</o-button>
    </o-field>
    <o-field label="Defaults">
      <div class="flex flex-col">
        <o-switch v-model="moviesDefault">Default for Movies</o-switch>
        <o-switch v-model="seriesDefault">Default for Series</o-switch>
      </div>
    </o-field>
    <template v-slot:footer>
      <o-button @click="$emit('close')">Cancel</o-button>
      <div>
        <o-button variant="primary" @click="submit"
          ><ActionButton :mode="loading ? 'loading' : 'off'"
            >Submit</ActionButton
          ></o-button
        >
      </div>
    </template>
  </Modal>
</template>

<script setup lang="ts">
import { NewPath, Path } from "@/types/api/path";
import { useComputedActive } from "@/util";
import { inject, ref, watch } from "vue";
import APIUtil from "../util/APIUtil";
import ActionButton from "./ActionButton.vue";
import Modal from "./Modal.vue";

const pathDir = ref<string | null>(null);
const moviesDefault = ref(false);
const seriesDefault = ref(false);
const loading = ref(false);

const props = defineProps<{
  path: Path | null;
  active: boolean;
}>();

const emit = defineEmits<{
  (e: "close"): void;
  (e: "submitted", path: NewPath): void;
  (e: "update:active", active: boolean): void;
}>();

const submit = () => {
  emit("submitted", {
    path: pathDir.value ?? "",
    moviesDefault: moviesDefault.value,
    seriesDefault: seriesDefault.value,
  });
};

const oruga = inject("oruga");

const testPath = async () => {
  try {
    await APIUtil.testPath(pathDir.value);
    oruga.notification.open({
      duration: 3000,
      message: `Path is OK`,
      position: "bottom-right",
      variant: "success",
      closable: false,
    });
  } catch (err) {
    oruga.notification.open({
      duration: 5000,
      message: `Test failed: ${err}`,
      position: "bottom-right",
      variant: "danger",
      closable: false,
    });
  }
};

watch(
  () => props.path,
  (newVal: Path | null) => {
    if (newVal) {
      pathDir.value = newVal.path;
      moviesDefault.value = newVal.moviesDefault;
      seriesDefault.value = newVal.seriesDefault;
    }
  },
  {
    immediate: true,
  }
);

const computedActive = useComputedActive(props, emit);
</script>
