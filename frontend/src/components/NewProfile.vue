<template>
  <modal title="New Profile" v-model="computedActive" @close="$emit('close')">
    <o-field label="Name">
      <o-input
        @keydown.enter="submit"
        v-model="name"
        placeholder="Profile name"
      />
    </o-field>
    <template v-slot:footer>
      <o-button @click="$emit('close')">Cancel</o-button>
      <div>
        <o-button variant="primary" @click="submit"
          ><action-button :mode="loading ? 'loading' : ''"
            >Submit</action-button
          ></o-button
        >
      </div>
    </template>
  </modal>
</template>

<script setup lang="ts">
import APIUtil from "../util/APIUtil";
import ActionButton from "./ActionButton.vue";
import Modal from "./Modal.vue";
import useComputedActive from "@/util/ComputedActive"
import { inject, ref } from "vue";

const name = ref("");
const loading = ref(false);
const oruga = inject("oruga")

const props = defineProps<{
  active: boolean
}>();

const emit = defineEmits<{
  (e: "close"): void
  (e: "submitted"): void
  (e: "update:active", newValue: boolean): void
}>()

const { computedActive } = useComputedActive(props, emit);

const sanitize = () => {
  name.value = name.value ? name.value.trim() : "";
}
const validate = () => {
  if (!name.value) {
    return "Name is required";
  }
}

const submit = async() => {
  sanitize();
  const validationErr = validate();
  if(validationErr) {
    oruga.notification.open({
      duration: 5000,
      message: validationErr,
      position: "bottom-right",
      variant: "danger",
      closable: false,
    })
    return;
  }
  loading.value = true
  try {
    await APIUtil.createNewProfile(name.value)
    oruga.notification.open({
      duration: 3000,
      message: `Created profile ${name.value}`,
      position: "bottom-right",
      variant: "success",
      closable: false,
    });
    emit("submitted")
  } finally {
    loading.value = false
  }
}
</script>
