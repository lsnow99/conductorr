import Vue from 'vue'
import App from './App.vue'
import router from './router'
import Buefy from 'buefy';
import { library } from '@fortawesome/fontawesome-svg-core'
import { fas } from '@fortawesome/free-solid-svg-icons'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'
import VueClipboard from 'vue-clipboard2'
import { AuthService } from '@/mixins/AuthService.js'
import VueTimeago from 'vue-timeago'
 
Vue.use(VueTimeago, {
  name: 'Timeago', // Component name, `Timeago` by default
  locale: 'en', // Default locale
  // We use `date-fns` under the hood
  // So you can use all locales from it
  locales: {
    'zh-CN': require('date-fns/locale/zh_cn'),
    ja: require('date-fns/locale/ja')
  }
})
 
Vue.use(VueClipboard)

library.add(fas)
Vue.component('vue-fontawesome', FontAwesomeIcon)

Vue.use(Buefy, {
  defaultIconPack: 'fas',
  defaultIconComponent: 'vue-fontawesome',
});

Vue.config.productionTip = false

Vue.mixin(AuthService)

new Vue({
  router,
  render: h => h(App)
}).$mount('#app')
