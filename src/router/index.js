import { createRouter, createWebHashHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import DirectionsView from '../views/DirectionsView.vue'

const router = createRouter({
  history: createWebHashHistory(),
  routes: [
    { path: '/', name: 'home', component: HomeView, meta: { title: '第九网络组' } },
    {
      path: '/directions',
      name: 'directions',
      component: DirectionsView,
      meta: { title: '兴趣方向 · 第九网络组' }
    },
    { path: '/:pathMatch(.*)*', redirect: '/' }
  ],
  scrollBehavior(to, from, savedPosition) {
    if (savedPosition) return savedPosition
    return { top: 0 }
  }
})

router.afterEach((to) => {
  document.title = to.meta.title || '第九网络组 | The 9th Network Team'
})

export default router
