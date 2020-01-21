import Vue from 'vue'
import VueRouter from 'vue-router'
import About from '@/views/About.vue'
import Configuration from '@/views/Configuration.vue'
import Overview from '@/views/Overview.vue'
import Authenticate from '@/views/Authenticate.vue'
import Statistics from '@/views/Statistics.vue'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'overview',
    component: Overview,
    meta: {
      requiresAuth: true
    }
  },
  {
    path: '/about',
    name: 'about',
    component: About
  },
  {
    path: '/configuration',
    name: 'configuration',
    component: Configuration,
    meta: {
      requiresAuth: true
    }
  },
  {
    path: '/statistics',
    name: 'statistics',
    component: Statistics,
    meta: {
      requiresAuth: true
    }
  },
  {
    path: '/auth',
    name: 'auth',
    component: Authenticate
  }
]

const router = new VueRouter({
  routes
})

router.beforeEach((to, from, next) => {
  if(to.matched.some(record => record.meta.requiresAuth)) {
    if (localStorage.getItem('jwt') == null) {
      next({
          path: '/auth',
          params: { nextUrl: to.fullPath }
      })
    } else {
      next()
    }
  } else {
    next()
  }
})

export default router
