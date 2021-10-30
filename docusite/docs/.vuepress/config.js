const path = require("path");

module.exports = {
  lang: "en-US",
  title: "Conductorr",
  description: "Smart, fast, and powerful PVR for movies and TV",

  bundler: "@vuepress/bundler-vite",
  bundlerConfig: {
    viteOptions: {
      css: {
        postcss: {
          plugins: [
            require("tailwindcss")("./tailwind.config.js"),
            require("autoprefixer"),
          ],
        },
      },
    },
  },

  plugins: [
    [
      "@vuepress/register-components",
      {
        componentsDir: path.resolve(__dirname, "./components"),
      },
    ],
  ],

  postcss: {
    plugins: [
      require("tailwindcss")("./tailwind.config.js"),
      require("autoprefixer"),
    ],
  },

  head: [["link", { rel: "shortcut icon", href: "/logo.svg" }]],
  themeConfig: {
    logo: "/logo-rounded.svg",
    overrideTheme: 'dark',
    navbar: [
      {
        text: "Guide",
        link: "/guide/",
      },
      {
        text: "Features",
        link: "/features/",
      },
      {
        text: "CSL",
        link: "/csl/",
      },
    ],
    sidebar: {
      "/": [
        {
          text: "Guide",
          children: ["/guide/download.md", "/guide/configuration.md"],
        },
        {
          text: "Features",
          children: [
            "/features/indexers.md",
            "/features/downloaders.md",
            "/features/profiles.md",
            "/features/notifications.md",
            "/features/post-processing.md",
            "/features/syncing.md",
          ],
        },
        {
          text: "CSL",
          children: [
            "/csl/introduction.md",
            "/csl/examples.md",
            "/csl/repl.md",
          ],
        },
      ],
    },
  },
};
