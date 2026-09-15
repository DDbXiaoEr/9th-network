/**
 * 站点运行时配置加载器
 * ---------------------------------------------------------------
 * 所有可配置内容优先从站点根目录的 `/config/*.json` 运行时加载，
 * 加载失败或字段缺省时回退到打包进应用的本地默认值：
 *   - site.js      站点通用数据（品牌、导航、页脚、首屏、社团、资源、加入）
 *   - services.js  社团公共服务
 *   - activities.js 活动通知
 *   - directions.js 兴趣方向与 CTF 战队
 *
 * 部署后修改服务器上 `dist/config/*.json` 即可更新内容，无需重新构建。
 * 使用方式：main.js 中 `await loadConfig()` 后再挂载应用。
 */

import { computed, reactive } from 'vue'
import {
  aboutTabs,
  brand,
  footerLinks,
  friendLinks,
  heroSlides,
  icp,
  joinInfo,
  navLinks,
  resourceArticles,
  resourceGroups,
  stats
} from '../data/site'
import { activitiesConfig as activitiesDefaults } from './activities'
import { directionsConfig as directionsDefaults } from './directions'
import { servicesConfig as servicesDefaults } from './services'

const CONFIG_FILES = {
  site: '/config/site.json',
  services: '/config/services.json',
  activities: '/config/activities.json',
  directions: '/config/directions.json'
}

const clone = (value) => JSON.parse(JSON.stringify(value))

const isPlainObject = (value) =>
  value !== null && typeof value === 'object' && !Array.isArray(value)

/**
 * 将 source 深度合并进 target（就地修改，保持引用稳定以维持响应式）。
 * 数组按整体替换处理，对象递归合并，标量直接覆盖。
 */
export const mergeConfig = (target, source) => {
  if (Array.isArray(source)) {
    if (Array.isArray(target)) {
      target.splice(0, target.length, ...clone(source))
      return target
    }
    return clone(source)
  }
  if (isPlainObject(source)) {
    if (!isPlainObject(target)) return clone(source)
    Object.entries(source).forEach(([key, value]) => {
      target[key] = key in target ? mergeConfig(target[key], value) : clone(value)
    })
    return target
  }
  return source
}

export const config = reactive({
  site: {
    brand: clone(brand),
    navLinks: clone(navLinks),
    footerLinks: clone(footerLinks),
    friendLinks: clone(friendLinks),
    icp: icp,
    heroSlides: clone(heroSlides),
    stats: clone(stats),
    aboutTabs: clone(aboutTabs),
    resourceGroups: clone(resourceGroups),
    resourceArticles: clone(resourceArticles),
    joinInfo: clone(joinInfo)
  },
  services: clone(servicesDefaults),
  activities: clone(activitiesDefaults),
  directions: clone(directionsDefaults),
  loaded: false,
  failed: []
})

export async function loadConfig() {
  await Promise.all(
    Object.entries(CONFIG_FILES).map(async ([key, url]) => {
      try {
        const response = await fetch(url, {
          cache: 'no-cache',
          headers: { Accept: 'application/json' }
        })
        if (!response.ok) throw new Error(`HTTP ${response.status}`)
        mergeConfig(config[key], await response.json())
      } catch (err) {
        config.failed.push({ key, url, message: err?.message || String(err) })
        console.warn(`[config] 加载 ${url} 失败，已回退本地默认配置`, err)
      }
    })
  )
  config.loaded = true
  return config
}

const toServiceId = (item, index) =>
  item.id ||
  `${item.category || 'service'}-${index}`.replace(/[^\w-]+/g, '-').toLowerCase()

export const services = computed(() =>
  (config.services.services || []).map((item, index) => ({
    ...item,
    id: toServiceId(item, index),
    tags: item.tags || [],
    hot: Boolean(item.hot)
  }))
)

export const serviceCategories = computed(() => [
  config.services.section.allLabel,
  ...(config.services.categories || [])
])

export default config
