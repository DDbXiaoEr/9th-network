<script setup>
import { useRouter } from 'vue-router'
import { config } from '../config'
import BaseIcon from './BaseIcon.vue'

const router = useRouter()
const joinInfo = config.site.joinInfo
</script>

<template>
  <section id="join" class="section join">
    <div class="container">
      <div class="section-head">
        <span class="eyebrow">{{ joinInfo.eyebrow }}</span>
        <h2 class="section-title">
          {{ joinInfo.title }} <span class="grad">{{ joinInfo.highlight }}</span>
        </h2>
        <p class="section-sub">{{ joinInfo.subtitle }}</p>
      </div>

      <div class="notice glass" v-reveal>
        <span class="notice-icon">
          <BaseIcon name="users" :size="30" />
        </span>
        <div class="notice-body">
          <h3>{{ joinInfo.notice.title }}</h3>
          <p>{{ joinInfo.notice.desc }}</p>
        </div>
        <span class="notice-badge">无需在线报名</span>
      </div>

      <ol class="steps">
        <li
          v-for="(step, i) in joinInfo.steps"
          :key="step.title"
          class="step glass"
          v-reveal="{ delay: i * 100 }"
        >
          <span class="step-index">{{ String(i + 1).padStart(2, '0') }}</span>
          <h4>{{ step.title }}</h4>
          <p>{{ step.desc }}</p>
        </li>
      </ol>

      <div class="join-foot" v-reveal>
        <p class="note">
          <BaseIcon name="check" :size="16" />
          {{ joinInfo.note }}
        </p>
        <button class="cta-btn" @click="router.push(joinInfo.cta.to)">
          {{ joinInfo.cta.label }}
          <BaseIcon name="arrow-right" :size="16" />
        </button>
      </div>
    </div>
  </section>
</template>

<style scoped>
.join {
  position: relative;
}

.notice {
  position: relative;
  display: flex;
  align-items: center;
  gap: 24px;
  padding: 32px 34px;
  margin-bottom: 24px;
  overflow: hidden;
  border-radius: 0;
  border-color: var(--border-strong);
}

.notice::before {
  display: none;
}

.notice-icon {
  position: relative;
  display: grid;
  place-items: center;
  width: 58px;
  height: 58px;
  border-radius: 0;
  color: var(--signal);
  background: rgba(232, 163, 23, 0.1);
  border: 1px solid var(--border-strong);
  flex-shrink: 0;
}

.notice-body {
  position: relative;
  flex: 1;
}

.notice-body h3 {
  font-size: clamp(18px, 2.4vw, 23px);
  color: var(--text);
  margin-bottom: 8px;
}

.notice-body p {
  color: var(--text-dim);
  font-size: 14.5px;
  line-height: 1.8;
  max-width: 760px;
}

.notice-badge {
  position: relative;
  font-family: var(--mono);
  font-size: 12px;
  letter-spacing: 0.06em;
  padding: 8px 14px;
  border-radius: 0;
  color: var(--ink);
  background: var(--signal);
  white-space: nowrap;
  flex-shrink: 0;
}

.steps {
  list-style: none;
  display: grid;
  grid-template-columns: repeat(4, 1fr);
  gap: 18px;
  margin-bottom: 28px;
}

.step {
  position: relative;
  padding: 26px 22px;
  border-radius: 0;
  transition: border-color 0.2s ease;
}

.step:hover {
  border-color: var(--border-strong);
}

.step-index {
  display: block;
  font-family: var(--serif);
  font-size: 28px;
  font-weight: 700;
  line-height: 1;
  margin-bottom: 14px;
  color: var(--signal);
}

.step h4 {
  font-size: 16px;
  color: var(--text);
  margin-bottom: 8px;
}

.step p {
  color: var(--text-dim);
  font-size: 13.5px;
  line-height: 1.7;
}

.join-foot {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 20px;
  flex-wrap: wrap;
}

.note {
  display: flex;
  align-items: center;
  gap: 8px;
  color: var(--text-mute);
  font-size: 13.5px;
}

.note :deep(.icon) {
  color: var(--signal);
}

.cta-btn {
  display: inline-flex;
  align-items: center;
  gap: 8px;
  padding: 13px 24px;
  border-radius: 0;
  border: 0;
  font-size: 14px;
  font-weight: 600;
  color: var(--ink);
  background: var(--signal);
  white-space: nowrap;
  transition: background 0.2s ease;
}

.cta-btn:hover {
  background: var(--signal-soft);
}

@media (max-width: 900px) {
  .steps {
    grid-template-columns: repeat(2, 1fr);
  }
}

@media (max-width: 640px) {
  .notice {
    flex-direction: column;
    align-items: flex-start;
    gap: 18px;
    padding: 26px 22px;
  }
  .steps {
    grid-template-columns: 1fr;
  }
  .join-foot {
    flex-direction: column;
    align-items: flex-start;
  }
}
</style>
