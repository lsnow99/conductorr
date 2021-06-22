<template>
  <header class="modal-card-header">
    <p class="modal-card-title">New Profile</p>
  </header>
  <section class="modal-card-content">
    <o-field label="Name">
      <o-input @keydown.enter="submit" v-model="name" />
    </o-field>
  </section>
  <footer class="modal-card-footer">
    <o-button @click="$emit('close')">Cancel</o-button>
    <div>
      <o-button variant="primary" @click="submit"
        ><action-button :mode="loading ? 'loading' : ''"
          >Submit</action-button
        ></o-button
      >
    </div>
  </footer>
</template>

<script>
import APIUtil from "../util/APIUtil";
import ActionButton from "./ActionButton.vue";

export default {
  components: { ActionButton },
  data() {
    return {
      name: "",
      loading: false,
    };
  },
  emits: ["close", "submitted"],
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
};
</script>
