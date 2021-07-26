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
  faArrowUp,
  faArrowDown,
  faCheck,
  faCogs,
  faTrash,
  faEdit,
  faBolt,
  faDice,
  faWrench,
  faChevronDown,
  faChevronUp,
  faCloudDownloadAlt,
  faBookmark,
  faCircleNotch,
  faSort,
} from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";

import { faGitAlt, faImdb } from "@fortawesome/free-brands-svg-icons";

import { faBookmark as faBookmarkRegular } from "@fortawesome/free-regular-svg-icons";

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
import { PrismEditor } from "vue-prism-editor";
import "vue-prism-editor/dist/prismeditor.min.css";

/*
Import VueX devtools fix from https://github.com/erdemefe07/vuex4-devtools-support
*/
import { addDevtools } from "./vuexdev";

const globalOptions = {
  debug: "info",
  modules: {
    syntax: true,
    toolbar: [["code-block"]],
  },
  readOnly: false,
  theme: "snow",
};

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
library.add(faChevronDown);
library.add(faChevronUp);
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
library.add(faArrowUp);
library.add(faArrowDown);
library.add(faCheck);
library.add(faCogs);
library.add(faTrash);
library.add(faEdit);
library.add(faBolt);
library.add(faDice);
library.add(faWrench);
library.add(faCloudDownloadAlt);
library.add(faBookmark);
library.add(faBookmarkRegular);
library.add(faCircleNotch);
library.add(faSort);

const app = createApp(App);
app.use(store);
app.use(router);
app.use(Oruga, {
  iconPack: "fas",
  iconComponent: "vue-fontawesome",
  statusIcon: false,
});
app.component("vue-fontawesome", FontAwesomeIcon);
app.component("PrismEditor", PrismEditor);
app.mount("#app");

if (import.meta.env.DEV) {
  addDevtools(app, store);
}
