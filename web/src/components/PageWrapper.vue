<template>
  <section class="container mx-auto mt-4">
    <section class="menu" :class="expandMenu ? 'expanded' : ''">
      <div class="visible xl:hidden pl-5 pr-5 pt-6 pb-6 text-xl">
        <div
          class="
            border-solid border-4
            rounded-md
            cursor-pointer
            inline-block
            pt-1
            pb-1
            pl-3
            pr-3
          "
          @click="expandMenu = !expandMenu"
        >
          <vue-fontawesome icon="bars" class="text-white" />
        </div>
      </div>
      <router-link
        v-for="route in routeTree"
        :key="route.name"
        class="route-item reg-route hidden xl:visible"
        :class="curRouteName === route.name ? 'cur-route' : ''"
        :to="{ name: route.name }"
      >
        <vue-fontawesome :icon="route.icon" />
        {{ route.title }}
      </router-link>
      <router-link
        :to="{ name: 'logout' }"
        class="route-item ml-auto logout align-middle flex items-center"
      >
        <vue-fontawesome icon="sign-out-alt" />
        Logout
      </router-link>
    </section>
    <section :class="class" class="mt-10">
      <slot />
    </section>
  </section>
</template>

<style scoped>
.route-item {
  @apply px-5 py-4 text-xl cursor-pointer hover:bg-yellow-400 focus:bg-yellow-400 uppercase text-white;
}

.reg-route {
  @apply hidden;
  @apply xl:block;
}

.expanded > .reg-route {
  @apply block;
}

.menu {
  @apply rounded-md flex flex-row xl:flex-row bg-gray-600 overflow-hidden sm:m-0 m-2;
}

.menu.expanded {
  @apply flex-col;
  @apply xl:flex-row;
}

.expanded > .logout {
  @apply ml-0;
  @apply xl:ml-auto;
}

.cur-route {
  @apply bg-yellow-500;
}
</style>

<script>
export default {
  data() {
    return {
      expandMenu: false,
      routeTree: [
        {
          title: "Home",
          name: "home",
          icon: "home",
        },
        {
          title: "Movies",
          name: "movies",
          icon: "film",
        },
        {
          title: "TV Shows",
          name: "tv",
          icon: "tv",
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
      ],
    };
  },
  props: {
    class: {
      type: String,
      default: function () {
        return "";
      },
    },
  },
  computed: {
    curRouteName() {
      if (this.$route.meta.navName) {
        return this.$route.meta.navName;
      }
      return this.$route.name
    },
  },
};
</script>
