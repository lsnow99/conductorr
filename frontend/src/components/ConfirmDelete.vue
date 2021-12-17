<template>
  <modal title="Confirm Delete" v-model="computedActive">
    <p class="text-2xl">{{ deleteMessage }}</p>
  <template v-slot:footer>
    <o-button @click="$emit('close')">Cancel</o-button>
    <o-button @click="$emit('delete')" variant="danger">Delete</o-button>
  </template>
  </modal>
</template>

<script>
import Modal from "./Modal.vue"

export default {
  props: {
    deleteMessage: {
      type: String,
      default: function () {
        return "You sure?";
      },
    },
    active: {
      type: Boolean,
      default: function () {
        return false;
      },
    },
  },
  components: {
    Modal,
  },
  computed: {
    computedActive: {
      get() {
        return this.active;
      },
      set(v) {
        this.$emit("update:active", v);
      },
    },
  },
};
</script>
