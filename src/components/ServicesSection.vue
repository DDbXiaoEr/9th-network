<script setup>
import { computed, ref } from 'vue'
import { serviceCategories, services, servicesConfig } from '../config/services'
import BaseIcon from './BaseIcon.vue'

const section = servicesConfig.section
const allLabel = section.allLabel

const activeCategory = ref(allLabel)
const selected = ref(null)

const filtered = computed(() =>
  activeCategory.value === allLabel
    ? services
    : services.filter((item) => item.category === activeCategory.value)
)

const counts = computed(() => {
  const map = { [allLabel]: services.length }
  serviceCategories.slice(1).forEach((cat) => {
    map[cat] = services.filter((s) => s.category === cat).length
  })
  return map
})

const goJoin = () => {
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
      <div v-if="selected" class="modal-mask" @click.self="selected = null">
        <div class="modal glass" v-reveal>
          <button class="close" @click="selected = null" aria-label="关闭">
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
          <button class="modal-btn" @click="goJoin">
            申请该服务
            <BaseIcon name="arrow-right" :size="16" />
          </button>
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
  justify-content: center;
  gap: 10px;
  margin-bottom: 40px;
}

.filter {
  display: inline-flex;
  align-items: center;
  gap: 8px;
  padding: 10px 18px;
  border-radius: 999px;
  border: 1px solid var(--border);
  background: rgba(255, 255, 255, 0.02);
  color: var(--text-dim);
  font-size: 14px;
  transition: all 0.22s ease;
}

.filter:hover {
  color: var(--text);
  border-color: var(--border-strong);
}

.filter.active {
  color: #04121a;
  background: linear-gradient(120deg, var(--cyan), var(--cyan-soft));
  border-color: transparent;
  font-weight: 600;
  box-shadow: 0 8px 26px rgba(34, 211, 238, 0.3);
}

.count {
  font-family: var(--mono);
  font-size: 11px;
  padding: 1px 7px;
  border-radius: 999px;
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
  transition: transform 0.28s ease, border-color 0.28s ease, box-shadow 0.28s ease;
}

.card::before {
  content: '';
  position: absolute;
  inset: 0;
  background: radial-gradient(circle at 100% 0%, rgba(34, 211, 238, 0.14), transparent 55%);
  opacity: 0;
  transition: opacity 0.28s ease;
}

.card:hover {
  transform: translateY(-6px);
  border-color: var(--border-strong);
  box-shadow: 0 22px 50px rgba(0, 0, 0, 0.45);
}

.card:hover::before {
  opacity: 1;
}

.hot {
  position: absolute;
  top: 0;
  right: 0;
  font-family: var(--mono);
  font-size: 10px;
  letter-spacing: 0.2em;
  padding: 4px 12px;
  border-radius: 0 var(--radius) 0 12px;
  background: linear-gradient(120deg, var(--amber), #fb7185);
  color: #1a0f00;
  font-weight: 700;
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
  width: 52px;
  height: 52px;
  border-radius: 14px;
  color: var(--cyan);
  background: linear-gradient(135deg, rgba(34, 211, 238, 0.14), rgba(139, 92, 246, 0.14));
  border: 1px solid var(--border);
}

.card-icon.large {
  width: 64px;
  height: 64px;
  color: var(--cyan-soft);
}

.status {
  font-family: var(--mono);
  font-size: 11px;
  color: var(--text-mute);
  padding: 4px 10px;
  border-radius: 999px;
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
  padding: 3px 9px;
  border-radius: 6px;
  color: var(--text-dim);
  background: rgba(255, 255, 255, 0.04);
  border: 1px solid var(--border);
}

.more {
  display: inline-flex;
  align-items: center;
  gap: 4px;
  font-size: 13px;
  color: var(--cyan);
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
  color: var(--cyan);
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
  background: rgba(3, 5, 12, 0.72);
  backdrop-filter: blur(6px);
}

.modal {
  position: relative;
  width: min(520px, 100%);
  padding: 34px 30px 30px;
  box-shadow: 0 40px 100px rgba(0, 0, 0, 0.6);
}

.close {
  position: absolute;
  top: 16px;
  right: 16px;
  display: grid;
  place-items: center;
  padding: 8px;
  border-radius: 10px;
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
  color: var(--cyan);
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
  border-radius: var(--radius-sm);
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
  border-radius: 12px;
  border: 0;
  font-size: 15px;
  font-weight: 600;
  color: #04121a;
  background: linear-gradient(120deg, var(--cyan), var(--cyan-soft));
  transition: transform 0.2s ease, box-shadow 0.2s ease;
}

.modal-btn:hover {
  transform: translateY(-2px);
  box-shadow: 0 14px 34px rgba(34, 211, 238, 0.4);
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
