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
      <o-button @click="submit">Submit</o-button>
    </div>
  </footer>
</template>

<script>
import APIUtil from "../util/APIUtil"

export default {
  data() {
    return {
      name: "",
    };
  },
  methods: {
    submit() {
      APIUtil.createNewProfile(this.name).then(() => {
        this.$emit("submitted");
      });
    },
  },
};
</script>
