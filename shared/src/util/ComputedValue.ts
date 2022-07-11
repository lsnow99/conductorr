import { computed, type WritableComputedRef } from "vue";


type ComputedValueProps<T> = Readonly<{modelValue: T}>
type ComputedValueEmit<T> = { (e: "update:modelValue", newValue: T): void}

export default function <T>(
  props: ComputedValueProps<T>,
  emit: ComputedValueEmit<T>
) {
  const computedValue: WritableComputedRef<T> = computed({
    get(): T {
      return props.modelValue;
    },
    set(newValue: T): void {
      emit("update:modelValue", newValue);
    },
  });

  return computedValue
}
