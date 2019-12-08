import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../views/Home.vue'
import About from '@/views/About.vue'
import Configuration from '@/views/Configuration.vue'
import Overview from '@/views/Overview.vue'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'home',
    component: Home
  },
  {
    path: '/about',
    name: 'about',
    component: About
  },
  {
    path: '/configuration',
    name: 'configuration',
    component: Configuration
  },
  {
    path: '/overview',
    name: 'overview',
    component: Overview
  }
]

const router = new VueRouter({
  routes
})

export default router
