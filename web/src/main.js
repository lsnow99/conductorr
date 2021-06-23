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
  faFilter,
  faPlusSquare,
  faMinusSquare,
  faPlus,
  faBook,
  faAngleRight,
  faAngleLeft,
  faTimesCircle,
  faHeart,
  faQuestionCircle,
  faExternalLinkAlt,
  faCaretUp,
  faCaretDown,
  faPlay,
  faCheckCircle,
  faExclamationCircle,
  faExclamationTriangle,
  faInfoCircle,
  faAngleDoubleDown,
  faStar,
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

/*
Import Vue Prism
*/
import { PrismEditor } from 'vue-prism-editor';
import 'vue-prism-editor/dist/prismeditor.min.css';

const globalOptions = {
  debug: 'info',
  modules: {
    syntax: true,
    toolbar: [['code-block']]
  },
  readOnly: false,
  theme: 'snow'
}

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
library.add(faFilter);
library.add(faPlusSquare);
library.add(faMinusSquare);
library.add(faPlus);
library.add(faBook);
library.add(faAngleRight);
library.add(faAngleLeft);
library.add(faTimesCircle);
library.add(faHeart);
library.add(faGitAlt);
library.add(faQuestionCircle);
library.add(faExternalLinkAlt);
library.add(faCaretUp);
library.add(faCaretDown);
library.add(faPlay);
library.add(faCheckCircle);
library.add(faExclamationCircle);
library.add(faExclamationTriangle);
library.add(faInfoCircle);
library.add(faAngleDoubleDown);
library.add(faImdb);
library.add(faStar);

const app = createApp(App);
app.use(store);
app.use(router);
app.use(Oruga, {
  iconPack: "fas",
  iconComponent: "vue-fontawesome",
  statusIcon: false,
});
app.component("vue-fontawesome", FontAwesomeIcon);
app.component("PrismEditor", PrismEditor)
app.mount("#app");

/*
Import VueX devtools fix from https://github.com/erdemefe07/vuex4-devtools-support
*/
import { addDevtools } from "./vuexdev";
import { faGitAlt, faImdb } from "@fortawesome/free-brands-svg-icons";
addDevtools(app, store);
