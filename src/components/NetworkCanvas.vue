<script setup>
import { onMounted, onUnmounted, ref } from 'vue'

const canvas = ref(null)
let ctx
let raf
let nodes = []
let w = 0
let h = 0

const setup = () => {
  const el = canvas.value
  if (!el) return
  const dpr = Math.min(window.devicePixelRatio || 1, 2)
  w = el.clientWidth
  h = el.clientHeight
  el.width = w * dpr
  el.height = h * dpr
  ctx = el.getContext('2d')
  ctx.scale(dpr, dpr)
  const count = Math.min(70, Math.floor((w * h) / 18000))
  nodes = Array.from({ length: count }, () => ({
    x: Math.random() * w,
    y: Math.random() * h,
    vx: (Math.random() - 0.5) * 0.4,
    vy: (Math.random() - 0.5) * 0.4,
    r: Math.random() * 1.6 + 0.6
  }))
}

const draw = () => {
  if (!ctx) return
  ctx.clearRect(0, 0, w, h)
  for (const n of nodes) {
    n.x += n.vx
    n.y += n.vy
    if (n.x < 0 || n.x > w) n.vx *= -1
    if (n.y < 0 || n.y > h) n.vy *= -1
  }
  for (let i = 0; i < nodes.length; i++) {
    for (let j = i + 1; j < nodes.length; j++) {
      const a = nodes[i]
      const b = nodes[j]
      const dist = Math.hypot(a.x - b.x, a.y - b.y)
      if (dist < 130) {
        ctx.strokeStyle = `rgba(34, 211, 238, ${0.16 * (1 - dist / 130)})`
        ctx.lineWidth = 0.7
        ctx.beginPath()
        ctx.moveTo(a.x, a.y)
        ctx.lineTo(b.x, b.y)
        ctx.stroke()
      }
    }
  }
  for (const n of nodes) {
    ctx.beginPath()
    ctx.arc(n.x, n.y, n.r, 0, Math.PI * 2)
    ctx.fillStyle = 'rgba(103, 232, 249, 0.75)'
    ctx.fill()
  }
  raf = requestAnimationFrame(draw)
}

const onResize = () => {
  setup()
}

onMounted(() => {
  setup()
  draw()
  window.addEventListener('resize', onResize)
})

onUnmounted(() => {
  cancelAnimationFrame(raf)
  window.removeEventListener('resize', onResize)
})
</script>

<template>
  <canvas ref="canvas" class="network-canvas" />
</template>

<style scoped>
.network-canvas {
  position: absolute;
  inset: 0;
  width: 100%;
  height: 100%;
}
</style>
