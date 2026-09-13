<script setup>
import { ref } from 'vue'
import { config } from '../config'
import BaseIcon from './BaseIcon.vue'

const selected = ref(null)
const resourceArticles = config.site.resourceArticles
const resourceGroups = config.site.resourceGroups

const openByCategory = (category) => {
  const article = resourceArticles.find((a) => a.category === category)
  if (article) selected.value = article
}

const openByTitle = (groupTitle, link) => {
  const article =
    resourceArticles.find((a) => a.title === link) ||
    resourceArticles.find((a) => a.category === groupTitle)
  if (article) selected.value = article
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
          v-for="(group, i) in resourceGroups"
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
      <div v-if="selected" class="reader-mask" @click.self="selected = null">
        <article class="reader glass">
          <button class="close" @click="selected = null" aria-label="关闭">
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
            <p v-for="(para, i) in selected.content" :key="i" class="reader-para">{{ para }}</p>
            <p class="reader-note">更多完整内容，欢迎加入社团一起学习交流。</p>
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
  transition: transform 0.28s ease, border-color 0.28s ease;
}

.group:hover {
  transform: translateY(-6px);
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
  border-radius: 12px;
  background: rgba(255, 255, 255, 0.04);
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
  border-radius: 8px;
  border: 0;
  background: none;
  color: var(--text-dim);
  font-size: 13.5px;
  transition: all 0.2s ease;
}

.group li button:hover {
  color: var(--cyan);
  background: rgba(34, 211, 238, 0.07);
  transform: translateX(4px);
}

.bullet {
  width: 6px;
  height: 6px;
  border-radius: 50%;
  background: var(--cyan);
  box-shadow: 0 0 8px var(--cyan);
  flex-shrink: 0;
}

.group.violet .bullet {
  background: var(--violet);
  box-shadow: 0 0 8px var(--violet);
}
.group.blue .bullet {
  background: var(--blue);
  box-shadow: 0 0 8px var(--blue);
}
.group.amber .bullet {
  background: var(--amber);
  box-shadow: 0 0 8px var(--amber);
}

.reader-mask {
  position: fixed;
  inset: 0;
  z-index: 200;
  display: grid;
  place-items: center;
  padding: 24px;
  background: rgba(10, 16, 30, 0.72);
  backdrop-filter: blur(8px);
}

.reader {
  position: relative;
  width: min(760px, 100%);
  max-height: 86vh;
  overflow-y: auto;
  padding: 0;
}

.reader-cover {
  width: 100%;
  height: 240px;
  object-fit: cover;
  border-radius: var(--radius) var(--radius) 0 0;
  opacity: 0.85;
}

.reader-body {
  padding: 30px 34px 36px;
}

.reader-cat {
  font-family: var(--mono);
  font-size: 11px;
  letter-spacing: 0.2em;
  color: var(--cyan);
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
  border-radius: 6px;
  background: rgba(34, 211, 238, 0.1);
  border: 1px solid var(--border);
  color: var(--cyan-soft);
}

.reader-summary {
  font-size: 15px;
  color: var(--text);
  padding-left: 16px;
  border-left: 3px solid var(--cyan);
  margin-bottom: 20px;
  line-height: 1.8;
}

.reader-para {
  color: var(--text-dim);
  font-size: 14.5px;
  line-height: 1.9;
  margin-bottom: 14px;
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
  border-radius: 10px;
  border: 1px solid rgba(255, 255, 255, 0.18);
  background: rgba(5, 7, 15, 0.6);
  color: #fff;
  transition: all 0.2s ease;
}

.close:hover {
  background: rgba(34, 211, 238, 0.2);
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
