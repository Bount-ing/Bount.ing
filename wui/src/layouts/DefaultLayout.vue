<script setup>
import Header from '@/components/UI/LoggedHeader.vue'
import Footer from '@/components/UI/Footer.vue'
import NotificationModal from '@/components/NotificationModal.vue'
import { useNotificationStore } from '@/stores/notification.ts'

const notificationStore = useNotificationStore()



</script>

<template>
  <div>
    <Header />
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
