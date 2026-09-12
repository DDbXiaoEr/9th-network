<script setup>
import { nextTick, onMounted, onUnmounted, ref, watch } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { navLinks } from '../data/site'
import BaseIcon from './BaseIcon.vue'

const route = useRoute()
const router = useRouter()

const scrolled = ref(false)
const menuOpen = ref(false)
const active = ref('home')

const sectionLinks = navLinks.filter((link) => link.kind === 'section')

const onScroll = () => {
  scrolled.value = window.scrollY > 24
  if (route.name === 'directions') {
    active.value = 'directions'
    return
  }
  const fromTop = window.scrollY + 140
  let current = 'home'
  sectionLinks.forEach((link) => {
    const el = document.getElementById(link.id)
    if (el && el.offsetTop <= fromTop) current = link.id
  })
  active.value = current
}

const scrollToSection = (id) =>
  document.getElementById(id)?.scrollIntoView({ behavior: 'smooth' })

const go = async (id) => {
  menuOpen.value = false
  if (route.name !== 'home') {
    await router.push('/')
    await nextTick()
    setTimeout(() => scrollToSection(id), 80)
  } else {
    scrollToSection(id)
  }
}

const goLink = (link) => {
  if (link.kind === 'route') {
    menuOpen.value = false
    router.push(link.to)
  } else {
    go(link.id)
  }
}

watch(
  () => route.fullPath,
  () => {
    if (route.name === 'directions') active.value = 'directions'
    onScroll()
  }
)

onMounted(() => {
  onScroll()
  window.addEventListener('scroll', onScroll, { passive: true })
})
onUnmounted(() => window.removeEventListener('scroll', onScroll))
</script>

<template>
  <header class="nav" :class="{ scrolled }">
    <div class="container nav-inner">
      <button class="brand" @click="go('home')" aria-label="返回首页">
        <span class="brand-mark">
          <img src="/images/t9.gif" alt="第九网络组" />
        </span>
        <span class="brand-text">
          <strong>第九网络组</strong>
          <em>THE 9TH NETWORK</em>
        </span>
      </button>

      <nav class="links" :class="{ open: menuOpen }">
        <button
          v-for="link in navLinks"
          :key="link.id"
          class="link"
          :class="{ active: active === link.id }"
          @click="goLink(link)"
        >
          {{ link.label }}
        </button>
        <button class="cta" @click="go('join')">
          <BaseIcon name="spark" :size="16" />
          加入我们
        </button>
      </nav>

      <button class="toggle" @click="menuOpen = !menuOpen" aria-label="切换菜单">
        <BaseIcon :name="menuOpen ? 'close' : 'menu'" :size="22" />
      </button>
    </div>
  </header>
</template>

<style scoped>
.nav {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  z-index: 100;
  transition: background 0.3s ease, border-color 0.3s ease, backdrop-filter 0.3s ease;
  border-bottom: 1px solid transparent;
}

.nav.scrolled {
  background: rgba(5, 7, 15, 0.82);
  backdrop-filter: blur(16px);
  border-bottom-color: var(--border);
}

.nav-inner {
  height: 72px;
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 20px;
}

.brand {
  display: flex;
  align-items: center;
  gap: 12px;
  background: none;
  border: 0;
  color: var(--text);
}

.brand-mark {
  width: 42px;
  height: 42px;
  border-radius: 12px;
  display: grid;
  place-items: center;
  background: linear-gradient(135deg, rgba(34, 211, 238, 0.2), rgba(139, 92, 246, 0.25));
  border: 1px solid var(--border-strong);
  box-shadow: 0 0 22px rgba(34, 211, 238, 0.25);
}

.brand-mark img {
  width: 26px;
  height: 26px;
  border-radius: 6px;
}

.brand-text {
  display: flex;
  flex-direction: column;
  align-items: flex-start;
  line-height: 1.1;
}

.brand-text strong {
  font-size: 16px;
  letter-spacing: 0.04em;
}

.brand-text em {
  font-family: var(--mono);
  font-style: normal;
  font-size: 9px;
  letter-spacing: 0.22em;
  color: var(--cyan);
}

.links {
  display: flex;
  align-items: center;
  gap: 4px;
}

.link {
  position: relative;
  background: none;
  border: 0;
  color: var(--text-dim);
  font-size: 14px;
  padding: 10px 16px;
  border-radius: 10px;
  transition: color 0.2s ease, background 0.2s ease;
}

.link:hover {
  color: var(--text);
  background: rgba(255, 255, 255, 0.04);
}

.link.active {
  color: var(--cyan);
}

.link.active::after {
  content: '';
  position: absolute;
  left: 50%;
  bottom: 4px;
  transform: translateX(-50%);
  width: 18px;
  height: 2px;
  border-radius: 2px;
  background: var(--cyan);
  box-shadow: 0 0 10px var(--cyan);
}

.cta {
  display: inline-flex;
  align-items: center;
  gap: 8px;
  margin-left: 12px;
  padding: 10px 18px;
  border-radius: 10px;
  border: 1px solid var(--border-strong);
  background: linear-gradient(135deg, rgba(34, 211, 238, 0.16), rgba(139, 92, 246, 0.16));
  color: var(--text);
  font-size: 14px;
  font-weight: 600;
  transition: transform 0.2s ease, box-shadow 0.2s ease;
}

.cta:hover {
  transform: translateY(-1px);
  box-shadow: 0 8px 26px rgba(34, 211, 238, 0.28);
}

.toggle {
  display: none;
  background: none;
  border: 1px solid var(--border);
  border-radius: 10px;
  padding: 8px;
  color: var(--text);
}

@media (max-width: 900px) {
  .toggle {
    display: grid;
    place-items: center;
  }

  .links {
    position: absolute;
    top: 72px;
    left: 4vw;
    right: 4vw;
    flex-direction: column;
    align-items: stretch;
    padding: 16px;
    gap: 6px;
    background: rgba(8, 12, 24, 0.96);
    border: 1px solid var(--border);
    border-radius: var(--radius);
    backdrop-filter: blur(18px);
    opacity: 0;
    transform: translateY(-12px);
    pointer-events: none;
    transition: opacity 0.25s ease, transform 0.25s ease;
  }

  .links.open {
    opacity: 1;
    transform: translateY(0);
    pointer-events: auto;
  }

  .link {
    text-align: left;
  }

  .cta {
    margin: 8px 0 0;
    justify-content: center;
  }
}
</style>
