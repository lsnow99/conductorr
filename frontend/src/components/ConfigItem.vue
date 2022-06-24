<template>
  <div class="w-full p-2 my-2 bg-gray-700 border-gray-100 rounded-lg border-1">
    <div
      @click="doExpand"
      @keydown.enter="doExpand"
      @keydown.space="doExpand"
      role="button"
      tabindex="0"
      :aria-label="ariaLabel"
      :aria-disabled="!collapsible"
      :class="collapsible ? 'cursor-pointer' : ''"
      class="flex flex-row justify-between w-full text-3xl border-b-2 border-gray-100 select-none focus:outline-white focus:opacity-80"
    >
      <div>
        {{ title }}
      </div>
      <div v-if="collapsible">
        <o-icon v-if="isExpanded" icon="chevron-up" />
        <o-icon v-else icon="chevron-down" />
      </div>
    </div>
    <transition name="fade">
      <slot v-if="!collapsible || isExpanded" />
    </transition>
    <div class="flex flex-row justify-between mt-4">
      <o-button variant="danger" @click="confirmDelete">Delete</o-button>
      <o-button variant="primary" @click="$emit('edit', $event)">Edit</o-button>
    </div>
  </div>
  <ConfirmDelete
    v-model:active="showConfirmDeleteModal"
    @delete="$emit('delete')"
    @close="closeDelete"
    :delete-message="deleteMessage"
  />
</template>

<script setup lang="ts">
import ConfirmDelete from "./ConfirmDelete.vue";
import useTabSaver from "../util/TabSaver";
import { computed, ref, watch } from "vue";

const props = defineProps<{
  title: string;
  collapsible: boolean;
  deleteMessage: string;
  expanded: boolean;
}>();

const emit = defineEmits<{
  (e: "update:expanded", expanded: boolean): void;
  (e: "delete"): void;
  (e: "edit", $event: Event): void;
}>();

const showConfirmDeleteModal = ref(false);
const isExpanded = ref(props.expanded);

const { lastButton, restoreFocus } = useTabSaver();

const closeDelete = () => {
  showConfirmDeleteModal.value = false;
  restoreFocus();
};

const confirmDelete = ($event: Event) => {
  lastButton.value = $event.currentTarget as HTMLElement;
  showConfirmDeleteModal.value = true;
};

const doExpand = () => {
  if (props.collapsible) {
    emit("update:expanded", !isExpanded.value);
    isExpanded.value = !isExpanded.value;
  }
};

watch(
  () => props.expanded,
  (v: boolean) => {
    isExpanded.value = v;
  }
);

const ariaLabel = computed(() => {
  isExpanded.value ? "Collapse" : "Expand";
});
</script>
