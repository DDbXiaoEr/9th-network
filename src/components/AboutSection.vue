<script setup>
import { ref } from 'vue'
import { config } from '../config'
import BaseIcon from './BaseIcon.vue'

const active = ref(0)
const aboutTabs = config.site.aboutTabs
</script>

<template>
  <section id="about" class="section">
    <div class="container">
      <div class="section-head">
        <span class="eyebrow">About Us</span>
        <h2 class="section-title">
          关于 <span class="grad">第九网络组</span>
        </h2>
        <p class="section-sub">
          我们是一群热爱计算机科学、有着刻苦专研精神的大学生，在技术里相遇，也在服务中成长。
        </p>
      </div>

      <div class="about glass" v-reveal>
        <div class="tabs">
          <button
            v-for="(tab, i) in aboutTabs"
            :key="tab.key"
            class="tab"
            :class="{ active: active === i }"
            @click="active = i"
          >
            {{ tab.label }}
          </button>
        </div>

        <div class="panel">
          <div class="media">
            <img
              v-for="(img, i) in aboutTabs[active].images"
              :key="i"
              :src="img"
              :alt="aboutTabs[active].label"
            />
            <div class="media-glow" />
          </div>

          <div class="content">
            <template v-if="aboutTabs[active].paragraphs">
              <p v-for="(p, i) in aboutTabs[active].paragraphs" :key="i" class="para">{{ p }}</p>
            </template>

            <template v-else-if="aboutTabs[active].groups">
              <div class="member-grid">
                <div v-for="group in aboutTabs[active].groups" :key="group.role" class="member">
                  <span class="role">{{ group.role }}</span>
                  <div class="names">
                    <span v-for="name in group.names" :key="name" class="name">{{ name }}</span>
                    <span v-if="!group.names.length" class="name muted">筹备中</span>
                  </div>
                </div>
              </div>
            </template>

            <template v-else>
              <ul class="dev-list">
                <li v-for="(item, i) in aboutTabs[active].items" :key="i">
                  <BaseIcon name="check" :size="16" />
                  <span>{{ item }}</span>
                </li>
              </ul>
            </template>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>

<style scoped>
.about {
  padding: 8px;
}

.tabs {
  display: flex;
  flex-wrap: wrap;
  gap: 8px;
  padding: 12px;
  border-bottom: 1px solid var(--border);
}

.tab {
  padding: 10px 20px;
  border-radius: 10px;
  border: 1px solid transparent;
  background: transparent;
  color: var(--text-dim);
  font-size: 14px;
  transition: all 0.22s ease;
}

.tab:hover {
  color: var(--text);
  background: rgba(255, 255, 255, 0.04);
}

.tab.active {
  color: var(--cyan);
  border-color: var(--border-strong);
  background: rgba(34, 211, 238, 0.08);
  box-shadow: 0 0 20px rgba(34, 211, 238, 0.15) inset;
}

.panel {
  display: grid;
  grid-template-columns: 0.8fr 1.2fr;
  gap: 36px;
  padding: 32px 24px;
}

.media {
  position: relative;
  display: flex;
  flex-direction: column;
  gap: 14px;
}

.media img {
  width: 100%;
  border-radius: var(--radius-sm);
  border: 1px solid var(--border);
  object-fit: cover;
  max-height: 260px;
}

.media-glow {
  position: absolute;
  inset: -20%;
  z-index: -1;
  background: radial-gradient(circle at 30% 30%, rgba(34, 211, 238, 0.18), transparent 60%);
}

.content {
  color: var(--text-dim);
  font-size: 15px;
}

.para + .para {
  margin-top: 14px;
}

.member-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(220px, 1fr));
  gap: 16px;
}

.member {
  padding: 16px;
  border-radius: var(--radius-sm);
  border: 1px solid var(--border);
  background: rgba(255, 255, 255, 0.02);
}

.role {
  display: block;
  font-family: var(--mono);
  font-size: 12px;
  letter-spacing: 0.16em;
  color: var(--cyan);
  margin-bottom: 10px;
}

.names {
  display: flex;
  flex-wrap: wrap;
  gap: 8px;
}

.name {
  padding: 4px 12px;
  border-radius: 999px;
  font-size: 13px;
  background: rgba(139, 92, 246, 0.12);
  border: 1px solid rgba(139, 92, 246, 0.3);
  color: var(--text);
}

.name.muted {
  color: var(--text-mute);
  background: rgba(255, 255, 255, 0.03);
  border-color: var(--border);
}

.dev-list {
  list-style: none;
  display: flex;
  flex-direction: column;
  gap: 14px;
}

.dev-list li {
  display: flex;
  gap: 12px;
  align-items: flex-start;
}

.dev-list li :deep(.icon) {
  margin-top: 5px;
  color: var(--cyan);
  flex-shrink: 0;
}

@media (max-width: 820px) {
  .panel {
    grid-template-columns: 1fr;
    padding: 24px 16px;
  }
}
</style>
