<script setup>
import NavBar from './components/NavBar.vue'
import SiteFooter from './components/SiteFooter.vue'
import BackToTop from './components/BackToTop.vue'
</script>

<template>
  <NavBar />
  <main>
    <router-view />
  </main>
  <SiteFooter />
  <BackToTop />
</template>
