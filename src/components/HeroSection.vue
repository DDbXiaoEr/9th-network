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
const position = ref({ x: 0, y: 0 })

let dragState = null
let rafId = null
let pendingPoint = null

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

// 拖动开始 - 缓存尺寸，避免拖动过程反复触发布局计算
const onDragStart = (e) => {
  const terminal = document.querySelector('.hero-terminal')
  const container = document.querySelector('.hero-inner')
  if (!terminal || !container) return

  isDragging.value = true
  const containerRect = container.getBoundingClientRect()
  const terminalRect = terminal.getBoundingClientRect()

  dragState = {
    containerLeft: containerRect.left,
    containerTop: containerRect.top,
    maxX: containerRect.width - terminalRect.width,
    maxY: containerRect.height - terminalRect.height,
    offsetX: e.clientX - containerRect.left - position.value.x,
    offsetY: e.clientY - containerRect.top - position.value.y
  }
}

// 拖动中 - 按帧节流，只做纯计算
const onDragMove = (e) => {
  if (!isDragging.value || !dragState) return
  pendingPoint = { x: e.clientX, y: e.clientY }
  if (rafId !== null) return

  rafId = requestAnimationFrame(() => {
    rafId = null
    if (!dragState || !pendingPoint) return
    const x = pendingPoint.x - dragState.containerLeft - dragState.offsetX
    const y = pendingPoint.y - dragState.containerTop - dragState.offsetY
    position.value = {
      x: Math.max(0, Math.min(x, dragState.maxX)),
      y: Math.max(0, Math.min(y, dragState.maxY))
    }
  })
}

// 拖动结束
const onDragEnd = () => {
  isDragging.value = false
  dragState = null
  pendingPoint = null
  if (rafId !== null) {
    cancelAnimationFrame(rafId)
    rafId = null
  }
}

// 关闭终端
const closeTerminal = () => {
  terminalVisible.value = false
}

onMounted(() => {
  timer = setInterval(next, 6000)
  document.addEventListener('mousemove', onDragMove)
  document.addEventListener('mouseup', onDragEnd)
  
  // 计算终端初始位置：相对于 .hero-inner 容器右对齐、垂直居中
  const terminal = document.querySelector('.hero-terminal')
  const container = document.querySelector('.hero-inner')
  if (terminal && container) {
    const terminalRect = terminal.getBoundingClientRect()
    const containerRect = container.getBoundingClientRect()
    position.value = {
      x: containerRect.width - terminalRect.width,
      y: (containerRect.height - terminalRect.height) / 2
    }
  }
})
onUnmounted(() => {
  clearInterval(timer)
  if (rafId !== null) cancelAnimationFrame(rafId)
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
        :class="{ dragging: isDragging }"
        :style="{ left: position.x + 'px', top: position.y + 'px' }"
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
  padding: 140px 0 56px;
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
  transform: scale(1.04);
  transition: opacity 1.2s ease, transform 8s ease;
  filter: saturate(0.55) brightness(0.72) sepia(0.22);
}

.hero-bg img.active {
  opacity: 0.9;
  transform: scale(1);
}

.hero-overlay {
  position: absolute;
  inset: 0;
  background:
    linear-gradient(
      105deg,
      rgba(12, 10, 8, 0.92) 0%,
      rgba(12, 10, 8, 0.62) 46%,
      rgba(12, 10, 8, 0.28) 100%
    ),
    linear-gradient(
      180deg,
      rgba(12, 10, 8, 0.2) 0%,
      rgba(12, 10, 8, 0.08) 40%,
      var(--bg) 100%
    );
}

.hero-canvas {
  opacity: 0.28;
  z-index: 0;
}

.hero-inner {
  position: relative;
  z-index: 1;
  display: grid;
  grid-template-columns: 1.1fr 0.9fr;
  gap: 48px;
  align-items: center;
  flex: 1;
}

.hero-content {
  display: flex;
  flex-direction: column;
  align-items: flex-start;
  gap: 16px;
}

.hero-title {
  font-family: var(--serif);
  font-size: clamp(44px, 7.2vw, 80px);
  line-height: 1.08;
  font-weight: 700;
  letter-spacing: 0.04em;
  color: var(--text);
}

.hero-subtitle {
  font-size: clamp(15px, 2vw, 18px);
  color: var(--signal);
  font-weight: 500;
  letter-spacing: 0.12em;
}

.hero-desc {
  color: var(--text-dim);
  max-width: 480px;
  font-size: 15px;
}

.hero-actions {
  display: flex;
  flex-wrap: wrap;
  gap: 12px;
  margin-top: 8px;
}

.btn {
  display: inline-flex;
  align-items: center;
  gap: 10px;
  padding: 13px 22px;
  border-radius: 0;
  font-size: 14px;
  font-weight: 600;
  border: 1px solid transparent;
  transition: background 0.2s ease, color 0.2s ease, border-color 0.2s ease;
}

.btn.primary {
  color: var(--ink);
  background: var(--signal);
  border-color: var(--signal);
}

.btn.primary:hover {
  background: var(--signal-soft);
}

.btn.ghost {
  color: var(--text);
  border-color: var(--border-strong);
  background: transparent;
}

.btn.ghost:hover {
  border-color: var(--signal);
  color: var(--signal);
}

.hero-dots {
  display: flex;
  gap: 8px;
  margin-top: 8px;
}

.hero-dots button {
  width: 22px;
  height: 2px;
  border-radius: 0;
  border: 0;
  background: rgba(255, 255, 255, 0.22);
  transition: background 0.25s ease, width 0.25s ease;
}

.hero-dots button.active {
  width: 40px;
  background: var(--signal);
}

.hero-terminal {
  position: absolute;
  z-index: 10;
  width: 460px;
  padding: 0;
  overflow: hidden;
  border-radius: 2px;
  background: rgba(10, 8, 6, 0.88);
  box-shadow: 8px 12px 0 rgba(196, 92, 38, 0.28);
  user-select: none;
  will-change: left, top;
}

.hero-terminal.dragging {
  box-shadow: 4px 6px 0 rgba(196, 92, 38, 0.4);
}

.terminal-bar {
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 12px 16px;
  border-bottom: 1px solid var(--border);
  background: rgba(232, 163, 23, 0.06);
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
  width: 10px;
  height: 10px;
  border-radius: 1px;
}
.dot.red {
  background: #c45c26;
  cursor: pointer;
  transition: opacity 0.2s ease;
}

.dot.red:hover {
  opacity: 0.7;
}

.dot.amber {
  background: #e8a317;
}
.dot.green {
  background: #9cbc4a;
}

.terminal-body {
  padding: 22px 22px 24px;
  font-family: var(--mono);
  font-size: 13.5px;
  line-height: 1.85;
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
  color: var(--signal-soft);
}
.cursor {
  color: var(--signal);
  animation: blink 1s steps(1) infinite;
}

.stats {
  position: relative;
  z-index: 1;
  margin-top: 56px;
  display: grid;
  grid-template-columns: repeat(4, 1fr);
  padding: 0;
  background: transparent;
  border: 0;
  border-top: 1px solid var(--border);
  border-radius: 0;
  backdrop-filter: none;
}

.stat {
  text-align: left;
  padding: 22px 8px 8px 0;
  border-right: 0;
  border-left: 1px solid var(--border);
  padding-left: 22px;
}
.stat:first-child {
  border-left: 0;
  padding-left: 0;
}

.stat-value {
  font-family: var(--serif);
  font-size: clamp(28px, 4vw, 42px);
  font-weight: 700;
  color: var(--text);
  line-height: 1;
}

.stat-value span {
  color: var(--signal);
  font-size: 0.42em;
  margin-left: 4px;
  font-family: var(--sans);
}

.stat-label {
  margin-top: 8px;
  color: var(--text-mute);
  font-size: 13px;
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
  }
  .stat {
    border-left: 0;
    padding-left: 0;
    padding-top: 16px;
  }
  .stat:nth-child(n + 3) {
    border-top: 1px solid var(--border);
  }
}
</style>
