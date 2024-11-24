<script setup lang="ts">
import { RouterLink, RouterView } from 'vue-router'
import Header from './components/Header.vue'
import Footer from './components/Footer.vue'
</script>

<template>
  <div class="flex flex-col fixed py-14 h-screen w-screen  text-blue min-w-screen min-h-screen p-0 m-0 top-0 left-0 center overflow-scroll scrollbar-hide">
    <Header />
    <main class="flex-grow">
      <RouterView />
    </main>
    <Footer />
  </div>
</template>

<style>

</style>