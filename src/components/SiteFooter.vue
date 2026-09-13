<script setup>
import { nextTick } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { config } from '../config'

const route = useRoute()
const router = useRouter()

const brand = config.site.brand
const links = config.site.footerLinks

const go = async (to) => {
  if (to.startsWith('/')) {
    router.push(to)
    return
  }
  if (route.name !== 'home') {
    await router.push('/')
    await nextTick()
    setTimeout(() => document.getElementById(to)?.scrollIntoView({ behavior: 'smooth' }), 80)
  } else {
    document.getElementById(to)?.scrollIntoView({ behavior: 'smooth' })
  }
}
</script>

<template>
  <footer class="footer">
    <div class="container footer-inner">
      <div class="brand-col">
        <div class="brand">
          <img :src="brand.logo" :alt="brand.name" />
          <div>
            <strong>{{ brand.name }}</strong>
            <em>{{ brand.fullEnName }}</em>
          </div>
        </div>
        <p>{{ brand.slogan }}</p>
        <p class="school">{{ brand.school }}</p>
      </div>

      <div class="link-cols">
        <div v-for="group in links" :key="group.title" class="link-col">
          <h4>{{ group.title }}</h4>
          <ul>
            <li v-for="item in group.items" :key="item.label">
              <button @click="go(item.to)">
                {{ item.label }}
              </button>
            </li>
          </ul>
        </div>
      </div>
    </div>

    <div class="container footer-bottom">
      <p>Copyright &copy; {{ new Date().getFullYear() }} {{ brand.name }} · {{ brand.fullEnName }}</p>
      <p class="mono">Designed &amp; Built with Vue 3</p>
    </div>
  </footer>
</template>

<style scoped>
.footer {
  position: relative;
  margin-top: 40px;
  border-top: 1px solid var(--border);
  background: linear-gradient(180deg, rgba(26, 36, 64, 0.35), rgba(18, 26, 45, 0.88));
}

.footer-inner {
  display: grid;
  grid-template-columns: 1.1fr 1.4fr;
  gap: 48px;
  padding: 64px 0 40px;
}

.brand {
  display: flex;
  align-items: center;
  gap: 14px;
  margin-bottom: 18px;
}

.brand img {
  width: 44px;
  height: 44px;
  border-radius: 12px;
  border: 1px solid var(--border-strong);
  padding: 6px;
  background: rgba(34, 211, 238, 0.06);
}

.brand strong {
  display: block;
  font-size: 17px;
  color: var(--text);
}

.brand em {
  font-family: var(--mono);
  font-style: normal;
  font-size: 9px;
  letter-spacing: 0.22em;
  color: var(--cyan);
}

.brand-col > p {
  color: var(--text-dim);
  font-size: 14px;
}

.school {
  margin-top: 6px;
  color: var(--text-mute) !important;
  font-size: 13px !important;
}

.link-cols {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 24px;
}

.link-col h4 {
  font-size: 13px;
  letter-spacing: 0.12em;
  font-family: var(--mono);
  color: var(--cyan);
  margin-bottom: 16px;
  text-transform: uppercase;
}

.link-col ul {
  list-style: none;
  display: flex;
  flex-direction: column;
  gap: 10px;
}

.link-col button {
  background: none;
  border: 0;
  color: var(--text-dim);
  font-size: 14px;
  padding: 0;
  transition: color 0.2s ease, transform 0.2s ease;
}

.link-col button:hover {
  color: var(--cyan-soft);
  transform: translateX(3px);
}

.footer-bottom {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 12px;
  padding: 20px 0 32px;
  border-top: 1px solid var(--border);
  color: var(--text-mute);
  font-size: 13px;
}

.mono {
  font-family: var(--mono);
  font-size: 11px;
  letter-spacing: 0.08em;
}

@media (max-width: 820px) {
  .footer-inner {
    grid-template-columns: 1fr;
    gap: 36px;
    padding-top: 48px;
  }
  .footer-bottom {
    flex-direction: column;
    align-items: flex-start;
  }
}

@media (max-width: 480px) {
  .link-cols {
    grid-template-columns: 1fr 1fr;
  }
}
</style>
