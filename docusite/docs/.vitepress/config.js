const path = require("path");

/**
 * @type {import('vitepress').UserConfig}
 */
const config = {
  lang: 'en-US',
  title: "Conductorr",
  description: "Smart, fast, and powerful PVR for movies and TV",
  ignoreDeadLinks: true,
  themeConfig: {
    logo: "/logo-rounded.svg",
    sidebar: {
      "/": [
        {
          text: "Guide",
          items: [
            { text: "Install", link: "/guide/install.md" },
            { text: "Configuration", link: "/guide/configuration.md" },
            { text: "Backups", link: "/guide/backups.md" },
          ],
        },
        {
          text: "Features",
          items: [
            { text: "Indexers", link: "/features/indexers.md" },
            { text: "Downloaders", link: "/features/downloaders.md" },
            { text: "Profiles", link: "/features/profiles.md" },
            // {text: "", link: "/features/integrations.md"},
            { text: "Post Processing", link: "/features/post-processing.md" },
            { text: "Mediaservers", link: "/features/mediaservers.md" },
          ],
        },
        {
          text: "CSL",
          items: [
            { text: "Reference", link: "/csl/reference.md" },
            { text: "Examples", link: "/csl/examples.md" },
            { text: "CLI", link: "/csl/cli.md" },
            { text: "REPL", link: "/csl/repl.md" },
          ],
        },
      ],
    },
  },
};

export default config;
