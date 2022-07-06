<template>
  <o-field label="Auth Token">
    <o-input
      expanded
      type="text"
      v-model="computedValue"
      placeholder="Auth Token"
    />
    <o-button @click="showFetchAuthTokenModal = true">Fetch New</o-button>
  </o-field>
  <modal
    v-model="showFetchAuthTokenModal"
    @close="showFetchAuthTokenModal = false"
    title="Fetch Plex Auth Token"
  >
    <div>
      <o-field label="Username">
        <o-input type="text" v-model="plexUsername" placeholder="Username" />
      </o-field>
      <o-field label="Password">
        <o-input
          type="password"
          v-model="plexPassword"
          placeholder="Password"
        />
      </o-field>
    </div>
    <template v-slot:footer>
      <o-button @click="showFetchAuthTokenModal = false">Cancel</o-button>
      <o-button variant="primary" @click="fetchPlexAuthToken">
        <action-button :mode="fetchTestingMode"> Fetch </action-button>
      </o-button>
    </template>
  </modal>
</template>

<script setup lang="ts">
import { useComputedValue } from "@/util";
import { inject, ref } from "vue";
import APIUtil from "../util/APIUtil";
import ActionButton from "./ActionButton.vue";
import Modal from "./Modal.vue";

const oruga = inject("oruga");
const plexUsername = ref("");
const plexPassword = ref("");
const fetchTestingMode = ref("");
const showFetchAuthTokenModal = ref(false);

const props = defineProps<{
  modelValue: string;
}>();

const emit = defineEmits<{
  (e: "update:modelValue"): void;
}>();

const computedValue = useComputedValue<string>(props, emit);

const fetchPlexAuthToken = async () => {
  fetchTestingMode.value = "loading";
  try {
    const response = await APIUtil.getPlexAuthToken(
      plexUsername.value,
      plexPassword.value
    );
    computedValue.value = response.token;
    showFetchAuthTokenModal.value = false;
    fetchTestingMode.value = "";
    oruga.notification.open({
      duration: 5000,
      message: `Successfully fetched Plex auth token`,
      position: "bottom-right",
      variant: "success",
      closable: false,
    });
  } catch (err) {
    fetchTestingMode.value = "danger";
    oruga.notification.open({
      duration: 5000,
      message: `Plex login failed: ${err}`,
      position: "bottom-right",
      variant: "danger",
      closable: false,
    });
  }
};
</script>
