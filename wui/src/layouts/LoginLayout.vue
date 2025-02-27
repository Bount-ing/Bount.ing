<script setup>
import LandingHeader from '@/components/UI/PublicHeader.vue'
import Footer from '@/components/UI/Footer.vue'
import NotificationModal from '@/components/NotificationModal.vue'
import { useNotificationStore } from '@/stores/notification.ts'

const notificationStore = useNotificationStore()
</script>

<template>
  <div>
    <LandingHeader />
    <div class="pt-8">
      <router-view />
    </div>
    <Footer />
    <NotificationModal
      :isOpen="notificationStore.isOpen"
      :message="notificationStore.message"
      :type="notificationStore.type"
      @close="notificationStore.closeNotification"
    />
  </div>
</template>

<style scoped></style>
