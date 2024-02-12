import { defineConfig } from "vitepress";

export default defineConfig({
  lang: "en-US",
  title: "Goverter",
  description: "Type-safe converter generation",
  cleanUrls: true,
  sitemap: {
    hostname: "https://goverter.jmattheis.de",
  },
  lastUpdated: true,
  head: [
    ["link", { rel: "icon", type: "image/svg+xml", href: "/favicon.svg" }],
    ["link", { rel: "icon", type: "image/png", href: "/favicon.png" }],
    ["meta", { name: "og:type", content: "website" }],
    ["meta", { name: "og:locale", content: "en" }],
    ["meta", { name: "og:site_name", content: "Goverter" }],
    ["meta", { name: "og:title", content: "Goverter" }],
    [
      "meta",
      { name: "og:description", content: "Type-safe Go converter generation" },
    ],
  ],
  themeConfig: {
    logo: "favicon.svg",
    editLink: {
      pattern: "https://github.com/summit-fi/goverter/tree/main/docs/:path",
    },
    search: {
      provider: "algolia",
      options: {
        appId: "ET81SWAMKQ",
        apiKey: "d21d3398e70912a7e1ef0eee97ee1376",
        indexName: "goverter-jmattheis",
      },
    },
    nav: [
      { text: "Getting Started", link: "/" },
      { text: "Settings", link: "/reference/settings" },
      { text: "FAQ", link: "/faq" },
    ],
    sidebar: [
      {
        text: "Guides",
        items: [
          { text: "Getting Started", link: "/" },
          { text: "Installation", link: "/guide/install" },
          { text: "Error early", link: "/guide/error-early" },
          {
            text: "Output into same package",
            link: "/guide/output-same-package",
          },
          {
            text: "Structs",
            items: [
              { text: "Basics", link: "/guide/struct" },
              { text: "Unexported field", link: "/guide/unexported-field" },
              { text: "Configure Nested", link: "/guide/configure-nested" },
              { text: "Embedded", link: "/guide/embedded-structs" },
            ],
          },
          { text: "Migrations", link: "/guide/migration" },
        ],
      },
      {
        text: "Reference",
        items: [
          { text: "Command Line Interface", link: "/reference/cli" },
          {
            text: "Build Constraint/Tags",
            link: "/reference/build-constraint",
          },
          { text: "Define Settings", link: "/reference/define-settings" },
          { text: "Settings Overview", link: "/reference/settings" },
          {
            text: "Converter",
            collapsed: true,
            items: [
              { text: "converter", link: "/reference/converter" },
              { text: "extend", link: "/reference/extend" },
              { text: "name", link: "/reference/name" },
              { text: "output", link: "/reference/output" },
              { text: "struct", link: "/reference/struct" },
            ],
          },
          {
            text: "Method",
            collapsed: true,
            items: [
              { text: "autoMap", link: "/reference/autoMap" },
              { text: "default", link: "/reference/default" },
              { text: "ignore", link: "/reference/ignore" },
              { text: "map", link: "/reference/map" },
            ],
          },
          {
            text: "Method (inheritable)",
            collapsed: true,
            items: [
              { text: "ignoreMissing", link: "/reference/ignoreMissing" },
              { text: "ignoreUnexported", link: "/reference/ignoreUnexported" },
              { text: "matchIgnoreCase", link: "/reference/matchIgnoreCase" },
              { text: "skipCopySameType", link: "/reference/skipCopySameType" },
              {
                text: "useUnderlyingTypeMethods",
                link: "/reference/useUnderlyingTypeMethods",
              },
              {
                text: "useZeroValueOnPointerInconsistency",
                link: "/reference/useZeroValueOnPointerInconsistency",
              },
              { text: "wrapErrors", link: "/reference/wrapErrors" },
            ],
          },
        ],
      },
      {
        text: "Explanations",
        items: [{ text: "Generation", link: "/explanation/generation" }],
      },
      { text: "FAQ", link: "/faq" },
      { text: "Changelog", link: "/changelog" },
      { text: "Alternatives", link: "/alternatives" },
      { text: "GitHub", link: "https://github.com/jmattheis/goverter" },
    ],
    socialLinks: [
      { icon: "github", link: "https://github.com/jmattheis/goverter" },
    ],
  },
});
