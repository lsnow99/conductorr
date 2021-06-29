<template>
  <div class="my-2 w-full rounded-lg p-2 border-gray-100 border-1 bg-gray-700">
    <div
      @click="doExpand"
      :class="collapsible?'cursor-pointer':''"
      class="
        text-3xl
        w-full
        select-none
        border-gray-100 border-b-2
        flex flex-row justify-between
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
  <o-modal v-model:active="showConfirmDeleteModal" @close="showConfirmDeleteModal = false">
    <header class="modal-card-header">
    <p class="modal-card-title">Confirm Delete</p>
  </header>
  <section class="modal-card-content">
    <p class="text-2xl">{{deleteMessage}}</p>
  </section>
  <footer class="modal-card-footer">
    <o-button @click="showConfirmDeleteModal = false">Cancel</o-button>
    <o-button @click="$emit('delete')" variant="danger">Delete</o-button>
  </footer>
  </o-modal>
</template>

<script>
export default {
  data() {
    return {
      showConfirmDeleteModal: false,
      expanded: false
    }
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
      default: function() {
        return false
      }
    },
    deleteMessage: {
      type: String,
      default: function() {
        return ""
      },
    }
  },
  emits: ["delete", "edit"],
  methods: {
    confirmDelete() {
      this.showConfirmDeleteModal = true
    },
    doExpand() {
      if(this.collapsible) {
        this.expanded = !this.expanded
      }
    }
  }
};
</script>
