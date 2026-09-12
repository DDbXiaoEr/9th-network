<script setup>
import { onMounted, onUnmounted, ref } from 'vue'
import BaseIcon from './BaseIcon.vue'

const visible = ref(false)
const progress = ref(0)

const onScroll = () => {
  const max = document.documentElement.scrollHeight - window.innerHeight
  progress.value = max > 0 ? (window.scrollY / max) * 100 : 0
  visible.value = window.scrollY > 400
}

const toTop = () => window.scrollTo({ top: 0, behavior: 'smooth' })

onMounted(() => {
  onScroll()
  window.addEventListener('scroll', onScroll, { passive: true })
})
onUnmounted(() => window.removeEventListener('scroll', onScroll))
</script>

<template>
  <transition name="pop">
    <div v-if="visible" class="tools">
      <button class="back" @click="toTop" aria-label="回到顶部">
        <svg class="ring" viewBox="0 0 44 44">
          <circle cx="22" cy="22" r="20" class="ring-bg" />
          <circle
            cx="22"
            cy="22"
            r="20"
            class="ring-fg"
            :style="{ strokeDashoffset: 126 - (126 * progress) / 100 }"
          />
        </svg>
        <BaseIcon name="arrow-up" :size="18" />
      </button>
    </div>
  </transition>
</template>

<style scoped>
.tools {
  position: fixed;
  right: 26px;
  bottom: 30px;
  z-index: 120;
}

.back {
  position: relative;
  display: grid;
  place-items: center;
  width: 46px;
  height: 46px;
  border-radius: 50%;
  border: 1px solid var(--border-strong);
  background: rgba(8, 12, 24, 0.85);
  backdrop-filter: blur(10px);
  color: var(--cyan);
  transition: transform 0.2s ease, box-shadow 0.2s ease;
}

.back:hover {
  transform: translateY(-3px);
  box-shadow: 0 10px 30px rgba(34, 211, 238, 0.3);
}

.ring {
  position: absolute;
  inset: 0;
  width: 100%;
  height: 100%;
  transform: rotate(-90deg);
  fill: none;
  stroke-width: 2;
}

.ring-bg {
  stroke: rgba(255, 255, 255, 0.08);
}

.ring-fg {
  stroke: var(--cyan);
  stroke-linecap: round;
  stroke-dasharray: 126;
  transition: stroke-dashoffset 0.15s linear;
}

.pop-enter-active,
.pop-leave-active {
  transition: opacity 0.25s ease, transform 0.25s ease;
}
.pop-enter-from,
.pop-leave-to {
  opacity: 0;
  transform: translateY(12px) scale(0.9);
}
</style>
