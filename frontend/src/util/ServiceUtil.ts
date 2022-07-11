import { computed, Ref, ref, WritableComputedRef } from "vue";
import { ConfigurableService, EditServiceMode } from "@/types/api/service";

export default function <T>(
  lastButton: Ref<HTMLElement | null>,
  restoreFocus: () => void,
  newServiceCallback: (service: T) => Promise<void>,
  editServiceCallback: (service: T) => Promise<void>
) {
  const editingService = ref<T | null>(null) as Ref<T | null>;
  const mode = ref<EditServiceMode>("");

  const showNewServiceModal: WritableComputedRef<boolean> = computed({
    get() {
      return mode.value === "new";
    },
    set(v) {
      if (v) {
        mode.value = "new";
      } else {
        mode.value = "";
      }
    },
  });

  const showEditServiceModal: WritableComputedRef<boolean> = computed({
    get() {
      return mode.value === "edit";
    },
    set(v) {
      if (v) {
        mode.value = "edit";
      } else {
        mode.value = "";
      }
    },
  });

  const closeModal = () => {
    mode.value = "";
    editingService.value = null;
    restoreFocus();
  };

  const openNewServiceModal = ($event: Event) => {
    lastButton.value = $event.currentTarget as HTMLElement;
    showNewServiceModal.value = true;
  };

  const openEditServiceModal = ($event: Event) => {
    lastButton.value = $event.currentTarget as HTMLElement;
    showEditServiceModal.value = true;
  };

  const editService = (service: T, $event: Event) => {
    editingService.value = service;
    openEditServiceModal($event);
  };

  const onSubmit = async(service: ConfigurableService) => {
    try {
        if (mode.value === "new") {
          await newServiceCallback(service as T);
          showNewServiceModal.value = false
        } else if (mode.value === "edit") {
          await editServiceCallback(service as T);
          showEditServiceModal.value = false
        }
    } catch {

    }
  };

  return {
    showNewServiceModal,
    showEditServiceModal,
    closeModal,
    openNewServiceModal,
    openEditServiceModal,
    editingService,
    editService,
    onSubmit,
  };
}
