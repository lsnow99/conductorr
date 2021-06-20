<template>
  <header class="modal-card-header">
    <p class="modal-card-title">New Profile</p>
    <o-icon
      clickable
      native-type="button"
      icon="times"
      @click="$emit('close')"
    />
  </header>
  <section class="modal-card-content">
    <o-field label="Name">
      <o-input v-model="name" />
    </o-field>
  </section>
  <footer class="modal-card-footer">
    <o-button @click="$emit('close')">Cancel</o-button>
    <div>
      <o-button @click="submit"><action-button :mode="loading?'loading':''">Submit</action-button></o-button>
    </div>
  </footer>
</template>

<script>
import APIUtil from "../util/APIUtil"
import ActionButton from './ActionButton.vue';

export default {
  components: { ActionButton },
  data() {
    return {
      name: "",
      loading: false,
    };
  },
  emits: ['close', 'submitted'],
  methods: {
    submit() {
      this.loading = true
      APIUtil.createNewProfile(this.name).then(() => {
        this.$store.commit("addToast", {
          type: "success",
          msg: `Created profile ${this.name}`
        })
        this.$emit("submitted");
      }).finally(() => {
        this.loading = false;
      });
    },
  },
};
</script>
