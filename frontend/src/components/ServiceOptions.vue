<template>
    <div class="grid md:grid-cols-2">
      <div
        v-for="service in services"
        :key="service.value"
        class="service"
        :class="computedValue == service.value ? 'active' : ''"
        @click="computedValue = service.value"
        @keydown.space="computedValue = service.value"
        @keydown.enter="computedValue = service.value"
        tabindex="0"
        role="button"
        :aria-label="service.name"
      >
        {{ service.name }}
      </div>
    </div>
</template>

<style scoped>
.service {
  @apply border-dashed border-4 border-gray-300 py-8 px-16 m-2 text-2xl cursor-pointer hover:opacity-70 flex justify-center font-bold select-none;
}

.service.active {
  @apply border-green-500 text-green-500 border-solid;
}
</style>

<script>
export default {
    props: {
        modelValue: {
            type: String,
            default: function() {
                return undefined
            }
        },
        services: {
            type: Array,
            default: function() {
                return []
            }
        }
    },
    emits: ['update:modelValue'],
    computed: {
        computedValue: {
            get() {
                return this.modelValue
            },
            set(newVal) {
                this.$emit('update:modelValue', newVal)
            }
        }
    }
}
</script>