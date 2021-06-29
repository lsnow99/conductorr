export default {
  props: {
    name: {
      type: String,
      default: function () {
        return "";
      },
    },
  },
  emits: ["update:name"],
  computed: {
    computedName: {
      get() {
        return this.name;
      },
      set(newVal) {
        console.log('updating name to ', newVal)
        this.$emit("update:name", newVal);
      },
    },
  },
};
