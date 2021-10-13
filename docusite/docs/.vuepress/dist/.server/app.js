"use strict";
var __defProp2 = Object.defineProperty;
var __defProps2 = Object.defineProperties;
var __getOwnPropDescs2 = Object.getOwnPropertyDescriptors;
var __getOwnPropSymbols2 = Object.getOwnPropertySymbols;
var __hasOwnProp2 = Object.prototype.hasOwnProperty;
var __propIsEnum2 = Object.prototype.propertyIsEnumerable;
var __defNormalProp2 = (obj, key, value) => key in obj ? __defProp2(obj, key, { enumerable: true, configurable: true, writable: true, value }) : obj[key] = value;
var __spreadValues2 = (a, b) => {
  for (var prop in b || (b = {}))
    if (__hasOwnProp2.call(b, prop))
      __defNormalProp2(a, prop, b[prop]);
  if (__getOwnPropSymbols2)
    for (var prop of __getOwnPropSymbols2(b)) {
      if (__propIsEnum2.call(b, prop))
        __defNormalProp2(a, prop, b[prop]);
    }
  return a;
};
var __spreadProps2 = (a, b) => __defProps2(a, __getOwnPropDescs2(b));
Object.defineProperty(exports, "__esModule", { value: true });
exports[Symbol.toStringTag] = "Module";
var vueRouter = require("vue-router");
var vue = require("vue");
var core = require("@vueuse/core");
var shared = require("@vuepress/shared");
var nprogress$1 = require("nprogress");
var serverRenderer = require("vue/server-renderer");
var luxon = require("luxon");
function _interopNamespace(e) {
  if (e && e.__esModule)
    return e;
  var n2 = { __proto__: null, [Symbol.toStringTag]: "Module" };
  if (e) {
    Object.keys(e).forEach(function(k) {
      if (k !== "default") {
        var d = Object.getOwnPropertyDescriptor(e, k);
        Object.defineProperty(n2, k, d.get ? d : {
          enumerable: true,
          get: function() {
            return e[k];
          }
        });
      }
    });
  }
  n2["default"] = e;
  return Object.freeze(n2);
}
var nprogress__namespace = /* @__PURE__ */ _interopNamespace(nprogress$1);
const ClientOnly = vue.defineComponent({
  setup(_, ctx) {
    const isMounted = vue.ref(false);
    vue.onMounted(() => {
      isMounted.value = true;
    });
    return () => {
      var _a, _b;
      return isMounted.value ? (_b = (_a = ctx.slots).default) === null || _b === void 0 ? void 0 : _b.call(_a) : null;
    };
  }
});
const pagesComponents = {
  "v-8daa1a0e": vue.defineAsyncComponent(() => Promise.resolve().then(function() {
    return index_html$2;
  })),
  "v-0c47f4f8": vue.defineAsyncComponent(() => Promise.resolve().then(function() {
    return examples_html$2;
  })),
  "v-4760ba53": vue.defineAsyncComponent(() => Promise.resolve().then(function() {
    return introduction_html$2;
  })),
  "v-dc33cfc4": vue.defineAsyncComponent(() => Promise.resolve().then(function() {
    return repl_html$2;
  })),
  "v-4f4ccb8f": vue.defineAsyncComponent(() => Promise.resolve().then(function() {
    return configuration_html$2;
  })),
  "v-62059836": vue.defineAsyncComponent(() => Promise.resolve().then(function() {
    return download_html$2;
  })),
  "v-8b3d4f7c": vue.defineAsyncComponent(() => Promise.resolve().then(function() {
    return calendar_html$2;
  })),
  "v-7e926318": vue.defineAsyncComponent(() => Promise.resolve().then(function() {
    return downloaders_html$2;
  })),
  "v-7440cc28": vue.defineAsyncComponent(() => Promise.resolve().then(function() {
    return indexers_html$2;
  })),
  "v-318bcd2c": vue.defineAsyncComponent(() => Promise.resolve().then(function() {
    return notifications_html$2;
  })),
  "v-0b1540b2": vue.defineAsyncComponent(() => Promise.resolve().then(function() {
    return postProcessing_html$2;
  })),
  "v-3bcd3316": vue.defineAsyncComponent(() => Promise.resolve().then(function() {
    return profiles_html$2;
  })),
  "v-5f90410b": vue.defineAsyncComponent(() => Promise.resolve().then(function() {
    return syncing_html$2;
  })),
  "v-3706649a": vue.defineAsyncComponent(() => Promise.resolve().then(function() {
    return _404_html$2;
  }))
};
const pagesData$1 = {
  "v-8daa1a0e": () => Promise.resolve().then(function() {
    return index_html;
  }).then(({ data: data3 }) => data3),
  "v-0c47f4f8": () => Promise.resolve().then(function() {
    return examples_html;
  }).then(({ data: data3 }) => data3),
  "v-4760ba53": () => Promise.resolve().then(function() {
    return introduction_html;
  }).then(({ data: data3 }) => data3),
  "v-dc33cfc4": () => Promise.resolve().then(function() {
    return repl_html;
  }).then(({ data: data3 }) => data3),
  "v-4f4ccb8f": () => Promise.resolve().then(function() {
    return configuration_html;
  }).then(({ data: data3 }) => data3),
  "v-62059836": () => Promise.resolve().then(function() {
    return download_html;
  }).then(({ data: data3 }) => data3),
  "v-8b3d4f7c": () => Promise.resolve().then(function() {
    return calendar_html;
  }).then(({ data: data3 }) => data3),
  "v-7e926318": () => Promise.resolve().then(function() {
    return downloaders_html;
  }).then(({ data: data3 }) => data3),
  "v-7440cc28": () => Promise.resolve().then(function() {
    return indexers_html;
  }).then(({ data: data3 }) => data3),
  "v-318bcd2c": () => Promise.resolve().then(function() {
    return notifications_html;
  }).then(({ data: data3 }) => data3),
  "v-0b1540b2": () => Promise.resolve().then(function() {
    return postProcessing_html;
  }).then(({ data: data3 }) => data3),
  "v-3bcd3316": () => Promise.resolve().then(function() {
    return profiles_html;
  }).then(({ data: data3 }) => data3),
  "v-5f90410b": () => Promise.resolve().then(function() {
    return syncing_html;
  }).then(({ data: data3 }) => data3),
  "v-3706649a": () => Promise.resolve().then(function() {
    return _404_html;
  }).then(({ data: data3 }) => data3)
};
const pagesData = vue.ref(pagesData$1);
const pageDataEmpty = vue.readonly({
  key: "",
  path: "",
  title: "",
  lang: "",
  frontmatter: {},
  excerpt: "",
  headers: []
});
const pageData = vue.ref(pageDataEmpty);
const usePageData = () => pageData;
const resolvePageData = async (pageKey) => {
  const pageDataResolver = pagesData.value[pageKey];
  if (!pageDataResolver) {
    return pageDataEmpty;
  }
  const pageData2 = await pageDataResolver();
  return pageData2 !== null && pageData2 !== void 0 ? pageData2 : pageDataEmpty;
};
if (false) {
  __VUE_HMR_RUNTIME__.updatePageData = (data3) => {
    pagesData.value[data3.key] = () => Promise.resolve(data3);
    if (data3.key === pageData.value.key) {
      pageData.value = data3;
    }
  };
}
const pageFrontmatterSymbol = Symbol("");
const usePageFrontmatter = () => {
  const pageFrontmatter = vue.inject(pageFrontmatterSymbol);
  if (!pageFrontmatter) {
    throw new Error("usePageFrontmatter() is called without provider.");
  }
  return pageFrontmatter;
};
const resolvePageFrontmatter = (pageData2) => pageData2.frontmatter;
const pageHeadSymbol = Symbol("");
const usePageHead = () => {
  const pageHead = vue.inject(pageHeadSymbol);
  if (!pageHead) {
    throw new Error("usePageHead() is called without provider.");
  }
  return pageHead;
};
const resolvePageHead = (headTitle, frontmatter, siteLocale) => {
  const description = shared.isString(frontmatter.description) ? frontmatter.description : siteLocale.description;
  const head = [
    ...shared.isArray(frontmatter.head) ? frontmatter.head : [],
    ...siteLocale.head,
    ["title", {}, headTitle],
    ["meta", { name: "description", content: description }]
  ];
  return shared.dedupeHead(head);
};
const pageHeadTitleSymbol = Symbol("");
const resolvePageHeadTitle = (page, siteLocale) => `${page.title ? `${page.title} | ` : ``}${siteLocale.title}`;
const pageLangSymbol = Symbol("");
const usePageLang = () => {
  const pageLang = vue.inject(pageLangSymbol);
  if (!pageLang) {
    throw new Error("usePageLang() is called without provider.");
  }
  return pageLang;
};
const resolvePageLang = (pageData2) => pageData2.lang || "en";
const routeLocaleSymbol = Symbol("");
const useRouteLocale = () => {
  const routeLocale = vue.inject(routeLocaleSymbol);
  if (!routeLocale) {
    throw new Error("useRouteLocale() is called without provider.");
  }
  return routeLocale;
};
const resolveRouteLocale = (locales, routePath) => shared.resolveLocalePath(locales, routePath);
const siteData$1 = {
  "base": "/",
  "lang": "en-US",
  "title": "Conductorr",
  "description": "Smart, fast, and powerful PVR for movies and TV",
  "head": [
    [
      "link",
      {
        "rel": "shortcut icon",
        "href": "/logo.svg"
      }
    ]
  ],
  "locales": {}
};
const siteData = vue.ref(siteData$1);
const useSiteData = () => siteData;
if (false) {
  __VUE_HMR_RUNTIME__.updateSiteData = (data3) => {
    siteData.value = data3;
  };
}
const siteLocaleDataSymbol = Symbol("");
const useSiteLocaleData = () => {
  const siteLocaleData = vue.inject(siteLocaleDataSymbol);
  if (!siteLocaleData) {
    throw new Error("useSiteLocaleData() is called without provider.");
  }
  return siteLocaleData;
};
const resolveSiteLocaleData = (site, routeLocale) => __spreadValues2(__spreadValues2({}, site), site.locales[routeLocale]);
const setupUpdateHead = () => {
  vueRouter.useRoute();
  const head = usePageHead();
  const lang = usePageLang();
  {
    const ssrContext = vue.useSSRContext();
    if (ssrContext) {
      ssrContext.head = head.value;
      ssrContext.lang = lang.value;
    }
    return;
  }
};
const Content = (props) => {
  let key;
  if (props.pageKey) {
    key = props.pageKey;
  } else {
    const page = usePageData();
    key = page.value.key;
  }
  const component = pagesComponents[key];
  if (component) {
    return vue.h(component);
  }
  return vue.h("div", "404 Not Found");
};
Content.displayName = "Content";
Content.props = {
  pageKey: {
    type: String,
    required: false
  }
};
var OutboundLink$1 = "";
const svg = vue.h("svg", {
  "class": "icon outbound",
  "xmlns": "http://www.w3.org/2000/svg",
  "aria-hidden": "true",
  "focusable": "false",
  "x": "0px",
  "y": "0px",
  "viewBox": "0 0 100 100",
  "width": "15",
  "height": "15"
}, [
  vue.h("path", {
    fill: "currentColor",
    d: "M18.8,85.1h56l0,0c2.2,0,4-1.8,4-4v-32h-8v28h-48v-48h28v-8h-32l0,0c-2.2,0-4,1.8-4,4v56C14.8,83.3,16.6,85.1,18.8,85.1z"
  }),
  vue.h("polygon", {
    fill: "currentColor",
    points: "45.7,48.7 51.3,54.3 77.2,28.5 77.2,37.2 85.2,37.2 85.2,14.9 62.8,14.9 62.8,22.9 71.5,22.9"
  })
]);
const OutboundLink = (_, { slots }) => {
  var _a;
  return vue.h("span", [svg, (_a = slots.default) === null || _a === void 0 ? void 0 : _a.call(slots)]);
};
OutboundLink.displayName = "OutboundLink";
const layoutComponents = {
  "404": vue.defineAsyncComponent(() => Promise.resolve().then(function() {
    return _404;
  })),
  "Layout": vue.defineAsyncComponent(() => Promise.resolve().then(function() {
    return Layout;
  }))
};
const Vuepress = vue.defineComponent({
  name: "Vuepress",
  setup() {
    const page = usePageData();
    const layoutComponent = vue.computed(() => {
      let layoutName;
      if (page.value.path) {
        const frontmatterLayout = page.value.frontmatter.layout;
        if (shared.isString(frontmatterLayout)) {
          layoutName = frontmatterLayout;
        } else {
          layoutName = "Layout";
        }
      } else {
        layoutName = "404";
      }
      return layoutComponents[layoutName] || vue.resolveComponent(layoutName, false);
    });
    return () => vue.h(layoutComponent.value);
  }
});
const defineClientAppEnhance = (clientAppEnhance) => clientAppEnhance;
const defineClientAppSetup = (clientAppSetup) => clientAppSetup;
const withBase = (url) => {
  if (shared.isLinkHttp(url))
    return url;
  const base = useSiteData().value.base;
  return `${base}${shared.removeLeadingSlash(url)}`;
};
const _sfc_main$v = /* @__PURE__ */ vue.defineComponent({
  __ssrInlineRender: true,
  props: {
    type: {
      type: String,
      required: false,
      default: "tip"
    },
    text: {
      type: String,
      required: false,
      default: ""
    },
    vertical: {
      type: String,
      required: false,
      default: void 0
    }
  },
  setup(__props) {
    return (_ctx, _push, _parent, _attrs) => {
      _push(`<span${serverRenderer.ssrRenderAttrs(vue.mergeProps({
        class: ["badge", __props.type],
        style: {
          verticalAlign: __props.vertical
        }
      }, _attrs))}>`);
      serverRenderer.ssrRenderSlot(_ctx.$slots, "default", {}, () => {
        _push(`${serverRenderer.ssrInterpolate(__props.text)}`);
      }, _push, _parent);
      _push(`</span>`);
    };
  }
});
const _sfc_setup$u = _sfc_main$v.setup;
_sfc_main$v.setup = (props, ctx) => {
  const ssrContext = vue.useSSRContext();
  (ssrContext.modules || (ssrContext.modules = new Set())).add("../../../../../node_modules/.pnpm/@vuepress+theme-default@2.0.0-beta.26/node_modules/@vuepress/theme-default/lib/client/components/global/Badge.vue");
  return _sfc_setup$u ? _sfc_setup$u(props, ctx) : void 0;
};
var CodeGroup = vue.defineComponent({
  name: "CodeGroup",
  setup(_, { slots }) {
    const activeIndex = vue.ref(-1);
    const tabRefs = vue.ref([]);
    const activateNext = (i = activeIndex.value) => {
      if (i < tabRefs.value.length - 1) {
        activeIndex.value = i + 1;
      } else {
        activeIndex.value = 0;
      }
      tabRefs.value[activeIndex.value].focus();
    };
    const activatePrev = (i = activeIndex.value) => {
      if (i > 0) {
        activeIndex.value = i - 1;
      } else {
        activeIndex.value = tabRefs.value.length - 1;
      }
      tabRefs.value[activeIndex.value].focus();
    };
    const keyboardHandler = (event, i) => {
      if (event.key === " " || event.key === "Enter") {
        event.preventDefault();
        activeIndex.value = i;
      } else if (event.key === "ArrowRight") {
        event.preventDefault();
        activateNext(i);
      } else if (event.key === "ArrowLeft") {
        event.preventDefault();
        activatePrev(i);
      }
    };
    return () => {
      var _a;
      const items2 = (((_a = slots.default) === null || _a === void 0 ? void 0 : _a.call(slots)) || []).filter((vnode) => vnode.type.name === "CodeGroupItem").map((vnode) => {
        if (vnode.props === null) {
          vnode.props = {};
        }
        return vnode;
      });
      if (items2.length === 0) {
        return null;
      }
      if (activeIndex.value < 0 || activeIndex.value > items2.length - 1) {
        activeIndex.value = items2.findIndex((vnode) => vnode.props.active === "" || vnode.props.active === true);
        if (activeIndex.value === -1) {
          activeIndex.value = 0;
        }
      } else {
        items2.forEach((vnode, i) => {
          vnode.props.active = i === activeIndex.value;
        });
      }
      return vue.h("div", { class: "code-group" }, [
        vue.h("div", { class: "code-group__nav" }, vue.h("ul", { class: "code-group__ul" }, items2.map((vnode, i) => {
          const isActive = i === activeIndex.value;
          return vue.h("li", { class: "code-group__li" }, vue.h("button", {
            ref: (element) => {
              if (element) {
                tabRefs.value[i] = element;
              }
            },
            class: {
              "code-group__nav-tab": true,
              "code-group__nav-tab-active": isActive
            },
            ariaPressed: isActive,
            ariaExpanded: isActive,
            onClick: () => activeIndex.value = i,
            onKeydown: (e) => keyboardHandler(e, i)
          }, vnode.props.title));
        }))),
        items2
      ]);
    };
  }
});
const __default__$1 = vue.defineComponent({
  name: "CodeGroupItem"
});
function setup$1(__props) {
  return (_ctx, _push, _parent, _attrs) => {
    _push(`<div${serverRenderer.ssrRenderAttrs(vue.mergeProps({
      class: ["code-group-item", { "code-group-item__active": __props.active }],
      "aria-selected": __props.active
    }, _attrs))}>`);
    serverRenderer.ssrRenderSlot(_ctx.$slots, "default", {}, null, _push, _parent);
    _push(`</div>`);
  };
}
const _sfc_main$u = /* @__PURE__ */ vue.defineComponent(__spreadProps2(__spreadValues2({}, __default__$1), {
  __ssrInlineRender: true,
  props: {
    title: {
      type: String,
      required: true
    },
    active: {
      type: Boolean,
      required: false,
      default: false
    }
  },
  setup: setup$1
}));
const _sfc_setup$t = _sfc_main$u.setup;
_sfc_main$u.setup = (props, ctx) => {
  const ssrContext = vue.useSSRContext();
  (ssrContext.modules || (ssrContext.modules = new Set())).add("../../../../../node_modules/.pnpm/@vuepress+theme-default@2.0.0-beta.26/node_modules/@vuepress/theme-default/lib/client/components/global/CodeGroupItem.vue");
  return _sfc_setup$t ? _sfc_setup$t(props, ctx) : void 0;
};
const darkModeSymbol = Symbol("");
const useDarkMode = () => {
  const isDarkMode = vue.inject(darkModeSymbol);
  if (!isDarkMode) {
    throw new Error("useDarkMode() is called without provider.");
  }
  return isDarkMode;
};
const setupDarkMode = () => {
  const themeLocale = useThemeLocaleData();
  const isDarkPreferred = core.usePreferredDark();
  const darkStorage = core.useStorage("vuepress-color-scheme", "auto");
  const isDarkMode = vue.computed({
    get() {
      if (!themeLocale.value.darkMode) {
        return false;
      }
      if (darkStorage.value === "auto") {
        return isDarkPreferred.value;
      }
      return darkStorage.value === "dark";
    },
    set(val) {
      if (val === isDarkPreferred.value) {
        darkStorage.value = "auto";
      } else {
        darkStorage.value = val ? "dark" : "light";
      }
    }
  });
  vue.provide(darkModeSymbol, isDarkMode);
  updateHtmlDarkClass(isDarkMode);
};
const updateHtmlDarkClass = (isDarkMode) => {
  const update = (value = isDarkMode.value) => {
    const htmlEl = window === null || window === void 0 ? void 0 : window.document.querySelector("html");
    htmlEl === null || htmlEl === void 0 ? void 0 : htmlEl.classList.toggle("dark", value);
  };
  vue.onMounted(() => {
    vue.watch(isDarkMode, update, { immediate: true });
  });
  vue.onUnmounted(() => update());
};
const useResolveRouteWithRedirect = (...args) => {
  const router = vueRouter.useRouter();
  const route = router.resolve(...args);
  const lastMatched = route.matched[route.matched.length - 1];
  if (!(lastMatched === null || lastMatched === void 0 ? void 0 : lastMatched.redirect)) {
    return route;
  }
  const { redirect } = lastMatched;
  const resolvedRedirect = shared.isFunction(redirect) ? redirect(route) : redirect;
  const resolvedRedirectObj = shared.isString(resolvedRedirect) ? { path: resolvedRedirect } : resolvedRedirect;
  return useResolveRouteWithRedirect(__spreadValues2({
    hash: route.hash,
    query: route.query,
    params: route.params
  }, resolvedRedirectObj));
};
const useNavLink = (item) => {
  const resolved = useResolveRouteWithRedirect(item);
  return {
    text: resolved.meta.title || item,
    link: resolved.name === "404" ? item : resolved.fullPath
  };
};
let promise = null;
let promiseResolve = null;
const scrollPromise = {
  wait: () => promise,
  pending: () => {
    promise = new Promise((resolve) => promiseResolve = resolve);
  },
  resolve: () => {
    promiseResolve === null || promiseResolve === void 0 ? void 0 : promiseResolve();
    promise = null;
    promiseResolve = null;
  }
};
const useScrollPromise = () => scrollPromise;
const sidebarItemsSymbol = Symbol("sidebarItems");
const useSidebarItems = () => {
  const sidebarItems = vue.inject(sidebarItemsSymbol);
  if (!sidebarItems) {
    throw new Error("useSidebarItems() is called without provider.");
  }
  return sidebarItems;
};
const setupSidebarItems = () => {
  const themeLocale = useThemeLocaleData();
  const frontmatter = usePageFrontmatter();
  const sidebarItems = vue.computed(() => resolveSidebarItems(frontmatter.value, themeLocale.value));
  vue.provide(sidebarItemsSymbol, sidebarItems);
};
const resolveSidebarItems = (frontmatter, themeLocale) => {
  var _a, _b, _c, _d;
  const sidebarConfig = (_b = (_a = frontmatter.sidebar) !== null && _a !== void 0 ? _a : themeLocale.sidebar) !== null && _b !== void 0 ? _b : "auto";
  const sidebarDepth = (_d = (_c = frontmatter.sidebarDepth) !== null && _c !== void 0 ? _c : themeLocale.sidebarDepth) !== null && _d !== void 0 ? _d : 2;
  if (frontmatter.home || sidebarConfig === false) {
    return [];
  }
  if (sidebarConfig === "auto") {
    return resolveAutoSidebarItems(sidebarDepth);
  }
  if (shared.isArray(sidebarConfig)) {
    return resolveArraySidebarItems(sidebarConfig, sidebarDepth);
  }
  if (shared.isPlainObject(sidebarConfig)) {
    return resolveMultiSidebarItems(sidebarConfig, sidebarDepth);
  }
  return [];
};
const headerToSidebarItem = (header, sidebarDepth) => ({
  text: header.title,
  link: `#${header.slug}`,
  children: headersToSidebarItemChildren(header.children, sidebarDepth)
});
const headersToSidebarItemChildren = (headers, sidebarDepth) => sidebarDepth > 0 ? headers.map((header) => headerToSidebarItem(header, sidebarDepth - 1)) : [];
const resolveAutoSidebarItems = (sidebarDepth) => {
  const page = usePageData();
  return [
    {
      text: page.value.title,
      children: headersToSidebarItemChildren(page.value.headers, sidebarDepth)
    }
  ];
};
const resolveArraySidebarItems = (sidebarConfig, sidebarDepth) => {
  const route = vueRouter.useRoute();
  const page = usePageData();
  const handleChildItem = (item) => {
    var _a;
    let childItem;
    if (shared.isString(item)) {
      childItem = useNavLink(item);
    } else {
      childItem = item;
    }
    if (childItem.children) {
      return __spreadProps2(__spreadValues2({}, childItem), {
        children: childItem.children.map((item2) => handleChildItem(item2))
      });
    }
    if (childItem.link === route.path) {
      const headers = ((_a = page.value.headers[0]) === null || _a === void 0 ? void 0 : _a.level) === 1 ? page.value.headers[0].children : page.value.headers;
      return __spreadProps2(__spreadValues2({}, childItem), {
        children: headersToSidebarItemChildren(headers, sidebarDepth)
      });
    }
    return childItem;
  };
  return sidebarConfig.map((item) => handleChildItem(item));
};
const resolveMultiSidebarItems = (sidebarConfig, sidebarDepth) => {
  var _a;
  const route = vueRouter.useRoute();
  const sidebarPath = shared.resolveLocalePath(sidebarConfig, route.path);
  const matchedSidebarConfig = (_a = sidebarConfig[sidebarPath]) !== null && _a !== void 0 ? _a : [];
  return resolveArraySidebarItems(matchedSidebarConfig, sidebarDepth);
};
const themeData$1 = {
  "logo": "/logo-rounded.svg",
  "navbar": [
    {
      "text": "Guide",
      "link": "/guide/"
    },
    {
      "text": "Features",
      "link": "/features/"
    },
    {
      "text": "CSL",
      "link": "/csl/"
    }
  ],
  "sidebar": {
    "/": [
      {
        "text": "Guide",
        "children": [
          "/guide/download.md",
          "/guide/configuration.md"
        ]
      },
      {
        "text": "Features",
        "children": [
          "/features/indexers.md",
          "/features/downloaders.md",
          "/features/profiles.md",
          "/features/notifications.md",
          "/features/post-processing.md",
          "/features/syncing.md"
        ]
      },
      {
        "text": "CSL",
        "children": [
          "/csl/introduction.md",
          "/csl/examples.md",
          "/csl/repl.md"
        ]
      }
    ]
  },
  "locales": {
    "/": {
      "selectLanguageName": "English"
    }
  },
  "darkMode": true,
  "repo": null,
  "selectLanguageText": "Languages",
  "selectLanguageAriaLabel": "Select language",
  "sidebarDepth": 2,
  "editLink": true,
  "editLinkText": "Edit this page",
  "lastUpdated": true,
  "lastUpdatedText": "Last Updated",
  "contributors": true,
  "contributorsText": "Contributors",
  "notFound": [
    "There's nothing here.",
    "How did we get here?",
    "That's a Four-Oh-Four.",
    "Looks like we've got some broken links."
  ],
  "backToHome": "Take me home",
  "openInNewWindow": "open in new window",
  "toggleDarkMode": "toggle dark mode",
  "toggleSidebar": "toggle sidebar"
};
const themeData = vue.ref(themeData$1);
const useThemeData = () => themeData;
if (false) {
  __VUE_HMR_RUNTIME__.updateThemeData = (data3) => {
    themeData.value = data3;
  };
}
const themeLocaleDataSymbol = Symbol("");
const useThemeLocaleData$1 = () => {
  const themeLocaleData = vue.inject(themeLocaleDataSymbol);
  if (!themeLocaleData) {
    throw new Error("useThemeLocaleData() is called without provider.");
  }
  return themeLocaleData;
};
const resolveThemeLocaleData = (theme, routeLocale) => {
  var _a;
  return __spreadValues2(__spreadValues2({}, theme), (_a = theme.locales) === null || _a === void 0 ? void 0 : _a[routeLocale]);
};
const useThemeLocaleData = () => useThemeLocaleData$1();
const _sfc_main$t = /* @__PURE__ */ vue.defineComponent({
  __ssrInlineRender: true,
  setup(__props) {
    const themeLocale = useThemeLocaleData();
    return (_ctx, _push, _parent, _attrs) => {
      _push(serverRenderer.ssrRenderComponent(vue.unref(OutboundLink), _attrs, {
        default: vue.withCtx((_, _push2, _parent2, _scopeId) => {
          if (_push2) {
            _push2(`<span class="sr-only"${_scopeId}>${serverRenderer.ssrInterpolate(vue.unref(themeLocale).openInNewWindow)}</span>`);
          } else {
            return [
              vue.createVNode("span", { class: "sr-only" }, vue.toDisplayString(vue.unref(themeLocale).openInNewWindow), 1)
            ];
          }
        }),
        _: 1
      }, _parent));
    };
  }
});
const _sfc_setup$s = _sfc_main$t.setup;
_sfc_main$t.setup = (props, ctx) => {
  const ssrContext = vue.useSSRContext();
  (ssrContext.modules || (ssrContext.modules = new Set())).add("../../../../../node_modules/.pnpm/@vuepress+theme-default@2.0.0-beta.26/node_modules/@vuepress/theme-default/lib/client/components/global/OutboundLink.vue");
  return _sfc_setup$s ? _sfc_setup$s(props, ctx) : void 0;
};
var index$F = "";
var clientAppEnhance0 = defineClientAppEnhance(({ app, router }) => {
  app.component("Badge", _sfc_main$v);
  app.component("CodeGroup", CodeGroup);
  app.component("CodeGroupItem", _sfc_main$u);
  delete app._context.components.OutboundLink;
  app.component("OutboundLink", _sfc_main$t);
  app.component("NavbarSearch", () => {
    const SearchComponent = app.component("Docsearch") || app.component("SearchBox");
    if (SearchComponent) {
      return vue.h(SearchComponent);
    }
    return null;
  });
  const scrollBehavior = router.options.scrollBehavior;
  router.options.scrollBehavior = async (...args) => {
    await useScrollPromise().wait();
    return scrollBehavior(...args);
  };
});
var vars$2 = "";
var mediumZoom = "";
var clientAppEnhance1 = defineClientAppEnhance(({ app, router }) => {
  return;
});
var clientAppEnhance2 = defineClientAppEnhance(({ app }) => {
  const themeData2 = useThemeData();
  const routeLocale = app._context.provides[routeLocaleSymbol];
  const themeLocaleData = vue.computed(() => resolveThemeLocaleData(themeData2.value, routeLocale.value));
  app.provide(themeLocaleDataSymbol, themeLocaleData);
  Object.defineProperties(app.config.globalProperties, {
    $theme: {
      get() {
        return themeData2.value;
      }
    },
    $themeLocale: {
      get() {
        return themeLocaleData.value;
      }
    }
  });
});
const clientAppEnhances = [
  clientAppEnhance0,
  clientAppEnhance1,
  clientAppEnhance2
];
function r(r2, e, n2) {
  var i, t, o;
  e === void 0 && (e = 50), n2 === void 0 && (n2 = {});
  var a = (i = n2.isImmediate) != null && i, u = (t = n2.callback) != null && t, c = n2.maxWait, v = Date.now(), l2 = [];
  function f() {
    if (c !== void 0) {
      var r3 = Date.now() - v;
      if (r3 + e >= c)
        return c - r3;
    }
    return e;
  }
  var d = function() {
    var e2 = [].slice.call(arguments), n3 = this;
    return new Promise(function(i2, t2) {
      var c2 = a && o === void 0;
      if (o !== void 0 && clearTimeout(o), o = setTimeout(function() {
        if (o = void 0, v = Date.now(), !a) {
          var i3 = r2.apply(n3, e2);
          u && u(i3), l2.forEach(function(r3) {
            return (0, r3.resolve)(i3);
          }), l2 = [];
        }
      }, f()), c2) {
        var d2 = r2.apply(n3, e2);
        return u && u(d2), i2(d2);
      }
      l2.push({ resolve: i2, reject: t2 });
    });
  };
  return d.cancel = function(r3) {
    o !== void 0 && clearTimeout(o), l2.forEach(function(e2) {
      return (0, e2.reject)(r3);
    }), l2 = [];
  }, d;
}
const getScrollTop = () => window.pageYOffset || document.documentElement.scrollTop || document.body.scrollTop || 0;
const scrollToTop = () => window.scrollTo({ top: 0, behavior: "smooth" });
var vars$1 = "";
var backToTop = "";
const BackToTop = vue.defineComponent({
  name: "BackToTop",
  setup() {
    const scrollTop = vue.ref(0);
    const show = vue.computed(() => scrollTop.value > 300);
    vue.onMounted(() => {
      scrollTop.value = getScrollTop();
      window.addEventListener("scroll", r(() => {
        scrollTop.value = getScrollTop();
      }, 100));
    });
    const backToTopEl = vue.h("div", { class: "back-to-top", onClick: scrollToTop });
    return () => vue.h(vue.Transition, {
      name: "back-to-top"
    }, {
      default: () => show.value ? backToTopEl : null
    });
  }
});
const clientAppRootComponents = [
  BackToTop
];
var clientAppSetup0 = defineClientAppSetup(() => {
  setupDarkMode();
  setupSidebarItems();
});
var clientAppSetup1 = defineClientAppSetup(() => {
  return;
});
var vars = "";
var nprogress = "";
const useNprogress = () => {
  vue.onMounted(() => {
    const router = vueRouter.useRouter();
    const loadedPages = new Set();
    loadedPages.add(router.currentRoute.value.path);
    nprogress__namespace.configure({ showSpinner: false });
    router.beforeEach((to) => {
      if (!loadedPages.has(to.path)) {
        nprogress__namespace.start();
      }
    });
    router.afterEach((to) => {
      loadedPages.add(to.path);
      nprogress__namespace.done();
    });
  });
};
var clientAppSetup2 = defineClientAppSetup(() => {
  useNprogress();
});
const clientAppSetups = [
  clientAppSetup0,
  clientAppSetup1,
  clientAppSetup2
];
const routeItems = [
  ["v-8daa1a0e", "/", "", ["/index.html", "/README.md"]],
  ["v-0c47f4f8", "/csl/examples.html", "Examples", ["/csl/examples", "/csl/examples.md"]],
  ["v-4760ba53", "/csl/introduction.html", "Introduction", ["/csl/introduction", "/csl/introduction.md"]],
  ["v-dc33cfc4", "/csl/repl.html", "REPL", ["/csl/repl", "/csl/repl.md"]],
  ["v-4f4ccb8f", "/guide/configuration.html", "Configuration", ["/guide/configuration", "/guide/configuration.md"]],
  ["v-62059836", "/guide/download.html", "Download and Installation", ["/guide/download", "/guide/download.md"]],
  ["v-8b3d4f7c", "/features/calendar.html", "Calendar", ["/features/calendar", "/features/calendar.md"]],
  ["v-7e926318", "/features/downloaders.html", "Downloaders", ["/features/downloaders", "/features/downloaders.md"]],
  ["v-7440cc28", "/features/indexers.html", "Indexers", ["/features/indexers", "/features/indexers.md"]],
  ["v-318bcd2c", "/features/notifications.html", "Notifications", ["/features/notifications", "/features/notifications.md"]],
  ["v-0b1540b2", "/features/post-processing.html", "Post-Processing", ["/features/post-processing", "/features/post-processing.md"]],
  ["v-3bcd3316", "/features/profiles.html", "Profiles", ["/features/profiles", "/features/profiles.md"]],
  ["v-5f90410b", "/features/syncing.html", "Syncing", ["/features/syncing", "/features/syncing.md"]],
  ["v-3706649a", "/404.html", "", ["/404"]]
];
const pagesRoutes = routeItems.reduce((result, [name, path, title, redirects]) => {
  result.push({
    name,
    path,
    component: Vuepress,
    meta: { title }
  }, ...redirects.map((item) => ({
    path: item,
    redirect: path
  })));
  return result;
}, [
  {
    name: "404",
    path: "/:catchAll(.*)",
    component: Vuepress
  }
]);
const provideGlobalComputed = (app, router) => {
  const routeLocale = vue.computed(() => resolveRouteLocale(siteData.value.locales, router.currentRoute.value.path));
  const siteLocaleData = vue.computed(() => resolveSiteLocaleData(siteData.value, routeLocale.value));
  const pageFrontmatter = vue.computed(() => resolvePageFrontmatter(pageData.value));
  const pageHeadTitle = vue.computed(() => resolvePageHeadTitle(pageData.value, siteLocaleData.value));
  const pageHead = vue.computed(() => resolvePageHead(pageHeadTitle.value, pageFrontmatter.value, siteLocaleData.value));
  const pageLang = vue.computed(() => resolvePageLang(pageData.value));
  app.provide(routeLocaleSymbol, routeLocale);
  app.provide(siteLocaleDataSymbol, siteLocaleData);
  app.provide(pageFrontmatterSymbol, pageFrontmatter);
  app.provide(pageHeadTitleSymbol, pageHeadTitle);
  app.provide(pageHeadSymbol, pageHead);
  app.provide(pageLangSymbol, pageLang);
  Object.defineProperties(app.config.globalProperties, {
    $frontmatter: { get: () => pageFrontmatter.value },
    $headTitle: { get: () => pageHeadTitle.value },
    $lang: { get: () => pageLang.value },
    $page: { get: () => pageData.value },
    $routeLocale: { get: () => routeLocale.value },
    $site: { get: () => siteData.value },
    $siteLocale: { get: () => siteLocaleData.value },
    $withBase: { get: () => withBase }
  });
};
const registerGlobalComponents = (app) => {
  app.component("ClientOnly", ClientOnly);
  app.component("Content", Content);
  app.component("OutboundLink", OutboundLink);
};
const appCreator = vue.createSSRApp;
const historyCreator = vueRouter.createMemoryHistory;
const createVueApp = async () => {
  const app = appCreator({
    name: "VuepressApp",
    setup() {
      setupUpdateHead();
      for (const clientAppSetup of clientAppSetups) {
        clientAppSetup();
      }
      return () => [
        vue.h(vueRouter.RouterView),
        ...clientAppRootComponents.map((comp) => vue.h(comp))
      ];
    }
  });
  const router = vueRouter.createRouter({
    history: historyCreator(shared.removeEndingSlash(siteData.value.base)),
    routes: pagesRoutes,
    scrollBehavior: (to, from, savedPosition) => {
      if (savedPosition)
        return savedPosition;
      if (to.hash)
        return { el: to.hash };
      return { top: 0 };
    }
  });
  router.beforeResolve(async (to, from) => {
    var _a;
    if (to.path !== from.path || from === vueRouter.START_LOCATION) {
      [pageData.value] = await Promise.all([
        resolvePageData(to.name),
        (_a = pagesComponents[to.name]) === null || _a === void 0 ? void 0 : _a.__asyncLoader()
      ]);
    }
  });
  provideGlobalComputed(app, router);
  registerGlobalComponents(app);
  for (const clientAppEnhance of clientAppEnhances) {
    await clientAppEnhance({ app, router, siteData });
  }
  app.use(router);
  return {
    app,
    router
  };
};
var _export_sfc$1 = (sfc, props) => {
  for (const [key, val] of props) {
    sfc[key] = val;
  }
  return sfc;
};
const _sfc_main$s = {};
function _sfc_ssrRender$d(_ctx, _push, _parent, _attrs) {
}
const _sfc_setup$r = _sfc_main$s.setup;
_sfc_main$s.setup = (props, ctx) => {
  const ssrContext = vue.useSSRContext();
  (ssrContext.modules || (ssrContext.modules = new Set())).add("../pages/index.html.vue");
  return _sfc_setup$r ? _sfc_setup$r(props, ctx) : void 0;
};
var index_html$1 = /* @__PURE__ */ _export_sfc$1(_sfc_main$s, [["ssrRender", _sfc_ssrRender$d]]);
var index_html$2 = /* @__PURE__ */ Object.freeze({
  __proto__: null,
  [Symbol.toStringTag]: "Module",
  "default": index_html$1
});
const _sfc_main$r = {};
function _sfc_ssrRender$c(_ctx, _push, _parent, _attrs) {
  _push(`<h1${serverRenderer.ssrRenderAttrs(vue.mergeProps({
    id: "examples",
    tabindex: "-1"
  }, _attrs))}><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h1>`);
}
const _sfc_setup$q = _sfc_main$r.setup;
_sfc_main$r.setup = (props, ctx) => {
  const ssrContext = vue.useSSRContext();
  (ssrContext.modules || (ssrContext.modules = new Set())).add("../pages/csl/examples.html.vue");
  return _sfc_setup$q ? _sfc_setup$q(props, ctx) : void 0;
};
var examples_html$1 = /* @__PURE__ */ _export_sfc$1(_sfc_main$r, [["ssrRender", _sfc_ssrRender$c]]);
var examples_html$2 = /* @__PURE__ */ Object.freeze({
  __proto__: null,
  [Symbol.toStringTag]: "Module",
  "default": examples_html$1
});
const _sfc_main$q = {};
function _sfc_ssrRender$b(_ctx, _push, _parent, _attrs) {
  const _component_OutboundLink = vue.resolveComponent("OutboundLink");
  _push(`<!--[--><h1 id="csl-conductorr-specific-language" tabindex="-1"><a class="header-anchor" href="#csl-conductorr-specific-language" aria-hidden="true">#</a> CSL (Conductorr-Specific Language)</h1><p>CSL is a stripped-down Lisp implementation intended to allow users of Conductorr to extend its capabilities and finely tune release processing.</p><p>The parsing capabilities are largely borrowed from <a href="https://github.com/janne" target="_blank" rel="noopener noreferrer">Jan Andersson`);
  _push(serverRenderer.ssrRenderComponent(_component_OutboundLink, null, null, _parent));
  _push(`</a>&#39;s <a href="https://github.com/janne/go-lisp" target="_blank" rel="noopener noreferrer">go-lisp`);
  _push(serverRenderer.ssrRenderComponent(_component_OutboundLink, null, null, _parent));
  _push(`</a> project.</p><p>The main differences between go-lisp and CSL are that CSL is <em>not</em> Turing-Complete, such that loops or recursion are impossible to implement. This allows us to keep the language simple and avoid the halting problem when running user-defined scripts. This means user scripts in CSL cannot contain loops or function definitions. The other main difference is that CSL has a number of predefined functions that are specifically written to aid in the processing of Newznab releases. This helps avoid the loss of functionality due to the restricted iteration and function-defining capabilities of CSL.</p><h2 id="syntax" tabindex="-1"><a class="header-anchor" href="#syntax" aria-hidden="true">#</a> Syntax</h2><p>If you are familiar with Lisp syntax, CSL follows it directly. Here is an example of a script in CSL:</p><div class="language-lisp ext-lisp line-numbers-mode"><pre class="language-lisp"><code><span class="token comment">;; Define a variable</span>
<span class="token punctuation">(</span><span class="token car">define</span> x <span class="token number">17</span><span class="token punctuation">)</span>

<span class="token comment">;; Return the result of x * 37</span>
<span class="token punctuation">(</span><span class="token car">*</span> x <span class="token number">37</span><span class="token punctuation">)</span>
</code></pre><div class="line-numbers"><span class="line-number">1</span><br><span class="line-number">2</span><br><span class="line-number">3</span><br><span class="line-number">4</span><br><span class="line-number">5</span><br></div></div><p>The value <code>629</code> is returned.</p><h3 id="data-types-and-literals" tabindex="-1"><a class="header-anchor" href="#data-types-and-literals" aria-hidden="true">#</a> Data Types and Literals</h3><p>CSL has supports the following datatypes:</p><ul><li>Integers</li><li>Strings</li><li>Booleans</li><li>Lists</li></ul><h4 id="integers" tabindex="-1"><a class="header-anchor" href="#integers" aria-hidden="true">#</a> Integers</h4><p>All CSL integers are 64-bit numbers. They can be negative, positive, or zero. One special feature of CSL is the ability to use suffixes that are a shorthand to multiply by an order of magnitude. For example, <code>3G</code> is equivalent to <code>3000000000</code>.</p><p>The full list of available suffixes is below:</p><ul><li><code>G</code> --&gt; 1000000000</li><li><code>Gi</code> --&gt; 2^30</li><li><code>M</code> --&gt; 1000000</li><li><code>Mi</code> --&gt; 2^20</li><li><code>K</code> --&gt; 1000</li><li><code>Ki</code> --&gt; 2^10</li><li><code>B</code> --&gt; 8</li></ul><h4 id="strings" tabindex="-1"><a class="header-anchor" href="#strings" aria-hidden="true">#</a> Strings</h4><p>String literals must be enclosed in double quotes (<code>&quot;</code>)</p><h4 id="booleans" tabindex="-1"><a class="header-anchor" href="#booleans" aria-hidden="true">#</a> Booleans</h4><p>Boolean literals are <code>true</code> and <code>false</code>, and are case-sensitive</p><h4 id="lists" tabindex="-1"><a class="header-anchor" href="#lists" aria-hidden="true">#</a> Lists</h4><p>Lists are implicitly defined any time expressions are joined within parentheses such as <code>(1 4 &quot;hello&quot; 5 2 &quot;world&quot; true)</code>. List elements are not required to have the same data type. Single-value lists may be treated as regular values or lists. (ie, <code>(+ (1) (2))</code> is equivalent to <code>(+ 1 2)</code> and <code>(in 1 (1))</code> evaluates correctly as well).</p><h2 id="built-in-functions" tabindex="-1"><a class="header-anchor" href="#built-in-functions" aria-hidden="true">#</a> Built-In Functions</h2><ul><li><code>(+ x y ...)</code> Adds all integer arguments together</li><li><code>(- x y ...)</code> Subtracts subsequent integer arguments from the first one</li><li><code>(* x y ...)</code> Multiplies all integer arguments together</li><li><code>(/ x y ...)</code> Divides subsequent integer arguments from the first one</li><li><code>(&gt; x y ...)</code> Returns true iff <code>x</code> is strictly greater than all subsequent integer arguments</li><li><code>(&gt;&gt; x y ...)</code> Returns true iff each integer argument is strictly greater than the ones before it</li><li><code>(&lt; x y ...)</code> Returns true iff <code>x</code> is strictly less than all subsequent integer arguments</li><li><code>(&lt;&lt; x y ...)</code> Returns true iff each integer argument is strictly less than the ones before it</li><li><code>(&gt;= x y ...)</code> Returns true iff <code>x</code> is greater than or equal to all subsequent integer arguments</li><li><code>(&gt;&gt;= x y ...)</code> Returns true iff each integer argument is strictly greater than or equal to the ones before it</li><li><code>(&lt;= x y ...)</code> Returns true iff <code>x</code> is less than or equal to all subsequent integer arguments</li><li><code>(&lt;&lt;= x y ...)</code> Returns true iff each integer argument is strictly less than or equal to the ones before it</li><li><code>(&gt;&lt; x y v ...)</code> Returns true iff for all integer arguments <code>v</code> that <code>v</code> \u2208 (<code>x</code>, <code>y</code>)</li><li><code>(&gt;&lt;= x y v ...)</code> Returns true iff for all integer arguments <code>v</code> that <code>v</code> \u2208 (<code>x</code>, <code>y</code>]</li><li><code>(&gt;=&lt; x y v ...)</code> Returns true iff for all integer arguments <code>v</code> that <code>v</code> \u2208 [<code>x</code>, <code>y</code>)</li><li><code>(&gt;=&lt;= x y v ...)</code> Returns true iff for all integer arguments <code>v</code> that <code>v</code> \u2208 [<code>x</code>, <code>y</code>]</li><li><code>(&lt;&gt; x y v ...)</code> Returns true iff for all integer arguments <code>v</code> that <code>v</code> \u2208 [-2<sup>63</sup>, x)\u2229(<code>y</code>, 2<sup>63</sup>-1]</li><li><code>(&lt;=&gt;= x y v ...)</code> Returns true iff for all integer arguments <code>v</code> that <code>v</code> \u2208 [-2<sup>63</sup>, <code>x</code>]\u2229[<code>y</code>, 2<sup>63</sup>-1]</li><li><code>(&lt;=&gt; x y v ...)</code> Returns true iff for all integer arguments <code>v</code> that <code>v</code> \u2208 [-2<sup>63</sup>, <code>x</code>]\u2229(<code>y</code>, 2<sup>63</sup>-1]</li><li><code>(&lt;&gt;= x y v ...)</code> Returns true iff for all integer arguments <code>v</code> that <code>v</code> \u2208 [-2<sup>63</sup>, <code>x</code>)\u2229[<code>y</code>, 2<sup>63</sup>-1]</li><li><code>(eq x y ...)</code> Returns true if all arguments are equal to each other in both type and value. In the case of lists, the elements must be in the same order to be considered equal. (Under the hood we are just using Go&#39;s <code>reflect.DeepEqual</code>)</li><li><code>(define x expr)</code> Defines a variable <code>x</code> initialized with the result of <code>expr</code></li><li><code>(in v l)</code> Returns true if the value of <code>v</code> appears in list <code>l</code></li><li><code>(in s1 s2)</code> Returns true iff <code>s1</code> is a substring of <code>s2</code></li><li><code>(nth i l)</code> Returns the <code>i</code>th value in list <code>l</code> or error if out of bounds</li><li><code>(nths i s)</code> Returns the <code>i</code>th character in the string <code>s</code></li><li><code>(len l)</code> Returns the length of list <code>l</code></li><li><code>(lens s)</code> Returns the length of string <code>s</code></li><li><code>(append l v ...)</code> Appends <code>v</code> and all subsequent arguments in order to the right end of list <code>l</code>. If the list <code>l</code> does not yet exist, then it is initialized</li><li><code>(appendleft l v ...)</code> Appends <code>v</code> and all subsequent arguments in order to the left end of list <code>l</code>. If the list <code>l</code> does not yet exist, then it is initialized</li><li><code>(pop l ...)</code> Removes the rightmost element in list <code>l</code>, returning the removed element</li><li><code>(popleft l ...)</code> Removes the leftmost element in list <code>l</code>, returning the removed element</li><li><code>(if p expr1 expr2)</code> Returns the result of <code>expr1</code> if <code>p</code> evaluates to <code>true</code>, and returns the result of <code>expr2</code> otherwise</li><li><code>(and p ...)</code> Returns the result of a conditional AND applied to all operands in the expression. Returns <code>true</code> if there is only one operand</li><li><code>(or p ...)</code> Returns the result of a conditional OR applied to all operands in the expression. Returns <code>true</code> if there is only one operand</li><li><code>(not p)</code> Returns the result of conditional inverse applied to p</li><li><code>(join s v ...)</code> Joins elements elements <code>v1</code>, <code>v2</code>, ... into a string using separator <code>s</code>. Does not add a separator to the end of the resultant string. Non-string elements such as integers will display using <code>fmt.Sprintf(%v, v)</code></li><li><code>(split s d)</code> Splits a string into a list of strings using delimiter <code>d</code></li></ul><blockquote><p>NOTE: All of the above functions are evaluated <em>eagerly</em>, meaning their arguments are evaluated before the function is called. The exceptions to this are the <code>if</code>, <code>append</code>, <code>appendleft</code>, <code>pop</code>, and <code>popleft</code> functions, which only evaluate their arguments as they are needed.</p></blockquote><h2 id="extension-functions" tabindex="-1"><a class="header-anchor" href="#extension-functions" aria-hidden="true">#</a> Extension Functions</h2><p>Conductorr extends the base CSL specification with functions specific to its own functionality for processing releases.</p><p>A <code>release</code> type is just a list defined like so:</p><div class="language-lisp ext-lisp line-numbers-mode"><pre class="language-lisp"><code><span class="token punctuation">(</span><span class="token car">define</span> a
  <span class="token punctuation">(</span>
    <span class="token string">&quot;Manos.The.Hands.of.Fate.1966.THEATRi<span class="token argument">CAL</span>.1080p.BluRay.x264-SADPANDA&quot;</span> 
    <span class="token string">&quot;TorrentLeech&quot;</span> 
    <span class="token string">&quot;torrent&quot;</span> 
    <span class="token string">&quot;movie&quot;</span> 
    <span class="token string">&quot;BDRip&quot;</span> 
    <span class="token string">&quot;1080p&quot;</span> 
    <span class="token string">&quot;x264&quot;</span> 
    <span class="token number">32</span> 
    <span class="token number">794</span> 
    <span class="token number">10737418240</span> 
    <span class="token number">288</span>
  <span class="token punctuation">)</span>
<span class="token punctuation">)</span>
</code></pre><div class="line-numbers"><span class="line-number">1</span><br><span class="line-number">2</span><br><span class="line-number">3</span><br><span class="line-number">4</span><br><span class="line-number">5</span><br><span class="line-number">6</span><br><span class="line-number">7</span><br><span class="line-number">8</span><br><span class="line-number">9</span><br><span class="line-number">10</span><br><span class="line-number">11</span><br><span class="line-number">12</span><br><span class="line-number">13</span><br><span class="line-number">14</span><br><span class="line-number">15</span><br></div></div><p>The properties are in order: Release title, indexer, download type (either torrent or nzb), content type (either movie or series), rip type, resolution, encoding, number of seeders, age in days, size in bytes, and runtime in minutes.</p><p>Possible values for rip type are <code>CAM</code>, <code>TELESYNC</code>, <code>SCR</code>, <code>HDTV</code>, <code>DVDRip</code>, <code>HDRip</code>,<code>WEBCap</code>, <code>WEBRip</code>, <code>WEBDL</code>, <code>BDRip</code>.</p><p>Possible values for resolution are <code>352p</code>, <code>360p</code>, <code>480i</code>, <code>480p</code>, <code>576i</code>, <code>576p</code>, <code>720p</code>, <code>1080p</code>, <code>2160p</code></p><p>Possible values for encoding are <code>Xvid</code>, <code>x264</code>, <code>x265</code>, <code>VP9</code></p><p>The following functions are defined on <code>release</code> types:</p><ul><li><code>r-title</code> Retrieves release title</li><li><code>r-indexer</code> Retrieves release indexer</li><li><code>r-downloadtype</code> Retrieves release download type</li><li><code>r-contenttype</code> Retrieves release content type</li><li><code>r-riptype</code> Retrieves release rip type</li><li><code>r-resolution</code> Retrieves release resolution</li><li><code>r-encoding</code> Retrieves release encoding</li><li><code>r-seeders</code> Retrieves number of seeders for release</li><li><code>r-age</code> Retrieves age in days for release</li><li><code>r-size</code> Retrieves size of release in bytes</li><li><code>r-runtime</code> Retrieves the runtime of the release in minutes</li><li><code>r-riptype-order</code> Returns a number designating the order of how preferrable a riptype is (<code>CAM &lt; TELESYNC &lt; SCR &lt; HDTV &lt; DVDRip &lt; HDRip &lt; WEBCap &lt; WEBRip &lt; WEBDL &lt; BDRip</code>)</li><li><code>r-resolution-order</code> Returns a number designating the order of how preferrable a resolution is (<code>352p &lt; 360p &lt; 480i &lt; 480p &lt; 576i &lt; 576p &lt; 720p &lt; 1080p &lt; 2160p</code>)</li><li><code>r-encoding-order</code> Returns a number designating the order of how preferrable an encodign is (<code>Xvid &lt; x264 &lt; x265 &lt; VP9</code>)</li></ul><p>To call use one of these functions, do the following:</p><p><code>(r-title a)</code> where <code>a</code> is a release as defined above</p><p>The above snippet returns &quot;Manos.The.Hands.of.Fate.1966.THEATRiCAL.1080p.BluRay.x264-SADPANDA&quot; when <code>a</code> is defined like in the earlier example</p><!--]-->`);
}
const _sfc_setup$p = _sfc_main$q.setup;
_sfc_main$q.setup = (props, ctx) => {
  const ssrContext = vue.useSSRContext();
  (ssrContext.modules || (ssrContext.modules = new Set())).add("../pages/csl/introduction.html.vue");
  return _sfc_setup$p ? _sfc_setup$p(props, ctx) : void 0;
};
var introduction_html$1 = /* @__PURE__ */ _export_sfc$1(_sfc_main$q, [["ssrRender", _sfc_ssrRender$b]]);
var introduction_html$2 = /* @__PURE__ */ Object.freeze({
  __proto__: null,
  [Symbol.toStringTag]: "Module",
  "default": introduction_html$1
});
var __defProp = Object.defineProperty;
var __defProps = Object.defineProperties;
var __getOwnPropDescs = Object.getOwnPropertyDescriptors;
var __getOwnPropSymbols = Object.getOwnPropertySymbols;
var __hasOwnProp = Object.prototype.hasOwnProperty;
var __propIsEnum = Object.prototype.propertyIsEnumerable;
var __defNormalProp = (obj, key, value) => key in obj ? __defProp(obj, key, { enumerable: true, configurable: true, writable: true, value }) : obj[key] = value;
var __spreadValues = (a, b) => {
  for (var prop in b || (b = {}))
    if (__hasOwnProp.call(b, prop))
      __defNormalProp(a, prop, b[prop]);
  if (__getOwnPropSymbols)
    for (var prop of __getOwnPropSymbols(b)) {
      if (__propIsEnum.call(b, prop))
        __defNormalProp(a, prop, b[prop]);
    }
  return a;
};
var __spreadProps = (a, b) => __defProps(a, __getOwnPropDescs(b));
var __objRest = (source, exclude) => {
  var target = {};
  for (var prop in source)
    if (__hasOwnProp.call(source, prop) && exclude.indexOf(prop) < 0)
      target[prop] = source[prop];
  if (source != null && __getOwnPropSymbols)
    for (var prop of __getOwnPropSymbols(source)) {
      if (exclude.indexOf(prop) < 0 && __propIsEnum.call(source, prop))
        target[prop] = source[prop];
    }
  return target;
};
function _extends() {
  _extends = Object.assign || function(target) {
    for (var i = 1; i < arguments.length; i++) {
      var source = arguments[i];
      for (var key in source) {
        if (Object.prototype.hasOwnProperty.call(source, key)) {
          target[key] = source[key];
        }
      }
    }
    return target;
  };
  return _extends.apply(this, arguments);
}
var KEYCODE_ENTER = 13;
var KEYCODE_TAB = 9;
var KEYCODE_BACKSPACE = 8;
var KEYCODE_Y = 89;
var KEYCODE_Z = 90;
var KEYCODE_M = 77;
var KEYCODE_PARENS = 57;
var KEYCODE_BRACKETS = 219;
var KEYCODE_QUOTE = 222;
var KEYCODE_BACK_QUOTE = 192;
var KEYCODE_ESCAPE = 27;
var HISTORY_LIMIT = 100;
var HISTORY_TIME_GAP = 3e3;
var isWindows = typeof window !== "undefined" && navigator && /* @__PURE__ */ /Win/i.test(navigator.platform);
var isMacLike = typeof window !== "undefined" && navigator && /* @__PURE__ */ /(Mac|iPhone|iPod|iPad)/i.test(navigator.platform);
var PrismEditor = /* @__PURE__ */ vue.defineComponent({
  props: {
    lineNumbers: {
      type: Boolean,
      "default": false
    },
    autoStyleLineNumbers: {
      type: Boolean,
      "default": true
    },
    readonly: {
      type: Boolean,
      "default": false
    },
    modelValue: {
      type: String,
      "default": ""
    },
    highlight: {
      type: Function,
      required: true
    },
    tabSize: {
      type: Number,
      "default": 2
    },
    insertSpaces: {
      type: Boolean,
      "default": true
    },
    ignoreTabKey: {
      type: Boolean,
      "default": false
    },
    placeholder: {
      type: String,
      "default": ""
    }
  },
  data: function data2() {
    return {
      capture: true,
      history: {
        stack: [],
        offset: -1
      },
      lineNumbersHeight: "20px",
      codeData: ""
    };
  },
  watch: {
    modelValue: {
      immediate: true,
      handler: function handler(newVal) {
        if (!newVal) {
          this.codeData = "";
        } else {
          this.codeData = newVal;
        }
      }
    },
    content: {
      immediate: true,
      handler: function handler2() {
        var _this = this;
        if (this.lineNumbers) {
          this.$nextTick(function() {
            _this.setLineNumbersHeight();
          });
        }
      }
    },
    lineNumbers: function lineNumbers() {
      var _this2 = this;
      this.$nextTick(function() {
        _this2.styleLineNumbers();
        _this2.setLineNumbersHeight();
      });
    }
  },
  computed: {
    isEmpty: function isEmpty() {
      return this.codeData.length === 0;
    },
    content: function content() {
      var result = this.highlight(this.codeData) + "<br />";
      return result;
    },
    lineNumbersCount: function lineNumbersCount() {
      var totalLines = this.codeData.split(/\r\n|\n/).length;
      return totalLines;
    }
  },
  mounted: function mounted() {
    this._recordCurrentState();
    this.styleLineNumbers();
  },
  methods: {
    setLineNumbersHeight: function setLineNumbersHeight() {
      this.lineNumbersHeight = getComputedStyle(this.$refs.pre).height;
    },
    styleLineNumbers: function styleLineNumbers() {
      if (!this.lineNumbers || !this.autoStyleLineNumbers)
        return;
      var $editor = this.$refs.pre;
      var $lineNumbers = this.$el.querySelector(".prism-editor__line-numbers");
      var editorStyles = window.getComputedStyle($editor);
      this.$nextTick(function() {
        var btlr = "border-top-left-radius";
        var bblr = "border-bottom-left-radius";
        if (!$lineNumbers)
          return;
        $lineNumbers.style[btlr] = editorStyles[btlr];
        $lineNumbers.style[bblr] = editorStyles[bblr];
        $editor.style[btlr] = "0";
        $editor.style[bblr] = "0";
        var stylesList = ["background-color", "margin-top", "padding-top", "font-family", "font-size", "line-height"];
        stylesList.forEach(function(style) {
          $lineNumbers.style[style] = editorStyles[style];
        });
        $lineNumbers.style["margin-bottom"] = "-" + editorStyles["padding-top"];
      });
    },
    _recordCurrentState: function _recordCurrentState() {
      var input = this.$refs.textarea;
      if (!input)
        return;
      var value = input.value, selectionStart = input.selectionStart, selectionEnd = input.selectionEnd;
      this._recordChange({
        value,
        selectionStart,
        selectionEnd
      });
    },
    _getLines: function _getLines(text, position) {
      return text.substring(0, position).split("\n");
    },
    _applyEdits: function _applyEdits(record) {
      var input = this.$refs.textarea;
      var last = this.history.stack[this.history.offset];
      if (last && input) {
        this.history.stack[this.history.offset] = _extends({}, last, {
          selectionStart: input.selectionStart,
          selectionEnd: input.selectionEnd
        });
      }
      this._recordChange(record);
      this._updateInput(record);
    },
    _recordChange: function _recordChange(record, overwrite) {
      if (overwrite === void 0) {
        overwrite = false;
      }
      var _this$history = this.history, stack = _this$history.stack, offset2 = _this$history.offset;
      if (stack.length && offset2 > -1) {
        this.history.stack = stack.slice(0, offset2 + 1);
        var count = this.history.stack.length;
        if (count > HISTORY_LIMIT) {
          var extras = count - HISTORY_LIMIT;
          this.history.stack = stack.slice(extras, count);
          this.history.offset = Math.max(this.history.offset - extras, 0);
        }
      }
      var timestamp = Date.now();
      if (overwrite) {
        var last = this.history.stack[this.history.offset];
        if (last && timestamp - last.timestamp < HISTORY_TIME_GAP) {
          var _this$_getLines$pop, _this$_getLines$pop2;
          var re = /[^a-z0-9]([a-z0-9]+)$/i;
          var previous = (_this$_getLines$pop = this._getLines(last.value, last.selectionStart).pop()) === null || _this$_getLines$pop === void 0 ? void 0 : _this$_getLines$pop.match(re);
          var current = (_this$_getLines$pop2 = this._getLines(record.value, record.selectionStart).pop()) === null || _this$_getLines$pop2 === void 0 ? void 0 : _this$_getLines$pop2.match(re);
          if (previous && current && current[1].startsWith(previous[1])) {
            this.history.stack[this.history.offset] = _extends({}, record, {
              timestamp
            });
            return;
          }
        }
      }
      this.history.stack.push(_extends({}, record, {
        timestamp
      }));
      this.history.offset++;
    },
    _updateInput: function _updateInput(record) {
      var input = this.$refs.textarea;
      if (!input)
        return;
      input.value = record.value;
      input.selectionStart = record.selectionStart;
      input.selectionEnd = record.selectionEnd;
      this.$emit("update:modelValue", record.value);
    },
    handleChange: function handleChange(e) {
      var _e$target = e.target, value = _e$target.value, selectionStart = _e$target.selectionStart, selectionEnd = _e$target.selectionEnd;
      this._recordChange({
        value,
        selectionStart,
        selectionEnd
      }, true);
      this.$emit("update:modelValue", value);
    },
    _undoEdit: function _undoEdit() {
      var _this$history2 = this.history, stack = _this$history2.stack, offset2 = _this$history2.offset;
      var record = stack[offset2 - 1];
      if (record) {
        this._updateInput(record);
        this.history.offset = Math.max(offset2 - 1, 0);
      }
    },
    _redoEdit: function _redoEdit() {
      var _this$history3 = this.history, stack = _this$history3.stack, offset2 = _this$history3.offset;
      var record = stack[offset2 + 1];
      if (record) {
        this._updateInput(record);
        this.history.offset = Math.min(offset2 + 1, stack.length - 1);
      }
    },
    handleKeyDown: function handleKeyDown(e) {
      var tabSize = this.tabSize, insertSpaces = this.insertSpaces, ignoreTabKey = this.ignoreTabKey;
      this.$emit("keydown", e);
      if (e.defaultPrevented) {
        return;
      }
      if (e.keyCode === KEYCODE_ESCAPE) {
        e.target.blur();
        this.$emit("blur", e);
      }
      var _e$target2 = e.target, value = _e$target2.value, selectionStart = _e$target2.selectionStart, selectionEnd = _e$target2.selectionEnd;
      var tabCharacter = (insertSpaces ? " " : "	").repeat(tabSize);
      if (e.keyCode === KEYCODE_TAB && !ignoreTabKey && this.capture) {
        e.preventDefault();
        if (e.shiftKey) {
          var linesBeforeCaret = this._getLines(value, selectionStart);
          var startLine = linesBeforeCaret.length - 1;
          var endLine = this._getLines(value, selectionEnd).length - 1;
          var nextValue = value.split("\n").map(function(line2, i) {
            if (i >= startLine && i <= endLine && line2.startsWith(tabCharacter)) {
              return line2.substring(tabCharacter.length);
            }
            return line2;
          }).join("\n");
          if (value !== nextValue) {
            var startLineText = linesBeforeCaret[startLine];
            this._applyEdits({
              value: nextValue,
              selectionStart: startLineText.startsWith(tabCharacter) ? selectionStart - tabCharacter.length : selectionStart,
              selectionEnd: selectionEnd - (value.length - nextValue.length)
            });
          }
        } else if (selectionStart !== selectionEnd) {
          var _linesBeforeCaret = this._getLines(value, selectionStart);
          var _startLine = _linesBeforeCaret.length - 1;
          var _endLine = this._getLines(value, selectionEnd).length - 1;
          var _startLineText = _linesBeforeCaret[_startLine];
          this._applyEdits({
            value: value.split("\n").map(function(line2, i) {
              if (i >= _startLine && i <= _endLine) {
                return tabCharacter + line2;
              }
              return line2;
            }).join("\n"),
            selectionStart: /\S/.test(_startLineText) ? selectionStart + tabCharacter.length : selectionStart,
            selectionEnd: selectionEnd + tabCharacter.length * (_endLine - _startLine + 1)
          });
        } else {
          var updatedSelection = selectionStart + tabCharacter.length;
          this._applyEdits({
            value: value.substring(0, selectionStart) + tabCharacter + value.substring(selectionEnd),
            selectionStart: updatedSelection,
            selectionEnd: updatedSelection
          });
        }
      } else if (e.keyCode === KEYCODE_BACKSPACE) {
        var hasSelection = selectionStart !== selectionEnd;
        var textBeforeCaret = value.substring(0, selectionStart);
        if (textBeforeCaret.endsWith(tabCharacter) && !hasSelection) {
          e.preventDefault();
          var _updatedSelection = selectionStart - tabCharacter.length;
          this._applyEdits({
            value: value.substring(0, selectionStart - tabCharacter.length) + value.substring(selectionEnd),
            selectionStart: _updatedSelection,
            selectionEnd: _updatedSelection
          });
        }
      } else if (e.keyCode === KEYCODE_ENTER) {
        if (selectionStart === selectionEnd) {
          var line = this._getLines(value, selectionStart).pop();
          var matches = line === null || line === void 0 ? void 0 : line.match(/^\s+/);
          if (matches && matches[0]) {
            e.preventDefault();
            var indent = "\n" + matches[0];
            var _updatedSelection2 = selectionStart + indent.length;
            this._applyEdits({
              value: value.substring(0, selectionStart) + indent + value.substring(selectionEnd),
              selectionStart: _updatedSelection2,
              selectionEnd: _updatedSelection2
            });
          }
        }
      } else if (e.keyCode === KEYCODE_PARENS || e.keyCode === KEYCODE_BRACKETS || e.keyCode === KEYCODE_QUOTE || e.keyCode === KEYCODE_BACK_QUOTE) {
        var chars;
        if (e.keyCode === KEYCODE_PARENS && e.shiftKey) {
          chars = ["(", ")"];
        } else if (e.keyCode === KEYCODE_BRACKETS) {
          if (e.shiftKey) {
            chars = ["{", "}"];
          } else {
            chars = ["[", "]"];
          }
        } else if (e.keyCode === KEYCODE_QUOTE) {
          if (e.shiftKey) {
            chars = ['"', '"'];
          } else {
            chars = ["'", "'"];
          }
        } else if (e.keyCode === KEYCODE_BACK_QUOTE && !e.shiftKey) {
          chars = ["`", "`"];
        }
        if (selectionStart !== selectionEnd && chars) {
          e.preventDefault();
          this._applyEdits({
            value: value.substring(0, selectionStart) + chars[0] + value.substring(selectionStart, selectionEnd) + chars[1] + value.substring(selectionEnd),
            selectionStart,
            selectionEnd: selectionEnd + 2
          });
        }
      } else if ((isMacLike ? e.metaKey && e.keyCode === KEYCODE_Z : e.ctrlKey && e.keyCode === KEYCODE_Z) && !e.shiftKey && !e.altKey) {
        e.preventDefault();
        this._undoEdit();
      } else if ((isMacLike ? e.metaKey && e.keyCode === KEYCODE_Z && e.shiftKey : isWindows ? e.ctrlKey && e.keyCode === KEYCODE_Y : e.ctrlKey && e.keyCode === KEYCODE_Z && e.shiftKey) && !e.altKey) {
        e.preventDefault();
        this._redoEdit();
      } else if (e.keyCode === KEYCODE_M && e.ctrlKey && (isMacLike ? e.shiftKey : true)) {
        e.preventDefault();
        this.capture = !this.capture;
      }
    }
  },
  render: function render() {
    var _this3 = this;
    var lineNumberWidthCalculator = vue.h("div", {
      "class": "prism-editor__line-width-calc",
      style: "height: 0px; visibility: hidden; pointer-events: none;"
    }, "999");
    var lineNumbers2 = vue.h("div", {
      "class": "prism-editor__line-numbers",
      style: {
        "min-height": this.lineNumbersHeight
      },
      "aria-hidden": "true"
    }, [lineNumberWidthCalculator, Array.from(Array(this.lineNumbersCount).keys()).map(function(_, index2) {
      return vue.h("div", {
        "class": "prism-editor__line-number token comment"
      }, "" + ++index2);
    })]);
    var textarea = vue.h("textarea", {
      ref: "textarea",
      onInput: this.handleChange,
      onKeydown: this.handleKeyDown,
      onClick: function onClick($event) {
        _this3.$emit("click", $event);
      },
      onKeyup: function onKeyup($event) {
        _this3.$emit("keyup", $event);
      },
      onFocus: function onFocus($event) {
        _this3.$emit("focus", $event);
      },
      onBlur: function onBlur($event) {
        _this3.$emit("blur", $event);
      },
      "class": {
        "prism-editor__textarea": true,
        "prism-editor__textarea--empty": this.isEmpty
      },
      spellCheck: "false",
      autocapitalize: "off",
      autocomplete: "off",
      autocorrect: "off",
      "data-gramm": "false",
      placeholder: this.placeholder,
      "data-testid": "textarea",
      readonly: this.readonly,
      value: this.codeData
    });
    var preview = vue.h("pre", {
      ref: "pre",
      "class": "prism-editor__editor",
      "data-testid": "preview",
      innerHTML: this.content
    });
    var editorContainer = vue.h("div", {
      "class": "prism-editor__container"
    }, [textarea, preview]);
    return vue.h("div", {
      "class": "prism-editor-wrapper"
    }, [this.lineNumbers && lineNumbers2, editorContainer]);
  }
});
var commonjsGlobal = typeof globalThis !== "undefined" ? globalThis : typeof window !== "undefined" ? window : typeof global !== "undefined" ? global : typeof self !== "undefined" ? self : {};
var prismCore = { exports: {} };
(function(module) {
  var _self = typeof window !== "undefined" ? window : typeof WorkerGlobalScope !== "undefined" && self instanceof WorkerGlobalScope ? self : {};
  /**
   * Prism: Lightweight, robust, elegant syntax highlighting
   *
   * @license MIT <https://opensource.org/licenses/MIT>
   * @author Lea Verou <https://lea.verou.me>
   * @namespace
   * @public
   */
  var Prism2 = function(_self2) {
    var lang = /\blang(?:uage)?-([\w-]+)\b/i;
    var uniqueId = 0;
    var plainTextGrammar = {};
    var _ = {
      manual: _self2.Prism && _self2.Prism.manual,
      disableWorkerMessageHandler: _self2.Prism && _self2.Prism.disableWorkerMessageHandler,
      util: {
        encode: function encode(tokens) {
          if (tokens instanceof Token) {
            return new Token(tokens.type, encode(tokens.content), tokens.alias);
          } else if (Array.isArray(tokens)) {
            return tokens.map(encode);
          } else {
            return tokens.replace(/&/g, "&amp;").replace(/</g, "&lt;").replace(/\u00a0/g, " ");
          }
        },
        type: function(o) {
          return Object.prototype.toString.call(o).slice(8, -1);
        },
        objId: function(obj) {
          if (!obj["__id"]) {
            Object.defineProperty(obj, "__id", { value: ++uniqueId });
          }
          return obj["__id"];
        },
        clone: function deepClone(o, visited) {
          visited = visited || {};
          var clone2;
          var id;
          switch (_.util.type(o)) {
            case "Object":
              id = _.util.objId(o);
              if (visited[id]) {
                return visited[id];
              }
              clone2 = {};
              visited[id] = clone2;
              for (var key in o) {
                if (o.hasOwnProperty(key)) {
                  clone2[key] = deepClone(o[key], visited);
                }
              }
              return clone2;
            case "Array":
              id = _.util.objId(o);
              if (visited[id]) {
                return visited[id];
              }
              clone2 = [];
              visited[id] = clone2;
              o.forEach(function(v, i) {
                clone2[i] = deepClone(v, visited);
              });
              return clone2;
            default:
              return o;
          }
        },
        getLanguage: function(element) {
          while (element && !lang.test(element.className)) {
            element = element.parentElement;
          }
          if (element) {
            return (element.className.match(lang) || [, "none"])[1].toLowerCase();
          }
          return "none";
        },
        currentScript: function() {
          if (typeof document === "undefined") {
            return null;
          }
          if ("currentScript" in document && 1 < 2) {
            return document.currentScript;
          }
          try {
            throw new Error();
          } catch (err) {
            var src = (/at [^(\r\n]*\((.*):[^:]+:[^:]+\)$/i.exec(err.stack) || [])[1];
            if (src) {
              var scripts = document.getElementsByTagName("script");
              for (var i in scripts) {
                if (scripts[i].src == src) {
                  return scripts[i];
                }
              }
            }
            return null;
          }
        },
        isActive: function(element, className, defaultActivation) {
          var no = "no-" + className;
          while (element) {
            var classList = element.classList;
            if (classList.contains(className)) {
              return true;
            }
            if (classList.contains(no)) {
              return false;
            }
            element = element.parentElement;
          }
          return !!defaultActivation;
        }
      },
      languages: {
        plain: plainTextGrammar,
        plaintext: plainTextGrammar,
        text: plainTextGrammar,
        txt: plainTextGrammar,
        extend: function(id, redef) {
          var lang2 = _.util.clone(_.languages[id]);
          for (var key in redef) {
            lang2[key] = redef[key];
          }
          return lang2;
        },
        insertBefore: function(inside, before, insert, root) {
          root = root || _.languages;
          var grammar = root[inside];
          var ret = {};
          for (var token in grammar) {
            if (grammar.hasOwnProperty(token)) {
              if (token == before) {
                for (var newToken in insert) {
                  if (insert.hasOwnProperty(newToken)) {
                    ret[newToken] = insert[newToken];
                  }
                }
              }
              if (!insert.hasOwnProperty(token)) {
                ret[token] = grammar[token];
              }
            }
          }
          var old = root[inside];
          root[inside] = ret;
          _.languages.DFS(_.languages, function(key, value) {
            if (value === old && key != inside) {
              this[key] = ret;
            }
          });
          return ret;
        },
        DFS: function DFS(o, callback, type, visited) {
          visited = visited || {};
          var objId = _.util.objId;
          for (var i in o) {
            if (o.hasOwnProperty(i)) {
              callback.call(o, i, o[i], type || i);
              var property = o[i];
              var propertyType = _.util.type(property);
              if (propertyType === "Object" && !visited[objId(property)]) {
                visited[objId(property)] = true;
                DFS(property, callback, null, visited);
              } else if (propertyType === "Array" && !visited[objId(property)]) {
                visited[objId(property)] = true;
                DFS(property, callback, i, visited);
              }
            }
          }
        }
      },
      plugins: {},
      highlightAll: function(async, callback) {
        _.highlightAllUnder(document, async, callback);
      },
      highlightAllUnder: function(container, async, callback) {
        var env = {
          callback,
          container,
          selector: 'code[class*="language-"], [class*="language-"] code, code[class*="lang-"], [class*="lang-"] code'
        };
        _.hooks.run("before-highlightall", env);
        env.elements = Array.prototype.slice.apply(env.container.querySelectorAll(env.selector));
        _.hooks.run("before-all-elements-highlight", env);
        for (var i = 0, element; element = env.elements[i++]; ) {
          _.highlightElement(element, async === true, env.callback);
        }
      },
      highlightElement: function(element, async, callback) {
        var language = _.util.getLanguage(element);
        var grammar = _.languages[language];
        element.className = element.className.replace(lang, "").replace(/\s+/g, " ") + " language-" + language;
        var parent = element.parentElement;
        if (parent && parent.nodeName.toLowerCase() === "pre") {
          parent.className = parent.className.replace(lang, "").replace(/\s+/g, " ") + " language-" + language;
        }
        var code = element.textContent;
        var env = {
          element,
          language,
          grammar,
          code
        };
        function insertHighlightedCode(highlightedCode) {
          env.highlightedCode = highlightedCode;
          _.hooks.run("before-insert", env);
          env.element.innerHTML = env.highlightedCode;
          _.hooks.run("after-highlight", env);
          _.hooks.run("complete", env);
          callback && callback.call(env.element);
        }
        _.hooks.run("before-sanity-check", env);
        parent = env.element.parentElement;
        if (parent && parent.nodeName.toLowerCase() === "pre" && !parent.hasAttribute("tabindex")) {
          parent.setAttribute("tabindex", "0");
        }
        if (!env.code) {
          _.hooks.run("complete", env);
          callback && callback.call(env.element);
          return;
        }
        _.hooks.run("before-highlight", env);
        if (!env.grammar) {
          insertHighlightedCode(_.util.encode(env.code));
          return;
        }
        if (async && _self2.Worker) {
          var worker = new Worker(_.filename);
          worker.onmessage = function(evt) {
            insertHighlightedCode(evt.data);
          };
          worker.postMessage(JSON.stringify({
            language: env.language,
            code: env.code,
            immediateClose: true
          }));
        } else {
          insertHighlightedCode(_.highlight(env.code, env.grammar, env.language));
        }
      },
      highlight: function(text, grammar, language) {
        var env = {
          code: text,
          grammar,
          language
        };
        _.hooks.run("before-tokenize", env);
        env.tokens = _.tokenize(env.code, env.grammar);
        _.hooks.run("after-tokenize", env);
        return Token.stringify(_.util.encode(env.tokens), env.language);
      },
      tokenize: function(text, grammar) {
        var rest = grammar.rest;
        if (rest) {
          for (var token in rest) {
            grammar[token] = rest[token];
          }
          delete grammar.rest;
        }
        var tokenList = new LinkedList();
        addAfter(tokenList, tokenList.head, text);
        matchGrammar(text, tokenList, grammar, tokenList.head, 0);
        return toArray(tokenList);
      },
      hooks: {
        all: {},
        add: function(name, callback) {
          var hooks = _.hooks.all;
          hooks[name] = hooks[name] || [];
          hooks[name].push(callback);
        },
        run: function(name, env) {
          var callbacks = _.hooks.all[name];
          if (!callbacks || !callbacks.length) {
            return;
          }
          for (var i = 0, callback; callback = callbacks[i++]; ) {
            callback(env);
          }
        }
      },
      Token
    };
    _self2.Prism = _;
    function Token(type, content2, alias, matchedStr) {
      this.type = type;
      this.content = content2;
      this.alias = alias;
      this.length = (matchedStr || "").length | 0;
    }
    Token.stringify = function stringify2(o, language) {
      if (typeof o == "string") {
        return o;
      }
      if (Array.isArray(o)) {
        var s2 = "";
        o.forEach(function(e) {
          s2 += stringify2(e, language);
        });
        return s2;
      }
      var env = {
        type: o.type,
        content: stringify2(o.content, language),
        tag: "span",
        classes: ["token", o.type],
        attributes: {},
        language
      };
      var aliases = o.alias;
      if (aliases) {
        if (Array.isArray(aliases)) {
          Array.prototype.push.apply(env.classes, aliases);
        } else {
          env.classes.push(aliases);
        }
      }
      _.hooks.run("wrap", env);
      var attributes = "";
      for (var name in env.attributes) {
        attributes += " " + name + '="' + (env.attributes[name] || "").replace(/"/g, "&quot;") + '"';
      }
      return "<" + env.tag + ' class="' + env.classes.join(" ") + '"' + attributes + ">" + env.content + "</" + env.tag + ">";
    };
    function matchPattern(pattern, pos, text, lookbehind) {
      pattern.lastIndex = pos;
      var match2 = pattern.exec(text);
      if (match2 && lookbehind && match2[1]) {
        var lookbehindLength = match2[1].length;
        match2.index += lookbehindLength;
        match2[0] = match2[0].slice(lookbehindLength);
      }
      return match2;
    }
    function matchGrammar(text, tokenList, grammar, startNode, startPos, rematch) {
      for (var token in grammar) {
        if (!grammar.hasOwnProperty(token) || !grammar[token]) {
          continue;
        }
        var patterns = grammar[token];
        patterns = Array.isArray(patterns) ? patterns : [patterns];
        for (var j = 0; j < patterns.length; ++j) {
          if (rematch && rematch.cause == token + "," + j) {
            return;
          }
          var patternObj = patterns[j];
          var inside = patternObj.inside;
          var lookbehind = !!patternObj.lookbehind;
          var greedy = !!patternObj.greedy;
          var alias = patternObj.alias;
          if (greedy && !patternObj.pattern.global) {
            var flags = patternObj.pattern.toString().match(/[imsuy]*$/)[0];
            patternObj.pattern = RegExp(patternObj.pattern.source, flags + "g");
          }
          var pattern = patternObj.pattern || patternObj;
          for (var currentNode = startNode.next, pos = startPos; currentNode !== tokenList.tail; pos += currentNode.value.length, currentNode = currentNode.next) {
            if (rematch && pos >= rematch.reach) {
              break;
            }
            var str = currentNode.value;
            if (tokenList.length > text.length) {
              return;
            }
            if (str instanceof Token) {
              continue;
            }
            var removeCount = 1;
            var match2;
            if (greedy) {
              match2 = matchPattern(pattern, pos, text, lookbehind);
              if (!match2) {
                break;
              }
              var from = match2.index;
              var to = match2.index + match2[0].length;
              var p = pos;
              p += currentNode.value.length;
              while (from >= p) {
                currentNode = currentNode.next;
                p += currentNode.value.length;
              }
              p -= currentNode.value.length;
              pos = p;
              if (currentNode.value instanceof Token) {
                continue;
              }
              for (var k = currentNode; k !== tokenList.tail && (p < to || typeof k.value === "string"); k = k.next) {
                removeCount++;
                p += k.value.length;
              }
              removeCount--;
              str = text.slice(pos, p);
              match2.index -= pos;
            } else {
              match2 = matchPattern(pattern, 0, str, lookbehind);
              if (!match2) {
                continue;
              }
            }
            var from = match2.index;
            var matchStr = match2[0];
            var before = str.slice(0, from);
            var after = str.slice(from + matchStr.length);
            var reach = pos + str.length;
            if (rematch && reach > rematch.reach) {
              rematch.reach = reach;
            }
            var removeFrom = currentNode.prev;
            if (before) {
              removeFrom = addAfter(tokenList, removeFrom, before);
              pos += before.length;
            }
            removeRange(tokenList, removeFrom, removeCount);
            var wrapped = new Token(token, inside ? _.tokenize(matchStr, inside) : matchStr, alias, matchStr);
            currentNode = addAfter(tokenList, removeFrom, wrapped);
            if (after) {
              addAfter(tokenList, currentNode, after);
            }
            if (removeCount > 1) {
              var nestedRematch = {
                cause: token + "," + j,
                reach
              };
              matchGrammar(text, tokenList, grammar, currentNode.prev, pos, nestedRematch);
              if (rematch && nestedRematch.reach > rematch.reach) {
                rematch.reach = nestedRematch.reach;
              }
            }
          }
        }
      }
    }
    function LinkedList() {
      var head = { value: null, prev: null, next: null };
      var tail = { value: null, prev: head, next: null };
      head.next = tail;
      this.head = head;
      this.tail = tail;
      this.length = 0;
    }
    function addAfter(list, node, value) {
      var next = node.next;
      var newNode = { value, prev: node, next };
      node.next = newNode;
      next.prev = newNode;
      list.length++;
      return newNode;
    }
    function removeRange(list, node, count) {
      var next = node.next;
      for (var i = 0; i < count && next !== list.tail; i++) {
        next = next.next;
      }
      node.next = next;
      next.prev = node;
      list.length -= i;
    }
    function toArray(list) {
      var array = [];
      var node = list.head.next;
      while (node !== list.tail) {
        array.push(node.value);
        node = node.next;
      }
      return array;
    }
    if (!_self2.document) {
      if (!_self2.addEventListener) {
        return _;
      }
      if (!_.disableWorkerMessageHandler) {
        _self2.addEventListener("message", function(evt) {
          var message = JSON.parse(evt.data);
          var lang2 = message.language;
          var code = message.code;
          var immediateClose = message.immediateClose;
          _self2.postMessage(_.highlight(code, _.languages[lang2], lang2));
          if (immediateClose) {
            _self2.close();
          }
        }, false);
      }
      return _;
    }
    var script2 = _.util.currentScript();
    if (script2) {
      _.filename = script2.src;
      if (script2.hasAttribute("data-manual")) {
        _.manual = true;
      }
    }
    function highlightAutomaticallyCallback() {
      if (!_.manual) {
        _.highlightAll();
      }
    }
    if (!_.manual) {
      var readyState = document.readyState;
      if (readyState === "loading" || readyState === "interactive" && script2 && script2.defer) {
        document.addEventListener("DOMContentLoaded", highlightAutomaticallyCallback);
      } else {
        if (window.requestAnimationFrame) {
          window.requestAnimationFrame(highlightAutomaticallyCallback);
        } else {
          window.setTimeout(highlightAutomaticallyCallback, 16);
        }
      }
    }
    return _;
  }(_self);
  if (module.exports) {
    module.exports = Prism2;
  }
  if (typeof commonjsGlobal !== "undefined") {
    commonjsGlobal.Prism = Prism2;
  }
})(prismCore);
Prism.languages.clike = {
  "comment": [
    {
      pattern: /(^|[^\\])\/\*[\s\S]*?(?:\*\/|$)/,
      lookbehind: true,
      greedy: true
    },
    {
      pattern: /(^|[^\\:])\/\/.*/,
      lookbehind: true,
      greedy: true
    }
  ],
  "string": {
    pattern: /(["'])(?:\\(?:\r\n|[\s\S])|(?!\1)[^\\\r\n])*\1/,
    greedy: true
  },
  "class-name": {
    pattern: /(\b(?:class|interface|extends|implements|trait|instanceof|new)\s+|\bcatch\s+\()[\w.\\]+/i,
    lookbehind: true,
    inside: {
      "punctuation": /[.\\]/
    }
  },
  "keyword": /\b(?:if|else|while|do|for|return|in|instanceof|function|new|try|throw|catch|finally|null|break|continue)\b/,
  "boolean": /\b(?:true|false)\b/,
  "function": /\b\w+(?=\()/,
  "number": /\b0x[\da-f]+\b|(?:\b\d+(?:\.\d*)?|\B\.\d+)(?:e[+-]?\d+)?/i,
  "operator": /[<>]=?|[!=]=?=?|--?|\+\+?|&&?|\|\|?|[?*/~^%]/,
  "punctuation": /[{}[\];(),.:]/
};
(function(Prism2) {
  function simple_form(name) {
    return RegExp("(\\()" + name + "(?=[\\s\\)])");
  }
  function primitive(pattern) {
    return RegExp("([\\s([])" + pattern + "(?=[\\s)])");
  }
  var symbol = "[-+*/~!@$%^=<>{}\\w]+";
  var marker = "&" + symbol;
  var par = "(\\()";
  var endpar = "(?=\\))";
  var space = "(?=\\s)";
  var language = {
    heading: {
      pattern: /;;;.*/,
      alias: ["comment", "title"]
    },
    comment: /;.*/,
    string: {
      pattern: /"(?:[^"\\]|\\.)*"/,
      greedy: true,
      inside: {
        argument: /[-A-Z]+(?=[.,\s])/,
        symbol: RegExp("`" + symbol + "'")
      }
    },
    "quoted-symbol": {
      pattern: RegExp("#?'" + symbol),
      alias: ["variable", "symbol"]
    },
    "lisp-property": {
      pattern: RegExp(":" + symbol),
      alias: "property"
    },
    splice: {
      pattern: RegExp(",@?" + symbol),
      alias: ["symbol", "variable"]
    },
    keyword: [
      {
        pattern: RegExp(par + "(?:(?:lexical-)?let\\*?|(?:cl-)?letf|if|when|while|unless|cons|cl-loop|and|or|not|cond|setq|error|message|null|require|provide|use-package)" + space),
        lookbehind: true
      },
      {
        pattern: RegExp(par + "(?:for|do|collect|return|finally|append|concat|in|by)" + space),
        lookbehind: true
      }
    ],
    declare: {
      pattern: simple_form("declare"),
      lookbehind: true,
      alias: "keyword"
    },
    interactive: {
      pattern: simple_form("interactive"),
      lookbehind: true,
      alias: "keyword"
    },
    boolean: {
      pattern: primitive("(?:t|nil)"),
      lookbehind: true
    },
    number: {
      pattern: primitive("[-+]?\\d+(?:\\.\\d*)?"),
      lookbehind: true
    },
    defvar: {
      pattern: RegExp(par + "def(?:var|const|custom|group)\\s+" + symbol),
      lookbehind: true,
      inside: {
        keyword: /^def[a-z]+/,
        variable: RegExp(symbol)
      }
    },
    defun: {
      pattern: RegExp(par + "(?:cl-)?(?:defun\\*?|defmacro)\\s+" + symbol + "\\s+\\([\\s\\S]*?\\)"),
      lookbehind: true,
      inside: {
        keyword: /^(?:cl-)?def\S+/,
        arguments: null,
        function: {
          pattern: RegExp("(^\\s)" + symbol),
          lookbehind: true
        },
        punctuation: /[()]/
      }
    },
    lambda: {
      pattern: RegExp(par + "lambda\\s+\\(\\s*(?:&?" + symbol + "(?:\\s+&?" + symbol + ")*\\s*)?\\)"),
      lookbehind: true,
      inside: {
        keyword: /^lambda/,
        arguments: null,
        punctuation: /[()]/
      }
    },
    car: {
      pattern: RegExp(par + symbol),
      lookbehind: true
    },
    punctuation: [
      /(?:['`,]?\(|[)\[\]])/,
      {
        pattern: /(\s)\.(?=\s)/,
        lookbehind: true
      }
    ]
  };
  var arg = {
    "lisp-marker": RegExp(marker),
    rest: {
      argument: {
        pattern: RegExp(symbol),
        alias: "variable"
      },
      varform: {
        pattern: RegExp(par + symbol + "\\s+\\S[\\s\\S]*" + endpar),
        lookbehind: true,
        inside: {
          string: language.string,
          boolean: language.boolean,
          number: language.number,
          symbol: language.symbol,
          punctuation: /[()]/
        }
      }
    }
  };
  var forms = "\\S+(?:\\s+\\S+)*";
  var arglist = {
    pattern: RegExp(par + "[\\s\\S]*" + endpar),
    lookbehind: true,
    inside: {
      "rest-vars": {
        pattern: RegExp("&(?:rest|body)\\s+" + forms),
        inside: arg
      },
      "other-marker-vars": {
        pattern: RegExp("&(?:optional|aux)\\s+" + forms),
        inside: arg
      },
      keys: {
        pattern: RegExp("&key\\s+" + forms + "(?:\\s+&allow-other-keys)?"),
        inside: arg
      },
      argument: {
        pattern: RegExp(symbol),
        alias: "variable"
      },
      punctuation: /[()]/
    }
  };
  language["lambda"].inside.arguments = arglist;
  language["defun"].inside.arguments = Prism2.util.clone(arglist);
  language["defun"].inside.arguments.inside.sublist = arglist;
  Prism2.languages.lisp = language;
  Prism2.languages.elisp = language;
  Prism2.languages.emacs = language;
  Prism2.languages["emacs-lisp"] = language;
})(Prism);
var _export_sfc = (sfc, props) => {
  for (const [key, val] of props) {
    sfc[key] = val;
  }
  return sfc;
};
const _sfc_main$1$1 = {
  name: "csl-editor",
  props: {
    modelValue: {
      type: String,
      default: function() {
        return "";
      }
    },
    readonly: {
      type: Boolean,
      default: function() {
        return false;
      }
    }
  },
  components: {
    PrismEditor
  },
  emits: ["update:modelValue", "submit"],
  methods: {
    highlighter(code) {
      return prismCore.exports.highlight(code, prismCore.exports.languages.lisp, "lisp");
    },
    enterPressed(event) {
      event.stopImmediatePropagation();
      this.$emit("submit");
    }
  },
  computed: {
    computedValue: {
      get() {
        return this.modelValue;
      },
      set(newVal) {
        this.$emit("update:modelValue", newVal);
      }
    }
  }
};
function _sfc_render$1(_ctx, _cache, $props, $setup, $data, $options) {
  const _component_PrismEditor = vue.resolveComponent("PrismEditor");
  return vue.openBlock(), vue.createBlock(_component_PrismEditor, {
    "aria-label": "Edit CSL",
    label: "edit csl",
    class: "my-editor h-full",
    modelValue: $options.computedValue,
    "onUpdate:modelValue": _cache[0] || (_cache[0] = ($event) => $options.computedValue = $event),
    onKeydown: vue.withKeys(vue.withModifiers($options.enterPressed, ["ctrl"]), ["enter"]),
    highlight: $options.highlighter,
    readonly: $props.readonly,
    "aria-disabled": $props.readonly,
    "line-numbers": ""
  }, null, 8, ["modelValue", "onKeydown", "highlight", "readonly", "aria-disabled"]);
}
var CSLEditor = /* @__PURE__ */ _export_sfc(_sfc_main$1$1, [["render", _sfc_render$1], ["__scopeId", "data-v-7519217c"]]);
class LuxonError extends Error {
}
class InvalidDateTimeError extends LuxonError {
  constructor(reason) {
    super(`Invalid DateTime: ${reason.toMessage()}`);
  }
}
class InvalidIntervalError extends LuxonError {
  constructor(reason) {
    super(`Invalid Interval: ${reason.toMessage()}`);
  }
}
class InvalidDurationError extends LuxonError {
  constructor(reason) {
    super(`Invalid Duration: ${reason.toMessage()}`);
  }
}
class ConflictingSpecificationError extends LuxonError {
}
class InvalidUnitError extends LuxonError {
  constructor(unit) {
    super(`Invalid unit ${unit}`);
  }
}
class InvalidArgumentError extends LuxonError {
}
class ZoneIsAbstractError extends LuxonError {
  constructor() {
    super("Zone is an abstract class");
  }
}
const n = "numeric", s = "short", l = "long";
const DATE_SHORT = {
  year: n,
  month: n,
  day: n
};
const DATE_MED = {
  year: n,
  month: s,
  day: n
};
const DATE_MED_WITH_WEEKDAY = {
  year: n,
  month: s,
  day: n,
  weekday: s
};
const DATE_FULL = {
  year: n,
  month: l,
  day: n
};
const DATE_HUGE = {
  year: n,
  month: l,
  day: n,
  weekday: l
};
const TIME_SIMPLE = {
  hour: n,
  minute: n
};
const TIME_WITH_SECONDS = {
  hour: n,
  minute: n,
  second: n
};
const TIME_WITH_SHORT_OFFSET = {
  hour: n,
  minute: n,
  second: n,
  timeZoneName: s
};
const TIME_WITH_LONG_OFFSET = {
  hour: n,
  minute: n,
  second: n,
  timeZoneName: l
};
const TIME_24_SIMPLE = {
  hour: n,
  minute: n,
  hour12: false
};
const TIME_24_WITH_SECONDS = {
  hour: n,
  minute: n,
  second: n,
  hour12: false
};
const TIME_24_WITH_SHORT_OFFSET = {
  hour: n,
  minute: n,
  second: n,
  hour12: false,
  timeZoneName: s
};
const TIME_24_WITH_LONG_OFFSET = {
  hour: n,
  minute: n,
  second: n,
  hour12: false,
  timeZoneName: l
};
const DATETIME_SHORT = {
  year: n,
  month: n,
  day: n,
  hour: n,
  minute: n
};
const DATETIME_SHORT_WITH_SECONDS = {
  year: n,
  month: n,
  day: n,
  hour: n,
  minute: n,
  second: n
};
const DATETIME_MED = {
  year: n,
  month: s,
  day: n,
  hour: n,
  minute: n
};
const DATETIME_MED_WITH_SECONDS = {
  year: n,
  month: s,
  day: n,
  hour: n,
  minute: n,
  second: n
};
const DATETIME_MED_WITH_WEEKDAY = {
  year: n,
  month: s,
  day: n,
  weekday: s,
  hour: n,
  minute: n
};
const DATETIME_FULL = {
  year: n,
  month: l,
  day: n,
  hour: n,
  minute: n,
  timeZoneName: s
};
const DATETIME_FULL_WITH_SECONDS = {
  year: n,
  month: l,
  day: n,
  hour: n,
  minute: n,
  second: n,
  timeZoneName: s
};
const DATETIME_HUGE = {
  year: n,
  month: l,
  day: n,
  weekday: l,
  hour: n,
  minute: n,
  timeZoneName: l
};
const DATETIME_HUGE_WITH_SECONDS = {
  year: n,
  month: l,
  day: n,
  weekday: l,
  hour: n,
  minute: n,
  second: n,
  timeZoneName: l
};
function isUndefined(o) {
  return typeof o === "undefined";
}
function isNumber(o) {
  return typeof o === "number";
}
function isInteger(o) {
  return typeof o === "number" && o % 1 === 0;
}
function isString(o) {
  return typeof o === "string";
}
function isDate(o) {
  return Object.prototype.toString.call(o) === "[object Date]";
}
function hasIntl() {
  try {
    return typeof Intl !== "undefined" && Intl.DateTimeFormat;
  } catch (e) {
    return false;
  }
}
function hasFormatToParts() {
  return !isUndefined(Intl.DateTimeFormat.prototype.formatToParts);
}
function hasRelative() {
  try {
    return typeof Intl !== "undefined" && !!Intl.RelativeTimeFormat;
  } catch (e) {
    return false;
  }
}
function maybeArray(thing) {
  return Array.isArray(thing) ? thing : [thing];
}
function bestBy(arr, by, compare) {
  if (arr.length === 0) {
    return void 0;
  }
  return arr.reduce((best, next) => {
    const pair = [by(next), next];
    if (!best) {
      return pair;
    } else if (compare(best[0], pair[0]) === best[0]) {
      return best;
    } else {
      return pair;
    }
  }, null)[1];
}
function pick(obj, keys) {
  return keys.reduce((a, k) => {
    a[k] = obj[k];
    return a;
  }, {});
}
function hasOwnProperty(obj, prop) {
  return Object.prototype.hasOwnProperty.call(obj, prop);
}
function integerBetween(thing, bottom, top) {
  return isInteger(thing) && thing >= bottom && thing <= top;
}
function floorMod(x, n2) {
  return x - n2 * Math.floor(x / n2);
}
function padStart(input, n2 = 2) {
  const minus = input < 0 ? "-" : "";
  const target = minus ? input * -1 : input;
  let result;
  if (target.toString().length < n2) {
    result = ("0".repeat(n2) + target).slice(-n2);
  } else {
    result = target.toString();
  }
  return `${minus}${result}`;
}
function parseInteger(string) {
  if (isUndefined(string) || string === null || string === "") {
    return void 0;
  } else {
    return parseInt(string, 10);
  }
}
function parseMillis(fraction) {
  if (isUndefined(fraction) || fraction === null || fraction === "") {
    return void 0;
  } else {
    const f = parseFloat("0." + fraction) * 1e3;
    return Math.floor(f);
  }
}
function roundTo(number, digits, towardZero = false) {
  const factor = 10 ** digits, rounder = towardZero ? Math.trunc : Math.round;
  return rounder(number * factor) / factor;
}
function isLeapYear(year) {
  return year % 4 === 0 && (year % 100 !== 0 || year % 400 === 0);
}
function daysInYear(year) {
  return isLeapYear(year) ? 366 : 365;
}
function daysInMonth(year, month) {
  const modMonth = floorMod(month - 1, 12) + 1, modYear = year + (month - modMonth) / 12;
  if (modMonth === 2) {
    return isLeapYear(modYear) ? 29 : 28;
  } else {
    return [31, null, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31][modMonth - 1];
  }
}
function objToLocalTS(obj) {
  let d = Date.UTC(obj.year, obj.month - 1, obj.day, obj.hour, obj.minute, obj.second, obj.millisecond);
  if (obj.year < 100 && obj.year >= 0) {
    d = new Date(d);
    d.setUTCFullYear(d.getUTCFullYear() - 1900);
  }
  return +d;
}
function weeksInWeekYear(weekYear) {
  const p1 = (weekYear + Math.floor(weekYear / 4) - Math.floor(weekYear / 100) + Math.floor(weekYear / 400)) % 7, last = weekYear - 1, p2 = (last + Math.floor(last / 4) - Math.floor(last / 100) + Math.floor(last / 400)) % 7;
  return p1 === 4 || p2 === 3 ? 53 : 52;
}
function untruncateYear(year) {
  if (year > 99) {
    return year;
  } else
    return year > 60 ? 1900 + year : 2e3 + year;
}
function parseZoneInfo(ts, offsetFormat, locale, timeZone = null) {
  const date = new Date(ts), intlOpts = {
    hour12: false,
    year: "numeric",
    month: "2-digit",
    day: "2-digit",
    hour: "2-digit",
    minute: "2-digit"
  };
  if (timeZone) {
    intlOpts.timeZone = timeZone;
  }
  const modified = Object.assign({ timeZoneName: offsetFormat }, intlOpts), intl = hasIntl();
  if (intl && hasFormatToParts()) {
    const parsed = new Intl.DateTimeFormat(locale, modified).formatToParts(date).find((m) => m.type.toLowerCase() === "timezonename");
    return parsed ? parsed.value : null;
  } else if (intl) {
    const without = new Intl.DateTimeFormat(locale, intlOpts).format(date), included = new Intl.DateTimeFormat(locale, modified).format(date), diffed = included.substring(without.length), trimmed = diffed.replace(/^[, \u200e]+/, "");
    return trimmed;
  } else {
    return null;
  }
}
function signedOffset(offHourStr, offMinuteStr) {
  let offHour = parseInt(offHourStr, 10);
  if (Number.isNaN(offHour)) {
    offHour = 0;
  }
  const offMin = parseInt(offMinuteStr, 10) || 0, offMinSigned = offHour < 0 || Object.is(offHour, -0) ? -offMin : offMin;
  return offHour * 60 + offMinSigned;
}
function asNumber(value) {
  const numericValue = Number(value);
  if (typeof value === "boolean" || value === "" || Number.isNaN(numericValue))
    throw new InvalidArgumentError(`Invalid unit value ${value}`);
  return numericValue;
}
function normalizeObject(obj, normalizer, nonUnitKeys) {
  const normalized = {};
  for (const u in obj) {
    if (hasOwnProperty(obj, u)) {
      if (nonUnitKeys.indexOf(u) >= 0)
        continue;
      const v = obj[u];
      if (v === void 0 || v === null)
        continue;
      normalized[normalizer(u)] = asNumber(v);
    }
  }
  return normalized;
}
function formatOffset(offset2, format) {
  const hours = Math.trunc(Math.abs(offset2 / 60)), minutes = Math.trunc(Math.abs(offset2 % 60)), sign = offset2 >= 0 ? "+" : "-";
  switch (format) {
    case "short":
      return `${sign}${padStart(hours, 2)}:${padStart(minutes, 2)}`;
    case "narrow":
      return `${sign}${hours}${minutes > 0 ? `:${minutes}` : ""}`;
    case "techie":
      return `${sign}${padStart(hours, 2)}${padStart(minutes, 2)}`;
    default:
      throw new RangeError(`Value format ${format} is out of range for property format`);
  }
}
function timeObject(obj) {
  return pick(obj, ["hour", "minute", "second", "millisecond"]);
}
const ianaRegex = /[A-Za-z_+-]{1,256}(:?\/[A-Za-z_+-]{1,256}(\/[A-Za-z_+-]{1,256})?)?/;
function stringify(obj) {
  return JSON.stringify(obj, Object.keys(obj).sort());
}
const monthsLong = [
  "January",
  "February",
  "March",
  "April",
  "May",
  "June",
  "July",
  "August",
  "September",
  "October",
  "November",
  "December"
];
const monthsShort = [
  "Jan",
  "Feb",
  "Mar",
  "Apr",
  "May",
  "Jun",
  "Jul",
  "Aug",
  "Sep",
  "Oct",
  "Nov",
  "Dec"
];
const monthsNarrow = ["J", "F", "M", "A", "M", "J", "J", "A", "S", "O", "N", "D"];
function months(length) {
  switch (length) {
    case "narrow":
      return [...monthsNarrow];
    case "short":
      return [...monthsShort];
    case "long":
      return [...monthsLong];
    case "numeric":
      return ["1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12"];
    case "2-digit":
      return ["01", "02", "03", "04", "05", "06", "07", "08", "09", "10", "11", "12"];
    default:
      return null;
  }
}
const weekdaysLong = [
  "Monday",
  "Tuesday",
  "Wednesday",
  "Thursday",
  "Friday",
  "Saturday",
  "Sunday"
];
const weekdaysShort = ["Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"];
const weekdaysNarrow = ["M", "T", "W", "T", "F", "S", "S"];
function weekdays(length) {
  switch (length) {
    case "narrow":
      return [...weekdaysNarrow];
    case "short":
      return [...weekdaysShort];
    case "long":
      return [...weekdaysLong];
    case "numeric":
      return ["1", "2", "3", "4", "5", "6", "7"];
    default:
      return null;
  }
}
const meridiems = ["AM", "PM"];
const erasLong = ["Before Christ", "Anno Domini"];
const erasShort = ["BC", "AD"];
const erasNarrow = ["B", "A"];
function eras(length) {
  switch (length) {
    case "narrow":
      return [...erasNarrow];
    case "short":
      return [...erasShort];
    case "long":
      return [...erasLong];
    default:
      return null;
  }
}
function meridiemForDateTime(dt) {
  return meridiems[dt.hour < 12 ? 0 : 1];
}
function weekdayForDateTime(dt, length) {
  return weekdays(length)[dt.weekday - 1];
}
function monthForDateTime(dt, length) {
  return months(length)[dt.month - 1];
}
function eraForDateTime(dt, length) {
  return eras(length)[dt.year < 0 ? 0 : 1];
}
function formatRelativeTime(unit, count, numeric = "always", narrow = false) {
  const units = {
    years: ["year", "yr."],
    quarters: ["quarter", "qtr."],
    months: ["month", "mo."],
    weeks: ["week", "wk."],
    days: ["day", "day", "days"],
    hours: ["hour", "hr."],
    minutes: ["minute", "min."],
    seconds: ["second", "sec."]
  };
  const lastable = ["hours", "minutes", "seconds"].indexOf(unit) === -1;
  if (numeric === "auto" && lastable) {
    const isDay = unit === "days";
    switch (count) {
      case 1:
        return isDay ? "tomorrow" : `next ${units[unit][0]}`;
      case -1:
        return isDay ? "yesterday" : `last ${units[unit][0]}`;
      case 0:
        return isDay ? "today" : `this ${units[unit][0]}`;
    }
  }
  const isInPast = Object.is(count, -0) || count < 0, fmtValue = Math.abs(count), singular = fmtValue === 1, lilUnits = units[unit], fmtUnit = narrow ? singular ? lilUnits[1] : lilUnits[2] || lilUnits[1] : singular ? units[unit][0] : unit;
  return isInPast ? `${fmtValue} ${fmtUnit} ago` : `in ${fmtValue} ${fmtUnit}`;
}
function formatString(knownFormat) {
  const filtered = pick(knownFormat, [
    "weekday",
    "era",
    "year",
    "month",
    "day",
    "hour",
    "minute",
    "second",
    "timeZoneName",
    "hour12"
  ]), key = stringify(filtered), dateTimeHuge = "EEEE, LLLL d, yyyy, h:mm a";
  switch (key) {
    case stringify(DATE_SHORT):
      return "M/d/yyyy";
    case stringify(DATE_MED):
      return "LLL d, yyyy";
    case stringify(DATE_MED_WITH_WEEKDAY):
      return "EEE, LLL d, yyyy";
    case stringify(DATE_FULL):
      return "LLLL d, yyyy";
    case stringify(DATE_HUGE):
      return "EEEE, LLLL d, yyyy";
    case stringify(TIME_SIMPLE):
      return "h:mm a";
    case stringify(TIME_WITH_SECONDS):
      return "h:mm:ss a";
    case stringify(TIME_WITH_SHORT_OFFSET):
      return "h:mm a";
    case stringify(TIME_WITH_LONG_OFFSET):
      return "h:mm a";
    case stringify(TIME_24_SIMPLE):
      return "HH:mm";
    case stringify(TIME_24_WITH_SECONDS):
      return "HH:mm:ss";
    case stringify(TIME_24_WITH_SHORT_OFFSET):
      return "HH:mm";
    case stringify(TIME_24_WITH_LONG_OFFSET):
      return "HH:mm";
    case stringify(DATETIME_SHORT):
      return "M/d/yyyy, h:mm a";
    case stringify(DATETIME_MED):
      return "LLL d, yyyy, h:mm a";
    case stringify(DATETIME_FULL):
      return "LLLL d, yyyy, h:mm a";
    case stringify(DATETIME_HUGE):
      return dateTimeHuge;
    case stringify(DATETIME_SHORT_WITH_SECONDS):
      return "M/d/yyyy, h:mm:ss a";
    case stringify(DATETIME_MED_WITH_SECONDS):
      return "LLL d, yyyy, h:mm:ss a";
    case stringify(DATETIME_MED_WITH_WEEKDAY):
      return "EEE, d LLL yyyy, h:mm a";
    case stringify(DATETIME_FULL_WITH_SECONDS):
      return "LLLL d, yyyy, h:mm:ss a";
    case stringify(DATETIME_HUGE_WITH_SECONDS):
      return "EEEE, LLLL d, yyyy, h:mm:ss a";
    default:
      return dateTimeHuge;
  }
}
function stringifyTokens(splits, tokenToString) {
  let s2 = "";
  for (const token of splits) {
    if (token.literal) {
      s2 += token.val;
    } else {
      s2 += tokenToString(token.val);
    }
  }
  return s2;
}
const macroTokenToFormatOpts = {
  D: DATE_SHORT,
  DD: DATE_MED,
  DDD: DATE_FULL,
  DDDD: DATE_HUGE,
  t: TIME_SIMPLE,
  tt: TIME_WITH_SECONDS,
  ttt: TIME_WITH_SHORT_OFFSET,
  tttt: TIME_WITH_LONG_OFFSET,
  T: TIME_24_SIMPLE,
  TT: TIME_24_WITH_SECONDS,
  TTT: TIME_24_WITH_SHORT_OFFSET,
  TTTT: TIME_24_WITH_LONG_OFFSET,
  f: DATETIME_SHORT,
  ff: DATETIME_MED,
  fff: DATETIME_FULL,
  ffff: DATETIME_HUGE,
  F: DATETIME_SHORT_WITH_SECONDS,
  FF: DATETIME_MED_WITH_SECONDS,
  FFF: DATETIME_FULL_WITH_SECONDS,
  FFFF: DATETIME_HUGE_WITH_SECONDS
};
class Formatter {
  static create(locale, opts = {}) {
    return new Formatter(locale, opts);
  }
  static parseFormat(fmt) {
    let current = null, currentFull = "", bracketed = false;
    const splits = [];
    for (let i = 0; i < fmt.length; i++) {
      const c = fmt.charAt(i);
      if (c === "'") {
        if (currentFull.length > 0) {
          splits.push({ literal: bracketed, val: currentFull });
        }
        current = null;
        currentFull = "";
        bracketed = !bracketed;
      } else if (bracketed) {
        currentFull += c;
      } else if (c === current) {
        currentFull += c;
      } else {
        if (currentFull.length > 0) {
          splits.push({ literal: false, val: currentFull });
        }
        currentFull = c;
        current = c;
      }
    }
    if (currentFull.length > 0) {
      splits.push({ literal: bracketed, val: currentFull });
    }
    return splits;
  }
  static macroTokenToFormatOpts(token) {
    return macroTokenToFormatOpts[token];
  }
  constructor(locale, formatOpts) {
    this.opts = formatOpts;
    this.loc = locale;
    this.systemLoc = null;
  }
  formatWithSystemDefault(dt, opts) {
    if (this.systemLoc === null) {
      this.systemLoc = this.loc.redefaultToSystem();
    }
    const df = this.systemLoc.dtFormatter(dt, Object.assign({}, this.opts, opts));
    return df.format();
  }
  formatDateTime(dt, opts = {}) {
    const df = this.loc.dtFormatter(dt, Object.assign({}, this.opts, opts));
    return df.format();
  }
  formatDateTimeParts(dt, opts = {}) {
    const df = this.loc.dtFormatter(dt, Object.assign({}, this.opts, opts));
    return df.formatToParts();
  }
  resolvedOptions(dt, opts = {}) {
    const df = this.loc.dtFormatter(dt, Object.assign({}, this.opts, opts));
    return df.resolvedOptions();
  }
  num(n2, p = 0) {
    if (this.opts.forceSimple) {
      return padStart(n2, p);
    }
    const opts = Object.assign({}, this.opts);
    if (p > 0) {
      opts.padTo = p;
    }
    return this.loc.numberFormatter(opts).format(n2);
  }
  formatDateTimeFromString(dt, fmt) {
    const knownEnglish = this.loc.listingMode() === "en", useDateTimeFormatter = this.loc.outputCalendar && this.loc.outputCalendar !== "gregory" && hasFormatToParts(), string = (opts, extract) => this.loc.extract(dt, opts, extract), formatOffset2 = (opts) => {
      if (dt.isOffsetFixed && dt.offset === 0 && opts.allowZ) {
        return "Z";
      }
      return dt.isValid ? dt.zone.formatOffset(dt.ts, opts.format) : "";
    }, meridiem = () => knownEnglish ? meridiemForDateTime(dt) : string({ hour: "numeric", hour12: true }, "dayperiod"), month = (length, standalone) => knownEnglish ? monthForDateTime(dt, length) : string(standalone ? { month: length } : { month: length, day: "numeric" }, "month"), weekday = (length, standalone) => knownEnglish ? weekdayForDateTime(dt, length) : string(standalone ? { weekday: length } : { weekday: length, month: "long", day: "numeric" }, "weekday"), maybeMacro = (token) => {
      const formatOpts = Formatter.macroTokenToFormatOpts(token);
      if (formatOpts) {
        return this.formatWithSystemDefault(dt, formatOpts);
      } else {
        return token;
      }
    }, era = (length) => knownEnglish ? eraForDateTime(dt, length) : string({ era: length }, "era"), tokenToString = (token) => {
      switch (token) {
        case "S":
          return this.num(dt.millisecond);
        case "u":
        case "SSS":
          return this.num(dt.millisecond, 3);
        case "s":
          return this.num(dt.second);
        case "ss":
          return this.num(dt.second, 2);
        case "m":
          return this.num(dt.minute);
        case "mm":
          return this.num(dt.minute, 2);
        case "h":
          return this.num(dt.hour % 12 === 0 ? 12 : dt.hour % 12);
        case "hh":
          return this.num(dt.hour % 12 === 0 ? 12 : dt.hour % 12, 2);
        case "H":
          return this.num(dt.hour);
        case "HH":
          return this.num(dt.hour, 2);
        case "Z":
          return formatOffset2({ format: "narrow", allowZ: this.opts.allowZ });
        case "ZZ":
          return formatOffset2({ format: "short", allowZ: this.opts.allowZ });
        case "ZZZ":
          return formatOffset2({ format: "techie", allowZ: this.opts.allowZ });
        case "ZZZZ":
          return dt.zone.offsetName(dt.ts, { format: "short", locale: this.loc.locale });
        case "ZZZZZ":
          return dt.zone.offsetName(dt.ts, { format: "long", locale: this.loc.locale });
        case "z":
          return dt.zoneName;
        case "a":
          return meridiem();
        case "d":
          return useDateTimeFormatter ? string({ day: "numeric" }, "day") : this.num(dt.day);
        case "dd":
          return useDateTimeFormatter ? string({ day: "2-digit" }, "day") : this.num(dt.day, 2);
        case "c":
          return this.num(dt.weekday);
        case "ccc":
          return weekday("short", true);
        case "cccc":
          return weekday("long", true);
        case "ccccc":
          return weekday("narrow", true);
        case "E":
          return this.num(dt.weekday);
        case "EEE":
          return weekday("short", false);
        case "EEEE":
          return weekday("long", false);
        case "EEEEE":
          return weekday("narrow", false);
        case "L":
          return useDateTimeFormatter ? string({ month: "numeric", day: "numeric" }, "month") : this.num(dt.month);
        case "LL":
          return useDateTimeFormatter ? string({ month: "2-digit", day: "numeric" }, "month") : this.num(dt.month, 2);
        case "LLL":
          return month("short", true);
        case "LLLL":
          return month("long", true);
        case "LLLLL":
          return month("narrow", true);
        case "M":
          return useDateTimeFormatter ? string({ month: "numeric" }, "month") : this.num(dt.month);
        case "MM":
          return useDateTimeFormatter ? string({ month: "2-digit" }, "month") : this.num(dt.month, 2);
        case "MMM":
          return month("short", false);
        case "MMMM":
          return month("long", false);
        case "MMMMM":
          return month("narrow", false);
        case "y":
          return useDateTimeFormatter ? string({ year: "numeric" }, "year") : this.num(dt.year);
        case "yy":
          return useDateTimeFormatter ? string({ year: "2-digit" }, "year") : this.num(dt.year.toString().slice(-2), 2);
        case "yyyy":
          return useDateTimeFormatter ? string({ year: "numeric" }, "year") : this.num(dt.year, 4);
        case "yyyyyy":
          return useDateTimeFormatter ? string({ year: "numeric" }, "year") : this.num(dt.year, 6);
        case "G":
          return era("short");
        case "GG":
          return era("long");
        case "GGGGG":
          return era("narrow");
        case "kk":
          return this.num(dt.weekYear.toString().slice(-2), 2);
        case "kkkk":
          return this.num(dt.weekYear, 4);
        case "W":
          return this.num(dt.weekNumber);
        case "WW":
          return this.num(dt.weekNumber, 2);
        case "o":
          return this.num(dt.ordinal);
        case "ooo":
          return this.num(dt.ordinal, 3);
        case "q":
          return this.num(dt.quarter);
        case "qq":
          return this.num(dt.quarter, 2);
        case "X":
          return this.num(Math.floor(dt.ts / 1e3));
        case "x":
          return this.num(dt.ts);
        default:
          return maybeMacro(token);
      }
    };
    return stringifyTokens(Formatter.parseFormat(fmt), tokenToString);
  }
  formatDurationFromString(dur, fmt) {
    const tokenToField = (token) => {
      switch (token[0]) {
        case "S":
          return "millisecond";
        case "s":
          return "second";
        case "m":
          return "minute";
        case "h":
          return "hour";
        case "d":
          return "day";
        case "M":
          return "month";
        case "y":
          return "year";
        default:
          return null;
      }
    }, tokenToString = (lildur) => (token) => {
      const mapped = tokenToField(token);
      if (mapped) {
        return this.num(lildur.get(mapped), token.length);
      } else {
        return token;
      }
    }, tokens = Formatter.parseFormat(fmt), realTokens = tokens.reduce((found, { literal, val }) => literal ? found : found.concat(val), []), collapsed = dur.shiftTo(...realTokens.map(tokenToField).filter((t) => t));
    return stringifyTokens(tokens, tokenToString(collapsed));
  }
}
class Invalid {
  constructor(reason, explanation) {
    this.reason = reason;
    this.explanation = explanation;
  }
  toMessage() {
    if (this.explanation) {
      return `${this.reason}: ${this.explanation}`;
    } else {
      return this.reason;
    }
  }
}
class Zone {
  get type() {
    throw new ZoneIsAbstractError();
  }
  get name() {
    throw new ZoneIsAbstractError();
  }
  get universal() {
    throw new ZoneIsAbstractError();
  }
  offsetName(ts, opts) {
    throw new ZoneIsAbstractError();
  }
  formatOffset(ts, format) {
    throw new ZoneIsAbstractError();
  }
  offset(ts) {
    throw new ZoneIsAbstractError();
  }
  equals(otherZone) {
    throw new ZoneIsAbstractError();
  }
  get isValid() {
    throw new ZoneIsAbstractError();
  }
}
let singleton$1 = null;
class LocalZone extends Zone {
  static get instance() {
    if (singleton$1 === null) {
      singleton$1 = new LocalZone();
    }
    return singleton$1;
  }
  get type() {
    return "local";
  }
  get name() {
    if (hasIntl()) {
      return new Intl.DateTimeFormat().resolvedOptions().timeZone;
    } else
      return "local";
  }
  get universal() {
    return false;
  }
  offsetName(ts, { format, locale }) {
    return parseZoneInfo(ts, format, locale);
  }
  formatOffset(ts, format) {
    return formatOffset(this.offset(ts), format);
  }
  offset(ts) {
    return -new Date(ts).getTimezoneOffset();
  }
  equals(otherZone) {
    return otherZone.type === "local";
  }
  get isValid() {
    return true;
  }
}
const matchingRegex = RegExp(`^${ianaRegex.source}$`);
let dtfCache = {};
function makeDTF(zone) {
  if (!dtfCache[zone]) {
    dtfCache[zone] = new Intl.DateTimeFormat("en-US", {
      hour12: false,
      timeZone: zone,
      year: "numeric",
      month: "2-digit",
      day: "2-digit",
      hour: "2-digit",
      minute: "2-digit",
      second: "2-digit"
    });
  }
  return dtfCache[zone];
}
const typeToPos = {
  year: 0,
  month: 1,
  day: 2,
  hour: 3,
  minute: 4,
  second: 5
};
function hackyOffset(dtf, date) {
  const formatted = dtf.format(date).replace(/\u200E/g, ""), parsed = /(\d+)\/(\d+)\/(\d+),? (\d+):(\d+):(\d+)/.exec(formatted), [, fMonth, fDay, fYear, fHour, fMinute, fSecond] = parsed;
  return [fYear, fMonth, fDay, fHour, fMinute, fSecond];
}
function partsOffset(dtf, date) {
  const formatted = dtf.formatToParts(date), filled = [];
  for (let i = 0; i < formatted.length; i++) {
    const { type, value } = formatted[i], pos = typeToPos[type];
    if (!isUndefined(pos)) {
      filled[pos] = parseInt(value, 10);
    }
  }
  return filled;
}
let ianaZoneCache = {};
class IANAZone extends Zone {
  static create(name) {
    if (!ianaZoneCache[name]) {
      ianaZoneCache[name] = new IANAZone(name);
    }
    return ianaZoneCache[name];
  }
  static resetCache() {
    ianaZoneCache = {};
    dtfCache = {};
  }
  static isValidSpecifier(s2) {
    return !!(s2 && s2.match(matchingRegex));
  }
  static isValidZone(zone) {
    try {
      new Intl.DateTimeFormat("en-US", { timeZone: zone }).format();
      return true;
    } catch (e) {
      return false;
    }
  }
  static parseGMTOffset(specifier) {
    if (specifier) {
      const match2 = specifier.match(/^Etc\/GMT(0|[+-]\d{1,2})$/i);
      if (match2) {
        return -60 * parseInt(match2[1]);
      }
    }
    return null;
  }
  constructor(name) {
    super();
    this.zoneName = name;
    this.valid = IANAZone.isValidZone(name);
  }
  get type() {
    return "iana";
  }
  get name() {
    return this.zoneName;
  }
  get universal() {
    return false;
  }
  offsetName(ts, { format, locale }) {
    return parseZoneInfo(ts, format, locale, this.name);
  }
  formatOffset(ts, format) {
    return formatOffset(this.offset(ts), format);
  }
  offset(ts) {
    const date = new Date(ts);
    if (isNaN(date))
      return NaN;
    const dtf = makeDTF(this.name), [year, month, day, hour, minute, second] = dtf.formatToParts ? partsOffset(dtf, date) : hackyOffset(dtf, date), adjustedHour = hour === 24 ? 0 : hour;
    const asUTC = objToLocalTS({
      year,
      month,
      day,
      hour: adjustedHour,
      minute,
      second,
      millisecond: 0
    });
    let asTS = +date;
    const over = asTS % 1e3;
    asTS -= over >= 0 ? over : 1e3 + over;
    return (asUTC - asTS) / (60 * 1e3);
  }
  equals(otherZone) {
    return otherZone.type === "iana" && otherZone.name === this.name;
  }
  get isValid() {
    return this.valid;
  }
}
let singleton = null;
class FixedOffsetZone extends Zone {
  static get utcInstance() {
    if (singleton === null) {
      singleton = new FixedOffsetZone(0);
    }
    return singleton;
  }
  static instance(offset2) {
    return offset2 === 0 ? FixedOffsetZone.utcInstance : new FixedOffsetZone(offset2);
  }
  static parseSpecifier(s2) {
    if (s2) {
      const r2 = s2.match(/^utc(?:([+-]\d{1,2})(?::(\d{2}))?)?$/i);
      if (r2) {
        return new FixedOffsetZone(signedOffset(r2[1], r2[2]));
      }
    }
    return null;
  }
  constructor(offset2) {
    super();
    this.fixed = offset2;
  }
  get type() {
    return "fixed";
  }
  get name() {
    return this.fixed === 0 ? "UTC" : `UTC${formatOffset(this.fixed, "narrow")}`;
  }
  offsetName() {
    return this.name;
  }
  formatOffset(ts, format) {
    return formatOffset(this.fixed, format);
  }
  get universal() {
    return true;
  }
  offset() {
    return this.fixed;
  }
  equals(otherZone) {
    return otherZone.type === "fixed" && otherZone.fixed === this.fixed;
  }
  get isValid() {
    return true;
  }
}
class InvalidZone extends Zone {
  constructor(zoneName) {
    super();
    this.zoneName = zoneName;
  }
  get type() {
    return "invalid";
  }
  get name() {
    return this.zoneName;
  }
  get universal() {
    return false;
  }
  offsetName() {
    return null;
  }
  formatOffset() {
    return "";
  }
  offset() {
    return NaN;
  }
  equals() {
    return false;
  }
  get isValid() {
    return false;
  }
}
function normalizeZone(input, defaultZone2) {
  let offset2;
  if (isUndefined(input) || input === null) {
    return defaultZone2;
  } else if (input instanceof Zone) {
    return input;
  } else if (isString(input)) {
    const lowered = input.toLowerCase();
    if (lowered === "local")
      return defaultZone2;
    else if (lowered === "utc" || lowered === "gmt")
      return FixedOffsetZone.utcInstance;
    else if ((offset2 = IANAZone.parseGMTOffset(input)) != null) {
      return FixedOffsetZone.instance(offset2);
    } else if (IANAZone.isValidSpecifier(lowered))
      return IANAZone.create(input);
    else
      return FixedOffsetZone.parseSpecifier(lowered) || new InvalidZone(input);
  } else if (isNumber(input)) {
    return FixedOffsetZone.instance(input);
  } else if (typeof input === "object" && input.offset && typeof input.offset === "number") {
    return input;
  } else {
    return new InvalidZone(input);
  }
}
let now = () => Date.now(), defaultZone = null, defaultLocale = null, defaultNumberingSystem = null, defaultOutputCalendar = null, throwOnInvalid = false;
class Settings {
  static get now() {
    return now;
  }
  static set now(n2) {
    now = n2;
  }
  static get defaultZoneName() {
    return Settings.defaultZone.name;
  }
  static set defaultZoneName(z) {
    if (!z) {
      defaultZone = null;
    } else {
      defaultZone = normalizeZone(z);
    }
  }
  static get defaultZone() {
    return defaultZone || LocalZone.instance;
  }
  static get defaultLocale() {
    return defaultLocale;
  }
  static set defaultLocale(locale) {
    defaultLocale = locale;
  }
  static get defaultNumberingSystem() {
    return defaultNumberingSystem;
  }
  static set defaultNumberingSystem(numberingSystem) {
    defaultNumberingSystem = numberingSystem;
  }
  static get defaultOutputCalendar() {
    return defaultOutputCalendar;
  }
  static set defaultOutputCalendar(outputCalendar) {
    defaultOutputCalendar = outputCalendar;
  }
  static get throwOnInvalid() {
    return throwOnInvalid;
  }
  static set throwOnInvalid(t) {
    throwOnInvalid = t;
  }
  static resetCaches() {
    Locale.resetCache();
    IANAZone.resetCache();
  }
}
let intlDTCache = {};
function getCachedDTF(locString, opts = {}) {
  const key = JSON.stringify([locString, opts]);
  let dtf = intlDTCache[key];
  if (!dtf) {
    dtf = new Intl.DateTimeFormat(locString, opts);
    intlDTCache[key] = dtf;
  }
  return dtf;
}
let intlNumCache = {};
function getCachedINF(locString, opts = {}) {
  const key = JSON.stringify([locString, opts]);
  let inf = intlNumCache[key];
  if (!inf) {
    inf = new Intl.NumberFormat(locString, opts);
    intlNumCache[key] = inf;
  }
  return inf;
}
let intlRelCache = {};
function getCachedRTF(locString, opts = {}) {
  const _a = opts, cacheKeyOpts = __objRest(_a, ["base"]);
  const key = JSON.stringify([locString, cacheKeyOpts]);
  let inf = intlRelCache[key];
  if (!inf) {
    inf = new Intl.RelativeTimeFormat(locString, opts);
    intlRelCache[key] = inf;
  }
  return inf;
}
let sysLocaleCache = null;
function systemLocale() {
  if (sysLocaleCache) {
    return sysLocaleCache;
  } else if (hasIntl()) {
    const computedSys = new Intl.DateTimeFormat().resolvedOptions().locale;
    sysLocaleCache = !computedSys || computedSys === "und" ? "en-US" : computedSys;
    return sysLocaleCache;
  } else {
    sysLocaleCache = "en-US";
    return sysLocaleCache;
  }
}
function parseLocaleString(localeStr) {
  const uIndex = localeStr.indexOf("-u-");
  if (uIndex === -1) {
    return [localeStr];
  } else {
    let options;
    const smaller = localeStr.substring(0, uIndex);
    try {
      options = getCachedDTF(localeStr).resolvedOptions();
    } catch (e) {
      options = getCachedDTF(smaller).resolvedOptions();
    }
    const { numberingSystem, calendar } = options;
    return [smaller, numberingSystem, calendar];
  }
}
function intlConfigString(localeStr, numberingSystem, outputCalendar) {
  if (hasIntl()) {
    if (outputCalendar || numberingSystem) {
      localeStr += "-u";
      if (outputCalendar) {
        localeStr += `-ca-${outputCalendar}`;
      }
      if (numberingSystem) {
        localeStr += `-nu-${numberingSystem}`;
      }
      return localeStr;
    } else {
      return localeStr;
    }
  } else {
    return [];
  }
}
function mapMonths(f) {
  const ms = [];
  for (let i = 1; i <= 12; i++) {
    const dt = DateTime.utc(2016, i, 1);
    ms.push(f(dt));
  }
  return ms;
}
function mapWeekdays(f) {
  const ms = [];
  for (let i = 1; i <= 7; i++) {
    const dt = DateTime.utc(2016, 11, 13 + i);
    ms.push(f(dt));
  }
  return ms;
}
function listStuff(loc, length, defaultOK, englishFn, intlFn) {
  const mode = loc.listingMode(defaultOK);
  if (mode === "error") {
    return null;
  } else if (mode === "en") {
    return englishFn(length);
  } else {
    return intlFn(length);
  }
}
function supportsFastNumbers(loc) {
  if (loc.numberingSystem && loc.numberingSystem !== "latn") {
    return false;
  } else {
    return loc.numberingSystem === "latn" || !loc.locale || loc.locale.startsWith("en") || hasIntl() && new Intl.DateTimeFormat(loc.intl).resolvedOptions().numberingSystem === "latn";
  }
}
class PolyNumberFormatter {
  constructor(intl, forceSimple, opts) {
    this.padTo = opts.padTo || 0;
    this.floor = opts.floor || false;
    if (!forceSimple && hasIntl()) {
      const intlOpts = { useGrouping: false };
      if (opts.padTo > 0)
        intlOpts.minimumIntegerDigits = opts.padTo;
      this.inf = getCachedINF(intl, intlOpts);
    }
  }
  format(i) {
    if (this.inf) {
      const fixed = this.floor ? Math.floor(i) : i;
      return this.inf.format(fixed);
    } else {
      const fixed = this.floor ? Math.floor(i) : roundTo(i, 3);
      return padStart(fixed, this.padTo);
    }
  }
}
class PolyDateFormatter {
  constructor(dt, intl, opts) {
    this.opts = opts;
    this.hasIntl = hasIntl();
    let z;
    if (dt.zone.universal && this.hasIntl) {
      const gmtOffset = -1 * (dt.offset / 60);
      const offsetZ = gmtOffset >= 0 ? `Etc/GMT+${gmtOffset}` : `Etc/GMT${gmtOffset}`;
      const isOffsetZoneSupported = IANAZone.isValidZone(offsetZ);
      if (dt.offset !== 0 && isOffsetZoneSupported) {
        z = offsetZ;
        this.dt = dt;
      } else {
        z = "UTC";
        if (opts.timeZoneName) {
          this.dt = dt;
        } else {
          this.dt = dt.offset === 0 ? dt : DateTime.fromMillis(dt.ts + dt.offset * 60 * 1e3);
        }
      }
    } else if (dt.zone.type === "local") {
      this.dt = dt;
    } else {
      this.dt = dt;
      z = dt.zone.name;
    }
    if (this.hasIntl) {
      const intlOpts = Object.assign({}, this.opts);
      if (z) {
        intlOpts.timeZone = z;
      }
      this.dtf = getCachedDTF(intl, intlOpts);
    }
  }
  format() {
    if (this.hasIntl) {
      return this.dtf.format(this.dt.toJSDate());
    } else {
      const tokenFormat = formatString(this.opts), loc = Locale.create("en-US");
      return Formatter.create(loc).formatDateTimeFromString(this.dt, tokenFormat);
    }
  }
  formatToParts() {
    if (this.hasIntl && hasFormatToParts()) {
      return this.dtf.formatToParts(this.dt.toJSDate());
    } else {
      return [];
    }
  }
  resolvedOptions() {
    if (this.hasIntl) {
      return this.dtf.resolvedOptions();
    } else {
      return {
        locale: "en-US",
        numberingSystem: "latn",
        outputCalendar: "gregory"
      };
    }
  }
}
class PolyRelFormatter {
  constructor(intl, isEnglish, opts) {
    this.opts = Object.assign({ style: "long" }, opts);
    if (!isEnglish && hasRelative()) {
      this.rtf = getCachedRTF(intl, opts);
    }
  }
  format(count, unit) {
    if (this.rtf) {
      return this.rtf.format(count, unit);
    } else {
      return formatRelativeTime(unit, count, this.opts.numeric, this.opts.style !== "long");
    }
  }
  formatToParts(count, unit) {
    if (this.rtf) {
      return this.rtf.formatToParts(count, unit);
    } else {
      return [];
    }
  }
}
class Locale {
  static fromOpts(opts) {
    return Locale.create(opts.locale, opts.numberingSystem, opts.outputCalendar, opts.defaultToEN);
  }
  static create(locale, numberingSystem, outputCalendar, defaultToEN = false) {
    const specifiedLocale = locale || Settings.defaultLocale, localeR = specifiedLocale || (defaultToEN ? "en-US" : systemLocale()), numberingSystemR = numberingSystem || Settings.defaultNumberingSystem, outputCalendarR = outputCalendar || Settings.defaultOutputCalendar;
    return new Locale(localeR, numberingSystemR, outputCalendarR, specifiedLocale);
  }
  static resetCache() {
    sysLocaleCache = null;
    intlDTCache = {};
    intlNumCache = {};
    intlRelCache = {};
  }
  static fromObject({ locale, numberingSystem, outputCalendar } = {}) {
    return Locale.create(locale, numberingSystem, outputCalendar);
  }
  constructor(locale, numbering, outputCalendar, specifiedLocale) {
    const [parsedLocale, parsedNumberingSystem, parsedOutputCalendar] = parseLocaleString(locale);
    this.locale = parsedLocale;
    this.numberingSystem = numbering || parsedNumberingSystem || null;
    this.outputCalendar = outputCalendar || parsedOutputCalendar || null;
    this.intl = intlConfigString(this.locale, this.numberingSystem, this.outputCalendar);
    this.weekdaysCache = { format: {}, standalone: {} };
    this.monthsCache = { format: {}, standalone: {} };
    this.meridiemCache = null;
    this.eraCache = {};
    this.specifiedLocale = specifiedLocale;
    this.fastNumbersCached = null;
  }
  get fastNumbers() {
    if (this.fastNumbersCached == null) {
      this.fastNumbersCached = supportsFastNumbers(this);
    }
    return this.fastNumbersCached;
  }
  listingMode(defaultOK = true) {
    const intl = hasIntl(), hasFTP = intl && hasFormatToParts(), isActuallyEn = this.isEnglish(), hasNoWeirdness = (this.numberingSystem === null || this.numberingSystem === "latn") && (this.outputCalendar === null || this.outputCalendar === "gregory");
    if (!hasFTP && !(isActuallyEn && hasNoWeirdness) && !defaultOK) {
      return "error";
    } else if (!hasFTP || isActuallyEn && hasNoWeirdness) {
      return "en";
    } else {
      return "intl";
    }
  }
  clone(alts) {
    if (!alts || Object.getOwnPropertyNames(alts).length === 0) {
      return this;
    } else {
      return Locale.create(alts.locale || this.specifiedLocale, alts.numberingSystem || this.numberingSystem, alts.outputCalendar || this.outputCalendar, alts.defaultToEN || false);
    }
  }
  redefaultToEN(alts = {}) {
    return this.clone(Object.assign({}, alts, { defaultToEN: true }));
  }
  redefaultToSystem(alts = {}) {
    return this.clone(Object.assign({}, alts, { defaultToEN: false }));
  }
  months(length, format = false, defaultOK = true) {
    return listStuff(this, length, defaultOK, months, () => {
      const intl = format ? { month: length, day: "numeric" } : { month: length }, formatStr = format ? "format" : "standalone";
      if (!this.monthsCache[formatStr][length]) {
        this.monthsCache[formatStr][length] = mapMonths((dt) => this.extract(dt, intl, "month"));
      }
      return this.monthsCache[formatStr][length];
    });
  }
  weekdays(length, format = false, defaultOK = true) {
    return listStuff(this, length, defaultOK, weekdays, () => {
      const intl = format ? { weekday: length, year: "numeric", month: "long", day: "numeric" } : { weekday: length }, formatStr = format ? "format" : "standalone";
      if (!this.weekdaysCache[formatStr][length]) {
        this.weekdaysCache[formatStr][length] = mapWeekdays((dt) => this.extract(dt, intl, "weekday"));
      }
      return this.weekdaysCache[formatStr][length];
    });
  }
  meridiems(defaultOK = true) {
    return listStuff(this, void 0, defaultOK, () => meridiems, () => {
      if (!this.meridiemCache) {
        const intl = { hour: "numeric", hour12: true };
        this.meridiemCache = [DateTime.utc(2016, 11, 13, 9), DateTime.utc(2016, 11, 13, 19)].map((dt) => this.extract(dt, intl, "dayperiod"));
      }
      return this.meridiemCache;
    });
  }
  eras(length, defaultOK = true) {
    return listStuff(this, length, defaultOK, eras, () => {
      const intl = { era: length };
      if (!this.eraCache[length]) {
        this.eraCache[length] = [DateTime.utc(-40, 1, 1), DateTime.utc(2017, 1, 1)].map((dt) => this.extract(dt, intl, "era"));
      }
      return this.eraCache[length];
    });
  }
  extract(dt, intlOpts, field) {
    const df = this.dtFormatter(dt, intlOpts), results = df.formatToParts(), matching = results.find((m) => m.type.toLowerCase() === field);
    return matching ? matching.value : null;
  }
  numberFormatter(opts = {}) {
    return new PolyNumberFormatter(this.intl, opts.forceSimple || this.fastNumbers, opts);
  }
  dtFormatter(dt, intlOpts = {}) {
    return new PolyDateFormatter(dt, this.intl, intlOpts);
  }
  relFormatter(opts = {}) {
    return new PolyRelFormatter(this.intl, this.isEnglish(), opts);
  }
  isEnglish() {
    return this.locale === "en" || this.locale.toLowerCase() === "en-us" || hasIntl() && new Intl.DateTimeFormat(this.intl).resolvedOptions().locale.startsWith("en-us");
  }
  equals(other) {
    return this.locale === other.locale && this.numberingSystem === other.numberingSystem && this.outputCalendar === other.outputCalendar;
  }
}
function combineRegexes(...regexes) {
  const full = regexes.reduce((f, r2) => f + r2.source, "");
  return RegExp(`^${full}$`);
}
function combineExtractors(...extractors) {
  return (m) => extractors.reduce(([mergedVals, mergedZone, cursor], ex) => {
    const [val, zone, next] = ex(m, cursor);
    return [Object.assign(mergedVals, val), mergedZone || zone, next];
  }, [{}, null, 1]).slice(0, 2);
}
function parse(s2, ...patterns) {
  if (s2 == null) {
    return [null, null];
  }
  for (const [regex, extractor] of patterns) {
    const m = regex.exec(s2);
    if (m) {
      return extractor(m);
    }
  }
  return [null, null];
}
function simpleParse(...keys) {
  return (match2, cursor) => {
    const ret = {};
    let i;
    for (i = 0; i < keys.length; i++) {
      ret[keys[i]] = parseInteger(match2[cursor + i]);
    }
    return [ret, null, cursor + i];
  };
}
const offsetRegex = /(?:(Z)|([+-]\d\d)(?::?(\d\d))?)/, isoTimeBaseRegex = /(\d\d)(?::?(\d\d)(?::?(\d\d)(?:[.,](\d{1,30}))?)?)?/, isoTimeRegex = RegExp(`${isoTimeBaseRegex.source}${offsetRegex.source}?`), isoTimeExtensionRegex = RegExp(`(?:T${isoTimeRegex.source})?`), isoYmdRegex = /([+-]\d{6}|\d{4})(?:-?(\d\d)(?:-?(\d\d))?)?/, isoWeekRegex = /(\d{4})-?W(\d\d)(?:-?(\d))?/, isoOrdinalRegex = /(\d{4})-?(\d{3})/, extractISOWeekData = simpleParse("weekYear", "weekNumber", "weekDay"), extractISOOrdinalData = simpleParse("year", "ordinal"), sqlYmdRegex = /(\d{4})-(\d\d)-(\d\d)/, sqlTimeRegex = RegExp(`${isoTimeBaseRegex.source} ?(?:${offsetRegex.source}|(${ianaRegex.source}))?`), sqlTimeExtensionRegex = RegExp(`(?: ${sqlTimeRegex.source})?`);
function int(match2, pos, fallback) {
  const m = match2[pos];
  return isUndefined(m) ? fallback : parseInteger(m);
}
function extractISOYmd(match2, cursor) {
  const item = {
    year: int(match2, cursor),
    month: int(match2, cursor + 1, 1),
    day: int(match2, cursor + 2, 1)
  };
  return [item, null, cursor + 3];
}
function extractISOTime(match2, cursor) {
  const item = {
    hours: int(match2, cursor, 0),
    minutes: int(match2, cursor + 1, 0),
    seconds: int(match2, cursor + 2, 0),
    milliseconds: parseMillis(match2[cursor + 3])
  };
  return [item, null, cursor + 4];
}
function extractISOOffset(match2, cursor) {
  const local = !match2[cursor] && !match2[cursor + 1], fullOffset = signedOffset(match2[cursor + 1], match2[cursor + 2]), zone = local ? null : FixedOffsetZone.instance(fullOffset);
  return [{}, zone, cursor + 3];
}
function extractIANAZone(match2, cursor) {
  const zone = match2[cursor] ? IANAZone.create(match2[cursor]) : null;
  return [{}, zone, cursor + 1];
}
const isoTimeOnly = RegExp(`^T?${isoTimeBaseRegex.source}$`);
const isoDuration = /^-?P(?:(?:(-?\d{1,9})Y)?(?:(-?\d{1,9})M)?(?:(-?\d{1,9})W)?(?:(-?\d{1,9})D)?(?:T(?:(-?\d{1,9})H)?(?:(-?\d{1,9})M)?(?:(-?\d{1,20})(?:[.,](-?\d{1,9}))?S)?)?)$/;
function extractISODuration(match2) {
  const [
    s2,
    yearStr,
    monthStr,
    weekStr,
    dayStr,
    hourStr,
    minuteStr,
    secondStr,
    millisecondsStr
  ] = match2;
  const hasNegativePrefix = s2[0] === "-";
  const negativeSeconds = secondStr && secondStr[0] === "-";
  const maybeNegate = (num, force = false) => num !== void 0 && (force || num && hasNegativePrefix) ? -num : num;
  return [
    {
      years: maybeNegate(parseInteger(yearStr)),
      months: maybeNegate(parseInteger(monthStr)),
      weeks: maybeNegate(parseInteger(weekStr)),
      days: maybeNegate(parseInteger(dayStr)),
      hours: maybeNegate(parseInteger(hourStr)),
      minutes: maybeNegate(parseInteger(minuteStr)),
      seconds: maybeNegate(parseInteger(secondStr), secondStr === "-0"),
      milliseconds: maybeNegate(parseMillis(millisecondsStr), negativeSeconds)
    }
  ];
}
const obsOffsets = {
  GMT: 0,
  EDT: -4 * 60,
  EST: -5 * 60,
  CDT: -5 * 60,
  CST: -6 * 60,
  MDT: -6 * 60,
  MST: -7 * 60,
  PDT: -7 * 60,
  PST: -8 * 60
};
function fromStrings(weekdayStr, yearStr, monthStr, dayStr, hourStr, minuteStr, secondStr) {
  const result = {
    year: yearStr.length === 2 ? untruncateYear(parseInteger(yearStr)) : parseInteger(yearStr),
    month: monthsShort.indexOf(monthStr) + 1,
    day: parseInteger(dayStr),
    hour: parseInteger(hourStr),
    minute: parseInteger(minuteStr)
  };
  if (secondStr)
    result.second = parseInteger(secondStr);
  if (weekdayStr) {
    result.weekday = weekdayStr.length > 3 ? weekdaysLong.indexOf(weekdayStr) + 1 : weekdaysShort.indexOf(weekdayStr) + 1;
  }
  return result;
}
const rfc2822 = /^(?:(Mon|Tue|Wed|Thu|Fri|Sat|Sun),\s)?(\d{1,2})\s(Jan|Feb|Mar|Apr|May|Jun|Jul|Aug|Sep|Oct|Nov|Dec)\s(\d{2,4})\s(\d\d):(\d\d)(?::(\d\d))?\s(?:(UT|GMT|[ECMP][SD]T)|([Zz])|(?:([+-]\d\d)(\d\d)))$/;
function extractRFC2822(match2) {
  const [
    ,
    weekdayStr,
    dayStr,
    monthStr,
    yearStr,
    hourStr,
    minuteStr,
    secondStr,
    obsOffset,
    milOffset,
    offHourStr,
    offMinuteStr
  ] = match2, result = fromStrings(weekdayStr, yearStr, monthStr, dayStr, hourStr, minuteStr, secondStr);
  let offset2;
  if (obsOffset) {
    offset2 = obsOffsets[obsOffset];
  } else if (milOffset) {
    offset2 = 0;
  } else {
    offset2 = signedOffset(offHourStr, offMinuteStr);
  }
  return [result, new FixedOffsetZone(offset2)];
}
function preprocessRFC2822(s2) {
  return s2.replace(/\([^)]*\)|[\n\t]/g, " ").replace(/(\s\s+)/g, " ").trim();
}
const rfc1123 = /^(Mon|Tue|Wed|Thu|Fri|Sat|Sun), (\d\d) (Jan|Feb|Mar|Apr|May|Jun|Jul|Aug|Sep|Oct|Nov|Dec) (\d{4}) (\d\d):(\d\d):(\d\d) GMT$/, rfc850 = /^(Monday|Tuesday|Wedsday|Thursday|Friday|Saturday|Sunday), (\d\d)-(Jan|Feb|Mar|Apr|May|Jun|Jul|Aug|Sep|Oct|Nov|Dec)-(\d\d) (\d\d):(\d\d):(\d\d) GMT$/, ascii = /^(Mon|Tue|Wed|Thu|Fri|Sat|Sun) (Jan|Feb|Mar|Apr|May|Jun|Jul|Aug|Sep|Oct|Nov|Dec) ( \d|\d\d) (\d\d):(\d\d):(\d\d) (\d{4})$/;
function extractRFC1123Or850(match2) {
  const [, weekdayStr, dayStr, monthStr, yearStr, hourStr, minuteStr, secondStr] = match2, result = fromStrings(weekdayStr, yearStr, monthStr, dayStr, hourStr, minuteStr, secondStr);
  return [result, FixedOffsetZone.utcInstance];
}
function extractASCII(match2) {
  const [, weekdayStr, monthStr, dayStr, hourStr, minuteStr, secondStr, yearStr] = match2, result = fromStrings(weekdayStr, yearStr, monthStr, dayStr, hourStr, minuteStr, secondStr);
  return [result, FixedOffsetZone.utcInstance];
}
const isoYmdWithTimeExtensionRegex = combineRegexes(isoYmdRegex, isoTimeExtensionRegex);
const isoWeekWithTimeExtensionRegex = combineRegexes(isoWeekRegex, isoTimeExtensionRegex);
const isoOrdinalWithTimeExtensionRegex = combineRegexes(isoOrdinalRegex, isoTimeExtensionRegex);
const isoTimeCombinedRegex = combineRegexes(isoTimeRegex);
const extractISOYmdTimeAndOffset = combineExtractors(extractISOYmd, extractISOTime, extractISOOffset);
const extractISOWeekTimeAndOffset = combineExtractors(extractISOWeekData, extractISOTime, extractISOOffset);
const extractISOOrdinalDateAndTime = combineExtractors(extractISOOrdinalData, extractISOTime, extractISOOffset);
const extractISOTimeAndOffset = combineExtractors(extractISOTime, extractISOOffset);
function parseISODate(s2) {
  return parse(s2, [isoYmdWithTimeExtensionRegex, extractISOYmdTimeAndOffset], [isoWeekWithTimeExtensionRegex, extractISOWeekTimeAndOffset], [isoOrdinalWithTimeExtensionRegex, extractISOOrdinalDateAndTime], [isoTimeCombinedRegex, extractISOTimeAndOffset]);
}
function parseRFC2822Date(s2) {
  return parse(preprocessRFC2822(s2), [rfc2822, extractRFC2822]);
}
function parseHTTPDate(s2) {
  return parse(s2, [rfc1123, extractRFC1123Or850], [rfc850, extractRFC1123Or850], [ascii, extractASCII]);
}
function parseISODuration(s2) {
  return parse(s2, [isoDuration, extractISODuration]);
}
const extractISOTimeOnly = combineExtractors(extractISOTime);
function parseISOTimeOnly(s2) {
  return parse(s2, [isoTimeOnly, extractISOTimeOnly]);
}
const sqlYmdWithTimeExtensionRegex = combineRegexes(sqlYmdRegex, sqlTimeExtensionRegex);
const sqlTimeCombinedRegex = combineRegexes(sqlTimeRegex);
const extractISOYmdTimeOffsetAndIANAZone = combineExtractors(extractISOYmd, extractISOTime, extractISOOffset, extractIANAZone);
const extractISOTimeOffsetAndIANAZone = combineExtractors(extractISOTime, extractISOOffset, extractIANAZone);
function parseSQL(s2) {
  return parse(s2, [sqlYmdWithTimeExtensionRegex, extractISOYmdTimeOffsetAndIANAZone], [sqlTimeCombinedRegex, extractISOTimeOffsetAndIANAZone]);
}
const INVALID$2 = "Invalid Duration";
const lowOrderMatrix = {
  weeks: {
    days: 7,
    hours: 7 * 24,
    minutes: 7 * 24 * 60,
    seconds: 7 * 24 * 60 * 60,
    milliseconds: 7 * 24 * 60 * 60 * 1e3
  },
  days: {
    hours: 24,
    minutes: 24 * 60,
    seconds: 24 * 60 * 60,
    milliseconds: 24 * 60 * 60 * 1e3
  },
  hours: { minutes: 60, seconds: 60 * 60, milliseconds: 60 * 60 * 1e3 },
  minutes: { seconds: 60, milliseconds: 60 * 1e3 },
  seconds: { milliseconds: 1e3 }
}, casualMatrix = Object.assign({
  years: {
    quarters: 4,
    months: 12,
    weeks: 52,
    days: 365,
    hours: 365 * 24,
    minutes: 365 * 24 * 60,
    seconds: 365 * 24 * 60 * 60,
    milliseconds: 365 * 24 * 60 * 60 * 1e3
  },
  quarters: {
    months: 3,
    weeks: 13,
    days: 91,
    hours: 91 * 24,
    minutes: 91 * 24 * 60,
    seconds: 91 * 24 * 60 * 60,
    milliseconds: 91 * 24 * 60 * 60 * 1e3
  },
  months: {
    weeks: 4,
    days: 30,
    hours: 30 * 24,
    minutes: 30 * 24 * 60,
    seconds: 30 * 24 * 60 * 60,
    milliseconds: 30 * 24 * 60 * 60 * 1e3
  }
}, lowOrderMatrix), daysInYearAccurate = 146097 / 400, daysInMonthAccurate = 146097 / 4800, accurateMatrix = Object.assign({
  years: {
    quarters: 4,
    months: 12,
    weeks: daysInYearAccurate / 7,
    days: daysInYearAccurate,
    hours: daysInYearAccurate * 24,
    minutes: daysInYearAccurate * 24 * 60,
    seconds: daysInYearAccurate * 24 * 60 * 60,
    milliseconds: daysInYearAccurate * 24 * 60 * 60 * 1e3
  },
  quarters: {
    months: 3,
    weeks: daysInYearAccurate / 28,
    days: daysInYearAccurate / 4,
    hours: daysInYearAccurate * 24 / 4,
    minutes: daysInYearAccurate * 24 * 60 / 4,
    seconds: daysInYearAccurate * 24 * 60 * 60 / 4,
    milliseconds: daysInYearAccurate * 24 * 60 * 60 * 1e3 / 4
  },
  months: {
    weeks: daysInMonthAccurate / 7,
    days: daysInMonthAccurate,
    hours: daysInMonthAccurate * 24,
    minutes: daysInMonthAccurate * 24 * 60,
    seconds: daysInMonthAccurate * 24 * 60 * 60,
    milliseconds: daysInMonthAccurate * 24 * 60 * 60 * 1e3
  }
}, lowOrderMatrix);
const orderedUnits$1 = [
  "years",
  "quarters",
  "months",
  "weeks",
  "days",
  "hours",
  "minutes",
  "seconds",
  "milliseconds"
];
const reverseUnits = orderedUnits$1.slice(0).reverse();
function clone$1(dur, alts, clear = false) {
  const conf = {
    values: clear ? alts.values : Object.assign({}, dur.values, alts.values || {}),
    loc: dur.loc.clone(alts.loc),
    conversionAccuracy: alts.conversionAccuracy || dur.conversionAccuracy
  };
  return new Duration(conf);
}
function antiTrunc(n2) {
  return n2 < 0 ? Math.floor(n2) : Math.ceil(n2);
}
function convert(matrix, fromMap, fromUnit, toMap, toUnit) {
  const conv = matrix[toUnit][fromUnit], raw = fromMap[fromUnit] / conv, sameSign = Math.sign(raw) === Math.sign(toMap[toUnit]), added = !sameSign && toMap[toUnit] !== 0 && Math.abs(raw) <= 1 ? antiTrunc(raw) : Math.trunc(raw);
  toMap[toUnit] += added;
  fromMap[fromUnit] -= added * conv;
}
function normalizeValues(matrix, vals) {
  reverseUnits.reduce((previous, current) => {
    if (!isUndefined(vals[current])) {
      if (previous) {
        convert(matrix, vals, previous, vals, current);
      }
      return current;
    } else {
      return previous;
    }
  }, null);
}
class Duration {
  constructor(config2) {
    const accurate = config2.conversionAccuracy === "longterm" || false;
    this.values = config2.values;
    this.loc = config2.loc || Locale.create();
    this.conversionAccuracy = accurate ? "longterm" : "casual";
    this.invalid = config2.invalid || null;
    this.matrix = accurate ? accurateMatrix : casualMatrix;
    this.isLuxonDuration = true;
  }
  static fromMillis(count, opts) {
    return Duration.fromObject(Object.assign({ milliseconds: count }, opts));
  }
  static fromObject(obj) {
    if (obj == null || typeof obj !== "object") {
      throw new InvalidArgumentError(`Duration.fromObject: argument expected to be an object, got ${obj === null ? "null" : typeof obj}`);
    }
    return new Duration({
      values: normalizeObject(obj, Duration.normalizeUnit, [
        "locale",
        "numberingSystem",
        "conversionAccuracy",
        "zone"
      ]),
      loc: Locale.fromObject(obj),
      conversionAccuracy: obj.conversionAccuracy
    });
  }
  static fromISO(text, opts) {
    const [parsed] = parseISODuration(text);
    if (parsed) {
      const obj = Object.assign(parsed, opts);
      return Duration.fromObject(obj);
    } else {
      return Duration.invalid("unparsable", `the input "${text}" can't be parsed as ISO 8601`);
    }
  }
  static fromISOTime(text, opts) {
    const [parsed] = parseISOTimeOnly(text);
    if (parsed) {
      const obj = Object.assign(parsed, opts);
      return Duration.fromObject(obj);
    } else {
      return Duration.invalid("unparsable", `the input "${text}" can't be parsed as ISO 8601`);
    }
  }
  static invalid(reason, explanation = null) {
    if (!reason) {
      throw new InvalidArgumentError("need to specify a reason the Duration is invalid");
    }
    const invalid = reason instanceof Invalid ? reason : new Invalid(reason, explanation);
    if (Settings.throwOnInvalid) {
      throw new InvalidDurationError(invalid);
    } else {
      return new Duration({ invalid });
    }
  }
  static normalizeUnit(unit) {
    const normalized = {
      year: "years",
      years: "years",
      quarter: "quarters",
      quarters: "quarters",
      month: "months",
      months: "months",
      week: "weeks",
      weeks: "weeks",
      day: "days",
      days: "days",
      hour: "hours",
      hours: "hours",
      minute: "minutes",
      minutes: "minutes",
      second: "seconds",
      seconds: "seconds",
      millisecond: "milliseconds",
      milliseconds: "milliseconds"
    }[unit ? unit.toLowerCase() : unit];
    if (!normalized)
      throw new InvalidUnitError(unit);
    return normalized;
  }
  static isDuration(o) {
    return o && o.isLuxonDuration || false;
  }
  get locale() {
    return this.isValid ? this.loc.locale : null;
  }
  get numberingSystem() {
    return this.isValid ? this.loc.numberingSystem : null;
  }
  toFormat(fmt, opts = {}) {
    const fmtOpts = Object.assign({}, opts, {
      floor: opts.round !== false && opts.floor !== false
    });
    return this.isValid ? Formatter.create(this.loc, fmtOpts).formatDurationFromString(this, fmt) : INVALID$2;
  }
  toObject(opts = {}) {
    if (!this.isValid)
      return {};
    const base = Object.assign({}, this.values);
    if (opts.includeConfig) {
      base.conversionAccuracy = this.conversionAccuracy;
      base.numberingSystem = this.loc.numberingSystem;
      base.locale = this.loc.locale;
    }
    return base;
  }
  toISO() {
    if (!this.isValid)
      return null;
    let s2 = "P";
    if (this.years !== 0)
      s2 += this.years + "Y";
    if (this.months !== 0 || this.quarters !== 0)
      s2 += this.months + this.quarters * 3 + "M";
    if (this.weeks !== 0)
      s2 += this.weeks + "W";
    if (this.days !== 0)
      s2 += this.days + "D";
    if (this.hours !== 0 || this.minutes !== 0 || this.seconds !== 0 || this.milliseconds !== 0)
      s2 += "T";
    if (this.hours !== 0)
      s2 += this.hours + "H";
    if (this.minutes !== 0)
      s2 += this.minutes + "M";
    if (this.seconds !== 0 || this.milliseconds !== 0)
      s2 += roundTo(this.seconds + this.milliseconds / 1e3, 3) + "S";
    if (s2 === "P")
      s2 += "T0S";
    return s2;
  }
  toISOTime(opts = {}) {
    if (!this.isValid)
      return null;
    const millis = this.toMillis();
    if (millis < 0 || millis >= 864e5)
      return null;
    opts = Object.assign({
      suppressMilliseconds: false,
      suppressSeconds: false,
      includePrefix: false,
      format: "extended"
    }, opts);
    const value = this.shiftTo("hours", "minutes", "seconds", "milliseconds");
    let fmt = opts.format === "basic" ? "hhmm" : "hh:mm";
    if (!opts.suppressSeconds || value.seconds !== 0 || value.milliseconds !== 0) {
      fmt += opts.format === "basic" ? "ss" : ":ss";
      if (!opts.suppressMilliseconds || value.milliseconds !== 0) {
        fmt += ".SSS";
      }
    }
    let str = value.toFormat(fmt);
    if (opts.includePrefix) {
      str = "T" + str;
    }
    return str;
  }
  toJSON() {
    return this.toISO();
  }
  toString() {
    return this.toISO();
  }
  toMillis() {
    return this.as("milliseconds");
  }
  valueOf() {
    return this.toMillis();
  }
  plus(duration) {
    if (!this.isValid)
      return this;
    const dur = friendlyDuration(duration), result = {};
    for (const k of orderedUnits$1) {
      if (hasOwnProperty(dur.values, k) || hasOwnProperty(this.values, k)) {
        result[k] = dur.get(k) + this.get(k);
      }
    }
    return clone$1(this, { values: result }, true);
  }
  minus(duration) {
    if (!this.isValid)
      return this;
    const dur = friendlyDuration(duration);
    return this.plus(dur.negate());
  }
  mapUnits(fn) {
    if (!this.isValid)
      return this;
    const result = {};
    for (const k of Object.keys(this.values)) {
      result[k] = asNumber(fn(this.values[k], k));
    }
    return clone$1(this, { values: result }, true);
  }
  get(unit) {
    return this[Duration.normalizeUnit(unit)];
  }
  set(values) {
    if (!this.isValid)
      return this;
    const mixed = Object.assign(this.values, normalizeObject(values, Duration.normalizeUnit, []));
    return clone$1(this, { values: mixed });
  }
  reconfigure({ locale, numberingSystem, conversionAccuracy } = {}) {
    const loc = this.loc.clone({ locale, numberingSystem }), opts = { loc };
    if (conversionAccuracy) {
      opts.conversionAccuracy = conversionAccuracy;
    }
    return clone$1(this, opts);
  }
  as(unit) {
    return this.isValid ? this.shiftTo(unit).get(unit) : NaN;
  }
  normalize() {
    if (!this.isValid)
      return this;
    const vals = this.toObject();
    normalizeValues(this.matrix, vals);
    return clone$1(this, { values: vals }, true);
  }
  shiftTo(...units) {
    if (!this.isValid)
      return this;
    if (units.length === 0) {
      return this;
    }
    units = units.map((u) => Duration.normalizeUnit(u));
    const built = {}, accumulated = {}, vals = this.toObject();
    let lastUnit;
    for (const k of orderedUnits$1) {
      if (units.indexOf(k) >= 0) {
        lastUnit = k;
        let own = 0;
        for (const ak in accumulated) {
          own += this.matrix[ak][k] * accumulated[ak];
          accumulated[ak] = 0;
        }
        if (isNumber(vals[k])) {
          own += vals[k];
        }
        const i = Math.trunc(own);
        built[k] = i;
        accumulated[k] = own - i;
        for (const down in vals) {
          if (orderedUnits$1.indexOf(down) > orderedUnits$1.indexOf(k)) {
            convert(this.matrix, vals, down, built, k);
          }
        }
      } else if (isNumber(vals[k])) {
        accumulated[k] = vals[k];
      }
    }
    for (const key in accumulated) {
      if (accumulated[key] !== 0) {
        built[lastUnit] += key === lastUnit ? accumulated[key] : accumulated[key] / this.matrix[lastUnit][key];
      }
    }
    return clone$1(this, { values: built }, true).normalize();
  }
  negate() {
    if (!this.isValid)
      return this;
    const negated = {};
    for (const k of Object.keys(this.values)) {
      negated[k] = -this.values[k];
    }
    return clone$1(this, { values: negated }, true);
  }
  get years() {
    return this.isValid ? this.values.years || 0 : NaN;
  }
  get quarters() {
    return this.isValid ? this.values.quarters || 0 : NaN;
  }
  get months() {
    return this.isValid ? this.values.months || 0 : NaN;
  }
  get weeks() {
    return this.isValid ? this.values.weeks || 0 : NaN;
  }
  get days() {
    return this.isValid ? this.values.days || 0 : NaN;
  }
  get hours() {
    return this.isValid ? this.values.hours || 0 : NaN;
  }
  get minutes() {
    return this.isValid ? this.values.minutes || 0 : NaN;
  }
  get seconds() {
    return this.isValid ? this.values.seconds || 0 : NaN;
  }
  get milliseconds() {
    return this.isValid ? this.values.milliseconds || 0 : NaN;
  }
  get isValid() {
    return this.invalid === null;
  }
  get invalidReason() {
    return this.invalid ? this.invalid.reason : null;
  }
  get invalidExplanation() {
    return this.invalid ? this.invalid.explanation : null;
  }
  equals(other) {
    if (!this.isValid || !other.isValid) {
      return false;
    }
    if (!this.loc.equals(other.loc)) {
      return false;
    }
    function eq(v1, v2) {
      if (v1 === void 0 || v1 === 0)
        return v2 === void 0 || v2 === 0;
      return v1 === v2;
    }
    for (const u of orderedUnits$1) {
      if (!eq(this.values[u], other.values[u])) {
        return false;
      }
    }
    return true;
  }
}
function friendlyDuration(durationish) {
  if (isNumber(durationish)) {
    return Duration.fromMillis(durationish);
  } else if (Duration.isDuration(durationish)) {
    return durationish;
  } else if (typeof durationish === "object") {
    return Duration.fromObject(durationish);
  } else {
    throw new InvalidArgumentError(`Unknown duration argument ${durationish} of type ${typeof durationish}`);
  }
}
const INVALID$1 = "Invalid Interval";
function validateStartEnd(start, end) {
  if (!start || !start.isValid) {
    return Interval.invalid("missing or invalid start");
  } else if (!end || !end.isValid) {
    return Interval.invalid("missing or invalid end");
  } else if (end < start) {
    return Interval.invalid("end before start", `The end of an interval must be after its start, but you had start=${start.toISO()} and end=${end.toISO()}`);
  } else {
    return null;
  }
}
class Interval {
  constructor(config2) {
    this.s = config2.start;
    this.e = config2.end;
    this.invalid = config2.invalid || null;
    this.isLuxonInterval = true;
  }
  static invalid(reason, explanation = null) {
    if (!reason) {
      throw new InvalidArgumentError("need to specify a reason the Interval is invalid");
    }
    const invalid = reason instanceof Invalid ? reason : new Invalid(reason, explanation);
    if (Settings.throwOnInvalid) {
      throw new InvalidIntervalError(invalid);
    } else {
      return new Interval({ invalid });
    }
  }
  static fromDateTimes(start, end) {
    const builtStart = friendlyDateTime(start), builtEnd = friendlyDateTime(end);
    const validateError = validateStartEnd(builtStart, builtEnd);
    if (validateError == null) {
      return new Interval({
        start: builtStart,
        end: builtEnd
      });
    } else {
      return validateError;
    }
  }
  static after(start, duration) {
    const dur = friendlyDuration(duration), dt = friendlyDateTime(start);
    return Interval.fromDateTimes(dt, dt.plus(dur));
  }
  static before(end, duration) {
    const dur = friendlyDuration(duration), dt = friendlyDateTime(end);
    return Interval.fromDateTimes(dt.minus(dur), dt);
  }
  static fromISO(text, opts) {
    const [s2, e] = (text || "").split("/", 2);
    if (s2 && e) {
      let start, startIsValid;
      try {
        start = DateTime.fromISO(s2, opts);
        startIsValid = start.isValid;
      } catch (e2) {
        startIsValid = false;
      }
      let end, endIsValid;
      try {
        end = DateTime.fromISO(e, opts);
        endIsValid = end.isValid;
      } catch (e2) {
        endIsValid = false;
      }
      if (startIsValid && endIsValid) {
        return Interval.fromDateTimes(start, end);
      }
      if (startIsValid) {
        const dur = Duration.fromISO(e, opts);
        if (dur.isValid) {
          return Interval.after(start, dur);
        }
      } else if (endIsValid) {
        const dur = Duration.fromISO(s2, opts);
        if (dur.isValid) {
          return Interval.before(end, dur);
        }
      }
    }
    return Interval.invalid("unparsable", `the input "${text}" can't be parsed as ISO 8601`);
  }
  static isInterval(o) {
    return o && o.isLuxonInterval || false;
  }
  get start() {
    return this.isValid ? this.s : null;
  }
  get end() {
    return this.isValid ? this.e : null;
  }
  get isValid() {
    return this.invalidReason === null;
  }
  get invalidReason() {
    return this.invalid ? this.invalid.reason : null;
  }
  get invalidExplanation() {
    return this.invalid ? this.invalid.explanation : null;
  }
  length(unit = "milliseconds") {
    return this.isValid ? this.toDuration(...[unit]).get(unit) : NaN;
  }
  count(unit = "milliseconds") {
    if (!this.isValid)
      return NaN;
    const start = this.start.startOf(unit), end = this.end.startOf(unit);
    return Math.floor(end.diff(start, unit).get(unit)) + 1;
  }
  hasSame(unit) {
    return this.isValid ? this.isEmpty() || this.e.minus(1).hasSame(this.s, unit) : false;
  }
  isEmpty() {
    return this.s.valueOf() === this.e.valueOf();
  }
  isAfter(dateTime) {
    if (!this.isValid)
      return false;
    return this.s > dateTime;
  }
  isBefore(dateTime) {
    if (!this.isValid)
      return false;
    return this.e <= dateTime;
  }
  contains(dateTime) {
    if (!this.isValid)
      return false;
    return this.s <= dateTime && this.e > dateTime;
  }
  set({ start, end } = {}) {
    if (!this.isValid)
      return this;
    return Interval.fromDateTimes(start || this.s, end || this.e);
  }
  splitAt(...dateTimes) {
    if (!this.isValid)
      return [];
    const sorted2 = dateTimes.map(friendlyDateTime).filter((d) => this.contains(d)).sort(), results = [];
    let { s: s2 } = this, i = 0;
    while (s2 < this.e) {
      const added = sorted2[i] || this.e, next = +added > +this.e ? this.e : added;
      results.push(Interval.fromDateTimes(s2, next));
      s2 = next;
      i += 1;
    }
    return results;
  }
  splitBy(duration) {
    const dur = friendlyDuration(duration);
    if (!this.isValid || !dur.isValid || dur.as("milliseconds") === 0) {
      return [];
    }
    let { s: s2 } = this, idx = 1, next;
    const results = [];
    while (s2 < this.e) {
      const added = this.start.plus(dur.mapUnits((x) => x * idx));
      next = +added > +this.e ? this.e : added;
      results.push(Interval.fromDateTimes(s2, next));
      s2 = next;
      idx += 1;
    }
    return results;
  }
  divideEqually(numberOfParts) {
    if (!this.isValid)
      return [];
    return this.splitBy(this.length() / numberOfParts).slice(0, numberOfParts);
  }
  overlaps(other) {
    return this.e > other.s && this.s < other.e;
  }
  abutsStart(other) {
    if (!this.isValid)
      return false;
    return +this.e === +other.s;
  }
  abutsEnd(other) {
    if (!this.isValid)
      return false;
    return +other.e === +this.s;
  }
  engulfs(other) {
    if (!this.isValid)
      return false;
    return this.s <= other.s && this.e >= other.e;
  }
  equals(other) {
    if (!this.isValid || !other.isValid) {
      return false;
    }
    return this.s.equals(other.s) && this.e.equals(other.e);
  }
  intersection(other) {
    if (!this.isValid)
      return this;
    const s2 = this.s > other.s ? this.s : other.s, e = this.e < other.e ? this.e : other.e;
    if (s2 >= e) {
      return null;
    } else {
      return Interval.fromDateTimes(s2, e);
    }
  }
  union(other) {
    if (!this.isValid)
      return this;
    const s2 = this.s < other.s ? this.s : other.s, e = this.e > other.e ? this.e : other.e;
    return Interval.fromDateTimes(s2, e);
  }
  static merge(intervals) {
    const [found, final] = intervals.sort((a, b) => a.s - b.s).reduce(([sofar, current], item) => {
      if (!current) {
        return [sofar, item];
      } else if (current.overlaps(item) || current.abutsStart(item)) {
        return [sofar, current.union(item)];
      } else {
        return [sofar.concat([current]), item];
      }
    }, [[], null]);
    if (final) {
      found.push(final);
    }
    return found;
  }
  static xor(intervals) {
    let start = null, currentCount = 0;
    const results = [], ends = intervals.map((i) => [{ time: i.s, type: "s" }, { time: i.e, type: "e" }]), flattened = Array.prototype.concat(...ends), arr = flattened.sort((a, b) => a.time - b.time);
    for (const i of arr) {
      currentCount += i.type === "s" ? 1 : -1;
      if (currentCount === 1) {
        start = i.time;
      } else {
        if (start && +start !== +i.time) {
          results.push(Interval.fromDateTimes(start, i.time));
        }
        start = null;
      }
    }
    return Interval.merge(results);
  }
  difference(...intervals) {
    return Interval.xor([this].concat(intervals)).map((i) => this.intersection(i)).filter((i) => i && !i.isEmpty());
  }
  toString() {
    if (!this.isValid)
      return INVALID$1;
    return `[${this.s.toISO()} \u2013 ${this.e.toISO()})`;
  }
  toISO(opts) {
    if (!this.isValid)
      return INVALID$1;
    return `${this.s.toISO(opts)}/${this.e.toISO(opts)}`;
  }
  toISODate() {
    if (!this.isValid)
      return INVALID$1;
    return `${this.s.toISODate()}/${this.e.toISODate()}`;
  }
  toISOTime(opts) {
    if (!this.isValid)
      return INVALID$1;
    return `${this.s.toISOTime(opts)}/${this.e.toISOTime(opts)}`;
  }
  toFormat(dateFormat, { separator = " \u2013 " } = {}) {
    if (!this.isValid)
      return INVALID$1;
    return `${this.s.toFormat(dateFormat)}${separator}${this.e.toFormat(dateFormat)}`;
  }
  toDuration(unit, opts) {
    if (!this.isValid) {
      return Duration.invalid(this.invalidReason);
    }
    return this.e.diff(this.s, unit, opts);
  }
  mapEndpoints(mapFn) {
    return Interval.fromDateTimes(mapFn(this.s), mapFn(this.e));
  }
}
class Info {
  static hasDST(zone = Settings.defaultZone) {
    const proto = DateTime.now().setZone(zone).set({ month: 12 });
    return !zone.universal && proto.offset !== proto.set({ month: 6 }).offset;
  }
  static isValidIANAZone(zone) {
    return IANAZone.isValidSpecifier(zone) && IANAZone.isValidZone(zone);
  }
  static normalizeZone(input) {
    return normalizeZone(input, Settings.defaultZone);
  }
  static months(length = "long", { locale = null, numberingSystem = null, locObj = null, outputCalendar = "gregory" } = {}) {
    return (locObj || Locale.create(locale, numberingSystem, outputCalendar)).months(length);
  }
  static monthsFormat(length = "long", { locale = null, numberingSystem = null, locObj = null, outputCalendar = "gregory" } = {}) {
    return (locObj || Locale.create(locale, numberingSystem, outputCalendar)).months(length, true);
  }
  static weekdays(length = "long", { locale = null, numberingSystem = null, locObj = null } = {}) {
    return (locObj || Locale.create(locale, numberingSystem, null)).weekdays(length);
  }
  static weekdaysFormat(length = "long", { locale = null, numberingSystem = null, locObj = null } = {}) {
    return (locObj || Locale.create(locale, numberingSystem, null)).weekdays(length, true);
  }
  static meridiems({ locale = null } = {}) {
    return Locale.create(locale).meridiems();
  }
  static eras(length = "short", { locale = null } = {}) {
    return Locale.create(locale, null, "gregory").eras(length);
  }
  static features() {
    let intl = false, intlTokens = false, zones = false, relative = false;
    if (hasIntl()) {
      intl = true;
      intlTokens = hasFormatToParts();
      relative = hasRelative();
      try {
        zones = new Intl.DateTimeFormat("en", { timeZone: "America/New_York" }).resolvedOptions().timeZone === "America/New_York";
      } catch (e) {
        zones = false;
      }
    }
    return { intl, intlTokens, zones, relative };
  }
}
function dayDiff(earlier, later) {
  const utcDayStart = (dt) => dt.toUTC(0, { keepLocalTime: true }).startOf("day").valueOf(), ms = utcDayStart(later) - utcDayStart(earlier);
  return Math.floor(Duration.fromMillis(ms).as("days"));
}
function highOrderDiffs(cursor, later, units) {
  const differs = [
    ["years", (a, b) => b.year - a.year],
    ["quarters", (a, b) => b.quarter - a.quarter],
    ["months", (a, b) => b.month - a.month + (b.year - a.year) * 12],
    [
      "weeks",
      (a, b) => {
        const days = dayDiff(a, b);
        return (days - days % 7) / 7;
      }
    ],
    ["days", dayDiff]
  ];
  const results = {};
  let lowestOrder, highWater;
  for (const [unit, differ] of differs) {
    if (units.indexOf(unit) >= 0) {
      lowestOrder = unit;
      let delta = differ(cursor, later);
      highWater = cursor.plus({ [unit]: delta });
      if (highWater > later) {
        cursor = cursor.plus({ [unit]: delta - 1 });
        delta -= 1;
      } else {
        cursor = highWater;
      }
      results[unit] = delta;
    }
  }
  return [cursor, results, highWater, lowestOrder];
}
function diff(earlier, later, units, opts) {
  let [cursor, results, highWater, lowestOrder] = highOrderDiffs(earlier, later, units);
  const remainingMillis = later - cursor;
  const lowerOrderUnits = units.filter((u) => ["hours", "minutes", "seconds", "milliseconds"].indexOf(u) >= 0);
  if (lowerOrderUnits.length === 0) {
    if (highWater < later) {
      highWater = cursor.plus({ [lowestOrder]: 1 });
    }
    if (highWater !== cursor) {
      results[lowestOrder] = (results[lowestOrder] || 0) + remainingMillis / (highWater - cursor);
    }
  }
  const duration = Duration.fromObject(Object.assign(results, opts));
  if (lowerOrderUnits.length > 0) {
    return Duration.fromMillis(remainingMillis, opts).shiftTo(...lowerOrderUnits).plus(duration);
  } else {
    return duration;
  }
}
const numberingSystems = {
  arab: "[\u0660-\u0669]",
  arabext: "[\u06F0-\u06F9]",
  bali: "[\u1B50-\u1B59]",
  beng: "[\u09E6-\u09EF]",
  deva: "[\u0966-\u096F]",
  fullwide: "[\uFF10-\uFF19]",
  gujr: "[\u0AE6-\u0AEF]",
  hanidec: "[\u3007|\u4E00|\u4E8C|\u4E09|\u56DB|\u4E94|\u516D|\u4E03|\u516B|\u4E5D]",
  khmr: "[\u17E0-\u17E9]",
  knda: "[\u0CE6-\u0CEF]",
  laoo: "[\u0ED0-\u0ED9]",
  limb: "[\u1946-\u194F]",
  mlym: "[\u0D66-\u0D6F]",
  mong: "[\u1810-\u1819]",
  mymr: "[\u1040-\u1049]",
  orya: "[\u0B66-\u0B6F]",
  tamldec: "[\u0BE6-\u0BEF]",
  telu: "[\u0C66-\u0C6F]",
  thai: "[\u0E50-\u0E59]",
  tibt: "[\u0F20-\u0F29]",
  latn: "\\d"
};
const numberingSystemsUTF16 = {
  arab: [1632, 1641],
  arabext: [1776, 1785],
  bali: [6992, 7001],
  beng: [2534, 2543],
  deva: [2406, 2415],
  fullwide: [65296, 65303],
  gujr: [2790, 2799],
  khmr: [6112, 6121],
  knda: [3302, 3311],
  laoo: [3792, 3801],
  limb: [6470, 6479],
  mlym: [3430, 3439],
  mong: [6160, 6169],
  mymr: [4160, 4169],
  orya: [2918, 2927],
  tamldec: [3046, 3055],
  telu: [3174, 3183],
  thai: [3664, 3673],
  tibt: [3872, 3881]
};
const hanidecChars = numberingSystems.hanidec.replace(/[\[|\]]/g, "").split("");
function parseDigits(str) {
  let value = parseInt(str, 10);
  if (isNaN(value)) {
    value = "";
    for (let i = 0; i < str.length; i++) {
      const code = str.charCodeAt(i);
      if (str[i].search(numberingSystems.hanidec) !== -1) {
        value += hanidecChars.indexOf(str[i]);
      } else {
        for (const key in numberingSystemsUTF16) {
          const [min, max] = numberingSystemsUTF16[key];
          if (code >= min && code <= max) {
            value += code - min;
          }
        }
      }
    }
    return parseInt(value, 10);
  } else {
    return value;
  }
}
function digitRegex({ numberingSystem }, append = "") {
  return new RegExp(`${numberingSystems[numberingSystem || "latn"]}${append}`);
}
const MISSING_FTP = "missing Intl.DateTimeFormat.formatToParts support";
function intUnit(regex, post = (i) => i) {
  return { regex, deser: ([s2]) => post(parseDigits(s2)) };
}
const NBSP = String.fromCharCode(160);
const spaceOrNBSP = `( |${NBSP})`;
const spaceOrNBSPRegExp = new RegExp(spaceOrNBSP, "g");
function fixListRegex(s2) {
  return s2.replace(/\./g, "\\.?").replace(spaceOrNBSPRegExp, spaceOrNBSP);
}
function stripInsensitivities(s2) {
  return s2.replace(/\./g, "").replace(spaceOrNBSPRegExp, " ").toLowerCase();
}
function oneOf(strings, startIndex) {
  if (strings === null) {
    return null;
  } else {
    return {
      regex: RegExp(strings.map(fixListRegex).join("|")),
      deser: ([s2]) => strings.findIndex((i) => stripInsensitivities(s2) === stripInsensitivities(i)) + startIndex
    };
  }
}
function offset(regex, groups) {
  return { regex, deser: ([, h2, m]) => signedOffset(h2, m), groups };
}
function simple(regex) {
  return { regex, deser: ([s2]) => s2 };
}
function escapeToken(value) {
  return value.replace(/[\-\[\]{}()*+?.,\\\^$|#\s]/g, "\\$&");
}
function unitForToken(token, loc) {
  const one = digitRegex(loc), two = digitRegex(loc, "{2}"), three = digitRegex(loc, "{3}"), four = digitRegex(loc, "{4}"), six = digitRegex(loc, "{6}"), oneOrTwo = digitRegex(loc, "{1,2}"), oneToThree = digitRegex(loc, "{1,3}"), oneToSix = digitRegex(loc, "{1,6}"), oneToNine = digitRegex(loc, "{1,9}"), twoToFour = digitRegex(loc, "{2,4}"), fourToSix = digitRegex(loc, "{4,6}"), literal = (t) => ({ regex: RegExp(escapeToken(t.val)), deser: ([s2]) => s2, literal: true }), unitate = (t) => {
    if (token.literal) {
      return literal(t);
    }
    switch (t.val) {
      case "G":
        return oneOf(loc.eras("short", false), 0);
      case "GG":
        return oneOf(loc.eras("long", false), 0);
      case "y":
        return intUnit(oneToSix);
      case "yy":
        return intUnit(twoToFour, untruncateYear);
      case "yyyy":
        return intUnit(four);
      case "yyyyy":
        return intUnit(fourToSix);
      case "yyyyyy":
        return intUnit(six);
      case "M":
        return intUnit(oneOrTwo);
      case "MM":
        return intUnit(two);
      case "MMM":
        return oneOf(loc.months("short", true, false), 1);
      case "MMMM":
        return oneOf(loc.months("long", true, false), 1);
      case "L":
        return intUnit(oneOrTwo);
      case "LL":
        return intUnit(two);
      case "LLL":
        return oneOf(loc.months("short", false, false), 1);
      case "LLLL":
        return oneOf(loc.months("long", false, false), 1);
      case "d":
        return intUnit(oneOrTwo);
      case "dd":
        return intUnit(two);
      case "o":
        return intUnit(oneToThree);
      case "ooo":
        return intUnit(three);
      case "HH":
        return intUnit(two);
      case "H":
        return intUnit(oneOrTwo);
      case "hh":
        return intUnit(two);
      case "h":
        return intUnit(oneOrTwo);
      case "mm":
        return intUnit(two);
      case "m":
        return intUnit(oneOrTwo);
      case "q":
        return intUnit(oneOrTwo);
      case "qq":
        return intUnit(two);
      case "s":
        return intUnit(oneOrTwo);
      case "ss":
        return intUnit(two);
      case "S":
        return intUnit(oneToThree);
      case "SSS":
        return intUnit(three);
      case "u":
        return simple(oneToNine);
      case "a":
        return oneOf(loc.meridiems(), 0);
      case "kkkk":
        return intUnit(four);
      case "kk":
        return intUnit(twoToFour, untruncateYear);
      case "W":
        return intUnit(oneOrTwo);
      case "WW":
        return intUnit(two);
      case "E":
      case "c":
        return intUnit(one);
      case "EEE":
        return oneOf(loc.weekdays("short", false, false), 1);
      case "EEEE":
        return oneOf(loc.weekdays("long", false, false), 1);
      case "ccc":
        return oneOf(loc.weekdays("short", true, false), 1);
      case "cccc":
        return oneOf(loc.weekdays("long", true, false), 1);
      case "Z":
      case "ZZ":
        return offset(new RegExp(`([+-]${oneOrTwo.source})(?::(${two.source}))?`), 2);
      case "ZZZ":
        return offset(new RegExp(`([+-]${oneOrTwo.source})(${two.source})?`), 2);
      case "z":
        return simple(/[a-z_+-/]{1,256}?/i);
      default:
        return literal(t);
    }
  };
  const unit = unitate(token) || {
    invalidReason: MISSING_FTP
  };
  unit.token = token;
  return unit;
}
const partTypeStyleToTokenVal = {
  year: {
    "2-digit": "yy",
    numeric: "yyyyy"
  },
  month: {
    numeric: "M",
    "2-digit": "MM",
    short: "MMM",
    long: "MMMM"
  },
  day: {
    numeric: "d",
    "2-digit": "dd"
  },
  weekday: {
    short: "EEE",
    long: "EEEE"
  },
  dayperiod: "a",
  dayPeriod: "a",
  hour: {
    numeric: "h",
    "2-digit": "hh"
  },
  minute: {
    numeric: "m",
    "2-digit": "mm"
  },
  second: {
    numeric: "s",
    "2-digit": "ss"
  }
};
function tokenForPart(part, locale, formatOpts) {
  const { type, value } = part;
  if (type === "literal") {
    return {
      literal: true,
      val: value
    };
  }
  const style = formatOpts[type];
  let val = partTypeStyleToTokenVal[type];
  if (typeof val === "object") {
    val = val[style];
  }
  if (val) {
    return {
      literal: false,
      val
    };
  }
  return void 0;
}
function buildRegex(units) {
  const re = units.map((u) => u.regex).reduce((f, r2) => `${f}(${r2.source})`, "");
  return [`^${re}$`, units];
}
function match(input, regex, handlers) {
  const matches = input.match(regex);
  if (matches) {
    const all = {};
    let matchIndex = 1;
    for (const i in handlers) {
      if (hasOwnProperty(handlers, i)) {
        const h2 = handlers[i], groups = h2.groups ? h2.groups + 1 : 1;
        if (!h2.literal && h2.token) {
          all[h2.token.val[0]] = h2.deser(matches.slice(matchIndex, matchIndex + groups));
        }
        matchIndex += groups;
      }
    }
    return [matches, all];
  } else {
    return [matches, {}];
  }
}
function dateTimeFromMatches(matches) {
  const toField = (token) => {
    switch (token) {
      case "S":
        return "millisecond";
      case "s":
        return "second";
      case "m":
        return "minute";
      case "h":
      case "H":
        return "hour";
      case "d":
        return "day";
      case "o":
        return "ordinal";
      case "L":
      case "M":
        return "month";
      case "y":
        return "year";
      case "E":
      case "c":
        return "weekday";
      case "W":
        return "weekNumber";
      case "k":
        return "weekYear";
      case "q":
        return "quarter";
      default:
        return null;
    }
  };
  let zone;
  if (!isUndefined(matches.Z)) {
    zone = new FixedOffsetZone(matches.Z);
  } else if (!isUndefined(matches.z)) {
    zone = IANAZone.create(matches.z);
  } else {
    zone = null;
  }
  if (!isUndefined(matches.q)) {
    matches.M = (matches.q - 1) * 3 + 1;
  }
  if (!isUndefined(matches.h)) {
    if (matches.h < 12 && matches.a === 1) {
      matches.h += 12;
    } else if (matches.h === 12 && matches.a === 0) {
      matches.h = 0;
    }
  }
  if (matches.G === 0 && matches.y) {
    matches.y = -matches.y;
  }
  if (!isUndefined(matches.u)) {
    matches.S = parseMillis(matches.u);
  }
  const vals = Object.keys(matches).reduce((r2, k) => {
    const f = toField(k);
    if (f) {
      r2[f] = matches[k];
    }
    return r2;
  }, {});
  return [vals, zone];
}
let dummyDateTimeCache = null;
function getDummyDateTime() {
  if (!dummyDateTimeCache) {
    dummyDateTimeCache = DateTime.fromMillis(1555555555555);
  }
  return dummyDateTimeCache;
}
function maybeExpandMacroToken(token, locale) {
  if (token.literal) {
    return token;
  }
  const formatOpts = Formatter.macroTokenToFormatOpts(token.val);
  if (!formatOpts) {
    return token;
  }
  const formatter = Formatter.create(locale, formatOpts);
  const parts = formatter.formatDateTimeParts(getDummyDateTime());
  const tokens = parts.map((p) => tokenForPart(p, locale, formatOpts));
  if (tokens.includes(void 0)) {
    return token;
  }
  return tokens;
}
function expandMacroTokens(tokens, locale) {
  return Array.prototype.concat(...tokens.map((t) => maybeExpandMacroToken(t, locale)));
}
function explainFromTokens(locale, input, format) {
  const tokens = expandMacroTokens(Formatter.parseFormat(format), locale), units = tokens.map((t) => unitForToken(t, locale)), disqualifyingUnit = units.find((t) => t.invalidReason);
  if (disqualifyingUnit) {
    return { input, tokens, invalidReason: disqualifyingUnit.invalidReason };
  } else {
    const [regexString, handlers] = buildRegex(units), regex = RegExp(regexString, "i"), [rawMatches, matches] = match(input, regex, handlers), [result, zone] = matches ? dateTimeFromMatches(matches) : [null, null];
    if (hasOwnProperty(matches, "a") && hasOwnProperty(matches, "H")) {
      throw new ConflictingSpecificationError("Can't include meridiem when specifying 24-hour format");
    }
    return { input, tokens, regex, rawMatches, matches, result, zone };
  }
}
function parseFromTokens(locale, input, format) {
  const { result, zone, invalidReason } = explainFromTokens(locale, input, format);
  return [result, zone, invalidReason];
}
const nonLeapLadder = [0, 31, 59, 90, 120, 151, 181, 212, 243, 273, 304, 334], leapLadder = [0, 31, 60, 91, 121, 152, 182, 213, 244, 274, 305, 335];
function unitOutOfRange(unit, value) {
  return new Invalid("unit out of range", `you specified ${value} (of type ${typeof value}) as a ${unit}, which is invalid`);
}
function dayOfWeek(year, month, day) {
  const js = new Date(Date.UTC(year, month - 1, day)).getUTCDay();
  return js === 0 ? 7 : js;
}
function computeOrdinal(year, month, day) {
  return day + (isLeapYear(year) ? leapLadder : nonLeapLadder)[month - 1];
}
function uncomputeOrdinal(year, ordinal) {
  const table = isLeapYear(year) ? leapLadder : nonLeapLadder, month0 = table.findIndex((i) => i < ordinal), day = ordinal - table[month0];
  return { month: month0 + 1, day };
}
function gregorianToWeek(gregObj) {
  const { year, month, day } = gregObj, ordinal = computeOrdinal(year, month, day), weekday = dayOfWeek(year, month, day);
  let weekNumber = Math.floor((ordinal - weekday + 10) / 7), weekYear;
  if (weekNumber < 1) {
    weekYear = year - 1;
    weekNumber = weeksInWeekYear(weekYear);
  } else if (weekNumber > weeksInWeekYear(year)) {
    weekYear = year + 1;
    weekNumber = 1;
  } else {
    weekYear = year;
  }
  return Object.assign({ weekYear, weekNumber, weekday }, timeObject(gregObj));
}
function weekToGregorian(weekData) {
  const { weekYear, weekNumber, weekday } = weekData, weekdayOfJan4 = dayOfWeek(weekYear, 1, 4), yearInDays = daysInYear(weekYear);
  let ordinal = weekNumber * 7 + weekday - weekdayOfJan4 - 3, year;
  if (ordinal < 1) {
    year = weekYear - 1;
    ordinal += daysInYear(year);
  } else if (ordinal > yearInDays) {
    year = weekYear + 1;
    ordinal -= daysInYear(weekYear);
  } else {
    year = weekYear;
  }
  const { month, day } = uncomputeOrdinal(year, ordinal);
  return Object.assign({ year, month, day }, timeObject(weekData));
}
function gregorianToOrdinal(gregData) {
  const { year, month, day } = gregData, ordinal = computeOrdinal(year, month, day);
  return Object.assign({ year, ordinal }, timeObject(gregData));
}
function ordinalToGregorian(ordinalData) {
  const { year, ordinal } = ordinalData, { month, day } = uncomputeOrdinal(year, ordinal);
  return Object.assign({ year, month, day }, timeObject(ordinalData));
}
function hasInvalidWeekData(obj) {
  const validYear = isInteger(obj.weekYear), validWeek = integerBetween(obj.weekNumber, 1, weeksInWeekYear(obj.weekYear)), validWeekday = integerBetween(obj.weekday, 1, 7);
  if (!validYear) {
    return unitOutOfRange("weekYear", obj.weekYear);
  } else if (!validWeek) {
    return unitOutOfRange("week", obj.week);
  } else if (!validWeekday) {
    return unitOutOfRange("weekday", obj.weekday);
  } else
    return false;
}
function hasInvalidOrdinalData(obj) {
  const validYear = isInteger(obj.year), validOrdinal = integerBetween(obj.ordinal, 1, daysInYear(obj.year));
  if (!validYear) {
    return unitOutOfRange("year", obj.year);
  } else if (!validOrdinal) {
    return unitOutOfRange("ordinal", obj.ordinal);
  } else
    return false;
}
function hasInvalidGregorianData(obj) {
  const validYear = isInteger(obj.year), validMonth = integerBetween(obj.month, 1, 12), validDay = integerBetween(obj.day, 1, daysInMonth(obj.year, obj.month));
  if (!validYear) {
    return unitOutOfRange("year", obj.year);
  } else if (!validMonth) {
    return unitOutOfRange("month", obj.month);
  } else if (!validDay) {
    return unitOutOfRange("day", obj.day);
  } else
    return false;
}
function hasInvalidTimeData(obj) {
  const { hour, minute, second, millisecond } = obj;
  const validHour = integerBetween(hour, 0, 23) || hour === 24 && minute === 0 && second === 0 && millisecond === 0, validMinute = integerBetween(minute, 0, 59), validSecond = integerBetween(second, 0, 59), validMillisecond = integerBetween(millisecond, 0, 999);
  if (!validHour) {
    return unitOutOfRange("hour", hour);
  } else if (!validMinute) {
    return unitOutOfRange("minute", minute);
  } else if (!validSecond) {
    return unitOutOfRange("second", second);
  } else if (!validMillisecond) {
    return unitOutOfRange("millisecond", millisecond);
  } else
    return false;
}
const INVALID = "Invalid DateTime";
const MAX_DATE = 864e13;
function unsupportedZone(zone) {
  return new Invalid("unsupported zone", `the zone "${zone.name}" is not supported`);
}
function possiblyCachedWeekData(dt) {
  if (dt.weekData === null) {
    dt.weekData = gregorianToWeek(dt.c);
  }
  return dt.weekData;
}
function clone(inst, alts) {
  const current = {
    ts: inst.ts,
    zone: inst.zone,
    c: inst.c,
    o: inst.o,
    loc: inst.loc,
    invalid: inst.invalid
  };
  return new DateTime(Object.assign({}, current, alts, { old: current }));
}
function fixOffset(localTS, o, tz) {
  let utcGuess = localTS - o * 60 * 1e3;
  const o2 = tz.offset(utcGuess);
  if (o === o2) {
    return [utcGuess, o];
  }
  utcGuess -= (o2 - o) * 60 * 1e3;
  const o3 = tz.offset(utcGuess);
  if (o2 === o3) {
    return [utcGuess, o2];
  }
  return [localTS - Math.min(o2, o3) * 60 * 1e3, Math.max(o2, o3)];
}
function tsToObj(ts, offset2) {
  ts += offset2 * 60 * 1e3;
  const d = new Date(ts);
  return {
    year: d.getUTCFullYear(),
    month: d.getUTCMonth() + 1,
    day: d.getUTCDate(),
    hour: d.getUTCHours(),
    minute: d.getUTCMinutes(),
    second: d.getUTCSeconds(),
    millisecond: d.getUTCMilliseconds()
  };
}
function objToTS(obj, offset2, zone) {
  return fixOffset(objToLocalTS(obj), offset2, zone);
}
function adjustTime(inst, dur) {
  const oPre = inst.o, year = inst.c.year + Math.trunc(dur.years), month = inst.c.month + Math.trunc(dur.months) + Math.trunc(dur.quarters) * 3, c = Object.assign({}, inst.c, {
    year,
    month,
    day: Math.min(inst.c.day, daysInMonth(year, month)) + Math.trunc(dur.days) + Math.trunc(dur.weeks) * 7
  }), millisToAdd = Duration.fromObject({
    years: dur.years - Math.trunc(dur.years),
    quarters: dur.quarters - Math.trunc(dur.quarters),
    months: dur.months - Math.trunc(dur.months),
    weeks: dur.weeks - Math.trunc(dur.weeks),
    days: dur.days - Math.trunc(dur.days),
    hours: dur.hours,
    minutes: dur.minutes,
    seconds: dur.seconds,
    milliseconds: dur.milliseconds
  }).as("milliseconds"), localTS = objToLocalTS(c);
  let [ts, o] = fixOffset(localTS, oPre, inst.zone);
  if (millisToAdd !== 0) {
    ts += millisToAdd;
    o = inst.zone.offset(ts);
  }
  return { ts, o };
}
function parseDataToDateTime(parsed, parsedZone, opts, format, text) {
  const { setZone, zone } = opts;
  if (parsed && Object.keys(parsed).length !== 0) {
    const interpretationZone = parsedZone || zone, inst = DateTime.fromObject(Object.assign(parsed, opts, {
      zone: interpretationZone,
      setZone: void 0
    }));
    return setZone ? inst : inst.setZone(zone);
  } else {
    return DateTime.invalid(new Invalid("unparsable", `the input "${text}" can't be parsed as ${format}`));
  }
}
function toTechFormat(dt, format, allowZ = true) {
  return dt.isValid ? Formatter.create(Locale.create("en-US"), {
    allowZ,
    forceSimple: true
  }).formatDateTimeFromString(dt, format) : null;
}
function toTechTimeFormat(dt, {
  suppressSeconds = false,
  suppressMilliseconds = false,
  includeOffset,
  includePrefix = false,
  includeZone = false,
  spaceZone = false,
  format = "extended"
}) {
  let fmt = format === "basic" ? "HHmm" : "HH:mm";
  if (!suppressSeconds || dt.second !== 0 || dt.millisecond !== 0) {
    fmt += format === "basic" ? "ss" : ":ss";
    if (!suppressMilliseconds || dt.millisecond !== 0) {
      fmt += ".SSS";
    }
  }
  if ((includeZone || includeOffset) && spaceZone) {
    fmt += " ";
  }
  if (includeZone) {
    fmt += "z";
  } else if (includeOffset) {
    fmt += format === "basic" ? "ZZZ" : "ZZ";
  }
  let str = toTechFormat(dt, fmt);
  if (includePrefix) {
    str = "T" + str;
  }
  return str;
}
const defaultUnitValues = {
  month: 1,
  day: 1,
  hour: 0,
  minute: 0,
  second: 0,
  millisecond: 0
}, defaultWeekUnitValues = {
  weekNumber: 1,
  weekday: 1,
  hour: 0,
  minute: 0,
  second: 0,
  millisecond: 0
}, defaultOrdinalUnitValues = {
  ordinal: 1,
  hour: 0,
  minute: 0,
  second: 0,
  millisecond: 0
};
const orderedUnits = ["year", "month", "day", "hour", "minute", "second", "millisecond"], orderedWeekUnits = [
  "weekYear",
  "weekNumber",
  "weekday",
  "hour",
  "minute",
  "second",
  "millisecond"
], orderedOrdinalUnits = ["year", "ordinal", "hour", "minute", "second", "millisecond"];
function normalizeUnit(unit) {
  const normalized = {
    year: "year",
    years: "year",
    month: "month",
    months: "month",
    day: "day",
    days: "day",
    hour: "hour",
    hours: "hour",
    minute: "minute",
    minutes: "minute",
    quarter: "quarter",
    quarters: "quarter",
    second: "second",
    seconds: "second",
    millisecond: "millisecond",
    milliseconds: "millisecond",
    weekday: "weekday",
    weekdays: "weekday",
    weeknumber: "weekNumber",
    weeksnumber: "weekNumber",
    weeknumbers: "weekNumber",
    weekyear: "weekYear",
    weekyears: "weekYear",
    ordinal: "ordinal"
  }[unit.toLowerCase()];
  if (!normalized)
    throw new InvalidUnitError(unit);
  return normalized;
}
function quickDT(obj, zone) {
  for (const u of orderedUnits) {
    if (isUndefined(obj[u])) {
      obj[u] = defaultUnitValues[u];
    }
  }
  const invalid = hasInvalidGregorianData(obj) || hasInvalidTimeData(obj);
  if (invalid) {
    return DateTime.invalid(invalid);
  }
  const tsNow = Settings.now(), offsetProvis = zone.offset(tsNow), [ts, o] = objToTS(obj, offsetProvis, zone);
  return new DateTime({
    ts,
    zone,
    o
  });
}
function diffRelative(start, end, opts) {
  const round = isUndefined(opts.round) ? true : opts.round, format = (c, unit) => {
    c = roundTo(c, round || opts.calendary ? 0 : 2, true);
    const formatter = end.loc.clone(opts).relFormatter(opts);
    return formatter.format(c, unit);
  }, differ = (unit) => {
    if (opts.calendary) {
      if (!end.hasSame(start, unit)) {
        return end.startOf(unit).diff(start.startOf(unit), unit).get(unit);
      } else
        return 0;
    } else {
      return end.diff(start, unit).get(unit);
    }
  };
  if (opts.unit) {
    return format(differ(opts.unit), opts.unit);
  }
  for (const unit of opts.units) {
    const count = differ(unit);
    if (Math.abs(count) >= 1) {
      return format(count, unit);
    }
  }
  return format(start > end ? -0 : 0, opts.units[opts.units.length - 1]);
}
class DateTime {
  constructor(config2) {
    const zone = config2.zone || Settings.defaultZone;
    let invalid = config2.invalid || (Number.isNaN(config2.ts) ? new Invalid("invalid input") : null) || (!zone.isValid ? unsupportedZone(zone) : null);
    this.ts = isUndefined(config2.ts) ? Settings.now() : config2.ts;
    let c = null, o = null;
    if (!invalid) {
      const unchanged = config2.old && config2.old.ts === this.ts && config2.old.zone.equals(zone);
      if (unchanged) {
        [c, o] = [config2.old.c, config2.old.o];
      } else {
        const ot = zone.offset(this.ts);
        c = tsToObj(this.ts, ot);
        invalid = Number.isNaN(c.year) ? new Invalid("invalid input") : null;
        c = invalid ? null : c;
        o = invalid ? null : ot;
      }
    }
    this._zone = zone;
    this.loc = config2.loc || Locale.create();
    this.invalid = invalid;
    this.weekData = null;
    this.c = c;
    this.o = o;
    this.isLuxonDateTime = true;
  }
  static now() {
    return new DateTime({});
  }
  static local(year, month, day, hour, minute, second, millisecond) {
    if (isUndefined(year)) {
      return DateTime.now();
    } else {
      return quickDT({
        year,
        month,
        day,
        hour,
        minute,
        second,
        millisecond
      }, Settings.defaultZone);
    }
  }
  static utc(year, month, day, hour, minute, second, millisecond) {
    if (isUndefined(year)) {
      return new DateTime({
        ts: Settings.now(),
        zone: FixedOffsetZone.utcInstance
      });
    } else {
      return quickDT({
        year,
        month,
        day,
        hour,
        minute,
        second,
        millisecond
      }, FixedOffsetZone.utcInstance);
    }
  }
  static fromJSDate(date, options = {}) {
    const ts = isDate(date) ? date.valueOf() : NaN;
    if (Number.isNaN(ts)) {
      return DateTime.invalid("invalid input");
    }
    const zoneToUse = normalizeZone(options.zone, Settings.defaultZone);
    if (!zoneToUse.isValid) {
      return DateTime.invalid(unsupportedZone(zoneToUse));
    }
    return new DateTime({
      ts,
      zone: zoneToUse,
      loc: Locale.fromObject(options)
    });
  }
  static fromMillis(milliseconds, options = {}) {
    if (!isNumber(milliseconds)) {
      throw new InvalidArgumentError(`fromMillis requires a numerical input, but received a ${typeof milliseconds} with value ${milliseconds}`);
    } else if (milliseconds < -MAX_DATE || milliseconds > MAX_DATE) {
      return DateTime.invalid("Timestamp out of range");
    } else {
      return new DateTime({
        ts: milliseconds,
        zone: normalizeZone(options.zone, Settings.defaultZone),
        loc: Locale.fromObject(options)
      });
    }
  }
  static fromSeconds(seconds, options = {}) {
    if (!isNumber(seconds)) {
      throw new InvalidArgumentError("fromSeconds requires a numerical input");
    } else {
      return new DateTime({
        ts: seconds * 1e3,
        zone: normalizeZone(options.zone, Settings.defaultZone),
        loc: Locale.fromObject(options)
      });
    }
  }
  static fromObject(obj) {
    const zoneToUse = normalizeZone(obj.zone, Settings.defaultZone);
    if (!zoneToUse.isValid) {
      return DateTime.invalid(unsupportedZone(zoneToUse));
    }
    const tsNow = Settings.now(), offsetProvis = zoneToUse.offset(tsNow), normalized = normalizeObject(obj, normalizeUnit, [
      "zone",
      "locale",
      "outputCalendar",
      "numberingSystem"
    ]), containsOrdinal = !isUndefined(normalized.ordinal), containsGregorYear = !isUndefined(normalized.year), containsGregorMD = !isUndefined(normalized.month) || !isUndefined(normalized.day), containsGregor = containsGregorYear || containsGregorMD, definiteWeekDef = normalized.weekYear || normalized.weekNumber, loc = Locale.fromObject(obj);
    if ((containsGregor || containsOrdinal) && definiteWeekDef) {
      throw new ConflictingSpecificationError("Can't mix weekYear/weekNumber units with year/month/day or ordinals");
    }
    if (containsGregorMD && containsOrdinal) {
      throw new ConflictingSpecificationError("Can't mix ordinal dates with month/day");
    }
    const useWeekData = definiteWeekDef || normalized.weekday && !containsGregor;
    let units, defaultValues, objNow = tsToObj(tsNow, offsetProvis);
    if (useWeekData) {
      units = orderedWeekUnits;
      defaultValues = defaultWeekUnitValues;
      objNow = gregorianToWeek(objNow);
    } else if (containsOrdinal) {
      units = orderedOrdinalUnits;
      defaultValues = defaultOrdinalUnitValues;
      objNow = gregorianToOrdinal(objNow);
    } else {
      units = orderedUnits;
      defaultValues = defaultUnitValues;
    }
    let foundFirst = false;
    for (const u of units) {
      const v = normalized[u];
      if (!isUndefined(v)) {
        foundFirst = true;
      } else if (foundFirst) {
        normalized[u] = defaultValues[u];
      } else {
        normalized[u] = objNow[u];
      }
    }
    const higherOrderInvalid = useWeekData ? hasInvalidWeekData(normalized) : containsOrdinal ? hasInvalidOrdinalData(normalized) : hasInvalidGregorianData(normalized), invalid = higherOrderInvalid || hasInvalidTimeData(normalized);
    if (invalid) {
      return DateTime.invalid(invalid);
    }
    const gregorian = useWeekData ? weekToGregorian(normalized) : containsOrdinal ? ordinalToGregorian(normalized) : normalized, [tsFinal, offsetFinal] = objToTS(gregorian, offsetProvis, zoneToUse), inst = new DateTime({
      ts: tsFinal,
      zone: zoneToUse,
      o: offsetFinal,
      loc
    });
    if (normalized.weekday && containsGregor && obj.weekday !== inst.weekday) {
      return DateTime.invalid("mismatched weekday", `you can't specify both a weekday of ${normalized.weekday} and a date of ${inst.toISO()}`);
    }
    return inst;
  }
  static fromISO(text, opts = {}) {
    const [vals, parsedZone] = parseISODate(text);
    return parseDataToDateTime(vals, parsedZone, opts, "ISO 8601", text);
  }
  static fromRFC2822(text, opts = {}) {
    const [vals, parsedZone] = parseRFC2822Date(text);
    return parseDataToDateTime(vals, parsedZone, opts, "RFC 2822", text);
  }
  static fromHTTP(text, opts = {}) {
    const [vals, parsedZone] = parseHTTPDate(text);
    return parseDataToDateTime(vals, parsedZone, opts, "HTTP", opts);
  }
  static fromFormat(text, fmt, opts = {}) {
    if (isUndefined(text) || isUndefined(fmt)) {
      throw new InvalidArgumentError("fromFormat requires an input string and a format");
    }
    const { locale = null, numberingSystem = null } = opts, localeToUse = Locale.fromOpts({
      locale,
      numberingSystem,
      defaultToEN: true
    }), [vals, parsedZone, invalid] = parseFromTokens(localeToUse, text, fmt);
    if (invalid) {
      return DateTime.invalid(invalid);
    } else {
      return parseDataToDateTime(vals, parsedZone, opts, `format ${fmt}`, text);
    }
  }
  static fromString(text, fmt, opts = {}) {
    return DateTime.fromFormat(text, fmt, opts);
  }
  static fromSQL(text, opts = {}) {
    const [vals, parsedZone] = parseSQL(text);
    return parseDataToDateTime(vals, parsedZone, opts, "SQL", text);
  }
  static invalid(reason, explanation = null) {
    if (!reason) {
      throw new InvalidArgumentError("need to specify a reason the DateTime is invalid");
    }
    const invalid = reason instanceof Invalid ? reason : new Invalid(reason, explanation);
    if (Settings.throwOnInvalid) {
      throw new InvalidDateTimeError(invalid);
    } else {
      return new DateTime({ invalid });
    }
  }
  static isDateTime(o) {
    return o && o.isLuxonDateTime || false;
  }
  get(unit) {
    return this[unit];
  }
  get isValid() {
    return this.invalid === null;
  }
  get invalidReason() {
    return this.invalid ? this.invalid.reason : null;
  }
  get invalidExplanation() {
    return this.invalid ? this.invalid.explanation : null;
  }
  get locale() {
    return this.isValid ? this.loc.locale : null;
  }
  get numberingSystem() {
    return this.isValid ? this.loc.numberingSystem : null;
  }
  get outputCalendar() {
    return this.isValid ? this.loc.outputCalendar : null;
  }
  get zone() {
    return this._zone;
  }
  get zoneName() {
    return this.isValid ? this.zone.name : null;
  }
  get year() {
    return this.isValid ? this.c.year : NaN;
  }
  get quarter() {
    return this.isValid ? Math.ceil(this.c.month / 3) : NaN;
  }
  get month() {
    return this.isValid ? this.c.month : NaN;
  }
  get day() {
    return this.isValid ? this.c.day : NaN;
  }
  get hour() {
    return this.isValid ? this.c.hour : NaN;
  }
  get minute() {
    return this.isValid ? this.c.minute : NaN;
  }
  get second() {
    return this.isValid ? this.c.second : NaN;
  }
  get millisecond() {
    return this.isValid ? this.c.millisecond : NaN;
  }
  get weekYear() {
    return this.isValid ? possiblyCachedWeekData(this).weekYear : NaN;
  }
  get weekNumber() {
    return this.isValid ? possiblyCachedWeekData(this).weekNumber : NaN;
  }
  get weekday() {
    return this.isValid ? possiblyCachedWeekData(this).weekday : NaN;
  }
  get ordinal() {
    return this.isValid ? gregorianToOrdinal(this.c).ordinal : NaN;
  }
  get monthShort() {
    return this.isValid ? Info.months("short", { locObj: this.loc })[this.month - 1] : null;
  }
  get monthLong() {
    return this.isValid ? Info.months("long", { locObj: this.loc })[this.month - 1] : null;
  }
  get weekdayShort() {
    return this.isValid ? Info.weekdays("short", { locObj: this.loc })[this.weekday - 1] : null;
  }
  get weekdayLong() {
    return this.isValid ? Info.weekdays("long", { locObj: this.loc })[this.weekday - 1] : null;
  }
  get offset() {
    return this.isValid ? +this.o : NaN;
  }
  get offsetNameShort() {
    if (this.isValid) {
      return this.zone.offsetName(this.ts, {
        format: "short",
        locale: this.locale
      });
    } else {
      return null;
    }
  }
  get offsetNameLong() {
    if (this.isValid) {
      return this.zone.offsetName(this.ts, {
        format: "long",
        locale: this.locale
      });
    } else {
      return null;
    }
  }
  get isOffsetFixed() {
    return this.isValid ? this.zone.universal : null;
  }
  get isInDST() {
    if (this.isOffsetFixed) {
      return false;
    } else {
      return this.offset > this.set({ month: 1 }).offset || this.offset > this.set({ month: 5 }).offset;
    }
  }
  get isInLeapYear() {
    return isLeapYear(this.year);
  }
  get daysInMonth() {
    return daysInMonth(this.year, this.month);
  }
  get daysInYear() {
    return this.isValid ? daysInYear(this.year) : NaN;
  }
  get weeksInWeekYear() {
    return this.isValid ? weeksInWeekYear(this.weekYear) : NaN;
  }
  resolvedLocaleOpts(opts = {}) {
    const { locale, numberingSystem, calendar } = Formatter.create(this.loc.clone(opts), opts).resolvedOptions(this);
    return { locale, numberingSystem, outputCalendar: calendar };
  }
  toUTC(offset2 = 0, opts = {}) {
    return this.setZone(FixedOffsetZone.instance(offset2), opts);
  }
  toLocal() {
    return this.setZone(Settings.defaultZone);
  }
  setZone(zone, { keepLocalTime = false, keepCalendarTime = false } = {}) {
    zone = normalizeZone(zone, Settings.defaultZone);
    if (zone.equals(this.zone)) {
      return this;
    } else if (!zone.isValid) {
      return DateTime.invalid(unsupportedZone(zone));
    } else {
      let newTS = this.ts;
      if (keepLocalTime || keepCalendarTime) {
        const offsetGuess = zone.offset(this.ts);
        const asObj = this.toObject();
        [newTS] = objToTS(asObj, offsetGuess, zone);
      }
      return clone(this, { ts: newTS, zone });
    }
  }
  reconfigure({ locale, numberingSystem, outputCalendar } = {}) {
    const loc = this.loc.clone({ locale, numberingSystem, outputCalendar });
    return clone(this, { loc });
  }
  setLocale(locale) {
    return this.reconfigure({ locale });
  }
  set(values) {
    if (!this.isValid)
      return this;
    const normalized = normalizeObject(values, normalizeUnit, []), settingWeekStuff = !isUndefined(normalized.weekYear) || !isUndefined(normalized.weekNumber) || !isUndefined(normalized.weekday), containsOrdinal = !isUndefined(normalized.ordinal), containsGregorYear = !isUndefined(normalized.year), containsGregorMD = !isUndefined(normalized.month) || !isUndefined(normalized.day), containsGregor = containsGregorYear || containsGregorMD, definiteWeekDef = normalized.weekYear || normalized.weekNumber;
    if ((containsGregor || containsOrdinal) && definiteWeekDef) {
      throw new ConflictingSpecificationError("Can't mix weekYear/weekNumber units with year/month/day or ordinals");
    }
    if (containsGregorMD && containsOrdinal) {
      throw new ConflictingSpecificationError("Can't mix ordinal dates with month/day");
    }
    let mixed;
    if (settingWeekStuff) {
      mixed = weekToGregorian(Object.assign(gregorianToWeek(this.c), normalized));
    } else if (!isUndefined(normalized.ordinal)) {
      mixed = ordinalToGregorian(Object.assign(gregorianToOrdinal(this.c), normalized));
    } else {
      mixed = Object.assign(this.toObject(), normalized);
      if (isUndefined(normalized.day)) {
        mixed.day = Math.min(daysInMonth(mixed.year, mixed.month), mixed.day);
      }
    }
    const [ts, o] = objToTS(mixed, this.o, this.zone);
    return clone(this, { ts, o });
  }
  plus(duration) {
    if (!this.isValid)
      return this;
    const dur = friendlyDuration(duration);
    return clone(this, adjustTime(this, dur));
  }
  minus(duration) {
    if (!this.isValid)
      return this;
    const dur = friendlyDuration(duration).negate();
    return clone(this, adjustTime(this, dur));
  }
  startOf(unit) {
    if (!this.isValid)
      return this;
    const o = {}, normalizedUnit = Duration.normalizeUnit(unit);
    switch (normalizedUnit) {
      case "years":
        o.month = 1;
      case "quarters":
      case "months":
        o.day = 1;
      case "weeks":
      case "days":
        o.hour = 0;
      case "hours":
        o.minute = 0;
      case "minutes":
        o.second = 0;
      case "seconds":
        o.millisecond = 0;
        break;
    }
    if (normalizedUnit === "weeks") {
      o.weekday = 1;
    }
    if (normalizedUnit === "quarters") {
      const q = Math.ceil(this.month / 3);
      o.month = (q - 1) * 3 + 1;
    }
    return this.set(o);
  }
  endOf(unit) {
    return this.isValid ? this.plus({ [unit]: 1 }).startOf(unit).minus(1) : this;
  }
  toFormat(fmt, opts = {}) {
    return this.isValid ? Formatter.create(this.loc.redefaultToEN(opts)).formatDateTimeFromString(this, fmt) : INVALID;
  }
  toLocaleString(opts = DATE_SHORT) {
    return this.isValid ? Formatter.create(this.loc.clone(opts), opts).formatDateTime(this) : INVALID;
  }
  toLocaleParts(opts = {}) {
    return this.isValid ? Formatter.create(this.loc.clone(opts), opts).formatDateTimeParts(this) : [];
  }
  toISO(opts = {}) {
    if (!this.isValid) {
      return null;
    }
    return `${this.toISODate(opts)}T${this.toISOTime(opts)}`;
  }
  toISODate({ format = "extended" } = {}) {
    let fmt = format === "basic" ? "yyyyMMdd" : "yyyy-MM-dd";
    if (this.year > 9999) {
      fmt = "+" + fmt;
    }
    return toTechFormat(this, fmt);
  }
  toISOWeekDate() {
    return toTechFormat(this, "kkkk-'W'WW-c");
  }
  toISOTime({
    suppressMilliseconds = false,
    suppressSeconds = false,
    includeOffset = true,
    includePrefix = false,
    format = "extended"
  } = {}) {
    return toTechTimeFormat(this, {
      suppressSeconds,
      suppressMilliseconds,
      includeOffset,
      includePrefix,
      format
    });
  }
  toRFC2822() {
    return toTechFormat(this, "EEE, dd LLL yyyy HH:mm:ss ZZZ", false);
  }
  toHTTP() {
    return toTechFormat(this.toUTC(), "EEE, dd LLL yyyy HH:mm:ss 'GMT'");
  }
  toSQLDate() {
    return toTechFormat(this, "yyyy-MM-dd");
  }
  toSQLTime({ includeOffset = true, includeZone = false } = {}) {
    return toTechTimeFormat(this, {
      includeOffset,
      includeZone,
      spaceZone: true
    });
  }
  toSQL(opts = {}) {
    if (!this.isValid) {
      return null;
    }
    return `${this.toSQLDate()} ${this.toSQLTime(opts)}`;
  }
  toString() {
    return this.isValid ? this.toISO() : INVALID;
  }
  valueOf() {
    return this.toMillis();
  }
  toMillis() {
    return this.isValid ? this.ts : NaN;
  }
  toSeconds() {
    return this.isValid ? this.ts / 1e3 : NaN;
  }
  toJSON() {
    return this.toISO();
  }
  toBSON() {
    return this.toJSDate();
  }
  toObject(opts = {}) {
    if (!this.isValid)
      return {};
    const base = Object.assign({}, this.c);
    if (opts.includeConfig) {
      base.outputCalendar = this.outputCalendar;
      base.numberingSystem = this.loc.numberingSystem;
      base.locale = this.loc.locale;
    }
    return base;
  }
  toJSDate() {
    return new Date(this.isValid ? this.ts : NaN);
  }
  diff(otherDateTime, unit = "milliseconds", opts = {}) {
    if (!this.isValid || !otherDateTime.isValid) {
      return Duration.invalid(this.invalid || otherDateTime.invalid, "created by diffing an invalid DateTime");
    }
    const durOpts = Object.assign({ locale: this.locale, numberingSystem: this.numberingSystem }, opts);
    const units = maybeArray(unit).map(Duration.normalizeUnit), otherIsLater = otherDateTime.valueOf() > this.valueOf(), earlier = otherIsLater ? this : otherDateTime, later = otherIsLater ? otherDateTime : this, diffed = diff(earlier, later, units, durOpts);
    return otherIsLater ? diffed.negate() : diffed;
  }
  diffNow(unit = "milliseconds", opts = {}) {
    return this.diff(DateTime.now(), unit, opts);
  }
  until(otherDateTime) {
    return this.isValid ? Interval.fromDateTimes(this, otherDateTime) : this;
  }
  hasSame(otherDateTime, unit) {
    if (!this.isValid)
      return false;
    const inputMs = otherDateTime.valueOf();
    const otherZoneDateTime = this.setZone(otherDateTime.zone, { keepLocalTime: true });
    return otherZoneDateTime.startOf(unit) <= inputMs && inputMs <= otherZoneDateTime.endOf(unit);
  }
  equals(other) {
    return this.isValid && other.isValid && this.valueOf() === other.valueOf() && this.zone.equals(other.zone) && this.loc.equals(other.loc);
  }
  toRelative(options = {}) {
    if (!this.isValid)
      return null;
    const base = options.base || DateTime.fromObject({ zone: this.zone }), padding = options.padding ? this < base ? -options.padding : options.padding : 0;
    let units = ["years", "months", "days", "hours", "minutes", "seconds"];
    let unit = options.unit;
    if (Array.isArray(options.unit)) {
      units = options.unit;
      unit = void 0;
    }
    return diffRelative(base, this.plus(padding), Object.assign(options, {
      numeric: "always",
      units,
      unit
    }));
  }
  toRelativeCalendar(options = {}) {
    if (!this.isValid)
      return null;
    return diffRelative(options.base || DateTime.fromObject({ zone: this.zone }), this, Object.assign(options, {
      numeric: "auto",
      units: ["years", "months", "days"],
      calendary: true
    }));
  }
  static min(...dateTimes) {
    if (!dateTimes.every(DateTime.isDateTime)) {
      throw new InvalidArgumentError("min requires all arguments be DateTimes");
    }
    return bestBy(dateTimes, (i) => i.valueOf(), Math.min);
  }
  static max(...dateTimes) {
    if (!dateTimes.every(DateTime.isDateTime)) {
      throw new InvalidArgumentError("max requires all arguments be DateTimes");
    }
    return bestBy(dateTimes, (i) => i.valueOf(), Math.max);
  }
  static fromFormatExplain(text, fmt, options = {}) {
    const { locale = null, numberingSystem = null } = options, localeToUse = Locale.fromOpts({
      locale,
      numberingSystem,
      defaultToEN: true
    });
    return explainFromTokens(localeToUse, text, fmt);
  }
  static fromStringExplain(text, fmt, options = {}) {
    return DateTime.fromFormatExplain(text, fmt, options);
  }
  static get DATE_SHORT() {
    return DATE_SHORT;
  }
  static get DATE_MED() {
    return DATE_MED;
  }
  static get DATE_MED_WITH_WEEKDAY() {
    return DATE_MED_WITH_WEEKDAY;
  }
  static get DATE_FULL() {
    return DATE_FULL;
  }
  static get DATE_HUGE() {
    return DATE_HUGE;
  }
  static get TIME_SIMPLE() {
    return TIME_SIMPLE;
  }
  static get TIME_WITH_SECONDS() {
    return TIME_WITH_SECONDS;
  }
  static get TIME_WITH_SHORT_OFFSET() {
    return TIME_WITH_SHORT_OFFSET;
  }
  static get TIME_WITH_LONG_OFFSET() {
    return TIME_WITH_LONG_OFFSET;
  }
  static get TIME_24_SIMPLE() {
    return TIME_24_SIMPLE;
  }
  static get TIME_24_WITH_SECONDS() {
    return TIME_24_WITH_SECONDS;
  }
  static get TIME_24_WITH_SHORT_OFFSET() {
    return TIME_24_WITH_SHORT_OFFSET;
  }
  static get TIME_24_WITH_LONG_OFFSET() {
    return TIME_24_WITH_LONG_OFFSET;
  }
  static get DATETIME_SHORT() {
    return DATETIME_SHORT;
  }
  static get DATETIME_SHORT_WITH_SECONDS() {
    return DATETIME_SHORT_WITH_SECONDS;
  }
  static get DATETIME_MED() {
    return DATETIME_MED;
  }
  static get DATETIME_MED_WITH_SECONDS() {
    return DATETIME_MED_WITH_SECONDS;
  }
  static get DATETIME_MED_WITH_WEEKDAY() {
    return DATETIME_MED_WITH_WEEKDAY;
  }
  static get DATETIME_FULL() {
    return DATETIME_FULL;
  }
  static get DATETIME_FULL_WITH_SECONDS() {
    return DATETIME_FULL_WITH_SECONDS;
  }
  static get DATETIME_HUGE() {
    return DATETIME_HUGE;
  }
  static get DATETIME_HUGE_WITH_SECONDS() {
    return DATETIME_HUGE_WITH_SECONDS;
  }
}
function friendlyDateTime(dateTimeish) {
  if (DateTime.isDateTime(dateTimeish)) {
    return dateTimeish;
  } else if (dateTimeish && dateTimeish.valueOf && isNumber(dateTimeish.valueOf())) {
    return DateTime.fromJSDate(dateTimeish);
  } else if (dateTimeish && typeof dateTimeish === "object") {
    return DateTime.fromObject(dateTimeish);
  } else {
    throw new InvalidArgumentError(`Unknown datetime argument: ${dateTimeish}, of type ${typeof dateTimeish}`);
  }
}
function hasFlag(val, flag) {
  return (val & flag) === flag;
}
function getValueByPath(obj, path, defaultValue = void 0) {
  const value = path.split(".").reduce((o, i) => typeof o !== "undefined" ? o[i] : void 0, obj);
  return typeof value !== "undefined" ? value : defaultValue;
}
function indexOf(array, obj, fn) {
  if (!array)
    return -1;
  if (!fn || typeof fn !== "function")
    return array.indexOf(obj);
  for (let i = 0; i < array.length; i++) {
    if (fn(array[i], obj)) {
      return i;
    }
  }
  return -1;
}
const isObject = (item) => typeof item === "object" && !Array.isArray(item);
const mergeFn = (target, source, deep = false) => {
  if (deep || !Object.assign) {
    const isDeep = (prop) => isObject(source[prop]) && target !== null && Object.prototype.hasOwnProperty.call(target, prop) && isObject(target[prop]);
    let replaced;
    if (source === null || typeof source === "undefined") {
      replaced = false;
    } else {
      replaced = Object.getOwnPropertyNames(source).map((prop) => ({ [prop]: isDeep(prop) ? mergeFn(target[prop], source[prop], deep) : source[prop] })).reduce((a, b) => __spreadValues(__spreadValues({}, a), b), {});
    }
    return __spreadValues(__spreadValues({}, target), replaced);
  } else {
    return Object.assign(target, source);
  }
};
const merge = mergeFn;
const isMobile = {
  Android: function() {
    return typeof window !== "undefined" && window.navigator.userAgent.match(/Android/i);
  },
  BlackBerry: function() {
    return typeof window !== "undefined" && window.navigator.userAgent.match(/BlackBerry/i);
  },
  iOS: function() {
    return typeof window !== "undefined" && window.navigator.userAgent.match(/iPhone|iPad|iPod/i);
  },
  Opera: function() {
    return typeof window !== "undefined" && window.navigator.userAgent.match(/Opera Mini/i);
  },
  Windows: function() {
    return typeof window !== "undefined" && window.navigator.userAgent.match(/IEMobile/i);
  },
  any: function() {
    return isMobile.Android() || isMobile.BlackBerry() || isMobile.iOS() || isMobile.Opera() || isMobile.Windows();
  }
};
function removeElement(el) {
  if (typeof el.remove !== "undefined") {
    el.remove();
  } else if (typeof el.parentNode !== "undefined" && el.parentNode !== null) {
    el.parentNode.removeChild(el);
  }
}
function createAbsoluteElement(el) {
  const root = document.createElement("div");
  root.style.position = "absolute";
  root.style.left = "0px";
  root.style.top = "0px";
  const wrapper = document.createElement("div");
  root.appendChild(wrapper);
  wrapper.appendChild(el);
  document.body.appendChild(root);
  return root;
}
function escapeRegExpChars(value) {
  if (!value)
    return value;
  return value.replace(/[\-\[\]\/\{\}\(\)\*\+\?\.\\\^\$\|]/g, "\\$&");
}
function toCssDimension(width) {
  return width === void 0 ? null : isNaN(width) ? width : width + "px";
}
function blankIfUndefined(value) {
  return typeof value !== "undefined" && value !== null ? value : "";
}
function getMonthNames(locale = void 0, format = "long") {
  const dates = [];
  for (let i = 0; i < 12; i++) {
    dates.push(new Date(2e3, i, 15));
  }
  const dtf = new Intl.DateTimeFormat(locale, {
    month: format
  });
  return dates.map((d) => dtf.format(d));
}
function getWeekdayNames(locale = void 0, firstDayOfWeek = 0, format = "narrow") {
  const dates = [];
  for (let i = 1, j = 0; j < 7; i++) {
    const d = new Date(2e3, 0, i);
    const day = d.getDay();
    if (day === firstDayOfWeek || j > 0) {
      dates.push(d);
      j++;
    }
  }
  const dtf = new Intl.DateTimeFormat(locale, {
    weekday: format
  });
  return dates.map((d) => dtf.format(d));
}
function matchWithGroups(pattern, str) {
  const matches = str.match(pattern);
  return pattern.toString().match(/<(.+?)>/g).map((group) => {
    const groupMatches = group.match(/<(.+)>/);
    if (!groupMatches || groupMatches.length <= 0) {
      return null;
    }
    return group.match(/<(.+)>/)[1];
  }).reduce((acc, curr, index2) => {
    if (matches && matches.length > index2) {
      acc[curr] = matches[index2 + 1];
    } else {
      acc[curr] = null;
    }
    return acc;
  }, {});
}
function debounce(func, wait, immediate) {
  let timeout;
  return function() {
    const context = this;
    const args = arguments;
    const later = function() {
      timeout = null;
      if (!immediate)
        func.apply(context, args);
    };
    const callNow = immediate && !timeout;
    clearTimeout(timeout);
    timeout = setTimeout(later, wait);
    if (callNow)
      func.apply(context, args);
  };
}
function endsWith(str, suffix) {
  return str.indexOf(suffix, str.length - suffix.length) !== -1;
}
let config$1 = {
  iconPack: "mdi",
  useHtml5Validation: true,
  statusIcon: true,
  transformClasses: void 0
};
const setOptions = (options) => {
  config$1 = options;
};
const getOptions$1 = () => {
  return config$1;
};
let VueInstance;
const setVueInstance = (Vue) => {
  VueInstance = Vue;
};
const Programmatic = {
  getOptions: getOptions$1,
  setOptions(options) {
    setOptions(merge(getOptions$1(), options, true));
  }
};
const _defaultSuffixProcessor = (input, suffix) => {
  return blankIfUndefined(input).split(" ").filter((cls) => cls.length > 0).map((cls) => cls + suffix).join(" ");
};
const _getContext = (vm) => {
  const computedNames = vm.$options.computed ? Object.keys(vm.$options.computed) : [];
  const computed = computedNames.filter((e) => !endsWith(e, "Classes")).reduce((o, key) => {
    o[key] = vm[key];
    return o;
  }, {});
  return { props: vm.$props, data: vm.$data, computed };
};
var BaseComponentMixin = vue.defineComponent({
  isOruga: true,
  props: {
    override: Boolean
  },
  methods: {
    computedClass(field, defaultValue, suffix = "") {
      const config2 = this.$props.override === true ? {} : getOptions$1();
      const override = this.$props.override || getValueByPath(config2, `${this.$options.configField}.override`, false);
      const overrideClass = getValueByPath(config2, `${this.$options.configField}.${field}.override`, override);
      const globalTransformClasses = getValueByPath(config2, `transformClasses`, void 0);
      const localTransformClasses = getValueByPath(config2, `${this.$options.configField}.transformClasses`, void 0);
      let globalClass = getValueByPath(config2, `${this.$options.configField}.${field}.class`, "") || getValueByPath(config2, `${this.$options.configField}.${field}`, "");
      let currentClass = getValueByPath(this.$props, field);
      if (Array.isArray(currentClass)) {
        currentClass = currentClass.join(" ");
      }
      if (defaultValue.search("{*}") !== -1) {
        defaultValue = defaultValue.replace(/\{\*\}/g, suffix);
      } else {
        defaultValue = defaultValue + suffix;
      }
      let context = null;
      if (typeof currentClass === "function") {
        context = _getContext(this);
        currentClass = currentClass(suffix, context);
      } else {
        currentClass = _defaultSuffixProcessor(currentClass, suffix);
      }
      if (typeof globalClass === "function") {
        globalClass = globalClass(suffix, context || _getContext(this));
      } else {
        globalClass = _defaultSuffixProcessor(globalClass, suffix);
      }
      let appliedClasses = `${override && !overrideClass || !override && !overrideClass ? defaultValue : ""} ${blankIfUndefined(globalClass)} ${blankIfUndefined(currentClass)}`.trim().replace(/\s\s+/g, " ");
      if (localTransformClasses) {
        appliedClasses = localTransformClasses(appliedClasses);
      }
      if (globalTransformClasses) {
        appliedClasses = globalTransformClasses(appliedClasses);
      }
      return appliedClasses;
    }
  }
});
const oruga = {};
function addProgrammatic(property, component) {
  oruga[property] = component;
}
function useProgrammatic() {
  return { oruga, addProgrammatic };
}
const registerPlugin = (app, plugin) => {
  app.use(plugin);
};
const registerComponent = (app, component) => {
  app.component(component.name, component);
};
const registerComponentProgrammatic = (app, property, component) => {
  const { oruga: oruga2, addProgrammatic: addProgrammatic2 } = useProgrammatic();
  addProgrammatic2(property, component);
  if (!(app._context.provides && app._context.provides.oruga))
    app.provide("oruga", oruga2);
  if (!app.config.globalProperties.$oruga)
    app.config.globalProperties.$oruga = oruga2;
};
const mdiIcons = {
  sizes: {
    "default": "mdi-24px",
    "small": null,
    "medium": "mdi-36px",
    "large": "mdi-48px"
  },
  iconPrefix: "mdi-"
};
const faIcons = () => {
  const iconComponent = getValueByPath(getOptions$1(), "iconComponent");
  const faIconPrefix = iconComponent ? "" : "fa-";
  return {
    sizes: {
      "default": null,
      "small": null,
      "medium": faIconPrefix + "lg",
      "large": faIconPrefix + "2x"
    },
    iconPrefix: faIconPrefix,
    internalIcons: {
      "check": "check",
      "information": "info-circle",
      "alert": "exclamation-triangle",
      "alert-circle": "exclamation-circle",
      "arrow-up": "arrow-up",
      "chevron-right": "angle-right",
      "chevron-left": "angle-left",
      "chevron-down": "angle-down",
      "chevron-up": "angle-up",
      "eye": "eye",
      "eye-off": "eye-slash",
      "caret-down": "caret-down",
      "caret-up": "caret-up",
      "close-circle": "times-circle",
      "loading": "circle-notch"
    }
  };
};
const getIcons = () => {
  let icons = {
    mdi: mdiIcons,
    fa: faIcons(),
    fas: faIcons(),
    far: faIcons(),
    fad: faIcons(),
    fab: faIcons(),
    fal: faIcons()
  };
  const customIconPacks = getValueByPath(getOptions$1(), "customIconPacks");
  if (customIconPacks) {
    icons = merge(icons, customIconPacks, true);
  }
  return icons;
};
var script$r = vue.defineComponent({
  name: "OIcon",
  mixins: [BaseComponentMixin],
  configField: "icon",
  props: {
    variant: [String, Object],
    component: String,
    pack: String,
    icon: String,
    size: String,
    customSize: String,
    customClass: String,
    clickable: Boolean,
    spin: Boolean,
    rotation: [Number, String],
    both: Boolean,
    rootClass: [String, Function, Array],
    clickableClass: [String, Function, Array],
    spinClass: [String, Function, Array],
    sizeClass: [String, Function, Array],
    variantClass: [String, Function, Array]
  },
  computed: {
    rootClasses() {
      return [
        this.computedClass("rootClass", "o-icon"),
        { [this.computedClass("clickableClass", "o-icon--clickable")]: this.clickable },
        { [this.computedClass("spinClass", "o-icon--spin")]: this.spin },
        { [this.computedClass("sizeClass", "o-icon--", this.size)]: this.size },
        { [this.computedClass("variantClass", "o-icon--", this.newVariant)]: this.newVariant }
      ];
    },
    rootStyle() {
      const style = {};
      if (this.rotation) {
        style["transform"] = `rotate(${this.rotation}deg)`;
      }
      return style;
    },
    iconConfig() {
      return getIcons()[this.newPack];
    },
    iconPrefix() {
      if (this.iconConfig && this.iconConfig.iconPrefix) {
        return this.iconConfig.iconPrefix;
      }
      return "";
    },
    newIcon() {
      return `${this.iconPrefix}${this.getEquivalentIconOf(this.icon)}`;
    },
    newPack() {
      return this.pack || getValueByPath(getOptions$1(), "iconPack", "mdi");
    },
    newVariant() {
      if (!this.variant)
        return;
      let newVariant = "";
      if (typeof this.variant === "string") {
        newVariant = this.variant;
      } else {
        newVariant = Object.keys(this.variant).filter((key) => this.variant[key])[0];
      }
      return newVariant;
    },
    newCustomSize() {
      return this.customSize || this.customSizeByPack;
    },
    customSizeByPack() {
      if (this.iconConfig && this.iconConfig.sizes) {
        if (this.size && this.iconConfig.sizes[this.size] !== void 0) {
          return this.iconConfig.sizes[this.size];
        } else if (this.iconConfig.sizes.default) {
          return this.iconConfig.sizes.default;
        }
      }
      return null;
    },
    useIconComponent() {
      if (this.component)
        return this.component;
      const component = getValueByPath(getOptions$1(), "iconComponent");
      if (component)
        return component;
      return null;
    }
  },
  methods: {
    getEquivalentIconOf(value) {
      if (!this.both) {
        return value;
      }
      if (this.iconConfig && this.iconConfig.internalIcons && this.iconConfig.internalIcons[value]) {
        return this.iconConfig.internalIcons[value];
      }
      return value;
    }
  }
});
function render$o(_ctx, _cache, $props, $setup, $data, $options) {
  return vue.openBlock(), vue.createBlock("span", {
    class: _ctx.rootClasses,
    style: _ctx.rootStyle
  }, [!_ctx.useIconComponent ? (vue.openBlock(), vue.createBlock("i", {
    key: 0,
    class: [_ctx.newPack, _ctx.newIcon, _ctx.newCustomSize, _ctx.customClass]
  }, null, 2)) : (vue.openBlock(), vue.createBlock(vue.Fragment, {
    key: 1
  }, [vue.createCommentVNode(" custom icon component "), (vue.openBlock(), vue.createBlock(vue.resolveDynamicComponent(_ctx.useIconComponent), {
    icon: [_ctx.newPack, _ctx.newIcon],
    size: _ctx.newCustomSize,
    class: [_ctx.customClass]
  }, null, 8, ["icon", "size", "class"]))], 64))], 6);
}
script$r.render = render$o;
script$r.__file = "src/components/icon/Icon.vue";
var FormElementMixin = vue.defineComponent({
  inject: {
    $field: { from: "$field", default: false },
    $elementRef: { from: "$elementRef", default: false }
  },
  emits: ["blur", "focus"],
  props: {
    expanded: Boolean,
    rounded: Boolean,
    icon: String,
    iconPack: String,
    autocomplete: String,
    maxlength: [Number, String],
    useHtml5Validation: {
      type: Boolean,
      default: () => {
        return getValueByPath(getOptions$1(), "useHtml5Validation", true);
      }
    },
    statusIcon: {
      type: Boolean,
      default: () => {
        return getValueByPath(getOptions$1(), "statusIcon", true);
      }
    },
    validationMessage: String
  },
  data() {
    return {
      isValid: true,
      isFocused: false,
      newIconPack: this.iconPack
    };
  },
  computed: {
    parentField() {
      return this.$field;
    },
    statusVariant() {
      if (!this.parentField)
        return;
      if (!this.parentField.newVariant)
        return;
      if (typeof this.parentField.newVariant === "string") {
        return this.parentField.newVariant;
      } else {
        for (const key in this.parentField.newVariant) {
          if (this.parentField.newVariant[key]) {
            return key;
          }
        }
      }
    },
    statusMessage() {
      if (!this.parentField)
        return;
      return this.parentField.newMessage || this.parentField.hasMessageSlot;
    },
    statusVariantIcon() {
      const statusVariantIcon = getValueByPath(getOptions$1(), "statusVariantIcon", {
        "success": "check",
        "danger": "alert-circle",
        "info": "information",
        "warning": "alert"
      });
      return statusVariantIcon[this.statusVariant] || "";
    }
  },
  methods: {
    focus() {
      const el = this.getElement();
      if (el === void 0)
        return;
      this.$nextTick(() => {
        if (el)
          el.focus();
      });
    },
    onBlur(event) {
      this.isFocused = false;
      if (this.parentField) {
        this.parentField.isFocused = false;
      }
      this.$emit("blur", event);
      this.checkHtml5Validity();
    },
    onFocus(event) {
      this.isFocused = true;
      if (this.parentField) {
        this.parentField.isFocused = true;
      }
      this.$emit("focus", event);
    },
    getElement() {
      let el = this.$refs[this.$elementRef];
      while (el && el.$elementRef) {
        el = el.$refs[el.$elementRef];
      }
      return el;
    },
    setInvalid() {
      const variant = "danger";
      const message = this.validationMessage || this.getElement().validationMessage;
      this.setValidity(variant, message);
    },
    setValidity(variant, message) {
      this.$nextTick(() => {
        if (this.parentField) {
          if (!this.parentField.variant) {
            this.parentField.newVariant = variant;
          }
          if (!this.parentField.message) {
            this.parentField.newMessage = message;
          }
        }
      });
    },
    checkHtml5Validity() {
      if (!this.useHtml5Validation)
        return;
      const el = this.getElement();
      if (el === void 0)
        return;
      if (!el.checkValidity()) {
        this.setInvalid();
        this.isValid = false;
      } else {
        this.setValidity(null, null);
        this.isValid = true;
      }
      return this.isValid;
    },
    syncFilled(value) {
      if (this.parentField) {
        this.parentField.isFilled = !!value;
      }
    }
  }
});
var script$q = vue.defineComponent({
  name: "OInput",
  components: {
    [script$r.name]: script$r
  },
  mixins: [BaseComponentMixin, FormElementMixin],
  configField: "input",
  inheritAttrs: false,
  provide() {
    return {
      $elementRef: this.type === "textarea" ? "textarea" : "input"
    };
  },
  emits: ["update:modelValue", "icon-click", "icon-right-click"],
  props: {
    modelValue: [Number, String],
    type: {
      type: String,
      default: "text"
    },
    size: String,
    passwordReveal: Boolean,
    iconClickable: Boolean,
    hasCounter: {
      type: Boolean,
      default: () => {
        return getValueByPath(getOptions$1(), "input.counter", false);
      }
    },
    iconRight: String,
    iconRightClickable: Boolean,
    iconRightVariant: String,
    clearable: {
      type: Boolean,
      default: () => {
        return getValueByPath(getOptions$1(), "input.clearable", false);
      }
    },
    rootClass: [String, Function, Array],
    expandedClass: [String, Function, Array],
    iconLeftSpaceClass: [String, Function, Array],
    iconRightSpaceClass: [String, Function, Array],
    inputClass: [String, Function, Array],
    roundedClass: [String, Function, Array],
    iconLeftClass: [String, Function, Array],
    iconRightClass: [String, Function, Array],
    counterClass: [String, Function, Array],
    sizeClass: [String, Function, Array],
    variantClass: [String, Function, Array]
  },
  data() {
    return {
      newValue: this.modelValue,
      newType: this.type,
      newAutocomplete: this.autocomplete || getValueByPath(getOptions$1(), "input.autocompletete", "off"),
      isPasswordVisible: false
    };
  },
  computed: {
    rootClasses() {
      return [
        this.computedClass("rootClass", "o-ctrl-input"),
        { [this.computedClass("expandedClass", "o-ctrl-input--expanded")]: this.expanded }
      ];
    },
    inputClasses() {
      return [
        this.computedClass("inputClass", "o-input"),
        { [this.computedClass("roundedClass", "o-input--rounded")]: this.rounded },
        { [this.computedClass("sizeClass", "o-input--", this.size)]: this.size },
        { [this.computedClass("variantClass", "o-input--", this.statusVariant)]: this.statusVariant },
        { [this.computedClass("textareaClass", "o-input__textarea")]: this.type === "textarea" },
        { [this.computedClass("iconLeftSpaceClass", "o-input-iconspace-left")]: this.icon },
        { [this.computedClass("iconRightSpaceClass", "o-input-iconspace-right")]: this.hasIconRight }
      ];
    },
    iconLeftClasses() {
      return [
        this.computedClass("iconLeftClass", "o-input__icon-left")
      ];
    },
    iconRightClasses() {
      return [
        this.computedClass("iconRightClass", "o-input__icon-right")
      ];
    },
    counterClasses() {
      return [
        this.computedClass("counterClass", "o-input__counter")
      ];
    },
    computedValue: {
      get() {
        return this.newValue;
      },
      set(value) {
        this.newValue = value;
        this.$emit("update:modelValue", this.newValue);
        this.syncFilled(this.newValue);
        !this.isValid && this.checkHtml5Validity();
      }
    },
    hasIconRight() {
      return this.passwordReveal || this.statusIcon && this.statusVariantIcon || this.clearable || this.iconRight;
    },
    rightIcon() {
      if (this.passwordReveal) {
        return this.passwordVisibleIcon;
      } else if (this.clearable && this.newValue) {
        return "close-circle";
      } else if (this.iconRight) {
        return this.iconRight;
      }
      return this.statusVariantIcon;
    },
    rightIconVariant() {
      if (this.passwordReveal || this.iconRight) {
        return this.iconRightVariant || null;
      }
      return this.statusVariant;
    },
    hasMessage() {
      return !!this.statusMessage;
    },
    passwordVisibleIcon() {
      return !this.isPasswordVisible ? "eye" : "eye-off";
    },
    valueLength() {
      if (typeof this.computedValue === "string") {
        return this.computedValue.length;
      } else if (typeof this.computedValue === "number") {
        return this.computedValue.toString().length;
      }
      return 0;
    }
  },
  watch: {
    modelValue: {
      immediate: true,
      handler(value) {
        this.newValue = value;
        this.syncFilled(this.newValue);
      }
    }
  },
  methods: {
    togglePasswordVisibility() {
      this.isPasswordVisible = !this.isPasswordVisible;
      this.newType = this.isPasswordVisible ? "text" : "password";
      this.$nextTick(() => {
        this.focus();
      });
    },
    iconClick(emit, event) {
      this.$emit(emit, event);
      this.$nextTick(() => {
        this.focus();
      });
    },
    rightIconClick(event) {
      if (this.passwordReveal) {
        this.togglePasswordVisibility();
      } else if (this.clearable) {
        this.computedValue = "";
      } else if (this.iconRightClickable) {
        this.iconClick("icon-right-click", event);
      }
    }
  }
});
function render$n(_ctx, _cache, $props, $setup, $data, $options) {
  const _component_o_icon = vue.resolveComponent("o-icon");
  return vue.openBlock(), vue.createBlock("div", {
    class: _ctx.rootClasses
  }, [_ctx.type !== "textarea" ? vue.withDirectives((vue.openBlock(), vue.createBlock("input", vue.mergeProps({
    key: 0
  }, _ctx.$attrs, {
    ref: "input",
    class: _ctx.inputClasses,
    type: _ctx.newType,
    autocomplete: _ctx.newAutocomplete,
    maxlength: _ctx.maxlength,
    "onUpdate:modelValue": _cache[1] || (_cache[1] = ($event) => _ctx.computedValue = $event),
    onBlur: _cache[2] || (_cache[2] = (...args) => _ctx.onBlur(...args)),
    onFocus: _cache[3] || (_cache[3] = (...args) => _ctx.onFocus(...args))
  }), null, 16, ["type", "autocomplete", "maxlength"])), [[vue.vModelDynamic, _ctx.computedValue]]) : vue.withDirectives((vue.openBlock(), vue.createBlock("textarea", vue.mergeProps({
    key: 1
  }, _ctx.$attrs, {
    ref: "textarea",
    class: _ctx.inputClasses,
    maxlength: _ctx.maxlength,
    "onUpdate:modelValue": _cache[4] || (_cache[4] = ($event) => _ctx.computedValue = $event),
    onBlur: _cache[5] || (_cache[5] = (...args) => _ctx.onBlur(...args)),
    onFocus: _cache[6] || (_cache[6] = (...args) => _ctx.onFocus(...args))
  }), null, 16, ["maxlength"])), [[vue.vModelText, _ctx.computedValue]]), _ctx.icon ? vue.createVNode(_component_o_icon, {
    key: 2,
    class: _ctx.iconLeftClasses,
    clickable: _ctx.iconClickable,
    icon: _ctx.icon,
    pack: _ctx.iconPack,
    size: _ctx.size,
    onClick: _cache[7] || (_cache[7] = ($event) => _ctx.iconClick("icon-click", $event))
  }, null, 8, ["class", "clickable", "icon", "pack", "size"]) : vue.createCommentVNode("v-if", true), _ctx.hasIconRight ? vue.createVNode(_component_o_icon, {
    key: 3,
    class: _ctx.iconRightClasses,
    clickable: _ctx.passwordReveal || _ctx.iconRightClickable,
    icon: _ctx.rightIcon,
    pack: _ctx.iconPack,
    size: _ctx.size,
    variant: _ctx.rightIconVariant,
    both: "",
    onClick: _ctx.rightIconClick
  }, null, 8, ["class", "clickable", "icon", "pack", "size", "variant", "onClick"]) : vue.createCommentVNode("v-if", true), _ctx.maxlength && _ctx.hasCounter && _ctx.isFocused && _ctx.type !== "number" ? (vue.openBlock(), vue.createBlock("small", {
    key: 4,
    class: _ctx.counterClasses
  }, vue.toDisplayString(_ctx.valueLength) + " / " + vue.toDisplayString(_ctx.maxlength), 3)) : vue.createCommentVNode("v-if", true)], 2);
}
script$q.render = render$n;
script$q.__file = "src/components/input/Input.vue";
let config = {
  iconPack: "mdi",
  useHtml5Validation: true,
  statusIcon: true
};
const getOptions = () => {
  return config;
};
var script$p = vue.defineComponent({
  name: "OAutocomplete",
  configField: "autocomplete",
  components: {
    [script$q.name]: script$q
  },
  mixins: [BaseComponentMixin, FormElementMixin],
  inheritAttrs: false,
  provide() {
    return {
      $elementRef: "input"
    };
  },
  emits: ["update:modelValue", "select", "infinite-scroll", "typing", "focus", "blur", "icon-click", "icon-right-click"],
  props: {
    modelValue: [Number, String],
    data: {
      type: Array,
      default: () => []
    },
    size: String,
    field: {
      type: String,
      default: "value"
    },
    keepFirst: Boolean,
    clearOnSelect: Boolean,
    openOnFocus: Boolean,
    customFormatter: Function,
    checkInfiniteScroll: Boolean,
    keepOpen: Boolean,
    clearable: Boolean,
    maxHeight: [String, Number],
    menuPosition: {
      type: String,
      default: "auto"
    },
    animation: {
      type: String,
      default: () => {
        return getValueByPath(getOptions(), "autocomplete.animation", "fade");
      }
    },
    groupField: String,
    groupOptions: String,
    debounceTyping: Number,
    iconRight: String,
    iconRightClickable: Boolean,
    appendToBody: Boolean,
    confirmKeys: {
      type: Array,
      default: () => ["Tab", "Enter"]
    },
    type: {
      type: String,
      default: "text"
    },
    selectOnClickOutside: Boolean,
    selectableHeader: Boolean,
    selectableFooter: Boolean,
    rootClass: [String, Function, Array],
    menuClass: [String, Function, Array],
    expandedClass: [String, Function, Array],
    menuPositionClass: [String, Function, Array],
    itemClass: [String, Function, Array],
    itemHoverClass: [String, Function, Array],
    itemGroupTitleClass: [String, Function, Array],
    itemEmptyClass: [String, Function, Array],
    inputClasses: Object
  },
  data() {
    return {
      selected: null,
      hovered: null,
      headerHovered: null,
      footerHovered: null,
      isActive: false,
      newValue: this.modelValue,
      newAutocomplete: this.autocomplete || "off",
      isListInViewportVertically: true,
      hasFocus: false,
      itemRefs: [],
      width: void 0,
      bodyEl: void 0
    };
  },
  computed: {
    rootClasses() {
      return [
        this.computedClass("rootClass", "o-acp"),
        { [this.computedClass("expandedClass", "o-acp--expanded")]: this.expanded }
      ];
    },
    menuClasses() {
      return [
        this.computedClass("menuClass", "o-acp__menu"),
        { [this.computedClass("menuPositionClass", "o-acp__menu--", this.newDropdownPosition)]: !this.appendToBody }
      ];
    },
    itemClasses() {
      return [
        this.computedClass("itemClass", "o-acp__item")
      ];
    },
    itemEmptyClasses() {
      return [
        ...this.itemClasses,
        this.computedClass("itemEmptyClass", "o-acp__item--empty")
      ];
    },
    itemGroupClasses() {
      return [
        ...this.itemClasses,
        this.computedClass("itemGroupTitleClass", "o-acp__item-group-title")
      ];
    },
    itemHeaderClasses() {
      return [
        ...this.itemClasses,
        { [this.computedClass("itemHoverClass", "o-acp__item--hover")]: this.headerHovered }
      ];
    },
    itemFooterClasses() {
      return [
        ...this.itemClasses,
        { [this.computedClass("itemHoverClass", "o-acp__item--hover")]: this.footerHovered }
      ];
    },
    inputBind() {
      return __spreadValues(__spreadValues({}, this.$attrs), this.inputClasses);
    },
    computedData() {
      if (this.groupField) {
        if (this.groupOptions) {
          const newData = [];
          this.data.forEach((option) => {
            const group = getValueByPath(option, this.groupField);
            const items2 = getValueByPath(option, this.groupOptions);
            newData.push({ group, items: items2 });
          });
          return newData;
        } else {
          const tmp = {};
          this.data.forEach((option) => {
            const group = getValueByPath(option, this.groupField);
            if (!tmp[group])
              tmp[group] = [];
            tmp[group].push(option);
          });
          const newData = [];
          Object.keys(this.data).forEach((group) => {
            newData.push({ group, items: this.data[group] });
          });
          return newData;
        }
      }
      return [{ items: this.data }];
    },
    isEmpty() {
      if (!this.computedData)
        return true;
      return !this.computedData.some((element) => element.items && element.items.length);
    },
    whiteList() {
      const whiteList = [];
      whiteList.push(this.$refs.input.$el.querySelector("input"));
      whiteList.push(this.$refs.dropdown);
      if (this.$refs.dropdown !== void 0) {
        const children = this.$refs.dropdown.querySelectorAll("*");
        for (const child of children) {
          whiteList.push(child);
        }
      }
      return whiteList;
    },
    newDropdownPosition() {
      if (this.menuPosition === "top" || this.menuPosition === "auto" && !this.isListInViewportVertically) {
        return "top";
      }
      return "bottom";
    },
    newIconRight() {
      if (this.clearable && this.newValue) {
        return "close-circle";
      }
      return this.iconRight;
    },
    newIconRightClickable() {
      if (this.clearable) {
        return true;
      }
      return this.iconRightClickable;
    },
    menuStyle() {
      return {
        maxHeight: toCssDimension(this.maxHeight)
      };
    }
  },
  watch: {
    modelValue(value) {
      this.newValue = value;
    },
    isActive(active) {
      if (this.menuPosition === "auto") {
        if (active) {
          this.calcDropdownInViewportVertical();
        } else {
          setTimeout(() => {
            this.calcDropdownInViewportVertical();
          }, 100);
        }
      }
    },
    newValue(value) {
      this.$emit("update:modelValue", value);
      const currentValue = this.getValue(this.selected);
      if (currentValue && currentValue !== value) {
        this.setSelected(null, false);
      }
      if (this.hasFocus && (!this.openOnFocus || value)) {
        this.isActive = !!value;
      }
    },
    data() {
      if (this.keepFirst) {
        this.$nextTick(() => {
          if (this.isActive) {
            this.selectFirstOption(this.computedData);
          } else {
            this.setHovered(null);
          }
        });
      }
    },
    debounceTyping: {
      handler(value) {
        this.debouncedEmitTyping = debounce(this.emitTyping, value);
      },
      immediate: true
    }
  },
  methods: {
    itemOptionClasses(option) {
      return [
        ...this.itemClasses,
        { [this.computedClass("itemHoverClass", "o-acp__item--hover")]: option === this.hovered }
      ];
    },
    setHovered(option) {
      if (option === void 0)
        return;
      this.hovered = option;
    },
    setSelected(option, closeDropdown = true, event = void 0) {
      if (option === void 0)
        return;
      this.selected = option;
      this.$emit("select", this.selected, event);
      if (this.selected !== null) {
        if (this.clearOnSelect) {
          const input = this.$refs.input.$refs.input;
          input.value = "";
        } else {
          this.newValue = this.getValue(this.selected);
        }
        this.setHovered(null);
      }
      closeDropdown && this.$nextTick(() => {
        this.isActive = false;
      });
      this.checkValidity();
    },
    selectFirstOption(computedData) {
      this.$nextTick(() => {
        const nonEmptyElements = computedData.filter((element) => element.items && element.items.length);
        if (nonEmptyElements.length) {
          const option = nonEmptyElements[0].items[0];
          this.setHovered(option);
        } else {
          this.setHovered(null);
        }
      });
    },
    keydown(event) {
      const { key } = event;
      if (key === "Enter")
        event.preventDefault();
      if (key === "Escape" || key === "Tab") {
        this.isActive = false;
      }
      if (this.confirmKeys.indexOf(key) >= 0) {
        if (key === ",")
          event.preventDefault();
        const closeDropdown = !this.keepOpen || key === "Tab";
        if (this.hovered === null) {
          this.checkIfHeaderOrFooterSelected(event, false, closeDropdown);
          return;
        }
        this.setSelected(this.hovered, closeDropdown, event);
      }
    },
    checkIfHeaderOrFooterSelected(event, triggeredByclick, closeDropdown = true) {
      if (this.selectableHeader && (this.headerHovered || triggeredByclick)) {
        this.$emit("select-header", event);
        this.headerHovered = false;
        if (triggeredByclick)
          this.setHovered(null);
        if (closeDropdown)
          this.isActive = false;
      }
      if (this.selectableFooter && (this.footerHovered || triggeredByclick)) {
        this.$emit("select-footer", event);
        this.footerHovered = false;
        if (triggeredByclick)
          this.setHovered(null);
        if (closeDropdown)
          this.isActive = false;
      }
    },
    clickedOutside(event) {
      if (!this.hasFocus && this.whiteList.indexOf(event.target) < 0) {
        if (this.keepFirst && this.hovered && this.selectOnClickOutside) {
          this.setSelected(this.hovered, true);
        } else {
          this.isActive = false;
        }
      }
    },
    getValue(option) {
      if (option === null)
        return;
      if (typeof this.customFormatter !== "undefined") {
        return this.customFormatter(option);
      }
      return typeof option === "object" ? getValueByPath(option, this.field) : option;
    },
    checkIfReachedTheEndOfScroll() {
      const list = this.$refs.dropdown;
      if (list.clientHeight !== list.scrollHeight && list.scrollTop + list.clientHeight >= list.scrollHeight) {
        this.$emit("infinite-scroll");
      }
    },
    calcDropdownInViewportVertical() {
      this.$nextTick(() => {
        if (this.$refs.dropdown === void 0)
          return;
        const rect = this.$refs.dropdown.getBoundingClientRect();
        this.isListInViewportVertically = rect.top >= 0 && rect.bottom <= (window.innerHeight || document.documentElement.clientHeight);
        if (this.appendToBody) {
          this.updateAppendToBody();
        }
      });
    },
    keyArrows(direction) {
      const sum = direction === "down" ? 1 : -1;
      if (this.isActive) {
        const data22 = this.computedData.map((d) => d.items).reduce((a, b) => [...a, ...b], []);
        if (this.$slots.header && this.selectableHeader) {
          data22.unshift(void 0);
        }
        if (this.$slots.footer && this.selectableFooter) {
          data22.push(void 0);
        }
        let index2;
        if (this.headerHovered) {
          index2 = 0 + sum;
        } else if (this.footerHovered) {
          index2 = data22.length - 1 + sum;
        } else {
          index2 = data22.indexOf(this.hovered) + sum;
        }
        index2 = index2 > data22.length - 1 ? data22.length - 1 : index2;
        index2 = index2 < 0 ? 0 : index2;
        this.footerHovered = false;
        this.headerHovered = false;
        this.setHovered(data22[index2] !== void 0 ? data22[index2] : null);
        if (this.$slots.header && this.selectableFooter && index2 === data22.length - 1) {
          this.footerHovered = true;
        }
        if (this.$slots.footer && this.selectableHeader && index2 === 0) {
          this.headerHovered = true;
        }
        const list = this.$refs.dropdown;
        let items2 = this.$refs.items || [];
        if (this.$slots.header && this.selectableHeader) {
          items2 = [this.$refs.header, ...items2];
        }
        if (this.$slots.footer && this.selectableFooter) {
          items2 = [...items2, this.$refs.footer];
        }
        const element = items2[index2];
        if (!element)
          return;
        const visMin = list.scrollTop;
        const visMax = list.scrollTop + list.clientHeight - element.clientHeight;
        if (element.offsetTop < visMin) {
          list.scrollTop = element.offsetTop;
        } else if (element.offsetTop >= visMax) {
          list.scrollTop = element.offsetTop - list.clientHeight + element.clientHeight;
        }
      } else {
        this.isActive = true;
      }
    },
    focused(event) {
      if (this.getValue(this.selected) === this.newValue) {
        this.$el.querySelector("input").select();
      }
      if (this.openOnFocus) {
        this.isActive = true;
        if (this.keepFirst) {
          this.selectFirstOption(this.computedData);
        }
      }
      this.hasFocus = true;
      this.$emit("focus", event);
    },
    onBlur(event) {
      this.hasFocus = false;
      this.$emit("blur", event);
    },
    onInput() {
      const currentValue = this.getValue(this.selected);
      if (currentValue && currentValue === this.newValue)
        return;
      if (this.debounceTyping) {
        this.debouncedEmitTyping();
      } else {
        this.emitTyping();
      }
    },
    emitTyping() {
      this.$emit("typing", this.newValue);
      this.checkValidity();
    },
    rightIconClick(event) {
      if (this.clearable) {
        this.newValue = "";
        this.setSelected(null, false);
        if (this.openOnFocus) {
          this.$refs.input.$el.focus();
        }
      } else {
        this.$emit("icon-right-click", event);
      }
    },
    checkValidity() {
      if (this.useHtml5Validation) {
        this.$nextTick(() => {
          this.checkHtml5Validity();
        });
      }
    },
    setItemRef(el) {
      if (el) {
        this.itemRefs.push(el);
      }
    },
    updateAppendToBody() {
      const dropdownMenu = this.$refs.dropdown;
      const trigger = this.$refs.input.$el;
      if (dropdownMenu && trigger) {
        const root = this.$data.bodyEl;
        root.classList.forEach((item) => root.classList.remove(item));
        this.rootClasses.forEach((item) => {
          if (item) {
            if (typeof item === "object") {
              Object.keys(item).filter((key) => item[key]).forEach((key) => root.classList.add(key));
            } else {
              root.classList.add(item);
            }
          }
        });
        const rect = trigger.getBoundingClientRect();
        let top = rect.top + window.scrollY;
        const left = rect.left + window.scrollX;
        if (this.newDropdownPosition !== "top") {
          top += trigger.clientHeight;
        } else {
          top -= dropdownMenu.clientHeight;
        }
        dropdownMenu.style.position = "absolute";
        dropdownMenu.style.top = `${top}px`;
        dropdownMenu.style.left = `${left}px`;
        dropdownMenu.style.width = `${trigger.clientWidth}px`;
        dropdownMenu.style.maxWidth = `${trigger.clientWidth}px`;
        dropdownMenu.style.zIndex = "9999";
      }
    }
  },
  created() {
    if (typeof window !== "undefined") {
      document.addEventListener("click", this.clickedOutside);
      if (this.menuPosition === "auto")
        window.addEventListener("resize", this.calcDropdownInViewportVertical);
    }
  },
  mounted() {
    const list = this.$refs.dropdown;
    if (this.checkInfiniteScroll && list) {
      list.addEventListener("scroll", this.checkIfReachedTheEndOfScroll);
    }
    if (this.appendToBody) {
      this.$data.bodyEl = createAbsoluteElement(list);
      this.updateAppendToBody();
    }
  },
  beforeUpdate() {
    this.width = this.$refs.input ? this.$refs.input.$el.clientWidth : void 0;
    this.itemRefs = [];
  },
  beforeUnmount() {
    if (typeof window !== "undefined") {
      document.removeEventListener("click", this.clickedOutside);
      if (this.menuPosition === "auto")
        window.removeEventListener("resize", this.calcDropdownInViewportVertical);
    }
    if (this.checkInfiniteScroll && this.$refs.dropdown) {
      const list = this.$refs.dropdown;
      list.removeEventListener("scroll", this.checkIfReachedTheEndOfScroll);
    }
    if (this.appendToBody) {
      removeElement(this.$data.bodyEl);
    }
  }
});
const _hoisted_1$8 = {
  key: 1
};
const _hoisted_2$4 = {
  key: 1
};
function render$m(_ctx, _cache, $props, $setup, $data, $options) {
  const _component_o_input = vue.resolveComponent("o-input");
  return vue.openBlock(), vue.createBlock("div", {
    class: _ctx.rootClasses
  }, [vue.createVNode(_component_o_input, vue.mergeProps(_ctx.inputBind, {
    modelValue: _ctx.newValue,
    "onUpdate:modelValue": [_cache[1] || (_cache[1] = ($event) => _ctx.newValue = $event), _ctx.onInput],
    ref: "input",
    type: _ctx.type,
    size: _ctx.size,
    rounded: _ctx.rounded,
    icon: _ctx.icon,
    "icon-right": _ctx.newIconRight,
    "icon-right-clickable": _ctx.newIconRightClickable,
    "icon-pack": _ctx.iconPack,
    maxlength: _ctx.maxlength,
    autocomplete: _ctx.newAutocomplete,
    "use-html5-validation": false,
    "aria-autocomplete": _ctx.ariaAutocomplete,
    expanded: _ctx.expanded,
    onFocus: _ctx.focused,
    onBlur: _ctx.onBlur,
    onKeydown: [_ctx.keydown, _cache[2] || (_cache[2] = vue.withKeys(vue.withModifiers(($event) => _ctx.keyArrows("up"), ["prevent"]), ["up"])), _cache[3] || (_cache[3] = vue.withKeys(vue.withModifiers(($event) => _ctx.keyArrows("down"), ["prevent"]), ["down"]))],
    "onIcon-right-click": _ctx.rightIconClick,
    "onIcon-click": _cache[4] || (_cache[4] = (event) => _ctx.$emit("icon-click", event))
  }), null, 16, ["modelValue", "type", "size", "rounded", "icon", "icon-right", "icon-right-clickable", "icon-pack", "maxlength", "autocomplete", "aria-autocomplete", "expanded", "onUpdate:modelValue", "onFocus", "onBlur", "onKeydown", "onIcon-right-click"]), vue.createVNode(vue.Transition, {
    name: _ctx.animation
  }, {
    default: vue.withCtx(() => [vue.withDirectives(vue.createVNode("div", {
      class: _ctx.menuClasses,
      style: _ctx.menuStyle,
      ref: "dropdown"
    }, [_ctx.$slots.header ? (vue.openBlock(), vue.createBlock("div", {
      key: 0,
      ref: "header",
      role: "button",
      tabindex: "0",
      onClick: _cache[5] || (_cache[5] = ($event) => _ctx.checkIfHeaderOrFooterSelected($event, true)),
      class: _ctx.itemHeaderClasses
    }, [vue.renderSlot(_ctx.$slots, "header")], 2)) : vue.createCommentVNode("v-if", true), (vue.openBlock(true), vue.createBlock(vue.Fragment, null, vue.renderList(_ctx.computedData, (element, groupindex) => {
      return vue.openBlock(), vue.createBlock(vue.Fragment, null, [element.group ? (vue.openBlock(), vue.createBlock("div", {
        key: groupindex + "group",
        class: _ctx.itemGroupClasses
      }, [_ctx.$slots.group ? vue.renderSlot(_ctx.$slots, "group", {
        key: 0,
        group: element.group,
        index: groupindex
      }) : (vue.openBlock(), vue.createBlock("span", _hoisted_1$8, vue.toDisplayString(element.group), 1))], 2)) : vue.createCommentVNode("v-if", true), (vue.openBlock(true), vue.createBlock(vue.Fragment, null, vue.renderList(element.items, (option, index2) => {
        return vue.openBlock(), vue.createBlock("div", {
          key: groupindex + ":" + index2,
          class: _ctx.itemOptionClasses(option),
          onClick: vue.withModifiers(($event) => _ctx.setSelected(option, !_ctx.keepOpen, $event), ["stop"]),
          ref: _ctx.setItemRef
        }, [_ctx.$slots.default ? vue.renderSlot(_ctx.$slots, "default", {
          key: 0,
          option,
          index: index2
        }) : (vue.openBlock(), vue.createBlock("span", _hoisted_2$4, vue.toDisplayString(_ctx.getValue(option, true)), 1))], 10, ["onClick"]);
      }), 128))], 64);
    }), 256)), _ctx.isEmpty && _ctx.$slots.empty ? (vue.openBlock(), vue.createBlock("div", {
      key: 1,
      class: _ctx.itemEmptyClasses
    }, [vue.renderSlot(_ctx.$slots, "empty")], 2)) : vue.createCommentVNode("v-if", true), _ctx.$slots.footer ? (vue.openBlock(), vue.createBlock("div", {
      key: 2,
      ref: "footer",
      role: "button",
      tabindex: "0",
      onClick: _cache[6] || (_cache[6] = ($event) => _ctx.checkIfHeaderOrFooterSelected($event, true)),
      class: _ctx.itemFooterClasses
    }, [vue.renderSlot(_ctx.$slots, "footer")], 2)) : vue.createCommentVNode("v-if", true)], 6), [[vue.vShow, _ctx.isActive && (!_ctx.isEmpty || _ctx.$slots.empty || _ctx.$slots.header)]])]),
    _: 1
  }, 8, ["name"])], 2);
}
script$p.render = render$m;
script$p.__file = "src/components/autocomplete/Autocomplete.vue";
var index$D = {
  install(app) {
    registerComponent(app, script$p);
  }
};
var index$E = index$D;
var script$o = vue.defineComponent({
  name: "OButton",
  components: {
    [script$r.name]: script$r
  },
  configField: "button",
  mixins: [BaseComponentMixin],
  inheritAttrs: false,
  props: {
    variant: [String, Object],
    size: String,
    label: String,
    iconPack: String,
    iconLeft: String,
    iconRight: String,
    rounded: {
      type: Boolean,
      default: () => {
        return getValueByPath(getOptions$1(), "button.rounded", false);
      }
    },
    outlined: Boolean,
    expanded: Boolean,
    inverted: Boolean,
    nativeType: {
      type: String,
      default: "button",
      validator: (value) => {
        return [
          "button",
          "submit",
          "reset"
        ].indexOf(value) >= 0;
      }
    },
    tag: {
      type: String,
      default: "button"
    },
    disabled: Boolean,
    iconBoth: Boolean,
    elementsWrapperClass: [String, Function, Array],
    rootClass: [String, Function, Array],
    outlinedClass: [String, Function, Array],
    invertedClass: [String, Function, Array],
    expandedClass: [String, Function, Array],
    roundedClass: [String, Function, Array],
    disabledClass: [String, Function, Array],
    iconClass: [String, Function, Array],
    iconLeftClass: [String, Function, Array],
    iconRightClass: [String, Function, Array],
    sizeClass: [String, Function, Array],
    variantClass: [String, Function, Array]
  },
  computed: {
    rootClasses() {
      return [
        this.computedClass("rootClass", "o-btn"),
        { [this.computedClass("sizeClass", "o-btn--", this.size)]: this.size },
        { [this.computedClass("variantClass", "o-btn--", this.variant)]: this.variant },
        { [this.computedClass("outlinedClass", "o-btn--outlined")]: this.outlined && !this.variant },
        { [this.computedClass("invertedClass", "o-btn--inverted")]: this.inverted && !this.variant },
        { [this.computedClass("outlinedClass", "o-btn--outlined-", this.variant)]: this.outlined && this.variant },
        { [this.computedClass("invertedClass", "o-btn--inverted-", this.variant)]: this.inverted && this.variant },
        { [this.computedClass("expandedClass", "o-btn--expanded")]: this.expanded },
        { [this.computedClass("roundedClass", "o-btn--rounded")]: this.rounded },
        { [this.computedClass("disabledClass", "o-btn--disabled")]: this.disabled }
      ];
    },
    iconClasses() {
      return [
        this.computedClass("iconClass", "o-btn__icon")
      ];
    },
    iconLeftClasses() {
      return [
        this.computedClass("iconClass", "o-btn__icon"),
        this.computedClass("iconLeftClass", "o-btn__icon-left")
      ];
    },
    iconRightClasses() {
      return [
        this.computedClass("iconClass", "o-btn__icon"),
        this.computedClass("iconRightClass", "o-btn__icon-right")
      ];
    },
    elementsWrapperClasses() {
      return [
        this.computedClass("elementsWrapperClass", "o-btn__wrapper")
      ];
    },
    computedTag() {
      if (typeof this.disabled !== "undefined" && this.disabled !== false) {
        return "button";
      }
      return this.tag;
    },
    computedNativeType() {
      if (this.tag === "button" || this.tag === "input") {
        return this.nativeType;
      }
      return null;
    },
    computedDisabled() {
      if (this.disabled)
        return true;
      return null;
    }
  }
});
const _hoisted_1$7 = {
  key: 1
};
const _hoisted_2$3 = {
  key: 2
};
function render$l(_ctx, _cache, $props, $setup, $data, $options) {
  const _component_o_icon = vue.resolveComponent("o-icon");
  return vue.openBlock(), vue.createBlock(vue.resolveDynamicComponent(_ctx.computedTag), vue.mergeProps(_ctx.$attrs, {
    disabled: _ctx.computedDisabled,
    type: _ctx.computedNativeType,
    class: _ctx.rootClasses
  }), {
    default: vue.withCtx(() => [vue.createVNode("span", {
      class: _ctx.elementsWrapperClasses
    }, [_ctx.iconLeft ? vue.createVNode(_component_o_icon, {
      key: 0,
      pack: _ctx.iconPack,
      icon: _ctx.iconLeft,
      size: _ctx.size,
      both: _ctx.iconBoth,
      class: _ctx.iconLeftClasses
    }, null, 8, ["pack", "icon", "size", "both", "class"]) : vue.createCommentVNode("v-if", true), _ctx.label ? (vue.openBlock(), vue.createBlock("span", _hoisted_1$7, vue.toDisplayString(_ctx.label), 1)) : _ctx.$slots.default ? (vue.openBlock(), vue.createBlock("span", _hoisted_2$3, [vue.renderSlot(_ctx.$slots, "default")])) : vue.createCommentVNode("v-if", true), _ctx.iconRight ? vue.createVNode(_component_o_icon, {
      key: 3,
      pack: _ctx.iconPack,
      icon: _ctx.iconRight,
      size: _ctx.size,
      both: _ctx.iconBoth,
      class: _ctx.iconRightClasses
    }, null, 8, ["pack", "icon", "size", "both", "class"]) : vue.createCommentVNode("v-if", true)], 2)]),
    _: 1
  }, 16, ["disabled", "type", "class"]);
}
script$o.render = render$l;
script$o.__file = "src/components/button/Button.vue";
var index$C = {
  install(app) {
    registerComponent(app, script$o);
  }
};
var index$1$1 = index$C;
var CheckRadioMixin = vue.defineComponent({
  emits: ["update:modelValue"],
  props: {
    modelValue: [String, Number, Boolean, Array],
    nativeValue: [String, Number, Boolean, Array],
    variant: String,
    disabled: Boolean,
    required: Boolean,
    name: String,
    size: String
  },
  data() {
    return {
      newValue: this.modelValue
    };
  },
  computed: {
    computedValue: {
      get() {
        return this.newValue;
      },
      set(value) {
        this.newValue = value;
        this.$emit("update:modelValue", this.newValue);
      }
    }
  },
  watch: {
    modelValue(value) {
      this.newValue = value;
    }
  },
  methods: {
    focus() {
      this.$refs.input.focus();
    }
  }
});
var script$n = vue.defineComponent({
  name: "OCheckbox",
  components: {
    [script$r.name]: script$r
  },
  mixins: [BaseComponentMixin, CheckRadioMixin],
  configField: "checkbox",
  emits: [
    "input"
  ],
  props: {
    indeterminate: {
      type: Boolean,
      default: false
    },
    trueValue: {
      type: [String, Number, Boolean],
      default: true
    },
    falseValue: {
      type: [String, Number, Boolean],
      default: false
    },
    iconPack: {
      type: String,
      default: () => {
        return getValueByPath(getOptions$1(), "checkbox.iconPack", void 0);
      }
    },
    iconCheck: {
      type: String,
      default: () => {
        return getValueByPath(getOptions$1(), "checkbox.iconCheck", void 0);
      }
    },
    ariaLabelledby: String,
    autocomplete: String,
    rootClass: [String, Function, Array],
    disabledClass: [String, Function, Array],
    checkClass: [String, Function, Array],
    checkCheckedClass: [String, Function, Array],
    checkIndeterminateClass: [String, Function, Array],
    labelClass: [String, Function, Array],
    sizeClass: [String, Function, Array],
    variantClass: [String, Function, Array],
    iconCheckClass: [String, Function, Array],
    iconCheckCheckedClass: [String, Function, Array]
  },
  watch: {
    indeterminate: {
      handler(val) {
        this.isIndeterminate = val;
      },
      immediate: true
    }
  },
  computed: {
    isChecked() {
      return this.computedValue === this.trueValue || Array.isArray(this.computedValue) && this.computedValue.indexOf(this.nativeValue) !== -1;
    },
    rootClasses() {
      return [
        this.computedClass("rootClass", "o-chk"),
        { [this.computedClass("checkedClass", "o-chk--checked")]: this.isChecked },
        { [this.computedClass("sizeClass", "o-chk--", this.size)]: this.size },
        { [this.computedClass("disabledClass", "o-chk--disabled")]: this.disabled },
        { [this.computedClass("variantClass", "o-chk--", this.variant)]: this.variant }
      ];
    },
    checkClasses() {
      return [
        this.computedClass("checkClass", "o-chk__check"),
        { [this.computedClass("checkCheckedClass", "o-chk__check--checked")]: this.isChecked },
        { [this.computedClass("checkIndeterminateClass", "o-chk__check--indeterminate")]: this.isIndeterminate }
      ];
    },
    labelClasses() {
      return [
        this.computedClass("labelClass", "o-chk__label")
      ];
    },
    iconCheckClasses() {
      return [
        this.computedClass("iconCheckClass", "o-chk__icon"),
        { [this.computedClass("iconCheckCheckedClass", "o-chk__icon--checked")]: this.isChecked }
      ];
    }
  }
});
function render$k(_ctx, _cache, $props, $setup, $data, $options) {
  const _component_o_icon = vue.resolveComponent("o-icon");
  return vue.openBlock(), vue.createBlock("label", {
    class: _ctx.rootClasses,
    ref: "label",
    onClick: _cache[3] || (_cache[3] = vue.withModifiers((...args) => _ctx.focus(...args), ["stop"])),
    onKeydown: _cache[4] || (_cache[4] = vue.withKeys(vue.withModifiers(($event) => _ctx.$refs.label.click(), ["prevent"]), ["enter"]))
  }, [_ctx.iconCheck ? vue.createVNode(_component_o_icon, {
    key: 0,
    icon: _ctx.iconCheck,
    pack: _ctx.iconPack,
    size: _ctx.size,
    class: _ctx.iconCheckClasses
  }, null, 8, ["icon", "pack", "size", "class"]) : vue.createCommentVNode("v-if", true), vue.withDirectives(vue.createVNode("input", {
    "onUpdate:modelValue": _cache[1] || (_cache[1] = ($event) => _ctx.computedValue = $event),
    indeterminate: _ctx.indeterminate,
    type: "checkbox",
    ref: "input",
    onClick: _cache[2] || (_cache[2] = vue.withModifiers(() => {
    }, ["stop"])),
    class: _ctx.checkClasses,
    disabled: _ctx.disabled,
    required: _ctx.required,
    name: _ctx.name,
    autocomplete: _ctx.autocomplete,
    value: _ctx.nativeValue,
    "true-value": _ctx.trueValue,
    "false-value": _ctx.falseValue,
    "aria-labelledby": _ctx.ariaLabelledby
  }, null, 10, ["indeterminate", "disabled", "required", "name", "autocomplete", "value", "true-value", "false-value", "aria-labelledby"]), [[vue.vModelCheckbox, _ctx.computedValue]]), vue.createVNode("span", {
    id: _ctx.ariaLabelledby,
    class: _ctx.labelClasses
  }, [vue.renderSlot(_ctx.$slots, "default")], 10, ["id"])], 34);
}
script$n.render = render$k;
script$n.__file = "src/components/checkbox/Checkbox.vue";
var index$B = {
  install(app) {
    registerComponent(app, script$n);
  }
};
var index$2$1 = index$B;
var script$m = vue.defineComponent({
  name: "OCollapse",
  mixins: [BaseComponentMixin],
  configField: "collapse",
  emits: ["update:open", "open", "close"],
  props: {
    open: {
      type: Boolean,
      default: true
    },
    animation: {
      type: String,
      default: () => {
        return getValueByPath(getOptions$1(), "collapse.animation", "fade");
      }
    },
    ariaId: {
      type: String,
      default: ""
    },
    position: {
      type: String,
      default: "top",
      validator: (value) => {
        return [
          "top",
          "bottom"
        ].indexOf(value) > -1;
      }
    },
    rootClass: [String, Function, Array],
    triggerClass: [String, Function, Array],
    contentClass: [String, Function, Array]
  },
  data() {
    return {
      isOpen: this.open
    };
  },
  watch: {
    open(value) {
      this.isOpen = value;
    }
  },
  methods: {
    toggle() {
      this.isOpen = !this.isOpen;
      this.$emit("update:open", this.isOpen);
      this.$emit(this.isOpen ? "open" : "close");
    }
  },
  render() {
    const trigger = vue.h("div", {
      class: this.computedClass("triggerClass", "o-clps__trigger"),
      onClick: this.toggle
    }, this.$slots.trigger({ open: this.isOpen }));
    const content2 = vue.h(vue.Transition, { name: this.animation }, () => vue.withDirectives(vue.h("div", {
      class: this.computedClass("contentClass", "o-clps__content"),
      "id": this.ariaId,
      "aria-expanded": this.isOpen
    }, this.$slots.default()), [[vue.vShow, this.isOpen]]));
    return vue.h("div", { class: this.computedClass("rootClass", "o-clps") }, this.position === "top" ? [trigger, content2] : [content2, trigger]);
  }
});
script$m.__file = "src/components/collapse/Collapse.vue";
var index$A = {
  install(app) {
    registerComponent(app, script$m);
  }
};
var index$3$1 = index$A;
var MatchMediaMixin = vue.defineComponent({
  props: {
    mobileBreakpoint: String
  },
  data() {
    return {
      $matchMediaRef: void 0,
      isMatchMedia: void 0
    };
  },
  methods: {
    onMatchMedia(event) {
      this.isMatchMedia = event.matches;
    }
  },
  created() {
    if (typeof window !== "undefined") {
      let width = this.mobileBreakpoint;
      if (!width) {
        const config2 = getOptions$1();
        const defaultWidth = getValueByPath(config2, `mobileBreakpoint`, "1023px");
        width = getValueByPath(config2, `${this.$options.configField}.mobileBreakpoint`, defaultWidth);
      }
      this.$matchMediaRef = window.matchMedia(`(max-width: ${width})`);
      this.isMatchMedia = this.$matchMediaRef.matches;
      this.$matchMediaRef.addListener(this.onMatchMedia, false);
    }
  },
  beforeUnmount() {
    if (typeof window !== "undefined") {
      this.$matchMediaRef.removeListener(this.checkMatchMedia);
    }
  }
});
const findFocusable = (element, programmatic = false) => {
  if (!element) {
    return null;
  }
  if (programmatic) {
    return element.querySelectorAll(`*[tabindex="-1"]`);
  }
  return element.querySelectorAll(`a[href]:not([tabindex="-1"]),
                                     area[href],
                                     input:not([disabled]),
                                     select:not([disabled]),
                                     textarea:not([disabled]),
                                     button:not([disabled]),
                                     iframe,
                                     object,
                                     embed,
                                     *[tabindex]:not([tabindex="-1"]),
                                     *[contenteditable]`);
};
let onKeyDown;
const bind = (el, { value = true }) => {
  if (value) {
    let focusable = findFocusable(el);
    let focusableProg = findFocusable(el, true);
    if (focusable && focusable.length > 0) {
      onKeyDown = (event) => {
        focusable = findFocusable(el);
        focusableProg = findFocusable(el, true);
        const firstFocusable = focusable[0];
        const lastFocusable = focusable[focusable.length - 1];
        if (event.target === firstFocusable && event.shiftKey && event.key === "Tab") {
          event.preventDefault();
          lastFocusable.focus();
        } else if ((event.target === lastFocusable || Array.from(focusableProg).indexOf(event.target) >= 0) && !event.shiftKey && event.key === "Tab") {
          event.preventDefault();
          firstFocusable.focus();
        }
      };
      el.addEventListener("keydown", onKeyDown);
    }
  }
};
const unbind = (el) => {
  el.removeEventListener("keydown", onKeyDown);
};
const directive = {
  beforeMount: bind,
  beforeUnmount: unbind
};
var script$l = vue.defineComponent({
  name: "ODropdown",
  directives: {
    trapFocus: directive
  },
  configField: "dropdown",
  mixins: [BaseComponentMixin, MatchMediaMixin],
  provide() {
    return {
      $dropdown: this
    };
  },
  emits: ["update:modelValue", "active-change", "change"],
  props: {
    modelValue: {
      type: [String, Number, Boolean, Object, Array],
      default: null
    },
    disabled: Boolean,
    inline: Boolean,
    scrollable: Boolean,
    maxHeight: {
      type: [String, Number],
      default: () => {
        return getValueByPath(getOptions$1(), "dropdown.maxHeight", 200);
      }
    },
    position: {
      type: String,
      validator(value) {
        return ["top-right", "top-left", "bottom-left", "bottom-right"].indexOf(value) > -1;
      }
    },
    mobileModal: {
      type: Boolean,
      default: () => {
        return getValueByPath(getOptions$1(), "dropdown.mobileModal", true);
      }
    },
    ariaRole: {
      type: String,
      validator(value) {
        return ["menu", "list", "dialog"].indexOf(value) > -1;
      },
      default: null
    },
    animation: {
      type: String,
      default: () => {
        return getValueByPath(getOptions$1(), "dropdown.animation", "fade");
      }
    },
    multiple: Boolean,
    trapFocus: {
      type: Boolean,
      default: () => {
        return getValueByPath(getOptions$1(), "dropdown.trapFocus", true);
      }
    },
    closeOnClick: {
      type: Boolean,
      default: true
    },
    canClose: {
      type: [Array, Boolean],
      default: true
    },
    expanded: Boolean,
    triggers: {
      type: Array,
      default: () => ["click"]
    },
    appendToBody: Boolean,
    appendToBodyCopyParent: Boolean,
    rootClass: [String, Function, Array],
    triggerClass: [String, Function, Array],
    inlineClass: [String, Function, Array],
    menuMobileOverlayClass: [String, Function, Array],
    menuClass: [String, Function, Array],
    menuPositionClass: [String, Function, Array],
    menuActiveClass: [String, Function, Array],
    mobileClass: [String, Function, Array],
    disabledClass: [String, Function, Array],
    expandedClass: [String, Function, Array]
  },
  data() {
    return {
      selected: this.modelValue,
      isActive: false,
      isHoverable: false,
      bodyEl: void 0
    };
  },
  computed: {
    rootClasses() {
      return [this.computedClass("rootClass", "o-drop"), {
        [this.computedClass("disabledClass", "o-drop--disabled")]: this.disabled
      }, {
        [this.computedClass("expandedClass", "o-drop--expanded")]: this.expanded
      }, {
        [this.computedClass("inlineClass", "o-drop--inline")]: this.inline
      }, {
        [this.computedClass("mobileClass", "o-drop--mobile")]: this.isMobileModal && this.isMatchMedia && !this.hoverable
      }];
    },
    triggerClasses() {
      return [this.computedClass("triggerClass", "o-drop__trigger")];
    },
    menuMobileOverlayClasses() {
      return [this.computedClass("menuMobileOverlayClass", "o-drop__overlay")];
    },
    menuClasses() {
      return [this.computedClass("menuClass", "o-drop__menu"), {
        [this.computedClass("menuPositionClass", "o-drop__menu--", this.position)]: this.position
      }, {
        [this.computedClass("menuActiveClass", "o-drop__menu--active")]: this.isActive || this.inline
      }];
    },
    isMobileModal() {
      return this.mobileModal && !this.inline;
    },
    cancelOptions() {
      return typeof this.canClose === "boolean" ? this.canClose ? ["escape", "outside"] : [] : this.canClose;
    },
    menuStyle() {
      return {
        maxHeight: this.scrollable ? toCssDimension(this.maxHeight) : null,
        overflow: this.scrollable ? "auto" : null
      };
    },
    hoverable() {
      return this.triggers.indexOf("hover") >= 0;
    }
  },
  watch: {
    modelValue(value) {
      this.selected = value;
    },
    isActive(value) {
      this.$emit("active-change", value);
      if (this.appendToBody) {
        this.$nextTick(() => {
          this.updateAppendToBody();
        });
      }
    }
  },
  methods: {
    selectItem(value) {
      if (this.multiple) {
        if (this.selected) {
          if (this.selected.indexOf(value) === -1) {
            this.selected = [...this.selected, value];
          } else {
            this.selected = this.selected.filter((val) => val !== value);
          }
        } else {
          this.selected = [value];
        }
        this.$emit("change", this.selected);
      } else {
        if (this.selected !== value) {
          this.selected = value;
          this.$emit("change", this.selected);
        }
      }
      this.$emit("update:modelValue", this.selected);
      if (!this.multiple) {
        this.isActive = !this.closeOnClick;
        if (this.hoverable && this.closeOnClick) {
          this.isHoverable = false;
        }
      }
    },
    isInWhiteList(el) {
      if (el === this.$refs.dropdownMenu)
        return true;
      if (el === this.$refs.trigger)
        return true;
      if (this.$refs.dropdownMenu !== void 0) {
        const children = this.$refs.dropdownMenu.querySelectorAll("*");
        for (const child of children) {
          if (el === child) {
            return true;
          }
        }
      }
      if (this.$refs.trigger !== void 0) {
        const children = this.$refs.trigger.querySelectorAll("*");
        for (const child of children) {
          if (el === child) {
            return true;
          }
        }
      }
      return false;
    },
    clickedOutside(event) {
      if (this.cancelOptions.indexOf("outside") < 0)
        return;
      if (this.inline)
        return;
      if (!this.isInWhiteList(event.target))
        this.isActive = false;
    },
    keyPress({
      key
    }) {
      if (this.isActive && (key === "Escape" || key === "Esc")) {
        if (this.cancelOptions.indexOf("escape") < 0)
          return;
        this.isActive = false;
      }
    },
    onClick() {
      if (this.triggers.indexOf("click") < 0)
        return;
      this.toggle();
    },
    onContextMenu() {
      if (this.triggers.indexOf("contextmenu") < 0)
        return;
      this.toggle();
    },
    onHover() {
      if (this.triggers.indexOf("hover") < 0)
        return;
      this.isHoverable = true;
    },
    onFocus() {
      if (this.triggers.indexOf("focus") < 0)
        return;
      this.toggle();
    },
    toggle() {
      if (this.disabled)
        return;
      if (!this.isActive) {
        this.$nextTick(() => {
          const value = !this.isActive;
          this.isActive = value;
          setTimeout(() => this.isActive = value);
        });
      } else {
        this.isActive = !this.isActive;
      }
    },
    updateAppendToBody() {
      const dropdownMenu = this.$refs.dropdownMenu;
      const trigger = this.$refs.trigger;
      if (dropdownMenu && trigger) {
        const dropdown = this.$data.bodyEl.children[0];
        dropdown.classList.forEach((item) => dropdown.classList.remove(item));
        this.rootClasses.forEach((item) => {
          if (item) {
            if (typeof item === "object") {
              Object.keys(item).filter((key) => item[key]).forEach((key) => dropdown.classList.add(key));
            } else {
              dropdown.classList.add(item);
            }
          }
        });
        if (this.appendToBodyCopyParent) {
          const parentNode = this.$refs.dropdown.parentNode;
          const parent = this.$data.bodyEl;
          parent.classList.forEach((item) => parent.classList.remove(item));
          parentNode.classList.forEach((item) => parent.classList.add(item));
        }
        const rect = trigger.getBoundingClientRect();
        let top = rect.top + window.scrollY;
        let left = rect.left + window.scrollX;
        if (!this.position || this.position.indexOf("bottom") >= 0) {
          top += trigger.clientHeight;
        } else {
          top -= dropdownMenu.clientHeight;
        }
        if (this.position && this.position.indexOf("left") >= 0) {
          left -= dropdownMenu.clientWidth - trigger.clientWidth;
        }
        dropdownMenu.style.position = "absolute";
        dropdownMenu.style.top = `${top}px`;
        dropdownMenu.style.left = `${left}px`;
        dropdownMenu.style.zIndex = "9999";
      }
    }
  },
  mounted() {
    if (this.appendToBody) {
      this.$data.bodyEl = createAbsoluteElement(this.$refs.dropdownMenu);
      this.updateAppendToBody();
    }
  },
  created() {
    if (typeof window !== "undefined") {
      document.addEventListener("click", this.clickedOutside);
      document.addEventListener("keyup", this.keyPress);
    }
  },
  beforeUnmount() {
    if (typeof window !== "undefined") {
      document.removeEventListener("click", this.clickedOutside);
      document.removeEventListener("keyup", this.keyPress);
    }
    if (this.appendToBody) {
      removeElement(this.$data.bodyEl);
    }
  }
});
function render$j(_ctx, _cache, $props, $setup, $data, $options) {
  const _directive_trap_focus = vue.resolveDirective("trap-focus");
  return vue.openBlock(), vue.createBlock("div", {
    ref: "dropdown",
    class: _ctx.rootClasses
  }, [!_ctx.inline ? (vue.openBlock(), vue.createBlock("div", {
    key: 0,
    role: "button",
    tabindex: _ctx.disabled ? false : 0,
    ref: "trigger",
    class: _ctx.triggerClasses,
    onClick: _cache[1] || (_cache[1] = (...args) => _ctx.onClick(...args)),
    onContextmenu: _cache[2] || (_cache[2] = vue.withModifiers((...args) => _ctx.onContextMenu(...args), ["prevent"])),
    onMouseenter: _cache[3] || (_cache[3] = (...args) => _ctx.onHover(...args)),
    onMouseleave: _cache[4] || (_cache[4] = ($event) => _ctx.isHoverable = false),
    onFocusCapture: _cache[5] || (_cache[5] = (...args) => _ctx.onFocus(...args)),
    "aria-haspopup": "true"
  }, [vue.renderSlot(_ctx.$slots, "trigger", {
    active: _ctx.isActive
  })], 42, ["tabindex"])) : vue.createCommentVNode("v-if", true), vue.createVNode(vue.Transition, {
    name: _ctx.animation
  }, {
    default: vue.withCtx(() => [_ctx.isMobileModal ? vue.withDirectives((vue.openBlock(), vue.createBlock("div", {
      key: 0,
      class: _ctx.menuMobileOverlayClasses,
      "aria-hidden": !_ctx.isActive
    }, null, 10, ["aria-hidden"])), [[vue.vShow, _ctx.isActive]]) : vue.createCommentVNode("v-if", true)]),
    _: 1
  }, 8, ["name"]), vue.createVNode(vue.Transition, {
    name: _ctx.animation
  }, {
    default: vue.withCtx(() => [vue.withDirectives(vue.createVNode("div", {
      ref: "dropdownMenu",
      class: _ctx.menuClasses,
      "aria-hidden": !_ctx.isActive,
      role: _ctx.ariaRole,
      style: _ctx.menuStyle,
      onMouseenter: _cache[6] || (_cache[6] = (...args) => _ctx.onHover(...args)),
      onMouseleave: _cache[7] || (_cache[7] = ($event) => _ctx.isHoverable = false)
    }, [vue.renderSlot(_ctx.$slots, "default")], 46, ["aria-hidden", "role"]), [[vue.vShow, !_ctx.disabled && (_ctx.isActive || _ctx.isHoverable) || _ctx.inline], [_directive_trap_focus, _ctx.trapFocus]])]),
    _: 3
  }, 8, ["name"])], 2);
}
script$l.render = render$j;
script$l.__file = "src/components/dropdown/Dropdown.vue";
var script$1$9 = vue.defineComponent({
  name: "ODropdownItem",
  mixins: [BaseComponentMixin],
  configField: "dropdown",
  inject: ["$dropdown"],
  emits: ["click"],
  props: {
    value: {
      type: [String, Number, Boolean, Object, Array]
    },
    disabled: Boolean,
    clickable: {
      type: Boolean,
      default: true
    },
    tag: {
      type: String,
      default: () => {
        return getValueByPath(getOptions$1(), "dropdown.itemTag", "div");
      }
    },
    tabindex: {
      type: [Number, String],
      default: 0
    },
    ariaRole: {
      type: String,
      default: ""
    },
    itemClass: [String, Function, Array],
    itemActiveClass: [String, Function, Array],
    itemDisabledClass: [String, Function, Array]
  },
  computed: {
    parent() {
      return this.$dropdown;
    },
    rootClasses() {
      return [this.computedClass("itemClass", "o-drop__item"), {
        [this.computedClass("itemDisabledClass", "o-drop__item--disabled")]: this.parent.disabled || this.disabled
      }, {
        [this.computedClass("itemActiveClass", "o-drop__item--active")]: this.isActive
      }];
    },
    ariaRoleItem() {
      return this.ariaRole === "menuitem" || this.ariaRole === "listitem" ? this.ariaRole : null;
    },
    isClickable() {
      return !this.parent.disabled && !this.disabled && this.clickable;
    },
    isActive() {
      if (this.parent.selected === null)
        return false;
      if (this.parent.multiple)
        return this.parent.selected.indexOf(this.value) >= 0;
      return this.value === this.parent.selected;
    }
  },
  methods: {
    selectItem() {
      if (!this.isClickable)
        return;
      this.parent.selectItem(this.value);
      this.$emit("click");
    }
  },
  created() {
    if (!this.parent) {
      throw new Error("You should wrap oDropdownItem on a oDropdown");
    }
  }
});
function render$1$6(_ctx, _cache, $props, $setup, $data, $options) {
  return vue.openBlock(), vue.createBlock(vue.resolveDynamicComponent(_ctx.tag), {
    class: _ctx.rootClasses,
    onClick: _ctx.selectItem,
    role: _ctx.ariaRoleItem,
    tabindex: _ctx.tabindex
  }, {
    default: vue.withCtx(() => [vue.renderSlot(_ctx.$slots, "default")]),
    _: 3
  }, 8, ["class", "onClick", "role", "tabindex"]);
}
script$1$9.render = render$1$6;
script$1$9.__file = "src/components/dropdown/DropdownItem.vue";
var script$k = vue.defineComponent({
  name: "OFieldBody",
  inject: ["$field"],
  configField: "field",
  computed: {
    parent() {
      return this.$field;
    }
  },
  render() {
    let first = true;
    const slot = this.$slots.default();
    const children = slot.length === 1 && Array.isArray(slot[0].children) ? slot[0].children : slot;
    return vue.h("div", { class: this.parent.bodyHorizontalClasses }, children.map((element) => {
      let message;
      if (first) {
        message = this.parent.newMessage;
        first = false;
      }
      return vue.h(vue.resolveComponent("OField"), { variant: this.parent.newVariant, message }, () => [element]);
    }));
  }
});
script$k.__file = "src/components/field/FieldBody.vue";
var script$1$8 = vue.defineComponent({
  name: "OField",
  components: {
    [script$k.name]: script$k
  },
  configField: "field",
  mixins: [BaseComponentMixin, MatchMediaMixin],
  provide() {
    return {
      $field: this
    };
  },
  inject: {
    $field: { from: "$field", default: false }
  },
  props: {
    variant: [String, Object],
    label: String,
    labelFor: String,
    message: String,
    grouped: Boolean,
    groupMultiline: Boolean,
    horizontal: Boolean,
    addons: {
      type: Boolean,
      default: true
    },
    labelSize: String,
    rootClass: [String, Function, Array],
    horizontalClass: [String, Function, Array],
    groupedClass: [String, Function, Array],
    groupMultilineClass: [String, Function, Array],
    labelClass: [String, Function, Array],
    labelSizeClass: [String, Function, Array],
    labelHorizontalClass: [String, Function, Array],
    bodyClass: [String, Function, Array],
    bodyHorizontalClass: [String, Function, Array],
    addonsClass: [String, Function, Array],
    messageClass: [String, Function, Array],
    variantClass: [String, Function, Array],
    mobileClass: [String, Function, Array],
    focusedClass: [String, Function, Array],
    filledClass: [String, Function, Array]
  },
  data() {
    return {
      newVariant: this.variant,
      newMessage: this.message,
      isFocused: false,
      isFilled: false
    };
  },
  computed: {
    rootClasses() {
      return [
        this.computedClass("rootClass", "o-field"),
        { [this.computedClass("horizontalClass", "o-field--horizontal")]: this.horizontal },
        { [this.computedClass("mobileClass", "o-field--mobile")]: this.isMatchMedia },
        { [this.computedClass("focusedClass", "o-field--focused")]: this.isFocused },
        { [this.computedClass("filledClass", "o-field--filled")]: this.isFilled }
      ];
    },
    messageClasses() {
      return [
        this.computedClass("messageClass", "o-field__message"),
        { [this.computedClass("variantClass", "o-field__message-", this.newVariant)]: this.newVariant }
      ];
    },
    labelClasses() {
      return [
        this.computedClass("labelClass", "o-field__label"),
        { [this.computedClass("labelSizeClass", "o-field__label-", this.labelSize)]: this.labelSize }
      ];
    },
    labelHorizontalClasses() {
      return [
        this.computedClass("labelHorizontalClass", "o-field__horizontal-label")
      ];
    },
    bodyClasses() {
      return [
        this.computedClass("bodyClass", "o-field__body")
      ];
    },
    bodyHorizontalClasses() {
      return [
        this.computedClass("bodyHorizontalClass", "o-field__horizontal-body")
      ];
    },
    innerFieldClasses() {
      return [
        { [this.computedClass("groupMultilineClass", "o-field--grouped-multiline")]: this.groupMultiline },
        { [this.computedClass("groupedClass", "o-field--grouped")]: this.grouped },
        { [this.computedClass("addonsClass", "o-field--addons")]: !this.grouped && this.hasAddons() }
      ];
    },
    parent() {
      return this.$field;
    },
    hasLabelSlot() {
      return this.$slots.label;
    },
    hasMessageSlot() {
      return this.$slots.message;
    },
    hasLabel() {
      return this.label || this.hasLabelSlot;
    },
    hasMessage() {
      return (!this.parent || !this.parent.hasInnerField) && this.newMessage || this.hasMessageSlot;
    },
    hasInnerField() {
      return this.grouped || this.groupMultiline || this.hasAddons();
    }
  },
  watch: {
    variant(value) {
      this.newVariant = value;
    },
    message(value) {
      this.newMessage = value;
    },
    newMessage(value) {
      if (this.parent && this.parent.hasInnerField) {
        this.parent.newMessage = value;
      }
    }
  },
  methods: {
    hasAddons() {
      let renderedNode = 0;
      const slot = this.$slots.default();
      if (slot) {
        const children = slot.length === 1 && Array.isArray(slot[0].children) ? slot[0].children : slot;
        renderedNode = children.reduce((i, node) => node ? i + 1 : i, 0);
      }
      return renderedNode > 1 && this.addons && !this.horizontal;
    }
  }
});
function render$i(_ctx, _cache, $props, $setup, $data, $options) {
  const _component_o_field_body = vue.resolveComponent("o-field-body");
  const _component_o_field = vue.resolveComponent("o-field");
  return vue.openBlock(), vue.createBlock("div", {
    class: _ctx.rootClasses
  }, [_ctx.horizontal ? (vue.openBlock(), vue.createBlock("div", {
    key: 0,
    class: _ctx.labelHorizontalClasses
  }, [_ctx.hasLabel ? (vue.openBlock(), vue.createBlock("label", {
    key: 0,
    for: _ctx.labelFor,
    class: _ctx.labelClasses
  }, [_ctx.hasLabelSlot ? vue.renderSlot(_ctx.$slots, "label", {
    key: 0
  }) : (vue.openBlock(), vue.createBlock(vue.Fragment, {
    key: 1
  }, [vue.createTextVNode(vue.toDisplayString(_ctx.label), 1)], 64))], 10, ["for"])) : vue.createCommentVNode("v-if", true)], 2)) : (vue.openBlock(), vue.createBlock(vue.Fragment, {
    key: 1
  }, [_ctx.hasLabel ? (vue.openBlock(), vue.createBlock("label", {
    key: 0,
    for: _ctx.labelFor,
    class: _ctx.labelClasses
  }, [_ctx.hasLabelSlot ? vue.renderSlot(_ctx.$slots, "label", {
    key: 0
  }) : (vue.openBlock(), vue.createBlock(vue.Fragment, {
    key: 1
  }, [vue.createTextVNode(vue.toDisplayString(_ctx.label), 1)], 64))], 10, ["for"])) : vue.createCommentVNode("v-if", true)], 64)), _ctx.horizontal ? vue.createVNode(_component_o_field_body, {
    key: 2
  }, {
    default: vue.withCtx(() => [vue.renderSlot(_ctx.$slots, "default")]),
    _: 3
  }) : _ctx.hasInnerField ? (vue.openBlock(), vue.createBlock("div", {
    key: 3,
    class: _ctx.bodyClasses
  }, [vue.createVNode(_component_o_field, {
    addons: false,
    variant: _ctx.newVariant,
    class: _ctx.innerFieldClasses
  }, {
    default: vue.withCtx(() => [vue.renderSlot(_ctx.$slots, "default")]),
    _: 3
  }, 8, ["variant", "class"])], 2)) : vue.renderSlot(_ctx.$slots, "default", {
    key: 4
  }), _ctx.hasMessage && !_ctx.horizontal ? (vue.openBlock(), vue.createBlock("p", {
    key: 5,
    class: _ctx.messageClasses
  }, [_ctx.hasMessageSlot ? vue.renderSlot(_ctx.$slots, "message", {
    key: 0
  }) : (vue.openBlock(), vue.createBlock(vue.Fragment, {
    key: 1
  }, [vue.createTextVNode(vue.toDisplayString(_ctx.message), 1)], 64))], 2)) : vue.createCommentVNode("v-if", true)], 2);
}
script$1$8.render = render$i;
script$1$8.__file = "src/components/field/Field.vue";
var script$j = vue.defineComponent({
  name: "OSelect",
  components: {
    [script$r.name]: script$r
  },
  mixins: [BaseComponentMixin, FormElementMixin],
  configField: "select",
  inheritAttrs: false,
  provide() {
    return {
      $elementRef: "select"
    };
  },
  emits: ["update:modelValue", "focus", "blur"],
  props: {
    modelValue: {
      type: [String, Number, Boolean, Object, Array],
      default: null
    },
    size: String,
    iconRight: String,
    placeholder: String,
    multiple: Boolean,
    nativeSize: [String, Number],
    rootClass: [String, Function, Array],
    selectClass: [String, Function, Array],
    iconLeftSpaceClass: [String, Function, Array],
    iconRightSpaceClass: [String, Function, Array],
    roundedClass: [String, Function, Array],
    multipleClass: [String, Function, Array],
    expandedClass: [String, Function, Array],
    iconLeftClass: [String, Function, Array],
    iconRightClass: [String, Function, Array],
    sizeClass: [String, Function, Array],
    variantClass: [String, Function, Array],
    placeholderClass: [String, Function, Array],
    arrowClass: [String, Function, Array]
  },
  data() {
    return {
      selected: this.modelValue
    };
  },
  computed: {
    rootClasses() {
      return [
        this.computedClass("rootClass", "o-ctrl-sel"),
        { [this.computedClass("expandedClass", "o-ctrl-sel--expanded")]: this.expanded }
      ];
    },
    selectClasses() {
      return [
        this.computedClass("selectClass", "o-sel"),
        { [this.computedClass("roundedClass", "o-sel--rounded")]: this.rounded },
        { [this.computedClass("multipleClass", "o-sel--multiple")]: this.multiple },
        { [this.computedClass("sizeClass", "o-sel--", this.size)]: this.size },
        { [this.computedClass("variantClass", "o-sel--", this.statusVariant)]: this.statusVariant },
        { [this.computedClass("iconLeftSpaceClass", "o-sel-iconspace-left")]: this.icon },
        { [this.computedClass("iconRightSpaceClass", "o-sel-iconspace-right")]: this.iconRight },
        { [this.computedClass("placeholderClass", "o-sel--placeholder")]: this.placeholderVisible },
        { [this.computedClass("arrowClass", "o-sel-arrow")]: !this.iconRight && !this.multiple }
      ];
    },
    iconLeftClasses() {
      return [
        this.computedClass("iconLeftClass", "o-sel__icon-left")
      ];
    },
    iconRightClasses() {
      return [
        this.computedClass("iconRightClass", "o-sel__icon-right")
      ];
    },
    placeholderVisible() {
      return this.computedValue === null;
    },
    computedValue: {
      get() {
        return this.selected;
      },
      set(value) {
        this.selected = value;
        this.$emit("update:modelValue", value);
        !this.isValid && this.checkHtml5Validity();
      }
    }
  },
  watch: {
    modelValue(value) {
      this.selected = value;
      !this.isValid && this.checkHtml5Validity();
    }
  }
});
const _hoisted_1$6 = {
  key: 0,
  value: null,
  disabled: "",
  hidden: ""
};
function render$h(_ctx, _cache, $props, $setup, $data, $options) {
  const _component_o_icon = vue.resolveComponent("o-icon");
  return vue.openBlock(), vue.createBlock("div", {
    class: _ctx.rootClasses
  }, [vue.withDirectives(vue.createVNode("select", vue.mergeProps(_ctx.$attrs, {
    "onUpdate:modelValue": _cache[1] || (_cache[1] = ($event) => _ctx.computedValue = $event),
    class: _ctx.selectClasses,
    ref: "select",
    multiple: _ctx.multiple,
    size: _ctx.nativeSize,
    onBlur: _cache[2] || (_cache[2] = ($event) => _ctx.$emit("blur", $event) && _ctx.checkHtml5Validity()),
    onFocus: _cache[3] || (_cache[3] = ($event) => _ctx.$emit("focus", $event))
  }), [_ctx.placeholder ? (vue.openBlock(), vue.createBlock(vue.Fragment, {
    key: 0
  }, [_ctx.placeholderVisible ? (vue.openBlock(), vue.createBlock("option", _hoisted_1$6, vue.toDisplayString(_ctx.placeholder), 1)) : vue.createCommentVNode("v-if", true)], 64)) : vue.createCommentVNode("v-if", true), vue.renderSlot(_ctx.$slots, "default")], 16, ["multiple", "size"]), [[vue.vModelSelect, _ctx.computedValue]]), _ctx.icon ? vue.createVNode(_component_o_icon, {
    key: 0,
    class: _ctx.iconLeftClasses,
    icon: _ctx.icon,
    pack: _ctx.iconPack,
    size: _ctx.size
  }, null, 8, ["class", "icon", "pack", "size"]) : vue.createCommentVNode("v-if", true), _ctx.iconRight && !_ctx.multiple ? vue.createVNode(_component_o_icon, {
    key: 1,
    class: _ctx.iconRightClasses,
    icon: _ctx.iconRight,
    pack: _ctx.iconPack,
    size: _ctx.size
  }, null, 8, ["class", "icon", "pack", "size"]) : vue.createCommentVNode("v-if", true)], 2);
}
script$j.render = render$h;
script$j.__file = "src/components/select/Select.vue";
var script$i = vue.defineComponent({
  name: "ODatepickerTableRow",
  mixins: [BaseComponentMixin],
  configField: "datepicker",
  inject: {
    $datepicker: {
      name: "$datepicker",
      default: false
    }
  },
  emits: ["select", "rangeHoverEndDate", "change-focus"],
  props: {
    selectedDate: {
      type: [Date, Array]
    },
    hoveredDateRange: Array,
    day: {
      type: Number
    },
    week: {
      type: Array,
      required: true
    },
    month: {
      type: Number,
      required: true
    },
    showWeekNumber: Boolean,
    minDate: Date,
    maxDate: Date,
    disabled: Boolean,
    unselectableDates: Array,
    unselectableDaysOfWeek: Array,
    selectableDates: Array,
    events: Array,
    indicators: String,
    dateCreator: Function,
    nearbyMonthDays: Boolean,
    nearbySelectableMonthDays: Boolean,
    weekNumberClickable: Boolean,
    range: Boolean,
    multiple: Boolean,
    rulesForFirstWeek: Number,
    firstDayOfWeek: Number,
    tableRowClass: [String, Function, Array],
    tableCellClass: [String, Function, Array],
    tableCellSelectedClass: [String, Function, Array],
    tableCellFirstSelectedClass: [String, Function, Array],
    tableCellWithinSelectedClass: [String, Function, Array],
    tableCellLastSelectedClass: [String, Function, Array],
    tableCellFirstHoveredClass: [String, Function, Array],
    tableCellInvisibleClass: [String, Function, Array],
    tableCellWithinHoveredClass: [String, Function, Array],
    tableCellLastHoveredClass: [String, Function, Array],
    tableCellTodayClass: [String, Function, Array],
    tableCellSelectableClass: [String, Function, Array],
    tableCellUnselectableClass: [String, Function, Array],
    tableCellNearbyClass: [String, Function, Array],
    tableCellEventsClass: [String, Function, Array],
    tableEventClass: [String, Function, Array],
    tableEventIndicatorsClass: [String, Function, Array],
    tableEventsClass: [String, Function, Array],
    tableEventVariantClass: [String, Function, Array]
  },
  computed: {
    tableRowClasses() {
      return [this.computedClass("tableRowClass", "o-dpck__table__row")];
    },
    tableCellClasses() {
      return [this.computedClass("tableCellClass", "o-dpck__table__cell")];
    },
    tableEventsClasses() {
      return [this.computedClass("tableEventsClass", "o-dpck__table__events")];
    },
    hasEvents() {
      return this.events && this.events.length;
    }
  },
  watch: {
    day(day) {
      const refName = `day-${this.month}-${day}`;
      this.$nextTick(() => {
        if (this.$refs[refName] && this.$refs[refName].length > 0) {
          if (this.$refs[refName][0]) {
            this.$refs[refName][0].focus();
          }
        }
      });
    }
  },
  methods: {
    firstWeekOffset(year, dow, doy) {
      const fwd = 7 + dow - doy;
      const firstJanuary = new Date(year, 0, fwd);
      const fwdlw = (7 + firstJanuary.getDay() - dow) % 7;
      return -fwdlw + fwd - 1;
    },
    daysInYear(year) {
      return this.isLeapYear(year) ? 366 : 365;
    },
    isLeapYear(year) {
      return year % 4 === 0 && year % 100 !== 0 || year % 400 === 0;
    },
    getSetDayOfYear(input) {
      return Math.round((input - new Date(input.getFullYear(), 0, 1)) / 864e5) + 1;
    },
    weeksInYear(year, dow, doy) {
      const weekOffset = this.firstWeekOffset(year, dow, doy);
      const weekOffsetNext = this.firstWeekOffset(year + 1, dow, doy);
      return (this.daysInYear(year) - weekOffset + weekOffsetNext) / 7;
    },
    getWeekNumber(mom) {
      const dow = this.firstDayOfWeek;
      const doy = this.rulesForFirstWeek;
      const weekOffset = this.firstWeekOffset(mom.getFullYear(), dow, doy);
      const week = Math.floor((this.getSetDayOfYear(mom) - weekOffset - 1) / 7) + 1;
      let resWeek;
      let resYear;
      if (week < 1) {
        resYear = mom.getFullYear() - 1;
        resWeek = week + this.weeksInYear(resYear, dow, doy);
      } else if (week > this.weeksInYear(mom.getFullYear(), dow, doy)) {
        resWeek = week - this.weeksInYear(mom.getFullYear(), dow, doy);
        resYear = mom.getFullYear() + 1;
      } else {
        resYear = mom.getFullYear();
        resWeek = week;
      }
      return resWeek;
    },
    clickWeekNumber(week) {
      if (this.weekNumberClickable) {
        this.$datepicker.$emit("week-number-click", week);
      }
    },
    selectableDate(day) {
      const validity = [];
      if (this.minDate) {
        validity.push(day >= this.minDate);
      }
      if (this.maxDate) {
        validity.push(day <= this.maxDate);
      }
      if (this.nearbyMonthDays && !this.nearbySelectableMonthDays) {
        validity.push(day.getMonth() === this.month);
      }
      if (this.selectableDates) {
        for (let i = 0; i < this.selectableDates.length; i++) {
          const enabledDate = this.selectableDates[i];
          if (day.getDate() === enabledDate.getDate() && day.getFullYear() === enabledDate.getFullYear() && day.getMonth() === enabledDate.getMonth()) {
            return true;
          } else {
            validity.push(false);
          }
        }
      }
      if (this.unselectableDates) {
        for (let i = 0; i < this.unselectableDates.length; i++) {
          const disabledDate = this.unselectableDates[i];
          validity.push(day.getDate() !== disabledDate.getDate() || day.getFullYear() !== disabledDate.getFullYear() || day.getMonth() !== disabledDate.getMonth());
        }
      }
      if (this.unselectableDaysOfWeek) {
        for (let i = 0; i < this.unselectableDaysOfWeek.length; i++) {
          const dayOfWeek2 = this.unselectableDaysOfWeek[i];
          validity.push(day.getDay() !== dayOfWeek2);
        }
      }
      return validity.indexOf(false) < 0;
    },
    emitChosenDate(day) {
      if (this.disabled)
        return;
      if (this.selectableDate(day)) {
        this.$emit("select", day);
      }
    },
    eventsDateMatch(day) {
      if (!this.events || !this.events.length)
        return false;
      const dayEvents = [];
      for (let i = 0; i < this.events.length; i++) {
        if (this.events[i].date.getDay() === day.getDay()) {
          dayEvents.push(this.events[i]);
        }
      }
      if (!dayEvents.length) {
        return false;
      }
      return dayEvents;
    },
    cellClasses(day) {
      function dateMatch(dateOne, dateTwo, multiple) {
        if (!dateOne || !dateTwo || multiple) {
          return false;
        }
        if (Array.isArray(dateTwo)) {
          return dateTwo.some((date) => dateOne.getDate() === date.getDate() && dateOne.getFullYear() === date.getFullYear() && dateOne.getMonth() === date.getMonth());
        }
        return dateOne.getDate() === dateTwo.getDate() && dateOne.getFullYear() === dateTwo.getFullYear() && dateOne.getMonth() === dateTwo.getMonth();
      }
      function dateWithin(dateOne, dates, multiple) {
        if (!Array.isArray(dates) || multiple) {
          return false;
        }
        return dateOne > dates[0] && dateOne < dates[1];
      }
      return [...this.tableCellClasses, {
        [this.computedClass("tableCellSelectedClass", "o-dpck__table__cell--selected")]: dateMatch(day, this.selectedDate) || dateWithin(day, this.selectedDate, this.multiple)
      }, {
        [this.computedClass("tableCellFirstSelectedClass", "o-dpck__table__cell--first-selected")]: dateMatch(day, Array.isArray(this.selectedDate) && this.selectedDate[0], this.multiple)
      }, {
        [this.computedClass("tableCellWithinSelectedClass", "o-dpck__table__cell--within-selected")]: dateWithin(day, this.selectedDate, this.multiple)
      }, {
        [this.computedClass("tableCellLastSelectedClass", "o-dpck__table__cell--last-selected")]: dateMatch(day, Array.isArray(this.selectedDate) && this.selectedDate[1], this.multiple)
      }, {
        [this.computedClass("tableCellFirstHoveredClass", "o-dpck__table__cell--first-hovered")]: dateMatch(day, Array.isArray(this.hoveredDateRange) && this.hoveredDateRange[0])
      }, {
        [this.computedClass("tableCellWithinHoveredClass", "o-dpck__table__cell--within-hovered")]: dateWithin(day, this.hoveredDateRange)
      }, {
        [this.computedClass("tableCellLastHoveredClass", "o-dpck__table__cell--last-hovered")]: dateMatch(day, Array.isArray(this.hoveredDateRange) && this.hoveredDateRange[1])
      }, {
        [this.computedClass("tableCellTodayClass", "o-dpck__table__cell--today")]: dateMatch(day, this.dateCreator())
      }, {
        [this.computedClass("tableCellSelectableClass", "o-dpck__table__cell--selectable")]: this.selectableDate(day) && !this.disabled
      }, {
        [this.computedClass("tableCellUnselectableClass", "o-dpck__table__cell--unselectable")]: !this.selectableDate(day) || this.disabled
      }, {
        [this.computedClass("tableCellInvisibleClass", "o-dpck__table__cell--invisible")]: !this.nearbyMonthDays && day.getMonth() !== this.month
      }, {
        [this.computedClass("tableCellNearbyClass", "o-dpck__table__cell--nearby")]: this.nearbySelectableMonthDays && day.getMonth() !== this.month
      }, {
        [this.computedClass("tableCellEventsClass", "o-dpck__table__cell--events")]: this.hasEvents
      }, {
        [this.computedClass("tableCellTodayClass", "o-dpck__table__cell--today")]: dateMatch(day, this.dateCreator())
      }];
    },
    eventClasses(event) {
      return [this.computedClass("tableEventClass", "o-dpck__table__event"), {
        [this.computedClass("tableEventVariantClass", "o-dpck__table__event--", event.type)]: event.type
      }, {
        [this.computedClass("tableEventIndicatorsClass", "o-dpck__table__event--", this.indicators)]: this.indicators
      }];
    },
    setRangeHoverEndDate(day) {
      if (this.range) {
        this.$emit("rangeHoverEndDate", day);
      }
    },
    manageKeydown(event, weekDay) {
      const {
        key
      } = event;
      let preventDefault = true;
      switch (key) {
        case "Tab": {
          preventDefault = false;
          break;
        }
        case " ":
        case "Space":
        case "Spacebar":
        case "Enter": {
          this.emitChosenDate(weekDay);
          break;
        }
        case "ArrowLeft":
        case "Left": {
          this.changeFocus(weekDay, -1);
          break;
        }
        case "ArrowRight":
        case "Right": {
          this.changeFocus(weekDay, 1);
          break;
        }
        case "ArrowUp":
        case "Up": {
          this.changeFocus(weekDay, -7);
          break;
        }
        case "ArrowDown":
        case "Down": {
          this.changeFocus(weekDay, 7);
          break;
        }
      }
      if (preventDefault) {
        event.preventDefault();
      }
    },
    changeFocus(day, inc) {
      const nextDay = new Date(day.getTime());
      nextDay.setDate(day.getDate() + inc);
      while ((!this.minDate || nextDay > this.minDate) && (!this.maxDate || nextDay < this.maxDate) && !this.selectableDate(nextDay)) {
        nextDay.setDate(day.getDate() + Math.sign(inc));
      }
      this.setRangeHoverEndDate(nextDay);
      this.$emit("change-focus", nextDay);
    }
  }
});
function render$g(_ctx, _cache, $props, $setup, $data, $options) {
  return vue.openBlock(), vue.createBlock("div", {
    class: _ctx.tableRowClasses
  }, [_ctx.showWeekNumber ? (vue.openBlock(), vue.createBlock("a", {
    key: 0,
    class: _ctx.tableCellClasses,
    style: {
      "cursor: pointer": _ctx.weekNumberClickable
    },
    onClick: _cache[1] || (_cache[1] = vue.withModifiers(($event) => _ctx.clickWeekNumber(_ctx.getWeekNumber(_ctx.week[6])), ["prevent"]))
  }, [vue.createVNode("span", null, vue.toDisplayString(_ctx.getWeekNumber(_ctx.week[6])), 1)], 6)) : vue.createCommentVNode("v-if", true), (vue.openBlock(true), vue.createBlock(vue.Fragment, null, vue.renderList(_ctx.week, (weekDay, index2) => {
    return vue.openBlock(), vue.createBlock(vue.Fragment, {
      key: index2
    }, [_ctx.selectableDate(weekDay) && !_ctx.disabled ? (vue.openBlock(), vue.createBlock("a", {
      key: 0,
      ref: `day-${weekDay.getMonth()}-${weekDay.getDate()}`,
      class: _ctx.cellClasses(weekDay),
      role: "button",
      href: "#",
      disabled: _ctx.disabled,
      onClick: vue.withModifiers(($event) => _ctx.emitChosenDate(weekDay), ["prevent"]),
      onMouseenter: ($event) => _ctx.setRangeHoverEndDate(weekDay),
      onKeydown: ($event) => _ctx.manageKeydown($event, weekDay),
      tabindex: _ctx.day === weekDay.getDate() && _ctx.month === weekDay.getMonth() ? null : -1
    }, [vue.createVNode("span", null, vue.toDisplayString(weekDay.getDate()), 1), _ctx.eventsDateMatch(weekDay) ? (vue.openBlock(), vue.createBlock("div", {
      key: 0,
      class: _ctx.tableEventsClasses
    }, [(vue.openBlock(true), vue.createBlock(vue.Fragment, null, vue.renderList(_ctx.eventsDateMatch(weekDay), (event, index3) => {
      return vue.openBlock(), vue.createBlock("div", {
        class: _ctx.eventClasses(event),
        key: index3
      }, null, 2);
    }), 128))], 2)) : vue.createCommentVNode("v-if", true)], 42, ["disabled", "onClick", "onMouseenter", "onKeydown", "tabindex"])) : (vue.openBlock(), vue.createBlock("div", {
      key: index2,
      class: _ctx.cellClasses(weekDay)
    }, [vue.createVNode("span", null, vue.toDisplayString(weekDay.getDate()), 1)], 2))], 64);
  }), 128))], 2);
}
script$i.render = render$g;
script$i.__file = "src/components/datepicker/DatepickerTableRow.vue";
var script$1$7 = vue.defineComponent({
  name: "ODatepickerTable",
  mixins: [BaseComponentMixin],
  configField: "datepicker",
  components: {
    [script$i.name]: script$i
  },
  emits: ["update:modelValue", "range-start", "range-end", "update:focused"],
  props: {
    modelValue: {
      type: [Date, Array]
    },
    dayNames: Array,
    monthNames: Array,
    firstDayOfWeek: Number,
    events: Array,
    indicators: String,
    minDate: Date,
    maxDate: Date,
    focused: Object,
    disabled: Boolean,
    dateCreator: Function,
    unselectableDates: Array,
    unselectableDaysOfWeek: Array,
    selectableDates: Array,
    nearbyMonthDays: Boolean,
    nearbySelectableMonthDays: Boolean,
    showWeekNumber: Boolean,
    weekNumberClickable: Boolean,
    rulesForFirstWeek: Number,
    range: Boolean,
    multiple: Boolean,
    tableClass: [String, Function, Array],
    tableHeadClass: [String, Function, Array],
    tableHeadCellClass: [String, Function, Array],
    tableBodyClass: [String, Function, Array],
    tableRowClass: [String, Function, Array],
    tableCellClass: [String, Function, Array],
    tableCellSelectedClass: [String, Function, Array],
    tableCellFirstSelectedClass: [String, Function, Array],
    tableCellInvisibleClass: [String, Function, Array],
    tableCellWithinSelectedClass: [String, Function, Array],
    tableCellLastSelectedClass: [String, Function, Array],
    tableCellFirstHoveredClass: [String, Function, Array],
    tableCellWithinHoveredClass: [String, Function, Array],
    tableCellLastHoveredClass: [String, Function, Array],
    tableCellTodayClass: [String, Function, Array],
    tableCellSelectableClass: [String, Function, Array],
    tableCellUnselectableClass: [String, Function, Array],
    tableCellNearbyClass: [String, Function, Array],
    tableCellEventsClass: [String, Function, Array],
    tableEventClass: [String, Function, Array],
    tableEventIndicatorsClass: [String, Function, Array],
    tableEventsClass: [String, Function, Array],
    tableEventVariantClass: [String, Function, Array]
  },
  data() {
    return {
      selectedBeginDate: void 0,
      selectedEndDate: void 0,
      hoveredEndDate: void 0
    };
  },
  computed: {
    tableClasses() {
      return [
        this.computedClass("tableClass", "o-dpck__table")
      ];
    },
    tableHeadClasses() {
      return [
        this.computedClass("tableHeadClass", "o-dpck__table__head")
      ];
    },
    tableHeadCellClasses() {
      return [
        this.computedClass("tableHeadCellClass", "o-dpck__table__head-cell"),
        ...this.tableCellClasses
      ];
    },
    tableBodyClasses() {
      return [
        this.computedClass("tableBodyClass", "o-dpck__table__body")
      ];
    },
    tableCellClasses() {
      return [
        this.computedClass("tableCellClass", "o-dpck__table__cell")
      ];
    },
    multipleSelectedDates: {
      get() {
        return this.multiple && this.value ? this.value : [];
      },
      set(value) {
        this.$emit("update:modelValue", value);
      }
    },
    visibleDayNames() {
      const visibleDayNames = [];
      let index2 = this.firstDayOfWeek;
      while (visibleDayNames.length < this.dayNames.length) {
        const currentDayName = this.dayNames[index2 % this.dayNames.length];
        visibleDayNames.push(currentDayName);
        index2++;
      }
      if (this.showWeekNumber)
        visibleDayNames.unshift("");
      return visibleDayNames;
    },
    eventsInThisMonth() {
      if (!this.events)
        return [];
      const monthEvents = [];
      for (let i = 0; i < this.events.length; i++) {
        let event = this.events[i];
        if (!Object.prototype.hasOwnProperty.call(event, "date")) {
          event = { date: event };
        }
        if (event.date.getMonth() === this.focused.month && event.date.getFullYear() === this.focused.year) {
          monthEvents.push(event);
        }
      }
      return monthEvents;
    },
    weeksInThisMonth() {
      this.validateFocusedDay();
      const month = this.focused.month;
      const year = this.focused.year;
      const weeksInThisMonth = [];
      let startingDay = 1;
      while (weeksInThisMonth.length < 6) {
        const newWeek = this.weekBuilder(startingDay, month, year);
        weeksInThisMonth.push(newWeek);
        startingDay += 7;
      }
      return weeksInThisMonth;
    },
    hoveredDateRange() {
      if (!this.range) {
        return [];
      }
      if (!isNaN(this.selectedEndDate)) {
        return [];
      }
      if (this.hoveredEndDate < this.selectedBeginDate) {
        return [this.hoveredEndDate, this.selectedBeginDate].filter((d) => d !== void 0);
      }
      return [this.selectedBeginDate, this.hoveredEndDate].filter((d) => d !== void 0);
    }
  },
  methods: {
    updateSelectedDate(date) {
      if (!this.range && !this.multiple) {
        this.$emit("update:modelValue", date);
      } else if (this.range) {
        this.handleSelectRangeDate(date);
      } else if (this.multiple) {
        this.handleSelectMultipleDates(date);
      }
    },
    handleSelectRangeDate(date) {
      if (this.selectedBeginDate && this.selectedEndDate) {
        this.selectedBeginDate = date;
        this.selectedEndDate = void 0;
        this.$emit("range-start", date);
      } else if (this.selectedBeginDate && !this.selectedEndDate) {
        if (this.selectedBeginDate > date) {
          this.selectedEndDate = this.selectedBeginDate;
          this.selectedBeginDate = date;
        } else {
          this.selectedEndDate = date;
        }
        this.$emit("range-end", date);
        this.$emit("update:modelValue", [this.selectedBeginDate, this.selectedEndDate]);
      } else {
        this.selectedBeginDate = date;
        this.$emit("range-start", date);
      }
    },
    handleSelectMultipleDates(date) {
      const multipleSelect = this.multipleSelectedDates.filter((selectedDate) => selectedDate.getDate() === date.getDate() && selectedDate.getFullYear() === date.getFullYear() && selectedDate.getMonth() === date.getMonth());
      if (multipleSelect.length) {
        this.multipleSelectedDates = this.multipleSelectedDates.filter((selectedDate) => selectedDate.getDate() !== date.getDate() || selectedDate.getFullYear() !== date.getFullYear() || selectedDate.getMonth() !== date.getMonth());
      } else {
        this.multipleSelectedDates = [...this.multipleSelectedDates, date];
      }
    },
    weekBuilder(startingDate, month, year) {
      const thisMonth = new Date(year, month);
      const thisWeek = [];
      const dayOfWeek2 = new Date(year, month, startingDate).getDay();
      const end = dayOfWeek2 >= this.firstDayOfWeek ? dayOfWeek2 - this.firstDayOfWeek : 7 - this.firstDayOfWeek + dayOfWeek2;
      let daysAgo = 1;
      for (let i = 0; i < end; i++) {
        thisWeek.unshift(new Date(thisMonth.getFullYear(), thisMonth.getMonth(), startingDate - daysAgo));
        daysAgo++;
      }
      thisWeek.push(new Date(year, month, startingDate));
      let daysForward = 1;
      while (thisWeek.length < 7) {
        thisWeek.push(new Date(year, month, startingDate + daysForward));
        daysForward++;
      }
      return thisWeek;
    },
    validateFocusedDay() {
      const focusedDate = new Date(this.focused.year, this.focused.month, this.focused.day);
      if (this.selectableDate(focusedDate))
        return;
      let day = 0;
      const monthDays = new Date(this.focused.year, this.focused.month + 1, 0).getDate();
      let firstFocusable = null;
      while (!firstFocusable && ++day < monthDays) {
        const date = new Date(this.focused.year, this.focused.month, day);
        if (this.selectableDate(date)) {
          firstFocusable = focusedDate;
          const focused = {
            day: date.getDate(),
            month: date.getMonth(),
            year: date.getFullYear()
          };
          this.$emit("update:focused", focused);
        }
      }
    },
    selectableDate(day) {
      const validity = [];
      if (this.minDate) {
        validity.push(day >= this.minDate);
      }
      if (this.maxDate) {
        validity.push(day <= this.maxDate);
      }
      if (this.nearbyMonthDays && !this.nearbySelectableMonthDays) {
        validity.push(day.getMonth() === this.focused.month);
      }
      if (this.selectableDates) {
        for (let i = 0; i < this.selectableDates.length; i++) {
          const enabledDate = this.selectableDates[i];
          if (day.getDate() === enabledDate.getDate() && day.getFullYear() === enabledDate.getFullYear() && day.getMonth() === enabledDate.getMonth()) {
            return true;
          } else {
            validity.push(false);
          }
        }
      }
      if (this.unselectableDates) {
        for (let i = 0; i < this.unselectableDates.length; i++) {
          const disabledDate = this.unselectableDates[i];
          validity.push(day.getDate() !== disabledDate.getDate() || day.getFullYear() !== disabledDate.getFullYear() || day.getMonth() !== disabledDate.getMonth());
        }
      }
      if (this.unselectableDaysOfWeek) {
        for (let i = 0; i < this.unselectableDaysOfWeek.length; i++) {
          const dayOfWeek2 = this.unselectableDaysOfWeek[i];
          validity.push(day.getDay() !== dayOfWeek2);
        }
      }
      return validity.indexOf(false) < 0;
    },
    eventsInThisWeek(week) {
      return this.eventsInThisMonth.filter((event) => {
        const stripped = new Date(Date.parse(event.date));
        stripped.setHours(0, 0, 0, 0);
        const timed = stripped.getTime();
        return week.some((weekDate) => weekDate.getTime() === timed);
      });
    },
    setRangeHoverEndDate(day) {
      this.hoveredEndDate = day;
    },
    changeFocus(day) {
      const focused = {
        day: day.getDate(),
        month: day.getMonth(),
        year: day.getFullYear()
      };
      this.$emit("update:focused", focused);
    }
  }
});
function render$1$5(_ctx, _cache, $props, $setup, $data, $options) {
  const _component_o_datepicker_table_row = vue.resolveComponent("o-datepicker-table-row");
  return vue.openBlock(), vue.createBlock("section", {
    class: _ctx.tableClasses
  }, [vue.createVNode("header", {
    class: _ctx.tableHeadClasses
  }, [(vue.openBlock(true), vue.createBlock(vue.Fragment, null, vue.renderList(_ctx.visibleDayNames, (day, index2) => {
    return vue.openBlock(), vue.createBlock("div", {
      key: index2,
      class: _ctx.tableHeadCellClasses
    }, [vue.createVNode("span", null, vue.toDisplayString(day), 1)], 2);
  }), 128))], 2), vue.createVNode("div", {
    class: _ctx.tableBodyClasses
  }, [(vue.openBlock(true), vue.createBlock(vue.Fragment, null, vue.renderList(_ctx.weeksInThisMonth, (week, index2) => {
    return vue.openBlock(), vue.createBlock(_component_o_datepicker_table_row, {
      key: index2,
      "selected-date": _ctx.modelValue,
      day: _ctx.focused.day,
      week,
      month: _ctx.focused.month,
      "min-date": _ctx.minDate,
      "max-date": _ctx.maxDate,
      disabled: _ctx.disabled,
      "unselectable-dates": _ctx.unselectableDates,
      "unselectable-days-of-week": _ctx.unselectableDaysOfWeek,
      "selectable-dates": _ctx.selectableDates,
      events: _ctx.eventsInThisWeek(week),
      indicators: _ctx.indicators,
      "date-creator": _ctx.dateCreator,
      "nearby-month-days": _ctx.nearbyMonthDays,
      "nearby-selectable-month-days": _ctx.nearbySelectableMonthDays,
      "show-week-number": _ctx.showWeekNumber,
      "week-number-clickable": _ctx.weekNumberClickable,
      "first-day-of-week": _ctx.firstDayOfWeek,
      "rules-for-first-week": _ctx.rulesForFirstWeek,
      range: _ctx.range,
      "hovered-date-range": _ctx.hoveredDateRange,
      multiple: _ctx.multiple,
      "table-row-class": _ctx.tableRowClass,
      "table-cell-class": _ctx.tableCellClass,
      "table-cell-selected-class": _ctx.tableCellSelectedClass,
      "table-cell-first-selected-class": _ctx.tableCellFirstSelectedClass,
      "table-cell-invisible-class": _ctx.tableCellInvisibleClass,
      "table-cell-within-selected-class": _ctx.tableCellWithinSelectedClass,
      "table-cell-last-selected-class": _ctx.tableCellLastSelectedClass,
      "table-cell-first-hovered-class": _ctx.tableCellFirstHoveredClass,
      "table-cell-within-hovered-class": _ctx.tableCellWithinHoveredClass,
      "table-cell-last-hovered-class": _ctx.tableCellLastHoveredClass,
      "table-cell-today-class": _ctx.tableCellTodayClass,
      "table-cell-selectable-class": _ctx.tableCellSelectableClass,
      "table-cell-unselectable-class": _ctx.tableCellUnselectableClass,
      "table-cell-nearby-class": _ctx.tableCellNearbyClass,
      "table-cell-events-class": _ctx.tableCellEventsClass,
      "table-events-class": _ctx.tableEventsClass,
      "table-event-variant-class": _ctx.tableEventVariantClass,
      "table-event-class": _ctx.tableEventClass,
      "table-event-indicators-class": _ctx.tableEventIndicatorsClass,
      onSelect: _ctx.updateSelectedDate,
      onRangeHoverEndDate: _ctx.setRangeHoverEndDate,
      "onChange-focus": _ctx.changeFocus
    }, null, 8, ["selected-date", "day", "week", "month", "min-date", "max-date", "disabled", "unselectable-dates", "unselectable-days-of-week", "selectable-dates", "events", "indicators", "date-creator", "nearby-month-days", "nearby-selectable-month-days", "show-week-number", "week-number-clickable", "first-day-of-week", "rules-for-first-week", "range", "hovered-date-range", "multiple", "table-row-class", "table-cell-class", "table-cell-selected-class", "table-cell-first-selected-class", "table-cell-invisible-class", "table-cell-within-selected-class", "table-cell-last-selected-class", "table-cell-first-hovered-class", "table-cell-within-hovered-class", "table-cell-last-hovered-class", "table-cell-today-class", "table-cell-selectable-class", "table-cell-unselectable-class", "table-cell-nearby-class", "table-cell-events-class", "table-events-class", "table-event-variant-class", "table-event-class", "table-event-indicators-class", "onSelect", "onRangeHoverEndDate", "onChange-focus"]);
  }), 128))], 2)], 2);
}
script$1$7.render = render$1$5;
script$1$7.__file = "src/components/datepicker/DatepickerTable.vue";
const defaultDateFormatter = (date, vm) => {
  const targetDates = Array.isArray(date) ? date : [date];
  const dates = targetDates.map((date2) => {
    const d = new Date(date2.getFullYear(), date2.getMonth(), date2.getDate(), 12);
    return !vm.isTypeMonth ? vm.dtf.format(d) : vm.dtfMonth.format(d);
  });
  return !vm.multiple ? dates.join(" - ") : dates.join(", ");
};
const defaultDateParser = (date, vm) => {
  if (vm.dtf.formatToParts && typeof vm.dtf.formatToParts === "function") {
    const formatRegex = (vm.isTypeMonth ? vm.dtfMonth : vm.dtf).formatToParts(new Date(2e3, 11, 25)).map((part) => {
      if (part.type === "literal") {
        return part.value;
      }
      return `((?!=<${part.type}>)\\d+)`;
    }).join("");
    const dateGroups = matchWithGroups(formatRegex, date);
    if (dateGroups.year && dateGroups.year.length === 4 && dateGroups.month && dateGroups.month <= 12) {
      if (vm.isTypeMonth)
        return new Date(dateGroups.year, dateGroups.month - 1);
      else if (dateGroups.day && dateGroups.day <= 31) {
        return new Date(dateGroups.year, dateGroups.month - 1, dateGroups.day, 12);
      }
    }
  }
  if (!vm.isTypeMonth)
    return new Date(Date.parse(date));
  if (date) {
    const s2 = date.split("/");
    const year = s2[0].length === 4 ? s2[0] : s2[1];
    const month = s2[0].length === 2 ? s2[0] : s2[1];
    if (year && month) {
      return new Date(parseInt(year, 10), parseInt(month, 10) - 1, 1, 0, 0, 0, 0);
    }
  }
  return null;
};
var script$2$3 = vue.defineComponent({
  name: "ODatepicker",
  components: {
    [script$1$7.name]: script$1$7,
    [script$1$8.name]: script$1$8,
    [script$q.name]: script$q,
    [script$j.name]: script$j,
    [script$r.name]: script$r,
    [script$l.name]: script$l,
    [script$1$9.name]: script$1$9
  },
  configField: "datepicker",
  mixins: [BaseComponentMixin, FormElementMixin, MatchMediaMixin],
  inheritAttrs: false,
  provide() {
    return {
      $datepicker: this
    };
  },
  emits: ["update:modelValue", "focus", "blur", "change-month", "change-year", "range-start", "range-end"],
  props: {
    modelValue: {
      type: [Date, Array]
    },
    dayNames: {
      type: Array,
      default: () => {
        return getValueByPath(getOptions$1(), "datepicker.dayNames", void 0);
      }
    },
    monthNames: {
      type: Array,
      default: () => {
        return getValueByPath(getOptions$1(), "datepicker.monthNames", void 0);
      }
    },
    firstDayOfWeek: {
      type: Number,
      default: () => {
        return getValueByPath(getOptions$1(), "datepicker.firstDayOfWeek", 0);
      }
    },
    size: String,
    inline: Boolean,
    minDate: Date,
    maxDate: Date,
    focusedDate: Date,
    placeholder: String,
    editable: Boolean,
    disabled: Boolean,
    unselectableDates: Array,
    unselectableDaysOfWeek: {
      type: Array,
      default: () => {
        return getValueByPath(getOptions$1(), "datepicker.unselectableDaysOfWeek", void 0);
      }
    },
    selectableDates: Array,
    dateFormatter: {
      type: Function,
      default: (date, vm) => {
        const dateFormatter = getValueByPath(getOptions$1(), "datepicker.dateFormatter", void 0);
        if (typeof dateFormatter === "function") {
          return dateFormatter(date);
        } else {
          return defaultDateFormatter(date, vm);
        }
      }
    },
    dateParser: {
      type: Function,
      default: (date, vm) => {
        const dateParser = getValueByPath(getOptions$1(), "datepicker.dateParser", void 0);
        if (typeof dateParser === "function") {
          return dateParser(date);
        } else {
          return defaultDateParser(date, vm);
        }
      }
    },
    dateCreator: {
      type: Function,
      default: () => {
        const dateCreator = getValueByPath(getOptions$1(), "datepicker.dateCreator", void 0);
        if (typeof dateCreator === "function") {
          return dateCreator();
        } else {
          return new Date();
        }
      }
    },
    mobileNative: {
      type: Boolean,
      default: () => {
        return getValueByPath(getOptions$1(), "datepicker.mobileNative", true);
      }
    },
    position: String,
    iconRight: String,
    iconRightClickable: Boolean,
    events: Array,
    indicators: {
      type: String,
      default: "dots"
    },
    openOnFocus: Boolean,
    iconPrev: {
      type: String,
      default: () => {
        return getValueByPath(getOptions$1(), "datepicker.iconPrev", "chevron-left");
      }
    },
    iconNext: {
      type: String,
      default: () => {
        return getValueByPath(getOptions$1(), "datepicker.iconNext", "chevron-right");
      }
    },
    yearsRange: {
      type: Array,
      default: () => {
        return getValueByPath(getOptions$1(), "datepicker.yearsRange", [-100, 10]);
      }
    },
    type: {
      type: String,
      validator: (value) => {
        return [
          "month"
        ].indexOf(value) >= 0;
      }
    },
    nearbyMonthDays: {
      type: Boolean,
      default: () => {
        return getValueByPath(getOptions$1(), "datepicker.nearbyMonthDays", true);
      }
    },
    nearbySelectableMonthDays: {
      type: Boolean,
      default: () => {
        return getValueByPath(getOptions$1(), "datepicker.nearbySelectableMonthDays", false);
      }
    },
    showWeekNumber: {
      type: Boolean,
      default: () => {
        return getValueByPath(getOptions$1(), "datepicker.showWeekNumber", false);
      }
    },
    weekNumberClickable: {
      type: Boolean,
      default: () => {
        return getValueByPath(getOptions$1(), "datepicker.weekNumberClickable", false);
      }
    },
    rulesForFirstWeek: {
      type: Number,
      default: () => 4
    },
    range: {
      type: Boolean,
      default: false
    },
    closeOnClick: {
      type: Boolean,
      default: true
    },
    multiple: {
      type: Boolean,
      default: false
    },
    mobileModal: {
      type: Boolean,
      default: () => {
        return getValueByPath(getOptions$1(), "datepicker.mobileModal", true);
      }
    },
    trapFocus: {
      type: Boolean,
      default: () => {
        return getValueByPath(getOptions$1(), "datepicker.trapFocus", true);
      }
    },
    locale: {
      type: [String, Array],
      default: () => {
        return getValueByPath(getOptions$1(), "locale");
      }
    },
    appendToBody: Boolean,
    ariaNextLabel: String,
    ariaPreviousLabel: String,
    rootClass: [String, Function, Array],
    sizeClass: [String, Function, Array],
    boxClass: [String, Function, Array],
    headerClass: [String, Function, Array],
    headerButtonsClass: [String, Function, Array],
    headerButtonsSizeClass: [String, Function, Array],
    prevBtnClass: [String, Function, Array],
    nextBtnClass: [String, Function, Array],
    listsClass: [String, Function, Array],
    footerClass: [String, Function, Array],
    tableClass: [String, Function, Array],
    tableHeadClass: [String, Function, Array],
    tableHeadCellClass: [String, Function, Array],
    tableBodyClass: [String, Function, Array],
    tableRowClass: [String, Function, Array],
    tableCellClass: [String, Function, Array],
    tableCellSelectedClass: [String, Function, Array],
    tableCellFirstSelectedClass: [String, Function, Array],
    tableCellInvisibleClass: [String, Function, Array],
    tableCellWithinSelectedClass: [String, Function, Array],
    tableCellLastSelectedClass: [String, Function, Array],
    tableCellFirstHoveredClass: [String, Function, Array],
    tableCellWithinHoveredClass: [String, Function, Array],
    tableCellLastHoveredClass: [String, Function, Array],
    tableCellTodayClass: [String, Function, Array],
    tableCellSelectableClass: [String, Function, Array],
    tableCellUnselectableClass: [String, Function, Array],
    tableCellNearbyClass: [String, Function, Array],
    tableCellEventsClass: [String, Function, Array],
    tableEventsClass: [String, Function, Array],
    tableEventVariantClass: [String, Function, Array],
    tableEventClass: [String, Function, Array],
    tableEventIndicatorsClass: [String, Function, Array],
    mobileClass: [String, Function, Array],
    inputClasses: Object,
    dropdownClasses: Object
  },
  data() {
    const focusedDate = (Array.isArray(this.modelValue) ? this.modelValue[0] : this.modelValue) || this.focusedDate || this.dateCreator();
    if (!this.modelValue && this.maxDate && this.maxDate.getFullYear() < focusedDate.getFullYear()) {
      focusedDate.setFullYear(this.maxDate.getFullYear());
    }
    return {
      dateSelected: this.modelValue,
      focusedDateData: {
        day: focusedDate.getDate(),
        month: focusedDate.getMonth(),
        year: focusedDate.getFullYear()
      },
      $elementRef: "input"
    };
  },
  computed: {
    inputBind() {
      return __spreadValues(__spreadValues({}, this.$attrs), this.inputClasses);
    },
    dropdownBind() {
      return __spreadValues({
        "root-class": this.computedClass("dropdownClasses.rootClass", "o-dpck__dropdown")
      }, this.dropdownClasses);
    },
    rootClasses() {
      return [
        this.computedClass("rootClass", "o-dpck"),
        { [this.computedClass("sizeClass", "o-dpck--", this.size)]: this.size },
        { [this.computedClass("mobileClass", "o-dpck--mobile")]: this.isMatchMedia }
      ];
    },
    boxClasses() {
      return [
        this.computedClass("boxClass", "o-dpck__box")
      ];
    },
    headerClasses() {
      return [
        this.computedClass("headerClass", "o-dpck__header")
      ];
    },
    headerButtonsClasses() {
      return [
        this.computedClass("headerButtonsClass", "o-dpck__header__buttons"),
        { [this.computedClass("headerButtonsSizeClass", "o-dpck__header__buttons--", this.size)]: this.size }
      ];
    },
    prevBtnClasses() {
      return [
        this.computedClass("prevBtnClass", "o-dpck__header__previous")
      ];
    },
    nextBtnClasses() {
      return [
        this.computedClass("nextBtnClass", "o-dpck__header__next")
      ];
    },
    listsClasses() {
      return [
        this.computedClass("listsClass", "o-dpck__header__list")
      ];
    },
    footerClasses() {
      return [
        this.computedClass("footerClass", "o-dpck__footer")
      ];
    },
    computedValue: {
      get() {
        return this.dateSelected;
      },
      set(value) {
        this.updateInternalState(value);
        if (!this.multiple)
          this.togglePicker(false);
        this.$emit("update:modelValue", value);
        if (this.useHtml5Validation) {
          this.$nextTick(() => {
            this.checkHtml5Validity();
          });
        }
      }
    },
    formattedValue() {
      return this.formatValue(this.computedValue);
    },
    localeOptions() {
      return new Intl.DateTimeFormat(this.locale, {
        year: "numeric",
        month: "numeric"
      }).resolvedOptions();
    },
    dtf() {
      return new Intl.DateTimeFormat(this.locale);
    },
    dtfMonth() {
      return new Intl.DateTimeFormat(this.locale, {
        year: this.localeOptions.year || "numeric",
        month: this.localeOptions.month || "2-digit"
      });
    },
    newMonthNames() {
      if (Array.isArray(this.monthNames)) {
        return this.monthNames;
      }
      return getMonthNames(this.locale);
    },
    newDayNames() {
      if (Array.isArray(this.dayNames)) {
        return this.dayNames;
      }
      return getWeekdayNames(this.locale);
    },
    listOfMonths() {
      let minMonth = 0;
      let maxMonth = 12;
      if (this.minDate && this.focusedDateData.year === this.minDate.getFullYear()) {
        minMonth = this.minDate.getMonth();
      }
      if (this.maxDate && this.focusedDateData.year === this.maxDate.getFullYear()) {
        maxMonth = this.maxDate.getMonth();
      }
      return this.newMonthNames.map((name, index2) => {
        return {
          name,
          index: index2,
          disabled: index2 < minMonth || index2 > maxMonth
        };
      });
    },
    listOfYears() {
      let latestYear = this.focusedDateData.year + this.yearsRange[1];
      if (this.maxDate && this.maxDate.getFullYear() < latestYear) {
        latestYear = Math.max(this.maxDate.getFullYear(), this.focusedDateData.year);
      }
      let earliestYear = this.focusedDateData.year + this.yearsRange[0];
      if (this.minDate && this.minDate.getFullYear() > earliestYear) {
        earliestYear = Math.min(this.minDate.getFullYear(), this.focusedDateData.year);
      }
      const arrayOfYears = [];
      for (let i = earliestYear; i <= latestYear; i++) {
        arrayOfYears.push(i);
      }
      return arrayOfYears.reverse();
    },
    showPrev() {
      if (!this.minDate)
        return false;
      if (this.isTypeMonth) {
        return this.focusedDateData.year <= this.minDate.getFullYear();
      }
      const dateToCheck = new Date(this.focusedDateData.year, this.focusedDateData.month);
      const date = new Date(this.minDate.getFullYear(), this.minDate.getMonth());
      return dateToCheck <= date;
    },
    showNext() {
      if (!this.maxDate)
        return false;
      if (this.isTypeMonth) {
        return this.focusedDateData.year >= this.maxDate.getFullYear();
      }
      const dateToCheck = new Date(this.focusedDateData.year, this.focusedDateData.month);
      const date = new Date(this.maxDate.getFullYear(), this.maxDate.getMonth());
      return dateToCheck >= date;
    },
    isMobile() {
      return this.mobileNative && isMobile.any();
    },
    isTypeMonth() {
      return this.type === "month";
    },
    ariaRole() {
      return !this.inline ? "dialog" : void 0;
    }
  },
  watch: {
    modelValue(value) {
      this.updateInternalState(value);
      if (!this.multiple)
        this.togglePicker(false);
    },
    focusedDate(value) {
      if (value) {
        this.focusedDateData = {
          day: value.getDate(),
          month: value.getMonth(),
          year: value.getFullYear()
        };
      }
    },
    "focusedDateData.month"(value) {
      this.$emit("change-month", value);
    },
    "focusedDateData.year"(value) {
      this.$emit("change-year", value);
    }
  },
  methods: {
    onChange(value) {
      const date = this.dateParser(value, this);
      if (date && (!isNaN(date) || Array.isArray(date) && date.length === 2 && !isNaN(date[0]) && !isNaN(date[1]))) {
        this.computedValue = date;
      } else {
        this.computedValue = null;
        if (this.$refs.input) {
          this.$refs.input.newValue = this.computedValue;
        }
      }
    },
    formatValue(value) {
      if (Array.isArray(value)) {
        const isArrayWithValidDates = Array.isArray(value) && value.every((v) => !isNaN(v));
        return isArrayWithValidDates ? this.dateFormatter([...value], this) : null;
      }
      return value && !isNaN(value) ? this.dateFormatter(value, this) : null;
    },
    prev() {
      if (this.disabled)
        return;
      if (this.isTypeMonth) {
        this.focusedDateData.year -= 1;
      } else {
        if (this.focusedDateData.month > 0) {
          this.focusedDateData.month -= 1;
        } else {
          this.focusedDateData.month = 11;
          this.focusedDateData.year -= 1;
        }
      }
    },
    next() {
      if (this.disabled)
        return;
      if (this.isTypeMonth) {
        this.focusedDateData.year += 1;
      } else {
        if (this.focusedDateData.month < 11) {
          this.focusedDateData.month += 1;
        } else {
          this.focusedDateData.month = 0;
          this.focusedDateData.year += 1;
        }
      }
    },
    formatNative(value) {
      return this.isTypeMonth ? this.formatYYYYMM(value) : this.formatYYYYMMDD(value);
    },
    formatYYYYMMDD(value) {
      const date = new Date(value);
      if (value && !isNaN(date.getTime())) {
        const year = date.getFullYear();
        const month = date.getMonth() + 1;
        const day = date.getDate();
        return year + "-" + ((month < 10 ? "0" : "") + month) + "-" + ((day < 10 ? "0" : "") + day);
      }
      return "";
    },
    formatYYYYMM(value) {
      const date = new Date(value);
      if (value && !isNaN(date.getTime())) {
        const year = date.getFullYear();
        const month = date.getMonth() + 1;
        return year + "-" + ((month < 10 ? "0" : "") + month);
      }
      return "";
    },
    onChangeNativePicker(event) {
      const date = event.target.value;
      const s2 = date ? date.split("-") : [];
      if (s2.length === 3) {
        const year = parseInt(s2[0], 10);
        const month = parseInt(s2[1]) - 1;
        const day = parseInt(s2[2]);
        this.computedValue = new Date(year, month, day);
      } else {
        this.computedValue = null;
      }
    },
    updateInternalState(value) {
      if (this.dateSelected === value)
        return;
      const currentDate = Array.isArray(value) ? !value.length ? this.dateCreator() : value[value.length - 1] : !value ? this.dateCreator() : value;
      if (Array.isArray(value) && this.dateSelected && value.length > this.dateSelected.length) {
        this.focusedDateData = {
          day: currentDate.getDate(),
          month: currentDate.getMonth(),
          year: currentDate.getFullYear()
        };
      }
      this.dateSelected = value;
    },
    togglePicker(active) {
      if (this.$refs.dropdown) {
        const isActive = typeof active === "boolean" ? active : !this.$refs.dropdown.isActive;
        if (isActive) {
          this.$refs.dropdown.isActive = isActive;
        } else if (this.closeOnClick) {
          this.$refs.dropdown.isActive = isActive;
        }
      }
    },
    handleOnFocus(event) {
      this.onFocus(event);
      if (this.openOnFocus) {
        this.togglePicker(true);
      }
    },
    toggle() {
      if (this.mobileNative && this.isMobile) {
        const input = this.$refs.input.$refs.input;
        input.focus();
        input.click();
        return;
      }
      this.$refs.dropdown.toggle();
    },
    onInputClick(event) {
      if (this.$refs.dropdown.isActive) {
        event.stopPropagation();
      }
    },
    keyPress({ key }) {
      if (this.$refs.dropdown && this.$refs.dropdown.isActive && (key === "Escape" || key === "Esc")) {
        this.togglePicker(false);
      }
    },
    onActiveChange(value) {
      if (!value) {
        this.onBlur();
      }
      this.$emit("active-change", value);
    },
    changeFocus(day) {
      this.focusedDateData = {
        day: day.getDate(),
        month: day.getMonth(),
        year: day.getFullYear()
      };
    }
  },
  created() {
    if (typeof window !== "undefined") {
      document.addEventListener("keyup", this.keyPress);
    }
  },
  beforeUnmount() {
    if (typeof window !== "undefined") {
      document.removeEventListener("keyup", this.keyPress);
    }
  }
});
function render$2$3(_ctx, _cache, $props, $setup, $data, $options) {
  const _component_o_input = vue.resolveComponent("o-input");
  const _component_o_icon = vue.resolveComponent("o-icon");
  const _component_o_select = vue.resolveComponent("o-select");
  const _component_o_field = vue.resolveComponent("o-field");
  const _component_o_datepicker_table = vue.resolveComponent("o-datepicker-table");
  const _component_o_dropdown_item = vue.resolveComponent("o-dropdown-item");
  const _component_o_dropdown = vue.resolveComponent("o-dropdown");
  return vue.openBlock(), vue.createBlock("div", {
    class: _ctx.rootClasses
  }, [!_ctx.isMobile || _ctx.inline ? vue.createVNode(_component_o_dropdown, vue.mergeProps({
    key: 0,
    ref: "dropdown"
  }, _ctx.dropdownBind, {
    position: _ctx.position,
    disabled: _ctx.disabled,
    inline: _ctx.inline,
    "mobile-modal": _ctx.mobileModal,
    "trap-focus": _ctx.trapFocus,
    "aria-role": _ctx.ariaRole,
    "aria-modal": !_ctx.inline,
    "append-to-body": _ctx.appendToBody,
    "append-to-body-copy-parent": "",
    "onActive-change": _ctx.onActiveChange
  }), vue.createSlots({
    default: vue.withCtx(() => [vue.createVNode(_component_o_dropdown_item, {
      override: "",
      disabled: _ctx.disabled,
      clickable: false
    }, {
      default: vue.withCtx(() => [vue.createVNode("div", {
        class: _ctx.boxClasses
      }, [vue.createVNode("header", {
        class: _ctx.headerClasses
      }, [_ctx.$slots.header !== void 0 && _ctx.$slots.header.length ? vue.renderSlot(_ctx.$slots, "header", {
        key: 0
      }) : (vue.openBlock(), vue.createBlock("div", {
        key: 1,
        class: _ctx.headerButtonsClasses
      }, [vue.withDirectives(vue.createVNode("a", {
        class: _ctx.prevBtnClasses,
        role: "button",
        href: "#",
        "aria-label": _ctx.ariaPreviousLabel,
        onClick: _cache[4] || (_cache[4] = vue.withModifiers((...args) => _ctx.prev(...args), ["prevent"])),
        onKeydown: [_cache[5] || (_cache[5] = vue.withKeys(vue.withModifiers((...args) => _ctx.prev(...args), ["prevent"]), ["enter"])), _cache[6] || (_cache[6] = vue.withKeys(vue.withModifiers((...args) => _ctx.prev(...args), ["prevent"]), ["space"]))]
      }, [vue.createVNode(_component_o_icon, {
        icon: _ctx.iconPrev,
        pack: _ctx.iconPack,
        both: "",
        clickable: ""
      }, null, 8, ["icon", "pack"])], 42, ["aria-label"]), [[vue.vShow, !_ctx.showPrev && !_ctx.disabled]]), vue.withDirectives(vue.createVNode("a", {
        class: _ctx.nextBtnClasses,
        role: "button",
        href: "#",
        "aria-label": _ctx.ariaNextLabel,
        onClick: _cache[7] || (_cache[7] = vue.withModifiers((...args) => _ctx.next(...args), ["prevent"])),
        onKeydown: [_cache[8] || (_cache[8] = vue.withKeys(vue.withModifiers((...args) => _ctx.next(...args), ["prevent"]), ["enter"])), _cache[9] || (_cache[9] = vue.withKeys(vue.withModifiers((...args) => _ctx.next(...args), ["prevent"]), ["space"]))]
      }, [vue.createVNode(_component_o_icon, {
        icon: _ctx.iconNext,
        pack: _ctx.iconPack,
        both: "",
        clickable: ""
      }, null, 8, ["icon", "pack"])], 42, ["aria-label"]), [[vue.vShow, !_ctx.showNext && !_ctx.disabled]]), vue.createVNode("div", {
        class: _ctx.listsClasses
      }, [vue.createVNode(_component_o_field, null, {
        default: vue.withCtx(() => [!_ctx.isTypeMonth ? vue.createVNode(_component_o_select, {
          key: 0,
          modelValue: _ctx.focusedDateData.month,
          "onUpdate:modelValue": _cache[10] || (_cache[10] = ($event) => _ctx.focusedDateData.month = $event),
          disabled: _ctx.disabled,
          size: _ctx.size
        }, {
          default: vue.withCtx(() => [(vue.openBlock(true), vue.createBlock(vue.Fragment, null, vue.renderList(_ctx.listOfMonths, (month) => {
            return vue.openBlock(), vue.createBlock("option", {
              value: month.index,
              key: month.name,
              disabled: month.disabled
            }, vue.toDisplayString(month.name), 9, ["value", "disabled"]);
          }), 128))]),
          _: 1
        }, 8, ["modelValue", "disabled", "size"]) : vue.createCommentVNode("v-if", true), vue.createVNode(_component_o_select, {
          modelValue: _ctx.focusedDateData.year,
          "onUpdate:modelValue": _cache[11] || (_cache[11] = ($event) => _ctx.focusedDateData.year = $event),
          disabled: _ctx.disabled,
          size: _ctx.size
        }, {
          default: vue.withCtx(() => [(vue.openBlock(true), vue.createBlock(vue.Fragment, null, vue.renderList(_ctx.listOfYears, (year) => {
            return vue.openBlock(), vue.createBlock("option", {
              value: year,
              key: year
            }, vue.toDisplayString(year), 9, ["value"]);
          }), 128))]),
          _: 1
        }, 8, ["modelValue", "disabled", "size"])]),
        _: 1
      })], 2)], 2))], 2), vue.createVNode(_component_o_datepicker_table, {
        modelValue: _ctx.computedValue,
        "onUpdate:modelValue": _cache[12] || (_cache[12] = ($event) => _ctx.computedValue = $event),
        "day-names": _ctx.newDayNames,
        "month-names": _ctx.newMonthNames,
        "first-day-of-week": _ctx.firstDayOfWeek,
        "rules-for-first-week": _ctx.rulesForFirstWeek,
        "min-date": _ctx.minDate,
        "max-date": _ctx.maxDate,
        focused: _ctx.focusedDateData,
        disabled: _ctx.disabled,
        "unselectable-dates": _ctx.unselectableDates,
        "unselectable-days-of-week": _ctx.unselectableDaysOfWeek,
        "selectable-dates": _ctx.selectableDates,
        events: _ctx.events,
        indicators: _ctx.indicators,
        "date-creator": _ctx.dateCreator,
        "type-month": _ctx.isTypeMonth,
        "nearby-month-days": _ctx.nearbyMonthDays,
        "nearby-selectable-month-days": _ctx.nearbySelectableMonthDays,
        "show-week-number": _ctx.showWeekNumber,
        "week-number-clickable": _ctx.weekNumberClickable,
        range: _ctx.range,
        multiple: _ctx.multiple,
        "table-class": _ctx.tableClass,
        "table-head-class": _ctx.tableHeadClass,
        "table-head-cell-class": _ctx.tableHeadCellClass,
        "table-body-class": _ctx.tableBodyClass,
        "table-row-class": _ctx.tableRowClass,
        "table-cell-class": _ctx.tableCellClass,
        "table-cell-selected-class": _ctx.tableCellSelectedClass,
        "table-cell-first-selected-class": _ctx.tableCellFirstSelectedClass,
        "table-cell-invisible-class": _ctx.tableCellInvisibleClass,
        "table-cell-within-selected-class": _ctx.tableCellWithinSelectedClass,
        "table-cell-last-selected-class": _ctx.tableCellLastSelectedClass,
        "table-cell-first-hovered-class": _ctx.tableCellFirstHoveredClass,
        "table-cell-within-hovered-class": _ctx.tableCellWithinHoveredClass,
        "table-cell-last-hovered-class": _ctx.tableCellLastHoveredClass,
        "table-cell-today-class": _ctx.tableCellTodayClass,
        "table-cell-selectable-class": _ctx.tableCellSelectableClass,
        "table-cell-unselectable-class": _ctx.tableCellUnselectableClass,
        "table-cell-nearby-class": _ctx.tableCellNearbyClass,
        "table-cell-events-class": _ctx.tableCellEventsClass,
        "table-events-class": _ctx.tableEventsClass,
        "table-event-variant-class": _ctx.tableEventVariantClass,
        "table-event-class": _ctx.tableEventClass,
        "table-event-indicators-class": _ctx.tableEventIndicatorsClass,
        "onRange-start": _cache[13] || (_cache[13] = (date) => _ctx.$emit("range-start", date)),
        "onRange-end": _cache[14] || (_cache[14] = (date) => _ctx.$emit("range-end", date)),
        onClose: _cache[15] || (_cache[15] = ($event) => _ctx.togglePicker(false)),
        "onUpdate:focused": _cache[16] || (_cache[16] = ($event) => _ctx.focusedDateData = $event)
      }, null, 8, ["modelValue", "day-names", "month-names", "first-day-of-week", "rules-for-first-week", "min-date", "max-date", "focused", "disabled", "unselectable-dates", "unselectable-days-of-week", "selectable-dates", "events", "indicators", "date-creator", "type-month", "nearby-month-days", "nearby-selectable-month-days", "show-week-number", "week-number-clickable", "range", "multiple", "table-class", "table-head-class", "table-head-cell-class", "table-body-class", "table-row-class", "table-cell-class", "table-cell-selected-class", "table-cell-first-selected-class", "table-cell-invisible-class", "table-cell-within-selected-class", "table-cell-last-selected-class", "table-cell-first-hovered-class", "table-cell-within-hovered-class", "table-cell-last-hovered-class", "table-cell-today-class", "table-cell-selectable-class", "table-cell-unselectable-class", "table-cell-nearby-class", "table-cell-events-class", "table-events-class", "table-event-variant-class", "table-event-class", "table-event-indicators-class"])], 2), _ctx.$slots.default !== void 0 ? (vue.openBlock(), vue.createBlock("footer", {
        key: 0,
        class: _ctx.footerClasses
      }, [vue.renderSlot(_ctx.$slots, "default")], 2)) : vue.createCommentVNode("v-if", true)]),
      _: 1
    }, 8, ["disabled"])]),
    _: 2
  }, [!_ctx.inline ? {
    name: "trigger",
    fn: vue.withCtx(() => [vue.renderSlot(_ctx.$slots, "trigger", {}, () => [vue.createVNode(_component_o_input, vue.mergeProps({
      ref: "input",
      autocomplete: "off",
      "model-value": _ctx.formattedValue,
      expanded: _ctx.expanded,
      placeholder: _ctx.placeholder,
      size: _ctx.size,
      icon: _ctx.icon,
      "icon-right": _ctx.iconRight,
      "icon-right-clickable": _ctx.iconRightClickable,
      "icon-pack": _ctx.iconPack,
      rounded: _ctx.rounded,
      disabled: _ctx.disabled,
      readonly: !_ctx.editable
    }, _ctx.inputBind, {
      "use-html5-validation": false,
      onClick: _ctx.onInputClick,
      "onIcon-right-click": _cache[1] || (_cache[1] = ($event) => _ctx.$emit("icon-right-click")),
      onKeyup: _cache[2] || (_cache[2] = vue.withKeys(($event) => _ctx.togglePicker(true), ["enter"])),
      onChange: _cache[3] || (_cache[3] = ($event) => _ctx.onChange($event.target.value)),
      onFocus: _ctx.handleOnFocus
    }), null, 16, ["model-value", "expanded", "placeholder", "size", "icon", "icon-right", "icon-right-clickable", "icon-pack", "rounded", "disabled", "readonly", "onClick", "onFocus"])])])
  } : void 0]), 1040, ["position", "disabled", "inline", "mobile-modal", "trap-focus", "aria-role", "aria-modal", "append-to-body", "onActive-change"]) : vue.createVNode(_component_o_input, vue.mergeProps({
    key: 1,
    ref: "input",
    type: !_ctx.isTypeMonth ? "date" : "month",
    autocomplete: "off",
    value: _ctx.formatNative(_ctx.computedValue),
    placeholder: _ctx.placeholder,
    size: _ctx.size,
    icon: _ctx.icon,
    "icon-pack": _ctx.iconPack,
    rounded: _ctx.rounded,
    max: _ctx.formatNative(_ctx.maxDate),
    min: _ctx.formatNative(_ctx.minDate),
    disabled: _ctx.disabled,
    readonly: false
  }, _ctx.$attrs, {
    "use-html5-validation": false,
    onChange: _ctx.onChangeNativePicker,
    onFocus: _ctx.onFocus,
    onBlur: _ctx.onBlur
  }), null, 16, ["type", "value", "placeholder", "size", "icon", "icon-pack", "rounded", "max", "min", "disabled", "onChange", "onFocus", "onBlur"])], 2);
}
script$2$3.render = render$2$3;
script$2$3.__file = "src/components/datepicker/Datepicker.vue";
var index$z = {
  install(app) {
    registerComponent(app, script$2$3);
  }
};
var index$4$1 = index$z;
const AM$1 = "AM";
const PM$1 = "PM";
const HOUR_FORMAT_24 = "24";
const HOUR_FORMAT_12 = "12";
const defaultTimeFormatter = (date, vm) => {
  return vm.dtf.format(date);
};
const defaultTimeParser = (timeString, vm) => {
  if (timeString) {
    let d = null;
    if (vm.computedValue && !isNaN(vm.computedValue)) {
      d = new Date(vm.computedValue);
    } else {
      d = vm.timeCreator();
      d.setMilliseconds(0);
    }
    if (vm.dtf.formatToParts && typeof vm.dtf.formatToParts === "function") {
      const formatRegex = vm.dtf.formatToParts(d).map((part) => {
        if (part.type === "literal") {
          return part.value.replace(/ /g, "\\s?");
        } else if (part.type === "dayPeriod") {
          return `((?!=<${part.type}>)(${vm.amString}|${vm.pmString}|${AM$1}|${PM$1}|${AM$1.toLowerCase()}|${PM$1.toLowerCase()})?)`;
        }
        return `((?!=<${part.type}>)\\d+)`;
      }).join("");
      const timeGroups = matchWithGroups(formatRegex, timeString);
      timeGroups.hour = timeGroups.hour ? parseInt(timeGroups.hour, 10) : null;
      timeGroups.minute = timeGroups.minute ? parseInt(timeGroups.minute, 10) : null;
      timeGroups.second = timeGroups.second ? parseInt(timeGroups.second, 10) : null;
      if (timeGroups.hour && timeGroups.hour >= 0 && timeGroups.hour < 24 && timeGroups.minute && timeGroups.minute >= 0 && timeGroups.minute < 59) {
        if (timeGroups.dayPeriod && (timeGroups.dayPeriod.toLowerCase() === vm.pmString.toLowerCase() || timeGroups.dayPeriod.toLowerCase() === PM$1.toLowerCase()) && timeGroups.hour < 12) {
          timeGroups.hour += 12;
        }
        d.setHours(timeGroups.hour);
        d.setMinutes(timeGroups.minute);
        d.setSeconds(timeGroups.second || 0);
        return d;
      }
    }
    let am = false;
    if (vm.hourFormat === HOUR_FORMAT_12) {
      const dateString12 = timeString.split(" ");
      timeString = dateString12[0];
      am = dateString12[1] === vm.amString || dateString12[1] === AM$1;
    }
    const time = timeString.split(":");
    let hours = parseInt(time[0], 10);
    const minutes = parseInt(time[1], 10);
    const seconds = vm.enableSeconds ? parseInt(time[2], 10) : 0;
    if (isNaN(hours) || hours < 0 || hours > 23 || vm.hourFormat === HOUR_FORMAT_12 && (hours < 1 || hours > 12) || isNaN(minutes) || minutes < 0 || minutes > 59) {
      return null;
    }
    d.setSeconds(seconds);
    d.setMinutes(minutes);
    if (vm.hourFormat === HOUR_FORMAT_12) {
      if (am && hours === 12) {
        hours = 0;
      } else if (!am && hours !== 12) {
        hours += 12;
      }
    }
    d.setHours(hours);
    return new Date(d.getTime());
  }
  return null;
};
var TimepickerMixin = {
  mixins: [FormElementMixin],
  inheritAttrs: false,
  props: {
    modelValue: Date,
    inline: Boolean,
    minTime: Date,
    maxTime: Date,
    placeholder: String,
    editable: Boolean,
    disabled: Boolean,
    size: String,
    hourFormat: {
      type: String,
      validator: (value) => {
        return value === HOUR_FORMAT_24 || value === HOUR_FORMAT_12;
      }
    },
    incrementHours: {
      type: Number,
      default: 1
    },
    incrementMinutes: {
      type: Number,
      default: 1
    },
    incrementSeconds: {
      type: Number,
      default: 1
    },
    timeFormatter: {
      type: Function,
      default: (date, vm) => {
        const timeFormatter = getValueByPath(getOptions$1(), "timepicker.timeFormatter", void 0);
        if (typeof timeFormatter === "function") {
          return timeFormatter(date);
        } else {
          return defaultTimeFormatter(date, vm);
        }
      }
    },
    timeParser: {
      type: Function,
      default: (date, vm) => {
        const timeParser = getValueByPath(getOptions$1(), "timepicker.timeParser", void 0);
        if (typeof timeParser === "function") {
          return timeParser(date);
        } else {
          return defaultTimeParser(date, vm);
        }
      }
    },
    mobileNative: {
      type: Boolean,
      default: () => {
        return getValueByPath(getOptions$1(), "timepicker.mobileNative", true);
      }
    },
    timeCreator: {
      type: Function,
      default: () => {
        const timeCreator = getValueByPath(getOptions$1(), "timepicker.timeCreator", void 0);
        if (typeof timeCreator === "function") {
          return timeCreator();
        } else {
          return new Date();
        }
      }
    },
    position: String,
    unselectableTimes: Array,
    openOnFocus: Boolean,
    enableSeconds: Boolean,
    defaultMinutes: Number,
    defaultSeconds: Number,
    appendToBody: Boolean,
    resetOnMeridianChange: {
      type: Boolean,
      default: false
    }
  },
  data() {
    return {
      dateSelected: this.modelValue,
      hoursSelected: null,
      minutesSelected: null,
      secondsSelected: null,
      meridienSelected: null,
      _elementRef: "input"
    };
  },
  computed: {
    computedValue: {
      get() {
        return this.dateSelected;
      },
      set(value) {
        this.dateSelected = value;
        this.$emit("update:modelValue", this.dateSelected);
      }
    },
    localeOptions() {
      return new Intl.DateTimeFormat(this.locale, {
        hour: "numeric",
        minute: "numeric",
        second: this.enableSeconds ? "numeric" : void 0
      }).resolvedOptions();
    },
    dtf() {
      return new Intl.DateTimeFormat(this.locale, {
        hour: this.localeOptions.hour || "numeric",
        minute: this.localeOptions.minute || "numeric",
        second: this.enableSeconds ? this.localeOptions.second || "numeric" : void 0,
        hour12: !this.isHourFormat24
      });
    },
    newHourFormat() {
      return this.hourFormat || (this.localeOptions.hour12 ? HOUR_FORMAT_12 : HOUR_FORMAT_24);
    },
    sampleTime() {
      let d = this.timeCreator();
      d.setHours(10);
      d.setSeconds(0);
      d.setMinutes(0);
      d.setMilliseconds(0);
      return d;
    },
    hourLiteral() {
      if (this.dtf.formatToParts && typeof this.dtf.formatToParts === "function") {
        let d = this.sampleTime;
        const parts = this.dtf.formatToParts(d);
        const literal = parts.find((part, idx) => idx > 0 && parts[idx - 1].type === "hour");
        if (literal) {
          return literal.value;
        }
      }
      return ":";
    },
    minuteLiteral() {
      if (this.dtf.formatToParts && typeof this.dtf.formatToParts === "function") {
        let d = this.sampleTime;
        const parts = this.dtf.formatToParts(d);
        const literal = parts.find((part, idx) => idx > 0 && parts[idx - 1].type === "minute");
        if (literal) {
          return literal.value;
        }
      }
      return ":";
    },
    secondLiteral() {
      if (this.dtf.formatToParts && typeof this.dtf.formatToParts === "function") {
        let d = this.sampleTime;
        const parts = this.dtf.formatToParts(d);
        const literal = parts.find((part, idx) => idx > 0 && parts[idx - 1].type === "second");
        if (literal) {
          return literal.value;
        }
      }
    },
    amString() {
      if (this.dtf.formatToParts && typeof this.dtf.formatToParts === "function") {
        let d = this.sampleTime;
        d.setHours(10);
        const dayPeriod = this.dtf.formatToParts(d).find((part) => part.type === "dayPeriod");
        if (dayPeriod) {
          return dayPeriod.value;
        }
      }
      return AM$1;
    },
    pmString() {
      if (this.dtf.formatToParts && typeof this.dtf.formatToParts === "function") {
        let d = this.sampleTime;
        d.setHours(20);
        const dayPeriod = this.dtf.formatToParts(d).find((part) => part.type === "dayPeriod");
        if (dayPeriod) {
          return dayPeriod.value;
        }
      }
      return PM$1;
    },
    hours() {
      if (!this.incrementHours || this.incrementHours < 1)
        throw new Error("Hour increment cannot be null or less than 1.");
      const hours = [];
      const numberOfHours = this.isHourFormat24 ? 24 : 12;
      for (let i = 0; i < numberOfHours; i += this.incrementHours) {
        let value = i;
        let label = value;
        if (!this.isHourFormat24) {
          value = i + 1;
          label = value;
          if (this.meridienSelected === this.amString) {
            if (value === 12) {
              value = 0;
            }
          } else if (this.meridienSelected === this.pmString) {
            if (value !== 12) {
              value += 12;
            }
          }
        }
        hours.push({
          label: this.formatNumber(label),
          value
        });
      }
      return hours;
    },
    minutes() {
      if (!this.incrementMinutes || this.incrementMinutes < 1)
        throw new Error("Minute increment cannot be null or less than 1.");
      const minutes = [];
      for (let i = 0; i < 60; i += this.incrementMinutes) {
        minutes.push({
          label: this.formatNumber(i, true),
          value: i
        });
      }
      return minutes;
    },
    seconds() {
      if (!this.incrementSeconds || this.incrementSeconds < 1)
        throw new Error("Second increment cannot be null or less than 1.");
      const seconds = [];
      for (let i = 0; i < 60; i += this.incrementSeconds) {
        seconds.push({
          label: this.formatNumber(i, true),
          value: i
        });
      }
      return seconds;
    },
    meridiens() {
      return [this.amString, this.pmString];
    },
    isMobile() {
      return this.mobileNative && isMobile.any();
    },
    isHourFormat24() {
      return this.newHourFormat === HOUR_FORMAT_24;
    }
  },
  watch: {
    hourFormat() {
      if (this.hoursSelected !== null) {
        this.meridienSelected = this.hoursSelected >= 12 ? this.pmString : this.amString;
      }
    },
    locale() {
      if (!this.value) {
        this.meridienSelected = this.amString;
      }
    },
    modelValue: {
      handler(value) {
        this.updateInternalState(value);
        !this.isValid && this.$refs.input.checkHtml5Validity();
      },
      immediate: true
    }
  },
  methods: {
    onMeridienChange(value) {
      if (this.hoursSelected !== null && this.resetOnMeridianChange) {
        this.hoursSelected = null;
        this.minutesSelected = null;
        this.secondsSelected = null;
        this.computedValue = null;
      } else if (this.hoursSelected !== null) {
        if (value === this.pmString) {
          this.hoursSelected += 12;
        } else if (value === this.amString) {
          this.hoursSelected -= 12;
        }
      }
      this.updateDateSelected(this.hoursSelected, this.minutesSelected, this.enableSeconds ? this.secondsSelected : 0, value);
    },
    onHoursChange(value) {
      if (!this.minutesSelected && typeof this.defaultMinutes !== "undefined") {
        this.minutesSelected = this.defaultMinutes;
      }
      if (!this.secondsSelected && typeof this.defaultSeconds !== "undefined") {
        this.secondsSelected = this.defaultSeconds;
      }
      this.updateDateSelected(parseInt(value, 10), this.minutesSelected, this.enableSeconds ? this.secondsSelected : 0, this.meridienSelected);
    },
    onMinutesChange(value) {
      if (!this.secondsSelected && this.defaultSeconds) {
        this.secondsSelected = this.defaultSeconds;
      }
      this.updateDateSelected(this.hoursSelected, parseInt(value, 10), this.enableSeconds ? this.secondsSelected : 0, this.meridienSelected);
    },
    onSecondsChange(value) {
      this.updateDateSelected(this.hoursSelected, this.minutesSelected, parseInt(value, 10), this.meridienSelected);
    },
    updateDateSelected(hours, minutes, seconds, meridiens) {
      if (hours != null && minutes != null && (!this.isHourFormat24 && meridiens !== null || this.isHourFormat24)) {
        let time = null;
        if (this.computedValue && !isNaN(this.computedValue)) {
          time = new Date(this.computedValue);
        } else {
          time = this.timeCreator();
          time.setMilliseconds(0);
        }
        time.setHours(hours);
        time.setMinutes(minutes);
        time.setSeconds(seconds);
        if (!isNaN(time.getTime())) {
          this.computedValue = new Date(time.getTime());
        }
      }
    },
    updateInternalState(value) {
      if (value) {
        this.hoursSelected = value.getHours();
        this.minutesSelected = value.getMinutes();
        this.secondsSelected = value.getSeconds();
        this.meridienSelected = value.getHours() >= 12 ? this.pmString : this.amString;
      } else {
        this.hoursSelected = null;
        this.minutesSelected = null;
        this.secondsSelected = null;
        this.meridienSelected = this.amString;
      }
      this.dateSelected = value;
    },
    isHourDisabled(hour) {
      let disabled = false;
      if (this.minTime) {
        const minHours = this.minTime.getHours();
        const noMinutesAvailable = this.minutes.every((minute) => {
          return this.isMinuteDisabledForHour(hour, minute.value);
        });
        disabled = hour < minHours || noMinutesAvailable;
      }
      if (this.maxTime) {
        if (!disabled) {
          const maxHours = this.maxTime.getHours();
          disabled = hour > maxHours;
        }
      }
      if (this.unselectableTimes) {
        if (!disabled) {
          const unselectable = this.unselectableTimes.filter((time) => {
            if (this.enableSeconds && this.secondsSelected !== null) {
              return time.getHours() === hour && time.getMinutes() === this.minutesSelected && time.getSeconds() === this.secondsSelected;
            } else if (this.minutesSelected !== null) {
              return time.getHours() === hour && time.getMinutes() === this.minutesSelected;
            }
            return false;
          });
          if (unselectable.length > 0) {
            disabled = true;
          } else {
            disabled = this.minutes.every((minute) => {
              return this.unselectableTimes.filter((time) => {
                return time.getHours() === hour && time.getMinutes() === minute.value;
              }).length > 0;
            });
          }
        }
      }
      return disabled;
    },
    isMinuteDisabledForHour(hour, minute) {
      let disabled = false;
      if (this.minTime) {
        const minHours = this.minTime.getHours();
        const minMinutes = this.minTime.getMinutes();
        disabled = hour === minHours && minute < minMinutes;
      }
      if (this.maxTime) {
        if (!disabled) {
          const maxHours = this.maxTime.getHours();
          const maxMinutes = this.maxTime.getMinutes();
          disabled = hour === maxHours && minute > maxMinutes;
        }
      }
      return disabled;
    },
    isMinuteDisabled(minute) {
      let disabled = false;
      if (this.hoursSelected !== null) {
        if (this.isHourDisabled(this.hoursSelected)) {
          disabled = true;
        } else {
          disabled = this.isMinuteDisabledForHour(this.hoursSelected, minute);
        }
        if (this.unselectableTimes) {
          if (!disabled) {
            const unselectable = this.unselectableTimes.filter((time) => {
              if (this.enableSeconds && this.secondsSelected !== null) {
                return time.getHours() === this.hoursSelected && time.getMinutes() === minute && time.getSeconds() === this.secondsSelected;
              } else {
                return time.getHours() === this.hoursSelected && time.getMinutes() === minute;
              }
            });
            disabled = unselectable.length > 0;
          }
        }
      }
      return disabled;
    },
    isSecondDisabled(second) {
      let disabled = false;
      if (this.minutesSelected !== null) {
        if (this.isMinuteDisabled(this.minutesSelected)) {
          disabled = true;
        } else {
          if (this.minTime) {
            const minHours = this.minTime.getHours();
            const minMinutes = this.minTime.getMinutes();
            const minSeconds = this.minTime.getSeconds();
            disabled = this.hoursSelected === minHours && this.minutesSelected === minMinutes && second < minSeconds;
          }
          if (this.maxTime) {
            if (!disabled) {
              const maxHours = this.maxTime.getHours();
              const maxMinutes = this.maxTime.getMinutes();
              const maxSeconds = this.maxTime.getSeconds();
              disabled = this.hoursSelected === maxHours && this.minutesSelected === maxMinutes && second > maxSeconds;
            }
          }
        }
        if (this.unselectableTimes) {
          if (!disabled) {
            const unselectable = this.unselectableTimes.filter((time) => {
              return time.getHours() === this.hoursSelected && time.getMinutes() === this.minutesSelected && time.getSeconds() === second;
            });
            disabled = unselectable.length > 0;
          }
        }
      }
      return disabled;
    },
    onChange(value) {
      const date = this.timeParser(value, this);
      this.updateInternalState(date);
      if (date && !isNaN(date)) {
        this.computedValue = date;
      } else {
        this.computedValue = null;
        this.$refs.input.newValue = this.computedValue;
      }
    },
    toggle(active) {
      if (this.$refs.dropdown) {
        this.$refs.dropdown.isActive = typeof active === "boolean" ? active : !this.$refs.dropdown.isActive;
      }
    },
    close() {
      this.toggle(false);
    },
    handleOnFocus() {
      this.onFocus();
      if (this.openOnFocus) {
        this.toggle(true);
      }
    },
    formatHHMMSS(value) {
      const date = new Date(value);
      if (value && !isNaN(date.getTime())) {
        const hours = date.getHours();
        const minutes = date.getMinutes();
        const seconds = date.getSeconds();
        return this.formatNumber(hours, true) + ":" + this.formatNumber(minutes, true) + ":" + this.formatNumber(seconds, true);
      }
      return "";
    },
    onChangeNativePicker(event) {
      const date = event.target.value;
      if (date) {
        let time = null;
        if (this.computedValue && !isNaN(this.computedValue)) {
          time = new Date(this.computedValue);
        } else {
          time = new Date();
          time.setMilliseconds(0);
        }
        const t = date.split(":");
        time.setHours(parseInt(t[0], 10));
        time.setMinutes(parseInt(t[1], 10));
        time.setSeconds(t[2] ? parseInt(t[2], 10) : 0);
        this.computedValue = new Date(time.getTime());
      } else {
        this.computedValue = null;
      }
    },
    formatNumber(value, prependZero) {
      return this.isHourFormat24 || prependZero ? this.pad(value) : value;
    },
    pad(value) {
      return (value < 10 ? "0" : "") + value;
    },
    formatValue(date) {
      if (date && !isNaN(date)) {
        return this.timeFormatter(date, this);
      } else {
        return null;
      }
    },
    keyPress({ key }) {
      if (this.$refs.dropdown && this.$refs.dropdown.isActive && (key === "Escape" || key === "Esc")) {
        this.toggle(false);
      }
    },
    onActiveChange(value) {
      if (!value) {
        this.onBlur();
      }
    }
  },
  created() {
    if (typeof window !== "undefined") {
      document.addEventListener("keyup", this.keyPress);
    }
  },
  beforeUnmount() {
    if (typeof window !== "undefined") {
      document.removeEventListener("keyup", this.keyPress);
    }
  }
};
var script$h = vue.defineComponent({
  name: "OTimepicker",
  components: {
    [script$q.name]: script$q,
    [script$j.name]: script$j,
    [script$r.name]: script$r,
    [script$l.name]: script$l,
    [script$1$9.name]: script$1$9
  },
  configField: "timepicker",
  mixins: [BaseComponentMixin, TimepickerMixin, MatchMediaMixin],
  inheritAttrs: false,
  props: {
    rootClass: [String, Function, Array],
    sizeClass: [String, Function, Array],
    boxClass: [String, Function, Array],
    selectClass: [String, Function, Array],
    selectPlaceholderClass: [String, Function, Array],
    separatorClass: [String, Function, Array],
    footerClass: [String, Function, Array],
    inputClasses: Object,
    dropdownClasses: Object
  },
  computed: {
    inputBind() {
      return __spreadValues(__spreadValues({}, this.$attrs), this.inputClasses);
    },
    dropdownBind() {
      return __spreadValues({
        "root-class": this.computedClass("dropdownClasses.rootClass", "o-tpck__dropdown")
      }, this.dropdownClasses);
    },
    rootClasses() {
      return [
        this.computedClass("rootClass", "o-tpck"),
        { [this.computedClass("sizeClass", "o-tpck--", this.size)]: this.size },
        { [this.computedClass("mobileClass", "o-tpck--mobile")]: this.isMatchMedia }
      ];
    },
    boxClasses() {
      return [
        this.computedClass("boxClass", "o-tpck__box")
      ];
    },
    selectClasses() {
      return [
        this.computedClass("selectClass", "o-tpck__select")
      ];
    },
    selectPlaceholderClasses() {
      return [
        this.computedClass("selectPlaceholderClass", "o-tpck__select-placeholder")
      ];
    },
    separatorClasses() {
      return [
        this.computedClass("separatorClass", "o-tpck__separator")
      ];
    },
    footerClasses() {
      return [
        this.computedClass("footerClass", "o-tpck__footer")
      ];
    },
    nativeStep() {
      if (this.enableSeconds)
        return "1";
      return null;
    }
  }
});
function render$f(_ctx, _cache, $props, $setup, $data, $options) {
  const _component_o_input = vue.resolveComponent("o-input");
  const _component_o_select = vue.resolveComponent("o-select");
  const _component_o_dropdown_item = vue.resolveComponent("o-dropdown-item");
  const _component_o_dropdown = vue.resolveComponent("o-dropdown");
  return vue.openBlock(), vue.createBlock("div", {
    class: _ctx.rootClasses
  }, [!_ctx.isMobile || _ctx.inline ? vue.createVNode(_component_o_dropdown, vue.mergeProps({
    key: 0,
    ref: "dropdown"
  }, _ctx.dropdownBind, {
    position: _ctx.position,
    disabled: _ctx.disabled,
    inline: _ctx.inline,
    "append-to-body": _ctx.appendToBody,
    "append-to-body-copy-parent": "",
    "onActive-change": _ctx.onActiveChange
  }), vue.createSlots({
    default: vue.withCtx(() => [vue.createVNode(_component_o_dropdown_item, {
      override: "",
      disabled: _ctx.disabled,
      clickable: false
    }, {
      default: vue.withCtx(() => [vue.createVNode("div", {
        class: _ctx.boxClasses
      }, [vue.createVNode(_component_o_select, {
        override: "",
        "select-class": _ctx.selectClasses,
        "placeholder-class": _ctx.selectPlaceholderClasses,
        modelValue: _ctx.hoursSelected,
        "onUpdate:modelValue": _cache[3] || (_cache[3] = ($event) => _ctx.hoursSelected = $event),
        onChange: _cache[4] || (_cache[4] = ($event) => _ctx.onHoursChange($event.target.value)),
        disabled: _ctx.disabled,
        placeholder: "00"
      }, {
        default: vue.withCtx(() => [(vue.openBlock(true), vue.createBlock(vue.Fragment, null, vue.renderList(_ctx.hours, (hour) => {
          return vue.openBlock(), vue.createBlock("option", {
            value: hour.value,
            key: hour.value,
            disabled: _ctx.isHourDisabled(hour.value)
          }, vue.toDisplayString(hour.label), 9, ["value", "disabled"]);
        }), 128))]),
        _: 1
      }, 8, ["select-class", "placeholder-class", "modelValue", "disabled"]), vue.createVNode("span", {
        class: _ctx.separatorClasses
      }, vue.toDisplayString(_ctx.hourLiteral), 3), vue.createVNode(_component_o_select, {
        override: "",
        "select-class": _ctx.selectClasses,
        "placeholder-class": _ctx.selectPlaceholderClasses,
        modelValue: _ctx.minutesSelected,
        "onUpdate:modelValue": _cache[5] || (_cache[5] = ($event) => _ctx.minutesSelected = $event),
        onChange: _cache[6] || (_cache[6] = ($event) => _ctx.onMinutesChange($event.target.value)),
        disabled: _ctx.disabled,
        placeholder: "00"
      }, {
        default: vue.withCtx(() => [(vue.openBlock(true), vue.createBlock(vue.Fragment, null, vue.renderList(_ctx.minutes, (minute) => {
          return vue.openBlock(), vue.createBlock("option", {
            value: minute.value,
            key: minute.value,
            disabled: _ctx.isMinuteDisabled(minute.value)
          }, vue.toDisplayString(minute.label), 9, ["value", "disabled"]);
        }), 128))]),
        _: 1
      }, 8, ["select-class", "placeholder-class", "modelValue", "disabled"]), _ctx.enableSeconds ? (vue.openBlock(), vue.createBlock(vue.Fragment, {
        key: 0
      }, [vue.createVNode("span", null, vue.toDisplayString(_ctx.minuteLiteral), 1), vue.createVNode(_component_o_select, {
        override: "",
        "select-class": _ctx.selectClasses,
        "placeholder-class": _ctx.selectPlaceholderClasses,
        modelValue: _ctx.secondsSelected,
        "onUpdate:modelValue": _cache[7] || (_cache[7] = ($event) => _ctx.secondsSelected = $event),
        onChange: _cache[8] || (_cache[8] = ($event) => _ctx.onSecondsChange($event.target.value)),
        disabled: _ctx.disabled,
        placeholder: "00"
      }, {
        default: vue.withCtx(() => [(vue.openBlock(true), vue.createBlock(vue.Fragment, null, vue.renderList(_ctx.seconds, (second) => {
          return vue.openBlock(), vue.createBlock("option", {
            value: second.value,
            key: second.value,
            disabled: _ctx.isSecondDisabled(second.value)
          }, vue.toDisplayString(second.label), 9, ["value", "disabled"]);
        }), 128))]),
        _: 1
      }, 8, ["select-class", "placeholder-class", "modelValue", "disabled"]), vue.createVNode("span", {
        class: _ctx.separatorClasses
      }, vue.toDisplayString(_ctx.secondLiteral), 3)], 64)) : vue.createCommentVNode("v-if", true), !_ctx.isHourFormat24 ? vue.createVNode(_component_o_select, {
        key: 1,
        override: "",
        "select-class": _ctx.selectClasses,
        "placeholder-class": _ctx.selectPlaceholderClasses,
        modelValue: _ctx.meridienSelected,
        "onUpdate:modelValue": _cache[9] || (_cache[9] = ($event) => _ctx.meridienSelected = $event),
        onChange: _cache[10] || (_cache[10] = ($event) => _ctx.onMeridienChange($event.target.value)),
        disabled: _ctx.disabled
      }, {
        default: vue.withCtx(() => [(vue.openBlock(true), vue.createBlock(vue.Fragment, null, vue.renderList(_ctx.meridiens, (meridien) => {
          return vue.openBlock(), vue.createBlock("option", {
            value: meridien,
            key: meridien
          }, vue.toDisplayString(meridien), 9, ["value"]);
        }), 128))]),
        _: 1
      }, 8, ["select-class", "placeholder-class", "modelValue", "disabled"]) : vue.createCommentVNode("v-if", true)], 2), _ctx.$slots.default !== void 0 ? (vue.openBlock(), vue.createBlock("footer", {
        key: 0,
        class: _ctx.footerClasses
      }, [vue.renderSlot(_ctx.$slots, "default")], 2)) : vue.createCommentVNode("v-if", true)]),
      _: 1
    }, 8, ["disabled"])]),
    _: 2
  }, [!_ctx.inline ? {
    name: "trigger",
    fn: vue.withCtx(() => [vue.renderSlot(_ctx.$slots, "trigger", {}, () => [vue.createVNode(_component_o_input, vue.mergeProps({
      ref: "input"
    }, _ctx.inputBind, {
      autocomplete: "off",
      value: _ctx.formatValue(_ctx.computedValue),
      placeholder: _ctx.placeholder,
      size: _ctx.size,
      icon: _ctx.icon,
      "icon-pack": _ctx.iconPack,
      disabled: _ctx.disabled,
      readonly: !_ctx.editable,
      rounded: _ctx.rounded,
      "use-html5-validation": _ctx.useHtml5Validation,
      onKeyup: _cache[1] || (_cache[1] = vue.withKeys(($event) => _ctx.toggle(true), ["enter"])),
      onChange: _cache[2] || (_cache[2] = ($event) => _ctx.onChange($event.target.value)),
      onFocus: _ctx.handleOnFocus
    }), null, 16, ["value", "placeholder", "size", "icon", "icon-pack", "disabled", "readonly", "rounded", "use-html5-validation", "onFocus"])])])
  } : void 0]), 1040, ["position", "disabled", "inline", "append-to-body", "onActive-change"]) : vue.createVNode(_component_o_input, vue.mergeProps({
    key: 1,
    ref: "input"
  }, _ctx.inputBind, {
    type: "time",
    step: _ctx.nativeStep,
    autocomplete: "off",
    value: _ctx.formatHHMMSS(_ctx.computedValue),
    placeholder: _ctx.placeholder,
    size: _ctx.size,
    icon: _ctx.icon,
    "icon-pack": _ctx.iconPack,
    rounded: _ctx.rounded,
    max: _ctx.formatHHMMSS(_ctx.maxTime),
    min: _ctx.formatHHMMSS(_ctx.minTime),
    disabled: _ctx.disabled,
    readonly: false,
    "use-html5-validation": _ctx.useHtml5Validation,
    onChange: _cache[11] || (_cache[11] = ($event) => _ctx.onChange($event.target.value)),
    onFocus: _ctx.handleOnFocus,
    onBlur: _cache[12] || (_cache[12] = ($event) => _ctx.onBlur() && _ctx.checkHtml5Validity())
  }), null, 16, ["step", "value", "placeholder", "size", "icon", "icon-pack", "rounded", "max", "min", "disabled", "use-html5-validation", "onFocus"])], 2);
}
script$h.render = render$f;
script$h.__file = "src/components/timepicker/Timepicker.vue";
const AM = "AM";
const PM = "PM";
var script$g = vue.defineComponent({
  name: "ODatetimepicker",
  components: {
    [script$2$3.name]: script$2$3,
    [script$h.name]: script$h
  },
  configField: "datetimepicker",
  mixins: [FormElementMixin, BaseComponentMixin],
  inheritAttrs: false,
  emits: ["update:modelValue"],
  props: {
    modelValue: {
      type: Date
    },
    editable: {
      type: Boolean,
      default: false
    },
    placeholder: String,
    disabled: Boolean,
    iconRight: String,
    iconRightClickable: Boolean,
    inline: Boolean,
    openOnFocus: Boolean,
    position: String,
    mobileNative: {
      type: Boolean,
      default: true
    },
    minDatetime: Date,
    maxDatetime: Date,
    datetimeFormatter: {
      type: Function
    },
    datetimeParser: {
      type: Function
    },
    datetimeCreator: {
      type: Function,
      default: (date) => {
        const datetimeCreator = getValueByPath(getOptions$1(), "datetimepicker.datetimeCreator", void 0);
        if (typeof datetimeCreator === "function") {
          return datetimeCreator(date);
        } else {
          return date;
        }
      }
    },
    datepicker: Object,
    timepicker: Object,
    locale: {
      type: [String, Array],
      default: () => {
        return getValueByPath(getOptions$1(), "locale");
      }
    },
    appendToBody: Boolean,
    datepickerWrapperClass: [String, Function, Array],
    timepickerWrapperClass: [String, Function, Array]
  },
  data() {
    return {
      newValue: this.modelValue
    };
  },
  computed: {
    datepickerWrapperClasses() {
      return [
        this.computedClass("datepickerWrapperClass", "o-dtpck__date")
      ];
    },
    timepickerWrapperClasses() {
      return [
        this.computedClass("timepickerWrapperClass", "o-dtpck__time")
      ];
    },
    computedValue: {
      get() {
        return this.newValue;
      },
      set(value) {
        if (value) {
          let val = new Date(value.getTime());
          if (this.newValue) {
            if ((value.getDate() !== this.newValue.getDate() || value.getMonth() !== this.newValue.getMonth() || value.getFullYear() !== this.newValue.getFullYear()) && value.getHours() === 0 && value.getMinutes() === 0 && value.getSeconds() === 0) {
              val.setHours(this.newValue.getHours(), this.newValue.getMinutes(), this.newValue.getSeconds(), 0);
            }
          } else {
            val = this.datetimeCreator(value);
          }
          if (this.minDatetime && val < this.minDatetime) {
            val = this.minDatetime;
          } else if (this.maxDatetime && val > this.maxDatetime) {
            val = this.maxDatetime;
          }
          this.newValue = new Date(val.getTime());
        } else {
          this.newValue = value;
        }
        this.$emit("update:modelValue", this.newValue);
      }
    },
    localeOptions() {
      return new Intl.DateTimeFormat(this.locale, {
        year: "numeric",
        month: "numeric",
        day: "numeric",
        hour: "numeric",
        minute: "numeric",
        second: this.enableSeconds() ? "numeric" : void 0
      }).resolvedOptions();
    },
    dtf() {
      return new Intl.DateTimeFormat(this.locale, {
        year: this.localeOptions.year || "numeric",
        month: this.localeOptions.month || "numeric",
        day: this.localeOptions.day || "numeric",
        hour: this.localeOptions.hour || "numeric",
        minute: this.localeOptions.minute || "numeric",
        second: this.enableSeconds() ? this.localeOptions.second || "numeric" : void 0,
        hour12: !this.isHourFormat24()
      });
    },
    isMobileNative() {
      return this.mobileNative;
    },
    isMobile() {
      return this.isMobileNative && isMobile.any();
    },
    minDate() {
      if (!this.minDatetime) {
        return this.datepicker ? this.datepicker.minDate : null;
      }
      return new Date(this.minDatetime.getFullYear(), this.minDatetime.getMonth(), this.minDatetime.getDate(), 0, 0, 0, 0);
    },
    maxDate() {
      if (!this.maxDatetime) {
        return this.datepicker ? this.datepicker.maxDate : null;
      }
      return new Date(this.maxDatetime.getFullYear(), this.maxDatetime.getMonth(), this.maxDatetime.getDate(), 0, 0, 0, 0);
    },
    minTime() {
      if (!this.minDatetime || (this.newValue === null || typeof this.newValue === "undefined")) {
        return this.timepicker ? this.timepicker.minTime : null;
      }
      return this.minDatetime;
    },
    maxTime() {
      if (!this.maxDatetime || (this.newValue === null || typeof this.newValue === "undefined")) {
        return this.timepicker ? this.timepicker.maxTime : null;
      }
      return this.maxDatetime;
    },
    datepickerSize() {
      return this.datepicker && this.datepicker.size ? this.datepicker.size : this.size;
    },
    timepickerSize() {
      return this.timepicker && this.timepicker.size ? this.timepicker.size : this.size;
    },
    timepickerDisabled() {
      return this.timepicker && this.timepicker.disabled ? this.timepicker.disabled : this.disabled;
    }
  },
  watch: {
    modelValue(value) {
      this.newValue = value;
    }
  },
  methods: {
    enableSeconds() {
      if (this.$refs.timepicker) {
        return this.$refs.timepicker.enableSeconds;
      }
      return false;
    },
    isHourFormat24() {
      if (this.$refs.timepicker) {
        return this.$refs.timepicker.isHourFormat24;
      }
      return !this.localeOptions.hour12;
    },
    defaultDatetimeParser(date) {
      const datetimeParser = getValueByPath(getOptions$1(), "datetimepicker.datetimeParser", void 0);
      if (typeof this.datetimeParser === "function") {
        return this.datetimeParser(date);
      } else if (typeof datetimeParser === "function") {
        return datetimeParser(date);
      } else {
        if (this.dtf.formatToParts && typeof this.dtf.formatToParts === "function") {
          let dayPeriods = [AM, PM, AM.toLowerCase(), PM.toLowerCase()];
          if (this.$refs.timepicker) {
            dayPeriods.push(this.$refs.timepicker.amString);
            dayPeriods.push(this.$refs.timepicker.pmString);
          }
          const parts = this.dtf.formatToParts(new Date());
          const formatRegex = parts.map((part, idx) => {
            if (part.type === "literal") {
              if (idx + 1 < parts.length && parts[idx + 1].type === "hour") {
                return `[^\\d]+`;
              }
              return part.value.replace(/ /g, "\\s?");
            } else if (part.type === "dayPeriod") {
              return `((?!=<${part.type}>)(${dayPeriods.join("|")})?)`;
            }
            return `((?!=<${part.type}>)\\d+)`;
          }).join("");
          const datetimeGroups = matchWithGroups(formatRegex, date);
          if (datetimeGroups.year && datetimeGroups.year.length === 4 && datetimeGroups.month && datetimeGroups.month <= 12 && datetimeGroups.day && datetimeGroups.day <= 31 && datetimeGroups.hour && datetimeGroups.hour >= 0 && datetimeGroups.hour < 24 && datetimeGroups.minute && datetimeGroups.minute >= 0 && datetimeGroups.minute < 59) {
            const d = new Date(datetimeGroups.year, datetimeGroups.month - 1, datetimeGroups.day, datetimeGroups.hour, datetimeGroups.minute, datetimeGroups.second || 0);
            return d;
          }
        }
        return new Date(Date.parse(date));
      }
    },
    defaultDatetimeFormatter(date) {
      const datetimeFormatter = getValueByPath(getOptions$1(), "datetimepicker.datetimeFormatter", void 0);
      if (typeof this.datetimeFormatter === "function") {
        return this.datetimeFormatter(date);
      } else if (typeof datetimeFormatter === "function") {
        return datetimeFormatter(date);
      } else {
        return this.dtf.format(date);
      }
    },
    onChangeNativePicker(event) {
      const date = event.target.value;
      const s2 = date ? date.split(/\D/) : [];
      if (s2.length >= 5) {
        const year = parseInt(s2[0], 10);
        const month = parseInt(s2[1], 10) - 1;
        const day = parseInt(s2[2], 10);
        const hours = parseInt(s2[3], 10);
        const minutes = parseInt(s2[4], 10);
        this.computedValue = new Date(year, month, day, hours, minutes);
      } else {
        this.computedValue = null;
      }
    },
    formatNative(value) {
      const date = new Date(value);
      if (value && !isNaN(date.getTime())) {
        const year = date.getFullYear();
        const month = date.getMonth() + 1;
        const day = date.getDate();
        const hours = date.getHours();
        const minutes = date.getMinutes();
        const seconds = date.getSeconds();
        return year + "-" + ((month < 10 ? "0" : "") + month) + "-" + ((day < 10 ? "0" : "") + day) + "T" + ((hours < 10 ? "0" : "") + hours) + ":" + ((minutes < 10 ? "0" : "") + minutes) + ":" + ((seconds < 10 ? "0" : "") + seconds);
      }
      return "";
    },
    toggle() {
      this.$refs.datepicker.toggle();
    }
  },
  mounted() {
    if (!this.isMobile || this.inline) {
      if (this.newValue) {
        this.$refs.datepicker.$forceUpdate();
      }
    }
  }
});
function render$e(_ctx, _cache, $props, $setup, $data, $options) {
  const _component_o_timepicker = vue.resolveComponent("o-timepicker");
  const _component_o_datepicker = vue.resolveComponent("o-datepicker");
  const _component_o_input = vue.resolveComponent("o-input");
  return !_ctx.isMobile || _ctx.inline ? vue.createVNode(_component_o_datepicker, vue.mergeProps({
    key: 0,
    ref: "datepicker",
    modelValue: _ctx.computedValue,
    "onUpdate:modelValue": _cache[2] || (_cache[2] = ($event) => _ctx.computedValue = $event)
  }, _ctx.datepicker, {
    class: _ctx.datepickerWrapperClasses,
    rounded: _ctx.rounded,
    "open-on-focus": _ctx.openOnFocus,
    position: _ctx.position,
    inline: _ctx.inline,
    editable: _ctx.editable,
    expanded: _ctx.expanded,
    "close-on-click": false,
    "date-formatter": _ctx.defaultDatetimeFormatter,
    "date-parser": _ctx.defaultDatetimeParser,
    "min-date": _ctx.minDate,
    "max-date": _ctx.maxDate,
    icon: _ctx.icon,
    "icon-right": _ctx.iconRight,
    "icon-right-clickable": _ctx.iconRightClickable,
    "icon-pack": _ctx.iconPack,
    size: _ctx.datepickerSize,
    placeholder: _ctx.placeholder,
    range: false,
    disabled: _ctx.disabled,
    "mobile-native": _ctx.isMobileNative,
    locale: _ctx.locale,
    "append-to-body": _ctx.appendToBody,
    onFocus: _ctx.onFocus,
    onBlur: _ctx.onBlur,
    "onActive-change": _cache[3] || (_cache[3] = ($event) => _ctx.$emit("active-change", $event)),
    "onIcon-right-click": _cache[4] || (_cache[4] = ($event) => _ctx.$emit("icon-right-click")),
    "onChange-month": _cache[5] || (_cache[5] = ($event) => _ctx.$emit("change-month", $event)),
    "onChange-year": _cache[6] || (_cache[6] = ($event) => _ctx.$emit("change-year", $event))
  }), {
    default: vue.withCtx(() => [vue.createVNode("div", {
      class: _ctx.timepickerWrapperClasses
    }, [vue.createVNode(_component_o_timepicker, vue.mergeProps({
      ref: "timepicker"
    }, _ctx.timepicker, {
      modelValue: _ctx.computedValue,
      "onUpdate:modelValue": _cache[1] || (_cache[1] = ($event) => _ctx.computedValue = $event),
      inline: "",
      editable: _ctx.editable,
      "min-time": _ctx.minTime,
      "max-time": _ctx.maxTime,
      size: _ctx.timepickerSize,
      disabled: _ctx.timepickerDisabled,
      "mobile-native": _ctx.isMobileNative,
      locale: _ctx.locale
    }), null, 16, ["modelValue", "editable", "min-time", "max-time", "size", "disabled", "mobile-native", "locale"])], 2), _ctx.$slots.default !== void 0 ? vue.renderSlot(_ctx.$slots, "default", {
      key: 0
    }) : vue.createCommentVNode("v-if", true)]),
    _: 1
  }, 16, ["modelValue", "class", "rounded", "open-on-focus", "position", "inline", "editable", "expanded", "date-formatter", "date-parser", "min-date", "max-date", "icon", "icon-right", "icon-right-clickable", "icon-pack", "size", "placeholder", "disabled", "mobile-native", "locale", "append-to-body", "onFocus", "onBlur"]) : vue.createVNode(_component_o_input, vue.mergeProps({
    key: 1,
    ref: "input",
    type: "datetime-local",
    autocomplete: "off",
    value: _ctx.formatNative(_ctx.computedValue),
    placeholder: _ctx.placeholder,
    size: _ctx.size,
    icon: _ctx.icon,
    "icon-pack": _ctx.iconPack,
    rounded: _ctx.rounded,
    max: _ctx.formatNative(_ctx.maxDate),
    min: _ctx.formatNative(_ctx.minDate),
    disabled: _ctx.disabled,
    readonly: false
  }, _ctx.$attrs, {
    "use-html5-validation": _ctx.useHtml5Validation,
    onChange: _ctx.onChangeNativePicker,
    onFocus: _ctx.onFocus,
    onBlur: _ctx.onBlur
  }), null, 16, ["value", "placeholder", "size", "icon", "icon-pack", "rounded", "max", "min", "disabled", "use-html5-validation", "onChange", "onFocus", "onBlur"]);
}
script$g.render = render$e;
script$g.__file = "src/components/datetimepicker/Datetimepicker.vue";
var index$y = {
  install(app) {
    registerComponent(app, script$g);
  }
};
var index$5$1 = index$y;
var index$x = {
  install(app) {
    registerComponent(app, script$l);
    registerComponent(app, script$1$9);
  }
};
var index$6$1 = index$x;
var index$w = {
  install(app) {
    registerComponent(app, script$1$8);
  }
};
var index$7$1 = index$w;
var index$v = {
  install(app) {
    registerComponent(app, script$r);
  }
};
var index$8$1 = index$v;
var index$u = {
  install(app) {
    registerComponent(app, script$q);
  }
};
var index$9$1 = index$u;
var script$f = vue.defineComponent({
  name: "OInputitems",
  components: {
    [script$p.name]: script$p,
    [script$r.name]: script$r
  },
  mixins: [FormElementMixin, BaseComponentMixin],
  inheritAttrs: false,
  configField: "inputitems",
  emits: ["update:modelValue", "focus", "blur", "add", "remove", "typing", "infinite-scroll", "icon-right-click"],
  props: {
    modelValue: {
      type: Array,
      default: () => []
    },
    size: String,
    data: {
      type: Array,
      default: () => []
    },
    variant: String,
    maxitems: {
      type: [Number, String],
      required: false
    },
    hasCounter: {
      type: Boolean,
      default: () => {
        return getValueByPath(getOptions$1(), "inputitems.hasCounter", true);
      }
    },
    field: {
      type: String,
      default: "value"
    },
    autocomplete: Boolean,
    groupField: String,
    groupOptions: String,
    nativeAutocomplete: String,
    openOnFocus: Boolean,
    disabled: Boolean,
    closable: {
      type: Boolean,
      default: () => {
        return getValueByPath(getOptions$1(), "inputitems.closable", true);
      }
    },
    confirmKeys: {
      type: Array,
      default: () => {
        return getValueByPath(getOptions$1(), "inputitems.confirmKeys", [",", "Tab", "Enter"]);
      }
    },
    removeOnKeys: {
      type: Array,
      default: () => {
        return getValueByPath(getOptions$1(), "inputitems.removeOnKeys", ["Backspace"]);
      }
    },
    allowNew: Boolean,
    onPasteSeparators: {
      type: Array,
      default: () => {
        return getValueByPath(getOptions$1(), "inputitems.onPasteSeparators", [","]);
      }
    },
    beforeAdding: {
      type: Function,
      default: () => true
    },
    allowDuplicates: {
      type: Boolean,
      default: false
    },
    checkInfiniteScroll: {
      type: Boolean,
      default: false
    },
    createItem: {
      type: Function,
      default: (item) => item
    },
    closeIcon: {
      type: String,
      default: () => {
        return getValueByPath(getOptions$1(), "inputitems.closeIcon", "times");
      }
    },
    keepFirst: Boolean,
    ariaCloseLabel: String,
    appendToBody: Boolean,
    rootClass: [String, Array, Function],
    expandedClass: [String, Array, Function],
    variantClass: [String, Array, Function],
    closeClass: [String, Array, Function],
    itemClass: [String, Array, Function],
    counterClass: [String, Array, Function],
    autocompleteClasses: Object
  },
  data() {
    return {
      items: Array.isArray(this.modelValue) ? this.modelValue.slice(0) : this.modelValue || [],
      newItem: "",
      isComposing: false,
      $elementRef: "autocomplete"
    };
  },
  computed: {
    rootClasses() {
      return [
        this.computedClass("rootClass", "o-inputit"),
        { [this.computedClass("expandedClass", "o-inputit--expanded")]: this.expanded }
      ];
    },
    containerClasses() {
      return [
        this.computedClass("containerClass", "o-inputit__container"),
        { [this.computedClass("sizeClass", "o-inputit__container--", this.size)]: this.size }
      ];
    },
    itemClasses() {
      return [
        this.computedClass("itemClass", "o-inputit__item"),
        { [this.computedClass("variantClass", "o-inputit__item--", this.variant)]: this.variant }
      ];
    },
    closeClasses() {
      return [
        this.computedClass("closeClass", "o-inputit__item__close")
      ];
    },
    counterClasses() {
      return [
        this.computedClass("counterClass", "o-inputit__counter")
      ];
    },
    autocompleteBind() {
      return __spreadValues(__spreadProps(__spreadValues({}, this.$attrs), {
        "root-class": this.computedClass("autocompleteClasses.rootClass", "o-inputit__autocomplete"),
        "input-classes": {
          "input-class": this.computedClass("autocompleteClasses.inputClasses.inputClass", "o-inputit__input")
        }
      }), this.autocompleteClasses);
    },
    valueLength() {
      return this.newItem.trim().length;
    },
    hasDefaultSlot() {
      return !!this.$slots.default;
    },
    hasEmptySlot() {
      return !!this.$slots.empty;
    },
    hasHeaderSlot() {
      return !!this.$slots.header;
    },
    hasFooterSlot() {
      return !!this.$slots.footer;
    },
    hasInput() {
      return this.maxitems == null || this.itemsLength < this.maxitems;
    },
    itemsLength() {
      return this.items.length;
    },
    separatorsAsRegExp() {
      const sep = this.onPasteSeparators;
      return sep.length ? new RegExp(sep.map((s2) => {
        return s2 ? s2.replace(/[-[\]{}()*+?.,\\^$|#\s]/g, "\\$&") : null;
      }).join("|"), "g") : null;
    }
  },
  watch: {
    modelValue(value) {
      this.items = Array.isArray(value) ? value.slice(0) : value || [];
    },
    hasInput() {
      if (!this.hasInput)
        this.onBlur();
    }
  },
  methods: {
    addItem(item) {
      const itemToAdd = item || this.newItem.trim();
      if (itemToAdd) {
        if (!this.autocomplete) {
          const reg = this.separatorsAsRegExp;
          if (reg && itemToAdd.match(reg)) {
            itemToAdd.split(reg).map((t) => t.trim()).filter((t) => t.length !== 0).map(this.addItem);
            return;
          }
        }
        const add = !this.allowDuplicates ? this.items.indexOf(this.createItem(itemToAdd)) === -1 : true;
        if (add && this.beforeAdding(itemToAdd)) {
          this.items.push(this.createItem(itemToAdd));
          this.$emit("update:modelValue", this.items);
          this.$emit("add", itemToAdd);
        }
      }
      requestAnimationFrame(() => {
        this.newItem = "";
        this.$emit("typing", "");
      });
    },
    getNormalizedItemText(item) {
      if (typeof item === "object") {
        item = getValueByPath(item, this.field);
      }
      return `${item}`;
    },
    customOnBlur(event) {
      if (!this.autocomplete)
        this.addItem();
      this.onBlur(event);
    },
    onSelect(option) {
      if (!option)
        return;
      this.addItem(option);
      this.$nextTick(() => {
        this.newItem = "";
      });
    },
    removeItem(index2, event) {
      const item = this.items.splice(index2, 1)[0];
      this.$emit("update:modelValue", this.items);
      this.$emit("remove", item);
      if (event)
        event.stopPropagation();
      if (this.openOnFocus && this.$refs.autocomplete) {
        this.$refs.autocomplete.focus();
      }
      return item;
    },
    removeLastItem() {
      if (this.itemsLength > 0) {
        this.removeItem(this.itemsLength - 1);
      }
    },
    keydown(event) {
      const { key } = event;
      if (this.removeOnKeys.indexOf(key) !== -1 && !this.newItem.length) {
        this.removeLastItem();
      }
      if (this.autocomplete && !this.allowNew)
        return;
      if (this.confirmKeys.indexOf(key) >= 0) {
        if (key !== "Tab")
          event.preventDefault();
        if (key === "Enter" && this.isComposing)
          return;
        this.addItem();
      }
    },
    onTyping(event) {
      this.$emit("typing", event.trim());
    }
  }
});
function render$d(_ctx, _cache, $props, $setup, $data, $options) {
  const _component_o_icon = vue.resolveComponent("o-icon");
  const _component_o_autocomplete = vue.resolveComponent("o-autocomplete");
  return vue.openBlock(), vue.createBlock("div", {
    class: _ctx.rootClasses
  }, [vue.createVNode("div", {
    class: _ctx.containerClasses,
    onClick: _cache[6] || (_cache[6] = ($event) => _ctx.hasInput && _ctx.focus($event))
  }, [vue.renderSlot(_ctx.$slots, "selected", {
    items: _ctx.items
  }, () => [(vue.openBlock(true), vue.createBlock(vue.Fragment, null, vue.renderList(_ctx.items, (item, index2) => {
    return vue.openBlock(), vue.createBlock("span", {
      key: _ctx.getNormalizedItemText(item) + index2,
      class: _ctx.itemClasses
    }, [vue.createVNode("span", null, vue.toDisplayString(_ctx.getNormalizedItemText(item)), 1), _ctx.closable ? vue.createVNode(_component_o_icon, {
      key: 0,
      class: _ctx.closeClasses,
      clickable: "",
      both: "",
      icon: _ctx.closeIcon,
      onClick: ($event) => _ctx.removeItem(index2, $event),
      "aria-label": _ctx.ariaCloseLabel
    }, null, 8, ["class", "icon", "onClick", "aria-label"]) : vue.createCommentVNode("v-if", true)], 2);
  }), 128))]), _ctx.hasInput ? vue.createVNode(_component_o_autocomplete, vue.mergeProps({
    key: 0,
    ref: "autocomplete",
    modelValue: _ctx.newItem,
    "onUpdate:modelValue": _cache[1] || (_cache[1] = ($event) => _ctx.newItem = $event)
  }, _ctx.autocompleteBind, {
    data: _ctx.data,
    field: _ctx.field,
    icon: _ctx.icon,
    "icon-pack": _ctx.iconPack,
    maxlength: _ctx.maxlength,
    "has-counter": false,
    size: _ctx.size,
    disabled: _ctx.disabled,
    autocomplete: _ctx.nativeAutocomplete,
    "open-on-focus": _ctx.openOnFocus,
    "keep-first": _ctx.keepFirst,
    "keep-open": _ctx.openOnFocus,
    "group-field": _ctx.groupField,
    "group-options": _ctx.groupOptions,
    "use-html5-validation": _ctx.useHtml5Validation,
    "check-infinite-scroll": _ctx.checkInfiniteScroll,
    "append-to-body": _ctx.appendToBody,
    "confirm-keys": _ctx.confirmKeys,
    onTyping: _ctx.onTyping,
    onFocus: _ctx.onFocus,
    onBlur: _ctx.customOnBlur,
    onKeydown: _ctx.keydown,
    onCompositionstart: _cache[2] || (_cache[2] = ($event) => _ctx.isComposing = true),
    onCompositionend: _cache[3] || (_cache[3] = ($event) => _ctx.isComposing = false),
    onSelect: _ctx.onSelect,
    "onInfinite-scroll": _cache[4] || (_cache[4] = ($event) => _ctx.$emit("infinite-scroll", $event)),
    "onIcon-right-click": _cache[5] || (_cache[5] = ($event) => _ctx.$emit("icon-right-click", $event))
  }), vue.createSlots({
    _: 2
  }, [_ctx.hasHeaderSlot ? {
    name: "header",
    fn: vue.withCtx(() => [vue.renderSlot(_ctx.$slots, "header")])
  } : void 0, _ctx.hasDefaultSlot ? {
    name: "default",
    fn: vue.withCtx((props) => [vue.renderSlot(_ctx.$slots, "default", {
      option: props.option,
      index: props.index
    })])
  } : void 0, _ctx.hasEmptySlot ? {
    name: "empty",
    fn: vue.withCtx(() => [vue.renderSlot(_ctx.$slots, "empty")])
  } : void 0, _ctx.hasFooterSlot ? {
    name: "footer",
    fn: vue.withCtx(() => [vue.renderSlot(_ctx.$slots, "footer")])
  } : void 0]), 1040, ["modelValue", "data", "field", "icon", "icon-pack", "maxlength", "size", "disabled", "autocomplete", "open-on-focus", "keep-first", "keep-open", "group-field", "group-options", "use-html5-validation", "check-infinite-scroll", "append-to-body", "confirm-keys", "onTyping", "onFocus", "onBlur", "onKeydown", "onSelect"]) : vue.createCommentVNode("v-if", true)], 2), _ctx.hasCounter && (_ctx.maxitems || _ctx.maxlength) ? (vue.openBlock(), vue.createBlock("small", {
    key: 0,
    class: _ctx.counterClasses
  }, [_ctx.maxlength && _ctx.valueLength > 0 ? (vue.openBlock(), vue.createBlock(vue.Fragment, {
    key: 0
  }, [vue.createTextVNode(vue.toDisplayString(_ctx.valueLength) + " / " + vue.toDisplayString(_ctx.maxlength), 1)], 64)) : _ctx.maxitems ? (vue.openBlock(), vue.createBlock(vue.Fragment, {
    key: 1
  }, [vue.createTextVNode(vue.toDisplayString(_ctx.itemsLength) + " / " + vue.toDisplayString(_ctx.maxitems), 1)], 64)) : vue.createCommentVNode("v-if", true)], 2)) : vue.createCommentVNode("v-if", true)], 2);
}
script$f.render = render$d;
script$f.__file = "src/components/inputitems/Inputitems.vue";
var index$t = {
  install(Vue) {
    registerComponent(Vue, script$f);
  }
};
var index$a$1 = index$t;
const isSSR = typeof window === "undefined";
const HTMLElement = isSSR ? Object : window.HTMLElement;
const File = isSSR ? Object : window.File;
var script$e = vue.defineComponent({
  name: "OLoading",
  components: {
    [script$r.name]: script$r
  },
  mixins: [BaseComponentMixin],
  configField: "loading",
  emits: ["update:active", "close", "update:full-page"],
  props: {
    active: Boolean,
    programmatic: Boolean,
    container: [Object, Function, HTMLElement],
    fullPage: {
      type: Boolean,
      default: true
    },
    animation: {
      type: String,
      default: () => {
        return getValueByPath(getOptions$1(), "loading.animation", "fade");
      }
    },
    canCancel: {
      type: Boolean,
      default: false
    },
    onCancel: {
      type: Function,
      default: () => {
      }
    },
    icon: {
      type: String,
      default: () => {
        return getValueByPath(getOptions$1(), "loading.icon", "loading");
      }
    },
    iconSpin: {
      type: Boolean,
      default: true
    },
    iconSize: {
      type: String,
      default: "medium"
    },
    rootClass: [String, Function, Array],
    overlayClass: [String, Function, Array],
    iconClass: [String, Function, Array]
  },
  data() {
    return {
      isActive: this.active || false,
      displayInFullPage: this.fullPage
    };
  },
  watch: {
    active(value) {
      this.isActive = value;
    },
    fullPage(value) {
      this.displayInFullPage = value;
    }
  },
  computed: {
    rootClasses() {
      return [
        this.computedClass("rootClass", "o-load"),
        { [this.computedClass("fullPageClass", "o-load--fullpage")]: this.displayInFullPage }
      ];
    },
    overlayClasses() {
      return [
        this.computedClass("overlayClass", "o-load__overlay")
      ];
    },
    iconClasses() {
      return [
        this.computedClass("iconClass", "o-load__icon")
      ];
    }
  },
  methods: {
    cancel() {
      if (!this.canCancel || !this.isActive)
        return;
      this.close();
    },
    close() {
      this.onCancel.apply(null, arguments);
      this.$emit("close");
      this.$emit("update:active", false);
      if (this.programmatic) {
        this.isActive = false;
        setTimeout(() => {
          this.$destroy();
          removeElement(this.$el);
        }, 150);
      }
    },
    keyPress({ key }) {
      if (key === "Escape" || key === "Esc")
        this.cancel();
    }
  },
  created() {
    if (typeof window !== "undefined") {
      document.addEventListener("keyup", this.keyPress);
    }
  },
  mounted() {
    if (this.programmatic) {
      if (!this.container) {
        document.body.appendChild(this.$el);
      } else {
        this.displayInFullPage = false;
        this.$emit("update:full-page", false);
        this.container.appendChild(this.$el);
      }
      this.isActive = true;
    }
  },
  beforeUnmount() {
    if (typeof window !== "undefined") {
      document.removeEventListener("keyup", this.keyPress);
    }
  }
});
function render$c(_ctx, _cache, $props, $setup, $data, $options) {
  const _component_o_icon = vue.resolveComponent("o-icon");
  return vue.openBlock(), vue.createBlock(vue.Transition, {
    name: _ctx.animation
  }, {
    default: vue.withCtx(() => [_ctx.isActive ? (vue.openBlock(), vue.createBlock("div", {
      key: 0,
      class: _ctx.rootClasses
    }, [vue.createVNode("div", {
      class: _ctx.overlayClasses,
      onClick: _cache[1] || (_cache[1] = (...args) => _ctx.cancel(...args))
    }, null, 2), vue.renderSlot(_ctx.$slots, "default", {}, () => [vue.createVNode(_component_o_icon, {
      icon: _ctx.icon,
      spin: _ctx.iconSpin,
      size: _ctx.iconSize,
      class: _ctx.iconClasses,
      both: ""
    }, null, 8, ["icon", "spin", "size", "class"])])], 2)) : vue.createCommentVNode("v-if", true)]),
    _: 1
  }, 8, ["name"]);
}
script$e.render = render$c;
script$e.__file = "src/components/loading/Loading.vue";
let localVueInstance$2;
const LoadingProgrammatic = {
  open(params) {
    const defaultParam = {
      programmatic: true
    };
    const propsData = merge(defaultParam, params);
    const app = localVueInstance$2 || VueInstance;
    const vnode = vue.createVNode(script$e, propsData);
    vnode.appContext = app._context;
    return vue.render(vnode, document.createElement("div"));
  }
};
var index$s = {
  install(app) {
    localVueInstance$2 = app;
    registerComponent(app, script$e);
    registerComponentProgrammatic(app, "loading", LoadingProgrammatic);
  }
};
var index$b$1 = index$s;
var script$d = vue.defineComponent({
  name: "OModal",
  components: {
    [script$r.name]: script$r
  },
  configField: "modal",
  directives: {
    trapFocus: directive
  },
  mixins: [BaseComponentMixin, MatchMediaMixin],
  emits: ["update:active", "close"],
  props: {
    active: Boolean,
    component: [Object, Function],
    content: String,
    closeButtonContent: {
      type: String,
      default: "\u2715"
    },
    programmatic: Boolean,
    props: Object,
    events: Object,
    width: {
      type: [String, Number],
      default: () => {
        return getValueByPath(getOptions$1(), "modal.width", 960);
      }
    },
    custom: Boolean,
    animation: {
      type: String,
      default: () => {
        return getValueByPath(getOptions$1(), "modal.animation", "zoom-out");
      }
    },
    canCancel: {
      type: [Array, Boolean],
      default: () => {
        return getValueByPath(getOptions$1(), "modal.canCancel", ["escape", "x", "outside", "button"]);
      }
    },
    onCancel: {
      type: Function,
      default: () => {
      }
    },
    onClose: {
      type: Function,
      default: () => {
      }
    },
    scroll: {
      type: String,
      default: () => {
        return getValueByPath(getOptions$1(), "modal.scroll", "keep");
      }
    },
    fullScreen: Boolean,
    trapFocus: {
      type: Boolean,
      default: () => {
        return getValueByPath(getOptions$1(), "modal.trapFocus", true);
      }
    },
    ariaRole: {
      type: String,
      validator: (value) => {
        return ["dialog", "alertdialog"].indexOf(value) >= 0;
      }
    },
    ariaModal: Boolean,
    ariaLabel: String,
    destroyOnHide: {
      type: Boolean,
      default: () => {
        return getValueByPath(getOptions$1(), "modal.destroyOnHide", true);
      }
    },
    autoFocus: {
      type: Boolean,
      default: () => {
        return getValueByPath(getOptions$1(), "modal.autoFocus", true);
      }
    },
    closeIcon: {
      type: String,
      default: () => {
        return getValueByPath(getOptions$1(), "modal.closeIcon", "times");
      }
    },
    closeIconSize: {
      type: String,
      default: "medium"
    },
    rootClass: [String, Function, Array],
    overlayClass: [String, Function, Array],
    contentClass: [String, Function, Array],
    closeClass: [String, Function, Array],
    fullScreenClass: [String, Function, Array],
    mobileClass: [String, Function, Array]
  },
  data() {
    return {
      isActive: this.active || false,
      savedScrollTop: null,
      newWidth: toCssDimension(this.width),
      animating: !this.active,
      destroyed: !this.active
    };
  },
  computed: {
    rootClasses() {
      return [
        this.computedClass("rootClass", "o-modal"),
        { [this.computedClass("mobileClass", "o-modal--mobile")]: this.isMatchMedia }
      ];
    },
    overlayClasses() {
      return [
        this.computedClass("overlayClass", "o-modal__overlay")
      ];
    },
    contentClasses() {
      return [
        { [this.computedClass("contentClass", "o-modal__content")]: !this.custom },
        { [this.computedClass("fullScreenClass", "o-modal__content--fullscreen")]: this.fullScreen }
      ];
    },
    closeClasses() {
      return [
        this.computedClass("closeClass", "o-modal__close")
      ];
    },
    cancelOptions() {
      return typeof this.canCancel === "boolean" ? this.canCancel ? getValueByPath(getOptions$1(), "modal.canCancel", ["escape", "x", "outside", "button"]) : [] : this.canCancel;
    },
    showX() {
      return this.cancelOptions.indexOf("x") >= 0;
    },
    customStyle() {
      if (!this.fullScreen) {
        return { maxWidth: this.newWidth };
      }
      return null;
    }
  },
  watch: {
    active(value) {
      this.isActive = value;
    },
    isActive(value) {
      if (value)
        this.destroyed = false;
      this.handleScroll();
      this.$nextTick(() => {
        if (value && this.$el && this.$el.focus && this.autoFocus) {
          this.$el.focus();
        }
      });
    }
  },
  methods: {
    handleScroll() {
      if (typeof window === "undefined")
        return;
      if (this.scroll === "clip") {
        if (this.isActive) {
          document.documentElement.classList.add("o-clipped");
        } else {
          document.documentElement.classList.remove("o-clipped");
        }
        return;
      }
      this.savedScrollTop = !this.savedScrollTop ? document.documentElement.scrollTop : this.savedScrollTop;
      if (this.isActive) {
        document.body.classList.add("o-noscroll");
      } else {
        document.body.classList.remove("o-noscroll");
      }
      if (this.isActive) {
        document.body.style.top = `-${this.savedScrollTop}px`;
        return;
      }
      document.documentElement.scrollTop = this.savedScrollTop;
      document.body.style.top = null;
      this.savedScrollTop = null;
    },
    cancel(method) {
      if (this.cancelOptions.indexOf(method) < 0)
        return;
      this.onCancel.apply(null, arguments);
      this.close();
    },
    close() {
      this.isActive = false;
      if (this.destroyOnHide) {
        this.destroyed = true;
      }
      this.$emit("close");
      this.$emit("update:active", false);
      this.onClose.apply(null, arguments);
      if (this.programmatic) {
        window.requestAnimationFrame(() => {
          removeElement(this.$el);
        });
      }
    },
    keyPress({ key }) {
      if (this.isActive && (key === "Escape" || key === "Esc"))
        this.cancel("escape");
    },
    afterEnter() {
      this.animating = false;
    },
    beforeLeave() {
      this.animating = true;
    }
  },
  created() {
    if (typeof window !== "undefined") {
      document.addEventListener("keyup", this.keyPress);
    }
  },
  mounted() {
    if (this.programmatic) {
      document.body.appendChild(this.$el);
      this.isActive = true;
    } else if (this.isActive)
      this.handleScroll();
  },
  beforeUnmount() {
    if (typeof window !== "undefined") {
      document.removeEventListener("keyup", this.keyPress);
      document.documentElement.classList.remove("o-clipped");
      const savedScrollTop = !this.savedScrollTop ? document.documentElement.scrollTop : this.savedScrollTop;
      document.body.classList.remove("o-noscroll");
      document.documentElement.scrollTop = savedScrollTop;
      document.body.style.top = null;
    }
  }
});
const _hoisted_1$5 = {
  key: 1
};
function render$b(_ctx, _cache, $props, $setup, $data, $options) {
  const _component_o_icon = vue.resolveComponent("o-icon");
  const _directive_trap_focus = vue.resolveDirective("trap-focus");
  return vue.openBlock(), vue.createBlock(vue.Transition, {
    name: _ctx.animation,
    "onAfter-enter": _ctx.afterEnter,
    "onBefore-leave": _ctx.beforeLeave
  }, {
    default: vue.withCtx(() => [!_ctx.destroyed ? vue.withDirectives((vue.openBlock(), vue.createBlock("div", {
      key: 0,
      class: _ctx.rootClasses,
      tabindex: "-1",
      role: _ctx.ariaRole,
      "aria-label": _ctx.ariaLabel,
      "aria-modal": _ctx.ariaModal
    }, [vue.createVNode("div", {
      class: _ctx.overlayClasses,
      onClick: _cache[1] || (_cache[1] = ($event) => _ctx.cancel("outside"))
    }, null, 2), vue.createVNode("div", {
      class: _ctx.contentClasses,
      style: _ctx.customStyle
    }, [_ctx.component ? (vue.openBlock(), vue.createBlock(vue.resolveDynamicComponent(_ctx.component), vue.mergeProps({
      key: 0
    }, _ctx.props, vue.toHandlers(_ctx.events || {}), {
      onClose: _ctx.close
    }), null, 16, ["onClose"])) : _ctx.content ? (vue.openBlock(), vue.createBlock("div", _hoisted_1$5, vue.toDisplayString(_ctx.content), 1)) : vue.renderSlot(_ctx.$slots, "default", {
      key: 2
    }), _ctx.showX ? vue.withDirectives(vue.createVNode(_component_o_icon, {
      key: 3,
      clickable: "",
      both: "",
      class: _ctx.closeClasses,
      icon: _ctx.closeIcon,
      size: _ctx.closeIconSize,
      onClick: _cache[2] || (_cache[2] = ($event) => _ctx.cancel("x"))
    }, null, 8, ["class", "icon", "size"]), [[vue.vShow, !_ctx.animating]]) : vue.createCommentVNode("v-if", true)], 6)], 10, ["role", "aria-label", "aria-modal"])), [[vue.vShow, _ctx.isActive], [_directive_trap_focus, _ctx.trapFocus]]) : vue.createCommentVNode("v-if", true)]),
    _: 1
  }, 8, ["name", "onAfter-enter", "onBefore-leave"]);
}
script$d.render = render$b;
script$d.__file = "src/components/modal/Modal.vue";
let localVueInstance$1;
const ModalProgrammatic = {
  open(params) {
    let newParams;
    if (typeof params === "string") {
      newParams = {
        content: params
      };
    } else {
      newParams = params;
    }
    const defaultParam = {
      programmatic: true
    };
    if (newParams.parent) {
      delete newParams.parent;
    }
    if (Array.isArray(newParams.content)) {
      delete newParams.content;
    }
    const propsData = merge(defaultParam, newParams);
    const app = localVueInstance$1 || VueInstance;
    const vnode = vue.createVNode(script$d, propsData);
    vnode.appContext = app._context;
    return vue.render(vnode, document.createElement("div"));
  }
};
var index$r = {
  install(app) {
    localVueInstance$1 = app;
    registerComponent(app, script$d);
    registerComponentProgrammatic(app, "modal", ModalProgrammatic);
  }
};
var index$c$1 = index$r;
var MessageMixin = {
  components: {
    [script$r.name]: script$r
  },
  props: {
    active: {
      type: Boolean,
      default: true
    },
    closable: {
      type: Boolean,
      default: false
    },
    message: String,
    type: String,
    hasIcon: Boolean,
    icon: String,
    iconPack: String,
    iconSize: {
      type: String,
      default: "large"
    },
    autoClose: {
      type: Boolean,
      default: false
    },
    duration: {
      type: Number,
      default: 2e3
    }
  },
  data() {
    return {
      isActive: this.active
    };
  },
  watch: {
    active(value) {
      this.isActive = value;
    },
    isActive(value) {
      if (value) {
        this.setAutoClose();
      } else {
        if (this.timer) {
          clearTimeout(this.timer);
        }
      }
    }
  },
  computed: {
    computedIcon() {
      if (this.icon) {
        return this.icon;
      }
      switch (this.type) {
        case "info":
          return "information";
        case "success":
          return "check-circle";
        case "warning":
          return "alert";
        case "danger":
          return "alert-circle";
        default:
          return null;
      }
    }
  },
  methods: {
    close() {
      this.isActive = false;
      this.$emit("close");
      this.$emit("update:active", false);
    },
    setAutoClose() {
      if (this.autoClose) {
        this.timer = setTimeout(() => {
          if (this.isActive) {
            this.close();
          }
        }, this.duration);
      }
    }
  },
  mounted() {
    this.setAutoClose();
  }
};
var script$c = {
  name: "ONotification",
  configField: "notification",
  mixins: [BaseComponentMixin, MessageMixin],
  props: {
    position: String,
    variant: [String, Object],
    ariaCloseLabel: String,
    animation: {
      type: String,
      default: "fade"
    },
    component: [Object, Function],
    props: Object,
    events: {
      type: Object,
      default: {}
    },
    closeIcon: {
      type: String,
      default: () => {
        return getValueByPath(getOptions$1(), "notification.closeIcon", "times");
      }
    },
    rootClass: [String, Function, Array],
    closeClass: [String, Function, Array],
    contentClass: [String, Function, Array],
    iconClass: [String, Function, Array],
    positionClass: [String, Function, Array],
    variantClass: [String, Function, Array],
    wrapperClass: [String, Function, Array]
  },
  computed: {
    rootClasses() {
      return [this.computedClass("rootClass", "o-notification"), {
        [this.computedClass("variantClass", "o-notification--", this.variant)]: this.variant
      }, {
        [this.computedClass("positionClass", "o-notification--", this.position)]: this.position
      }];
    },
    wrapperClasses() {
      return [this.computedClass("wrapperClass", "o-notification__wrapper")];
    },
    iconClasses() {
      return [this.computedClass("iconClass", "o-notification__icon")];
    },
    contentClasses() {
      return [this.computedClass("contentClass", "o-notification__content")];
    },
    closeClasses() {
      return [this.computedClass("closeClass", "o-notification__close")];
    }
  }
};
function render$a(_ctx, _cache, $props, $setup, $data, $options) {
  const _component_o_icon = vue.resolveComponent("o-icon");
  return vue.openBlock(), vue.createBlock(vue.Transition, {
    name: $props.animation
  }, {
    default: vue.withCtx(() => [vue.withDirectives(vue.createVNode("article", {
      class: $options.rootClasses
    }, [_ctx.closable ? (vue.openBlock(), vue.createBlock("button", {
      key: 0,
      class: $options.closeClasses,
      type: "button",
      onClick: _cache[1] || (_cache[1] = (...args) => _ctx.close(...args)),
      "aria-label": $props.ariaCloseLabel
    }, [vue.createVNode(_component_o_icon, {
      clickable: "",
      icon: $props.closeIcon,
      size: "small"
    }, null, 8, ["icon"])], 10, ["aria-label"])) : vue.createCommentVNode("v-if", true), $props.component ? (vue.openBlock(), vue.createBlock(vue.resolveDynamicComponent($props.component), vue.mergeProps({
      key: 1
    }, $props.props, vue.toHandlers($props.events), {
      onClose: _ctx.close
    }), null, 16, ["onClose"])) : vue.createCommentVNode("v-if", true), _ctx.$slots.default || _ctx.message ? (vue.openBlock(), vue.createBlock("div", {
      key: 2,
      class: $options.wrapperClasses
    }, [_ctx.computedIcon ? vue.createVNode(_component_o_icon, {
      key: 0,
      icon: _ctx.computedIcon,
      pack: _ctx.iconPack,
      class: $options.iconClasses,
      both: "",
      size: _ctx.iconSize,
      "aria-hidden": ""
    }, null, 8, ["icon", "pack", "class", "size"]) : vue.createCommentVNode("v-if", true), vue.createVNode("div", {
      class: $options.contentClasses
    }, [_ctx.message ? (vue.openBlock(), vue.createBlock("span", {
      key: 0,
      innerHTML: _ctx.message
    }, null, 8, ["innerHTML"])) : vue.renderSlot(_ctx.$slots, "default", {
      key: 1,
      closeNotification: _ctx.close
    })], 2)], 2)) : vue.createCommentVNode("v-if", true)], 2), [[vue.vShow, _ctx.isActive]])]),
    _: 1
  }, 8, ["name"]);
}
script$c.render = render$a;
script$c.__file = "src/components/notification/Notification.vue";
var NoticeMixin = {
  props: {
    type: {
      type: String
    },
    message: [String, Array],
    duration: {
      type: Number
    },
    queue: {
      type: Boolean,
      default: void 0
    },
    indefinite: {
      type: Boolean,
      default: false
    },
    position: {
      type: String,
      default: "top",
      validator(value) {
        return [
          "top-right",
          "top",
          "top-left",
          "bottom-right",
          "bottom",
          "bottom-left"
        ].indexOf(value) > -1;
      }
    },
    container: String,
    onClose: {
      type: Function,
      default: () => {
      }
    }
  },
  data() {
    return {
      isActive: false,
      parentTop: null,
      parentBottom: null,
      newContainer: this.container || getValueByPath(getOptions$1(), "notification.defaultContainerElement", void 0)
    };
  },
  computed: {
    correctParent() {
      switch (this.position) {
        case "top-right":
        case "top":
        case "top-left":
          return this.parentTop;
        case "bottom-right":
        case "bottom":
        case "bottom-left":
          return this.parentBottom;
      }
    },
    transition() {
      switch (this.position) {
        case "top-right":
        case "top":
        case "top-left":
          return {
            enter: "fadeInDown",
            leave: "fadeOut"
          };
        case "bottom-right":
        case "bottom":
        case "bottom-left":
          return {
            enter: "fadeInUp",
            leave: "fadeOut"
          };
      }
    }
  },
  methods: {
    shouldQueue() {
      const queue = this.queue !== void 0 ? this.queue : getValueByPath(getOptions$1(), "notification.defaultNoticeQueue", void 0);
      if (!queue)
        return false;
      return this.parentTop.childElementCount > 0 || this.parentBottom.childElementCount > 0;
    },
    close() {
      clearTimeout(this.timer);
      this.isActive = false;
      this.$emit("close");
      this.onClose.apply(null, arguments);
      setTimeout(() => {
        removeElement(this.$el);
      }, 150);
    },
    showNotice() {
      if (this.shouldQueue()) {
        setTimeout(() => this.showNotice(), 250);
        return;
      }
      this.correctParent.insertAdjacentElement("afterbegin", this.$el);
      this.isActive = true;
      if (!this.indefinite) {
        this.timer = setTimeout(() => this.timeoutCallback(), this.newDuration);
      }
    },
    setupContainer() {
      this.parentTop = document.querySelector((this.newContainer ? this.newContainer : "body") + `>.${this.rootClasses().join(".")}.${this.positionClasses("top").join(".")}`);
      this.parentBottom = document.querySelector((this.newContainer ? this.newContainer : "body") + `>.${this.rootClasses().join(".")}.${this.positionClasses("bottom").join(".")}`);
      if (this.parentTop && this.parentBottom)
        return;
      if (!this.parentTop) {
        this.parentTop = document.createElement("div");
        this.parentTop.className = `${this.rootClasses().join(" ")} ${this.positionClasses("top").join(" ")}`;
      }
      if (!this.parentBottom) {
        this.parentBottom = document.createElement("div");
        this.parentBottom.className = `${this.rootClasses().join(" ")} ${this.positionClasses("bottom").join(" ")}`;
      }
      const container = document.querySelector(this.newContainer) || document.body;
      container.appendChild(this.parentTop);
      container.appendChild(this.parentBottom);
      if (this.newContainer) {
        this.parentTop.classList.add("has-custom-container");
        this.parentBottom.classList.add("has-custom-container");
      }
    },
    timeoutCallback() {
      return this.close();
    }
  },
  beforeMount() {
    this.setupContainer();
  },
  mounted() {
    this.showNotice();
  }
};
var script$1$6 = {
  name: "ONotificationNotice",
  configField: "notice",
  mixins: [BaseComponentMixin, NoticeMixin],
  emits: ["update:active", "close"],
  props: {
    propsNotification: Object
  },
  data() {
    return {
      newDuration: this.duration || getValueByPath(getOptions$1(), "notification.duration", 1e3)
    };
  },
  methods: {
    rootClasses() {
      return [this.computedClass("rootClass", "o-notices")];
    },
    positionClasses(position) {
      return [this.computedClass("positionClass", "o-notices--", position)];
    },
    timeoutCallback() {
      return this.$refs.notification.close();
    }
  }
};
function render$1$4(_ctx, _cache, $props, $setup, $data, $options) {
  const _component_o_notification = vue.resolveComponent("o-notification");
  return vue.openBlock(), vue.createBlock(_component_o_notification, vue.mergeProps($props.propsNotification, {
    ref: "notification",
    onClose: _ctx.close
  }), {
    default: vue.withCtx(() => [vue.renderSlot(_ctx.$slots, "default")]),
    _: 3
  }, 16, ["onClose"]);
}
script$1$6.render = render$1$4;
script$1$6.__file = "src/components/notification/NotificationNotice.vue";
let localVueInstance;
const NotificationProgrammatic = {
  open(params) {
    let newParams;
    if (typeof params === "string") {
      newParams = {
        message: params
      };
    } else {
      newParams = params;
    }
    const defaultParam = {
      programmatic: true,
      position: getValueByPath(getOptions$1(), "notification.position", "top-right"),
      closable: params.closable || getValueByPath(getOptions$1(), "notification.closable", false)
    };
    if (newParams.parent) {
      delete params.parent;
    }
    if (Array.isArray(newParams.message)) {
      delete newParams.message;
    }
    newParams.active = true;
    const propsData = merge(defaultParam, newParams);
    const app = localVueInstance || VueInstance;
    propsData.propsNotification = Object.assign({}, propsData);
    propsData.propsNotification.isActive = true;
    const defaultSlot = () => {
      return newParams.message;
    };
    const vnode = vue.createVNode(script$1$6, propsData, defaultSlot);
    vnode.appContext = app._context;
    return vue.render(vnode, document.createElement("div"));
  }
};
var index$d = {
  install(app) {
    localVueInstance = app;
    registerComponent(app, script$c);
    registerComponentProgrammatic(app, "notification", NotificationProgrammatic);
  }
};
var index$d$1 = index$d;
var script$b = vue.defineComponent({
  name: "OPaginationButton",
  inject: ["$pagination"],
  configField: "pagination",
  props: {
    page: {
      type: Object,
      required: true
    },
    tag: {
      type: String,
      default: "a",
      validator: (value) => getValueByPath(getOptions$1(), "linkTags", ["a", "button", "input", "router-link", "nuxt-link"]).indexOf(value) >= 0
    },
    disabled: {
      type: Boolean,
      default: false
    },
    linkClass: [String, Array, Object],
    linkCurrentClass: [String, Array, Object]
  },
  computed: {
    linkClasses() {
      return [
        this.linkClass || [...this.$pagination.linkClasses],
        this.page.class,
        { [this.linkCurrentClass || this.$pagination.linkCurrentClasses]: this.page.isCurrent }
      ];
    },
    href() {
      if (this.tag === "a") {
        return "#";
      }
      return "";
    },
    isDisabled() {
      if (this.tag === "a")
        return null;
      return this.disabled || this.page.disabled;
    }
  }
});
function render$9(_ctx, _cache, $props, $setup, $data, $options) {
  return vue.openBlock(), vue.createBlock(vue.resolveDynamicComponent(_ctx.tag), vue.mergeProps({
    role: "button",
    href: _ctx.href,
    disabled: _ctx.isDisabled,
    class: _ctx.linkClasses
  }, _ctx.$attrs, {
    onClick: vue.withModifiers(_ctx.page.click, ["prevent"]),
    "aria-label": _ctx.page["aria-label"],
    "aria-current": _ctx.page.isCurrent
  }), {
    default: vue.withCtx(() => [vue.renderSlot(_ctx.$slots, "default", {}, () => [vue.createTextVNode(vue.toDisplayString(_ctx.page.number), 1)])]),
    _: 3
  }, 16, ["href", "disabled", "class", "onClick", "aria-label", "aria-current"]);
}
script$b.render = render$9;
script$b.__file = "src/components/pagination/PaginationButton.vue";
var script$1$5 = vue.defineComponent({
  name: "OPagination",
  components: {
    [script$r.name]: script$r,
    [script$b.name]: script$b
  },
  configField: "pagination",
  mixins: [BaseComponentMixin, MatchMediaMixin],
  provide() {
    return {
      $pagination: this
    };
  },
  emits: ["update:active", "change", "update:current"],
  props: {
    total: [Number, String],
    perPage: {
      type: [Number, String],
      default: () => {
        return getValueByPath(getOptions$1(), "pagination.perPage", 20);
      }
    },
    current: {
      type: [Number, String],
      default: 1
    },
    rangeBefore: {
      type: [Number, String],
      default: 1
    },
    rangeAfter: {
      type: [Number, String],
      default: 1
    },
    size: String,
    simple: Boolean,
    rounded: Boolean,
    order: String,
    iconPack: String,
    iconPrev: {
      type: String,
      default: () => {
        return getValueByPath(getOptions$1(), "pagination.iconPrev", "chevron-left");
      }
    },
    iconNext: {
      type: String,
      default: () => {
        return getValueByPath(getOptions$1(), "pagination.iconNext", "chevron-right");
      }
    },
    ariaNextLabel: String,
    ariaPreviousLabel: String,
    ariaPageLabel: String,
    ariaCurrentLabel: String,
    rootClass: [String, Function, Array],
    prevBtnClass: [String, Function, Array],
    nextBtnClass: [String, Function, Array],
    listClass: [String, Function, Array],
    linkClass: [String, Function, Array],
    linkCurrentClass: [String, Function, Array],
    ellipsisClass: [String, Function, Array],
    infoClass: [String, Function, Array],
    orderClass: [String, Function, Array],
    simpleClass: [String, Function, Array],
    roundedClass: [String, Function, Array],
    linkDisabledClass: [String, Function, Array],
    sizeClass: [String, Function, Array],
    mobileClass: [String, Function, Array]
  },
  computed: {
    rootClasses() {
      return [
        this.computedClass("rootClass", "o-pag"),
        { [this.computedClass("orderClass", "o-pag--", this.order)]: this.order },
        { [this.computedClass("sizeClass", "o-pag--", this.size)]: this.size },
        { [this.computedClass("simpleClass", "o-pag--simple")]: this.simple },
        { [this.computedClass("mobileClass", "o-pag--mobile")]: this.isMatchMedia }
      ];
    },
    prevBtnClasses() {
      return [
        this.computedClass("prevBtnClass", "o-pag__previous"),
        { [this.computedClass("linkDisabledClass", "o-pag__link--disabled")]: !this.hasPrev }
      ];
    },
    nextBtnClasses() {
      return [
        this.computedClass("nextBtnClass", "o-pag__next"),
        { [this.computedClass("linkDisabledClass", "o-pag__link--disabled")]: !this.hasNext }
      ];
    },
    infoClasses() {
      return [
        this.computedClass("infoClass", "o-pag__info")
      ];
    },
    ellipsisClasses() {
      return [
        this.computedClass("ellipsisClass", "o-pag__ellipsis")
      ];
    },
    listClasses() {
      return [
        this.computedClass("listClass", "o-pag__list")
      ];
    },
    linkClasses() {
      return [
        this.computedClass("linkClass", "o-pag__link"),
        { [this.computedClass("roundedClass", "o-pag__link--rounded")]: this.rounded }
      ];
    },
    linkCurrentClasses() {
      return [
        this.computedClass("linkCurrentClass", "o-pag__link--current")
      ];
    },
    beforeCurrent() {
      return parseInt(this.rangeBefore);
    },
    afterCurrent() {
      return parseInt(this.rangeAfter);
    },
    pageCount() {
      return Math.ceil(this.total / this.perPage);
    },
    firstItem() {
      const firstItem = this.current * this.perPage - this.perPage + 1;
      return firstItem >= 0 ? firstItem : 0;
    },
    hasPrev() {
      return this.current > 1;
    },
    hasFirst() {
      return this.current >= 2 + this.beforeCurrent;
    },
    hasFirstEllipsis() {
      return this.current >= this.beforeCurrent + 4;
    },
    hasLast() {
      return this.current <= this.pageCount - (1 + this.afterCurrent);
    },
    hasLastEllipsis() {
      return this.current < this.pageCount - (2 + this.afterCurrent);
    },
    hasNext() {
      return this.current < this.pageCount;
    },
    pagesInRange() {
      if (this.simple)
        return;
      let left = Math.max(1, this.current - this.beforeCurrent);
      if (left - 1 === 2) {
        left--;
      }
      let right = Math.min(this.current + this.afterCurrent, this.pageCount);
      if (this.pageCount - right === 2) {
        right++;
      }
      const pages = [];
      for (let i = left; i <= right; i++) {
        pages.push(this.getPage(i));
      }
      return pages;
    },
    hasDefaultSlot() {
      return this.$slots.default;
    },
    hasPreviousSlot() {
      return this.$slots.previous;
    },
    hasNextSlot() {
      return this.$slots.next;
    }
  },
  watch: {
    pageCount(value) {
      if (this.current > value)
        this.last();
    }
  },
  methods: {
    prev(event) {
      this.changePage(this.current - 1, event);
    },
    next(event) {
      this.changePage(this.current + 1, event);
    },
    first(event) {
      this.changePage(1, event);
    },
    last(event) {
      this.changePage(this.pageCount, event);
    },
    changePage(num, event) {
      if (this.current === num || num < 1 || num > this.pageCount)
        return;
      this.$emit("change", num);
      this.$emit("update:current", num);
      if (event && event.target) {
        this.$nextTick(() => event.target.focus());
      }
    },
    getPage(num, options = {}) {
      return {
        number: num,
        isCurrent: this.current === num,
        click: (event) => this.changePage(num, event),
        disabled: options.disabled || false,
        class: options.class || "",
        "aria-label": options["aria-label"] || this.getAriaPageLabel(num, this.current === num)
      };
    },
    getAriaPageLabel(pageNumber, isCurrent) {
      if (this.ariaPageLabel && (!isCurrent || !this.ariaCurrentLabel)) {
        return this.ariaPageLabel + " " + pageNumber + ".";
      } else if (this.ariaPageLabel && isCurrent && this.ariaCurrentLabel) {
        return this.ariaCurrentLabel + ", " + this.ariaPageLabel + " " + pageNumber + ".";
      }
      return null;
    }
  }
});
const _hoisted_1$4 = {
  key: 0
};
const _hoisted_2$2 = {
  key: 1
};
const _hoisted_3$2 = {
  key: 2
};
const _hoisted_4$2 = {
  key: 3
};
function render$1$3(_ctx, _cache, $props, $setup, $data, $options) {
  const _component_o_icon = vue.resolveComponent("o-icon");
  const _component_o_pagination_button = vue.resolveComponent("o-pagination-button");
  return vue.openBlock(), vue.createBlock("nav", {
    class: _ctx.rootClasses
  }, [_ctx.hasPreviousSlot ? vue.renderSlot(_ctx.$slots, "previous", {
    key: 0,
    linkClass: _ctx.linkClasses,
    linkCurrentClass: _ctx.linkCurrentClasses,
    page: _ctx.getPage(_ctx.current - 1, {
      class: _ctx.prevBtnClasses,
      "aria-label": _ctx.ariaPreviousLabel
    })
  }, () => [vue.createVNode(_component_o_icon, {
    icon: _ctx.iconPrev,
    pack: _ctx.iconPack,
    both: "",
    "aria-hidden": "true"
  }, null, 8, ["icon", "pack"])]) : vue.createVNode(_component_o_pagination_button, {
    key: 1,
    class: _ctx.prevBtnClasses,
    linkClass: _ctx.linkClasses,
    linkCurrentClass: _ctx.linkCurrentClasses,
    page: _ctx.getPage(_ctx.current - 1)
  }, {
    default: vue.withCtx(() => [vue.createVNode(_component_o_icon, {
      icon: _ctx.iconPrev,
      pack: _ctx.iconPack,
      both: "",
      "aria-hidden": "true"
    }, null, 8, ["icon", "pack"])]),
    _: 1
  }, 8, ["class", "linkClass", "linkCurrentClass", "page"]), _ctx.hasNextSlot ? vue.renderSlot(_ctx.$slots, "next", {
    key: 2,
    linkClass: _ctx.linkClasses,
    linkCurrentClass: _ctx.linkCurrentClasses,
    page: _ctx.getPage(_ctx.current + 1, {
      class: _ctx.nextBtnClasses,
      "aria-label": _ctx.ariaNextLabel
    })
  }, () => [vue.createVNode(_component_o_icon, {
    icon: _ctx.iconNext,
    pack: _ctx.iconPack,
    both: "",
    "aria-hidden": "true"
  }, null, 8, ["icon", "pack"])]) : vue.createVNode(_component_o_pagination_button, {
    key: 3,
    class: _ctx.nextBtnClasses,
    linkClass: _ctx.linkClasses,
    linkCurrentClass: _ctx.linkCurrentClasses,
    page: _ctx.getPage(_ctx.current + 1)
  }, {
    default: vue.withCtx(() => [vue.createVNode(_component_o_icon, {
      icon: _ctx.iconNext,
      pack: _ctx.iconPack,
      both: "",
      "aria-hidden": "true"
    }, null, 8, ["icon", "pack"])]),
    _: 1
  }, 8, ["class", "linkClass", "linkCurrentClass", "page"]), _ctx.simple ? (vue.openBlock(), vue.createBlock("small", {
    key: 4,
    class: _ctx.infoClasses
  }, [_ctx.perPage == 1 ? (vue.openBlock(), vue.createBlock(vue.Fragment, {
    key: 0
  }, [vue.createTextVNode(vue.toDisplayString(_ctx.firstItem) + " / " + vue.toDisplayString(_ctx.total), 1)], 64)) : (vue.openBlock(), vue.createBlock(vue.Fragment, {
    key: 1
  }, [vue.createTextVNode(vue.toDisplayString(_ctx.firstItem) + "-" + vue.toDisplayString(Math.min(_ctx.current * _ctx.perPage, _ctx.total)) + " / " + vue.toDisplayString(_ctx.total), 1)], 64))], 2)) : (vue.openBlock(), vue.createBlock("ul", {
    key: 5,
    class: _ctx.listClasses
  }, [vue.createCommentVNode("First"), _ctx.hasFirst ? (vue.openBlock(), vue.createBlock("li", _hoisted_1$4, [_ctx.hasDefaultSlot ? vue.renderSlot(_ctx.$slots, "default", {
    key: 0,
    page: _ctx.getPage(1),
    linkClass: _ctx.linkClasses,
    linkCurrentClass: _ctx.linkCurrentClasses
  }) : vue.createVNode(_component_o_pagination_button, {
    key: 1,
    linkClass: _ctx.linkClasses,
    linkCurrentClass: _ctx.linkCurrentClasses,
    page: _ctx.getPage(1)
  }, null, 8, ["linkClass", "linkCurrentClass", "page"])])) : vue.createCommentVNode("v-if", true), _ctx.hasFirstEllipsis ? (vue.openBlock(), vue.createBlock("li", _hoisted_2$2, [vue.createVNode("span", {
    class: _ctx.ellipsisClasses
  }, "\u2026", 2)])) : vue.createCommentVNode("v-if", true), vue.createCommentVNode("Pages"), (vue.openBlock(true), vue.createBlock(vue.Fragment, null, vue.renderList(_ctx.pagesInRange, (page) => {
    return vue.openBlock(), vue.createBlock("li", {
      key: page.number
    }, [_ctx.hasDefaultSlot ? vue.renderSlot(_ctx.$slots, "default", {
      key: 0,
      page,
      linkClass: _ctx.linkClasses,
      linkCurrentClass: _ctx.linkCurrentClasses
    }) : vue.createVNode(_component_o_pagination_button, {
      key: 1,
      linkClass: _ctx.linkClasses,
      linkCurrentClass: _ctx.linkCurrentClasses,
      page
    }, null, 8, ["linkClass", "linkCurrentClass", "page"])]);
  }), 128)), vue.createCommentVNode("Last"), _ctx.hasLastEllipsis ? (vue.openBlock(), vue.createBlock("li", _hoisted_3$2, [vue.createVNode("span", {
    class: _ctx.ellipsisClasses
  }, "\u2026", 2)])) : vue.createCommentVNode("v-if", true), _ctx.hasLast ? (vue.openBlock(), vue.createBlock("li", _hoisted_4$2, [_ctx.hasDefaultSlot ? vue.renderSlot(_ctx.$slots, "default", {
    key: 0,
    page: _ctx.getPage(_ctx.pageCount),
    linkClass: _ctx.linkClasses,
    linkCurrentClass: _ctx.linkCurrentClasses
  }) : vue.createVNode(_component_o_pagination_button, {
    key: 1,
    linkClass: _ctx.linkClasses,
    linkCurrentClass: _ctx.linkCurrentClasses,
    page: _ctx.getPage(_ctx.pageCount)
  }, null, 8, ["linkClass", "linkCurrentClass", "page"])])) : vue.createCommentVNode("v-if", true)], 2))], 2);
}
script$1$5.render = render$1$3;
script$1$5.__file = "src/components/pagination/Pagination.vue";
var index$c = {
  install(app) {
    registerComponent(app, script$1$5);
    registerComponent(app, script$b);
  }
};
var index$e = index$c;
var script$a = vue.defineComponent({
  name: "ORadio",
  components: {
    [script$r.name]: script$r
  },
  mixins: [BaseComponentMixin, CheckRadioMixin],
  configField: "radio",
  emits: [
    "input"
  ],
  props: {
    iconPack: {
      type: String,
      default: () => {
        return getValueByPath(getOptions$1(), "radio.iconPack", void 0);
      }
    },
    iconCheck: {
      type: String,
      default: () => {
        return getValueByPath(getOptions$1(), "radio.iconCheck", void 0);
      }
    },
    rootClass: [String, Function, Array],
    disabledClass: [String, Function, Array],
    checkCheckedClass: [String, Function, Array],
    checkClass: [String, Function, Array],
    labelClass: [String, Function, Array],
    sizeClass: [String, Function, Array],
    variantClass: [String, Function, Array],
    iconCheckClass: [String, Function, Array],
    iconCheckCheckedClass: [String, Function, Array]
  },
  computed: {
    isChecked() {
      return this.modelValue === this.nativeValue;
    },
    rootClasses() {
      return [
        this.computedClass("rootClass", "o-radio"),
        { [this.computedClass("checkedClass", "o-radio--checked")]: this.isChecked },
        { [this.computedClass("sizeClass", "o-radio--", this.size)]: this.size },
        { [this.computedClass("disabledClass", "o-radio--disabled")]: this.disabled },
        { [this.computedClass("variantClass", "o-radio--", this.variant)]: this.variant }
      ];
    },
    checkClasses() {
      return [
        this.computedClass("checkClass", "o-radio__check"),
        { [this.computedClass("checkCheckedClass", "o-radio__check--checked")]: this.isChecked }
      ];
    },
    labelClasses() {
      return [
        this.computedClass("labelClass", "o-radio__label")
      ];
    },
    iconCheckClasses() {
      return [
        this.computedClass("iconCheckClass", "o-radio__icon"),
        { [this.computedClass("iconCheckCheckedClass", "o-radio__icon--checked")]: this.isChecked }
      ];
    }
  }
});
function render$8(_ctx, _cache, $props, $setup, $data, $options) {
  const _component_o_icon = vue.resolveComponent("o-icon");
  return vue.openBlock(), vue.createBlock("label", {
    class: _ctx.rootClasses,
    ref: "label",
    onClick: _cache[3] || (_cache[3] = vue.withModifiers((...args) => _ctx.focus(...args), ["stop"])),
    onKeydown: _cache[4] || (_cache[4] = vue.withKeys(vue.withModifiers(($event) => _ctx.$refs.label.click(), ["prevent"]), ["enter"]))
  }, [_ctx.iconCheck ? vue.createVNode(_component_o_icon, {
    key: 0,
    icon: _ctx.iconCheck,
    pack: _ctx.iconPack,
    size: _ctx.size,
    class: _ctx.iconCheckClasses
  }, null, 8, ["icon", "pack", "size", "class"]) : vue.createCommentVNode("v-if", true), vue.withDirectives(vue.createVNode("input", {
    "onUpdate:modelValue": _cache[1] || (_cache[1] = ($event) => _ctx.computedValue = $event),
    type: "radio",
    ref: "input",
    onClick: _cache[2] || (_cache[2] = vue.withModifiers(() => {
    }, ["stop"])),
    class: _ctx.checkClasses,
    disabled: _ctx.disabled,
    required: _ctx.required,
    name: _ctx.name,
    value: _ctx.nativeValue
  }, null, 10, ["disabled", "required", "name", "value"]), [[vue.vModelRadio, _ctx.computedValue]]), vue.createVNode("span", {
    class: _ctx.labelClasses
  }, [vue.renderSlot(_ctx.$slots, "default")], 2)], 34);
}
script$a.render = render$8;
script$a.__file = "src/components/radio/Radio.vue";
var index$b = {
  install(app) {
    registerComponent(app, script$a);
  }
};
var index$f = index$b;
var index$a = {
  install(app) {
    registerComponent(app, script$j);
  }
};
var index$g = index$a;
var script$9 = vue.defineComponent({
  name: "OSkeleton",
  mixins: [BaseComponentMixin],
  configField: "skeleton",
  props: {
    active: {
      type: Boolean,
      default: true
    },
    animated: {
      type: Boolean,
      default: true
    },
    width: [Number, String],
    height: [Number, String],
    circle: Boolean,
    rounded: {
      type: Boolean,
      default: true
    },
    count: {
      type: Number,
      default: 1
    },
    position: {
      type: String,
      default: "left",
      validator(value) {
        return [
          "left",
          "centered",
          "right"
        ].indexOf(value) > -1;
      }
    },
    size: String,
    rootClass: [String, Function, Array],
    animationClass: [String, Function, Array],
    positionClass: [String, Function, Array],
    itemClass: [String, Function, Array],
    itemRoundedClass: [String, Function, Array],
    sizeClass: [String, Function, Array]
  },
  render() {
    if (!this.active)
      return;
    const items2 = [];
    const width = this.width;
    const height = this.height;
    for (let i = 0; i < this.count; i++) {
      items2.push(vue.h("div", {
        class: [
          this.computedClass("itemClass", "o-sklt__item"),
          { [this.computedClass("itemRoundedClass", "o-sklt__item--rounded")]: this.rounded },
          { [this.computedClass("animationClass", "o-sklt__item--animated")]: this.animated },
          { [this.computedClass("sizeClass", "o-sklt__item--", this.size)]: this.size }
        ],
        key: i,
        style: {
          height: toCssDimension(height),
          width: toCssDimension(width),
          borderRadius: this.circle ? "50%" : null
        }
      }));
    }
    return vue.h("div", {
      class: [
        this.computedClass("rootClass", "o-sklt"),
        { [this.computedClass("positionClass", "o-sklt--", this.position)]: this.position }
      ]
    }, items2);
  }
});
script$9.__file = "src/components/skeleton/Skeleton.vue";
var index$9 = {
  install(app) {
    registerComponent(app, script$9);
  }
};
var index$h = index$9;
var script$8 = vue.defineComponent({
  name: "OSidebar",
  mixins: [BaseComponentMixin, MatchMediaMixin],
  configField: "sidebar",
  emits: ["update:open", "close"],
  props: {
    open: Boolean,
    variant: [String, Object],
    overlay: Boolean,
    position: {
      type: String,
      default: () => {
        return getValueByPath(getOptions$1(), "sidebar.position", "fixed");
      },
      validator: (value) => {
        return [
          "fixed",
          "absolute",
          "static"
        ].indexOf(value) >= 0;
      }
    },
    fullheight: Boolean,
    fullwidth: Boolean,
    right: Boolean,
    mobile: {
      type: String,
      validator: (value) => {
        return [
          "",
          "fullwidth",
          "reduced",
          "hidden"
        ].indexOf(value) >= 0;
      }
    },
    reduce: Boolean,
    expandOnHover: Boolean,
    expandOnHoverFixed: Boolean,
    canCancel: {
      type: [Array, Boolean],
      default: () => {
        return getValueByPath(getOptions$1(), "sidebar.canCancel", ["escape", "outside"]);
      }
    },
    onCancel: {
      type: Function,
      default: () => {
      }
    },
    scroll: {
      type: String,
      default: () => {
        return getValueByPath(getOptions$1(), "sidebar.scroll", "clip");
      },
      validator: (value) => {
        return [
          "clip",
          "keep"
        ].indexOf(value) >= 0;
      }
    },
    rootClass: [String, Function, Array],
    overlayClass: [String, Function, Array],
    contentClass: [String, Function, Array],
    fixedClass: [String, Function, Array],
    staticClass: [String, Function, Array],
    absoluteClass: [String, Function, Array],
    fullheightClass: [String, Function, Array],
    fullwidthClass: [String, Function, Array],
    rightClass: [String, Function, Array],
    reduceClass: [String, Function, Array],
    expandOnHoverClass: [String, Function, Array],
    expandOnHoverFixedClass: [String, Function, Array],
    variantClass: [String, Function, Array],
    mobileClass: [String, Function, Array]
  },
  data() {
    return {
      isOpen: this.open,
      transitionName: null,
      animating: true,
      savedScrollTop: null
    };
  },
  computed: {
    rootClasses() {
      return [
        this.computedClass("rootClass", "o-side"),
        { [this.computedClass("mobileClass", "o-side--mobile")]: this.isMatchMedia }
      ];
    },
    overlayClasses() {
      return [
        this.computedClass("overlayClass", "o-side__overlay")
      ];
    },
    contentClasses() {
      return [
        this.computedClass("contentClass", "o-side__content"),
        { [this.computedClass("variantClass", "o-side__content--", this.variant)]: this.variant },
        { [this.computedClass("fixedClass", "o-side__content--fixed")]: this.isFixed },
        { [this.computedClass("staticClass", "o-side__content--static")]: this.isStatic },
        { [this.computedClass("absoluteClass", "o-side__content--absolute")]: this.isAbsolute },
        { [this.computedClass("fullheightClass", "o-side__content--fullheight")]: this.fullheight },
        { [this.computedClass("fullwidthClass", "o-side__content--fullwidth")]: this.fullwidth || this.mobile === "fullwidth" && this.isMatchMedia },
        { [this.computedClass("rightClass", "o-side__content--right")]: this.right },
        { [this.computedClass("reduceClass", "o-side__content--mini")]: this.reduce || this.mobile === "reduced" && this.isMatchMedia },
        { [this.computedClass("expandOnHoverClass", "o-side__content--mini-expand")]: this.expandOnHover && this.mobile !== "fullwidth" },
        { [this.computedClass("expandOnHoverFixedClass", "o-side__content--expand-mini-hover-fixed")]: this.expandOnHover && this.expandOnHoverFixed && this.mobile !== "fullwidth" }
      ];
    },
    cancelOptions() {
      return typeof this.canCancel === "boolean" ? this.canCancel ? getValueByPath(getOptions$1(), "sidebar.canCancel", ["escape", "outside"]) : [] : this.canCancel;
    },
    isStatic() {
      return this.position === "static";
    },
    isFixed() {
      return this.position === "fixed";
    },
    isAbsolute() {
      return this.position === "absolute";
    },
    hideOnMobile() {
      return this.mobile === "hidden" && this.isMatchMedia;
    }
  },
  watch: {
    open: {
      handler(value) {
        this.isOpen = value;
        if (this.overlay) {
          this.handleScroll();
        }
        const open = this.right ? !value : value;
        this.transitionName = !open ? "slide-prev" : "slide-next";
      },
      immediate: true
    }
  },
  methods: {
    whiteList() {
      const whiteList = [];
      whiteList.push(this.$refs.sidebarContent);
      if (this.$refs.sidebarContent !== void 0) {
        const children = this.$refs.sidebarContent.querySelectorAll("*");
        for (const child of children) {
          whiteList.push(child);
        }
      }
      return whiteList;
    },
    keyPress({ key }) {
      if (this.isFixed) {
        if (this.isOpen && (key === "Escape" || key === "Esc"))
          this.cancel("escape");
      }
    },
    cancel(method) {
      if (this.cancelOptions.indexOf(method) < 0)
        return;
      if (this.isStatic)
        return;
      this.onCancel.apply(null, arguments);
      this.close();
    },
    close() {
      this.isOpen = false;
      this.$emit("close");
      this.$emit("update:open", false);
    },
    clickedOutside(event) {
      if (this.isFixed) {
        if (this.isOpen && !this.animating) {
          if (this.whiteList().indexOf(event.target) < 0) {
            this.cancel("outside");
          }
        }
      }
    },
    beforeEnter() {
      this.animating = true;
    },
    afterEnter() {
      this.animating = false;
    },
    handleScroll() {
      if (typeof window === "undefined")
        return;
      if (this.scroll === "clip") {
        if (this.open) {
          document.documentElement.classList.add("o-clipped");
        } else {
          document.documentElement.classList.remove("o-clipped");
        }
        return;
      }
      this.savedScrollTop = !this.savedScrollTop ? document.documentElement.scrollTop : this.savedScrollTop;
      if (this.open) {
        document.body.classList.add("o-noscroll");
      } else {
        document.body.classList.remove("o-noscroll");
      }
      if (this.open) {
        document.body.style.top = `-${this.savedScrollTop}px`;
        return;
      }
      document.documentElement.scrollTop = this.savedScrollTop;
      document.body.style.top = null;
      this.savedScrollTop = null;
    }
  },
  created() {
    if (typeof window !== "undefined") {
      document.addEventListener("keyup", this.keyPress);
      document.addEventListener("click", this.clickedOutside);
    }
  },
  mounted() {
    if (typeof window !== "undefined") {
      if (this.isFixed) {
        document.body.appendChild(this.$el);
      }
      if (this.overlay && this.open) {
        this.handleScroll();
      }
    }
  },
  beforeUnmount() {
    if (typeof window !== "undefined") {
      document.removeEventListener("keyup", this.keyPress);
      document.removeEventListener("click", this.clickedOutside);
      if (this.overlay) {
        document.documentElement.classList.remove("o-clipped");
        const savedScrollTop = !this.savedScrollTop ? document.documentElement.scrollTop : this.savedScrollTop;
        document.body.classList.remove("o-noscroll");
        document.documentElement.scrollTop = savedScrollTop;
        document.body.style.top = null;
      }
    }
    if (this.isFixed) {
      removeElement(this.$el);
    }
  }
});
function render$7(_ctx, _cache, $props, $setup, $data, $options) {
  return vue.withDirectives((vue.openBlock(), vue.createBlock("div", {
    class: _ctx.rootClasses
  }, [_ctx.overlay && _ctx.isOpen ? (vue.openBlock(), vue.createBlock("div", {
    key: 0,
    class: _ctx.overlayClasses
  }, null, 2)) : vue.createCommentVNode("v-if", true), vue.createVNode(vue.Transition, {
    name: _ctx.transitionName,
    "onBefore-enter": _ctx.beforeEnter,
    "onAfter-enter": _ctx.afterEnter
  }, {
    default: vue.withCtx(() => [vue.withDirectives(vue.createVNode("div", {
      ref: "sidebarContent",
      class: _ctx.contentClasses
    }, [vue.renderSlot(_ctx.$slots, "default")], 2), [[vue.vShow, _ctx.isOpen]])]),
    _: 3
  }, 8, ["name", "onBefore-enter", "onAfter-enter"])], 2)), [[vue.vShow, !_ctx.hideOnMobile]]);
}
script$8.render = render$7;
script$8.__file = "src/components/sidebar/Sidebar.vue";
var index$8 = {
  install(app) {
    registerComponent(app, script$8);
  }
};
var index$i = index$8;
var script$7 = vue.defineComponent({
  name: "OTooltip",
  mixins: [BaseComponentMixin],
  configField: "tooltip",
  props: {
    active: {
      type: Boolean,
      default: true
    },
    label: String,
    delay: Number,
    position: {
      type: String,
      default: () => {
        return getValueByPath(getOptions$1(), "tooltip.position", "top");
      },
      validator(value) {
        return [
          "top",
          "bottom",
          "left",
          "right"
        ].indexOf(value) > -1;
      }
    },
    triggers: {
      type: Array,
      default: () => {
        return getValueByPath(getOptions$1(), "tooltip.triggers", ["hover"]);
      }
    },
    always: Boolean,
    animated: {
      type: Boolean,
      default: true
    },
    animation: {
      type: String,
      default: () => {
        return getValueByPath(getOptions$1(), "tooltip.animation", "fade");
      }
    },
    autoClose: {
      type: [Array, Boolean],
      default: true
    },
    multiline: Boolean,
    appendToBody: Boolean,
    variant: [String, Function, Array],
    rootClass: [String, Function, Array],
    contentClass: [String, Function, Array],
    orderClass: [String, Function, Array],
    triggerClass: [String, Function, Array],
    multilineClass: [String, Function, Array],
    alwaysClass: [String, Function, Array],
    variantClass: [String, Function, Array],
    arrowClass: [String, Function, Array],
    arrowOrderClass: [String, Function, Array]
  },
  data() {
    return {
      isActive: false,
      triggerStyle: {},
      bodyEl: void 0
    };
  },
  computed: {
    rootClasses() {
      return [
        this.computedClass("rootClass", "o-tip")
      ];
    },
    triggerClasses() {
      return [
        this.computedClass("triggerClass", "o-tip__trigger")
      ];
    },
    arrowClasses() {
      return [
        this.computedClass("arrowClass", "o-tip__arrow"),
        { [this.computedClass("arrowOrderClass", "o-tip__arrow--", this.position)]: this.position },
        { [this.computedClass("variantArrowClass", "o-tip__arrow--", this.variant)]: this.variant }
      ];
    },
    contentClasses() {
      return [
        this.computedClass("contentClass", "o-tip__content"),
        { [this.computedClass("orderClass", "o-tip__content--", this.position)]: this.position },
        { [this.computedClass("variantClass", "o-tip__content--", this.variant)]: this.variant },
        { [this.computedClass("multilineClass", "o-tip__content--multiline")]: this.multiline },
        { [this.computedClass("alwaysClass", "o-tip__content--always")]: this.always }
      ];
    },
    newAnimation() {
      return this.animated ? this.animation : void 0;
    }
  },
  watch: {
    isActive(value) {
      if (value && this.appendToBody) {
        this.updateAppendToBody();
      }
    }
  },
  methods: {
    updateAppendToBody() {
      const tooltip = this.$refs.tooltip;
      const trigger = this.$refs.trigger;
      if (tooltip && trigger) {
        const tooltipEl = this.$data.bodyEl.children[0];
        tooltipEl.classList.forEach((item) => tooltipEl.classList.remove(item));
        if (this.$vnode && this.$vnode.data && this.$vnode.data.staticClass) {
          tooltipEl.classList.add(this.$vnode.data.staticClass);
        }
        this.rootClasses.forEach((item) => {
          if (typeof item === "object") {
            Object.keys(item).filter((key) => !!item[key]).forEach((key) => tooltipEl.classList.add(key));
          } else {
            tooltipEl.classList.add(item);
          }
        });
        tooltipEl.style.width = `${trigger.clientWidth}px`;
        tooltipEl.style.height = `${trigger.clientHeight}px`;
        const rect = trigger.getBoundingClientRect();
        const top = rect.top + window.scrollY;
        const left = rect.left + window.scrollX;
        const wrapper = this.$data.bodyEl;
        wrapper.style.position = "absolute";
        wrapper.style.top = `${top}px`;
        wrapper.style.left = `${left}px`;
        wrapper.style.zIndex = this.isActive || this.always ? "99" : "-1";
        this.triggerStyle = { zIndex: this.isActive || this.always ? "100" : void 0 };
      }
    },
    onClick() {
      if (this.triggers.indexOf("click") < 0)
        return;
      this.$nextTick(() => {
        setTimeout(() => this.open());
      });
    },
    onHover() {
      if (this.triggers.indexOf("hover") < 0)
        return;
      this.open();
    },
    onFocus() {
      if (this.triggers.indexOf("focus") < 0)
        return;
      this.open();
    },
    onContextMenu(event) {
      if (this.triggers.indexOf("contextmenu") < 0)
        return;
      event.preventDefault();
      this.open();
    },
    open() {
      if (this.delay) {
        this.timer = setTimeout(() => {
          this.isActive = true;
          this.timer = null;
        }, this.delay);
      } else {
        this.isActive = true;
      }
    },
    close() {
      if (typeof this.autoClose === "boolean") {
        this.isActive = !this.autoClose;
      }
      if (this.autoClose && this.timer)
        clearTimeout(this.timer);
    },
    clickedOutside(event) {
      if (this.isActive) {
        if (Array.isArray(this.autoClose)) {
          if (this.autoClose.indexOf("outside") >= 0) {
            if (!this.isInWhiteList(event.target))
              this.isActive = false;
          }
          if (this.autoClose.indexOf("inside") >= 0) {
            if (this.isInWhiteList(event.target))
              this.isActive = false;
          }
        }
      }
    },
    keyPress({ key }) {
      if (this.isActive && (key === "Escape" || key === "Esc")) {
        if (Array.isArray(this.autoClose)) {
          if (this.autoClose.indexOf("escape") >= 0)
            this.isActive = false;
        }
      }
    },
    isInWhiteList(el) {
      if (el === this.$refs.content)
        return true;
      if (this.$refs.content !== void 0) {
        const children = this.$refs.content.querySelectorAll("*");
        for (const child of children) {
          if (el === child) {
            return true;
          }
        }
      }
      return false;
    }
  },
  mounted() {
    if (this.appendToBody) {
      this.$data.bodyEl = createAbsoluteElement(this.$refs.content);
      this.updateAppendToBody();
    }
  },
  created() {
    if (typeof window !== "undefined") {
      document.addEventListener("click", this.clickedOutside);
      document.addEventListener("keyup", this.keyPress);
    }
  },
  beforeUnmount() {
    if (typeof window !== "undefined") {
      document.removeEventListener("click", this.clickedOutside);
      document.removeEventListener("keyup", this.keyPress);
    }
    if (this.appendToBody) {
      removeElement(this.$data.bodyEl);
    }
  }
});
function render$6(_ctx, _cache, $props, $setup, $data, $options) {
  return vue.openBlock(), vue.createBlock("span", {
    ref: "tooltip",
    class: _ctx.rootClasses
  }, [vue.createVNode(vue.Transition, {
    name: _ctx.newAnimation
  }, {
    default: vue.withCtx(() => [vue.withDirectives(vue.createVNode("div", {
      ref: "content",
      class: _ctx.contentClasses
    }, [vue.createVNode("span", {
      class: _ctx.arrowClasses
    }, null, 2), _ctx.label ? (vue.openBlock(), vue.createBlock(vue.Fragment, {
      key: 0
    }, [vue.createTextVNode(vue.toDisplayString(_ctx.label), 1)], 64)) : _ctx.$slots.default ? vue.renderSlot(_ctx.$slots, "content", {
      key: 1
    }) : vue.createCommentVNode("v-if", true)], 2), [[vue.vShow, _ctx.active && (_ctx.isActive || _ctx.always)]])]),
    _: 1
  }, 8, ["name"]), vue.createVNode("div", {
    ref: "trigger",
    class: _ctx.triggerClasses,
    style: _ctx.triggerStyle,
    onClick: _cache[1] || (_cache[1] = (...args) => _ctx.onClick(...args)),
    onContextmenu: _cache[2] || (_cache[2] = (...args) => _ctx.onContextMenu(...args)),
    onMouseenter: _cache[3] || (_cache[3] = (...args) => _ctx.onHover(...args)),
    onFocusCapture: _cache[4] || (_cache[4] = (...args) => _ctx.onFocus(...args)),
    onBlurCapture: _cache[5] || (_cache[5] = (...args) => _ctx.close(...args)),
    onMouseleave: _cache[6] || (_cache[6] = (...args) => _ctx.close(...args))
  }, [vue.renderSlot(_ctx.$slots, "default", {
    ref: "slot"
  })], 38)], 2);
}
script$7.render = render$6;
script$7.__file = "src/components/tooltip/Tooltip.vue";
var script$6 = vue.defineComponent({
  name: "OSliderThumb",
  components: {
    [script$7.name]: script$7
  },
  configField: "slider",
  inheritAttrs: false,
  inject: ["$slider"],
  emits: ["update:modelValue", "dragstart", "dragend"],
  props: {
    modelValue: {
      type: Number,
      default: 0
    },
    variant: {
      type: String,
      default: ""
    },
    tooltip: {
      type: Boolean,
      default: true
    },
    indicator: {
      type: Boolean,
      default: false
    },
    customFormatter: Function,
    format: {
      type: String,
      default: "raw",
      validator: (value) => {
        return [
          "raw",
          "percent"
        ].indexOf(value) >= 0;
      }
    },
    locale: {
      type: [String, Array],
      default: () => {
        return getValueByPath(getOptions$1(), "locale");
      }
    },
    tooltipAlways: {
      type: Boolean,
      default: false
    }
  },
  data() {
    return {
      isFocused: false,
      dragging: false,
      startX: 0,
      startPosition: 0,
      newPosition: null,
      oldValue: this.modelValue
    };
  },
  computed: {
    disabled() {
      return this.$parent.disabled;
    },
    max() {
      return this.$parent.max;
    },
    min() {
      return this.$parent.min;
    },
    step() {
      return this.$parent.step;
    },
    precision() {
      return this.$parent.precision;
    },
    currentPosition() {
      return `${(this.modelValue - this.min) / (this.max - this.min) * 100}%`;
    },
    wrapperStyle() {
      return { left: this.currentPosition };
    },
    formattedValue() {
      if (typeof this.customFormatter !== "undefined") {
        return this.customFormatter(this.modelValue);
      }
      if (this.format === "percent") {
        return new Intl.NumberFormat(this.locale, {
          style: "percent"
        }).format((this.modelValue - this.min) / (this.max - this.min));
      }
      return new Intl.NumberFormat(this.locale).format(this.modelValue);
    }
  },
  methods: {
    onFocus() {
      this.isFocused = true;
    },
    onBlur() {
      this.isFocused = false;
    },
    onButtonDown(event) {
      if (this.disabled)
        return;
      event.preventDefault();
      this.onDragStart(event);
      if (typeof window !== "undefined") {
        document.addEventListener("mousemove", this.onDragging);
        document.addEventListener("touchmove", this.onDragging);
        document.addEventListener("mouseup", this.onDragEnd);
        document.addEventListener("touchend", this.onDragEnd);
        document.addEventListener("contextmenu", this.onDragEnd);
      }
    },
    onLeftKeyDown() {
      if (this.disabled || this.modelvalue === this.min)
        return;
      this.newPosition = parseFloat(this.currentPosition) - this.step / (this.max - this.min) * 100;
      this.setPosition(this.newPosition);
      this.$parent.emitValue("change");
    },
    onRightKeyDown() {
      if (this.disabled || this.modelvalue === this.max)
        return;
      this.newPosition = parseFloat(this.currentPosition) + this.step / (this.max - this.min) * 100;
      this.setPosition(this.newPosition);
      this.$parent.emitValue("change");
    },
    onHomeKeyDown() {
      if (this.disabled || this.modelvalue === this.min)
        return;
      this.newPosition = 0;
      this.setPosition(this.newPosition);
      this.$parent.emitValue("change");
    },
    onEndKeyDown() {
      if (this.disabled || this.modelvalue === this.max)
        return;
      this.newPosition = 100;
      this.setPosition(this.newPosition);
      this.$parent.emitValue("change");
    },
    onDragStart(event) {
      this.dragging = true;
      this.$emit("dragstart");
      if (event.type === "touchstart") {
        event.clientX = event.touches[0].clientX;
      }
      this.startX = event.clientX;
      this.startPosition = parseFloat(this.currentPosition);
      this.newPosition = this.startPosition;
    },
    onDragging(event) {
      if (this.dragging) {
        if (event.type === "touchmove") {
          event.clientX = event.touches[0].clientX;
        }
        const diff2 = (event.clientX - this.startX) / this.$parent.sliderSize() * 100;
        this.newPosition = this.startPosition + diff2;
        this.setPosition(this.newPosition);
      }
    },
    onDragEnd() {
      this.dragging = false;
      this.$emit("dragend");
      if (this.modelvalue !== this.oldValue) {
        this.$parent.emitValue("change");
      }
      this.setPosition(this.newPosition);
      if (typeof window !== "undefined") {
        document.removeEventListener("mousemove", this.onDragging);
        document.removeEventListener("touchmove", this.onDragging);
        document.removeEventListener("mouseup", this.onDragEnd);
        document.removeEventListener("touchend", this.onDragEnd);
        document.removeEventListener("contextmenu", this.onDragEnd);
      }
    },
    setPosition(percent) {
      if (percent === null || isNaN(percent))
        return;
      if (percent < 0) {
        percent = 0;
      } else if (percent > 100) {
        percent = 100;
      }
      const stepLength = 100 / ((this.max - this.min) / this.step);
      const steps = Math.round(percent / stepLength);
      let value = steps * stepLength / 100 * (this.max - this.min) + this.min;
      value = parseFloat(value.toFixed(this.precision));
      this.$emit("update:modelValue", value);
      if (!this.dragging && value !== this.oldValue) {
        this.oldValue = value;
      }
    }
  }
});
const _hoisted_1$3 = {
  key: 0
};
function render$5(_ctx, _cache, $props, $setup, $data, $options) {
  const _component_o_tooltip = vue.resolveComponent("o-tooltip");
  return vue.openBlock(), vue.createBlock("div", {
    class: _ctx.$slider.thumbWrapperClasses,
    style: _ctx.wrapperStyle
  }, [vue.createVNode(_component_o_tooltip, {
    label: _ctx.formattedValue,
    variant: _ctx.variant,
    always: _ctx.dragging || _ctx.isFocused || _ctx.tooltipAlways,
    active: !_ctx.disabled && _ctx.tooltip
  }, {
    default: vue.withCtx(() => [vue.createVNode("div", vue.mergeProps(_ctx.$attrs, {
      class: _ctx.$slider.thumbClasses,
      tabindex: _ctx.disabled ? false : 0,
      onMousedown: _cache[1] || (_cache[1] = (...args) => _ctx.onButtonDown(...args)),
      onTouchstart: _cache[2] || (_cache[2] = (...args) => _ctx.onButtonDown(...args)),
      onFocus: _cache[3] || (_cache[3] = (...args) => _ctx.onFocus(...args)),
      onBlur: _cache[4] || (_cache[4] = (...args) => _ctx.onBlur(...args)),
      onKeydown: [_cache[5] || (_cache[5] = vue.withKeys(vue.withModifiers((...args) => _ctx.onLeftKeyDown(...args), ["prevent"]), ["left"])), _cache[6] || (_cache[6] = vue.withKeys(vue.withModifiers((...args) => _ctx.onRightKeyDown(...args), ["prevent"]), ["right"])), _cache[7] || (_cache[7] = vue.withKeys(vue.withModifiers((...args) => _ctx.onLeftKeyDown(...args), ["prevent"]), ["down"])), _cache[8] || (_cache[8] = vue.withKeys(vue.withModifiers((...args) => _ctx.onRightKeyDown(...args), ["prevent"]), ["up"])), _cache[9] || (_cache[9] = vue.withKeys(vue.withModifiers((...args) => _ctx.onHomeKeyDown(...args), ["prevent"]), ["home"])), _cache[10] || (_cache[10] = vue.withKeys(vue.withModifiers((...args) => _ctx.onEndKeyDown(...args), ["prevent"]), ["end"]))]
    }), [_ctx.indicator ? (vue.openBlock(), vue.createBlock("span", _hoisted_1$3, vue.toDisplayString(_ctx.formattedValue), 1)) : vue.createCommentVNode("v-if", true)], 16, ["tabindex"])]),
    _: 1
  }, 8, ["label", "variant", "always", "active"])], 6);
}
script$6.render = render$5;
script$6.__file = "src/components/slider/SliderThumb.vue";
var script$1$4 = vue.defineComponent({
  name: "OSliderTick",
  mixins: [BaseComponentMixin],
  configField: "slider",
  inject: ["$slider"],
  props: {
    value: {
      variant: Number,
      default: 0
    },
    tickClass: [String, Function, Array],
    tickHiddenClass: [String, Function, Array],
    tickLabelClass: [String, Function, Array]
  },
  computed: {
    rootClasses() {
      return [
        this.computedClass("tickClass", "o-slide__tick"),
        { [this.computedClass("tickHiddenClass", "o-slide__tick--hidden")]: this.hidden }
      ];
    },
    tickLabelClasses() {
      return [
        this.computedClass("tickLabelClass", "o-slide__tick-label")
      ];
    },
    position() {
      const pos = (this.value - this.$parent.min) / (this.$parent.max - this.$parent.min) * 100;
      return pos >= 0 && pos <= 100 ? pos : 0;
    },
    hidden() {
      return this.value === this.$parent.min || this.value === this.$parent.max;
    },
    tickStyle() {
      return { "left": this.position + "%" };
    }
  },
  created() {
    if (!this.$slider) {
      throw new Error("You should wrap oSliderTick on a oSlider");
    }
  }
});
function render$1$2(_ctx, _cache, $props, $setup, $data, $options) {
  return vue.openBlock(), vue.createBlock("div", {
    class: _ctx.rootClasses,
    style: _ctx.tickStyle
  }, [_ctx.$slots.default ? (vue.openBlock(), vue.createBlock("span", {
    key: 0,
    class: _ctx.tickLabelClasses
  }, [vue.renderSlot(_ctx.$slots, "default")], 2)) : vue.createCommentVNode("v-if", true)], 6);
}
script$1$4.render = render$1$2;
script$1$4.__file = "src/components/slider/SliderTick.vue";
var script$2$2 = vue.defineComponent({
  name: "OSlider",
  components: {
    [script$6.name]: script$6,
    [script$1$4.name]: script$1$4
  },
  configField: "slider",
  mixins: [BaseComponentMixin],
  provide() {
    return {
      $slider: this
    };
  },
  emits: ["update:modelValue", "change", "dragging", "dragstart", "dragend"],
  props: {
    modelValue: {
      type: [Number, Array],
      default: 0
    },
    min: {
      type: Number,
      default: 0
    },
    max: {
      type: Number,
      default: 100
    },
    step: {
      type: Number,
      default: 1
    },
    variant: {
      type: String
    },
    size: String,
    ticks: {
      type: Boolean,
      default: false
    },
    tooltip: {
      type: Boolean,
      default: true
    },
    tooltipVariant: String,
    rounded: {
      type: Boolean,
      default: () => {
        return getValueByPath(getOptions$1(), "slider.rounded", false);
      }
    },
    disabled: {
      type: Boolean,
      default: false
    },
    lazy: {
      type: Boolean,
      default: false
    },
    customFormatter: Function,
    ariaLabel: [String, Array],
    biggerSliderFocus: {
      type: Boolean,
      default: false
    },
    indicator: {
      type: Boolean,
      default: false
    },
    format: {
      type: String,
      default: "raw",
      validator: (value) => {
        return [
          "raw",
          "percent"
        ].indexOf(value) >= 0;
      }
    },
    locale: {
      type: [String, Array],
      default: () => {
        return getValueByPath(getOptions$1(), "locale");
      }
    },
    tooltipAlways: {
      type: Boolean,
      default: false
    },
    rootClass: [String, Function, Array],
    sizeClass: [String, Function, Array],
    trackClass: [String, Function, Array],
    fillClass: [String, Function, Array],
    thumbRoundedClass: [String, Function, Array],
    thumbDraggingClass: [String, Function, Array],
    disabledClass: [String, Function, Array],
    thumbWrapperClass: [String, Function, Array],
    thumbClass: [String, Function, Array],
    variantClass: [String, Function, Array]
  },
  data() {
    return {
      value1: null,
      value2: null,
      dragging: false,
      isRange: false
    };
  },
  computed: {
    rootClasses() {
      return [
        this.computedClass("rootClass", "o-slide"),
        { [this.computedClass("sizeClass", "o-slide--", this.size)]: this.size },
        { [this.computedClass("disabledClass", "o-slide--disabled")]: this.disabled }
      ];
    },
    trackClasses() {
      return [
        this.computedClass("trackClass", "o-slide__track")
      ];
    },
    fillClasses() {
      return [
        this.computedClass("fillClass", "o-slide__fill"),
        { [this.computedClass("variantClass", "o-slide__fill--", this.variant)]: this.variant }
      ];
    },
    thumbClasses() {
      return [
        this.computedClass("thumbClass", "o-slide__thumb"),
        { [this.computedClass("thumbDraggingClass", "o-slide__thumb--dragging")]: this.dragging },
        { [this.computedClass("thumbRoundedClass", "o-slide__thumb--rounded")]: this.rounded }
      ];
    },
    thumbWrapperClasses() {
      return [
        this.computedClass("thumbWrapperClass", "o-slide__thumb-wrapper")
      ];
    },
    newTooltipVariant() {
      return this.tooltipVariant ? this.tooltipVariant : this.variant;
    },
    tickValues() {
      if (!this.ticks || this.min > this.max || this.step === 0)
        return [];
      const result = [];
      for (let i = this.min + this.step; i < this.max; i = i + this.step) {
        result.push(i);
      }
      return result;
    },
    minValue() {
      return Math.min(this.value1, this.value2);
    },
    maxValue() {
      return Math.max(this.value1, this.value2);
    },
    barSize() {
      return this.isRange ? `${100 * (this.maxValue - this.minValue) / (this.max - this.min)}%` : `${100 * (this.value1 - this.min) / (this.max - this.min)}%`;
    },
    barStart() {
      return this.isRange ? `${100 * (this.minValue - this.min) / (this.max - this.min)}%` : "0%";
    },
    precision() {
      const precisions = [this.min, this.max, this.step].map((item) => {
        const decimal = ("" + item).split(".")[1];
        return decimal ? decimal.length : 0;
      });
      return Math.max(...precisions);
    },
    barStyle() {
      return {
        width: this.barSize,
        left: this.barStart
      };
    }
  },
  watch: {
    value1() {
      this.onInternalValueUpdate();
    },
    value2() {
      this.onInternalValueUpdate();
    },
    min() {
      this.setValues(this.value);
    },
    max() {
      this.setValues(this.value);
    },
    modelValue(value) {
      this.setValues(value);
    }
  },
  methods: {
    setValues(newValue) {
      if (this.min > this.max) {
        return;
      }
      if (Array.isArray(newValue)) {
        this.isRange = true;
        const smallValue = typeof newValue[0] !== "number" || isNaN(newValue[0]) ? this.min : Math.min(Math.max(this.min, newValue[0]), this.max);
        const largeValue = typeof newValue[1] !== "number" || isNaN(newValue[1]) ? this.max : Math.max(Math.min(this.max, newValue[1]), this.min);
        this.value1 = this.isThumbReversed ? largeValue : smallValue;
        this.value2 = this.isThumbReversed ? smallValue : largeValue;
      } else {
        this.isRange = false;
        this.value1 = isNaN(newValue) ? this.min : Math.min(this.max, Math.max(this.min, newValue));
        this.value2 = null;
      }
    },
    onInternalValueUpdate() {
      if (this.isRange) {
        this.isThumbReversed = this.value1 > this.value2;
      }
      if (!this.lazy || !this.dragging) {
        this.emitValue("update:modelValue");
      }
      if (this.dragging) {
        this.emitValue("dragging");
      }
    },
    sliderSize() {
      return this.$refs.slider.getBoundingClientRect().width;
    },
    onSliderClick(event) {
      if (this.disabled || this.isTrackClickDisabled)
        return;
      const sliderOffsetLeft = this.$refs.slider.getBoundingClientRect().left;
      const percent = (event.clientX - sliderOffsetLeft) / this.sliderSize() * 100;
      const targetValue = this.min + percent * (this.max - this.min) / 100;
      const diffFirst = Math.abs(targetValue - this.value1);
      if (!this.isRange) {
        if (diffFirst < this.step / 2)
          return;
        this.$refs.button1.setPosition(percent);
      } else {
        const diffSecond = Math.abs(targetValue - this.value2);
        if (diffFirst <= diffSecond) {
          if (diffFirst < this.step / 2)
            return;
          this.$refs["button1"].setPosition(percent);
        } else {
          if (diffSecond < this.step / 2)
            return;
          this.$refs["button2"].setPosition(percent);
        }
      }
      this.emitValue("change");
    },
    onDragStart() {
      this.dragging = true;
      this.$emit("dragstart");
    },
    onDragEnd() {
      this.isTrackClickDisabled = true;
      setTimeout(() => {
        this.isTrackClickDisabled = false;
      }, 0);
      this.dragging = false;
      this.$emit("dragend");
      if (this.lazy) {
        this.emitValue("update:modelValue");
      }
    },
    emitValue(event) {
      const val = this.isRange ? [this.minValue, this.maxValue] : this.value1;
      this.$emit(event, val);
    }
  },
  created() {
    this.isThumbReversed = false;
    this.isTrackClickDisabled = false;
    this.setValues(this.modelValue);
  }
});
function render$2$2(_ctx, _cache, $props, $setup, $data, $options) {
  const _component_o_slider_tick = vue.resolveComponent("o-slider-tick");
  const _component_o_slider_thumb = vue.resolveComponent("o-slider-thumb");
  return vue.openBlock(), vue.createBlock("div", {
    onClick: _cache[3] || (_cache[3] = (...args) => _ctx.onSliderClick(...args)),
    class: _ctx.rootClasses
  }, [vue.createVNode("div", {
    class: _ctx.trackClasses,
    ref: "slider"
  }, [vue.createVNode("div", {
    class: _ctx.fillClasses,
    style: _ctx.barStyle
  }, null, 6), _ctx.ticks ? (vue.openBlock(true), vue.createBlock(vue.Fragment, {
    key: 0
  }, vue.renderList(_ctx.tickValues, (val, key) => {
    return vue.openBlock(), vue.createBlock(_component_o_slider_tick, {
      key,
      value: val
    }, null, 8, ["value"]);
  }), 128)) : vue.createCommentVNode("v-if", true), vue.renderSlot(_ctx.$slots, "default"), vue.createVNode(_component_o_slider_thumb, {
    modelValue: _ctx.value1,
    "onUpdate:modelValue": _cache[1] || (_cache[1] = ($event) => _ctx.value1 = $event),
    variant: _ctx.newTooltipVariant,
    tooltip: _ctx.tooltip,
    "custom-formatter": _ctx.customFormatter,
    indicator: _ctx.indicator,
    ref: "button1",
    role: "slider",
    format: _ctx.format,
    locale: _ctx.locale,
    "tooltip-always": _ctx.tooltipAlways,
    "aria-valuenow": _ctx.value1,
    "aria-valuemin": _ctx.min,
    "aria-valuemax": _ctx.max,
    "aria-orientation": "horizontal",
    "aria-label": Array.isArray(_ctx.ariaLabel) ? _ctx.ariaLabel[0] : _ctx.ariaLabel,
    "aria-disabled": _ctx.disabled,
    onDragstart: _ctx.onDragStart,
    onDragend: _ctx.onDragEnd
  }, null, 8, ["modelValue", "variant", "tooltip", "custom-formatter", "indicator", "format", "locale", "tooltip-always", "aria-valuenow", "aria-valuemin", "aria-valuemax", "aria-label", "aria-disabled", "onDragstart", "onDragend"]), _ctx.isRange ? vue.createVNode(_component_o_slider_thumb, {
    key: 1,
    modelValue: _ctx.value2,
    "onUpdate:modelValue": _cache[2] || (_cache[2] = ($event) => _ctx.value2 = $event),
    variant: _ctx.newTooltipVariant,
    tooltip: _ctx.tooltip,
    "custom-formatter": _ctx.customFormatter,
    indicator: _ctx.indicator,
    ref: "button2",
    role: "slider",
    format: _ctx.format,
    locale: _ctx.locale,
    "tooltip-always": _ctx.tooltipAlways,
    "aria-valuenow": _ctx.value2,
    "aria-valuemin": _ctx.min,
    "aria-valuemax": _ctx.max,
    "aria-orientation": "horizontal",
    "aria-label": Array.isArray(_ctx.ariaLabel) ? _ctx.ariaLabel[1] : "",
    "aria-disabled": _ctx.disabled,
    onDragstart: _ctx.onDragStart,
    onDragend: _ctx.onDragEnd
  }, null, 8, ["modelValue", "variant", "tooltip", "custom-formatter", "indicator", "format", "locale", "tooltip-always", "aria-valuenow", "aria-valuemin", "aria-valuemax", "aria-label", "aria-disabled", "onDragstart", "onDragend"]) : vue.createCommentVNode("v-if", true)], 2)], 2);
}
script$2$2.render = render$2$2;
script$2$2.__file = "src/components/slider/Slider.vue";
var index$7 = {
  install(app) {
    registerComponent(app, script$2$2);
    registerComponent(app, script$1$4);
  }
};
var index$j = index$7;
var SlotComponent = vue.defineComponent({
  name: "OSlotComponent",
  props: {
    component: {
      type: Object,
      required: true
    },
    name: {
      type: String,
      default: "default"
    },
    props: {
      type: Object
    },
    tag: {
      type: String,
      default: "div"
    }
  },
  render() {
    const slot = this.component.$slots[this.name](this.props);
    return vue.h(this.tag, {}, slot);
  }
});
const items = 1;
const sorted = 3;
const Sorted = sorted;
var ProviderParentMixin = (itemName, flags = 0) => {
  const mixin = vue.defineComponent({
    provide() {
      return {
        ["o" + itemName]: this
      };
    }
  });
  if (hasFlag(flags, items)) {
    mixin.data = function() {
      return {
        childItems: [],
        sequence: 1
      };
    };
    mixin.methods = {
      _registerItem(item) {
        this.$nextTick(() => {
          item.index = this.childItems.length;
          this.childItems.push(item);
        });
      },
      _unregisterItem(item) {
        this.$nextTick(() => {
          this.childItems = this.childItems.filter((i) => i !== item);
          let index2 = 0;
          this.childItems.forEach((it) => {
            it.index = index2++;
          });
        });
      },
      _nextSequence() {
        return this.sequence++;
      }
    };
    if (hasFlag(flags, sorted)) {
      mixin.computed = {
        sortedItems() {
          return this.childItems.slice().sort((i1, i2) => {
            return i1.index - i2.index;
          });
        }
      };
    }
  }
  return mixin;
};
var TabbedMixin = (cmp) => vue.defineComponent({
  mixins: [ProviderParentMixin(cmp, Sorted)],
  components: {
    [script$r.name]: script$r,
    [SlotComponent.name]: SlotComponent
  },
  emits: ["update:modelValue"],
  props: {
    modelValue: [String, Number],
    variant: [String, Object],
    size: String,
    animated: {
      type: Boolean,
      default: true
    },
    vertical: {
      type: Boolean,
      default: false
    },
    position: String,
    destroyOnHide: {
      type: Boolean,
      default: false
    }
  },
  data() {
    return {
      activeId: this.modelValue,
      contentHeight: 0,
      isTransitioning: false
    };
  },
  computed: {
    activeItem() {
      return this.activeId !== void 0 && this.activeId !== null ? this.childItems.filter((i) => i.newValue === this.activeId)[0] : this.items[0];
    },
    items() {
      return this.sortedItems;
    }
  },
  watch: {
    modelValue(value) {
      this.performAction();
      this.activeId = value;
      this.performAction();
    }
  },
  methods: {
    childClick(child) {
      if (this.activeId !== child.newValue) {
        this.performAction();
        this.activeId = child.newValue;
        this.performAction();
        this.$emit("update:modelValue", this.activeId);
      }
    },
    performAction() {
      const oldValue = this.activeId;
      const oldTab = oldValue !== void 0 && oldValue !== null ? this.childItems.filter((i) => i.newValue === oldValue)[0] : this.items[0];
      if (oldTab && this.activeItem) {
        oldTab.deactivate(this.activeItem.index);
        this.activeItem.activate(oldTab.index);
      }
    }
  }
});
const sorted$1 = 1;
const optional = 2;
const Sorted$1 = sorted$1;
var InjectedChildMixin = (parentItemName, flags = 0) => {
  const mixin = vue.defineComponent({
    inject: {
      parent: { from: "o" + parentItemName }
    },
    created() {
      this.newValue = typeof this.value === "undefined" ? this.parent._nextSequence() : this.value;
      if (!this.parent) {
        if (!hasFlag(flags, optional)) {
          throw new Error("You should wrap " + this.$options.name + " in a " + parentItemName);
        }
      } else if (this.parent._registerItem) {
        this.parent._registerItem(this);
      }
    },
    beforeUnmount() {
      if (this.parent && this.parent._unregisterItem) {
        this.parent._unregisterItem(this);
      }
    }
  });
  if (hasFlag(flags, sorted$1)) {
    mixin.data = () => {
      return {
        index: null
      };
    };
  }
  return mixin;
};
var TabbedChildMixin = (parentCmp) => vue.defineComponent({
  mixins: [InjectedChildMixin(parentCmp, Sorted$1)],
  props: {
    label: String,
    icon: String,
    iconPack: String,
    visible: {
      type: Boolean,
      default: true
    },
    value: [String, Number],
    headerClass: {
      type: [String, Array, Object]
    }
  },
  data() {
    return {
      transitionName: void 0,
      newValue: this.value
    };
  },
  computed: {
    isActive() {
      return this.parent.activeItem === this;
    },
    elementClasses() {
      return [];
    }
  },
  methods: {
    activate(oldIndex) {
      this.transitionName = this.index < oldIndex ? this.parent.vertical ? "slide-down" : "slide-next" : this.parent.vertical ? "slide-up" : "slide-prev";
    },
    deactivate(newIndex) {
      this.transitionName = newIndex < this.index ? this.parent.vertical ? "slide-down" : "slide-next" : this.parent.vertical ? "slide-up" : "slide-prev";
    }
  },
  render() {
    if (this.parent.destroyOnHide) {
      if (!this.isActive || !this.visible)
        return;
    }
    const vnode = vue.withDirectives(vue.h("div", {
      class: this.elementClasses
    }, this.$slots.default()), [[vue.vShow, this.isActive && this.visible]]);
    if (this.parent.animated) {
      return vue.h(vue.Transition, {
        "name": this.transitionName,
        "onBeforeEnter": () => {
          this.parent.isTransitioning = true;
        },
        "onAfterEnter": () => {
          this.parent.isTransitioning = false;
        }
      }, () => [vnode]);
    }
    return vnode;
  }
});
var script$5 = vue.defineComponent({
  name: "OSteps",
  components: {
    [script$o.name]: script$o,
    [script$r.name]: script$r
  },
  configField: "steps",
  mixins: [BaseComponentMixin, MatchMediaMixin, TabbedMixin("step")],
  props: {
    iconPack: String,
    iconPrev: {
      type: String,
      default: () => {
        return getValueByPath(getOptions$1(), "steps.iconPrev", "chevron-left");
      }
    },
    iconNext: {
      type: String,
      default: () => {
        return getValueByPath(getOptions$1(), "steps.iconNext", "chevron-right");
      }
    },
    hasNavigation: {
      type: Boolean,
      default: true
    },
    animated: {
      type: Boolean,
      default: true
    },
    labelPosition: {
      type: String,
      validator(value) {
        return [
          "bottom",
          "right",
          "left"
        ].indexOf(value) > -1;
      },
      default: "bottom"
    },
    rounded: {
      type: Boolean,
      default: true
    },
    ariaNextLabel: String,
    ariaPreviousLabel: String,
    rootClass: [String, Function, Array],
    sizeClass: [String, Function, Array],
    verticalClass: [String, Function, Array],
    positionClass: [String, Function, Array],
    stepsClass: [String, Function, Array],
    animatedClass: [String, Function, Array],
    stepMarkerRoundedClass: [String, Function, Array],
    stepDividerClass: [String, Function, Array],
    stepMarkerClass: [String, Function, Array],
    stepContentClass: [String, Function, Array],
    stepContentTransitioningClass: [String, Function, Array],
    stepNavigationClass: [String, Function, Array],
    stepLinkClass: [String, Function, Array],
    stepLinkClickableClass: [String, Function, Array],
    stepLinkLabelClass: [String, Function, Array],
    stepLinkLabelPositionClass: [String, Function, Array],
    mobileClass: [String, Function, Array]
  },
  computed: {
    wrapperClasses() {
      return [
        this.computedClass("rootClass", "o-steps__wrapper"),
        { [this.computedClass("sizeClass", "o-steps--", this.size)]: this.size },
        { [this.computedClass("verticalClass", "o-steps__wrapper-vertical")]: this.vertical },
        { [this.computedClass("positionClass", "o-steps__wrapper-position-", this.position)]: this.position && this.vertical },
        { [this.computedClass("mobileClass", "o-steps--mobile")]: this.isMatchMedia }
      ];
    },
    mainClasses() {
      return [
        this.computedClass("stepsClass", "o-steps"),
        { [this.computedClass("animatedClass", "o-steps--animated")]: this.animated }
      ];
    },
    stepDividerClasses() {
      return [
        this.computedClass("stepDividerClass", "o-steps__divider")
      ];
    },
    stepMarkerClasses() {
      return [
        this.computedClass("stepMarkerClass", "o-steps__marker"),
        { [this.computedClass("stepMarkerRoundedClass", "o-steps__marker--rounded")]: this.rounded }
      ];
    },
    stepContentClasses() {
      return [
        this.computedClass("stepContentClass", "o-steps__content"),
        { [this.computedClass("stepContentTransitioningClass", "o-steps__content-transitioning")]: this.isTransitioning }
      ];
    },
    stepNavigationClasses() {
      return [
        this.computedClass("stepNavigationClass", "o-steps__navigation")
      ];
    },
    stepLinkLabelClasses() {
      return [
        this.computedClass("stepLinkLabelClass", "o-steps__title")
      ];
    },
    activeItem() {
      return this.childItems.filter((i) => i.newValue === this.activeId)[0] || this.items[0];
    },
    hasPrev() {
      return !!this.prevItem;
    },
    nextItem() {
      let nextItem = null;
      let idx = this.activeItem ? this.items.indexOf(this.activeItem) + 1 : 0;
      for (; idx < this.items.length; idx++) {
        if (this.items[idx].visible) {
          nextItem = this.items[idx];
          break;
        }
      }
      return nextItem;
    },
    prevItem() {
      if (!this.activeItem) {
        return null;
      }
      let prevItem = null;
      for (let idx = this.items.indexOf(this.activeItem) - 1; idx >= 0; idx--) {
        if (this.items[idx].visible) {
          prevItem = this.items[idx];
          break;
        }
      }
      return prevItem;
    },
    hasNext() {
      return !!this.nextItem;
    },
    navigationProps() {
      return {
        previous: {
          disabled: !this.hasPrev,
          action: this.prev
        },
        next: {
          disabled: !this.hasNext,
          action: this.next
        }
      };
    }
  },
  methods: {
    stepLinkClasses(childItem) {
      return [
        this.computedClass("stepLinkClass", "o-steps__link"),
        { [this.computedClass("stepLinkLabelPositionClass", "o-steps__link-label-", this.labelPosition)]: this.labelPosition },
        { [this.computedClass("stepLinkClickableClass", "o-steps__link-clickable")]: this.isItemClickable(childItem) }
      ];
    },
    isItemClickable(stepItem) {
      if (stepItem.clickable === void 0) {
        return stepItem.index < this.activeItem.index;
      }
      return stepItem.clickable;
    },
    prev() {
      if (this.hasPrev) {
        this.childClick(this.prevItem);
      }
    },
    next() {
      if (this.hasNext) {
        this.childClick(this.nextItem);
      }
    }
  }
});
const _hoisted_1$2 = {
  key: 1
};
function render$4(_ctx, _cache, $props, $setup, $data, $options) {
  const _component_o_icon = vue.resolveComponent("o-icon");
  const _component_o_button = vue.resolveComponent("o-button");
  return vue.openBlock(), vue.createBlock("div", {
    class: _ctx.wrapperClasses
  }, [vue.createVNode("nav", {
    class: _ctx.mainClasses
  }, [(vue.openBlock(true), vue.createBlock(vue.Fragment, null, vue.renderList(_ctx.items, (childItem, index2) => {
    return vue.withDirectives((vue.openBlock(), vue.createBlock("div", {
      key: childItem.newValue,
      class: childItem.itemClasses
    }, [index2 > 0 ? (vue.openBlock(), vue.createBlock("span", {
      key: 0,
      class: _ctx.stepDividerClasses
    }, null, 2)) : vue.createCommentVNode("v-if", true), vue.createVNode("a", {
      class: _ctx.stepLinkClasses(childItem),
      onClick: ($event) => _ctx.isItemClickable(childItem) && _ctx.childClick(childItem)
    }, [vue.createVNode("div", {
      class: _ctx.stepMarkerClasses
    }, [childItem.icon ? vue.createVNode(_component_o_icon, {
      key: 0,
      icon: childItem.icon,
      pack: childItem.iconPack,
      size: _ctx.size
    }, null, 8, ["icon", "pack", "size"]) : childItem.step ? (vue.openBlock(), vue.createBlock("span", _hoisted_1$2, vue.toDisplayString(childItem.step), 1)) : vue.createCommentVNode("v-if", true)], 2), vue.createVNode("div", {
      class: _ctx.stepLinkLabelClasses
    }, vue.toDisplayString(childItem.label), 3)], 10, ["onClick"])], 2)), [[vue.vShow, childItem.visible]]);
  }), 128))], 2), vue.createVNode("section", {
    class: _ctx.stepContentClasses
  }, [vue.renderSlot(_ctx.$slots, "default")], 2), vue.renderSlot(_ctx.$slots, "navigation", {
    previous: _ctx.navigationProps.previous,
    next: _ctx.navigationProps.next
  }, () => [_ctx.hasNavigation ? (vue.openBlock(), vue.createBlock("nav", {
    key: 0,
    class: _ctx.stepNavigationClasses
  }, [vue.createVNode(_component_o_button, {
    role: "button",
    "icon-left": _ctx.iconPrev,
    "icon-pack": _ctx.iconPack,
    "icon-both": "",
    disabled: _ctx.navigationProps.previous.disabled,
    onClick: vue.withModifiers(_ctx.navigationProps.previous.action, ["prevent"]),
    "aria-label": _ctx.ariaPreviousLabel
  }, null, 8, ["icon-left", "icon-pack", "disabled", "onClick", "aria-label"]), vue.createVNode(_component_o_button, {
    role: "button",
    "icon-left": _ctx.iconNext,
    "icon-pack": _ctx.iconPack,
    "icon-both": "",
    disabled: _ctx.navigationProps.next.disabled,
    onClick: vue.withModifiers(_ctx.navigationProps.next.action, ["prevent"]),
    "aria-label": _ctx.ariaNextLabel
  }, null, 8, ["icon-left", "icon-pack", "disabled", "onClick", "aria-label"])], 2)) : vue.createCommentVNode("v-if", true)])], 2);
}
script$5.render = render$4;
script$5.__file = "src/components/steps/Steps.vue";
var script$1$3 = vue.defineComponent({
  name: "OStepItem",
  mixins: [BaseComponentMixin, TabbedChildMixin("step")],
  configField: "steps",
  props: {
    step: [String, Number],
    variant: [String, Object],
    clickable: {
      type: Boolean,
      default: void 0
    },
    itemClass: [String, Function, Array],
    itemHeaderClass: [String, Function, Array],
    itemHeaderActiveClass: [String, Function, Array],
    itemHeaderPreviousClass: [String, Function, Array],
    itemHeaderVariantClass: [String, Function, Array]
  },
  computed: {
    elementClasses() {
      return [this.computedClass("itemClass", "o-steps__item")];
    },
    itemClasses() {
      return [this.headerClass, this.computedClass("itemHeaderClass", "o-steps__nav-item"), {
        [this.computedClass("itemHeaderVariantClass", "o-steps__nav-item--", this.variant || this.parent.variant)]: this.variant || this.parent.variant
      }, {
        [this.computedClass("itemHeaderActiveClass", "o-steps__nav-item-active")]: this.isActive
      }, {
        [this.computedClass("itemHeaderPreviousClass", "o-steps__nav-item-previous")]: this.parent.activeItem.index > this.index
      }];
    }
  }
});
script$1$3.__file = "src/components/steps/StepItem.vue";
var index$6 = {
  install(app) {
    registerComponent(app, script$5);
    registerComponent(app, script$1$3);
  }
};
var index$k = index$6;
var script$4 = vue.defineComponent({
  name: "OSwitch",
  mixins: [BaseComponentMixin],
  configField: "switch",
  emits: ["update:modelValue"],
  props: {
    modelValue: [String, Number, Boolean],
    nativeValue: [String, Number, Boolean],
    disabled: Boolean,
    variant: String,
    passiveVariant: String,
    name: String,
    required: Boolean,
    size: String,
    trueValue: {
      type: [String, Number, Boolean],
      default: true
    },
    falseValue: {
      type: [String, Number, Boolean],
      default: false
    },
    rounded: {
      type: Boolean,
      default: true
    },
    leftLabel: {
      type: Boolean,
      default: false
    },
    ariaLabelledby: String,
    rootClass: [String, Function, Array],
    disabledClass: [String, Function, Array],
    checkClass: [String, Function, Array],
    checkCheckedClass: [String, Function, Array],
    checkSwitchClass: [String, Function, Array],
    roundedClass: [String, Function, Array],
    labelClass: [String, Function, Array],
    sizeClass: [String, Function, Array],
    variantClass: [String, Function, Array],
    elementsWrapperClass: [String, Function, Array],
    passiveVariantClass: [String, Function, Array],
    leftLabelClass: [String, Function, Array]
  },
  data() {
    return {
      newValue: this.modelValue,
      isMouseDown: false
    };
  },
  computed: {
    rootClasses() {
      return [
        this.computedClass("rootClass", "o-switch"),
        { [this.computedClass("sizeClass", "o-switch--", this.size)]: this.size },
        { [this.computedClass("disabledClass", "o-switch--disabled")]: this.disabled },
        { [this.computedClass("variantClass", "o-switch--", this.variant)]: this.variant },
        { [this.computedClass("passiveVariantClass", "o-switch--", this.passiveVariant + "-passive")]: this.passiveVariant }
      ];
    },
    checkClasses() {
      return [
        this.computedClass("checkClass", "o-switch__check"),
        { [this.computedClass("checkCheckedClass", "o-switch__check--checked")]: this.newValue === this.trueValue },
        { [this.computedClass("roundedClass", "o-switch--rounded")]: this.rounded }
      ];
    },
    elementsWrapperClasses() {
      return [
        this.computedClass("elementsWrapperClass", "o-switch__wrapper"),
        { [this.computedClass("leftLabelClass", "o-switch__wrapper--left")]: this.leftLabel }
      ];
    },
    checkSwitchClasses() {
      return [
        this.computedClass("checkSwitchClass", "o-switch__check-switch"),
        { [this.computedClass("roundedClass", "o-switch--rounded")]: this.rounded }
      ];
    },
    labelClasses() {
      return [
        this.computedClass("labelClass", "o-switch__label")
      ];
    },
    computedValue: {
      get() {
        return this.newValue;
      },
      set(value) {
        this.newValue = value;
        this.$emit("update:modelValue", this.newValue);
      }
    }
  },
  watch: {
    modelValue(value) {
      this.newValue = value;
    }
  },
  methods: {
    focus() {
      this.$refs.input.focus();
    }
  }
});
function render$3(_ctx, _cache, $props, $setup, $data, $options) {
  return vue.openBlock(), vue.createBlock("label", {
    class: _ctx.rootClasses,
    ref: "label",
    onClick: _cache[3] || (_cache[3] = (...args) => _ctx.focus(...args)),
    onKeydown: _cache[4] || (_cache[4] = vue.withKeys(vue.withModifiers(($event) => _ctx.$refs.label.click(), ["prevent"]), ["enter"])),
    onMousedown: _cache[5] || (_cache[5] = ($event) => _ctx.isMouseDown = true),
    onMouseup: _cache[6] || (_cache[6] = ($event) => _ctx.isMouseDown = false),
    onMouseout: _cache[7] || (_cache[7] = ($event) => _ctx.isMouseDown = false),
    onBlur: _cache[8] || (_cache[8] = ($event) => _ctx.isMouseDown = false)
  }, [vue.withDirectives(vue.createVNode("input", {
    "onUpdate:modelValue": _cache[1] || (_cache[1] = ($event) => _ctx.computedValue = $event),
    type: "checkbox",
    ref: "input",
    onClick: _cache[2] || (_cache[2] = vue.withModifiers(() => {
    }, ["stop"])),
    disabled: _ctx.disabled,
    name: _ctx.name,
    required: _ctx.required,
    value: _ctx.nativeValue,
    "true-value": _ctx.trueValue,
    "false-value": _ctx.falseValue,
    "aria-labelledby": _ctx.ariaLabelledby
  }, null, 8, ["disabled", "name", "required", "value", "true-value", "false-value", "aria-labelledby"]), [[vue.vModelCheckbox, _ctx.computedValue]]), vue.createVNode("span", {
    class: _ctx.elementsWrapperClasses
  }, [vue.createVNode("span", {
    class: _ctx.checkClasses
  }, [vue.createVNode("span", {
    class: _ctx.checkSwitchClasses
  }, null, 2)], 2), vue.createVNode("span", {
    id: _ctx.ariaLabelledby,
    class: _ctx.labelClasses
  }, [vue.renderSlot(_ctx.$slots, "default")], 10, ["id"])], 2)], 34);
}
script$4.render = render$3;
script$4.__file = "src/components/switch/Switch.vue";
var index$5 = {
  install(app) {
    registerComponent(app, script$4);
  }
};
var index$l = index$5;
var script$2 = vue.defineComponent({
  name: "OTableMobileSort",
  components: {
    [script$o.name]: script$o,
    [script$j.name]: script$j,
    [script$r.name]: script$r,
    [script$1$8.name]: script$1$8
  },
  inject: ["$table"],
  emits: ["sort"],
  props: {
    currentSortColumn: Object,
    columns: Array,
    placeholder: String,
    iconPack: String,
    sortIcon: {
      type: String,
      default: "arrow-up"
    },
    sortIconSize: {
      type: String,
      default: "small"
    },
    isAsc: Boolean
  },
  data() {
    return {
      mobileSort: getValueByPath(this.currentSortColumn, "newKey"),
      defaultEvent: {
        shiftKey: true,
        altKey: true,
        ctrlKey: true
      },
      ignoreSort: false
    };
  },
  computed: {
    showPlaceholder() {
      return !this.columns || !this.columns.some((column) => getValueByPath(column, "newKey") === this.mobileSort);
    },
    sortableColumns() {
      if (!this.columns)
        return [];
      return this.columns.filter((c) => c.sortable);
    },
    isCurrentSort() {
      return getValueByPath(this.currentSortColumn, "newKey") === this.mobileSort;
    }
  },
  watch: {
    mobileSort(value) {
      if (this.currentSortColumn.newKey === value)
        return;
      const column = this.sortableColumns.filter((c) => getValueByPath(c, "newKey") === value)[0];
      this.$emit("sort", column, this.defaultEvent);
    },
    currentSortColumn(column) {
      this.mobileSort = getValueByPath(column, "newKey");
    }
  },
  methods: {
    sort() {
      const column = this.sortableColumns.filter((c) => getValueByPath(c, "newKey") === this.mobileSort)[0];
      this.$emit("sort", column, this.defaultEvent);
    }
  }
});
function render$2(_ctx, _cache, $props, $setup, $data, $options) {
  const _component_o_select = vue.resolveComponent("o-select");
  const _component_o_icon = vue.resolveComponent("o-icon");
  const _component_o_button = vue.resolveComponent("o-button");
  const _component_o_field = vue.resolveComponent("o-field");
  return vue.openBlock(), vue.createBlock("div", {
    class: _ctx.$table.mobileSortClasses
  }, [vue.createVNode(_component_o_field, null, {
    default: vue.withCtx(() => [vue.createVNode(_component_o_select, {
      modelValue: _ctx.mobileSort,
      "onUpdate:modelValue": _cache[1] || (_cache[1] = ($event) => _ctx.mobileSort = $event),
      expanded: ""
    }, {
      default: vue.withCtx(() => [_ctx.placeholder ? vue.withDirectives((vue.openBlock(), vue.createBlock("option", {
        key: 0,
        value: {},
        selected: "",
        disabled: "",
        hidden: ""
      }, vue.toDisplayString(_ctx.placeholder), 513)), [[vue.vShow, _ctx.showPlaceholder]]) : vue.createCommentVNode("v-if", true), (vue.openBlock(true), vue.createBlock(vue.Fragment, null, vue.renderList(_ctx.sortableColumns, (column, index2) => {
        return vue.openBlock(), vue.createBlock("option", {
          key: index2,
          value: column.newKey
        }, vue.toDisplayString(column.label), 9, ["value"]);
      }), 128))]),
      _: 1
    }, 8, ["modelValue"]), vue.createVNode(_component_o_button, {
      onClick: _ctx.sort
    }, {
      default: vue.withCtx(() => [vue.withDirectives(vue.createVNode(_component_o_icon, {
        icon: _ctx.sortIcon,
        pack: _ctx.iconPack,
        size: _ctx.sortIconSize,
        both: "",
        rotation: !_ctx.isAsc ? 180 : 0
      }, null, 8, ["icon", "pack", "size", "rotation"]), [[vue.vShow, _ctx.isCurrentSort]])]),
      _: 1
    }, 8, ["onClick"])]),
    _: 1
  })], 2);
}
script$2.render = render$2;
script$2.__file = "src/components/table/TableMobileSort.vue";
var script$1$2 = vue.defineComponent({
  name: "OTableColumn",
  inject: ["$table"],
  props: {
    label: String,
    customKey: [String, Number],
    field: String,
    meta: [String, Number, Boolean, Function, Object, Array],
    width: [Number, String],
    numeric: Boolean,
    position: {
      type: String,
      validator(value) {
        return [
          "left",
          "centered",
          "right"
        ].indexOf(value) > -1;
      }
    },
    searchable: Boolean,
    sortable: Boolean,
    visible: {
      type: Boolean,
      default: true
    },
    customSort: Function,
    customSearch: Function,
    sticky: Boolean,
    headerSelectable: Boolean,
    thAttrs: {
      type: Function,
      default: () => ({})
    },
    tdAttrs: {
      type: Function,
      default: () => ({})
    }
  },
  data() {
    return {
      newKey: void 0,
      _isTableColumn: true
    };
  },
  computed: {
    style() {
      return {
        width: toCssDimension(this.width)
      };
    },
    hasDefaultSlot() {
      return this.$slots.default;
    },
    hasSearchableSlot() {
      return this.$slots.searchable;
    },
    hasHeaderSlot() {
      return this.$slots.header;
    },
    isHeaderUnselectable() {
      return !this.headerSelectable && this.sortable;
    }
  },
  created() {
    if (!this.$table) {
      throw new Error("You should wrap oTableColumn on a oTable");
    }
    this.newKey = this.$table._nextSequence();
    this.$table._addColumn(this);
  },
  beforeUnmount() {
    this.$table._removeColumn(this);
  },
  render() {
    return vue.h("span", { "data-id": this.newKey }, this.label);
  }
});
script$1$2.__file = "src/components/table/TableColumn.vue";
var script$2$1 = vue.defineComponent({
  name: "OTablePagination",
  components: {
    [script$1$5.name]: script$1$5
  },
  emits: ["update:currentPage", "page-change"],
  props: {
    paginated: Boolean,
    total: [Number, String],
    perPage: [Number, String],
    currentPage: [Number, String],
    paginationSimple: Boolean,
    paginationSize: String,
    rounded: Boolean,
    iconPack: String,
    rootClass: [String, Array, Object],
    ariaNextLabel: String,
    ariaPreviousLabel: String,
    ariaPageLabel: String,
    ariaCurrentLabel: String
  },
  data() {
    return {
      newCurrentPage: this.currentPage
    };
  },
  watch: {
    currentPage(newVal) {
      this.newCurrentPage = newVal;
    }
  },
  methods: {
    pageChanged(page) {
      this.newCurrentPage = page > 0 ? page : 1;
      this.$emit("update:currentPage", this.newCurrentPage);
      this.$emit("page-change", this.newCurrentPage);
    }
  }
});
const _hoisted_1$1 = {
  key: 0
};
function render$1$1(_ctx, _cache, $props, $setup, $data, $options) {
  const _component_o_pagination = vue.resolveComponent("o-pagination");
  return vue.openBlock(), vue.createBlock("div", {
    class: _ctx.rootClass
  }, [vue.createVNode("div", null, [vue.renderSlot(_ctx.$slots, "default")]), vue.createVNode("div", null, [_ctx.paginated ? (vue.openBlock(), vue.createBlock("div", _hoisted_1$1, [vue.createVNode(_component_o_pagination, {
    "icon-pack": _ctx.iconPack,
    total: _ctx.total,
    "per-page": _ctx.perPage,
    simple: _ctx.paginationSimple,
    size: _ctx.paginationSize,
    current: _ctx.newCurrentPage,
    rounded: _ctx.rounded,
    onChange: _ctx.pageChanged,
    "aria-next-label": _ctx.ariaNextLabel,
    "aria-previous-label": _ctx.ariaPreviousLabel,
    "aria-page-label": _ctx.ariaPageLabel,
    "aria-current-label": _ctx.ariaCurrentLabel
  }, null, 8, ["icon-pack", "total", "per-page", "simple", "size", "current", "rounded", "onChange", "aria-next-label", "aria-previous-label", "aria-page-label", "aria-current-label"])])) : vue.createCommentVNode("v-if", true)])], 2);
}
script$2$1.render = render$1$1;
script$2$1.__file = "src/components/table/TablePagination.vue";
var script$3 = vue.defineComponent({
  name: "OTable",
  components: {
    [script$o.name]: script$o,
    [script$n.name]: script$n,
    [script$r.name]: script$r,
    [script$q.name]: script$q,
    [script$e.name]: script$e,
    [SlotComponent.name]: SlotComponent,
    [script$2.name]: script$2,
    [script$1$2.name]: script$1$2,
    [script$2$1.name]: script$2$1
  },
  mixins: [BaseComponentMixin, MatchMediaMixin],
  configField: "table",
  inheritAttrs: false,
  provide() {
    return {
      $table: this
    };
  },
  emits: [
    "page-change",
    "click",
    "dblclick",
    "contextmenu",
    "check",
    "check-all",
    "update:checkedRows",
    "select",
    "update:selected",
    "filters-change",
    "details-close",
    "update:openedDetailed",
    "mouseenter",
    "mouseleave",
    "sort",
    "sorting-priority-removed",
    "dragstart",
    "dragend",
    "drop",
    "dragleave",
    "dragover",
    "cell-click"
  ],
  props: {
    data: {
      type: Array,
      default: () => []
    },
    columns: {
      type: Array,
      default: () => []
    },
    bordered: Boolean,
    striped: Boolean,
    narrowed: Boolean,
    hoverable: Boolean,
    loading: Boolean,
    detailed: Boolean,
    checkable: Boolean,
    headerCheckable: {
      type: Boolean,
      default: true
    },
    checkboxPosition: {
      type: String,
      default: "left",
      validator: (value) => {
        return [
          "left",
          "right"
        ].indexOf(value) >= 0;
      }
    },
    selected: Object,
    isRowSelectable: {
      type: Function,
      default: () => true
    },
    focusable: Boolean,
    customIsChecked: Function,
    isRowCheckable: {
      type: Function,
      default: () => true
    },
    checkedRows: {
      type: Array,
      default: () => []
    },
    mobileCards: {
      type: Boolean,
      default: () => {
        return getValueByPath(getOptions$1(), "table.mobileCards", true);
      }
    },
    defaultSort: [String, Array],
    defaultSortDirection: {
      type: String,
      default: "asc"
    },
    sortIcon: {
      type: String,
      default: () => {
        return getValueByPath(getOptions$1(), "table.sortIcon", "arrow-up");
      }
    },
    sortIconSize: {
      type: String,
      default: () => {
        return getValueByPath(getOptions$1(), "table.sortIconSize", "small");
      }
    },
    paginated: Boolean,
    currentPage: {
      type: Number,
      default: 1
    },
    perPage: {
      type: [Number, String],
      default: () => {
        return getValueByPath(getOptions$1(), "table.perPage", 20);
      }
    },
    showDetailIcon: {
      type: Boolean,
      default: true
    },
    detailIcon: {
      type: String,
      default: "chevron-right"
    },
    paginationPosition: {
      type: String,
      default: () => {
        return getValueByPath(getOptions$1(), "table.paginationPosition", "bottom");
      },
      validator: (value) => {
        return [
          "bottom",
          "top",
          "both"
        ].indexOf(value) >= 0;
      }
    },
    backendSorting: Boolean,
    backendFiltering: Boolean,
    rowClass: {
      type: Function,
      default: () => ""
    },
    openedDetailed: {
      type: Array,
      default: () => []
    },
    hasDetailedVisible: {
      type: Function,
      default: () => true
    },
    detailKey: {
      type: String,
      default: ""
    },
    customDetailRow: {
      type: Boolean,
      default: false
    },
    detailTransition: {
      type: String,
      default: ""
    },
    backendPagination: Boolean,
    total: {
      type: [Number, String],
      default: 0
    },
    iconPack: String,
    mobileSortPlaceholder: String,
    customRowKey: String,
    draggable: {
      type: Boolean,
      default: false
    },
    scrollable: Boolean,
    ariaNextLabel: String,
    ariaPreviousLabel: String,
    ariaPageLabel: String,
    ariaCurrentLabel: String,
    stickyHeader: Boolean,
    height: [Number, String],
    filtersEvent: {
      type: String,
      default: ""
    },
    debounceSearch: Number,
    showHeader: {
      type: Boolean,
      default: () => {
        return getValueByPath(getOptions$1(), "table.showHeader", true);
      }
    },
    stickyCheckbox: {
      type: Boolean,
      default: false
    },
    paginationRounded: Boolean,
    tableClass: [String, Function, Array],
    wrapperClass: [String, Function, Array],
    footerClass: [String, Function, Array],
    emptyClass: [String, Function, Array],
    detailedClass: [String, Function, Array],
    borderedClass: [String, Function, Array],
    stripedClass: [String, Function, Array],
    narrowedClass: [String, Function, Array],
    hoverableClass: [String, Function, Array],
    thClass: [String, Function, Array],
    tdClass: [String, Function, Array],
    thPositionClass: [String, Function, Array],
    thStickyClass: [String, Function, Array],
    thCheckboxClass: [String, Function, Array],
    thCurrentSortClass: [String, Function, Array],
    thSortableClass: [String, Function, Array],
    thUnselectableClass: [String, Function, Array],
    thSortIconClass: [String, Function, Array],
    thDetailedClass: [String, Function, Array],
    tdPositionClass: [String, Function, Array],
    tdStickyClass: [String, Function, Array],
    tdCheckboxClass: [String, Function, Array],
    tdDetailedChevronClass: [String, Function, Array],
    trSelectedClass: [String, Function, Array],
    stickyHeaderClass: [String, Function, Array],
    scrollableClass: [String, Function, Array],
    mobileSortClass: [String, Function, Array],
    paginationWrapperClass: [String, Function, Array],
    mobileClass: [String, Function, Array]
  },
  data() {
    return {
      getValueByPath,
      visibleDetailRows: this.openedDetailed,
      newData: this.data,
      newDataTotal: this.backendPagination ? this.total : this.data.length,
      newCheckedRows: [...this.checkedRows],
      lastCheckedRowIndex: null,
      newCurrentPage: this.currentPage,
      currentSortColumn: {},
      isAsc: true,
      filters: {},
      defaultSlots: [],
      firstTimeSort: true,
      sequence: 1
    };
  },
  mounted() {
    this.$nextTick(() => {
      this.checkSort();
    });
  },
  computed: {
    tableClasses() {
      return [
        this.computedClass("tableClass", "o-table"),
        { [this.computedClass("borderedClass", "o-table--bordered")]: this.bordered },
        { [this.computedClass("stripedClass", "o-table--striped")]: this.striped },
        { [this.computedClass("narrowedClass", "o-table--narrowed")]: this.narrowed },
        { [this.computedClass("hoverableClass", "o-table--hoverable")]: (this.hoverable || this.focusable) && this.visibleData.length },
        { [this.computedClass("emptyClass", "o-table--table__empty")]: !this.visibleData.length }
      ];
    },
    tableWrapperClasses() {
      return [
        this.computedClass("wrapperClass", "o-table__wrapper"),
        { [this.computedClass("stickyHeaderClass", "o-table__wrapper--sticky-header")]: this.stickyHeader },
        { [this.computedClass("scrollableClass", "o-table__wrapper--scrollable")]: this.isScrollable },
        { [this.computedClass("mobileClass", "o-table__wrapper--mobile")]: this.isMobile }
      ];
    },
    footerClasses() {
      return [
        this.computedClass("footerClass", "o-table__footer")
      ];
    },
    thBaseClasses() {
      return [
        this.computedClass("thClass", "o-table__th")
      ];
    },
    tdBaseClasses() {
      return [
        this.computedClass("tdClass", "o-table__td")
      ];
    },
    thCheckboxClasses() {
      return [
        ...this.thBaseClasses,
        this.computedClass("thCheckboxClass", "o-table__th-checkbox")
      ];
    },
    thDetailedClasses() {
      return [
        ...this.thBaseClasses,
        this.computedClass("thDetailedClass", "o-table__th--detailed")
      ];
    },
    tdCheckboxClasses() {
      return [
        ...this.tdBaseClasses,
        this.computedClass("tdCheckboxClass", "o-table__td-checkbox"),
        ...this.thStickyClasses({ sticky: this.stickyCheckbox })
      ];
    },
    detailedClasses() {
      return [
        this.computedClass("detailedClass", "o-table__detail")
      ];
    },
    tdDetailedChevronClasses() {
      return [
        ...this.tdBaseClasses,
        this.computedClass("tdDetailedChevronClass", "o-table__td-chevron")
      ];
    },
    mobileSortClasses() {
      return [
        this.computedClass("mobileSortClass", "o-table__mobile-sort")
      ];
    },
    paginationWrapperClasses() {
      return [
        this.computedClass("paginationWrapperClass", "o-table__pagination")
      ];
    },
    tableWrapperStyle() {
      return {
        height: toCssDimension(this.height)
      };
    },
    visibleData() {
      if (!this.paginated)
        return this.newData;
      const currentPage = this.newCurrentPage;
      const perPage = this.perPage;
      if (this.newData.length <= perPage) {
        return this.newData;
      } else {
        const start = (currentPage - 1) * perPage;
        const end = start + parseInt(perPage, 10);
        return this.newData.slice(start, end);
      }
    },
    visibleColumns() {
      if (!this.newColumns)
        return this.newColumns;
      return this.newColumns.filter((column) => {
        return column.visible || column.visible === void 0;
      });
    },
    isAllChecked() {
      const validVisibleData = this.visibleData.filter((row) => this.isRowCheckable(row));
      if (validVisibleData.length === 0)
        return false;
      const isAllChecked = validVisibleData.some((currentVisibleRow) => {
        return indexOf(this.newCheckedRows, currentVisibleRow, this.customIsChecked) < 0;
      });
      return !isAllChecked;
    },
    isAllUncheckable() {
      const validVisibleData = this.visibleData.filter((row) => this.isRowCheckable(row));
      return validVisibleData.length === 0;
    },
    hasSortablenewColumns() {
      return this.newColumns.some((column) => {
        return column.sortable;
      });
    },
    hasSearchablenewColumns() {
      return this.newColumns.some((column) => {
        return column.searchable;
      });
    },
    columnCount() {
      let count = this.visibleColumns.length;
      count += this.checkable ? 1 : 0;
      count += this.detailed && this.showDetailIcon ? 1 : 0;
      return count;
    },
    showDetailRowIcon() {
      return this.detailed && this.showDetailIcon;
    },
    isScrollable() {
      if (this.scrollable)
        return true;
      if (!this.newColumns)
        return false;
      return this.newColumns.some((column) => {
        return column.sticky;
      });
    },
    newColumns() {
      if (this.columns && this.columns.length) {
        return this.columns.map((column) => {
          const vnode = vue.createVNode(script$1$2, column, (props) => {
            const vnode2 = vue.h("span", {}, getValueByPath(props.row, column.field));
            return [vnode2];
          });
          return vue.createApp(vnode).provide("$table", this).mount(document.createElement("div"));
        });
      }
      let defaultSlots = this.defaultSlots.filter((vnode) => vnode && vnode.$data && vnode.$data._isTableColumn);
      return defaultSlots;
    },
    isMobile() {
      return this.mobileCards && this.isMatchMedia;
    }
  },
  watch: {
    data: {
      handler(value) {
        this.newData = value;
        if (!this.backendFiltering) {
          this.newData = value.filter((row) => this.isRowFiltered(row));
        }
        if (!this.backendSorting) {
          this.sort(this.currentSortColumn, true);
        }
        if (!this.backendPagination) {
          this.newDataTotal = this.newData.length;
        }
      },
      deep: true
    },
    total(newTotal) {
      if (!this.backendPagination)
        return;
      this.newDataTotal = newTotal;
    },
    currentPage(newValue) {
      this.newCurrentPage = newValue;
    },
    checkedRows: {
      handler(rows) {
        this.newCheckedRows = [...rows];
      },
      deep: true
    },
    debounceSearch: {
      handler(value) {
        this.debouncedHandleFiltersChange = debounce(this.handleFiltersChange, value);
      },
      immediate: true
    },
    filters: {
      handler(value) {
        if (this.debounceSearch) {
          this.debouncedHandleFiltersChange(value);
        } else {
          this.handleFiltersChange(value);
        }
      },
      deep: true
    },
    openedDetailed(expandedRows) {
      this.visibleDetailRows = expandedRows;
    },
    newCurrentPage(newVal) {
      this.$emit("update:currentPage", newVal);
    }
  },
  methods: {
    thClasses(column) {
      return [
        ...this.thBaseClasses,
        ...this.thStickyClasses(column),
        column.thAttrs && getValueByPath(column.thAttrs(column), "class"),
        { [this.computedClass("thCurrentSortClass", "o-table__th-current-sort")]: this.currentSortColumn === column },
        { [this.computedClass("thSortableClass", "o-table__th--sortable")]: column.sortable },
        { [this.computedClass("thUnselectableClass", "o-table__th--unselectable")]: column.isHeaderUnselectable },
        { [this.computedClass("thPositionClass", "o-table__th--", column.position)]: column.position }
      ];
    },
    thStickyClasses(column) {
      return [
        { [this.computedClass("thStickyClass", "o-table__th--sticky")]: column.sticky }
      ];
    },
    rowClasses(row, index2) {
      return [
        this.rowClass(row, index2),
        { [this.computedClass("trSelectedClass", "o-table__tr--selected")]: this.isRowSelected(row, this.selected) }
      ];
    },
    thSortIconClasses() {
      return [
        this.computedClass("thSortIconClass", "o-table__th__sort-icon")
      ];
    },
    tdClasses(row, column) {
      return [
        ...this.tdBaseClasses,
        column.tdAttrs && getValueByPath(column.tdAttrs(row, column), "class"),
        { [this.computedClass("tdPositionClass", "o-table__td--", column.position)]: column.position },
        { [this.computedClass("tdStickyClass", "o-table__td--sticky")]: column.sticky }
      ];
    },
    onFiltersEvent(event) {
      this.$emit(`filters-event-${this.filtersEvent}`, { event, filters: this.filters });
    },
    handleFiltersChange(value) {
      if (this.backendFiltering) {
        this.$emit("filters-change", value);
      } else {
        this.newData = this.data.filter((row) => this.isRowFiltered(row));
        if (!this.backendPagination) {
          this.newDataTotal = this.newData.length;
        }
        if (!this.backendSorting) {
          if (Object.keys(this.currentSortColumn).length > 0) {
            this.doSortSingleColumn(this.currentSortColumn);
          }
        }
      }
    },
    sortBy(array, key, fn, isAsc) {
      let sorted2 = [];
      if (fn && typeof fn === "function") {
        sorted2 = [...array].sort((a, b) => fn(a, b, isAsc));
      } else {
        sorted2 = [...array].sort((a, b) => {
          let newA = getValueByPath(a, key);
          let newB = getValueByPath(b, key);
          if (typeof newA === "boolean" && typeof newB === "boolean") {
            return isAsc ? newA - newB : newB - newA;
          }
          if (!newA && newA !== 0)
            return 1;
          if (!newB && newB !== 0)
            return -1;
          if (newA === newB)
            return 0;
          newA = typeof newA === "string" ? newA.toUpperCase() : newA;
          newB = typeof newB === "string" ? newB.toUpperCase() : newB;
          return isAsc ? newA > newB ? 1 : -1 : newA > newB ? -1 : 1;
        });
      }
      return sorted2;
    },
    sort(column, updatingData = false, event = null) {
      if (!column || !column.sortable)
        return;
      if (!updatingData) {
        this.isAsc = column === this.currentSortColumn ? !this.isAsc : this.defaultSortDirection.toLowerCase() !== "desc";
      }
      if (!this.firstTimeSort) {
        this.$emit("sort", column.field, this.isAsc ? "asc" : "desc", event);
      }
      if (!this.backendSorting) {
        this.doSortSingleColumn(column);
      }
      this.currentSortColumn = column;
    },
    doSortSingleColumn(column) {
      this.newData = this.sortBy(this.newData, column.field, column.customSort, this.isAsc);
    },
    isRowSelected(row, selected) {
      if (!selected) {
        return false;
      }
      if (this.customRowKey) {
        return row[this.customRowKey] === selected[this.customRowKey];
      }
      return row === selected;
    },
    isRowChecked(row) {
      return indexOf(this.newCheckedRows, row, this.customIsChecked) >= 0;
    },
    removeCheckedRow(row) {
      const index2 = indexOf(this.newCheckedRows, row, this.customIsChecked);
      if (index2 >= 0) {
        this.newCheckedRows.splice(index2, 1);
      }
    },
    checkAll() {
      const isAllChecked = this.isAllChecked;
      this.visibleData.forEach((currentRow) => {
        if (this.isRowCheckable(currentRow)) {
          this.removeCheckedRow(currentRow);
        }
        if (!isAllChecked) {
          if (this.isRowCheckable(currentRow)) {
            this.newCheckedRows.push(currentRow);
          }
        }
      });
      this.$emit("check", this.newCheckedRows);
      this.$emit("check-all", this.newCheckedRows);
      this.$emit("update:checkedRows", this.newCheckedRows);
    },
    checkRow(row, index2, event) {
      if (!this.isRowCheckable(row))
        return;
      const lastIndex = this.lastCheckedRowIndex;
      this.lastCheckedRowIndex = index2;
      if (event.shiftKey && lastIndex !== null && index2 !== lastIndex) {
        this.shiftCheckRow(row, index2, lastIndex);
      } else if (!this.isRowChecked(row)) {
        this.newCheckedRows.push(row);
      } else {
        this.removeCheckedRow(row);
      }
      this.$emit("check", this.newCheckedRows, row);
      this.$emit("update:checkedRows", this.newCheckedRows);
    },
    shiftCheckRow(row, index2, lastCheckedRowIndex) {
      const subset = this.visibleData.slice(Math.min(index2, lastCheckedRowIndex), Math.max(index2, lastCheckedRowIndex) + 1);
      const shouldCheck = !this.isRowChecked(row);
      subset.forEach((item) => {
        this.removeCheckedRow(item);
        if (shouldCheck && this.isRowCheckable(item)) {
          this.newCheckedRows.push(item);
        }
      });
    },
    selectRow(row, index2) {
      this.$emit("click", row, index2);
      if (this.selected === row)
        return;
      if (!this.isRowSelectable(row))
        return;
      this.$emit("select", row, this.selected);
      this.$emit("update:selected", row);
    },
    toggleDetails(obj) {
      const found = this.isVisibleDetailRow(obj);
      if (found) {
        this.closeDetailRow(obj);
        this.$emit("details-close", obj);
      } else {
        this.openDetailRow(obj);
        this.$emit("details-open", obj);
      }
      this.$emit("update:openedDetailed", this.visibleDetailRows);
    },
    openDetailRow(obj) {
      const index2 = this.handleDetailKey(obj);
      this.visibleDetailRows.push(index2);
    },
    closeDetailRow(obj) {
      const index2 = this.handleDetailKey(obj);
      const i = this.visibleDetailRows.indexOf(index2);
      if (i >= 0) {
        this.visibleDetailRows.splice(i, 1);
      }
    },
    isVisibleDetailRow(obj) {
      const index2 = this.handleDetailKey(obj);
      return this.visibleDetailRows.indexOf(index2) >= 0;
    },
    isActiveDetailRow(row) {
      return this.detailed && !this.customDetailRow && this.isVisibleDetailRow(row);
    },
    isActiveCustomDetailRow(row) {
      return this.detailed && this.customDetailRow && this.isVisibleDetailRow(row);
    },
    isRowFiltered(row) {
      for (const key in this.filters) {
        if (!this.filters[key]) {
          delete this.filters[key];
          return true;
        }
        const input = this.filters[key];
        const column = this.newColumns.filter((c) => c.field === key)[0];
        if (column && column.customSearch && typeof column.customSearch === "function") {
          if (!column.customSearch(row, input))
            return false;
        } else {
          let value = this.getValueByPath(row, key);
          if (value == null)
            return false;
          if (Number.isInteger(value)) {
            if (value !== Number(input))
              return false;
          } else {
            const re = new RegExp(escapeRegExpChars(input), "i");
            if (!re.test(value))
              return false;
          }
        }
      }
      return true;
    },
    handleDetailKey(index2) {
      const key = this.detailKey;
      return !key.length || !index2 ? index2 : index2[key];
    },
    checkSort() {
      if (this.newColumns.length && this.firstTimeSort) {
        this.initSort();
        this.firstTimeSort = false;
      } else if (this.newColumns.length) {
        if (Object.keys(this.currentSortColumn).length > 0) {
          for (let i = 0; i < this.newColumns.length; i++) {
            if (this.newColumns[i].field === this.currentSortColumn.field) {
              this.currentSortColumn = this.newColumns[i];
              break;
            }
          }
        }
      }
    },
    hasCustomFooterSlot() {
      const footer = this.$slots.footer;
      if (footer.length > 1)
        return true;
      const tag = footer[0].tag;
      if (tag !== "th" && tag !== "td")
        return false;
      return true;
    },
    pressedArrow(pos) {
      if (!this.visibleData.length)
        return;
      let index2 = this.visibleData.indexOf(this.selected) + pos;
      index2 = index2 < 0 ? 0 : index2 > this.visibleData.length - 1 ? this.visibleData.length - 1 : index2;
      const row = this.visibleData[index2];
      if (!this.isRowSelectable(row)) {
        let newIndex = null;
        if (pos > 0) {
          for (let i = index2; i < this.visibleData.length && newIndex === null; i++) {
            if (this.isRowSelectable(this.visibleData[i]))
              newIndex = i;
          }
        } else {
          for (let i = index2; i >= 0 && newIndex === null; i--) {
            if (this.isRowSelectable(this.visibleData[i]))
              newIndex = i;
          }
        }
        if (newIndex >= 0) {
          this.selectRow(this.visibleData[newIndex]);
        }
      } else {
        this.selectRow(row);
      }
    },
    focus() {
      if (!this.focusable)
        return;
      this.$el.querySelector("table").focus();
    },
    initSort() {
      if (!this.defaultSort)
        return;
      let sortField = "";
      let sortDirection = this.defaultSortDirection;
      if (Array.isArray(this.defaultSort)) {
        sortField = this.defaultSort[0];
        if (this.defaultSort[1]) {
          sortDirection = this.defaultSort[1];
        }
      } else {
        sortField = this.defaultSort;
      }
      const sortColumn = this.newColumns.filter((column) => column.field === sortField)[0];
      if (sortColumn) {
        this.isAsc = sortDirection.toLowerCase() !== "desc";
        this.sort(sortColumn, true);
      }
    },
    handleDragStart(event, row, index2) {
      if (!this.draggable)
        return;
      this.$emit("dragstart", { event, row, index: index2 });
    },
    handleDragEnd(event, row, index2) {
      if (!this.draggable)
        return;
      this.$emit("dragend", { event, row, index: index2 });
    },
    handleDrop(event, row, index2) {
      if (!this.draggable)
        return;
      this.$emit("drop", { event, row, index: index2 });
    },
    handleDragOver(event, row, index2) {
      if (!this.draggable)
        return;
      this.$emit("dragover", { event, row, index: index2 });
    },
    handleDragLeave(event, row, index2) {
      if (!this.draggable)
        return;
      this.$emit("dragleave", { event, row, index: index2 });
    },
    emitEventForRow(eventName, event, row) {
      return this.$attrs[eventName] ? this.$emit(eventName, row, event) : null;
    },
    _addColumn(column) {
      if (typeof window !== "undefined") {
        this.$nextTick(() => {
          this.defaultSlots.push(column);
          requestAnimationFrame(() => {
            const div = this.$refs["slot"];
            if (div && div.children) {
              const position = [...div.children].map((c) => parseInt(c.getAttribute("data-id"), 10)).indexOf(column.newKey);
              if (position !== this.defaultSlots.length) {
                this.defaultSlots.splice(position, 0, column);
                this.defaultSlots = this.defaultSlots.slice(0, this.defaultSlots.length - 1);
              }
            }
          });
        });
      }
    },
    _removeColumn(column) {
      this.$nextTick(() => {
        this.defaultSlots = this.defaultSlots.filter((d) => d.newKey !== column.newKey);
      });
    },
    _nextSequence() {
      return this.sequence++;
    }
  }
});
const _hoisted_1$1$1 = {
  ref: "slot",
  style: {
    "display": "none"
  }
};
const _hoisted_2$1 = {
  key: 0
};
const _hoisted_3$1 = {
  key: 1
};
const _hoisted_4$1 = {
  key: 0,
  width: "40px"
};
const _hoisted_5$1 = {
  key: 1
};
const _hoisted_6$1 = {
  key: 0
};
const _hoisted_7 = {
  key: 1
};
const _hoisted_8 = {
  key: 2
};
const _hoisted_9 = {
  key: 0
};
const _hoisted_10 = {
  key: 2
};
function render$2$1(_ctx, _cache, $props, $setup, $data, $options) {
  const _component_o_table_mobile_sort = vue.resolveComponent("o-table-mobile-sort");
  const _component_o_table_pagination = vue.resolveComponent("o-table-pagination");
  const _component_o_checkbox = vue.resolveComponent("o-checkbox");
  const _component_o_slot_component = vue.resolveComponent("o-slot-component");
  const _component_o_icon = vue.resolveComponent("o-icon");
  const _component_o_input = vue.resolveComponent("o-input");
  const _component_o_loading = vue.resolveComponent("o-loading");
  return vue.openBlock(), vue.createBlock("div", {
    class: _ctx.tableWrapperClasses,
    style: _ctx.tableWrapperStyle
  }, [vue.createVNode("div", _hoisted_1$1$1, [vue.renderSlot(_ctx.$slots, "default")], 512), _ctx.isMobile && _ctx.hasSortablenewColumns ? vue.createVNode(_component_o_table_mobile_sort, {
    key: 0,
    "current-sort-column": _ctx.currentSortColumn,
    columns: _ctx.newColumns,
    placeholder: _ctx.mobileSortPlaceholder,
    "icon-pack": _ctx.iconPack,
    "sort-icon": _ctx.sortIcon,
    "sort-icon-size": _ctx.sortIconSize,
    "is-asc": _ctx.isAsc,
    onSort: _cache[1] || (_cache[1] = (column, event) => _ctx.sort(column, null, event)),
    "onRemove-priority": _cache[2] || (_cache[2] = (column) => _ctx.removeSortingPriority(column))
  }, null, 8, ["current-sort-column", "columns", "placeholder", "icon-pack", "sort-icon", "sort-icon-size", "is-asc"]) : vue.createCommentVNode("v-if", true), _ctx.paginated && (_ctx.paginationPosition === "top" || _ctx.paginationPosition === "both") ? vue.renderSlot(_ctx.$slots, "pagination", {
    key: 1
  }, () => [vue.createVNode(_component_o_table_pagination, vue.mergeProps(_ctx.$attrs, {
    "per-page": _ctx.perPage,
    paginated: _ctx.paginated,
    total: _ctx.newDataTotal,
    "current-page": _ctx.newCurrentPage,
    "onUpdate:currentPage": _cache[3] || (_cache[3] = ($event) => _ctx.newCurrentPage = $event),
    "root-class": _ctx.paginationWrapperClasses,
    "icon-pack": _ctx.iconPack,
    rounded: _ctx.paginationRounded,
    "onPage-change": _cache[4] || (_cache[4] = (event) => _ctx.$emit("page-change", event)),
    "aria-next-label": _ctx.ariaNextLabel,
    "aria-previous-label": _ctx.ariaPreviousLabel,
    "aria-page-label": _ctx.ariaPageLabel,
    "aria-current-label": _ctx.ariaCurrentLabel
  }), {
    default: vue.withCtx(() => [vue.renderSlot(_ctx.$slots, "top-left")]),
    _: 3
  }, 16, ["per-page", "paginated", "total", "current-page", "root-class", "icon-pack", "rounded", "aria-next-label", "aria-previous-label", "aria-page-label", "aria-current-label"])]) : vue.createCommentVNode("v-if", true), vue.createVNode("table", {
    class: _ctx.tableClasses,
    tabindex: !_ctx.focusable ? false : 0,
    onKeydown: [_cache[5] || (_cache[5] = vue.withKeys(vue.withModifiers(($event) => _ctx.pressedArrow(-1), ["self", "prevent"]), ["up"])), _cache[6] || (_cache[6] = vue.withKeys(vue.withModifiers(($event) => _ctx.pressedArrow(1), ["self", "prevent"]), ["down"]))]
  }, [_ctx.$slots.caption ? (vue.openBlock(), vue.createBlock("caption", _hoisted_2$1, [vue.renderSlot(_ctx.$slots, "caption")])) : vue.createCommentVNode("v-if", true), _ctx.newColumns.length && _ctx.showHeader ? (vue.openBlock(), vue.createBlock("thead", _hoisted_3$1, [vue.createVNode("tr", null, [_ctx.showDetailRowIcon ? (vue.openBlock(), vue.createBlock("th", _hoisted_4$1)) : vue.createCommentVNode("v-if", true), _ctx.checkable && _ctx.checkboxPosition === "left" ? (vue.openBlock(), vue.createBlock("th", {
    key: 1,
    class: _ctx.thCheckboxClasses
  }, [_ctx.headerCheckable ? vue.createVNode(_component_o_checkbox, {
    key: 0,
    autocomplete: "off",
    modelValue: _ctx.isAllChecked,
    disabled: _ctx.isAllUncheckable,
    onChange: _ctx.checkAll
  }, null, 8, ["modelValue", "disabled", "onChange"]) : vue.createCommentVNode("v-if", true)], 2)) : vue.createCommentVNode("v-if", true), (vue.openBlock(true), vue.createBlock(vue.Fragment, null, vue.renderList(_ctx.visibleColumns, (column, index2) => {
    return vue.openBlock(), vue.createBlock("th", vue.mergeProps({
      key: column.newKey + ":" + index2 + "header"
    }, column.thAttrs && column.thAttrs(column), {
      class: _ctx.thClasses(column),
      style: column.style,
      onClick: vue.withModifiers(($event) => _ctx.sort(column, null, $event), ["stop"])
    }), [column.hasHeaderSlot ? vue.createVNode(_component_o_slot_component, {
      key: 0,
      component: column,
      scoped: "",
      name: "header",
      tag: "span",
      props: {
        column,
        index: index2
      }
    }, null, 8, ["component", "props"]) : (vue.openBlock(), vue.createBlock("span", _hoisted_5$1, [vue.createTextVNode(vue.toDisplayString(column.label) + " ", 1), vue.withDirectives(vue.createVNode("span", {
      class: _ctx.thSortIconClasses(column)
    }, [vue.createVNode(_component_o_icon, {
      icon: _ctx.sortIcon,
      pack: _ctx.iconPack,
      both: "",
      size: _ctx.sortIconSize,
      rotation: !_ctx.isAsc ? 180 : 0
    }, null, 8, ["icon", "pack", "size", "rotation"])], 2), [[vue.vShow, column.sortable && _ctx.currentSortColumn === column]])]))], 16, ["onClick"]);
  }), 128)), _ctx.checkable && _ctx.checkboxPosition === "right" ? (vue.openBlock(), vue.createBlock("th", {
    key: 2,
    class: _ctx.thCheckboxClasses
  }, [_ctx.headerCheckable ? vue.createVNode(_component_o_checkbox, {
    key: 0,
    autocomplete: "off",
    modelValue: _ctx.isAllChecked,
    disabled: _ctx.isAllUncheckable,
    onChange: _ctx.checkAll
  }, null, 8, ["modelValue", "disabled", "onChange"]) : vue.createCommentVNode("v-if", true)], 2)) : vue.createCommentVNode("v-if", true)]), _ctx.hasSearchablenewColumns ? (vue.openBlock(), vue.createBlock("tr", _hoisted_6$1, [_ctx.showDetailRowIcon ? (vue.openBlock(), vue.createBlock("th", {
    key: 0,
    class: _ctx.thDetailedClasses
  }, null, 2)) : vue.createCommentVNode("v-if", true), _ctx.checkable && _ctx.checkboxPosition === "left" ? (vue.openBlock(), vue.createBlock("th", _hoisted_7)) : vue.createCommentVNode("v-if", true), (vue.openBlock(true), vue.createBlock(vue.Fragment, null, vue.renderList(_ctx.visibleColumns, (column, index2) => {
    return vue.openBlock(), vue.createBlock("th", vue.mergeProps({
      key: column.newKey + ":" + index2 + "searchable"
    }, column.thAttrs && column.thAttrs(column), {
      class: _ctx.thClasses(column),
      style: column.style
    }), [column.searchable ? (vue.openBlock(), vue.createBlock(vue.Fragment, {
      key: 0
    }, [column.hasSearchableSlot ? vue.createVNode(_component_o_slot_component, {
      key: 0,
      component: column,
      scoped: "",
      name: "searchable",
      tag: "span",
      props: {
        column,
        filters: _ctx.filters
      }
    }, null, 8, ["component", "props"]) : vue.createVNode(_component_o_input, {
      key: 1,
      ["on" + vue.capitalize(_ctx.filtersEvent)]: _ctx.onFiltersEvent,
      modelValue: _ctx.filters[column.field],
      "onUpdate:modelValue": ($event) => _ctx.filters[column.field] = $event,
      type: column.numeric ? "number" : "text"
    }, null, 16, ["modelValue", "onUpdate:modelValue", "type"])], 64)) : vue.createCommentVNode("v-if", true)], 16);
  }), 128)), _ctx.checkable && _ctx.checkboxPosition === "right" ? (vue.openBlock(), vue.createBlock("th", _hoisted_8)) : vue.createCommentVNode("v-if", true)])) : vue.createCommentVNode("v-if", true)])) : vue.createCommentVNode("v-if", true), vue.createVNode("tbody", null, [(vue.openBlock(true), vue.createBlock(vue.Fragment, null, vue.renderList(_ctx.visibleData, (row, index2) => {
    return vue.openBlock(), vue.createBlock(vue.Fragment, {
      key: this.customRowKey ? row[this.customRowKey] : index2
    }, [vue.createVNode("tr", {
      class: _ctx.rowClasses(row, index2),
      onClick: ($event) => _ctx.selectRow(row),
      onDblclick: ($event) => _ctx.$emit("dblclick", row),
      onMouseenter: ($event) => _ctx.emitEventForRow("mouseenter", $event, row),
      onMouseleave: ($event) => _ctx.emitEventForRow("mouseleave", $event, row),
      onContextmenu: ($event) => _ctx.$emit("contextmenu", row, $event),
      draggable: _ctx.draggable,
      onDragstart: ($event) => _ctx.handleDragStart($event, row, index2),
      onDragend: ($event) => _ctx.handleDragEnd($event, row, index2),
      onDrop: ($event) => _ctx.handleDrop($event, row, index2),
      onDragover: ($event) => _ctx.handleDragOver($event, row, index2),
      onDragleave: ($event) => _ctx.handleDragLeave($event, row, index2)
    }, [_ctx.showDetailRowIcon ? (vue.openBlock(), vue.createBlock("td", {
      key: 0,
      class: _ctx.tdDetailedChevronClasses
    }, [_ctx.hasDetailedVisible(row) ? vue.createVNode(_component_o_icon, {
      key: 0,
      icon: _ctx.detailIcon,
      pack: _ctx.iconPack,
      rotation: _ctx.isVisibleDetailRow(row) ? 90 : 0,
      role: "button",
      onClick: vue.withModifiers(($event) => _ctx.toggleDetails(row), ["stop"]),
      clickable: "",
      both: ""
    }, null, 8, ["icon", "pack", "rotation", "onClick"]) : vue.createCommentVNode("v-if", true)], 2)) : vue.createCommentVNode("v-if", true), _ctx.checkable && _ctx.checkboxPosition === "left" ? (vue.openBlock(), vue.createBlock("td", {
      key: 1,
      class: _ctx.tdCheckboxClasses
    }, [vue.createVNode(_component_o_checkbox, {
      autocomplete: "off",
      disabled: !_ctx.isRowCheckable(row),
      modelValue: _ctx.isRowChecked(row),
      "onUpdate:modelValue": ($event) => _ctx.checkRow(row, index2, $event)
    }, null, 8, ["disabled", "modelValue", "onUpdate:modelValue"])], 2)) : vue.createCommentVNode("v-if", true), (vue.openBlock(true), vue.createBlock(vue.Fragment, null, vue.renderList(_ctx.visibleColumns, (column, colindex) => {
      return vue.openBlock(), vue.createBlock(_component_o_slot_component, vue.mergeProps({
        key: column.newKey + index2 + ":" + colindex
      }, column.tdAttrs && column.tdAttrs(row, column), {
        component: column,
        scoped: "",
        name: "default",
        tag: "td",
        class: _ctx.tdClasses(row, column),
        "data-label": column.label,
        props: {
          row,
          column,
          index: index2,
          colindex,
          toggleDetails: _ctx.toggleDetails
        },
        onClick: ($event) => _ctx.$emit("cell-click", row, column, index2, colindex, $event)
      }), null, 16, ["component", "class", "data-label", "props", "onClick"]);
    }), 128)), _ctx.checkable && _ctx.checkboxPosition === "right" ? (vue.openBlock(), vue.createBlock("td", {
      key: 2,
      class: _ctx.tdCheckboxClasses
    }, [vue.createVNode(_component_o_checkbox, {
      autocomplete: "off",
      disabled: !_ctx.isRowCheckable(row),
      modelValue: _ctx.isRowChecked(row),
      "onUpdate:modelvalue": ($event) => _ctx.checkRow(row, index2, $event)
    }, null, 8, ["disabled", "modelValue", "onUpdate:modelvalue"])], 2)) : vue.createCommentVNode("v-if", true)], 42, ["onClick", "onDblclick", "onMouseenter", "onMouseleave", "onContextmenu", "draggable", "onDragstart", "onDragend", "onDrop", "onDragover", "onDragleave"]), vue.createVNode(vue.Transition, {
      name: _ctx.detailTransition
    }, {
      default: vue.withCtx(() => [_ctx.isActiveDetailRow(row) ? (vue.openBlock(), vue.createBlock("tr", {
        key: (this.customRowKey ? row[this.customRowKey] : index2) + "detail",
        class: _ctx.detailedClasses
      }, [vue.createVNode("td", {
        colspan: _ctx.columnCount
      }, [vue.renderSlot(_ctx.$slots, "detail", {
        row,
        index: index2
      })], 8, ["colspan"])], 2)) : vue.createCommentVNode("v-if", true)]),
      _: 2
    }, 1032, ["name"]), _ctx.isActiveCustomDetailRow(row) ? vue.renderSlot(_ctx.$slots, "detail", {
      key: 0,
      row,
      index: index2
    }) : vue.createCommentVNode("v-if", true)], 64);
  }), 128)), !_ctx.visibleData.length ? (vue.openBlock(), vue.createBlock("tr", _hoisted_9, [vue.createVNode("td", {
    colspan: _ctx.columnCount
  }, [vue.renderSlot(_ctx.$slots, "empty")], 8, ["colspan"])])) : vue.createCommentVNode("v-if", true)]), _ctx.$slots.footer ? (vue.openBlock(), vue.createBlock("tfoot", _hoisted_10, [vue.createVNode("tr", {
    class: _ctx.footerClasses
  }, [_ctx.hasCustomFooterSlot() ? vue.renderSlot(_ctx.$slots, "footer", {
    key: 0
  }) : (vue.openBlock(), vue.createBlock("th", {
    key: 1,
    colspan: _ctx.columnCount
  }, [vue.renderSlot(_ctx.$slots, "footer")], 8, ["colspan"]))], 2)])) : vue.createCommentVNode("v-if", true)], 42, ["tabindex"]), _ctx.loading ? vue.renderSlot(_ctx.$slots, "loading", {
    key: 2
  }, () => [vue.createVNode(_component_o_loading, {
    "full-page": false,
    active: _ctx.loading
  }, null, 8, ["active"])]) : vue.createCommentVNode("v-if", true), _ctx.checkable && _ctx.$slots["bottom-left"] || _ctx.paginated && (_ctx.paginationPosition === "bottom" || _ctx.paginationPosition === "both") ? vue.renderSlot(_ctx.$slots, "pagination", {
    key: 3
  }, () => [vue.createVNode(_component_o_table_pagination, vue.mergeProps(_ctx.$attrs, {
    "per-page": _ctx.perPage,
    paginated: _ctx.paginated,
    total: _ctx.newDataTotal,
    "current-page": _ctx.newCurrentPage,
    "onUpdate:currentPage": _cache[7] || (_cache[7] = ($event) => _ctx.newCurrentPage = $event),
    "root-class": _ctx.paginationWrapperClasses,
    "icon-pack": _ctx.iconPack,
    rounded: _ctx.paginationRounded,
    "onPage-change": _cache[8] || (_cache[8] = (event) => _ctx.$emit("page-change", event)),
    "aria-next-label": _ctx.ariaNextLabel,
    "aria-previous-label": _ctx.ariaPreviousLabel,
    "aria-page-label": _ctx.ariaPageLabel,
    "aria-current-label": _ctx.ariaCurrentLabel
  }), {
    default: vue.withCtx(() => [vue.renderSlot(_ctx.$slots, "bottom-left")]),
    _: 3
  }, 16, ["per-page", "paginated", "total", "current-page", "root-class", "icon-pack", "rounded", "aria-next-label", "aria-previous-label", "aria-page-label", "aria-current-label"])]) : vue.createCommentVNode("v-if", true)], 6);
}
script$3.render = render$2$1;
script$3.__file = "src/components/table/Table.vue";
var index$4 = {
  install(app) {
    registerComponent(app, script$3);
    registerComponent(app, script$1$2);
  }
};
var index$m = index$4;
var script$1 = vue.defineComponent({
  name: "OTabs",
  mixins: [BaseComponentMixin, TabbedMixin("tab")],
  configField: "tabs",
  props: {
    type: {
      type: String,
      default: "default"
    },
    expanded: Boolean,
    animated: {
      type: Boolean,
      default: () => {
        return getValueByPath(getOptions$1(), "tabs.animated", true);
      }
    },
    multiline: Boolean,
    rootClass: [String, Function, Array],
    positionClass: [String, Function, Array],
    expandedClass: [String, Function, Array],
    verticalClass: [String, Function, Array],
    multilineClass: [String, Function, Array],
    navTabsClass: [String, Function, Array],
    navSizeClass: [String, Function, Array],
    navPositionClass: [String, Function, Array],
    contentClass: [String, Function, Array],
    transitioningClass: [String, Function, Array],
    tabItemWrapperClass: [String, Function, Array]
  },
  computed: {
    rootClasses() {
      return [this.computedClass("rootClass", "o-tabs"), {
        [this.computedClass("positionClass", "o-tabs--", this.position)]: this.position && this.vertical
      }, {
        [this.computedClass("expandedClass", "o-tabs--fullwidth")]: this.expanded
      }, {
        [this.computedClass("verticalClass", "o-tabs--vertical")]: this.vertical
      }, {
        [this.computedClass("multilineClass", "o-tabs--multiline")]: this.multiline
      }];
    },
    itemWrapperClasses() {
      return [this.computedClass("tabItemWrapperClass", "o-tabs__nav-item-wrapper")];
    },
    navClasses() {
      return [this.computedClass("navTabsClass", "o-tabs__nav"), {
        [this.computedClass("navSizeClass", "o-tabs__nav--", this.size)]: this.size
      }, {
        [this.computedClass("navPositionClass", "o-tabs__nav--", this.position)]: this.position && !this.vertical
      }];
    },
    contentClasses() {
      return [this.computedClass("contentClass", "o-tabs__content"), {
        [this.computedClass("transitioningClass", "o-tabs__content--transitioning")]: this.isTransitioning
      }];
    }
  }
});
function render$1(_ctx, _cache, $props, $setup, $data, $options) {
  const _component_o_slot_component = vue.resolveComponent("o-slot-component");
  const _component_o_icon = vue.resolveComponent("o-icon");
  return vue.openBlock(), vue.createBlock("div", {
    class: _ctx.rootClasses
  }, [vue.createVNode("nav", {
    class: _ctx.navClasses
  }, [(vue.openBlock(true), vue.createBlock(vue.Fragment, null, vue.renderList(_ctx.items, (childItem) => {
    return vue.withDirectives((vue.openBlock(), vue.createBlock("div", {
      key: childItem.newValue,
      class: _ctx.itemWrapperClasses
    }, [childItem.$slots.header ? vue.createVNode(_component_o_slot_component, {
      key: 0,
      component: childItem,
      tag: childItem.tag,
      name: "header",
      onClick: ($event) => _ctx.childClick(childItem),
      class: childItem.headerClasses
    }, null, 8, ["component", "tag", "onClick", "class"]) : (vue.openBlock(), vue.createBlock(vue.resolveDynamicComponent(childItem.tag), {
      key: 1,
      onClick: ($event) => _ctx.childClick(childItem),
      class: childItem.headerClasses
    }, {
      default: vue.withCtx(() => [childItem.icon ? vue.createVNode(_component_o_icon, {
        key: 0,
        rootClass: childItem.headerIconClasses,
        icon: childItem.icon,
        pack: childItem.iconPack,
        size: _ctx.size
      }, null, 8, ["rootClass", "icon", "pack", "size"]) : vue.createCommentVNode("v-if", true), vue.createVNode("span", {
        class: childItem.headerTextClasses
      }, vue.toDisplayString(childItem.label), 3)]),
      _: 2
    }, 1032, ["onClick", "class"]))], 2)), [[vue.vShow, childItem.visible]]);
  }), 128))], 2), vue.createVNode("section", {
    class: _ctx.contentClasses
  }, [vue.renderSlot(_ctx.$slots, "default")], 2)], 2);
}
script$1.render = render$1;
script$1.__file = "src/components/tabs/Tabs.vue";
var script$1$1 = vue.defineComponent({
  name: "OTabItem",
  mixins: [BaseComponentMixin, TabbedChildMixin("tab")],
  configField: "tabs",
  props: {
    disabled: Boolean,
    itemClass: [String, Function, Array],
    itemHeaderClass: [String, Function, Array],
    itemHeaderActiveClass: [String, Function, Array],
    itemHeaderDisabledClass: [String, Function, Array],
    itemHeaderTypeClass: [String, Function, Array],
    itemHeaderIconClass: [String, Function, Array],
    itemHeaderTextClass: [String, Function, Array],
    tag: {
      type: String,
      default: () => {
        return getValueByPath(getOptions$1(), "dropdown.itemTag", "div");
      }
    }
  },
  computed: {
    elementClasses() {
      return [
        this.computedClass("itemClass", "o-tab-item__content")
      ];
    },
    headerClasses() {
      return [
        this.computedClass("itemHeaderClass", "o-tabs__nav-item"),
        { [this.computedClass("itemHeaderActiveClass", "o-tabs__nav-item-{*}--active", this.$parent.type)]: this.isActive },
        { [this.computedClass("itemHeaderDisabledClass", "o-tabs__nav-item-{*}--disabled", this.$parent.type)]: this.disabled },
        { [this.computedClass("itemHeaderTypeClass", "o-tabs__nav-item-", this.$parent.type)]: this.$parent.type }
      ];
    },
    headerIconClasses() {
      return [
        this.computedClass("itemHeaderIconClass", "o-tabs__nav-item-icon")
      ];
    },
    headerTextClasses() {
      return [
        this.computedClass("itemHeaderTextClass", "o-tabs__nav-item-text")
      ];
    }
  }
});
script$1$1.__file = "src/components/tabs/TabItem.vue";
var index$3 = {
  install(app) {
    registerComponent(app, script$1);
    registerComponent(app, script$1$1);
  }
};
var index$n = index$3;
var index$2 = {
  install(app) {
    registerComponent(app, script$h);
  }
};
var index$o = index$2;
var index$1 = {
  install(app) {
    registerComponent(app, script$7);
  }
};
var index$p = index$1;
var script = vue.defineComponent({
  name: "OUpload",
  mixins: [BaseComponentMixin, FormElementMixin],
  configField: "upload",
  inheritAttrs: false,
  provide() {
    return {
      $elementRef: "input"
    };
  },
  emits: ["update:modelValue"],
  props: {
    modelValue: [Object, File, Array],
    multiple: Boolean,
    disabled: Boolean,
    accept: String,
    dragDrop: Boolean,
    variant: {
      type: String
    },
    native: {
      type: Boolean,
      default: false
    },
    expanded: {
      type: Boolean,
      default: false
    },
    rootClass: [String, Function, Array],
    draggableClass: [String, Function, Array],
    variantClass: [String, Function, Array],
    expandedClass: [String, Function, Array],
    disabledClass: [String, Function, Array],
    hoveredClass: [String, Function, Array]
  },
  data() {
    return {
      newValue: this.modelValue,
      dragDropFocus: false
    };
  },
  computed: {
    rootClasses() {
      return [this.computedClass("rootClass", "o-upl"), {
        [this.computedClass("expandedClass", "o-upl--expanded")]: this.expanded
      }, {
        [this.computedClass("disabledClass", "o-upl--disabled")]: this.disabled
      }];
    },
    draggableClasses() {
      return [this.computedClass("draggableClass", "o-upl__draggable"), {
        [this.computedClass("hoveredClass", "o-upl__draggable--hovered")]: !this.variant && this.dragDropFocus
      }, {
        [this.computedClass("variantClass", "o-upl__draggable--hovered-", this.variant)]: this.variant && this.dragDropFocus
      }];
    }
  },
  watch: {
    modelValue(value) {
      this.newValue = value;
      if (!value || Array.isArray(value) && value.length === 0) {
        this.$refs.input.value = null;
      }
      !this.isValid && !this.dragDrop && this.checkHtml5Validity();
    }
  },
  methods: {
    onFileChange(event) {
      if (this.disabled)
        return;
      if (this.dragDrop)
        this.updateDragDropFocus(false);
      const value = event.target.files || event.dataTransfer.files;
      if (value.length === 0) {
        if (!this.newValue)
          return;
        if (this.native)
          this.newValue = null;
      } else if (!this.multiple) {
        if (this.dragDrop && value.length !== 1)
          return;
        else {
          const file = value[0];
          if (this.checkType(file))
            this.newValue = file;
          else if (this.newValue)
            this.newValue = null;
          else
            return;
        }
      } else {
        let newValues = false;
        if (this.native || !this.newValue) {
          this.newValue = [];
          newValues = true;
        }
        for (let i = 0; i < value.length; i++) {
          const file = value[i];
          if (this.checkType(file)) {
            this.newValue.push(file);
            newValues = true;
          }
        }
        if (!newValues)
          return;
      }
      this.$emit("update:modelValue", this.newValue);
      !this.dragDrop && this.checkHtml5Validity();
    },
    updateDragDropFocus(focus) {
      if (!this.disabled) {
        this.dragDropFocus = focus;
      }
    },
    checkType(file) {
      if (!this.accept)
        return true;
      const types = this.accept.split(",");
      if (types.length === 0)
        return true;
      for (let i = 0; i < types.length; i++) {
        const type = types[i].trim();
        if (type) {
          if (type.substring(0, 1) === ".") {
            const extIndex = file.name.lastIndexOf(".");
            const extension = extIndex >= 0 ? file.name.substring(extIndex) : "";
            if (extension.toLowerCase() === type.toLowerCase()) {
              return true;
            }
          } else {
            if (file.type.match(type))
              return true;
          }
        }
      }
      return false;
    }
  }
});
function render2(_ctx, _cache, $props, $setup, $data, $options) {
  return vue.openBlock(), vue.createBlock("label", {
    class: _ctx.rootClasses
  }, [!_ctx.dragDrop ? vue.renderSlot(_ctx.$slots, "default", {
    key: 0
  }) : (vue.openBlock(), vue.createBlock("div", {
    key: 1,
    class: _ctx.draggableClasses,
    onMouseenter: _cache[1] || (_cache[1] = ($event) => _ctx.updateDragDropFocus(true)),
    onMouseleave: _cache[2] || (_cache[2] = ($event) => _ctx.updateDragDropFocus(false)),
    onDragover: _cache[3] || (_cache[3] = vue.withModifiers(($event) => _ctx.updateDragDropFocus(true), ["prevent"])),
    onDragleave: _cache[4] || (_cache[4] = vue.withModifiers(($event) => _ctx.updateDragDropFocus(false), ["prevent"])),
    onDragenter: _cache[5] || (_cache[5] = vue.withModifiers(($event) => _ctx.updateDragDropFocus(true), ["prevent"])),
    onDrop: _cache[6] || (_cache[6] = vue.withModifiers((...args) => _ctx.onFileChange(...args), ["prevent"]))
  }, [vue.renderSlot(_ctx.$slots, "default")], 34)), vue.createVNode("input", vue.mergeProps({
    ref: "input",
    type: "file"
  }, _ctx.$attrs, {
    multiple: _ctx.multiple,
    accept: _ctx.accept,
    disabled: _ctx.disabled,
    onChange: _cache[7] || (_cache[7] = (...args) => _ctx.onFileChange(...args))
  }), null, 16, ["multiple", "accept", "disabled"])], 2);
}
script.render = render2;
script.__file = "src/components/upload/Upload.vue";
var index = {
  install(app) {
    registerComponent(app, script);
  }
};
var index$q = index;
var components = /* @__PURE__ */ Object.freeze({
  __proto__: null,
  Autocomplete: index$E,
  Button: index$1$1,
  Checkbox: index$2$1,
  Collapse: index$3$1,
  Datepicker: index$4$1,
  Datetimepicker: index$5$1,
  Dropdown: index$6$1,
  Field: index$7$1,
  Icon: index$8$1,
  Input: index$9$1,
  Inputitems: index$a$1,
  Loading: index$b$1,
  Modal: index$c$1,
  Notification: index$d$1,
  Pagination: index$e,
  Radio: index$f,
  Select: index$g,
  Skeleton: index$h,
  Sidebar: index$i,
  Slider: index$j,
  Steps: index$k,
  Switch: index$l,
  Table: index$m,
  Tabs: index$n,
  Timepicker: index$o,
  Tooltip: index$p,
  Upload: index$q
});
const Oruga = {
  install(app, options = {}) {
    setVueInstance(app);
    const defaultConfig = getOptions$1();
    setOptions(merge(defaultConfig, options, true));
    for (const componentKey in components) {
      registerPlugin(app, components[componentKey]);
    }
    registerComponentProgrammatic(app, "config", Programmatic);
    app.provide("oruga", app.config.globalProperties.$oruga);
  }
};
var OIcon = Oruga;
const _sfc_main$p = {
  props: {
    logs: {
      type: Array,
      default: function() {
        return [];
      }
    }
  },
  components: {
    OIcon
  },
  methods: {
    outputClass(variant) {
      if (variant == "success") {
        return `bg-green-600 text-green-300`;
      } else if (variant == "danger") {
        return `bg-red-600 text-red-300`;
      } else if (variant == "warning") {
        return `bg-yellow-600 text-yellow-300`;
      } else {
        return `bg-gray-600 text-gray-300`;
      }
    },
    formatTime(timestamp) {
      return timestamp.toLocaleString(DateTime.TIME_WITH_SECONDS);
    }
  }
};
const _hoisted_1 = {
  key: 0,
  class: "absolute mt-1 ml-1"
};
const _hoisted_2 = {
  key: 1,
  class: "absolute mt-1 ml-1"
};
const _hoisted_3 = {
  key: 2,
  class: "absolute mt-1 ml-1"
};
const _hoisted_4 = {
  key: 3,
  class: "absolute mt-1 ml-1"
};
const _hoisted_5 = { class: "mr-3 ml-8 float-left" };
const _hoisted_6 = { class: "text-gray-100" };
function _sfc_render(_ctx, _cache, $props, $setup, $data, $options) {
  const _component_OIcon = vue.resolveComponent("OIcon");
  return vue.openBlock(true), vue.createElementBlock(vue.Fragment, null, vue.renderList($props.logs, (output) => {
    return vue.openBlock(), vue.createElementBlock("div", {
      key: output.timestamp,
      class: vue.normalizeClass([$options.outputClass(output.variant), "p-2 text-lg font-semibold relative"])
    }, [
      output.variant == "success" ? (vue.openBlock(), vue.createElementBlock("div", _hoisted_1, [
        vue.createVNode(_component_OIcon, { icon: "check-circle" })
      ])) : vue.createCommentVNode("", true),
      output.variant == "danger" ? (vue.openBlock(), vue.createElementBlock("div", _hoisted_2, [
        vue.createVNode(_component_OIcon, { icon: "exclamation-circle" })
      ])) : vue.createCommentVNode("", true),
      output.variant == "warning" ? (vue.openBlock(), vue.createElementBlock("div", _hoisted_3, [
        vue.createVNode(_component_OIcon, { icon: "exclamation-triangle" })
      ])) : vue.createCommentVNode("", true),
      output.variant == "" ? (vue.openBlock(), vue.createElementBlock("div", _hoisted_4, [
        vue.createVNode(_component_OIcon, { icon: "info-circle" })
      ])) : vue.createCommentVNode("", true),
      vue.createElementVNode("div", _hoisted_5, vue.toDisplayString($options.formatTime(output.timestamp)), 1),
      vue.createElementVNode("div", _hoisted_6, vue.toDisplayString(output.msg), 1)
    ], 2);
  }), 128);
}
var LogPane = /* @__PURE__ */ _export_sfc(_sfc_main$p, [["render", _sfc_render]]);
var lib = {
  CSLEditor,
  LogPane
};
var prismeditor_min = "";
var repl_html_vue_vue_type_style_index_0_scoped_true_lang = "";
console.log(lib);
const _sfc_main$o = {
  data() {
    return {
      code: `this
            is
            code`,
      logs: [
        {
          msg: "Hi",
          variant: "danger",
          timestamp: luxon.DateTime.now()
        }
      ]
    };
  },
  components: {
    CSLEditor: lib.CSLEditor,
    LogPane: lib.LogPane
  }
};
function _sfc_ssrRender$a(_ctx, _push, _parent, _attrs, $props, $setup, $data, $options) {
  const _component_CSLEditor = vue.resolveComponent("CSLEditor");
  const _component_LogPane = vue.resolveComponent("LogPane");
  _push(`<!--[--><h1 id="repl" tabindex="-1" data-v-429ef458><a class="header-anchor" href="#repl" aria-hidden="true" data-v-429ef458>#</a> REPL</h1>`);
  _push(serverRenderer.ssrRenderComponent(_component_CSLEditor, {
    modelValue: $data.code,
    "onUpdate:modelValue": ($event) => $data.code = $event
  }, null, _parent));
  _push(serverRenderer.ssrRenderComponent(_component_LogPane, { logs: $data.logs }, null, _parent));
  _push(`<!--]-->`);
}
const _sfc_setup$o = _sfc_main$o.setup;
_sfc_main$o.setup = (props, ctx) => {
  const ssrContext = vue.useSSRContext();
  (ssrContext.modules || (ssrContext.modules = new Set())).add("../pages/csl/repl.html.vue");
  return _sfc_setup$o ? _sfc_setup$o(props, ctx) : void 0;
};
var repl_html$1 = /* @__PURE__ */ _export_sfc$1(_sfc_main$o, [["ssrRender", _sfc_ssrRender$a], ["__scopeId", "data-v-429ef458"]]);
var repl_html$2 = /* @__PURE__ */ Object.freeze({
  __proto__: null,
  [Symbol.toStringTag]: "Module",
  "default": repl_html$1
});
const _sfc_main$n = {};
function _sfc_ssrRender$9(_ctx, _push, _parent, _attrs) {
  const _component_OutboundLink = vue.resolveComponent("OutboundLink");
  _push(`<!--[--><h1 id="configuration" tabindex="-1"><a class="header-anchor" href="#configuration" aria-hidden="true">#</a> Configuration</h1><p>Conductorr uses the following environment variables for configuration</p><table><thead><tr><th>Environment Variable</th><th>Description</th></tr></thead><tbody><tr><td>TMDB_API_KEY</td><td>API key for searching media and fetching media metadata from <a href="https://themoviedb.org" target="_blank" rel="noopener noreferrer">The Movie Database`);
  _push(serverRenderer.ssrRenderComponent(_component_OutboundLink, null, null, _parent));
  _push(`</a></td></tr><tr><td>CONDUCTORR_DEBUG</td><td>Run Conductorr in debug mode if this environment variable is set to anything</td></tr><tr><td>JWT_EXP_DAYS</td><td>Number of days a user should stay logged in for without logging in</td></tr><tr><td>RESET_ADMIN_USER</td><td>If you forgot your password, set this environment variable to true and Conductorr will prompt you to set up a new user next time it starts</td></tr><tr><td>PG_USER</td><td>PostgreSQL user. If set, Conductorr will not use SQLite and will instead use PostgreSQL (optional - required for postgres support)</td></tr><tr><td>PG_PASS</td><td>PostgreSQL user password (optional - required for postgres support)</td></tr><tr><td>PG_DATABASE</td><td>PostgreSQL database name (optional - required for postgres support)</td></tr><tr><td>PG_PORT</td><td>PostgreSQL server port (optional - default is 5432)</td></tr><tr><td>PG_SSL</td><td>If set, Conductorr will force SSL for PostgreSQL connections (optional)</td></tr><tr><td>DB_PATH</td><td>When using SQLite, this is the path for the database file. Do not append any URL parameters to the end</td></tr></tbody></table><!--]-->`);
}
const _sfc_setup$n = _sfc_main$n.setup;
_sfc_main$n.setup = (props, ctx) => {
  const ssrContext = vue.useSSRContext();
  (ssrContext.modules || (ssrContext.modules = new Set())).add("../pages/guide/configuration.html.vue");
  return _sfc_setup$n ? _sfc_setup$n(props, ctx) : void 0;
};
var configuration_html$1 = /* @__PURE__ */ _export_sfc$1(_sfc_main$n, [["ssrRender", _sfc_ssrRender$9]]);
var configuration_html$2 = /* @__PURE__ */ Object.freeze({
  __proto__: null,
  [Symbol.toStringTag]: "Module",
  "default": configuration_html$1
});
const _sfc_main$m = {};
function _sfc_ssrRender$8(_ctx, _push, _parent, _attrs) {
  const _component_RouterLink = vue.resolveComponent("RouterLink");
  const _component_OutboundLink = vue.resolveComponent("OutboundLink");
  _push(`<!--[--><h1 id="download-and-installation" tabindex="-1"><a class="header-anchor" href="#download-and-installation" aria-hidden="true">#</a> Download and Installation</h1><p>Follow the instructions below depending on your system and preferred installation method. For configuration instructions post-installation, check `);
  _push(serverRenderer.ssrRenderComponent(_component_RouterLink, { to: "/guide/configuration.html" }, {
    default: vue.withCtx((_, _push2, _parent2, _scopeId) => {
      if (_push2) {
        _push2(`here`);
      } else {
        return [
          vue.createTextVNode("here")
        ];
      }
    }),
    _: 1
  }, _parent));
  _push(`.</p><h2 id="manual-installation" tabindex="-1"><a class="header-anchor" href="#manual-installation" aria-hidden="true">#</a> Manual Installation</h2><p>Download the latest release for your operating system <a href="https://github.com/lsnow99/conductorr/releases" target="_blank" rel="noopener noreferrer">here`);
  _push(serverRenderer.ssrRenderComponent(_component_OutboundLink, null, null, _parent));
  _push(`</a>.</p><p>Put the binary file in a suitable location. Simply double click the binary or run <code>./conductorr</code> in your terminal to start Conductorr. On some systems you may need to first run <code>chmod +x conductorr</code> in your terminal before launching.</p><p>To launch automatically at startup, please refer to your operating system&#39;s instructions.</p><p>Navigate to <a href="http://localhost:6416/" target="_blank" rel="noopener noreferrer">http://localhost:6416/`);
  _push(serverRenderer.ssrRenderComponent(_component_OutboundLink, null, null, _parent));
  _push(`</a> and you should see the setup screen prompting you to create the admin user.</p><h2 id="docker-compose" tabindex="-1"><a class="header-anchor" href="#docker-compose" aria-hidden="true">#</a> Docker Compose</h2><p>Create a <code>docker-compose.yml</code> file with the following contents:</p><div class="language-yaml ext-yml line-numbers-mode"><pre class="language-yaml"><code><span class="token key atrule">version</span><span class="token punctuation">:</span> <span class="token string">&quot;3.9&quot;</span>
<span class="token key atrule">services</span><span class="token punctuation">:</span>
  <span class="token key atrule">conductorr</span><span class="token punctuation">:</span>
    <span class="token key atrule">image</span><span class="token punctuation">:</span> <span class="token string">&quot;conductorr&quot;</span>
    <span class="token key atrule">ports</span><span class="token punctuation">:</span>
      <span class="token punctuation">-</span> <span class="token string">&quot;6416:6416&quot;</span>
    <span class="token key atrule">volumes</span><span class="token punctuation">:</span>
      <span class="token comment"># Replace the . with your preferred location on your system for the database file</span>
      <span class="token punctuation">-</span> .<span class="token punctuation">:</span>/app/conductorr.db  
    <span class="token key atrule">environment</span><span class="token punctuation">:</span>
      <span class="token key atrule">TMDB_API_KEY</span><span class="token punctuation">:</span> yourapikeyhere
      <span class="token comment"># Add any other environment variables for configuration here</span>
</code></pre><div class="line-numbers"><span class="line-number">1</span><br><span class="line-number">2</span><br><span class="line-number">3</span><br><span class="line-number">4</span><br><span class="line-number">5</span><br><span class="line-number">6</span><br><span class="line-number">7</span><br><span class="line-number">8</span><br><span class="line-number">9</span><br><span class="line-number">10</span><br><span class="line-number">11</span><br><span class="line-number">12</span><br></div></div><!--]-->`);
}
const _sfc_setup$m = _sfc_main$m.setup;
_sfc_main$m.setup = (props, ctx) => {
  const ssrContext = vue.useSSRContext();
  (ssrContext.modules || (ssrContext.modules = new Set())).add("../pages/guide/download.html.vue");
  return _sfc_setup$m ? _sfc_setup$m(props, ctx) : void 0;
};
var download_html$1 = /* @__PURE__ */ _export_sfc$1(_sfc_main$m, [["ssrRender", _sfc_ssrRender$8]]);
var download_html$2 = /* @__PURE__ */ Object.freeze({
  __proto__: null,
  [Symbol.toStringTag]: "Module",
  "default": download_html$1
});
const _sfc_main$l = {};
function _sfc_ssrRender$7(_ctx, _push, _parent, _attrs) {
  _push(`<h1${serverRenderer.ssrRenderAttrs(vue.mergeProps({
    id: "calendar",
    tabindex: "-1"
  }, _attrs))}><a class="header-anchor" href="#calendar" aria-hidden="true">#</a> Calendar</h1>`);
}
const _sfc_setup$l = _sfc_main$l.setup;
_sfc_main$l.setup = (props, ctx) => {
  const ssrContext = vue.useSSRContext();
  (ssrContext.modules || (ssrContext.modules = new Set())).add("../pages/features/calendar.html.vue");
  return _sfc_setup$l ? _sfc_setup$l(props, ctx) : void 0;
};
var calendar_html$1 = /* @__PURE__ */ _export_sfc$1(_sfc_main$l, [["ssrRender", _sfc_ssrRender$7]]);
var calendar_html$2 = /* @__PURE__ */ Object.freeze({
  __proto__: null,
  [Symbol.toStringTag]: "Module",
  "default": calendar_html$1
});
const _sfc_main$k = {};
function _sfc_ssrRender$6(_ctx, _push, _parent, _attrs) {
  _push(`<h1${serverRenderer.ssrRenderAttrs(vue.mergeProps({
    id: "downloaders",
    tabindex: "-1"
  }, _attrs))}><a class="header-anchor" href="#downloaders" aria-hidden="true">#</a> Downloaders</h1>`);
}
const _sfc_setup$k = _sfc_main$k.setup;
_sfc_main$k.setup = (props, ctx) => {
  const ssrContext = vue.useSSRContext();
  (ssrContext.modules || (ssrContext.modules = new Set())).add("../pages/features/downloaders.html.vue");
  return _sfc_setup$k ? _sfc_setup$k(props, ctx) : void 0;
};
var downloaders_html$1 = /* @__PURE__ */ _export_sfc$1(_sfc_main$k, [["ssrRender", _sfc_ssrRender$6]]);
var downloaders_html$2 = /* @__PURE__ */ Object.freeze({
  __proto__: null,
  [Symbol.toStringTag]: "Module",
  "default": downloaders_html$1
});
const _sfc_main$j = {};
function _sfc_ssrRender$5(_ctx, _push, _parent, _attrs) {
  _push(`<h1${serverRenderer.ssrRenderAttrs(vue.mergeProps({
    id: "indexers",
    tabindex: "-1"
  }, _attrs))}><a class="header-anchor" href="#indexers" aria-hidden="true">#</a> Indexers</h1>`);
}
const _sfc_setup$j = _sfc_main$j.setup;
_sfc_main$j.setup = (props, ctx) => {
  const ssrContext = vue.useSSRContext();
  (ssrContext.modules || (ssrContext.modules = new Set())).add("../pages/features/indexers.html.vue");
  return _sfc_setup$j ? _sfc_setup$j(props, ctx) : void 0;
};
var indexers_html$1 = /* @__PURE__ */ _export_sfc$1(_sfc_main$j, [["ssrRender", _sfc_ssrRender$5]]);
var indexers_html$2 = /* @__PURE__ */ Object.freeze({
  __proto__: null,
  [Symbol.toStringTag]: "Module",
  "default": indexers_html$1
});
const _sfc_main$i = {};
function _sfc_ssrRender$4(_ctx, _push, _parent, _attrs) {
  _push(`<h1${serverRenderer.ssrRenderAttrs(vue.mergeProps({
    id: "notifications",
    tabindex: "-1"
  }, _attrs))}><a class="header-anchor" href="#notifications" aria-hidden="true">#</a> Notifications</h1>`);
}
const _sfc_setup$i = _sfc_main$i.setup;
_sfc_main$i.setup = (props, ctx) => {
  const ssrContext = vue.useSSRContext();
  (ssrContext.modules || (ssrContext.modules = new Set())).add("../pages/features/notifications.html.vue");
  return _sfc_setup$i ? _sfc_setup$i(props, ctx) : void 0;
};
var notifications_html$1 = /* @__PURE__ */ _export_sfc$1(_sfc_main$i, [["ssrRender", _sfc_ssrRender$4]]);
var notifications_html$2 = /* @__PURE__ */ Object.freeze({
  __proto__: null,
  [Symbol.toStringTag]: "Module",
  "default": notifications_html$1
});
const _sfc_main$h = {};
function _sfc_ssrRender$3(_ctx, _push, _parent, _attrs) {
  _push(`<h1${serverRenderer.ssrRenderAttrs(vue.mergeProps({
    id: "post-processing",
    tabindex: "-1"
  }, _attrs))}><a class="header-anchor" href="#post-processing" aria-hidden="true">#</a> Post-Processing</h1>`);
}
const _sfc_setup$h = _sfc_main$h.setup;
_sfc_main$h.setup = (props, ctx) => {
  const ssrContext = vue.useSSRContext();
  (ssrContext.modules || (ssrContext.modules = new Set())).add("../pages/features/post-processing.html.vue");
  return _sfc_setup$h ? _sfc_setup$h(props, ctx) : void 0;
};
var postProcessing_html$1 = /* @__PURE__ */ _export_sfc$1(_sfc_main$h, [["ssrRender", _sfc_ssrRender$3]]);
var postProcessing_html$2 = /* @__PURE__ */ Object.freeze({
  __proto__: null,
  [Symbol.toStringTag]: "Module",
  "default": postProcessing_html$1
});
const _sfc_main$g = {};
function _sfc_ssrRender$2(_ctx, _push, _parent, _attrs) {
  _push(`<h1${serverRenderer.ssrRenderAttrs(vue.mergeProps({
    id: "profiles",
    tabindex: "-1"
  }, _attrs))}><a class="header-anchor" href="#profiles" aria-hidden="true">#</a> Profiles</h1>`);
}
const _sfc_setup$g = _sfc_main$g.setup;
_sfc_main$g.setup = (props, ctx) => {
  const ssrContext = vue.useSSRContext();
  (ssrContext.modules || (ssrContext.modules = new Set())).add("../pages/features/profiles.html.vue");
  return _sfc_setup$g ? _sfc_setup$g(props, ctx) : void 0;
};
var profiles_html$1 = /* @__PURE__ */ _export_sfc$1(_sfc_main$g, [["ssrRender", _sfc_ssrRender$2]]);
var profiles_html$2 = /* @__PURE__ */ Object.freeze({
  __proto__: null,
  [Symbol.toStringTag]: "Module",
  "default": profiles_html$1
});
const _sfc_main$f = {};
function _sfc_ssrRender$1(_ctx, _push, _parent, _attrs) {
  _push(`<h1${serverRenderer.ssrRenderAttrs(vue.mergeProps({
    id: "syncing",
    tabindex: "-1"
  }, _attrs))}><a class="header-anchor" href="#syncing" aria-hidden="true">#</a> Syncing</h1>`);
}
const _sfc_setup$f = _sfc_main$f.setup;
_sfc_main$f.setup = (props, ctx) => {
  const ssrContext = vue.useSSRContext();
  (ssrContext.modules || (ssrContext.modules = new Set())).add("../pages/features/syncing.html.vue");
  return _sfc_setup$f ? _sfc_setup$f(props, ctx) : void 0;
};
var syncing_html$1 = /* @__PURE__ */ _export_sfc$1(_sfc_main$f, [["ssrRender", _sfc_ssrRender$1]]);
var syncing_html$2 = /* @__PURE__ */ Object.freeze({
  __proto__: null,
  [Symbol.toStringTag]: "Module",
  "default": syncing_html$1
});
const _sfc_main$e = {};
function _sfc_ssrRender(_ctx, _push, _parent, _attrs) {
}
const _sfc_setup$e = _sfc_main$e.setup;
_sfc_main$e.setup = (props, ctx) => {
  const ssrContext = vue.useSSRContext();
  (ssrContext.modules || (ssrContext.modules = new Set())).add("../pages/404.html.vue");
  return _sfc_setup$e ? _sfc_setup$e(props, ctx) : void 0;
};
var _404_html$1 = /* @__PURE__ */ _export_sfc$1(_sfc_main$e, [["ssrRender", _sfc_ssrRender]]);
var _404_html$2 = /* @__PURE__ */ Object.freeze({
  __proto__: null,
  [Symbol.toStringTag]: "Module",
  "default": _404_html$1
});
const data$d = {
  "key": "v-8daa1a0e",
  "path": "/",
  "title": "",
  "lang": "en-US",
  "frontmatter": {
    "home": true,
    "heroImage": "/logo-rect.svg",
    "actions": [
      {
        "text": "Quick Start \u2192",
        "link": "/guide/download",
        "type": "primary"
      }
    ],
    "features": [
      {
        "title": "Full Pipeline Integration",
        "details": "Handles everything from searching, queueing downloads, post-processing, and even syncing with your media server!"
      },
      {
        "title": "Fine-Tune Your Searches",
        "details": "Powerful scripting language gives you full control over which releases to filter and which to prioritize"
      },
      {
        "title": "Single Binary",
        "details": "Conductorr is shipped with binaries for all major platforms, and has public Docker images to make setup easy"
      }
    ]
  },
  "excerpt": "",
  "headers": [],
  "filePathRelative": "README.md",
  "git": {
    "updatedTime": 1633915073e3,
    "contributors": [
      {
        "name": "Linus Torvalds",
        "email": "torvalds@linux-foundation.org",
        "commits": 1
      }
    ]
  }
};
var index_html = /* @__PURE__ */ Object.freeze({
  __proto__: null,
  [Symbol.toStringTag]: "Module",
  data: data$d
});
const data$c = {
  "key": "v-0c47f4f8",
  "path": "/csl/examples.html",
  "title": "Examples",
  "lang": "en-US",
  "frontmatter": {},
  "excerpt": "",
  "headers": [],
  "filePathRelative": "csl/examples.md",
  "git": {
    "updatedTime": null,
    "contributors": []
  }
};
var examples_html = /* @__PURE__ */ Object.freeze({
  __proto__: null,
  [Symbol.toStringTag]: "Module",
  data: data$c
});
const data$b = {
  "key": "v-4760ba53",
  "path": "/csl/introduction.html",
  "title": "Introduction",
  "lang": "en-US",
  "frontmatter": {
    "title": "Introduction"
  },
  "excerpt": "",
  "headers": [
    {
      "level": 2,
      "title": "Syntax",
      "slug": "syntax",
      "children": [
        {
          "level": 3,
          "title": "Data Types and Literals",
          "slug": "data-types-and-literals",
          "children": []
        }
      ]
    },
    {
      "level": 2,
      "title": "Built-In Functions",
      "slug": "built-in-functions",
      "children": []
    },
    {
      "level": 2,
      "title": "Extension Functions",
      "slug": "extension-functions",
      "children": []
    }
  ],
  "filePathRelative": "csl/introduction.md",
  "git": {
    "updatedTime": null,
    "contributors": []
  }
};
var introduction_html = /* @__PURE__ */ Object.freeze({
  __proto__: null,
  [Symbol.toStringTag]: "Module",
  data: data$b
});
const data$a = {
  "key": "v-dc33cfc4",
  "path": "/csl/repl.html",
  "title": "REPL",
  "lang": "en-US",
  "frontmatter": {},
  "excerpt": "",
  "headers": [],
  "filePathRelative": "csl/repl.md",
  "git": {
    "updatedTime": null,
    "contributors": []
  }
};
var repl_html = /* @__PURE__ */ Object.freeze({
  __proto__: null,
  [Symbol.toStringTag]: "Module",
  data: data$a
});
const data$9 = {
  "key": "v-4f4ccb8f",
  "path": "/guide/configuration.html",
  "title": "Configuration",
  "lang": "en-US",
  "frontmatter": {},
  "excerpt": "",
  "headers": [],
  "filePathRelative": "guide/configuration.md",
  "git": {
    "updatedTime": null,
    "contributors": []
  }
};
var configuration_html = /* @__PURE__ */ Object.freeze({
  __proto__: null,
  [Symbol.toStringTag]: "Module",
  data: data$9
});
const data$8 = {
  "key": "v-62059836",
  "path": "/guide/download.html",
  "title": "Download and Installation",
  "lang": "en-US",
  "frontmatter": {},
  "excerpt": "",
  "headers": [
    {
      "level": 2,
      "title": "Manual Installation",
      "slug": "manual-installation",
      "children": []
    },
    {
      "level": 2,
      "title": "Docker Compose",
      "slug": "docker-compose",
      "children": []
    }
  ],
  "filePathRelative": "guide/download.md",
  "git": {
    "updatedTime": 1633915073e3,
    "contributors": [
      {
        "name": "Linus Torvalds",
        "email": "torvalds@linux-foundation.org",
        "commits": 1
      }
    ]
  }
};
var download_html = /* @__PURE__ */ Object.freeze({
  __proto__: null,
  [Symbol.toStringTag]: "Module",
  data: data$8
});
const data$7 = {
  "key": "v-8b3d4f7c",
  "path": "/features/calendar.html",
  "title": "Calendar",
  "lang": "en-US",
  "frontmatter": {},
  "excerpt": "",
  "headers": [],
  "filePathRelative": "features/calendar.md",
  "git": {
    "updatedTime": null,
    "contributors": []
  }
};
var calendar_html = /* @__PURE__ */ Object.freeze({
  __proto__: null,
  [Symbol.toStringTag]: "Module",
  data: data$7
});
const data$6 = {
  "key": "v-7e926318",
  "path": "/features/downloaders.html",
  "title": "Downloaders",
  "lang": "en-US",
  "frontmatter": {},
  "excerpt": "",
  "headers": [],
  "filePathRelative": "features/downloaders.md",
  "git": {
    "updatedTime": null,
    "contributors": []
  }
};
var downloaders_html = /* @__PURE__ */ Object.freeze({
  __proto__: null,
  [Symbol.toStringTag]: "Module",
  data: data$6
});
const data$5 = {
  "key": "v-7440cc28",
  "path": "/features/indexers.html",
  "title": "Indexers",
  "lang": "en-US",
  "frontmatter": {},
  "excerpt": "",
  "headers": [],
  "filePathRelative": "features/indexers.md",
  "git": {
    "updatedTime": null,
    "contributors": []
  }
};
var indexers_html = /* @__PURE__ */ Object.freeze({
  __proto__: null,
  [Symbol.toStringTag]: "Module",
  data: data$5
});
const data$4 = {
  "key": "v-318bcd2c",
  "path": "/features/notifications.html",
  "title": "Notifications",
  "lang": "en-US",
  "frontmatter": {},
  "excerpt": "",
  "headers": [],
  "filePathRelative": "features/notifications.md",
  "git": {
    "updatedTime": null,
    "contributors": []
  }
};
var notifications_html = /* @__PURE__ */ Object.freeze({
  __proto__: null,
  [Symbol.toStringTag]: "Module",
  data: data$4
});
const data$3 = {
  "key": "v-0b1540b2",
  "path": "/features/post-processing.html",
  "title": "Post-Processing",
  "lang": "en-US",
  "frontmatter": {},
  "excerpt": "",
  "headers": [],
  "filePathRelative": "features/post-processing.md",
  "git": {
    "updatedTime": null,
    "contributors": []
  }
};
var postProcessing_html = /* @__PURE__ */ Object.freeze({
  __proto__: null,
  [Symbol.toStringTag]: "Module",
  data: data$3
});
const data$2 = {
  "key": "v-3bcd3316",
  "path": "/features/profiles.html",
  "title": "Profiles",
  "lang": "en-US",
  "frontmatter": {},
  "excerpt": "",
  "headers": [],
  "filePathRelative": "features/profiles.md",
  "git": {
    "updatedTime": null,
    "contributors": []
  }
};
var profiles_html = /* @__PURE__ */ Object.freeze({
  __proto__: null,
  [Symbol.toStringTag]: "Module",
  data: data$2
});
const data$1 = {
  "key": "v-5f90410b",
  "path": "/features/syncing.html",
  "title": "Syncing",
  "lang": "en-US",
  "frontmatter": {},
  "excerpt": "",
  "headers": [],
  "filePathRelative": "features/syncing.md",
  "git": {
    "updatedTime": null,
    "contributors": []
  }
};
var syncing_html = /* @__PURE__ */ Object.freeze({
  __proto__: null,
  [Symbol.toStringTag]: "Module",
  data: data$1
});
const data = {
  "key": "v-3706649a",
  "path": "/404.html",
  "title": "",
  "lang": "en-US",
  "frontmatter": {
    "layout": "404"
  },
  "excerpt": "",
  "headers": [],
  "filePathRelative": null,
  "git": {}
};
var _404_html = /* @__PURE__ */ Object.freeze({
  __proto__: null,
  [Symbol.toStringTag]: "Module",
  data
});
const _sfc_main$d = /* @__PURE__ */ vue.defineComponent({
  __ssrInlineRender: true,
  setup(__props) {
    var _a, _b, _c;
    const routeLocale = useRouteLocale();
    const themeLocale = useThemeLocaleData();
    const messages = (_a = themeLocale.value.notFound) != null ? _a : ["Not Found"];
    const getMsg = () => messages[Math.floor(Math.random() * messages.length)];
    const homeLink = (_b = themeLocale.value.home) != null ? _b : routeLocale.value;
    const homeText = (_c = themeLocale.value.backToHome) != null ? _c : "Back to home";
    return (_ctx, _push, _parent, _attrs) => {
      const _component_RouterLink = vue.resolveComponent("RouterLink");
      _push(`<div${serverRenderer.ssrRenderAttrs(vue.mergeProps({ class: "theme-container" }, _attrs))}><div class="theme-default-content"><h1>404</h1><blockquote>${serverRenderer.ssrInterpolate(getMsg())}</blockquote>`);
      _push(serverRenderer.ssrRenderComponent(_component_RouterLink, { to: vue.unref(homeLink) }, {
        default: vue.withCtx((_, _push2, _parent2, _scopeId) => {
          if (_push2) {
            _push2(`${serverRenderer.ssrInterpolate(vue.unref(homeText))}`);
          } else {
            return [
              vue.createTextVNode(vue.toDisplayString(vue.unref(homeText)), 1)
            ];
          }
        }),
        _: 1
      }, _parent));
      _push(`</div></div>`);
    };
  }
});
const _sfc_setup$d = _sfc_main$d.setup;
_sfc_main$d.setup = (props, ctx) => {
  const ssrContext = vue.useSSRContext();
  (ssrContext.modules || (ssrContext.modules = new Set())).add("../../../../../node_modules/.pnpm/@vuepress+theme-default@2.0.0-beta.26/node_modules/@vuepress/theme-default/lib/client/layouts/404.vue");
  return _sfc_setup$d ? _sfc_setup$d(props, ctx) : void 0;
};
var _404 = /* @__PURE__ */ Object.freeze({
  __proto__: null,
  [Symbol.toStringTag]: "Module",
  "default": _sfc_main$d
});
const __default__ = vue.defineComponent({
  inheritAttrs: false
});
function setup(__props) {
  const props = __props;
  const route = vueRouter.useRoute();
  const site = useSiteData();
  const { item } = vue.toRefs(props);
  const hasHttpProtocol = vue.computed(() => shared.isLinkHttp(item.value.link));
  const hasNonHttpProtocal = vue.computed(() => shared.isLinkMailto(item.value.link) || shared.isLinkTel(item.value.link));
  const linkTarget = vue.computed(() => {
    if (hasNonHttpProtocal.value)
      return void 0;
    if (item.value.target)
      return item.value.target;
    if (hasHttpProtocol.value)
      return "_blank";
    return void 0;
  });
  const isBlankTarget = vue.computed(() => linkTarget.value === "_blank");
  const isRouterLink = vue.computed(() => !hasHttpProtocol.value && !hasNonHttpProtocal.value && !isBlankTarget.value);
  const linkRel = vue.computed(() => {
    if (hasNonHttpProtocal.value)
      return void 0;
    if (item.value.rel)
      return item.value.rel;
    if (isBlankTarget.value)
      return "noopener noreferrer";
    return void 0;
  });
  const linkAriaLabel = vue.computed(() => item.value.ariaLabel || item.value.text);
  const shouldBeActiveInSubpath = vue.computed(() => {
    const localeKeys = Object.keys(site.value.locales);
    if (localeKeys.length) {
      return !localeKeys.some((key) => key === item.value.link);
    }
    return item.value.link !== "/";
  });
  const isActiveInSubpath = vue.computed(() => {
    if (!shouldBeActiveInSubpath.value) {
      return false;
    }
    return route.path.startsWith(item.value.link);
  });
  const isActive = vue.computed(() => {
    if (!isRouterLink.value) {
      return false;
    }
    if (item.value.activeMatch) {
      return new RegExp(item.value.activeMatch).test(route.path);
    }
    return isActiveInSubpath.value;
  });
  return (_ctx, _push, _parent, _attrs) => {
    const _component_RouterLink = vue.resolveComponent("RouterLink");
    const _component_OutboundLink = vue.resolveComponent("OutboundLink");
    if (vue.unref(isRouterLink)) {
      _push(serverRenderer.ssrRenderComponent(_component_RouterLink, vue.mergeProps({
        class: ["nav-link", { "router-link-active": vue.unref(isActive) }],
        to: vue.unref(item).link,
        "aria-label": vue.unref(linkAriaLabel)
      }, _ctx.$attrs, _attrs), {
        default: vue.withCtx((_, _push2, _parent2, _scopeId) => {
          if (_push2) {
            serverRenderer.ssrRenderSlot(_ctx.$slots, "before", {}, null, _push2, _parent2, _scopeId);
            _push2(` ${serverRenderer.ssrInterpolate(vue.unref(item).text)} `);
            serverRenderer.ssrRenderSlot(_ctx.$slots, "after", {}, null, _push2, _parent2, _scopeId);
          } else {
            return [
              vue.renderSlot(_ctx.$slots, "before"),
              vue.createTextVNode(" " + vue.toDisplayString(vue.unref(item).text) + " ", 1),
              vue.renderSlot(_ctx.$slots, "after")
            ];
          }
        }),
        _: 3
      }, _parent));
    } else {
      _push(`<a${serverRenderer.ssrRenderAttrs(vue.mergeProps({
        class: "nav-link external",
        href: vue.unref(item).link,
        rel: vue.unref(linkRel),
        target: vue.unref(linkTarget),
        "aria-label": vue.unref(linkAriaLabel)
      }, _ctx.$attrs, _attrs))}>`);
      serverRenderer.ssrRenderSlot(_ctx.$slots, "before", {}, null, _push, _parent);
      _push(` ${serverRenderer.ssrInterpolate(vue.unref(item).text)} `);
      if (vue.unref(isBlankTarget)) {
        _push(serverRenderer.ssrRenderComponent(_component_OutboundLink, null, null, _parent));
      } else {
        _push(`<!---->`);
      }
      serverRenderer.ssrRenderSlot(_ctx.$slots, "after", {}, null, _push, _parent);
      _push(`</a>`);
    }
  };
}
const _sfc_main$c = /* @__PURE__ */ vue.defineComponent(__spreadProps2(__spreadValues2({}, __default__), {
  __ssrInlineRender: true,
  props: {
    item: {
      type: Object,
      required: true
    }
  },
  setup
}));
const _sfc_setup$c = _sfc_main$c.setup;
_sfc_main$c.setup = (props, ctx) => {
  const ssrContext = vue.useSSRContext();
  (ssrContext.modules || (ssrContext.modules = new Set())).add("../../../../../node_modules/.pnpm/@vuepress+theme-default@2.0.0-beta.26/node_modules/@vuepress/theme-default/lib/client/components/NavLink.vue");
  return _sfc_setup$c ? _sfc_setup$c(props, ctx) : void 0;
};
const _sfc_main$b = /* @__PURE__ */ vue.defineComponent({
  __ssrInlineRender: true,
  setup(__props) {
    const frontmatter = usePageFrontmatter();
    const siteLocale = useSiteLocaleData();
    const heroImage = vue.computed(() => {
      if (!frontmatter.value.heroImage) {
        return null;
      }
      return withBase(frontmatter.value.heroImage);
    });
    const heroText = vue.computed(() => {
      if (frontmatter.value.heroText === null) {
        return null;
      }
      return frontmatter.value.heroText || siteLocale.value.title || "Hello";
    });
    const heroAlt = vue.computed(() => frontmatter.value.heroAlt || heroText.value || "hero");
    const tagline = vue.computed(() => {
      if (frontmatter.value.tagline === null) {
        return null;
      }
      return frontmatter.value.tagline || siteLocale.value.description || "Welcome to your VuePress site";
    });
    const actions = vue.computed(() => {
      if (!shared.isArray(frontmatter.value.actions)) {
        return [];
      }
      return frontmatter.value.actions.map(({ text, link, type = "primary" }) => ({
        text,
        link,
        type
      }));
    });
    const features = vue.computed(() => {
      if (shared.isArray(frontmatter.value.features)) {
        return frontmatter.value.features;
      }
      return [];
    });
    const footer = vue.computed(() => frontmatter.value.footer);
    const footerHtml = vue.computed(() => frontmatter.value.footerHtml);
    return (_ctx, _push, _parent, _attrs) => {
      const _component_Content = vue.resolveComponent("Content");
      _push(`<main${serverRenderer.ssrRenderAttrs(vue.mergeProps({
        class: "home",
        "aria-labelledby": vue.unref(heroText) ? "main-title" : void 0
      }, _attrs))}><header class="hero">`);
      if (vue.unref(heroImage)) {
        _push(`<img${serverRenderer.ssrRenderAttr("src", vue.unref(heroImage))}${serverRenderer.ssrRenderAttr("alt", vue.unref(heroAlt))}>`);
      } else {
        _push(`<!---->`);
      }
      if (vue.unref(heroText)) {
        _push(`<h1 id="main-title">${serverRenderer.ssrInterpolate(vue.unref(heroText))}</h1>`);
      } else {
        _push(`<!---->`);
      }
      if (vue.unref(tagline)) {
        _push(`<p class="description">${serverRenderer.ssrInterpolate(vue.unref(tagline))}</p>`);
      } else {
        _push(`<!---->`);
      }
      if (vue.unref(actions).length) {
        _push(`<p class="actions"><!--[-->`);
        serverRenderer.ssrRenderList(vue.unref(actions), (action) => {
          _push(serverRenderer.ssrRenderComponent(_sfc_main$c, {
            key: action.text,
            class: ["action-button", [action.type]],
            item: action
          }, null, _parent));
        });
        _push(`<!--]--></p>`);
      } else {
        _push(`<!---->`);
      }
      _push(`</header>`);
      if (vue.unref(features).length) {
        _push(`<div class="features"><!--[-->`);
        serverRenderer.ssrRenderList(vue.unref(features), (feature) => {
          _push(`<div class="feature"><h2>${serverRenderer.ssrInterpolate(feature.title)}</h2><p>${serverRenderer.ssrInterpolate(feature.details)}</p></div>`);
        });
        _push(`<!--]--></div>`);
      } else {
        _push(`<!---->`);
      }
      _push(`<div class="theme-default-content custom">`);
      _push(serverRenderer.ssrRenderComponent(_component_Content, null, null, _parent));
      _push(`</div>`);
      if (vue.unref(footer)) {
        _push(`<!--[-->`);
        if (vue.unref(footerHtml)) {
          _push(`<div class="footer">${vue.unref(footer)}</div>`);
        } else {
          _push(`<div class="footer">${serverRenderer.ssrInterpolate(vue.unref(footer))}</div>`);
        }
        _push(`<!--]-->`);
      } else {
        _push(`<!---->`);
      }
      _push(`</main>`);
    };
  }
});
const _sfc_setup$b = _sfc_main$b.setup;
_sfc_main$b.setup = (props, ctx) => {
  const ssrContext = vue.useSSRContext();
  (ssrContext.modules || (ssrContext.modules = new Set())).add("../../../../../node_modules/.pnpm/@vuepress+theme-default@2.0.0-beta.26/node_modules/@vuepress/theme-default/lib/client/components/Home.vue");
  return _sfc_setup$b ? _sfc_setup$b(props, ctx) : void 0;
};
const resolveRepoType = (repo) => {
  if (!shared.isLinkHttp(repo) || /github\.com/.test(repo))
    return "GitHub";
  if (/bitbucket\.org/.test(repo))
    return "Bitbucket";
  if (/gitlab\.com/.test(repo))
    return "GitLab";
  if (/gitee\.com/.test(repo))
    return "Gitee";
  return null;
};
const editLinkPatterns = {
  GitHub: ":repo/edit/:branch/:path",
  GitLab: ":repo/-/edit/:branch/:path",
  Gitee: ":repo/edit/:branch/:path",
  Bitbucket: ":repo/src/:branch/:path?mode=edit&spa=0&at=:branch&fileviewer=file-view-default"
};
const resolveEditLink = ({ docsRepo, docsBranch, docsDir, filePathRelative, editLinkPattern }) => {
  const repoType = resolveRepoType(docsRepo);
  let pattern;
  if (editLinkPattern) {
    pattern = editLinkPattern;
  } else if (repoType !== null) {
    pattern = editLinkPatterns[repoType];
  }
  if (!pattern)
    return null;
  return pattern.replace(/:repo/, shared.isLinkHttp(docsRepo) ? docsRepo : `https://github.com/${docsRepo}`).replace(/:branch/, docsBranch).replace(/:path/, shared.removeLeadingSlash(`${shared.removeEndingSlash(docsDir)}/${filePathRelative}`));
};
const _sfc_main$a = /* @__PURE__ */ vue.defineComponent({
  __ssrInlineRender: true,
  setup(__props) {
    return (_ctx, _push, _parent, _attrs) => {
      serverRenderer.ssrRenderSlot(_ctx.$slots, "default", {}, null, _push, _parent);
    };
  }
});
const _sfc_setup$a = _sfc_main$a.setup;
_sfc_main$a.setup = (props, ctx) => {
  const ssrContext = vue.useSSRContext();
  (ssrContext.modules || (ssrContext.modules = new Set())).add("../../../../../node_modules/.pnpm/@vuepress+theme-default@2.0.0-beta.26/node_modules/@vuepress/theme-default/lib/client/components/DropdownTransition.vue");
  return _sfc_setup$a ? _sfc_setup$a(props, ctx) : void 0;
};
const _sfc_main$9 = /* @__PURE__ */ vue.defineComponent({
  __ssrInlineRender: true,
  props: {
    item: {
      type: Object,
      required: true
    }
  },
  setup(__props) {
    const props = __props;
    const { item } = vue.toRefs(props);
    const dropdownAriaLabel = vue.computed(() => item.value.ariaLabel || item.value.text);
    const open = vue.ref(false);
    const route = vueRouter.useRoute();
    vue.watch(() => route.path, () => {
      open.value = false;
    });
    const isLastItemOfArray = (item2, arr) => arr[arr.length - 1] === item2;
    return (_ctx, _push, _parent, _attrs) => {
      _push(`<div${serverRenderer.ssrRenderAttrs(vue.mergeProps({
        class: ["dropdown-wrapper", { open: open.value }]
      }, _attrs))}><button class="dropdown-title" type="button"${serverRenderer.ssrRenderAttr("aria-label", vue.unref(dropdownAriaLabel))}><span class="title">${serverRenderer.ssrInterpolate(vue.unref(item).text)}</span><span class="arrow down"></span></button><button class="mobile-dropdown-title" type="button"${serverRenderer.ssrRenderAttr("aria-label", vue.unref(dropdownAriaLabel))}><span class="title">${serverRenderer.ssrInterpolate(vue.unref(item).text)}</span><span class="${serverRenderer.ssrRenderClass([open.value ? "down" : "right", "arrow"])}"></span></button>`);
      _push(serverRenderer.ssrRenderComponent(_sfc_main$a, null, {
        default: vue.withCtx((_, _push2, _parent2, _scopeId) => {
          if (_push2) {
            _push2(`<ul style="${serverRenderer.ssrRenderStyle(open.value ? null : { display: "none" })}" class="nav-dropdown"${_scopeId}><!--[-->`);
            serverRenderer.ssrRenderList(vue.unref(item).children, (child, index2) => {
              _push2(`<li class="dropdown-item"${_scopeId}>`);
              if (child.children) {
                _push2(`<!--[--><h4 class="dropdown-subtitle"${_scopeId}>`);
                if (child.link) {
                  _push2(serverRenderer.ssrRenderComponent(_sfc_main$c, { item: child }, null, _parent2, _scopeId));
                } else {
                  _push2(`<span${_scopeId}>${serverRenderer.ssrInterpolate(child.text)}</span>`);
                }
                _push2(`</h4><ul class="dropdown-subitem-wrapper"${_scopeId}><!--[-->`);
                serverRenderer.ssrRenderList(child.children, (grandchild) => {
                  _push2(`<li class="dropdown-subitem"${_scopeId}>`);
                  _push2(serverRenderer.ssrRenderComponent(_sfc_main$c, { item: grandchild }, null, _parent2, _scopeId));
                  _push2(`</li>`);
                });
                _push2(`<!--]--></ul><!--]-->`);
              } else {
                _push2(serverRenderer.ssrRenderComponent(_sfc_main$c, { item: child }, null, _parent2, _scopeId));
              }
              _push2(`</li>`);
            });
            _push2(`<!--]--></ul>`);
          } else {
            return [
              vue.withDirectives(vue.createVNode("ul", { class: "nav-dropdown" }, [
                (vue.openBlock(true), vue.createBlock(vue.Fragment, null, vue.renderList(vue.unref(item).children, (child, index2) => {
                  return vue.openBlock(), vue.createBlock("li", {
                    key: child.link || index2,
                    class: "dropdown-item"
                  }, [
                    child.children ? (vue.openBlock(), vue.createBlock(vue.Fragment, { key: 0 }, [
                      vue.createVNode("h4", { class: "dropdown-subtitle" }, [
                        child.link ? (vue.openBlock(), vue.createBlock(_sfc_main$c, {
                          key: 0,
                          item: child,
                          onFocusout: ($event) => isLastItemOfArray(child, vue.unref(item).children) && child.children.length === 0 && (open.value = false)
                        }, null, 8, ["item", "onFocusout"])) : (vue.openBlock(), vue.createBlock("span", { key: 1 }, vue.toDisplayString(child.text), 1))
                      ]),
                      vue.createVNode("ul", { class: "dropdown-subitem-wrapper" }, [
                        (vue.openBlock(true), vue.createBlock(vue.Fragment, null, vue.renderList(child.children, (grandchild) => {
                          return vue.openBlock(), vue.createBlock("li", {
                            key: grandchild.link,
                            class: "dropdown-subitem"
                          }, [
                            vue.createVNode(_sfc_main$c, {
                              item: grandchild,
                              onFocusout: ($event) => isLastItemOfArray(grandchild, child.children) && isLastItemOfArray(child, vue.unref(item).children) && (open.value = false)
                            }, null, 8, ["item", "onFocusout"])
                          ]);
                        }), 128))
                      ])
                    ], 64)) : (vue.openBlock(), vue.createBlock(_sfc_main$c, {
                      key: 1,
                      item: child,
                      onFocusout: ($event) => isLastItemOfArray(child, vue.unref(item).children) && (open.value = false)
                    }, null, 8, ["item", "onFocusout"]))
                  ]);
                }), 128))
              ], 512), [
                [vue.vShow, open.value]
              ])
            ];
          }
        }),
        _: 1
      }, _parent));
      _push(`</div>`);
    };
  }
});
const _sfc_setup$9 = _sfc_main$9.setup;
_sfc_main$9.setup = (props, ctx) => {
  const ssrContext = vue.useSSRContext();
  (ssrContext.modules || (ssrContext.modules = new Set())).add("../../../../../node_modules/.pnpm/@vuepress+theme-default@2.0.0-beta.26/node_modules/@vuepress/theme-default/lib/client/components/DropdownLink.vue");
  return _sfc_setup$9 ? _sfc_setup$9(props, ctx) : void 0;
};
const _sfc_main$8 = /* @__PURE__ */ vue.defineComponent({
  __ssrInlineRender: true,
  setup(__props) {
    const useNavbarSelectLanguage = () => {
      const router = vueRouter.useRouter();
      const routeLocale = useRouteLocale();
      const siteLocale = useSiteLocaleData();
      const themeLocale = useThemeLocaleData();
      return vue.computed(() => {
        var _a, _b;
        const localePaths = Object.keys(siteLocale.value.locales);
        if (localePaths.length < 2) {
          return [];
        }
        const currentPath = router.currentRoute.value.path;
        const currentFullPath = router.currentRoute.value.fullPath;
        const languageDropdown = {
          text: (_a = themeLocale.value.selectLanguageText) != null ? _a : "unkown language",
          ariaLabel: (_b = themeLocale.value.selectLanguageAriaLabel) != null ? _b : "unkown language",
          children: localePaths.map((targetLocalePath) => {
            var _a2, _b2, _c, _d, _e, _f;
            const targetSiteLocale = (_b2 = (_a2 = siteLocale.value.locales) == null ? void 0 : _a2[targetLocalePath]) != null ? _b2 : {};
            const targetThemeLocale = (_d = (_c = themeLocale.value.locales) == null ? void 0 : _c[targetLocalePath]) != null ? _d : {};
            const targetLang = `${targetSiteLocale.lang}`;
            const text = (_e = targetThemeLocale.selectLanguageName) != null ? _e : targetLang;
            let link;
            if (targetLang === siteLocale.value.lang) {
              link = currentFullPath;
            } else {
              const targetLocalePage = currentPath.replace(routeLocale.value, targetLocalePath);
              if (router.getRoutes().some((item) => item.path === targetLocalePage)) {
                link = targetLocalePage;
              } else {
                link = (_f = targetThemeLocale.home) != null ? _f : targetLocalePath;
              }
            }
            return {
              text,
              link
            };
          })
        };
        return [languageDropdown];
      });
    };
    const useNavbarRepo = () => {
      const themeLocale = useThemeLocaleData();
      const repo = vue.computed(() => themeLocale.value.repo);
      const repoType = vue.computed(() => repo.value ? resolveRepoType(repo.value) : null);
      const repoLink = vue.computed(() => {
        if (repo.value && !shared.isLinkHttp(repo.value)) {
          return `https://github.com/${repo.value}`;
        }
        return repo.value;
      });
      const repoLabel = vue.computed(() => {
        if (!repoLink.value)
          return null;
        if (themeLocale.value.repoLabel)
          return themeLocale.value.repoLabel;
        if (repoType.value === null)
          return "Source";
        return repoType.value;
      });
      return vue.computed(() => {
        if (!repoLink.value || !repoLabel.value) {
          return [];
        }
        return [
          {
            text: repoLabel.value,
            link: repoLink.value
          }
        ];
      });
    };
    const resolveNavbarItem = (item) => {
      if (shared.isString(item)) {
        return useNavLink(item);
      }
      if (item.children) {
        return __spreadProps2(__spreadValues2({}, item), {
          children: item.children.map(resolveNavbarItem)
        });
      }
      return item;
    };
    const useNavbarConfig = () => {
      const themeLocale = useThemeLocaleData();
      return vue.computed(() => (themeLocale.value.navbar || []).map(resolveNavbarItem));
    };
    const navbarConfig = useNavbarConfig();
    const navbarSelectLanguage = useNavbarSelectLanguage();
    const navbarRepo = useNavbarRepo();
    const navbarLinks = vue.computed(() => [
      ...navbarConfig.value,
      ...navbarSelectLanguage.value,
      ...navbarRepo.value
    ]);
    return (_ctx, _push, _parent, _attrs) => {
      if (vue.unref(navbarLinks).length) {
        _push(`<nav${serverRenderer.ssrRenderAttrs(vue.mergeProps({ class: "navbar-links" }, _attrs))}><!--[-->`);
        serverRenderer.ssrRenderList(vue.unref(navbarLinks), (item) => {
          _push(`<div class="navbar-links-item">`);
          if (item.children) {
            _push(serverRenderer.ssrRenderComponent(_sfc_main$9, { item }, null, _parent));
          } else {
            _push(serverRenderer.ssrRenderComponent(_sfc_main$c, { item }, null, _parent));
          }
          _push(`</div>`);
        });
        _push(`<!--]--></nav>`);
      } else {
        _push(`<!---->`);
      }
    };
  }
});
const _sfc_setup$8 = _sfc_main$8.setup;
_sfc_main$8.setup = (props, ctx) => {
  const ssrContext = vue.useSSRContext();
  (ssrContext.modules || (ssrContext.modules = new Set())).add("../../../../../node_modules/.pnpm/@vuepress+theme-default@2.0.0-beta.26/node_modules/@vuepress/theme-default/lib/client/components/NavbarLinks.vue");
  return _sfc_setup$8 ? _sfc_setup$8(props, ctx) : void 0;
};
const _sfc_main$7 = /* @__PURE__ */ vue.defineComponent({
  __ssrInlineRender: true,
  setup(__props) {
    const themeLocale = useThemeLocaleData();
    const isDarkMode = useDarkMode();
    return (_ctx, _push, _parent, _attrs) => {
      _push(`<button${serverRenderer.ssrRenderAttrs(vue.mergeProps({
        class: "toggle-dark-button",
        title: vue.unref(themeLocale).toggleDarkMode
      }, _attrs))}><svg style="${serverRenderer.ssrRenderStyle(!vue.unref(isDarkMode) ? null : { display: "none" })}" class="icon" focusable="false" viewBox="0 0 32 32"><path d="M16 12.005a4 4 0 1 1-4 4a4.005 4.005 0 0 1 4-4m0-2a6 6 0 1 0 6 6a6 6 0 0 0-6-6z" fill="currentColor"></path><path d="M5.394 6.813l1.414-1.415l3.506 3.506L8.9 10.318z" fill="currentColor"></path><path d="M2 15.005h5v2H2z" fill="currentColor"></path><path d="M5.394 25.197L8.9 21.691l1.414 1.415l-3.506 3.505z" fill="currentColor"></path><path d="M15 25.005h2v5h-2z" fill="currentColor"></path><path d="M21.687 23.106l1.414-1.415l3.506 3.506l-1.414 1.414z" fill="currentColor"></path><path d="M25 15.005h5v2h-5z" fill="currentColor"></path><path d="M21.687 8.904l3.506-3.506l1.414 1.415l-3.506 3.505z" fill="currentColor"></path><path d="M15 2.005h2v5h-2z" fill="currentColor"></path></svg><svg style="${serverRenderer.ssrRenderStyle(vue.unref(isDarkMode) ? null : { display: "none" })}" class="icon" focusable="false" viewBox="0 0 32 32"><path d="M13.502 5.414a15.075 15.075 0 0 0 11.594 18.194a11.113 11.113 0 0 1-7.975 3.39c-.138 0-.278.005-.418 0a11.094 11.094 0 0 1-3.2-21.584M14.98 3a1.002 1.002 0 0 0-.175.016a13.096 13.096 0 0 0 1.825 25.981c.164.006.328 0 .49 0a13.072 13.072 0 0 0 10.703-5.555a1.01 1.01 0 0 0-.783-1.565A13.08 13.08 0 0 1 15.89 4.38A1.015 1.015 0 0 0 14.98 3z" fill="currentColor"></path></svg></button>`);
    };
  }
});
const _sfc_setup$7 = _sfc_main$7.setup;
_sfc_main$7.setup = (props, ctx) => {
  const ssrContext = vue.useSSRContext();
  (ssrContext.modules || (ssrContext.modules = new Set())).add("../../../../../node_modules/.pnpm/@vuepress+theme-default@2.0.0-beta.26/node_modules/@vuepress/theme-default/lib/client/components/ToggleDarkModeButton.vue");
  return _sfc_setup$7 ? _sfc_setup$7(props, ctx) : void 0;
};
const _sfc_main$6 = /* @__PURE__ */ vue.defineComponent({
  __ssrInlineRender: true,
  emits: ["toggle"],
  setup(__props) {
    const themeLocale = useThemeLocaleData();
    return (_ctx, _push, _parent, _attrs) => {
      _push(`<div${serverRenderer.ssrRenderAttrs(vue.mergeProps({
        class: "toggle-sidebar-button",
        title: vue.unref(themeLocale).toggleSidebar,
        "aria-expanded": "false",
        role: "button",
        tabindex: "0"
      }, _attrs))}><div class="icon" aria-hidden="true"><span></span><span></span><span></span></div></div>`);
    };
  }
});
const _sfc_setup$6 = _sfc_main$6.setup;
_sfc_main$6.setup = (props, ctx) => {
  const ssrContext = vue.useSSRContext();
  (ssrContext.modules || (ssrContext.modules = new Set())).add("../../../../../node_modules/.pnpm/@vuepress+theme-default@2.0.0-beta.26/node_modules/@vuepress/theme-default/lib/client/components/ToggleSidebarButton.vue");
  return _sfc_setup$6 ? _sfc_setup$6(props, ctx) : void 0;
};
const _sfc_main$5 = /* @__PURE__ */ vue.defineComponent({
  __ssrInlineRender: true,
  emits: ["toggle-sidebar"],
  setup(__props) {
    const routeLocale = useRouteLocale();
    const siteLocale = useSiteLocaleData();
    const themeLocale = useThemeLocaleData();
    const isDarkMode = useDarkMode();
    const navbar = vue.ref(null);
    const siteBrand = vue.ref(null);
    const siteBrandLink = vue.computed(() => themeLocale.value.home || routeLocale.value);
    const siteBrandLogo = vue.computed(() => {
      if (isDarkMode.value && themeLocale.value.logoDark !== void 0) {
        return themeLocale.value.logoDark;
      }
      return themeLocale.value.logo;
    });
    const siteBrandTitle = vue.computed(() => siteLocale.value.title);
    const linksWrapperMaxWidth = vue.ref(0);
    const linksWrapperStyle = vue.computed(() => {
      if (!linksWrapperMaxWidth.value) {
        return {};
      }
      return {
        maxWidth: linksWrapperMaxWidth.value + "px"
      };
    });
    const enableDarkMode = vue.computed(() => themeLocale.value.darkMode);
    vue.onMounted(() => {
      const MOBILE_DESKTOP_BREAKPOINT = 719;
      const navbarHorizontalPadding = getCssValue(navbar.value, "paddingLeft") + getCssValue(navbar.value, "paddingRight");
      const handleLinksWrapWidth = () => {
        var _a;
        if (window.innerWidth <= MOBILE_DESKTOP_BREAKPOINT) {
          linksWrapperMaxWidth.value = 0;
        } else {
          linksWrapperMaxWidth.value = navbar.value.offsetWidth - navbarHorizontalPadding - (((_a = siteBrand.value) == null ? void 0 : _a.offsetWidth) || 0);
        }
      };
      handleLinksWrapWidth();
      window.addEventListener("resize", handleLinksWrapWidth, false);
      window.addEventListener("orientationchange", handleLinksWrapWidth, false);
    });
    function getCssValue(el, property) {
      var _a, _b, _c;
      const val = (_c = (_b = (_a = el == null ? void 0 : el.ownerDocument) == null ? void 0 : _a.defaultView) == null ? void 0 : _b.getComputedStyle(el, null)) == null ? void 0 : _c[property];
      const num = Number.parseInt(val, 10);
      return Number.isNaN(num) ? 0 : num;
    }
    return (_ctx, _push, _parent, _attrs) => {
      const _component_RouterLink = vue.resolveComponent("RouterLink");
      const _component_NavbarSearch = vue.resolveComponent("NavbarSearch");
      _push(`<header${serverRenderer.ssrRenderAttrs(vue.mergeProps({
        ref: (_value, _refs) => {
          _refs["navbar"] = _value;
          navbar.value = _value;
        },
        class: "navbar"
      }, _attrs))}>`);
      _push(serverRenderer.ssrRenderComponent(_sfc_main$6, null, null, _parent));
      _push(`<span>`);
      _push(serverRenderer.ssrRenderComponent(_component_RouterLink, { to: vue.unref(siteBrandLink) }, {
        default: vue.withCtx((_, _push2, _parent2, _scopeId) => {
          if (_push2) {
            if (vue.unref(siteBrandLogo)) {
              _push2(`<img class="logo"${serverRenderer.ssrRenderAttr("src", vue.unref(withBase)(vue.unref(siteBrandLogo)))}${serverRenderer.ssrRenderAttr("alt", vue.unref(siteBrandTitle))}${_scopeId}>`);
            } else {
              _push2(`<!---->`);
            }
            if (vue.unref(siteBrandTitle)) {
              _push2(`<span class="${serverRenderer.ssrRenderClass([{ "can-hide": vue.unref(siteBrandLogo) }, "site-name"])}"${_scopeId}>${serverRenderer.ssrInterpolate(vue.unref(siteBrandTitle))}</span>`);
            } else {
              _push2(`<!---->`);
            }
          } else {
            return [
              vue.unref(siteBrandLogo) ? (vue.openBlock(), vue.createBlock("img", {
                key: 0,
                class: "logo",
                src: vue.unref(withBase)(vue.unref(siteBrandLogo)),
                alt: vue.unref(siteBrandTitle)
              }, null, 8, ["src", "alt"])) : vue.createCommentVNode("", true),
              vue.unref(siteBrandTitle) ? (vue.openBlock(), vue.createBlock("span", {
                key: 1,
                class: ["site-name", { "can-hide": vue.unref(siteBrandLogo) }]
              }, vue.toDisplayString(vue.unref(siteBrandTitle)), 3)) : vue.createCommentVNode("", true)
            ];
          }
        }),
        _: 1
      }, _parent));
      _push(`</span><div class="navbar-links-wrapper" style="${serverRenderer.ssrRenderStyle(vue.unref(linksWrapperStyle))}">`);
      serverRenderer.ssrRenderSlot(_ctx.$slots, "before", {}, null, _push, _parent);
      _push(serverRenderer.ssrRenderComponent(_sfc_main$8, { class: "can-hide" }, null, _parent));
      serverRenderer.ssrRenderSlot(_ctx.$slots, "after", {}, null, _push, _parent);
      if (vue.unref(enableDarkMode)) {
        _push(serverRenderer.ssrRenderComponent(_sfc_main$7, null, null, _parent));
      } else {
        _push(`<!---->`);
      }
      _push(serverRenderer.ssrRenderComponent(_component_NavbarSearch, null, null, _parent));
      _push(`</div></header>`);
    };
  }
});
const _sfc_setup$5 = _sfc_main$5.setup;
_sfc_main$5.setup = (props, ctx) => {
  const ssrContext = vue.useSSRContext();
  (ssrContext.modules || (ssrContext.modules = new Set())).add("../../../../../node_modules/.pnpm/@vuepress+theme-default@2.0.0-beta.26/node_modules/@vuepress/theme-default/lib/client/components/Navbar.vue");
  return _sfc_setup$5 ? _sfc_setup$5(props, ctx) : void 0;
};
const _sfc_main$4 = /* @__PURE__ */ vue.defineComponent({
  __ssrInlineRender: true,
  setup(__props) {
    const useEditNavLink = () => {
      const themeLocale2 = useThemeLocaleData();
      const page = usePageData();
      const frontmatter = usePageFrontmatter();
      return vue.computed(() => {
        var _a, _b;
        const showEditLink = (_b = (_a = frontmatter.value.editLink) != null ? _a : themeLocale2.value.editLink) != null ? _b : true;
        if (!showEditLink) {
          return null;
        }
        const {
          repo,
          docsRepo = repo,
          docsBranch = "main",
          docsDir = "",
          editLinkText
        } = themeLocale2.value;
        if (!docsRepo)
          return null;
        const editLink = resolveEditLink({
          docsRepo,
          docsBranch,
          docsDir,
          filePathRelative: page.value.filePathRelative,
          editLinkPattern: themeLocale2.value.editLinkPattern
        });
        if (!editLink)
          return null;
        return {
          text: editLinkText != null ? editLinkText : "Edit this page",
          link: editLink
        };
      });
    };
    const useLastUpdated = () => {
      const siteLocale = useSiteLocaleData();
      const themeLocale2 = useThemeLocaleData();
      const page = usePageData();
      const frontmatter = usePageFrontmatter();
      return vue.computed(() => {
        var _a, _b, _c, _d;
        const showLastUpdated = (_b = (_a = frontmatter.value.lastUpdated) != null ? _a : themeLocale2.value.lastUpdated) != null ? _b : true;
        if (!showLastUpdated)
          return null;
        if (!((_c = page.value.git) == null ? void 0 : _c.updatedTime))
          return null;
        const updatedDate = new Date((_d = page.value.git) == null ? void 0 : _d.updatedTime);
        return updatedDate.toLocaleString(siteLocale.value.lang);
      });
    };
    const useContributors = () => {
      const themeLocale2 = useThemeLocaleData();
      const page = usePageData();
      const frontmatter = usePageFrontmatter();
      return vue.computed(() => {
        var _a, _b, _c, _d;
        const showContributors = (_b = (_a = frontmatter.value.contributors) != null ? _a : themeLocale2.value.contributors) != null ? _b : true;
        if (!showContributors)
          return null;
        return (_d = (_c = page.value.git) == null ? void 0 : _c.contributors) != null ? _d : null;
      });
    };
    const themeLocale = useThemeLocaleData();
    const editNavLink = useEditNavLink();
    const lastUpdated = useLastUpdated();
    const contributors = useContributors();
    return (_ctx, _push, _parent, _attrs) => {
      _push(`<footer${serverRenderer.ssrRenderAttrs(vue.mergeProps({ class: "page-meta" }, _attrs))}>`);
      if (vue.unref(editNavLink)) {
        _push(`<div class="meta-item edit-link">`);
        _push(serverRenderer.ssrRenderComponent(_sfc_main$c, {
          class: "meta-item-label",
          item: vue.unref(editNavLink)
        }, null, _parent));
        _push(`</div>`);
      } else {
        _push(`<!---->`);
      }
      if (vue.unref(lastUpdated)) {
        _push(`<div class="meta-item last-updated"><span class="meta-item-label">${serverRenderer.ssrInterpolate(vue.unref(themeLocale).lastUpdatedText)}: </span><span class="meta-item-info">${serverRenderer.ssrInterpolate(vue.unref(lastUpdated))}</span></div>`);
      } else {
        _push(`<!---->`);
      }
      if (vue.unref(contributors) && vue.unref(contributors).length) {
        _push(`<div class="meta-item contributors"><span class="meta-item-label">${serverRenderer.ssrInterpolate(vue.unref(themeLocale).contributorsText)}: </span><span class="meta-item-info"><!--[-->`);
        serverRenderer.ssrRenderList(vue.unref(contributors), (contributor, index2) => {
          _push(`<!--[--><span class="contributor"${serverRenderer.ssrRenderAttr("title", `email: ${contributor.email}`)}>${serverRenderer.ssrInterpolate(contributor.name)}</span>`);
          if (index2 !== vue.unref(contributors).length - 1) {
            _push(`<!--[-->, <!--]-->`);
          } else {
            _push(`<!---->`);
          }
          _push(`<!--]-->`);
        });
        _push(`<!--]--></span></div>`);
      } else {
        _push(`<!---->`);
      }
      _push(`</footer>`);
    };
  }
});
const _sfc_setup$4 = _sfc_main$4.setup;
_sfc_main$4.setup = (props, ctx) => {
  const ssrContext = vue.useSSRContext();
  (ssrContext.modules || (ssrContext.modules = new Set())).add("../../../../../node_modules/.pnpm/@vuepress+theme-default@2.0.0-beta.26/node_modules/@vuepress/theme-default/lib/client/components/PageMeta.vue");
  return _sfc_setup$4 ? _sfc_setup$4(props, ctx) : void 0;
};
const _sfc_main$3 = /* @__PURE__ */ vue.defineComponent({
  __ssrInlineRender: true,
  setup(__props) {
    const resolveFromFrontmatterConfig = (conf) => {
      if (conf === false) {
        return null;
      }
      if (shared.isString(conf)) {
        return useNavLink(conf);
      }
      if (shared.isPlainObject(conf)) {
        return conf;
      }
      return false;
    };
    const resolveFromSidebarItems = (sidebarItems2, currentPath, offset2) => {
      const index2 = sidebarItems2.findIndex((item) => item.link === currentPath);
      if (index2 !== -1) {
        const targetItem = sidebarItems2[index2 + offset2];
        if (!(targetItem == null ? void 0 : targetItem.link)) {
          return null;
        }
        return targetItem;
      }
      for (const item of sidebarItems2) {
        if (item.children) {
          const childResult = resolveFromSidebarItems(item.children, currentPath, offset2);
          if (childResult) {
            return childResult;
          }
        }
      }
      return null;
    };
    const frontmatter = usePageFrontmatter();
    const sidebarItems = useSidebarItems();
    const route = vueRouter.useRoute();
    const prevNavLink = vue.computed(() => {
      const prevConfig = resolveFromFrontmatterConfig(frontmatter.value.prev);
      if (prevConfig !== false) {
        return prevConfig;
      }
      return resolveFromSidebarItems(sidebarItems.value, route.path, -1);
    });
    const nextNavLink = vue.computed(() => {
      const nextConfig = resolveFromFrontmatterConfig(frontmatter.value.next);
      if (nextConfig !== false) {
        return nextConfig;
      }
      return resolveFromSidebarItems(sidebarItems.value, route.path, 1);
    });
    return (_ctx, _push, _parent, _attrs) => {
      if (vue.unref(prevNavLink) || vue.unref(nextNavLink)) {
        _push(`<nav${serverRenderer.ssrRenderAttrs(vue.mergeProps({ class: "page-nav" }, _attrs))}><p class="inner">`);
        if (vue.unref(prevNavLink)) {
          _push(`<span class="prev"> \u2190 `);
          _push(serverRenderer.ssrRenderComponent(_sfc_main$c, { item: vue.unref(prevNavLink) }, null, _parent));
          _push(`</span>`);
        } else {
          _push(`<!---->`);
        }
        if (vue.unref(nextNavLink)) {
          _push(`<span class="next">`);
          _push(serverRenderer.ssrRenderComponent(_sfc_main$c, { item: vue.unref(nextNavLink) }, null, _parent));
          _push(` \u2192 </span>`);
        } else {
          _push(`<!---->`);
        }
        _push(`</p></nav>`);
      } else {
        _push(`<!---->`);
      }
    };
  }
});
const _sfc_setup$3 = _sfc_main$3.setup;
_sfc_main$3.setup = (props, ctx) => {
  const ssrContext = vue.useSSRContext();
  (ssrContext.modules || (ssrContext.modules = new Set())).add("../../../../../node_modules/.pnpm/@vuepress+theme-default@2.0.0-beta.26/node_modules/@vuepress/theme-default/lib/client/components/PageNav.vue");
  return _sfc_setup$3 ? _sfc_setup$3(props, ctx) : void 0;
};
const _sfc_main$2 = /* @__PURE__ */ vue.defineComponent({
  __ssrInlineRender: true,
  setup(__props) {
    return (_ctx, _push, _parent, _attrs) => {
      const _component_Content = vue.resolveComponent("Content");
      _push(`<main${serverRenderer.ssrRenderAttrs(vue.mergeProps({ class: "page" }, _attrs))}>`);
      serverRenderer.ssrRenderSlot(_ctx.$slots, "top", {}, null, _push, _parent);
      _push(`<div class="theme-default-content">`);
      _push(serverRenderer.ssrRenderComponent(_component_Content, null, null, _parent));
      _push(`</div>`);
      _push(serverRenderer.ssrRenderComponent(_sfc_main$4, null, null, _parent));
      _push(serverRenderer.ssrRenderComponent(_sfc_main$3, null, null, _parent));
      serverRenderer.ssrRenderSlot(_ctx.$slots, "bottom", {}, null, _push, _parent);
      _push(`</main>`);
    };
  }
});
const _sfc_setup$2 = _sfc_main$2.setup;
_sfc_main$2.setup = (props, ctx) => {
  const ssrContext = vue.useSSRContext();
  (ssrContext.modules || (ssrContext.modules = new Set())).add("../../../../../node_modules/.pnpm/@vuepress+theme-default@2.0.0-beta.26/node_modules/@vuepress/theme-default/lib/client/components/Page.vue");
  return _sfc_setup$2 ? _sfc_setup$2(props, ctx) : void 0;
};
const normalizePath = (path) => decodeURI(path).replace(/#.*$/, "").replace(/(index)?\.(md|html)$/, "");
const isActiveLink = (route, link) => {
  if (link === void 0) {
    return false;
  }
  if (route.hash === link) {
    return true;
  }
  const currentPath = normalizePath(route.path);
  const targetPath = normalizePath(link);
  return currentPath === targetPath;
};
const isActiveItem = (route, item) => {
  if (isActiveLink(route, item.link)) {
    return true;
  }
  if (item.children) {
    return item.children.some((child) => isActiveItem(route, child));
  }
  return false;
};
const renderItem = (item, props) => {
  if (item.link) {
    return vue.h(_sfc_main$c, __spreadProps2(__spreadValues2({}, props), {
      item
    }));
  }
  return vue.h("p", props, item.text);
};
const renderChildren = (item, depth) => {
  var _a;
  if (!((_a = item.children) === null || _a === void 0 ? void 0 : _a.length)) {
    return null;
  }
  return vue.h("ul", {
    class: {
      "sidebar-sub-items": depth > 0
    }
  }, item.children.map((child) => vue.h("li", vue.h(SidebarChild, {
    item: child,
    depth: depth + 1
  }))));
};
const SidebarChild = ({ item, depth = 0 }) => {
  const route = vueRouter.useRoute();
  const active = isActiveItem(route, item);
  return [
    renderItem(item, {
      class: {
        "sidebar-heading": depth === 0,
        "sidebar-item": true,
        active
      }
    }),
    renderChildren(item, depth)
  ];
};
SidebarChild.displayName = "SidebarChild";
SidebarChild.props = {
  item: {
    type: Object,
    required: true
  },
  depth: {
    type: Number,
    required: false
  }
};
const _sfc_main$1 = /* @__PURE__ */ vue.defineComponent({
  __ssrInlineRender: true,
  setup(__props) {
    const sidebarItems = useSidebarItems();
    return (_ctx, _push, _parent, _attrs) => {
      _push(`<aside${serverRenderer.ssrRenderAttrs(vue.mergeProps({ class: "sidebar" }, _attrs))}>`);
      _push(serverRenderer.ssrRenderComponent(_sfc_main$8, null, null, _parent));
      serverRenderer.ssrRenderSlot(_ctx.$slots, "top", {}, null, _push, _parent);
      _push(`<ul class="sidebar-links"><!--[-->`);
      serverRenderer.ssrRenderList(vue.unref(sidebarItems), (item) => {
        _push(serverRenderer.ssrRenderComponent(vue.unref(SidebarChild), {
          key: item.link || item.text,
          item
        }, null, _parent));
      });
      _push(`<!--]--></ul>`);
      serverRenderer.ssrRenderSlot(_ctx.$slots, "bottom", {}, null, _push, _parent);
      _push(`</aside>`);
    };
  }
});
const _sfc_setup$1 = _sfc_main$1.setup;
_sfc_main$1.setup = (props, ctx) => {
  const ssrContext = vue.useSSRContext();
  (ssrContext.modules || (ssrContext.modules = new Set())).add("../../../../../node_modules/.pnpm/@vuepress+theme-default@2.0.0-beta.26/node_modules/@vuepress/theme-default/lib/client/components/Sidebar.vue");
  return _sfc_setup$1 ? _sfc_setup$1(props, ctx) : void 0;
};
const _sfc_main = /* @__PURE__ */ vue.defineComponent({
  __ssrInlineRender: true,
  setup(__props) {
    const page = usePageData();
    const frontmatter = usePageFrontmatter();
    const themeLocale = useThemeLocaleData();
    const shouldShowNavbar = vue.computed(() => frontmatter.value.navbar !== false && themeLocale.value.navbar !== false);
    const sidebarItems = useSidebarItems();
    const isSidebarOpen = vue.ref(false);
    const toggleSidebar = (to) => {
      isSidebarOpen.value = typeof to === "boolean" ? to : !isSidebarOpen.value;
    };
    const containerClass = vue.computed(() => [
      {
        "no-navbar": !shouldShowNavbar.value,
        "no-sidebar": !sidebarItems.value.length,
        "sidebar-open": isSidebarOpen.value
      },
      frontmatter.value.pageClass
    ]);
    let unregisterRouterHook;
    vue.onMounted(() => {
      const router = vueRouter.useRouter();
      unregisterRouterHook = router.afterEach(() => {
        toggleSidebar(false);
      });
    });
    vue.onUnmounted(() => {
      unregisterRouterHook();
    });
    return (_ctx, _push, _parent, _attrs) => {
      _push(`<div${serverRenderer.ssrRenderAttrs(vue.mergeProps({
        class: ["theme-container", vue.unref(containerClass)]
      }, _attrs))}>`);
      serverRenderer.ssrRenderSlot(_ctx.$slots, "navbar", {}, () => {
        if (vue.unref(shouldShowNavbar)) {
          _push(serverRenderer.ssrRenderComponent(_sfc_main$5, null, {
            before: vue.withCtx((_, _push2, _parent2, _scopeId) => {
              if (_push2) {
                serverRenderer.ssrRenderSlot(_ctx.$slots, "navbar-before", {}, null, _push2, _parent2, _scopeId);
              } else {
                return [
                  vue.renderSlot(_ctx.$slots, "navbar-before")
                ];
              }
            }),
            after: vue.withCtx((_, _push2, _parent2, _scopeId) => {
              if (_push2) {
                serverRenderer.ssrRenderSlot(_ctx.$slots, "navbar-after", {}, null, _push2, _parent2, _scopeId);
              } else {
                return [
                  vue.renderSlot(_ctx.$slots, "navbar-after")
                ];
              }
            }),
            _: 3
          }, _parent));
        } else {
          _push(`<!---->`);
        }
      }, _push, _parent);
      _push(`<div class="sidebar-mask"></div>`);
      serverRenderer.ssrRenderSlot(_ctx.$slots, "sidebar", {}, () => {
        _push(serverRenderer.ssrRenderComponent(_sfc_main$1, null, {
          top: vue.withCtx((_, _push2, _parent2, _scopeId) => {
            if (_push2) {
              serverRenderer.ssrRenderSlot(_ctx.$slots, "sidebar-top", {}, null, _push2, _parent2, _scopeId);
            } else {
              return [
                vue.renderSlot(_ctx.$slots, "sidebar-top")
              ];
            }
          }),
          bottom: vue.withCtx((_, _push2, _parent2, _scopeId) => {
            if (_push2) {
              serverRenderer.ssrRenderSlot(_ctx.$slots, "sidebar-bottom", {}, null, _push2, _parent2, _scopeId);
            } else {
              return [
                vue.renderSlot(_ctx.$slots, "sidebar-bottom")
              ];
            }
          }),
          _: 3
        }, _parent));
      }, _push, _parent);
      serverRenderer.ssrRenderSlot(_ctx.$slots, "page", {}, () => {
        if (vue.unref(frontmatter).home) {
          _push(serverRenderer.ssrRenderComponent(_sfc_main$b, null, null, _parent));
        } else {
          _push(serverRenderer.ssrRenderComponent(_sfc_main$2, vue.mergeProps({
            key: vue.unref(page).path
          }, _attrs), {
            top: vue.withCtx((_, _push2, _parent2, _scopeId) => {
              if (_push2) {
                serverRenderer.ssrRenderSlot(_ctx.$slots, "page-top", {}, null, _push2, _parent2, _scopeId);
              } else {
                return [
                  vue.renderSlot(_ctx.$slots, "page-top")
                ];
              }
            }),
            bottom: vue.withCtx((_, _push2, _parent2, _scopeId) => {
              if (_push2) {
                serverRenderer.ssrRenderSlot(_ctx.$slots, "page-bottom", {}, null, _push2, _parent2, _scopeId);
              } else {
                return [
                  vue.renderSlot(_ctx.$slots, "page-bottom")
                ];
              }
            }),
            _: 3
          }, _parent));
        }
      }, _push, _parent);
      _push(`</div>`);
    };
  }
});
const _sfc_setup = _sfc_main.setup;
_sfc_main.setup = (props, ctx) => {
  const ssrContext = vue.useSSRContext();
  (ssrContext.modules || (ssrContext.modules = new Set())).add("../../../../../node_modules/.pnpm/@vuepress+theme-default@2.0.0-beta.26/node_modules/@vuepress/theme-default/lib/client/layouts/Layout.vue");
  return _sfc_setup ? _sfc_setup(props, ctx) : void 0;
};
var Layout = /* @__PURE__ */ Object.freeze({
  __proto__: null,
  [Symbol.toStringTag]: "Module",
  "default": _sfc_main
});
exports.createVueApp = createVueApp;
