<script setup>
import Header from '@/components/UI/PublicHeader.vue'
import Footer from '@/components/UI/Footer.vue'
import NotificationModal from '@/components/NotificationModal.vue'
import { useNotificationStore } from '@/stores/notification.ts'
import { onMounted } from 'vue'

const notificationStore = useNotificationStore()

onMounted(() => {
  notificationStore.showNotification('Welcome to the Public Layout!', 'info')
})
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
