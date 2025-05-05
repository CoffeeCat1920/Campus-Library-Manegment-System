<script setup lang="ts">
import { RouterLink, RouterView } from 'vue-router'
import HomeView from './views/Home.vue'
import Navbar from './components/Navbar.vue'
</script>

<template>
  <header>
    <Navbar />
  </header>

  <RouterView />
</template>
