<script setup>
import Header from '@/components/UI/PublicHeader.vue'
import Footer from '@/components/UI/Footer.vue'
import ErrorModal from '@/components/ErrorModal.vue'
import { useErrorStore } from '@/stores/errors'

const errorStore = useErrorStore()
</script>

<template>
  <div>
    <Header />

    <router-view />

    <Footer />
    <!-- Global Error Modal -->
    <ErrorModal
      :isOpen="errorStore.isOpen"
      :message="errorStore.message"
      :isError="errorStore.isError"
      @close="errorStore.closeError"
    />
  </div>
</template>
