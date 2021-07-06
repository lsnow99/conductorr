<template>
  <div class="my-2 w-full rounded-lg p-2 border-gray-100 border-1 bg-gray-700">
    <div
      @click="doExpand"
      :class="collapsible ? 'cursor-pointer' : ''"
      class="
        text-3xl
        w-full
        select-none
        border-gray-100 border-b-2
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
      <o-button variant="primary" @click="$emit('edit')">Edit</o-button>
    </div>
  </div>
  <o-modal
    v-model:active="showConfirmDeleteModal"
    @close="showConfirmDeleteModal = false"
  >
    <ConfirmDelete
      @delete="$emit('delete')"
      @close="showConfirmDeleteModal = false"
      :delete-message="deleteMessage"
    />
  </o-modal>
</template>

<script>
import ConfirmDelete from "./ConfirmDelete.vue";

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
  methods: {
    confirmDelete() {
      this.showConfirmDeleteModal = true;
    },
    doExpand() {
      if (this.collapsible) {
        this.expanded = !this.expanded;
      }
    },
  },
};
</script>
