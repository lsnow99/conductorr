import { nextTick, ref } from "vue";

export default () => {
  const lastButton = ref<HTMLElement | null>(null);
  
  const restoreFocus = () => {
    nextTick(() => {
      if(lastButton.value) {
        lastButton.value.focus()
      }
    })
  }

  return { restoreFocus, lastButton }
};
