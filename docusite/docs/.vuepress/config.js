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
    navbar: [],
    oldnav: [
      {
        text: "Guide",
        link: "/guide/install.md",
      },
      {
        text: "Features",
        link: "/features/indexers.md",
      },
      {
        text: "CSL",
        link: "/csl/reference.md",
      },
    ],
    sidebar:{
      "/": [
        {
          text: "Guide",
          children: ["/guide/install.md", "/guide/configuration.md", "/guide/backups.md"],
        },
        {
          text: "Features",
          children: [
            "/features/indexers.md",
            "/features/downloaders.md",
            "/features/profiles.md",
            // "/features/integrations.md",
            "/features/post-processing.md",
            "/features/mediaservers.md",
          ],
        },
        {
          text: "CSL",
          children: [
            "/csl/reference.md",
            "/csl/examples.md",
            "/csl/cli.md",
            "/csl/repl.md",
          ],
        },
      ],
    },
  },
};
