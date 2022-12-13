<template>
  <div class="flex flex-col min-h-screen">
    <nav class="container pb-4 mx-auto mt-4">
      <div class="cursor-auto menu" :class="expandMenu ? 'expanded' : ''">
        <div
          class="visible w-full px-3 py-2 text-xl lg:hidden focus:outline-white"
          @click="toggleMenu"
          @keydown.enter="toggleMenu"
          @keydown.space="toggleMenu"
          role="button"
          tabindex="0"
          aria-label="Expand navigation menu"
        >
          <div class="inline-block px-2 py-px rounded-md cursor-pointer">
            <vue-fontawesome icon="bars" class="text-white" />
          </div>
        </div>
        <router-link
          v-for="route in routeTree"
          :key="route.name"
          class="hidden route-item reg-route lg:visible"
          :class="curRouteName === route.name ? 'cur-route' : ''"
          :to="{ name: route.name }"
        >
          <vue-fontawesome :icon="route.icon" />
          <span class="ml-2">{{ route.title }}</span>
        </router-link>
        <router-link
          :to="{ name: 'logout' }"
          class="flex items-center ml-auto route-item logout"
        >
          <vue-fontawesome icon="sign-out-alt" />
          <span class="ml-1">Logout</span>
        </router-link>
      </div>
    </nav>
    <main :class="class" class="container mx-auto mt-4">
      <slot />
    </main>
    <footer class="mt-auto">
      <div
        class="flex flex-row w-64 items-center justify-center mx-auto rounded-t-md h-14"
      >
        <o-tooltip variant="info">
          <a
            href="javascript:void()"
            class="text-gray-500 hover:text-red-500"
            aria-label="Donate"
          >
            <vue-fontawesome class="footer-icon" icon="heart" />
          </a>
          <template v-slot:content
            >Donate <vue-fontawesome class="ml-1" icon="external-link-alt"
          /></template>
        </o-tooltip>
        <o-tooltip variant="info">
          <a
            href="https://github.com/lsnow99/conductorr"
            target="_blank"
            class="text-gray-500 hover:text-blue-500"
            rel="noreferrer"
            aria-label="Source code"
          >
            <vue-fontawesome class="footer-icon" :icon="['fab', 'git-alt']" />
          </a>
          <template v-slot:content
            >Source Code <vue-fontawesome class="ml-1" icon="external-link-alt"
          /></template>
        </o-tooltip>
        <o-tooltip variant="info">
          <a
            href="https://github.com/lsnow99/conductorr/issues"
            target="_blank"
            class="text-gray-500 hover:text-green-500"
            rel="noreferrer"
            aria-label="Support"
          >
            <vue-fontawesome class="footer-icon" icon="question-circle" />
          </a>
          <template v-slot:content
            >Support <vue-fontawesome class="ml-1" icon="external-link-alt"
          /></template>
        </o-tooltip>
      </div>
    </footer>
  </div>
</template>

<style scoped>
.footer-icon {
  @apply text-2xl;
  @apply mx-3;
}

.route-item {
  @apply px-3 py-1 text-xl cursor-pointer hover:bg-yellow-400 focus:bg-yellow-400 uppercase text-white;
}

.reg-route {
  @apply hidden;
  @apply lg:block;
}

.expanded > .reg-route {
  @apply block;
}

.menu {
  @apply rounded-md flex flex-row lg:flex-row bg-gray-700 overflow-hidden sm:m-0 m-2;
}

.menu.expanded {
  @apply flex-col;
  @apply lg:flex-row;
}

.expanded > .logout {
  @apply ml-0;
  @apply lg:ml-auto;
}

.cur-route {
  @apply bg-yellow-500;
}
</style>

<script setup lang="ts">
import { computed, ref } from "vue";
import { useRoute } from "vue-router";

const expandMenu = ref(false);
const routeTree = ref([
  {
    title: "Home",
    name: "home",
    icon: "home",
  },
  {
    title: "Library",
    name: "library",
    icon: "book",
  },
  {
    title: "Calendar",
    name: "calendar",
    icon: "calendar-alt",
  },
  {
    title: "Configuration",
    name: "configuration",
    icon: "sliders-h",
  },
  {
    title: "System",
    name: "system",
    icon: "server",
  },
]);

const props = defineProps<{
  class?: string;
}>();

const toggleMenu = ($event?: Event) => {
  if ($event) {
    $event.preventDefault();
  }
  expandMenu.value = !expandMenu.value;
};

const route = useRoute();
const curRouteName = computed(() => {
  return route.meta.navName ?? route.name;
});
</script>
