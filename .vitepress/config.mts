import { DefaultTheme, defineConfig } from 'vitepress'

/**
 * 侧边栏配置: https://vitepress.dev/zh/reference/default-theme-sidebar
 */
const sidebar: DefaultTheme.Sidebar = [
  {
    collapsed: false,
    text: '环境准备',
    base: '/starter/',
    items: [
      { text: 'Go 的安装', link: '/install' },
      { text: 'Go 多版本管理', link: '/multi-version' },
    ]
  },
  {
    collapsed: false,
    text: 'Go 工具链',
    base: '/toolchain/',
    items: [
      { text: 'Go 环境变量', link: '/env' },
      { text: 'Go 包管理', link: '/mod' },
      // { text: '跨服务 全链路追踪', link: '/OpenTelemetry' },
      // { text: '格式化', link: '/gofmt' },
    ]
  },
  {
    collapsed: false,
    text: '语言基础',
    base: '/basics/', // 子文档全部从二级标题，便于合并到一整个文档，`outline: [3,6]`
    items: [
      { text: '函数', link: '/function' },
      { text: '结构体', link: '/struct' },
      { text: '接口', link: '/interface' },
      { text: '异常处理', link: '/exception_handling' },
    ]
  },
  {
    collapsed: false,
    text: '标准库',
    base: '/stdlib/',
    items: [
      { text: '命令行参数解析', link: '/flag' },
    ]
  },
  {
    collapsed: false,
    text: '开源项目',
    base: '/3rdparty/',
    items: [
      { text: 'go-zero 微服务框架', link: '/gozero/gozero' },
    ]
  },
]

/**
 * 主题配置: https://vitepress.dev/zh/reference/default-theme-config
 */
const themeConfig: DefaultTheme.Config = {
  logo: '/go.svg',
  sidebar: sidebar, // 侧边栏配置
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
