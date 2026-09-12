<script setup>
import { nextTick } from 'vue'
import { useRouter } from 'vue-router'
import { directionsConfig } from '../config/directions'
import BaseIcon from '../components/BaseIcon.vue'

const router = useRouter()
const { section, directions, ctf } = directionsConfig

const goHomeSection = async (id) => {
  await router.push('/')
  await nextTick()
  setTimeout(() => document.getElementById(id)?.scrollIntoView({ behavior: 'smooth' }), 80)
}
</script>

<template>
  <div class="directions-page">
    <section class="page-hero">
      <div class="container">
        <span class="eyebrow">{{ section.eyebrow }}</span>
        <h1 class="page-title">
          {{ section.title }}<span class="grad">{{ section.highlight }}</span>
        </h1>
        <p class="page-sub">{{ section.subtitle }}</p>
        <button class="back-home" @click="goHomeSection('home')">
          <BaseIcon name="chevron-left" :size="16" />
          返回首页
        </button>
      </div>
    </section>

    <section class="section directions">
      <div class="container">
        <div class="dir-grid">
          <article
            v-for="(dir, i) in directions"
            :key="dir.id"
            class="dir-card glass"
            v-reveal="{ delay: i * 120 }"
          >
            <span class="dir-glow" />
            <div class="dir-top">
              <span class="dir-icon">
                <BaseIcon :name="dir.icon" :size="28" />
              </span>
              <div>
                <h3>{{ dir.name }}</h3>
                <p class="dir-tagline">{{ dir.tagline }}</p>
              </div>
            </div>
            <p class="dir-desc">{{ dir.desc }}</p>
            <ul class="dir-list">
              <li v-for="item in dir.highlights" :key="item">
                <BaseIcon name="check" :size="15" />
                <span>{{ item }}</span>
              </li>
            </ul>
            <div class="dir-stack">
              <span v-for="tech in dir.stack" :key="tech">{{ tech }}</span>
            </div>
          </article>
        </div>
      </div>
    </section>

    <section v-if="ctf.enabled" class="section ctf">
      <div class="container">
        <div class="section-head">
          <span class="eyebrow">{{ ctf.eyebrow }}</span>
          <h2 class="section-title">
            <span class="grad">{{ ctf.title }}</span>
          </h2>
          <p class="section-sub">{{ ctf.subtitle }}</p>
        </div>

        <div class="rank-banner glass" v-reveal>
          <div class="rank-left">
            <span class="rank-tag">{{ ctf.bestRank.label }}</span>
            <div class="rank-value">
              {{ ctf.bestRank.value }}<em>{{ ctf.bestRank.unit }}</em>
            </div>
          </div>
          <div class="rank-right">
            <h3>{{ ctf.bestRank.title }}</h3>
            <p>{{ ctf.bestRank.note }}</p>
          </div>
          <BaseIcon class="rank-icon" name="trophy" :size="120" />
        </div>

        <div class="ctf-grid">
          <div class="team-card glass" v-reveal>
            <div class="team-head">
              <span class="team-logo">
                <BaseIcon name="shield" :size="24" />
              </span>
              <div>
                <h3>{{ ctf.team.name }}</h3>
                <p>{{ ctf.team.fullName }}</p>
              </div>
            </div>
            <p class="team-slogan">“{{ ctf.team.slogan }}”</p>
            <ul class="team-fields">
              <li v-for="field in ctf.team.fields" :key="field.label">
                <span>{{ field.label }}</span>
                <strong>{{ field.value }}</strong>
              </li>
            </ul>
            <a
              v-if="ctf.team.url"
              class="team-link"
              :href="ctf.team.url"
              target="_blank"
              rel="noopener noreferrer"
            >
              <BaseIcon name="external" :size="16" />
              CTFTime 战队主页
            </a>
          </div>

          <div class="members" v-reveal="{ delay: 120 }">
            <div class="members-head">
              <h3>
                <BaseIcon name="users" :size="18" />
                战队队员
              </h3>
              <span>共 {{ ctf.members.length }} 人</span>
            </div>
            <div class="member-grid">
              <div v-for="m in ctf.members" :key="m.handle" class="member glass">
                <span class="avatar">{{ m.name.slice(0, 1) }}</span>
                <div class="member-info">
                  <div class="member-line">
                    <strong>{{ m.name }}</strong>
                    <span class="handle">{{ m.handle }}</span>
                  </div>
                  <div class="member-line sub">
                    <span class="role">{{ m.role }}</span>
                    <span class="direction">{{ m.direction }}</span>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>

        <div class="records" v-reveal>
          <div class="records-head">
            <h3>
              <BaseIcon name="trophy" :size="18" />
              比赛记录与成绩
            </h3>
          </div>
          <div class="table-wrap glass">
            <table>
              <thead>
                <tr>
                  <th>时间</th>
                  <th>赛事</th>
                  <th>类型</th>
                  <th>成绩</th>
                  <th>备注</th>
                </tr>
              </thead>
              <tbody>
                <tr v-for="(c, i) in ctf.competitions" :key="i" :class="{ best: c.highlight }">
                  <td class="date">{{ c.date }}</td>
                  <td class="name">{{ c.name }}</td>
                  <td><span class="type">{{ c.type }}</span></td>
                  <td class="result">{{ c.result }}</td>
                  <td class="note">{{ c.note || '—' }}</td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>

        <div class="ctf-cta" v-reveal>
          <div>
            <h3>想一起打 CTF？</h3>
            <p>无论你是零基础还是老选手，欢迎加入战队一起训练、一起参赛。</p>
          </div>
          <button class="cta-btn" @click="goHomeSection('join')">
            加入我们
            <BaseIcon name="arrow-right" :size="16" />
          </button>
        </div>
      </div>
    </section>
  </div>
</template>

<style scoped>
.page-hero {
  position: relative;
  padding: 150px 0 40px;
  text-align: center;
  overflow: hidden;
}

.page-hero::before {
  content: '';
  position: absolute;
  top: -40%;
  left: 50%;
  width: 720px;
  height: 720px;
  transform: translateX(-50%);
  background: radial-gradient(circle, rgba(34, 211, 238, 0.16), transparent 62%);
  z-index: -1;
}

.page-title {
  margin-top: 20px;
  font-size: clamp(34px, 6vw, 60px);
  font-weight: 900;
  letter-spacing: -0.02em;
  line-height: 1.1;
  color: var(--text);
}

.page-title .grad {
  background: linear-gradient(120deg, var(--cyan) 0%, var(--cyan-soft) 40%, var(--violet) 100%);
  -webkit-background-clip: text;
  background-clip: text;
  color: transparent;
}

.page-sub {
  max-width: 720px;
  margin: 18px auto 0;
  color: var(--text-dim);
  font-size: 15px;
}

.back-home {
  display: inline-flex;
  align-items: center;
  gap: 6px;
  margin-top: 24px;
  padding: 9px 18px;
  border-radius: 999px;
  border: 1px solid var(--border);
  background: rgba(255, 255, 255, 0.02);
  color: var(--text-dim);
  font-size: 13px;
  transition: all 0.2s ease;
}

.back-home:hover {
  color: var(--cyan);
  border-color: var(--border-strong);
}

.dir-grid {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: 24px;
}

.dir-card {
  position: relative;
  padding: 34px 32px;
  overflow: hidden;
  transition: transform 0.3s ease, border-color 0.3s ease, box-shadow 0.3s ease;
}

.dir-card:hover {
  transform: translateY(-6px);
  border-color: var(--border-strong);
  box-shadow: 0 26px 60px rgba(0, 0, 0, 0.45);
}

.dir-glow {
  position: absolute;
  top: -60px;
  right: -60px;
  width: 220px;
  height: 220px;
  border-radius: 50%;
  background: radial-gradient(circle, rgba(34, 211, 238, 0.18), transparent 68%);
  pointer-events: none;
}

.dir-top {
  display: flex;
  align-items: center;
  gap: 16px;
  margin-bottom: 20px;
}

.dir-icon {
  display: grid;
  place-items: center;
  width: 58px;
  height: 58px;
  border-radius: 16px;
  color: var(--cyan);
  background: linear-gradient(135deg, rgba(34, 211, 238, 0.16), rgba(139, 92, 246, 0.16));
  border: 1px solid var(--border-strong);
  flex-shrink: 0;
}

.dir-top h3 {
  font-size: 21px;
  color: var(--text);
}

.dir-tagline {
  margin-top: 4px;
  font-family: var(--mono);
  font-size: 12px;
  letter-spacing: 0.06em;
  color: var(--cyan-soft);
}

.dir-desc {
  color: var(--text-dim);
  font-size: 14.5px;
  line-height: 1.8;
  margin-bottom: 20px;
}

.dir-list {
  list-style: none;
  display: flex;
  flex-direction: column;
  gap: 11px;
  margin-bottom: 22px;
}

.dir-list li {
  display: flex;
  gap: 10px;
  align-items: flex-start;
  color: var(--text-dim);
  font-size: 14px;
}

.dir-list li :deep(.icon) {
  margin-top: 4px;
  color: var(--cyan);
  flex-shrink: 0;
}

.dir-stack {
  display: flex;
  flex-wrap: wrap;
  gap: 8px;
  padding-top: 18px;
  border-top: 1px solid var(--border);
}

.dir-stack span {
  font-family: var(--mono);
  font-size: 11.5px;
  padding: 4px 11px;
  border-radius: 7px;
  color: var(--text-dim);
  background: rgba(255, 255, 255, 0.04);
  border: 1px solid var(--border);
}

.rank-banner {
  position: relative;
  display: flex;
  align-items: center;
  gap: 36px;
  padding: 34px 40px;
  margin-bottom: 24px;
  overflow: hidden;
  border-color: var(--border-strong);
}

.rank-left {
  padding-right: 36px;
  border-right: 1px solid var(--border);
  flex-shrink: 0;
}

.rank-tag {
  display: block;
  font-family: var(--mono);
  font-size: 11px;
  letter-spacing: 0.2em;
  color: var(--cyan);
  margin-bottom: 8px;
}

.rank-value {
  font-family: var(--mono);
  font-size: clamp(48px, 8vw, 76px);
  font-weight: 800;
  line-height: 1;
  background: linear-gradient(120deg, var(--cyan), var(--violet));
  -webkit-background-clip: text;
  background-clip: text;
  color: transparent;
}

.rank-value em {
  font-size: 0.28em;
  font-style: normal;
  margin-left: 6px;
  color: var(--text-dim);
  -webkit-text-fill-color: var(--text-dim);
}

.rank-right h3 {
  font-size: clamp(18px, 2.4vw, 24px);
  color: var(--text);
  margin-bottom: 8px;
}

.rank-right p {
  color: var(--text-dim);
  font-size: 14px;
}

.rank-icon {
  position: absolute;
  right: 28px;
  bottom: -26px;
  color: rgba(34, 211, 238, 0.12);
}

.ctf-grid {
  display: grid;
  grid-template-columns: 0.85fr 1.15fr;
  gap: 24px;
  margin-bottom: 24px;
}

.team-card {
  padding: 30px 28px;
}

.team-head {
  display: flex;
  align-items: center;
  gap: 16px;
  margin-bottom: 18px;
}

.team-logo {
  display: grid;
  place-items: center;
  width: 52px;
  height: 52px;
  border-radius: 14px;
  color: var(--cyan);
  background: rgba(34, 211, 238, 0.1);
  border: 1px solid var(--border-strong);
}

.team-head h3 {
  font-size: 20px;
  color: var(--text);
}

.team-head p {
  font-size: 13px;
  color: var(--text-mute);
}

.team-slogan {
  font-size: 13.5px;
  color: var(--cyan-soft);
  font-style: italic;
  margin-bottom: 20px;
}

.team-fields {
  list-style: none;
  display: flex;
  flex-direction: column;
}

.team-fields li {
  display: flex;
  justify-content: space-between;
  gap: 16px;
  padding: 12px 0;
  border-bottom: 1px dashed var(--border);
  font-size: 13.5px;
}

.team-fields li:last-child {
  border-bottom: 0;
}

.team-fields span {
  color: var(--text-mute);
  flex-shrink: 0;
}

.team-fields strong {
  color: var(--text);
  font-weight: 600;
  text-align: right;
}

.team-link {
  display: inline-flex;
  align-items: center;
  gap: 8px;
  margin-top: 20px;
  padding: 10px 18px;
  border-radius: 10px;
  border: 1px solid var(--border-strong);
  background: rgba(34, 211, 238, 0.08);
  color: var(--cyan-soft);
  font-size: 13.5px;
  font-weight: 600;
  transition: background 0.2s ease, color 0.2s ease, transform 0.2s ease, box-shadow 0.2s ease;
}

.team-link:hover {
  background: rgba(34, 211, 238, 0.16);
  color: var(--text);
  transform: translateY(-1px);
  box-shadow: 0 10px 26px rgba(34, 211, 238, 0.25);
}

.members {
  padding: 0;
}

.members-head {
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-bottom: 16px;
}

.members-head h3 {
  display: flex;
  align-items: center;
  gap: 8px;
  font-size: 18px;
  color: var(--text);
}

.members-head h3 :deep(.icon) {
  color: var(--cyan);
}

.members-head > span {
  font-family: var(--mono);
  font-size: 12px;
  color: var(--text-mute);
}

.member-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(230px, 1fr));
  gap: 14px;
}

.member {
  display: flex;
  align-items: center;
  gap: 14px;
  padding: 16px;
  transition: transform 0.24s ease, border-color 0.24s ease;
}

.member:hover {
  transform: translateY(-3px);
  border-color: var(--border-strong);
}

.avatar {
  display: grid;
  place-items: center;
  width: 44px;
  height: 44px;
  border-radius: 12px;
  font-size: 18px;
  font-weight: 700;
  color: #04121a;
  background: linear-gradient(135deg, var(--cyan), var(--cyan-soft));
  flex-shrink: 0;
}

.member-info {
  min-width: 0;
  display: flex;
  flex-direction: column;
  gap: 5px;
}

.member-line {
  display: flex;
  align-items: baseline;
  gap: 8px;
}

.member-line strong {
  font-size: 15px;
  color: var(--text);
}

.handle {
  font-family: var(--mono);
  font-size: 11.5px;
  color: var(--text-mute);
  overflow: hidden;
  text-overflow: ellipsis;
}

.member-line.sub {
  font-size: 11.5px;
}

.role {
  padding: 2px 8px;
  border-radius: 999px;
  color: var(--cyan-soft);
  background: rgba(34, 211, 238, 0.1);
  border: 1px solid var(--border);
}

.direction {
  color: var(--text-dim);
}

.records-head h3 {
  display: flex;
  align-items: center;
  gap: 8px;
  font-size: 18px;
  color: var(--text);
  margin-bottom: 16px;
}

.records-head h3 :deep(.icon) {
  color: var(--amber);
}

.table-wrap {
  overflow-x: auto;
  padding: 4px;
}

table {
  width: 100%;
  border-collapse: collapse;
  min-width: 640px;
}

th,
td {
  text-align: left;
  padding: 15px 18px;
  font-size: 14px;
  white-space: nowrap;
}

thead th {
  font-family: var(--mono);
  font-size: 11px;
  letter-spacing: 0.14em;
  text-transform: uppercase;
  color: var(--text-mute);
  border-bottom: 1px solid var(--border);
}

tbody tr {
  border-bottom: 1px solid rgba(110, 231, 249, 0.08);
  transition: background 0.2s ease;
}

tbody tr:last-child {
  border-bottom: 0;
}

tbody tr:hover {
  background: rgba(34, 211, 238, 0.05);
}

tbody tr.best {
  background: linear-gradient(90deg, rgba(34, 211, 238, 0.12), transparent 70%);
}

.date {
  font-family: var(--mono);
  color: var(--text-mute);
}

.name {
  color: var(--text);
  font-weight: 600;
}

.type {
  font-size: 11.5px;
  padding: 3px 10px;
  border-radius: 999px;
  color: var(--text-dim);
  background: rgba(255, 255, 255, 0.04);
  border: 1px solid var(--border);
}

.result {
  color: var(--cyan-soft);
  font-weight: 600;
}

tr.best .result {
  color: var(--cyan);
}

.note {
  color: var(--text-mute);
}

.ctf-cta {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 20px;
  margin-top: 32px;
  padding: 30px 34px;
  border-radius: var(--radius);
  border: 1px solid var(--border-strong);
  background: linear-gradient(120deg, rgba(34, 211, 238, 0.1), rgba(139, 92, 246, 0.12));
}

.ctf-cta h3 {
  font-size: 20px;
  color: var(--text);
  margin-bottom: 6px;
}

.ctf-cta p {
  color: var(--text-dim);
  font-size: 14px;
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
  .dir-grid,
  .ctf-grid {
    grid-template-columns: 1fr;
  }
}

@media (max-width: 640px) {
  .rank-banner {
    flex-direction: column;
    align-items: flex-start;
    gap: 20px;
    padding: 26px 24px;
  }
  .rank-left {
    padding-right: 0;
    padding-bottom: 16px;
    border-right: 0;
    border-bottom: 1px solid var(--border);
    width: 100%;
  }
  .ctf-cta {
    flex-direction: column;
    align-items: flex-start;
  }
}
</style>
