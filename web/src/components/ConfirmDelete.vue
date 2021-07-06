<template>
  <header class="modal-card-header">
    <p class="modal-card-title">Confirm Delete</p>
  </header>
  <section class="modal-card-content">
    <p class="text-2xl">{{ deleteMessage }}</p>
  </section>
  <footer class="modal-card-footer">
    <o-button @click="$emit('close')">Cancel</o-button>
    <o-button @click="$emit('delete')" variant="danger">Delete</o-button>
  </footer>
</template>

<script>
export default {
    props: {
        deleteMessage: {
            type: String,
            default: function() {
                return "You sure?"
            }
        }
    }
}
</script>