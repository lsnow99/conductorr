import { nextTick } from "@vue/runtime-core";

export default {
  data() {
    return {
      lastButton: null,
    };
  },
  methods: {
    restoreFocus() {
      setTimeout(() => {
        if (this.lastButton) {
          console.log(this.lastButton);
          this.lastButton.focus();
        }
      }, 0);
    },
  },
};
