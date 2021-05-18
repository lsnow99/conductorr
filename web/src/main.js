import { createApp } from 'vue'
import App from './App.vue'

import '@oruga-ui/oruga-next/dist/oruga.css'
import './index.css'

/*
Import libraries
*/
import Oruga from '@oruga-ui/oruga-next'

import { library } from '@fortawesome/fontawesome-svg-core'
import { faSyncAlt } from '@fortawesome/free-solid-svg-icons'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'

/*
Import our router
*/
import router from './router'

library.add(faSyncAlt)

const app = createApp(App)
app.use(router)
app.use(Oruga, {
    iconPack: "fas",
    iconComponent: "vue-fontawesome",
    statusIcon: false
})
app.component('vue-fontawesome', FontAwesomeIcon)
app.mount('#app')
