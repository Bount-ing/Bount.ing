<script setup>
import LandingHeader from '@/components/UI/PublicHeader.vue'
import NotificationModal from '@/components/NotificationModal.vue'
import { useNotificationStore } from '@/stores/notification.ts'

const notificationStore = useNotificationStore()
</script>

<template>
  <div class="fullscreen">
    <section class="user-form-part">
      <div class="user-form-banner">
        <LandingHeader />
        <div class="user-form-content"></div>
      </div>
      <div class="user-form-category">
        <RouterView />
      </div>
    </section>
    <NotificationModal
      :isOpen="notificationStore.isOpen"
      :message="notificationStore.message"
      :type="notificationStore.type"
      @close="notificationStore.closeNotification"
    />
  </div>
</template>

<style scoped></style>
