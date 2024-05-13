<template>
  <div class="bg-secondary text-sans min-h-screen flex flex-col p-8">
    <NavBar />
    <Nuxt /> <!-- This tag renders the page content -->
    <Footer />
  </div>
</template>

<script>
import NavBar from '~/components/NavBar.vue'
import Footer from '~/components/Footer.vue'
export default {
  // You can add JavaScript here if needed
}
</script>

<style>

</style>
