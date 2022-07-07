import { computed, type WritableComputedRef } from "vue";
import ComputedValue from "./ComputedValue";

export type ActiveProps = Readonly<{
    active: boolean;
}>

export type ActiveEmit = {
    (e: "update:active", newValue: boolean): void
}

export default function (props: ActiveProps, emit: ActiveEmit) {

  const computedActive: WritableComputedRef<boolean> = computed({
    get(): boolean {
      return props.active;
    },
    set(newValue: boolean): void {
      emit("update:active", newValue);
    },
  });

  return computedActive
}
