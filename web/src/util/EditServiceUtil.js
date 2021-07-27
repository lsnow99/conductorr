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
        this.$emit("update:name", newVal);
      },
    },
  },
  created() {
    console.log("HELLLO")
  }
};
