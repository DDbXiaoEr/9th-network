<script setup>
import { computed, onMounted, onUnmounted, ref } from 'vue'
import { config, services } from '../config'
import BaseIcon from './BaseIcon.vue'
import NetworkCanvas from './NetworkCanvas.vue'

const current = ref(0)
let timer

// 终端状态
const terminalVisible = ref(true)
const isDragging = ref(false)
const dragOffset = ref({ x: 0, y: 0 })
const position = ref({ x: 0, y: 0 })
const isInitialPosition = ref(true)

const heroSlides = config.site.heroSlides
const stats = config.site.stats

const heroStats = computed(() =>
  stats.map((item) =>
    item.key === 'services' ? { ...item, value: String(services.value.length) } : item
  )
)

const go = (id) => document.getElementById(id)?.scrollIntoView({ behavior: 'smooth' })

const next = () => {
  current.value = (current.value + 1) % heroSlides.length
}

// 拖动开始
const onDragStart = (e) => {
  isDragging.value = true
  const terminal = document.querySelector('.hero-terminal')
  if (!terminal) return
  
  // 第一次拖动时，获取当前实际位置
  if (isInitialPosition.value) {
    const rect = terminal.getBoundingClientRect()
    position.value = { x: rect.left, y: rect.top }
    isInitialPosition.value = false
  }
  
  dragOffset.value = {
    x: e.clientX - position.value.x,
    y: e.clientY - position.value.y
  }
}

// 拖动中 - 限制在窗口范围内
const onDragMove = (e) => {
  if (!isDragging.value) return
  const terminal = document.querySelector('.hero-terminal')
  if (!terminal) return

  const rect = terminal.getBoundingClientRect()
  const maxX = window.innerWidth - rect.width
  const maxY = window.innerHeight - rect.height

  position.value = {
    x: Math.max(0, Math.min(e.clientX - dragOffset.value.x, maxX)),
    y: Math.max(0, Math.min(e.clientY - dragOffset.value.y, maxY))
  }
}

// 拖动结束
const onDragEnd = () => {
  isDragging.value = false
}

// 关闭终端
const closeTerminal = () => {
  terminalVisible.value = false
}

onMounted(() => {
  timer = setInterval(next, 6000)
  document.addEventListener('mousemove', onDragMove)
  document.addEventListener('mouseup', onDragEnd)
})
onUnmounted(() => {
  clearInterval(timer)
  document.removeEventListener('mousemove', onDragMove)
  document.removeEventListener('mouseup', onDragEnd)
})
</script>

<template>
  <section id="home" class="hero">
    <div class="hero-bg">
      <img
        v-for="(slide, i) in heroSlides"
        :key="i"
        :src="slide.image"
        :class="{ active: i === current }"
        alt=""
      />
      <div class="hero-overlay" />
    </div>

    <NetworkCanvas class="hero-canvas" />

    <div class="container hero-inner">
      <div class="hero-content">
        <span class="eyebrow">{{ heroSlides[current].tag }}</span>
        <h1 class="hero-title">
          {{ heroSlides[current].title }}
        </h1>
        <p class="hero-subtitle">{{ heroSlides[current].subtitle }}</p>
        <p class="hero-desc">{{ heroSlides[current].desc }}</p>

        <div class="hero-actions">
          <button class="btn primary" @click="go('services')">
            <BaseIcon name="spark" :size="18" />
            查看社团服务
          </button>
          <button class="btn ghost" @click="go('about')">
            了解第九网络组
            <BaseIcon name="arrow-right" :size="18" />
          </button>
        </div>

        <div class="hero-dots">
          <button
            v-for="(slide, i) in heroSlides"
            :key="i"
            :class="{ active: i === current }"
            :aria-label="`切换到第 ${i + 1} 张`"
            @click="current = i"
          />
        </div>
      </div>

      <div
        v-if="terminalVisible"
        class="hero-terminal glass"
        :class="{ dragging: isDragging, 'has-moved': !isInitialPosition }"
        :style="isInitialPosition ? {} : { left: position.x + 'px', top: position.y + 'px' }"
      >
        <div class="terminal-bar" @mousedown="onDragStart">
          <span class="dot red" @click.stop="closeTerminal" />
          <span class="dot amber" />
          <span class="dot green" />
          <em>the9@xauat:~</em>
        </div>
        <pre class="terminal-body"><code><span class="c">$</span> whoami
<span class="o">第九网络组 / The 9th Network Team</span>
<span class="c">$</span> cat mission.txt
<span class="o">资源共享 · 共学习 · 共提高 · 共进步</span>
<span class="c">$</span> ls ./services
<span class="k">系统运维  网络安全  开发支持</span>
<span class="k">技术培训  硬件服务  竞赛组队</span>
<span class="c">$</span> ./join --now<span class="cursor">▊</span></code></pre>
      </div>
    </div>

    <div class="container">
      <div class="stats glass">
        <div v-for="item in heroStats" :key="item.label" class="stat">
          <div class="stat-value">
            {{ item.value }}<span>{{ item.suffix }}</span>
          </div>
          <div class="stat-label">{{ item.label }}</div>
        </div>
      </div>
    </div>
  </section>
</template>

<style scoped>
.hero {
  position: relative;
  padding: 140px 0 72px;
  min-height: 100vh;
  display: flex;
  flex-direction: column;
  justify-content: center;
  overflow: hidden;
}

.hero-bg {
  position: absolute;
  inset: 0;
  z-index: -1;
}

.hero-bg img {
  position: absolute;
  inset: 0;
  width: 100%;
  height: 100%;
  object-fit: cover;
  opacity: 0;
  transform: scale(1.06);
  transition: opacity 1.2s ease, transform 7s ease;
  filter: saturate(1) brightness(1.12);
}

.hero-bg img.active {
  opacity: 0.82;
  transform: scale(1);
}

.hero-overlay {
  position: absolute;
  inset: 0;
  background:
    linear-gradient(
      90deg,
      rgba(10, 16, 30, 0.82) 0%,
      rgba(10, 16, 30, 0.48) 42%,
      rgba(10, 16, 30, 0.1) 100%
    ),
    linear-gradient(
      180deg,
      rgba(10, 16, 30, 0.28) 0%,
      rgba(10, 16, 30, 0.08) 38%,
      rgba(10, 16, 30, 0.7) 82%,
      var(--bg) 100%
    );
}

.hero-canvas {
  opacity: 0.5;
  z-index: 0;
}

.hero-inner {
  position: relative;
  z-index: 1;
  display: grid;
  grid-template-columns: 1.15fr 0.85fr;
  gap: 56px;
  align-items: center;
  flex: 1;
}

.hero-content {
  display: flex;
  flex-direction: column;
  align-items: flex-start;
  gap: 18px;
}

.hero-title {
  font-size: clamp(40px, 7vw, 76px);
  line-height: 1.05;
  font-weight: 900;
  letter-spacing: -0.02em;
  background: linear-gradient(120deg, #ffffff 0%, var(--cyan-soft) 55%, var(--violet) 100%);
  -webkit-background-clip: text;
  background-clip: text;
  color: transparent;
}

.hero-subtitle {
  font-size: clamp(16px, 2.2vw, 22px);
  color: var(--cyan-soft);
  font-weight: 600;
  letter-spacing: 0.02em;
}

.hero-desc {
  color: var(--text-dim);
  max-width: 520px;
  font-size: 15px;
}

.hero-actions {
  display: flex;
  flex-wrap: wrap;
  gap: 14px;
  margin-top: 8px;
}

.btn {
  display: inline-flex;
  align-items: center;
  gap: 10px;
  padding: 14px 26px;
  border-radius: 12px;
  font-size: 15px;
  font-weight: 600;
  border: 1px solid transparent;
  transition: transform 0.2s ease, box-shadow 0.25s ease, background 0.25s ease;
}

.btn.primary {
  color: #04121a;
  background: linear-gradient(120deg, var(--cyan), var(--cyan-soft));
  box-shadow: 0 12px 34px rgba(34, 211, 238, 0.35);
}

.btn.primary:hover {
  transform: translateY(-2px);
  box-shadow: 0 18px 44px rgba(34, 211, 238, 0.5);
}

.btn.ghost {
  color: var(--text);
  border-color: var(--border-strong);
  background: rgba(255, 255, 255, 0.03);
}

.btn.ghost:hover {
  transform: translateY(-2px);
  background: rgba(34, 211, 238, 0.1);
}

.hero-dots {
  display: flex;
  gap: 10px;
  margin-top: 10px;
}

.hero-dots button {
  width: 30px;
  height: 4px;
  border-radius: 4px;
  border: 0;
  background: rgba(255, 255, 255, 0.18);
  transition: background 0.25s ease, width 0.25s ease;
}

.hero-dots button.active {
  width: 46px;
  background: var(--cyan);
  box-shadow: 0 0 12px var(--cyan);
}

.hero-terminal {
  position: absolute;
  right: 0;
  top: 50%;
  transform: translateY(-50%);
  z-index: 10;
  width: 480px;
  padding: 0;
  overflow: hidden;
  box-shadow: 0 30px 80px rgba(0, 0, 0, 0.55);
  animation: float 7s ease-in-out infinite;
  user-select: none;
}

.hero-terminal.dragging {
  animation: none;
}

.hero-terminal.has-moved {
  right: auto;
  transform: none;
}

.terminal-bar {
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 14px 18px;
  border-bottom: 1px solid var(--border);
  background: rgba(255, 255, 255, 0.02);
  cursor: grab;
}

.terminal-bar:active {
  cursor: grabbing;
}

.terminal-bar em {
  margin-left: auto;
  font-family: var(--mono);
  font-size: 11px;
  color: var(--text-mute);
  font-style: normal;
}

.dot {
  width: 12px;
  height: 12px;
  border-radius: 50%;
}
.dot.red {
  background: #ff5f57;
  cursor: pointer;
  transition: transform 0.2s ease;
}

.dot.red:hover {
  transform: scale(1.3);
}

.dot.amber {
  background: #febc2e;
}
.dot.green {
  background: #28c840;
}

.terminal-body {
  padding: 26px 24px;
  font-family: var(--mono);
  font-size: 15px;
  line-height: 1.9;
  color: var(--text-dim);
  overflow-x: auto;
}

.terminal-body .c {
  color: var(--green);
}
.terminal-body .o {
  color: var(--text);
}
.terminal-body .k {
  color: var(--cyan-soft);
}
.cursor {
  color: var(--cyan);
  animation: glow-pulse 1s steps(2) infinite;
}

.stats {
  position: relative;
  z-index: 1;
  margin-top: 64px;
  display: grid;
  grid-template-columns: repeat(4, 1fr);
  padding: 28px 12px;
}

.stat {
  text-align: center;
  padding: 6px 12px;
  border-right: 1px solid var(--border);
}
.stat:last-child {
  border-right: 0;
}

.stat-value {
  font-family: var(--mono);
  font-size: clamp(26px, 4vw, 40px);
  font-weight: 700;
  color: var(--text);
  line-height: 1;
}

.stat-value span {
  color: var(--cyan);
  font-size: 0.5em;
  margin-left: 2px;
}

.stat-label {
  margin-top: 8px;
  color: var(--text-mute);
  font-size: 13px;
  letter-spacing: 0.08em;
}

@media (max-width: 900px) {
  .hero-inner {
    grid-template-columns: 1fr;
  }
  .hero-terminal {
    display: none;
  }
}

@media (max-width: 560px) {
  .stats {
    grid-template-columns: repeat(2, 1fr);
    gap: 20px 0;
  }
  .stat:nth-child(2) {
    border-right: 0;
  }
}
</style>
