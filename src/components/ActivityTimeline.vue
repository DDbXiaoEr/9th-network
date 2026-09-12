<script setup>
import { computed } from 'vue'
import { activitiesConfig } from '../config/activities'
import { parseDate, useActivities } from '../composables/useActivities'
import BaseIcon from './BaseIcon.vue'

const { section } = activitiesConfig
const { activities, loading, error } = useActivities()

const upcoming = computed(() =>
  activities.value
    .filter((item) => item.status === 'upcoming')
    .sort((a, b) => parseDate(a.date) - parseDate(b.date))
)

const past = computed(() =>
  activities.value
    .filter((item) => item.status !== 'upcoming')
    .sort((a, b) => parseDate(b.date) - parseDate(a.date))
)

const isEmpty = computed(() => !loading.value && !activities.value.length)
</script>

<template>
  <section id="activities" class="section activities">
    <div class="container">
      <div class="section-head">
        <span class="eyebrow">{{ section.eyebrow }}</span>
        <h2 class="section-title">
          {{ section.title }} <span class="grad">{{ section.highlight }}</span>
        </h2>
        <p class="section-sub">{{ section.subtitle }}</p>
      </div>

      <p v-if="error" class="load-note" v-reveal>
        <BaseIcon name="network" :size="15" />
        在线活动接口加载失败，已展示本地配置数据。
      </p>

      <div v-if="loading" class="skeleton" v-reveal>
        <div v-for="n in 3" :key="n" class="skeleton-item glass" />
      </div>

      <p v-else-if="isEmpty" class="empty glass" v-reveal>{{ section.emptyText }}</p>

      <template v-else>
        <div v-if="upcoming.length" class="group" v-reveal>
          <div class="group-label">
            <span class="pulse" />
            {{ section.upcomingLabel }}
            <span class="count">{{ upcoming.length }}</span>
          </div>
          <ol class="timeline">
            <li v-for="item in upcoming" :key="item.id" class="tl-item">
              <span class="tl-dot upcoming" />
              <div class="tl-card glass">
                <div class="tl-top">
                  <span class="tl-type">{{ item.type }}</span>
                  <span class="tl-status upcoming">即将开始</span>
                </div>
                <h3>{{ item.title }}</h3>
                <p v-if="item.summary" class="tl-summary">{{ item.summary }}</p>
                <div class="tl-meta">
                  <span>
                    <BaseIcon name="calendar" :size="15" />
                    {{ item.date }}
                  </span>
                  <span v-if="item.location">
                    <BaseIcon name="location" :size="15" />
                    {{ item.location }}
                  </span>
                </div>
              </div>
            </li>
          </ol>
        </div>

        <div v-if="past.length" class="group" v-reveal>
          <div class="group-label muted">
            <span class="dot-label" />
            {{ section.pastLabel }}
            <span class="count">{{ past.length }}</span>
          </div>
          <ol class="timeline">
            <li v-for="item in past" :key="item.id" class="tl-item">
              <span class="tl-dot past" />
              <div class="tl-card glass">
                <div class="tl-top">
                  <span class="tl-type">{{ item.type }}</span>
                  <span class="tl-status past">往期</span>
                </div>
                <h3>{{ item.title }}</h3>
                <p v-if="item.summary" class="tl-summary">{{ item.summary }}</p>
                <div class="tl-meta">
                  <span>
                    <BaseIcon name="calendar" :size="15" />
                    {{ item.date }}
                  </span>
                  <span v-if="item.location">
                    <BaseIcon name="location" :size="15" />
                    {{ item.location }}
                  </span>
                </div>
              </div>
            </li>
          </ol>
        </div>
      </template>
    </div>
  </section>
</template>

<style scoped>
.load-note {
  display: inline-flex;
  align-items: center;
  gap: 8px;
  margin: 0 auto 28px;
  padding: 8px 16px;
  border-radius: 999px;
  font-size: 13px;
  color: var(--amber);
  background: rgba(245, 158, 11, 0.1);
  border: 1px solid rgba(245, 158, 11, 0.3);
}

.activities :deep(.container) {
  display: flex;
  flex-direction: column;
  align-items: center;
}

.group {
  width: 100%;
  margin-bottom: 12px;
}

.group-label {
  display: flex;
  align-items: center;
  gap: 10px;
  font-family: var(--mono);
  font-size: 13px;
  letter-spacing: 0.16em;
  color: var(--cyan);
  margin: 0 0 22px 4px;
}

.group-label.muted {
  color: var(--text-dim);
}

.pulse {
  width: 9px;
  height: 9px;
  border-radius: 50%;
  background: var(--cyan);
  box-shadow: 0 0 0 0 rgba(34, 211, 238, 0.6);
  animation: pulse 2s infinite;
}

.dot-label {
  width: 9px;
  height: 9px;
  border-radius: 50%;
  background: var(--text-mute);
}

@keyframes pulse {
  0% {
    box-shadow: 0 0 0 0 rgba(34, 211, 238, 0.55);
  }
  70% {
    box-shadow: 0 0 0 10px rgba(34, 211, 238, 0);
  }
  100% {
    box-shadow: 0 0 0 0 rgba(34, 211, 238, 0);
  }
}

.count {
  padding: 1px 9px;
  border-radius: 999px;
  font-size: 11px;
  background: rgba(255, 255, 255, 0.06);
  border: 1px solid var(--border);
  color: var(--text-dim);
}

.timeline {
  position: relative;
  list-style: none;
  padding-left: 30px;
  margin-bottom: 34px;
}

.timeline::before {
  content: '';
  position: absolute;
  left: 6px;
  top: 8px;
  bottom: 8px;
  width: 2px;
  background: linear-gradient(180deg, var(--cyan), rgba(139, 92, 246, 0.5), transparent);
}

.tl-item {
  position: relative;
  margin-bottom: 20px;
}

.tl-item:last-child {
  margin-bottom: 0;
}

.tl-dot {
  position: absolute;
  left: -30px;
  top: 26px;
  width: 14px;
  height: 14px;
  border-radius: 50%;
  background: var(--bg);
  border: 2px solid var(--text-mute);
  transform: translateX(-50%);
}

.tl-dot.upcoming {
  border-color: var(--cyan);
  box-shadow: 0 0 12px rgba(34, 211, 238, 0.7);
}

.tl-card {
  padding: 22px 24px;
  transition: transform 0.26s ease, border-color 0.26s ease, box-shadow 0.26s ease;
}

.tl-card:hover {
  transform: translateX(6px);
  border-color: var(--border-strong);
  box-shadow: 0 18px 40px rgba(0, 0, 0, 0.4);
}

.tl-top {
  display: flex;
  align-items: center;
  gap: 10px;
  margin-bottom: 12px;
}

.tl-type {
  font-size: 11.5px;
  padding: 3px 10px;
  border-radius: 999px;
  color: var(--text-dim);
  background: rgba(255, 255, 255, 0.04);
  border: 1px solid var(--border);
}

.tl-status {
  font-family: var(--mono);
  font-size: 11px;
  letter-spacing: 0.12em;
  padding: 3px 10px;
  border-radius: 999px;
}

.tl-status.upcoming {
  color: #04121a;
  background: linear-gradient(120deg, var(--cyan), var(--cyan-soft));
}

.tl-status.past {
  color: var(--text-mute);
  border: 1px solid var(--border);
}

.tl-card h3 {
  font-size: 17.5px;
  color: var(--text);
  margin-bottom: 8px;
}

.tl-summary {
  color: var(--text-dim);
  font-size: 13.5px;
  line-height: 1.75;
  margin-bottom: 14px;
}

.tl-meta {
  display: flex;
  flex-wrap: wrap;
  gap: 8px 22px;
  padding-top: 14px;
  border-top: 1px solid var(--border);
}

.tl-meta span {
  display: inline-flex;
  align-items: center;
  gap: 7px;
  font-size: 13px;
  color: var(--text-dim);
}

.tl-meta span :deep(.icon) {
  color: var(--cyan);
}

.empty {
  width: 100%;
  padding: 40px;
  text-align: center;
  color: var(--text-mute);
}

.skeleton {
  width: 100%;
  display: flex;
  flex-direction: column;
  gap: 18px;
}

.skeleton-item {
  height: 118px;
  position: relative;
  overflow: hidden;
}

.skeleton-item::after {
  content: '';
  position: absolute;
  inset: 0;
  background: linear-gradient(90deg, transparent, rgba(110, 231, 249, 0.08), transparent);
  animation: scan 1.4s infinite;
}

@media (max-width: 560px) {
  .timeline {
    padding-left: 22px;
  }
  .tl-dot {
    left: -22px;
  }
  .tl-card {
    padding: 18px 18px;
  }
  .tl-card:hover {
    transform: none;
  }
}
</style>
