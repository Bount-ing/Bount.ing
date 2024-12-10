<script setup lang="ts">
import { RouterLink, RouterView } from 'vue-router'
import Footer from './components/UI/Footer.vue'
</script>

<template>
  <div class="flex flex-col fixed  h-screen w-screen bg-secondary-dark text-blue min-w-screen min-h-screen p-0 m-0 top-0 left-0 center overflow-scroll scrollbar-hide">
    <main class="flex-grow">
      <RouterView />
    </main>
  </div>
</template>

<style>

</style>