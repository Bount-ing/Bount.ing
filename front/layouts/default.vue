<template>
  <div class="bg-gradient-to-tr from-primary to-secondary  text-sans min-h-screen flex flex-col  min-w-screen">
    <NavBar />
    <div class="flex-grow pt-8">
      <Nuxt /> <!-- This tag renders the page content -->
    </div>
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
