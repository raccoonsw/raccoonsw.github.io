// @ts-check
// Note: type annotations allow type checking and IDEs autocompletion

import type * as Preset from "@docusaurus/preset-classic";
import type { Config } from "@docusaurus/types";
import type * as Plugin from "@docusaurus/types/src/plugin";
import type * as OpenApiPlugin from "docusaurus-plugin-openapi-docs";

const config: Config = {
  title: "OpenAPI",
  tagline: "Dinosaurs are cool",
  url: "https://your-docusaurus-test-site.com",
  baseUrl: "/",
  onBrokenLinks: "throw",
  onBrokenMarkdownLinks: "warn",
  favicon: "img/favicon.ico",

  // GitHub pages deployment config.
  // If you aren't using GitHub pages, you don't need these.
  organizationName: "test-api", // Usually your GitHub org/user name.
  projectName: "test-api", // Usually your repo name.

  presets: [
    [
      "classic",
      {
        docs: {
          sidebarPath: require.resolve("./sidebars.ts"),
          // Please change this to your repo.
          // Remove this to remove the "edit this page" links.
          editUrl:
            "https://github.com/raccoonsw/raccoonsw.github.io/",
          docItemComponent: "@theme/ApiItem", // Derived from docusaurus-theme-openapi
        },
        blog: {
          showReadingTime: true,
          // Please change this to your repo.
          // Remove this to remove the "edit this page" links.
          editUrl:
            "https://github.com/raccoonsw/raccoonsw.github.io/",
          onInlineAuthors: "ignore",
          onUntruncatedBlogPosts: "ignore",
        },
        theme: {
          customCss: require.resolve("./src/css/custom.css"),
        },
      } satisfies Preset.Options,
    ],
  ],

  themeConfig:
    {
      docs: {
        sidebar: {
          hideable: true,
        },
      },
      navbar: {
        title: "OpenAPI",
        logo: {
          alt: "OpenAPI Logo",
          src: "img/logo.svg",
        },
        items: [
          {
            type: "doc",
            docId: "intro",
            position: "left",
            label: "Docs",
          },
          {
            href: "https://github.com/raccoonsw/raccoonsw.github.io/",
            "aria-label": "GitHub",
            position: "right",
            className: "header-github-link",
          },
        ],
      },
      footer: {
        style: "dark",
        copyright: `Copyright © ${new Date().getFullYear()} My Project, Inc. Built with Docusaurus.`,
      },
      prism: {
        additionalLanguages: [
          "ruby",
          "csharp",
          "php",
          "java",
          "powershell",
          "json",
          "bash",
          "dart",
          "objectivec",
          "r",
          "diff"
        ],
      },
      languageTabs: [
        {
          highlight: "python",
          language: "python",
          logoClass: "python",
        },
        {
          highlight: "bash",
          language: "curl",
          logoClass: "curl",
        },
        {
          highlight: "csharp",
          language: "csharp",
          logoClass: "csharp",
        },
        {
          highlight: "go",
          language: "go",
          logoClass: "go",
        },
        {
          highlight: "javascript",
          language: "nodejs",
          logoClass: "nodejs",
        },
        {
          highlight: "ruby",
          language: "ruby",
          logoClass: "ruby",
        },
        {
          highlight: "php",
          language: "php",
          logoClass: "php",
        },
        {
          highlight: "java",
          language: "java",
          logoClass: "java",
          variant: "unirest",
        },
        {
          highlight: "powershell",
          language: "powershell",
          logoClass: "powershell",
        },
        {
          highlight: "dart",
          language: "dart",
          logoClass: "dart",
        },
        {
          highlight: "javascript",
          language: "javascript",
          logoClass: "javascript",
        },
        {
          highlight: "c",
          language: "c",
          logoClass: "c",
        },
        {
          highlight: "objective-c",
          language: "objective-c",
          logoClass: "objective-c",
        },
        {
          highlight: "ocaml",
          language: "ocaml",
          logoClass: "ocaml",
        },
        {
          highlight: "r",
          language: "r",
          logoClass: "r",
        },
        {
          highlight: "swift",
          language: "swift",
          logoClass: "swift",
        },
        {
          highlight: "kotlin",
          language: "kotlin",
          logoClass: "kotlin",
        },
        {
          highlight: "rust",
          language: "rust",
          logoClass: "rust",
        },
      ],
    } satisfies Preset.ThemeConfig,

  plugins: [
    [
      "docusaurus-plugin-openapi-docs",
      {
        id: "openapi",
        docsPluginId: "classic",
        config: {
          payments: {
            specPath: "examples/payments.yaml",
            outputDir: "docs/api-reference/payments",
            downloadUrl:
              "https://raw.githubusercontent.com/PaloAltoNetworks/docusaurus-template-openapi-docs/main/examples/petstore.yaml",
            sidebarOptions: {
              groupPathsBy: "tag",
              categoryLinkSource: "tag",
            },
            version: "1.0.0",
            label: "v1.0.0",
          } satisfies OpenApiPlugin.Options,
        } satisfies Plugin.PluginOptions,
      },
    ],
  ],

  themes: ["docusaurus-theme-openapi-docs"],
};

export default async function createConfig() {
  return config;
}
