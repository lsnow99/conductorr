import { nextTick, ref } from "vue";

export default function() {
    const tooltipActive = ref(true)

    const resetTooltips = () => {
        tooltipActive.value = false;
        nextTick(() => {
            tooltipActive.value = true
        })
    }

    return {
        tooltipActive,
        resetTooltips
    }
}
