import { nextTick } from "vue";

export default {
  data() {
    return {
      tooltipActive: true,
    };
  },
  methods: {
    resetTooltips() {
      this.tooltipActive = false;
      nextTick(() => {
        this.tooltipActive = true;
      });
    },
  },
};
