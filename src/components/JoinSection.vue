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
  padding: 34px 36px;
  margin-bottom: 24px;
  overflow: hidden;
  border-color: var(--border-strong);
}

.notice::before {
  content: '';
  position: absolute;
  inset: 0;
  background: linear-gradient(120deg, rgba(34, 211, 238, 0.1), rgba(139, 92, 246, 0.1));
  pointer-events: none;
}

.notice-icon {
  position: relative;
  display: grid;
  place-items: center;
  width: 66px;
  height: 66px;
  border-radius: 18px;
  color: var(--cyan);
  background: rgba(34, 211, 238, 0.12);
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
  letter-spacing: 0.08em;
  padding: 8px 16px;
  border-radius: 999px;
  color: #04121a;
  background: linear-gradient(120deg, var(--cyan), var(--cyan-soft));
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
  transition: transform 0.26s ease, border-color 0.26s ease;
}

.step:hover {
  transform: translateY(-5px);
  border-color: var(--border-strong);
}

.step-index {
  display: block;
  font-family: var(--mono);
  font-size: 26px;
  font-weight: 800;
  line-height: 1;
  margin-bottom: 14px;
  background: linear-gradient(120deg, var(--cyan), var(--violet));
  -webkit-background-clip: text;
  background-clip: text;
  color: transparent;
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
  color: var(--green);
}

.cta-btn {
  display: inline-flex;
  align-items: center;
  gap: 8px;
  padding: 13px 28px;
  border-radius: 12px;
  border: 0;
  font-size: 15px;
  font-weight: 600;
  color: #04121a;
  background: linear-gradient(120deg, var(--cyan), var(--cyan-soft));
  white-space: nowrap;
  transition: transform 0.2s ease, box-shadow 0.2s ease;
}

.cta-btn:hover {
  transform: translateY(-2px);
  box-shadow: 0 16px 40px rgba(34, 211, 238, 0.4);
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
