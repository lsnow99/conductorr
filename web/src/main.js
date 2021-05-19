import { createApp } from "vue";
import App from "./App.vue";

import "@oruga-ui/oruga-next/dist/oruga.css";
import "./index.css";

/*
Import third party libraries
*/
import Oruga from "@oruga-ui/oruga-next";

import { library } from "@fortawesome/fontawesome-svg-core";
import {
  faServer,
  faHome,
  faSignOutAlt,
  faSyncAlt,
  faTimes,
  faFilm,
  faTv,
  faCalendarAlt,
  faBars,
  faChevronLeft,
  faChevronRight,
  faSearch,
  faDownload,
  faSlidersH,
  faPlayCircle,
} from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";

/*
Import our router
*/
import router from "./router";

/*
Import our store
*/
import store from "./store";

library.add(faSyncAlt);
library.add(faTimes);
library.add(faHome);
library.add(faFilm);
library.add(faTv);
library.add(faCalendarAlt);
library.add(faSignOutAlt);
library.add(faServer);
library.add(faBars);
library.add(faChevronLeft);
library.add(faChevronRight);
library.add(faSearch);
library.add(faDownload);
library.add(faSlidersH);
library.add(faPlayCircle);

const app = createApp(App);
app.use(store);
app.use(router);
app.use(Oruga, {
  iconPack: "fas",
  iconComponent: "vue-fontawesome",
  statusIcon: false,
});
app.component("vue-fontawesome", FontAwesomeIcon);
app.mount("#app");

/*
Import VueX devtools fix from https://github.com/erdemefe07/vuex4-devtools-support
*/
import { addDevtools } from "./vuexdev";
addDevtools(app, store);
