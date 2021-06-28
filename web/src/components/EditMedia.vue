<template>
    <header class="modal-card-header">
    <p class="modal-card-title">{{ media.title }}</p>
  </header>
  <section class="modal-card-content w-96 min-w-full">
    
  </section>
  <footer class="modal-card-footer">
    <o-button @click="$emit('close')">Cancel</o-button>
    <div>
      <o-button variant="primary" @click="submit">Save</o-button>
    </div>
  </footer>
</template>

<script>
export default {
    props: {
        media: {
            type: Object,
            default: function() {
                return {}
            }
        }
    },
    emits: ['close', 'submit']
}
</script>