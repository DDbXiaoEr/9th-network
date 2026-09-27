<script setup>
import { computed, ref } from 'vue'
import { config } from '../config'
import { renderMarkdown } from '../lib/markdown'
import BaseIcon from './BaseIcon.vue'

const selected = ref(null)
const markdown = ref('')
const markdownError = ref('')
const resourceArticles = config.articles.articles
const resourceGroups = config.articles.groups
const cache = new Map()

const html = computed(() => renderMarkdown(markdown.value))

const displayGroups = computed(() =>
  resourceGroups.map((group) => {
    const titles = [...(group.links || [])]
    resourceArticles.forEach((article) => {
      if (article.category === group.title && article.title && !titles.includes(article.title)) {
        titles.push(article.title)
      }
    })
    return { ...group, links: titles }
  })
)

const articleFile = (article) => {
  const file = article?.file || (article?.id ? `/articles/${article.id}.md` : '')
  if (!file) return ''
  return file.startsWith('/') ? file : `/${file}`
}

const openArticle = async (article) => {
  if (!article) return
  selected.value = article
  markdownError.value = ''
  const file = articleFile(article)
  if (!file) {
    markdown.value = ''
    markdownError.value = '未配置文章文件'
    return
  }
  if (cache.has(file)) {
    markdown.value = cache.get(file)
    return
  }
  markdown.value = ''
  try {
    const response = await fetch(file, { cache: 'no-cache' })
    if (!response.ok) throw new Error(`HTTP ${response.status}`)
    const text = await response.text()
    cache.set(file, text)
    if (selected.value === article) markdown.value = text
  } catch (err) {
    if (selected.value === article) {
      markdownError.value = '文章加载失败，请稍后重试。'
      console.warn(`[resources] 加载 ${file} 失败`, err)
    }
  }
}

const openByCategory = (category) => {
  openArticle(resourceArticles.find((a) => a.category === category))
}

const openByTitle = (groupTitle, link) => {
  openArticle(
    resourceArticles.find((a) => a.title === link) ||
      resourceArticles.find((a) => a.category === groupTitle)
  )
}

const closeReader = () => {
  selected.value = null
  markdown.value = ''
  markdownError.value = ''
}
</script>

<template>
  <section id="resources" class="section resources">
    <div class="container">
      <div class="section-head">
        <span class="eyebrow">Resources</span>
        <h2 class="section-title">
          学习 <span class="grad">资源</span>
        </h2>
        <p class="section-sub">系统教程、网络安全、校园热点与社团动态，持续更新中。</p>
      </div>

      <div class="group-grid">
        <div
          v-for="(group, i) in displayGroups"
          :key="group.title"
          class="group glass"
          :class="group.accent"
          v-reveal="{ delay: i * 80 }"
        >
          <div class="group-head" @click="openByCategory(group.title)">
            <span class="group-icon">
              <img :src="group.icon" :alt="group.title" />
            </span>
            <h3>{{ group.title }}</h3>
          </div>
          <ul>
            <li v-for="link in group.links" :key="link">
              <button @click="openByTitle(group.title, link)">
                <span class="bullet" />
                {{ link }}
              </button>
            </li>
          </ul>
        </div>
      </div>
    </div>

    <transition name="fade">
      <div v-if="selected" class="reader-mask" @click.self="closeReader">
        <article class="reader glass">
          <button class="close" @click="closeReader" aria-label="关闭">
            <BaseIcon name="close" :size="20" />
          </button>
          <img class="reader-cover" :src="selected.cover" :alt="selected.title" />
          <div class="reader-body">
            <span class="reader-cat">{{ selected.category }}</span>
            <h3>{{ selected.title }}</h3>
            <div class="reader-meta">
              <span>作者 · {{ selected.author }}</span>
              <span class="reader-tags">
                <span v-for="tag in selected.tags" :key="tag">{{ tag }}</span>
              </span>
            </div>
            <p class="reader-summary">{{ selected.summary }}</p>
            <p v-if="markdownError" class="reader-error">{{ markdownError }}</p>
            <div v-else-if="html" class="reader-md" v-html="html" />
            <p class="reader-note">想接触了解更多内容，欢迎加入社团一起学习交流。</p>
          </div>
        </article>
      </div>
    </transition>
  </section>
</template>

<style scoped>
.group-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
  gap: 20px;
}

.group {
  padding: 26px 22px;
  border-radius: 0;
  transition: border-color 0.2s ease;
}

.group:hover {
  border-color: var(--border-strong);
}

.group-head {
  display: flex;
  align-items: center;
  gap: 14px;
  margin-bottom: 20px;
  cursor: pointer;
}

.group-icon {
  display: grid;
  place-items: center;
  width: 46px;
  height: 46px;
  border-radius: 0;
  background: rgba(232, 163, 23, 0.08);
  border: 1px solid var(--border);
}

.group-icon img {
  width: 28px;
  height: 28px;
  border-radius: 6px;
}

.group-head h3 {
  font-size: 17px;
  color: var(--text);
}

.group ul {
  list-style: none;
  display: flex;
  flex-direction: column;
  gap: 4px;
}

.group li button {
  display: flex;
  align-items: center;
  gap: 10px;
  width: 100%;
  text-align: left;
  padding: 10px 8px;
  border-radius: 0;
  border: 0;
  background: none;
  color: var(--text-dim);
  font-size: 13.5px;
  transition: color 0.2s ease, transform 0.2s ease;
}

.group li button:hover {
  color: var(--signal);
  transform: translateX(4px);
}

.bullet {
  width: 6px;
  height: 1px;
  border-radius: 0;
  background: var(--signal);
  flex-shrink: 0;
}

.group.violet .bullet {
  background: var(--brick);
}
.group.blue .bullet {
  background: var(--blue);
}
.group.amber .bullet {
  background: var(--amber);
}

.reader-mask {
  position: fixed;
  inset: 0;
  z-index: 200;
  display: grid;
  place-items: center;
  padding: 24px;
  background: rgba(12, 10, 8, 0.78);
  backdrop-filter: blur(8px);
}

.reader {
  position: relative;
  width: min(760px, 100%);
  max-height: 86vh;
  overflow-y: auto;
  padding: 0;
  border-radius: 0;
}

.reader-cover {
  width: 100%;
  height: 240px;
  object-fit: cover;
  border-radius: 0;
  opacity: 0.85;
  filter: saturate(0.7) sepia(0.12);
}

.reader-body {
  padding: 30px 34px 36px;
}

.reader-cat {
  font-family: var(--mono);
  font-size: 11px;
  letter-spacing: 0.2em;
  color: var(--signal);
}

.reader-body h3 {
  font-size: 26px;
  margin: 8px 0 14px;
  color: var(--text);
}

.reader-meta {
  display: flex;
  flex-wrap: wrap;
  gap: 14px;
  align-items: center;
  color: var(--text-mute);
  font-size: 13px;
  padding-bottom: 18px;
  border-bottom: 1px solid var(--border);
  margin-bottom: 20px;
}

.reader-tags {
  display: flex;
  gap: 6px;
}

.reader-tags span {
  font-size: 11px;
  padding: 3px 9px;
  border-radius: 0;
  background: rgba(232, 163, 23, 0.1);
  border: 1px solid var(--border);
  color: var(--signal-soft);
}

.reader-summary {
  font-size: 15px;
  color: var(--text);
  padding-left: 16px;
  border-left: 2px solid var(--signal);
  margin-bottom: 20px;
  line-height: 1.8;
}

.reader-error {
  color: var(--brick);
  font-size: 14px;
  margin-bottom: 16px;
}

.reader-md {
  color: var(--text-dim);
  font-size: 14.5px;
  line-height: 1.9;
}

.reader-md :deep(h1),
.reader-md :deep(h2),
.reader-md :deep(h3),
.reader-md :deep(h4) {
  color: var(--text);
  font-weight: 600;
  margin: 22px 0 10px;
  line-height: 1.4;
}

.reader-md :deep(h1) { font-size: 22px; }
.reader-md :deep(h2) { font-size: 18px; }
.reader-md :deep(h3) { font-size: 16px; }

.reader-md :deep(p) {
  margin-bottom: 14px;
}

.reader-md :deep(a) {
  color: var(--signal);
  text-decoration: underline;
  text-underline-offset: 3px;
}

.reader-md :deep(ul),
.reader-md :deep(ol) {
  padding-left: 22px;
  margin: 0 0 14px;
}

.reader-md :deep(li) {
  margin-bottom: 6px;
}

.reader-md :deep(blockquote) {
  margin: 0 0 14px;
  padding: 4px 0 4px 14px;
  border-left: 2px solid var(--signal);
  color: var(--text);
}

.reader-md :deep(code) {
  font-family: var(--mono);
  font-size: 13px;
  padding: 1px 6px;
  background: rgba(232, 163, 23, 0.1);
  border: 1px solid var(--border);
  color: var(--signal-soft);
}

.reader-md :deep(pre) {
  margin: 0 0 16px;
  padding: 14px 16px;
  overflow-x: auto;
  background: rgba(12, 10, 8, 0.55);
  border: 1px solid var(--border);
}

.reader-md :deep(pre code) {
  padding: 0;
  background: none;
  border: 0;
  color: var(--text-dim);
}

.reader-md :deep(img) {
  display: block;
  max-width: 100%;
  margin: 12px 0 16px;
  border: 1px solid var(--border);
}

.reader-md :deep(hr) {
  border: 0;
  border-top: 1px solid var(--border);
  margin: 20px 0;
}

.reader-md :deep(table) {
  width: 100%;
  border-collapse: collapse;
  margin: 0 0 16px;
  font-size: 13.5px;
}

.reader-md :deep(th),
.reader-md :deep(td) {
  border: 1px solid var(--border);
  padding: 8px 10px;
  text-align: left;
}

.reader-md :deep(th) {
  color: var(--text);
  background: rgba(232, 163, 23, 0.08);
}

.reader-note {
  margin-top: 22px;
  font-size: 13px;
  color: var(--text-mute);
  font-style: italic;
}

.close {
  position: absolute;
  top: 16px;
  right: 16px;
  z-index: 2;
  display: grid;
  place-items: center;
  padding: 9px;
  border-radius: 0;
  border: 1px solid rgba(255, 255, 255, 0.18);
  background: rgba(12, 10, 8, 0.7);
  color: #fff;
  transition: background 0.2s ease;
}

.close:hover {
  background: rgba(232, 163, 23, 0.28);
}

.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.24s ease;
}
.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}

@media (max-width: 560px) {
  .reader-cover {
    height: 160px;
  }
  .reader-body {
    padding: 24px 20px 28px;
  }
}
</style>
