import { createRouter, createWebHistory } from "vue-router";
import Home from "./views/Home.vue";
import Auth from "./views/Auth.vue";
import TV from "./views/TV.vue";
import Movies from "./views/Movies.vue";
import Configuration from "./views/Configuration.vue";
import Calendar from "./views/Calendar.vue";
import System from "./views/System.vue";
import AuthUtil from "./util/AuthUtil.js";
import { faTachometerAlt } from "@fortawesome/free-solid-svg-icons";

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
    path: "/movies",
    name: 'movies',
    component: Movies,
    meta: {
      requiresAuth: true,
    },
  },
  {
    path: "/tv",
    name: 'tv',
    component: TV,
    meta: {
      requiresAuth: true,
    },
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
    path: "/configuration",
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
    meta: {
      logout: true
    }
  }
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

router.beforeEach((to, from, next) => {
  /*
      If navigating to a special route with the meta.logout property set,
      we log out the user and bring them to the sign in page.
      */
  if (to.matched.some((record) => record.meta.logout)) {
    // TODO: Do logout
    console.log('logging out')
    AuthUtil.logout();
    next({
      name: "auth"
    })
    return
  }
  /*
      If a path requires authentication, we first check that the user is logged in
      and if the check fails, send them to the login page.
      TODO: implement
      */
  if (to.matched.some((record) => record.meta.requiresAuth)) {
    // Only allow the requested route if the logged in check passes
    AuthUtil.getLoggedInID()
      .then(() => {
        next()
      }).catch(() => {
        AuthUtil.logout();
        next({
          name: 'auth'
        })
      })
  } else {
    next();
  }
});

export default router;
