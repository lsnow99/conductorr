import { computed, type WritableComputedRef } from "vue";

export default function <T>(
  props: Readonly<{ modelValue: T }>,
  emit: { (e: "update:modelValue", newValue: T): void }
) {
  const computedValue: WritableComputedRef<T> = computed({
    get(): T {
      return props.modelValue;
    },
    set(newValue: T): void {
      emit("update:modelValue", newValue);
    },
  });

  return {
    computedValue,
  };
}
