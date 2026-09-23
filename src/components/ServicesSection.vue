<script setup>
import { computed, ref } from 'vue'
import { config, serviceCategories, services } from '../config'
import BaseIcon from './BaseIcon.vue'

const section = config.services.section
const allLabel = section.allLabel

const activeCategory = ref(allLabel)
const selected = ref(null)
const prompt = ref(null)

const accessOf = (item) => {
  const access = item?.access || {}
  const type = access.type === 'link' && access.url ? 'link' : 'modal'
  return {
    type,
    url: access.url || '',
    label: access.label || '申请该服务',
    title: access.title || '申请提示',
    note: access.note || '请联系社团管理员了解该服务的申请方式。'
  }
}

const selectedAccess = computed(() => (selected.value ? accessOf(selected.value) : null))

const filtered = computed(() =>
  activeCategory.value === allLabel
    ? services.value
    : services.value.filter((item) => item.category === activeCategory.value)
)

const counts = computed(() => {
  const map = { [allLabel]: services.value.length }
  serviceCategories.value.slice(1).forEach((cat) => {
    map[cat] = services.value.filter((s) => s.category === cat).length
  })
  return map
})

const closeSelected = () => {
  selected.value = null
}

const applyService = () => {
  if (!selected.value) return
  const access = accessOf(selected.value)
  if (access.type === 'link') {
    window.open(access.url, '_blank', 'noopener,noreferrer')
    return
  }
  prompt.value = access
}

const closePrompt = () => {
  prompt.value = null
}

const goJoin = () => {
  prompt.value = null
  selected.value = null
  document.getElementById('join')?.scrollIntoView({ behavior: 'smooth' })
}
</script>

<template>
  <section id="services" class="section services">
    <div class="container">
      <div class="section-head">
        <span class="eyebrow">{{ section.eyebrow }}</span>
        <h2 class="section-title">
          {{ section.title }} <span class="grad">{{ section.highlight }}</span>
        </h2>
        <p class="section-sub">{{ section.subtitle }}</p>
      </div>

      <div class="filters" v-reveal>
        <button
          v-for="cat in serviceCategories"
          :key="cat"
          class="filter"
          :class="{ active: activeCategory === cat }"
          @click="activeCategory = cat"
        >
          {{ cat }}
          <span class="count">{{ counts[cat] }}</span>
        </button>
      </div>

      <div class="grid">
        <article
          v-for="(item, i) in filtered"
          :key="item.id"
          class="card glass"
          v-reveal="{ delay: (i % 3) * 90 }"
          @click="selected = item"
        >
          <span v-if="item.hot" class="hot">HOT</span>
          <div class="card-top">
            <span class="card-icon">
              <BaseIcon :name="item.icon" :size="24" />
            </span>
            <span class="status">{{ item.status }}</span>
          </div>
          <h3 class="card-title">{{ item.title }}</h3>
          <p class="card-desc">{{ item.desc }}</p>
          <div class="card-foot">
            <div class="tags">
              <span v-for="tag in item.tags" :key="tag">{{ tag }}</span>
            </div>
            <span class="more">
              详情
              <BaseIcon name="arrow-right" :size="15" />
            </span>
          </div>
        </article>
      </div>

      <p class="hint">
        {{ section.hint }}<button class="link-btn" @click="goJoin">{{ section.hintAction }}</button>
      </p>
    </div>

    <transition name="fade">
      <div v-if="selected" class="modal-mask" @click.self="closeSelected">
        <div class="modal glass" v-reveal>
          <button class="close" @click="closeSelected" aria-label="关闭">
            <BaseIcon name="close" :size="20" />
          </button>
          <div class="modal-head">
            <span class="card-icon large">
              <BaseIcon :name="selected.icon" :size="30" />
            </span>
            <div>
              <span class="modal-cat">{{ selected.category }}</span>
              <h3>{{ selected.title }}</h3>
            </div>
          </div>
          <p class="modal-desc">{{ selected.desc }}</p>
          <div class="modal-meta">
            <div>
              <span class="meta-label">服务状态</span>
              <span class="meta-value">{{ selected.status }}</span>
            </div>
            <div>
              <span class="meta-label">相关方向</span>
              <span class="meta-value">{{ selected.tags.join(' · ') }}</span>
            </div>
          </div>
          <button class="modal-btn" @click="applyService">
            {{ selectedAccess.label }}
            <BaseIcon :name="selectedAccess.type === 'link' ? 'external' : 'arrow-right'" :size="16" />
          </button>
        </div>
      </div>
    </transition>

    <transition name="fade">
      <div v-if="prompt" class="modal-mask prompt-mask" @click.self="closePrompt">
        <div class="prompt glass" v-reveal>
          <button class="close" @click="closePrompt" aria-label="关闭">
            <BaseIcon name="close" :size="20" />
          </button>
          <h3 class="prompt-title">{{ prompt.title }}</h3>
          <p class="prompt-note">{{ prompt.note }}</p>
          <button class="modal-btn" @click="closePrompt">知道了</button>
        </div>
      </div>
    </transition>
  </section>
</template>

<style scoped>
.services {
  position: relative;
}

.filters {
  display: flex;
  flex-wrap: wrap;
  justify-content: flex-start;
  gap: 10px;
  margin-bottom: 40px;
}

.filter {
  display: inline-flex;
  align-items: center;
  gap: 8px;
  padding: 8px 16px;
  border-radius: 0;
  border: 1px solid var(--border);
  background: transparent;
  color: var(--text-dim);
  font-size: 13px;
  transition: color 0.2s ease, border-color 0.2s ease, background 0.2s ease;
}

.filter:hover {
  color: var(--text);
  border-color: var(--border-strong);
}

.filter.active {
  color: var(--ink);
  background: var(--signal);
  border-color: var(--signal);
  font-weight: 600;
}

.count {
  font-family: var(--mono);
  font-size: 11px;
  padding: 1px 7px;
  border-radius: 0;
  background: rgba(0, 0, 0, 0.18);
}

.grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(300px, 1fr));
  gap: 20px;
}

.card {
  position: relative;
  padding: 26px 24px;
  cursor: pointer;
  overflow: hidden;
  border-radius: 0;
  transition: border-color 0.2s ease, transform 0.2s ease;
}

.card::before {
  display: none;
}

.card:hover {
  transform: translateY(-3px);
  border-color: var(--border-strong);
}

.hot {
  position: absolute;
  top: 0;
  right: 0;
  font-family: var(--mono);
  font-size: 10px;
  letter-spacing: 0.12em;
  padding: 4px 10px;
  border-radius: 0;
  background: var(--brick);
  color: #fff8ee;
  font-weight: 600;
}

.card-top {
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-bottom: 18px;
}

.card-icon {
  display: grid;
  place-items: center;
  width: 48px;
  height: 48px;
  border-radius: 0;
  color: var(--signal);
  background: rgba(232, 163, 23, 0.08);
  border: 1px solid var(--border);
}

.card-icon.large {
  width: 58px;
  height: 58px;
  color: var(--signal-soft);
}

.status {
  font-family: var(--mono);
  font-size: 11px;
  color: var(--text-mute);
  padding: 4px 10px;
  border-radius: 0;
  border: 1px solid var(--border);
}

.card-title {
  font-size: 18px;
  margin-bottom: 10px;
  color: var(--text);
}

.card-desc {
  color: var(--text-dim);
  font-size: 13.5px;
  line-height: 1.7;
  min-height: 68px;
}

.card-foot {
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-top: 18px;
  padding-top: 16px;
  border-top: 1px solid var(--border);
}

.tags {
  display: flex;
  flex-wrap: wrap;
  gap: 6px;
}

.tags span {
  font-size: 11.5px;
  padding: 3px 8px;
  border-radius: 0;
  color: var(--text-dim);
  background: transparent;
  border: 1px solid var(--border);
}

.more {
  display: inline-flex;
  align-items: center;
  gap: 4px;
  font-size: 13px;
  color: var(--signal);
  white-space: nowrap;
}

.hint {
  text-align: center;
  margin-top: 36px;
  color: var(--text-mute);
  font-size: 14px;
}

.link-btn {
  background: none;
  border: 0;
  color: var(--signal);
  font-size: 14px;
  text-decoration: underline;
  text-underline-offset: 3px;
}

.modal-mask {
  position: fixed;
  inset: 0;
  z-index: 200;
  display: grid;
  place-items: center;
  padding: 20px;
  background: rgba(12, 10, 8, 0.78);
  backdrop-filter: blur(6px);
}

.modal {
  position: relative;
  width: min(520px, 100%);
  padding: 34px 30px 30px;
  border-radius: 0;
  box-shadow: 10px 14px 0 rgba(196, 92, 38, 0.22);
}

.close {
  position: absolute;
  top: 16px;
  right: 16px;
  display: grid;
  place-items: center;
  padding: 8px;
  border-radius: 0;
  border: 1px solid var(--border);
  background: rgba(255, 255, 255, 0.03);
  color: var(--text-dim);
  transition: all 0.2s ease;
}

.close:hover {
  color: var(--text);
  border-color: var(--border-strong);
}

.modal-head {
  display: flex;
  align-items: center;
  gap: 18px;
  margin-bottom: 20px;
}

.modal-cat {
  font-family: var(--mono);
  font-size: 11px;
  letter-spacing: 0.18em;
  color: var(--signal);
}

.modal-head h3 {
  font-size: 22px;
  color: var(--text);
}

.modal-desc {
  color: var(--text-dim);
  font-size: 14.5px;
  line-height: 1.8;
  margin-bottom: 22px;
}

.modal-meta {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 14px;
  margin-bottom: 26px;
}

.modal-meta > div {
  padding: 14px;
  border-radius: 0;
  border: 1px solid var(--border);
  background: rgba(255, 255, 255, 0.02);
}

.meta-label {
  display: block;
  font-size: 11px;
  color: var(--text-mute);
  letter-spacing: 0.1em;
  margin-bottom: 6px;
}

.meta-value {
  font-size: 14px;
  color: var(--text);
}

.modal-btn {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  gap: 8px;
  width: 100%;
  padding: 14px;
  border-radius: 0;
  border: 0;
  font-size: 15px;
  font-weight: 600;
  color: var(--ink);
  background: var(--signal);
  transition: background 0.2s ease;
}

.modal-btn:hover {
  background: var(--signal-soft);
}

.prompt-mask {
  z-index: 300;
  background: rgba(12, 10, 8, 0.82);
}

.prompt {
  position: relative;
  width: min(440px, 100%);
  padding: 34px 30px 28px;
  text-align: center;
  border-radius: 0;
  box-shadow: 10px 14px 0 rgba(196, 92, 38, 0.22);
}

.prompt-title {
  font-size: 20px;
  color: var(--text);
  margin-bottom: 14px;
}

.prompt-note {
  color: var(--text-dim);
  font-size: 14.5px;
  line-height: 1.8;
  margin-bottom: 24px;
  white-space: pre-line;
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
  .modal-meta {
    grid-template-columns: 1fr;
  }
}
</style>
