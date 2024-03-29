import { createRouter, createWebHistory } from "vue-router";
import Home from "./views/Home.vue";
import Auth from "./views/Auth.vue";
import Configuration from "./views/Configuration.vue";
import EditProfile from "./views/EditProfile.vue";
import Library from "./views/Library.vue";
import Media from "./views/Media.vue";
import Calendar from "./views/Calendar.vue";
import System from "./views/System.vue";
import Logout from "./views/Logout.vue";
import NotFound from "./views/NotFound.vue";
import AuthUtil from "./util/AuthUtil.js";
import { useAppStore } from "@/store"

const routes = [
  {
    path: "/",
    name: 'home',
    component: Home,
    meta: {
      requiresAuth: true,
    },
  },
  {
    path: "/auth",
    name: 'auth',
    component: Auth,
    meta: {
      requiresAuth: false,
    },
  },
  {
    path: "/library",
    name: 'library',
    component: Library,
    meta: {
      requiresAuth: true,
    },
  },
  {
    path: "/media/:media_id",
    name: 'media',
    component: Media,
    meta: {
      requiresAuth: true,
      navName: 'library'
    }
  },
  {
    path: "/editProfile/:profile_id",
    name: 'editProfile',
    component: EditProfile,
    meta: {
      requiresAuth: true,
    }
  },
  {
    path: "/calendar",
    name: 'calendar',
    component: Calendar,
    meta: {
      requiresAuth: true,
    },
  },
  {
    path: "/configuration/:subpath?",
    name: 'configuration',
    component: Configuration,
    meta: {
      requiresAuth: true,
    },
  },
  {
    path: "/system",
    name: 'system',
    component: System,
    meta: {
      requiresAuth: true,
    }
  },
  {
    path: "/logout",
    name: 'logout',
    component: Logout,
    meta: {
      requiresAuth: false
    }
  },
  {
    path: '/:pathMatch(.*)*',
    name: "404",
    component: NotFound,
    meta: {
      requiresAuth: false
    }
  }
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

router.beforeEach((to, from, next) => {
  const appStore = useAppStore()
  if (to.matched.some((record) => record.meta.requiresAuth)) {
    // Only allow the requested route if the logged in check passes
    if (appStore.loggedIn) {
      next()
    } else {
      AuthUtil.logout();
      next({
        name: 'auth'
      })
    }
  } else {
    next();
  }
});

export default router;
