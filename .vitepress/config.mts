import { DefaultTheme, defineConfig } from 'vitepress'

/**
 * 侧边栏配置: https://vitepress.dev/zh/reference/default-theme-sidebar
 */
const sidebar: DefaultTheme.Sidebar = [
  {
    text: '首页',
    items: [
      { text: '首页', link: '/' },
    ]
  },
  {
    text: '环境准备',
    items: [
      { text: 'Go 的安装', link: '/install' },
      { text: 'Go 的包管理', link: '/pkg-management' },
    ]
  },
  {
    text: '语言基础',
    items: [
      { text: 'Go 的基础语法', link: '/grammar' },
    ]
  },
]

/**
 * 主题配置: https://vitepress.dev/zh/reference/default-theme-config
 */
const themeConfig: DefaultTheme.Config = {
  // nav: [
  //   { text: 'Home', link: '/' },
  //   { text: 'Examples', link: '/markdown-examples' }
  // ],
  sidebar: sidebar,
  socialLinks: [
    { icon: 'github', link: 'https://henryzhuhr.github.io/hello-go/' }
  ],
  externalLinkIcon: true,
  footer: {
    message: 'Powered By <a href="https://vitepress.dev/">Vitepress</a>',
    copyright: `All rights reserved © 2024-${new Date().getFullYear()} <a href="https://github.com/HenryZhuHR?tab=repositories">HenryZhuHR</a>`
  },
  outline: {
    label: '页面导航'
  },
  lastUpdated: {
    text: '最后更新于',
    formatOptions: {
      dateStyle: 'short',
      timeStyle: 'medium'
    }
  },
  search: {   // 本地搜索: https://vitepress.dev/zh/reference/default-theme-search#local-search
    provider: 'local',
  },
}
// https://vitepress.dev/reference/site-config
export default defineConfig({
  srcDir: 'docs',
  base: '/hello-go/',
  title: "Hello Go",
  description: "Go Learning Log",
  themeConfig: themeConfig,
})
