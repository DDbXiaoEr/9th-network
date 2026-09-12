import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import './style.css'

const reveal = {
  mounted(el, binding) {
    el.classList.add('reveal')
    if (binding.value?.delay) {
      el.style.transitionDelay = `${binding.value.delay}ms`
    }
    const io = new IntersectionObserver(
      (entries) => {
        entries.forEach((entry) => {
          if (entry.isIntersecting) {
            entry.target.classList.add('is-visible')
            io.unobserve(entry.target)
          }
        })
      },
      { threshold: 0.1, rootMargin: '0px 0px -40px 0px' }
    )
    io.observe(el)
  }
}

createApp(App).use(router).directive('reveal', reveal).mount('#app')
