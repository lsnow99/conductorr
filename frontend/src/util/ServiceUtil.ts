import { computed, Ref, ref, WritableComputedRef } from "vue";
import { ConfigurableService, EditServiceMode } from "@/types/api/service";

export default function <T extends ConfigurableService>(
  lastButton: Ref<HTMLElement | null>,
  restoreFocus: () => void,
  newServiceCallback: (service: T) => Promise<void>,
  editServiceCallback: (service: T) => Promise<void>
) {
  const editingService = ref<T | null>(null) as Ref<T | null>;
  const mode = ref<EditServiceMode>("");

  const closeModal = () => {
    mode.value = "";
    editingService.value = null;
    restoreFocus();
  };

  const showModal = computed(() => mode.value !== "")

  const openNewServiceModal = ($event?: Event) => {
    lastButton.value = $event?.currentTarget as HTMLElement | undefined ?? null;
    mode.value = "new"
  };

  const openEditServiceModal = ($event?: Event) => {
    lastButton.value = $event?.currentTarget as HTMLElement | undefined ?? null;
    mode.value = "edit"
  };

  const editService = (service: T, $event?: Event) => {
    editingService.value = service;
    openEditServiceModal($event);
  };

  const onSubmit = async(service: ConfigurableService) => {
    try {
        if (mode.value === "new") {
          await newServiceCallback(service as T);
          console.log('done')
          mode.value = ""
        } else if (mode.value === "edit") {
          await editServiceCallback(service as T);
          mode.value = ""
        }
    } catch {

    }
  };

  return {
    showModal,
    closeModal,
    openNewServiceModal,
    editingService,
    editService,
    onSubmit,
    mode
  };
}
