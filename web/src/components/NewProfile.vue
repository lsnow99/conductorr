<template>
  <modal title="New Profile" v-model="computedActive" @close="$emit('close')">
    <o-field label="Name">
      <o-input @keydown.enter="submit" v-model="name" />
    </o-field>
    <template v-slot:footer>
      <o-button @click="$emit('close')">Cancel</o-button>
      <div>
        <o-button variant="primary" @click="submit"
          ><action-button :mode="loading ? 'loading' : ''"
            >Submit</action-button
          ></o-button
        >
      </div>
    </template>
  </modal>
</template>

<script>
import APIUtil from "../util/APIUtil";
import ActionButton from "./ActionButton.vue";
import Modal from "./Modal.vue";

export default {
  components: { ActionButton, Modal },
  data() {
    return {
      name: "",
      loading: false,
    };
  },
  props: {
    active: {
      type: Boolean,
      default: function() {
        return false
      }
    }
  },
  emits: ["close", "submitted", "update:active"],
  methods: {
    submit() {
      this.loading = true;
      APIUtil.createNewProfile(this.name)
        .then(() => {
          this.$oruga.notification.open({
            duration: 3000,
            message: `Created profile ${this.name}`,
            position: "bottom-right",
            variant: "success",
            closable: false,
          });
          this.$emit("submitted");
        })
        .finally(() => {
          this.loading = false;
        });
    },
  },
  computed: {
    computedActive: {
      get() {
        return this.active
      },
      set(v) {
        this.$emit('update:active', v)
      }
    }
  }
};
</script>
