<script setup>
import Header from '@/components/UI/LoggedHeader.vue'
import Footer from '@/components/UI/Footer.vue'
import ErrorModal from '@/components/ErrorModal.vue'
import { useErrorStore } from '@/stores/errors'

const errorStore = useErrorStore()
</script>

<template>
  <div>
    <Header />
    <div class="pt-8">
      <router-view />
    </div>
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
