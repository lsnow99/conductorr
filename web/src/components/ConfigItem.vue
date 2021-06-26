<template>
  <div class="my-2 w-full rounded-lg p-2 border-gray-100 border-1 bg-gray-700">
    <div
      class="
        text-3xl
        w-full
        cursor-pointer
        select-none
        border-gray-100 border-b-2
      "
    >
      {{ title }}
    </div>
    <slot />
    <div class="flex flex-row justify-between mt-4">
      <o-button variant="danger" @click="$emit('delete')">Delete</o-button>
      <o-button variant="primary" @click="$emit('edit')">Edit</o-button>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    title: {
      type: String,
      default: function () {
        return "";
      },
    },
  },
  emits: ["delete", "edit"],
};
</script>
