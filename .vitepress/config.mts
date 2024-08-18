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
  {
    text: '语言进阶',
    items: [
      { text: '异常处理', link: '/advance/exception_handling.md' },
    ]
  },
]

/**
 * 主题配置: https://vitepress.dev/zh/reference/default-theme-config
 */
const themeConfig: DefaultTheme.Config = {
  logo: '/go.svg',
  // nav: [
  //   { text: 'Home', link: '/' },
  //   { text: 'Examples', link: '/markdown-examples' }
  // ],
  sidebar: sidebar,
  socialLinks: [
    { icon: 'github', link: 'https://henryzhuhr.github.io/hello-go/' }
  ],
  darkModeSwitchLabel: '外观',          // 用于自定义深色模式开关标签
  lightModeSwitchTitle: '切换到浅色模式', // 用于自定义悬停时显示的浅色模式开关标题
  darkModeSwitchTitle: '切换到深色模式',  // 用于自定义悬停时显示的深色模式开关标题
  returnToTopLabel: '返回顶部',          // 用于自定义返回顶部按钮的标题
  langMenuLabel: '选择语言',             // 用于自定义语言选择菜单的标题
  externalLinkIcon: true,
  docFooter: {
    prev: '⏪️ 上一页',
    next: '下一页 ⏩️'
  },
  footer: {
    message: 'Powered By <a href="https://vitepress.dev/">Vitepress</a>',
    copyright: `All rights reserved © 2024-${new Date().getFullYear()} <a href="https://github.com/HenryZhuHR?tab=repositories">HenryZhuHR</a>`
  },
  outline: {
    label: '页面导航'
  },
  lastUpdated: {
    text: '⏰ 内容最后更新于',
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
  lastUpdated: true,
  vite: {// Vite 配置选项
    publicDir: '../.vitepress/public', // 相对于 docs 目录
  },
})
