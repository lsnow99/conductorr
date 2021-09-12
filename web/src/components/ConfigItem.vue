<template>
  <div class="my-2 w-full rounded-lg p-2 border-gray-100 border-1 bg-gray-700">
    <div
      @click="doExpand"
      @keydown.enter="doExpand"
      @keydown.space="doExpand"
      role="button"
      tabindex="0"
      :aria-label="ariaLabel"
      :aria-disabled="!collapsible"
      :class="collapsible ? 'cursor-pointer' : ''"
      class="
        text-3xl
        w-full
        select-none
        border-gray-100 border-b-2
        focus:outline-white focus:opacity-80
        flex flex-row
        justify-between
      "
    >
      <div>
        {{ title }}
      </div>
      <div v-if="collapsible">
        <o-icon v-if="expanded" icon="chevron-up" />
        <o-icon v-else icon="chevron-down" />
      </div>
    </div>
    <transition name="fade">
      <slot v-if="!collapsible || expanded" />
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

<style scoped></style>

<script>
import ConfirmDelete from "./ConfirmDelete.vue";
import TabSaver from "../util/TabSaver";

export default {
  data() {
    return {
      showConfirmDeleteModal: false,
      expanded: false,
    };
  },
  props: {
    title: {
      type: String,
      default: function () {
        return "";
      },
    },
    collapsible: {
      type: Boolean,
      default: function () {
        return false;
      },
    },
    deleteMessage: {
      type: String,
      default: function () {
        return "";
      },
    },
  },
  components: { ConfirmDelete },
  emits: ["delete", "edit"],
  mixins: [TabSaver],
  methods: {
    closeDelete() {
      this.showConfirmDeleteModal = false;
      this.restoreFocus();
    },
    confirmDelete($event) {
      this.lastButton = $event.currentTarget;
      this.showConfirmDeleteModal = true;
    },
    doExpand() {
      if (this.collapsible) {
        this.expanded = !this.expanded;
      }
    },
  },
  computed: {
    ariaLabel() {
      return this.expanded ? "Collapse" : "Expand";
    },
  },
};
</script>
