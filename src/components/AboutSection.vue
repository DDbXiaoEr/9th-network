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
  border-radius: 0;
}

.tabs {
  display: flex;
  flex-wrap: wrap;
  gap: 0;
  padding: 0 8px;
  border-bottom: 1px solid var(--border);
}

.tab {
  padding: 14px 22px;
  border-radius: 0;
  border: 0;
  border-bottom: 2px solid transparent;
  background: transparent;
  color: var(--text-dim);
  font-size: 14px;
  margin-bottom: -1px;
  transition: color 0.2s ease, border-color 0.2s ease;
}

.tab:hover {
  color: var(--text);
}

.tab.active {
  color: var(--signal);
  border-bottom-color: var(--signal);
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
  border-radius: 0;
  border: 1px solid var(--border);
  object-fit: cover;
  max-height: 260px;
  filter: saturate(0.7) sepia(0.12);
}

.media-glow {
  display: none;
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
  padding: 16px 0;
  border-radius: 0;
  border: 0;
  border-bottom: 1px solid var(--border);
  background: none;
}

.role {
  display: block;
  font-family: var(--mono);
  font-size: 12px;
  letter-spacing: 0.08em;
  color: var(--signal);
  margin-bottom: 10px;
}

.names {
  display: flex;
  flex-wrap: wrap;
  gap: 8px;
}

.name {
  padding: 3px 10px;
  border-radius: 0;
  font-size: 13px;
  background: rgba(232, 163, 23, 0.08);
  border: 1px solid var(--border);
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
  color: var(--signal);
  flex-shrink: 0;
}

@media (max-width: 820px) {
  .panel {
    grid-template-columns: 1fr;
    padding: 24px 16px;
  }
}
</style>
