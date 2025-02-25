<script setup>
import Header from '@/components/UI/PublicHeader.vue'
import Footer from '@/components/UI/Footer.vue'
import NotificationModal from '@/components/NotificationModal.vue'
import { useNotificationStore } from '@/stores/notification.ts'

const notificationStore = useNotificationStore()
</script>

<template>
  <div>
    <Header />
    <router-view />
    <Footer />
    <NotificationModal
      :isOpen="notificationStore.isOpen"
      :message="notificationStore.message"
      :type="notificationStore.type"
      @close="notificationStore.closeNotification"
    />
  </div>
</template>
